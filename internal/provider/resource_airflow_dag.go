package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAirflowDAG() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceAirflowDAGCreate,
		ReadContext:   resourceAirflowDAGRead,
		UpdateContext: resourceAirflowDAGUpdate,
		DeleteContext: resourceAirflowDAGDelete,
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
						"filename": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"code": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceAirflowDAGCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "AirflowDAG"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0]

	resource.Spec = spec.(map[string]interface{})

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
	d.Set("spec", []interface{}{ret.Spec})

	return nil
}

func resourceAirflowDAGRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "AirflowDAG"
	name := d.Id()

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
	d.Set("spec", []interface{}{ret.Spec})

	return nil
}

func resourceAirflowDAGUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "AirflowDAG"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").([]interface{})[0]

	resource.Spec = spec.(map[string]interface{})

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
	d.Set("spec", []interface{}{ret.Spec})

	return nil
}

func resourceAirflowDAGDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "AirflowDAG"
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
