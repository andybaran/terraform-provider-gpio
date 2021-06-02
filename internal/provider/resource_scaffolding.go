package provider

import (
	"context"

	pb "github.com/andybaran/fictional-goggles/terragpio"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePWM() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceScaffoldingCreate,
		ReadContext:   resourceScaffoldingRead,
		UpdateContext: resourceScaffoldingUpdate,
		DeleteContext: resourceScaffoldingDelete,

		Schema: map[string]*schema.Schema{
			"Pin": {
				// This description is used by the documentation generator and the language server.
				Description: "GPIO Pin",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"Dutycycle": {
				// This description is used by the documentation generator and the language server.
				Description: "Duty cycle",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"Frequency": {
				// This description is used by the documentation generator and the language server.
				Description: "Frequency",
				Type:        schema.TypeString,
				Optional:    false,
			},
		},
	}
}

func resourceScaffoldingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}

func resourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceScaffoldingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceScaffoldingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
