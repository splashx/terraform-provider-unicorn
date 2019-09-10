package unicorn

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a schema.Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITLAB_TOKEN", nil),
				Description: "Gitlab Rest API token",
			},
			"server_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITLAB_URL", nil),
				Description: "Gitlab Server URL, defaults to gitlab.com",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"snippet": resourceSnippet(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("api_key").(string),
		BaseURL: d.Get("server_url").(string),
	}

	return config.Client()
}
