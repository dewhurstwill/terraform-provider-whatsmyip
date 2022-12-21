package whatsmyip

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// dataSourceWhatsmyip is where we define the schema of the Terraform data source
func dataSourceWhatsmyip() *schema.Resource {
	return &schema.Resource{
		Read: whatsmyipRead,
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func whatsmyipRead(d *schema.ResourceData, meta interface{}) (err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.myip.com", nil)
	if err != nil {
		// Throw error
	}

	r, err := client.Do(req)
	if err != nil {
		// Throw error
	}
	defer r.Body.Close()

	outputs := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&outputs)
	// If an error occured, throw it
	if err != nil {
		// Throw error
	}

	outputs["id"] = outputs["ip"]
	marshalData(d, outputs)

	return
}