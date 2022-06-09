package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHudiSinkJob() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceHudiSinkJobCreate,
		ReadContext:   resourceHudiSinkJobRead,
		UpdateContext: resourceHudiSinkJobUpdate,
		DeleteContext: resourceHudiSinkJobDelete,

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
						"delta_streamer_task": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"operation": {
										Type:     schema.TypeString,
										Required: true,
									},
									"interval": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"batch_size_limit": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"batch_records_limit": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"backoff": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"max_backoff": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"default_lag_records": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"default_avg_record_size": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"parallelism": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"extra_properties": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"sync_task": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"interval": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"backoff": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"max_backoff": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"extra_properties": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceHudiSinkJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "HudiSinkJob"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := make(map[string]interface{})
	spec["table"] = d.Get("spec.0.table.0")
	spec["source"] = d.Get("spec.0.source.0")
	spec["delta_streamer_task"] = d.Get("spec.0.delta_streamer_task.0")
	spec["sync_task"] = d.Get("spec.0.sync_task.0")

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
			"table":               []interface{}{spec["table"]},
			"source":              []interface{}{spec["source"]},
			"delta_streamer_task": []interface{}{spec["delta_streamer_task"]},
			"sync_task":           []interface{}{spec["sync_task"]},
		},
	})

	return nil
}

func resourceHudiSinkJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "HudiSinkJob"
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
			"table":               []interface{}{spec["table"]},
			"source":              []interface{}{spec["source"]},
			"delta_streamer_task": []interface{}{spec["delta_streamer_task"]},
			"sync_task":           []interface{}{spec["sync_task"]},
		},
	})

	return nil
}

func resourceHudiSinkJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "HudiSinkJob"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := make(map[string]interface{})
	spec["table"] = d.Get("spec.0.table.0")
	spec["source"] = d.Get("spec.0.source.0")
	spec["delta_streamer_task"] = d.Get("spec.0.delta_streamer_task.0")
	spec["sync_task"] = d.Get("spec.0.sync_task.0")

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
			"table":               []interface{}{spec["table"]},
			"source":              []interface{}{spec["source"]},
			"delta_streamer_task": []interface{}{spec["delta_streamer_task"]},
			"sync_task":           []interface{}{spec["sync_task"]},
		},
	})

	return nil
}

func resourceHudiSinkJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "HudiSinkJob"
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
