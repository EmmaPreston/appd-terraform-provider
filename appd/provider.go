package appd

import (
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"appd_health_rule": resourceHealthRule(),
			"appd_action":      resourceAction(),
			"appd_policy":      resourcePolicy(),
		},
		Schema: map[string]*schema.Schema{
			"secret":              {Type: schema.TypeString, Sensitive: true, Required: true},
			"controller_base_url": {Type: schema.TypeString, Required: true},
		},
		ConfigureFunc: func(data *schema.ResourceData) (interface{}, error) {
			return &client.AppDClient{
				BaseUrl: data.Get("controller_base_url").(string),
				Secret:  data.Get("secret").(string),
			}, nil
		},
	}
}
