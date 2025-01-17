package papertrail

import (
	"fmt"

	"github.com/Etienne-Carriere/terraform-provider-papertrail/lowlevel/goptrail"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token, ok := d.Get("token").(string)
	if !ok {
		return nil, fmt.Errorf("Cannot convert token %v to string", d.Get("token"))
	}
	client := goptrail.NewClient(token)
	return client, nil
}
