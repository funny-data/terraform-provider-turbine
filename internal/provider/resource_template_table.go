package provider

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateTable() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceTemplateTableCreate,
		ReadContext:   resourceTemplateTableRead,
		UpdateContext: resourceTemplateTableUpdate,
		DeleteContext: resourceTemplateTableDelete,

		Schema: map[string]*schema.Schema{
			"metadata": {
				Type:     schema.TypeList,
				MaxItems: 1,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"labels": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"spec": {
				Type:     schema.TypeList,
				MaxItems: 1,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								var oldJSON, newJSON interface{}
								if err := json.Unmarshal([]byte(old), &oldJSON); err != nil {
									return false
								}
								if err := json.Unmarshal([]byte(new), &newJSON); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSON, newJSON)
							},
						},
						"data": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								var oldJSON, newJSON interface{}
								if err := json.Unmarshal([]byte(old), &oldJSON); err != nil {
									return false
								}
								if err := json.Unmarshal([]byte(new), &newJSON); err != nil {
									return false
								}

								return reflect.DeepEqual(oldJSON, newJSON)
							},
						},
					},
				},
			},
		},
	}
}

func resourceTemplateTableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "TemplateTable"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0].(map[string]interface{})

	spec["schema"] = json.RawMessage(spec["schema"].(string))
	spec["data"] = json.RawMessage(spec["data"].(string))

	resource.Spec = spec

	d.SetId(name)

	ret, err := client.Create(ctx, &resource)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("metadata", []interface{}{
		map[string]interface{}{
			"name":   ret.Metadata.Name,
			"labels": ret.Metadata.Labels,
		},
	})
	spec = ret.Spec.(map[string]interface{})
	schemaBytes, err := json.Marshal(spec["schema"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["schema"] = string(schemaBytes)
	dataBytes, err := json.Marshal(spec["data"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["data"] = string(dataBytes)
	d.Set("spec", []interface{}{spec})

	return nil
}

func resourceTemplateTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "TemplateTable"
	name := d.Get("metadata.0.name").(string)

	ret, err := client.Retrieve(ctx, kind, name)
	if err != nil {
		if err == errResourceNotFound {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("metadata", []interface{}{
		map[string]interface{}{
			"name":   ret.Metadata.Name,
			"labels": ret.Metadata.Labels,
		},
	})
	spec := ret.Spec.(map[string]interface{})
	schemaBytes, err := json.Marshal(spec["schema"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["schema"] = string(schemaBytes)
	dataBytes, err := json.Marshal(spec["data"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["data"] = string(dataBytes)
	d.Set("spec", []interface{}{spec})

	return nil
}

func resourceTemplateTableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "TemplateTable"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0].(map[string]interface{})

	spec["schema"] = json.RawMessage(spec["schema"].(string))
	spec["data"] = json.RawMessage(spec["data"].(string))

	resource.Spec = spec

	d.SetId(name)

	ret, err := client.Update(ctx, &resource)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("metadata", []interface{}{
		map[string]interface{}{
			"name":   ret.Metadata.Name,
			"labels": ret.Metadata.Labels,
		},
	})
	spec = ret.Spec.(map[string]interface{})
	schemaBytes, err := json.Marshal(spec["schema"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["schema"] = string(schemaBytes)
	dataBytes, err := json.Marshal(spec["data"])
	if err != nil {
		return diag.FromErr(err)
	}
	spec["data"] = string(dataBytes)
	d.Set("spec", []interface{}{spec})

	return nil
}

func resourceTemplateTableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "TemplateTable"
	name := d.Get("metadata.0.name").(string)

	err := client.Destroy(ctx, kind, name)
	if err != nil {
		if errors.Is(err, errResourceNotFound) {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
