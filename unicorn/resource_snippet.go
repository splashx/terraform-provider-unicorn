package unicorn

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	gitlab "github.com/xanzy/go-gitlab"
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
	client := m.(*gitlab.Client)

	createSnippetOptions := &gitlab.CreateSnippetOptions{
		Title:       gitlab.String(d.Get("title").(string)),
		FileName:    gitlab.String(d.Get("file_name").(string)),
		Content:     gitlab.String(d.Get("content").(string)),
		Description: gitlab.String("Created by terraform-unicorn"),
		Visibility:  gitlab.Visibility(gitlab.PrivateVisibility),
	}
	snippet, _, err := client.Snippets.CreateSnippet(createSnippetOptions)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(snippet.ID))

	return resourceSnippetRead(d, m)
}

func resourceSnippetRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitlab.Client)

	snippetID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	snippet, resp, err := client.Snippets.GetSnippet(snippetID)

	if err != nil {
		if resp.StatusCode == 404 || snippet == nil {
			fmt.Printf("[WARN] removing snippet id %d from state - not found in gitlab", snippetID)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnippetDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gitlab.Client)

	snippetID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	_, err = client.Snippets.DeleteSnippet(snippetID)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

// func resourceSnippetUpdate(d *schema.ResourceData, m interface{}) error {
// 	return resourceSnippetRead(d, m)
// }

// func resourceSnippetExists(d *schema.ResourceData, m interface{}) (bool, error) {
// 	return true, nil
// }
