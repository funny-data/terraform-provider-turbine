package provider

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIngestMetadata() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceIngestMetadataCreate,
		ReadContext:   resourceIngestMetadataRead,
		UpdateContext: resourceIngestMetadataUpdate,
		DeleteContext: resourceIngestMetadataDelete,

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
	}
}

func resourceIngestMetadataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "IngestMetadata"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").(string)

	resource.Spec = json.RawMessage(spec)

	d.SetId(name)

	ret, err := client.Create(ctx, &resource)
	if err != nil {
		return diag.FromErr(err)
	}

	specBytes, err := json.Marshal(ret.Spec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("metadata.0.name", ret.Metadata.Name)
	d.Set("metadata.0.labels", ret.Metadata.Labels)
	d.Set("spec.0", string(specBytes))

	return nil
}

func resourceIngestMetadataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "IngestMetadata"
	name := d.Get("metadata.0.name").(string)

	ret, err := client.Retrieve(ctx, kind, name)
	if err != nil {
		if err == errResourceNotFound {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	specBytes, err := json.Marshal(ret.Spec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("metadata.0.name", ret.Metadata.Name)
	d.Set("metadata.0.labels", ret.Metadata.Labels)
	d.Set("spec.0", string(specBytes))

	return nil
}

func resourceIngestMetadataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("metadata.0.name").(string)
	labels := make(map[string]string)
	if v, ok := d.GetOk("metadata.0.labels"); ok && v != nil {
		for k, v := range v.(map[string]interface{}) {
			labels[k] = v.(string)
		}
	}

	resource := Resource{}
	resource.Kind = "IngestMetadata"
	resource.Metadata.Name = name
	resource.Metadata.Labels = labels

	spec := d.Get("spec").(string)

	resource.Spec = json.RawMessage(spec)

	d.SetId(name)

	ret, err := client.Update(ctx, &resource)
	if err != nil {
		return diag.FromErr(err)
	}

	specBytes, err := json.Marshal(ret.Spec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("metadata.0.name", ret.Metadata.Name)
	d.Set("metadata.0.labels", ret.Metadata.Labels)
	d.Set("spec.0", string(specBytes))

	return nil
}

func resourceIngestMetadataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	kind := "IngestMetadata"
	name := d.Get("metadata.0.name").(string)

	err := client.Destroy(ctx, kind, name)
	if err != nil {
		if err == errResourceNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
