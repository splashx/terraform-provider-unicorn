package unicorn

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSnippet() *schema.Resource {
	return &schema.Resource{
		// The golang module used for GitlabAPI can handle these fields:
		// https://github.com/xanzy/go-gitlab/blob/master/snippets.go
		// Required is:
		// Title
		// Filename
		// content
		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},

		Create: resourceSnippetCreate,
		Read:   resourceSnippetRead,
		Delete: resourceSnippetDelete,
		//Update: resourceSnippetUpdate,
		//Exists: resourceSnippetExists,
	}
}

func resourceSnippetCreate(d *schema.ResourceData, m interface{}) error {
	return resourceSnippetRead(d, m)
}

func resourceSnippetRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSnippetDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSnippetUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSnippetRead(d, m)
}

func resourceSnippetExists(d *schema.ResourceData, m interface{}) (bool, error) {
	return true, nil
}
