package provider

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAPIClient(t *testing.T) {
	endpoint := os.Getenv("TURBINE_ENDPOINT")
	username := os.Getenv("TURBINE_USERNAME")
	password := os.Getenv("TURBINE_PASSWORD")

	if endpoint == "" {
		t.Skip("env TURBINE_ENDPOINT no provided")
	}

	kind := "Dummy"
	name := "test-" + strconv.FormatInt(time.Now().UnixMilli(), 10)

	ctx := context.Background()

	client := newApiClient("test-client/0", endpoint, username, password)

	t.Run("create", func(t *testing.T) {
		res := Resource{}
		res.Kind = kind
		res.Metadata.Name = name
		res.Spec = map[string]interface{}{"arg": "foo"}

		ret, err := client.Create(ctx, &res)
		assert.NoError(t, err)
		if assert.NotEmpty(t, ret) {
			assert.Equal(t, name, ret.Metadata.Name)
			assert.Equal(t, res.Spec, ret.Spec)
		}
	})

	t.Run("retrieve", func(t *testing.T) {
		ret, err := client.Retrieve(ctx, kind, name)
		assert.NoError(t, err)
		if assert.NotEmpty(t, ret) {
			assert.Equal(t, name, ret.Metadata.Name)
			assert.Equal(t, map[string]interface{}{"arg": "foo"}, ret.Spec)
		}
	})

	t.Run("update", func(t *testing.T) {
		res := Resource{}
		res.Kind = kind
		res.Metadata.Name = name
		res.Spec = map[string]interface{}{"arg": "bar"}

		ret, err := client.Update(ctx, &res)
		assert.NoError(t, err)
		if assert.NotEmpty(t, ret) {
			assert.Equal(t, name, ret.Metadata.Name)
			assert.Equal(t, res.Spec, ret.Spec)
		}
	})

	t.Run("destroy", func(t *testing.T) {
		assert.NoError(t, client.Destroy(ctx, kind, name))

		_, err := client.Retrieve(ctx, kind, name)
		assert.Error(t, errResourceNotFound, err)
	})
}
