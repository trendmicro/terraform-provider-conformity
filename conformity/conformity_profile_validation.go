package conformity

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CustomizeDiffValidateProfileValue(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if _, ok := diff.GetOk("included"); !ok {
		return nil
	}
	included := diff.Get("included").([]interface{})
	for i, v := range included {
		incItem := v.(map[string]interface{})
		ess := incItem["extra_settings"].(*schema.Set).List()
		if len(ess) != 0 {
			for esi, es := range ess {
				esItem := es.(map[string]interface{})
				value := esItem["value"].(string)
				log.Printf("[DEBUG] customize type: %v", esItem["type"].(string))

				if esItem["type"].(string) == "single-number-value" || esItem["type"].(string) == "ttl" {
					if _, err := strconv.Atoi(value); err != nil {
						return fmt.Errorf(`included.%d.extra_settings.%d.value is not valid. Must follow the valid syntax "<number>", got: "%s"`, i, esi, value)
					}
				}

			}
		}
	}
	return nil
}
