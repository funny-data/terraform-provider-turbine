package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIcebergSinkJob() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceIcebergSinkJobCreate,
		ReadContext:   resourceIcebergSinkJobRead,
		UpdateContext: resourceIcebergSinkJobUpdate,
		DeleteContext: resourceIcebergSinkJobDelete,
		Importer:      &schema.ResourceImporter{},

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
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"table": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"table": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"source": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"format": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"topic": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"starting_offsets": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"target_interval": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"event_time_field": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceIcebergSinkJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "IcebergSinkJob"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0].(map[string]interface{})
	spec["table"] = d.Get("spec.0.table.0")
	spec["source"] = d.Get("spec.0.source.0")

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
	d.Set("spec", []interface{}{
		map[string]interface{}{
			"enabled":          spec["enabled"],
			"table":            []interface{}{spec["table"]},
			"source":           []interface{}{spec["source"]},
			"target_interval":  spec["target_interval"],
			"event_time_field": spec["event_time_field"],
		},
	})

	return nil
}

func resourceIcebergSinkJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "IcebergSinkJob"
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
	d.Set("spec", []interface{}{
		map[string]interface{}{
			"enabled":          spec["enabled"],
			"table":            []interface{}{spec["table"]},
			"source":           []interface{}{spec["source"]},
			"target_interval":  spec["target_interval"],
			"event_time_field": spec["event_time_field"],
		},
	})

	return nil
}

func resourceIcebergSinkJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "IcebergSinkJob"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0].(map[string]interface{})
	spec["table"] = d.Get("spec.0.table.0")
	spec["source"] = d.Get("spec.0.source.0")

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
	d.Set("spec", []interface{}{
		map[string]interface{}{
			"enabled":          spec["enabled"],
			"table":            []interface{}{spec["table"]},
			"source":           []interface{}{spec["source"]},
			"target_interval":  spec["target_interval"],
			"event_time_field": spec["event_time_field"],
		},
	})

	return nil
}

func resourceIcebergSinkJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "IcebergSinkJob"
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
