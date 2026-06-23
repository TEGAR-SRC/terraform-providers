package objectstorage

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudObjectStorage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudObjectStorageCreate,
		ReadContext:   resourceOnidelCloudObjectStorageRead,
		DeleteContext: resourceOnidelCloudObjectStorageDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Service name.",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Location/region.",
			},
			"bucket_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Initial bucket name to create.",
			},
			"versioning": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Enable bucket versioning.",
			},
			"object_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Enable object lock.",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "S3-compatible region.",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "S3-compatible endpoint URL.",
			},
			"capacity": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Storage capacity in KB.",
			},
			"used_capacity": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Used capacity in KB.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service status.",
			},
		},
	}
}

type objectStorageResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Region       string `json:"region"`
	Endpoint     string `json:"endpoint"`
	Capacity     int    `json:"capacity"`
	UsedCapacity int    `json:"used_capacity"`
	Status       string `json:"status"`
}

func resourceOnidelCloudObjectStorageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	// Create object storage service (assuming a POST to create a service,
	// if not available, we create via bucket creation)
	// The spec shows bucket creation under /object-storage/{service_id}/buckets,
	// so the service itself may be created separately or auto-provisioned.
	// For now, we handle as an existing service lookup.

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var listResult struct {
		Services []objectStorageResponse `json:"services"`
	}
	err := client.Get("/object-storage", params, &listResult)
	if err != nil {
		return diag.Errorf("Error listing object storage services: %s", err)
	}

	if len(listResult.Services) == 0 {
		return diag.Errorf("No object storage services available. Create one via Onidel Cloud console.")
	}

	svc := listResult.Services[0]
	d.SetId(svc.ID)
	log.Printf("[INFO] Object Storage service: %s", svc.ID)

	// Create bucket if bucket_name specified
	if bucketName, ok := d.GetOk("bucket_name"); ok {
		bucketBody := map[string]interface{}{
			"bucket_name": bucketName.(string),
			"versioning":  d.Get("versioning").(bool),
			"object_lock": d.Get("object_lock").(bool),
		}
		if v, ok := d.GetOk("team_id"); ok {
			bucketBody["team_id"] = v.(string)
		}
		err = client.Post("/object-storage/"+svc.ID+"/buckets", bucketBody, nil)
		if err != nil {
			return diag.Errorf("Error creating bucket: %s", err)
		}
		log.Printf("[INFO] Bucket %s created", bucketName.(string))
	}

	return resourceOnidelCloudObjectStorageRead(ctx, d, meta)
}

func resourceOnidelCloudObjectStorageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result objectStorageResponse
	err := client.Get("/object-storage/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] Object Storage not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.Name)
	d.Set("location", result.Location)
	d.Set("region", result.Region)
	d.Set("endpoint", result.Endpoint)
	d.Set("capacity", result.Capacity)
	d.Set("used_capacity", result.UsedCapacity)
	d.Set("status", result.Status)
	return nil
}

func resourceOnidelCloudObjectStorageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Object storage service deletion is not available via API (managed via console).
	// Remove from state.
	log.Printf("[INFO] Object Storage service %s removed from state", d.Id())
	d.SetId("")
	return nil
}
