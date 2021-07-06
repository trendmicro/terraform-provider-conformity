package conformity

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CustomizeDiffValidateConfiguration(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {

	list := []string{"email", "sms", "ms_teams", "slack", "sns", "pager_duty", "webhook"}

	channels := []string{}

	for _, ch := range list {
		if _, ok := diff.GetOk(ch); ok {
			channels = append(channels, ch)
		}
	}

	if len(channels) == 0 {
		return fmt.Errorf("no channel configuration set found, please provide one of this: %s", strings.Join(list, ", "))
	}
	if len(channels) > 1 {
		return fmt.Errorf(`found multiple channel configuration set, please provide only one; got: %s`, strings.Join(channels, ", "))
	}
	return nil
}
