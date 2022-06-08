package provider

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func CheckResourceAttrEqualToJson(name, key, jsonStr string) func(*terraform.State) error {
	return resource.TestCheckResourceAttrWith(
		name, key, func(value string) error {
			var (
				expected, actual interface{}
			)
			if err := json.Unmarshal([]byte(jsonStr), &expected); err != nil {
				return err
			}
			if err := json.Unmarshal([]byte(value), &actual); err != nil {
				return err
			}
			if !reflect.DeepEqual(expected, actual) {
				return fmt.Errorf("value mismatch")
			}
			return nil
		})
}
