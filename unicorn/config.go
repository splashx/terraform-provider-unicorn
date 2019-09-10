package unicorn

import (
	"strings"

	gitlab "github.com/xanzy/go-gitlab"
)

// Config defines the Gitlab client config
type Config struct {
	Token   string
	BaseURL string
}

// Client returns a Gitlab API client
func (c *Config) Client() (interface{}, error) {

	client := gitlab.NewClient(nil, c.Token)
	baseURL := strings.Trim(c.BaseURL, " ")

	if baseURL != "" {
		err := client.SetBaseURL(baseURL)
		if err != nil {
			// The BaseURL supplied wasn't valid, bail.
			return nil, err
		}
	}

	// Test the credentials by checking we can get information about the authenticated user.
	_, _, err := client.Users.CurrentUser()
	if err != nil {
		return nil, err
	}
	return client, nil
}
