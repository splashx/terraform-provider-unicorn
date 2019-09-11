package unicorn

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		// Create: resourceFileCreate,
		// Read:   resourceFileRead,
		// Update: resourceFileUpdate,
		// Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
		},
	}
}
