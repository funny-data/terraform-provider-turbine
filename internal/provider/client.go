package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	errResourceNotFound = errors.New("resource not found")
)

type apiClient struct {
	http *http.Client

	userAgent string
	endpoint  string
	username  string
	password  string

	pollInterval time.Duration
	pollTimeout  time.Duration
}

func newApiClient(userAgent, endpoint, username, password string) *apiClient {
	return &apiClient{
		http:         &http.Client{},
		userAgent:    userAgent,
		endpoint:     endpoint,
		username:     username,
		password:     password,
		pollInterval: 1 * time.Second,
		pollTimeout:  5 * time.Minute,
	}
}

func (c *apiClient) Create(ctx context.Context, resource *Resource) (*Resource, error) {
	ret := Resource{}
	err := c.do(ctx, "POST", "v1/resources", resource, http.StatusCreated, &ret)
	if err != nil {
		return nil, fmt.Errorf("create resource: %s", err)
	}
	err = c.waitStatus(ctx, &ret, "LIVE")
	if err != nil {
		return nil, fmt.Errorf("wait resource become LIVE: %s", err)
	}
	return &ret, err
}

func (c *apiClient) Retrieve(ctx context.Context, kind string, name string) (*Resource, error) {
	ret := Resource{}
	err := c.do(ctx, "GET", fmt.Sprintf("v1/resources/%s/%s", kind, name), nil, http.StatusOK, &ret)
	return &ret, err
}

func (c *apiClient) Update(ctx context.Context, resource *Resource) (*Resource, error) {
	ret := Resource{}
	err := c.do(ctx, "PUT", fmt.Sprintf("v1/resources/%s/%s", resource.Kind, resource.Metadata.Name), resource, http.StatusOK, &ret)
	if err != nil {
		return nil, fmt.Errorf("update resource: %s", err)
	}
	err = c.waitStatus(ctx, &ret, "LIVE")
	if err != nil {
		return nil, fmt.Errorf("wait resource become LIVE: %s", err)
	}
	return &ret, err
}

func (c *apiClient) Destroy(ctx context.Context, kind string, name string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("v1/resources/%s/%s", kind, name), nil, http.StatusAccepted, nil)
	if err != nil {
		return fmt.Errorf("delete resource: %s", err)
	}

	pollCtx, cancel := context.WithTimeout(ctx, c.pollTimeout)
	defer cancel()
	for {
		time.Sleep(c.pollInterval)

		r, err := c.Retrieve(pollCtx, kind, name)
		if err != nil {
			if err == errResourceNotFound {
				return nil
			}
			if err == context.DeadlineExceeded && pollCtx.Err() != nil && ctx.Err() == nil {
				return fmt.Errorf("timeout while polling resource status: %s", err)
			}
			return fmt.Errorf("poll resource statusp: %s", err)
		}
		if r.Status.State == "FAILURE" {
			return fmt.Errorf("resource failure with: %s", r.Status.Reason)
		}
	}
}

func (c *apiClient) waitStatus(ctx context.Context, resource *Resource, expectedStatus string) error {
	pollCtx, cancel := context.WithTimeout(ctx, c.pollTimeout)
	defer cancel()
	for {
		time.Sleep(c.pollInterval)

		r, err := c.Retrieve(pollCtx, resource.Kind, resource.Metadata.Name)
		if err != nil {
			if err == context.DeadlineExceeded && pollCtx.Err() != nil && ctx.Err() == nil {
				return fmt.Errorf("timeout while polling resource status: %s", err)
			}
			return fmt.Errorf("poll resource statusp: %s", err)
		}
		if r.Status.State == expectedStatus {
			return nil
		}
		if r.Status.State == "FAILURE" {
			return fmt.Errorf("resource failure with: %s", r.Status.Reason)
		}
	}
}

func (c *apiClient) newRequest(ctx context.Context, method string, api string, body io.Reader) (*http.Request, error) {
	url := strings.TrimRight(c.endpoint, "/") + "/" + api
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", c.userAgent)
	if c.username != "" {
		req.SetBasicAuth(c.username, c.password)
	}
	return req, nil
}

func (c *apiClient) do(ctx context.Context, method, api string, args interface{}, expectedStatusCode int, returns interface{}) error {
	var body io.Reader
	if args != nil {
		b := bytes.NewBuffer(nil)
		if err := json.NewEncoder(b).Encode(&args); err != nil {
			return err
		}
		body = b
	}

	req, err := c.newRequest(ctx, method, api, body)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, "sending api request", map[string]interface{}{
		"method":          method,
		"api":             api,
		"args":            args,
		"expected_status": expectedStatusCode,
	})

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, "received api response", map[string]interface{}{
		"method":          method,
		"api":             api,
		"args":            args,
		"expected_status": expectedStatusCode,
		"actual_status":   resp.StatusCode,
	})
	tflog.Trace(ctx, "api response content", map[string]interface{}{
		"content": string(content),
	})

	if resp.StatusCode != expectedStatusCode {
		if resp.StatusCode == http.StatusNotFound {
			return errResourceNotFound
		}
		return fmt.Errorf("unexpected response status (%d != %d): %s", resp.StatusCode, expectedStatusCode, string(content))
	}

	if returns != nil {
		if err := json.Unmarshal(content, &returns); err != nil {
			return err
		}
	}
	return nil
}

type Resource struct {
	Metadata struct {
		Name   string            `json:"name"`
		Labels map[string]string `json:"labels"`
	} `json:"metadata"`
	Kind   string                 `json:"kind"`
	Spec   map[string]interface{} `json:"spec"`
	Status struct {
		State  string `json:"state"`
		Reason string `json:"reason"`
	} `json:"status"`
}
