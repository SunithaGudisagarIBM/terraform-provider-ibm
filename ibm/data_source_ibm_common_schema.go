package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func filterIBMDataSources(set *schema.Set, data *[]map[string]interface{}) *[]map[string]interface{} {
	var filteredVals []map[string]interface{}
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var key string
		if k, ok := m["name"]; ok {
			key = k.(string)
		}
		var values []interface{}
		if v, ok := m["values"]; ok {
			values = v.([]interface{})
		}
		vals := make([]string, len(values))
		for i, v := range values {
			vals[i] = v.(string)
		}
		for _, each := range *data {
			var dataval string
			if v, ok := each[key]; ok {
				dataval = v.(string)
				for _, v := range vals {
					if dataval == v {
						filteredVals = append(filteredVals, each)
					}
				}
			}
		}
	}
	return &filteredVals
}

func dataSourceFiltersSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Computed:    false,
		Description: "List of filters",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of the filter",
				},

				"values": {
					Type:        schema.TypeList,
					Required:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "Values of the filter",
				},
			},
		},
	}
}
