package conformity

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestAccResourceConformityCustomRule(t *testing.T) {
	name := "S3 Bucket Custom Rule update name"
	description := "This custom rule ensures S3 buckets follow our best practice"
	notes := "If this is broken, please follow these steps:\\n1. Step one \\n2. Step two\\n"
	resourceType := "s3-bucket"
	service := "S3"
	categories := []string{"security", "sustainability"}
	updatedCategories := []string{"reliability", "security", "performance-efficiency", "operational-excellence"}
	severity := "HIGH"
	updatedSeverity := "EXTREME"
	cloud := "aws"
	updatedCloud := "azure"
	attributes := []interface{}{map[string]interface{}{
		"name":     "bucketName",
		"path":     "bucketPath",
		"required": "true",
	}}
	updatedAttributes := []interface{}{
		map[string]interface{}{
			"name":     "bucketName",
			"path":     "data.Name",
			"required": "true",
		},
		map[string]interface{}{
			"name":     "ownerName",
			"path":     "data.OwnerName",
			"required": "true",
		},
		map[string]interface{}{
			"name":     "ownerEmail",
			"path":     "data.OwnerEmail",
			"required": "true",
		},
		map[string]interface{}{
			"name":     "ownerMobile",
			"path":     "data.OwnerMobile",
			"required": "true",
		},
	}
	rules := map[string]interface{}{
		"operation": "any",
		"conditions": []interface{}{map[string]interface{}{
			"fact":     "bucketName",
			"operator": "pattern",
			"value":    "^([a-zA-Z0-9_-]){1,32}$",
		},
		},
		"event_type": "Bucket name is longer than 32 characters",
	}
	updatedRules := map[string]interface{}{
		"operation": "all",
		"conditions": []interface{}{
			map[string]interface{}{
				"fact":     "bucketName",
				"operator": "pattern",
				"value":    "^([a-zA-Z0-9_-]){1,32}$",
			},
			map[string]interface{}{
				"fact":     "OwnerName",
				"operator": "pattern",
				"value":    "^([a-zA-Z0-9_-]){1,32}$",
			},
			map[string]interface{}{
				"fact":     "OwnerEmail",
				"operator": "pattern",
				"value":    "^[a-zA-Z0-9_\\\\.+-]+@[a-zA-Z0-9-]+\\\\.[a-zA-Z0-9-.]+$",
			},
			map[string]interface{}{
				"fact":     "OwnerMobile",
				"operator": "pattern",
				"value":    "^([0-9]){1,12}$",
			},
		},
		"event_type": "Bucket name is longer than 32 characters",
	}
	sort.Strings(updatedCategories)
	sort.Slice(updatedAttributes, func(i, j int) bool {
		return updatedAttributes[i].(map[string]interface{})["name"].(string) < updatedAttributes[j].(map[string]interface{})["name"].(string)
	})
	sort.Slice(updatedRules["conditions"], func(i, j int) bool {
		return updatedRules["conditions"].([]interface{})[i].(map[string]interface{})["fact"].(string) < updatedRules["conditions"].([]interface{})[j].(map[string]interface{})["fact"].(string)
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingName(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingDescription(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingService(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingResourceType(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingCategories(),
				ExpectError: regexp.MustCompile("Not enough list items"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingAttributes(),
				ExpectError: regexp.MustCompile("Insufficient attributes blocks"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigMissingConditions(),
				ExpectError: regexp.MustCompile("Insufficient conditions blocks"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigInvalidProvider(),
				ExpectError: regexp.MustCompile("expected cloud_provider to be one of \\[aws azure gcp\\], got invalid"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigInvalidCategory(),
				ExpectError: regexp.MustCompile("expected categories.0 to be one of \\[security cost-optimisation reliability performance-efficiency operational-excellence sustainability\\], got invalid"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigInvalidOperation(),
				ExpectError: regexp.MustCompile("expected rules.0.operation to be one of \\[any all\\], got invalid"),
			},
			{
				Config:      testAccCheckConformityCustomRuleConfigInvalidSeverity(),
				ExpectError: regexp.MustCompile("expected severity to be one of \\[LOW MEDIUM HIGH VERY_HIGH EXTREME\\], got invalid"),
			},
			{
				Config: testAccCheckConformityCustomRuleConfigBasic(
					name, description, resourceType, service, categories, severity, cloud, true, attributes, rules,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "id", "some_id"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "type", "CustomRule"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "name", name),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "description", description),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "service", service),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "resource_type", resourceType),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "severity", severity),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "cloud_provider", cloud),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "enabled", "true"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.#", "2"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.0", categories[0]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.1", categories[1]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.#", "1"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.name", attributes[0].(map[string]interface{})["name"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.path", attributes[0].(map[string]interface{})["path"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.required", attributes[0].(map[string]interface{})["required"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.#", "1"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.operation", rules["operation"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.event_type", rules["event_type"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.#", "1"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.fact", rules["conditions"].([]interface{})[0].(map[string]interface{})["fact"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.operator", rules["conditions"].([]interface{})[0].(map[string]interface{})["operator"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.value", rules["conditions"].([]interface{})[0].(map[string]interface{})["value"].(string)),
				),
			},
			{
				Config: testAccCheckConformityCustomRuleConfigFull(
					transform(name), transform(description), transform(resourceType), transform(service), updatedCategories,
					updatedSeverity, transform(notes), updatedCloud, false, updatedAttributes, updatedRules,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "id", "some_id"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "type", "CustomRule"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "name", transform(name)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "description", transform(description)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "service", transform(service)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "resource_type", transform(resourceType)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "severity", updatedSeverity),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "cloud_provider", updatedCloud),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.#", "4"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.0", updatedCategories[0]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.1", updatedCategories[1]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.2", updatedCategories[2]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "categories.3", updatedCategories[3]),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.#", "4"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.name", updatedAttributes[0].(map[string]interface{})["name"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.path", updatedAttributes[0].(map[string]interface{})["path"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.0.required", updatedAttributes[0].(map[string]interface{})["required"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.1.name", updatedAttributes[1].(map[string]interface{})["name"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.1.path", updatedAttributes[1].(map[string]interface{})["path"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.1.required", updatedAttributes[1].(map[string]interface{})["required"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.2.name", updatedAttributes[2].(map[string]interface{})["name"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.2.path", updatedAttributes[2].(map[string]interface{})["path"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.2.required", updatedAttributes[2].(map[string]interface{})["required"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.3.name", updatedAttributes[3].(map[string]interface{})["name"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.3.path", updatedAttributes[3].(map[string]interface{})["path"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "attributes.3.required", updatedAttributes[3].(map[string]interface{})["required"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.#", "1"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.operation", updatedRules["operation"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.event_type", updatedRules["event_type"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.#", "4"),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.fact", updatedRules["conditions"].([]interface{})[0].(map[string]interface{})["fact"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.operator", updatedRules["conditions"].([]interface{})[0].(map[string]interface{})["operator"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.0.value", strings.ReplaceAll(updatedRules["conditions"].([]interface{})[0].(map[string]interface{})["value"].(string), "\\\\.", "\\.")),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.1.fact", updatedRules["conditions"].([]interface{})[1].(map[string]interface{})["fact"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.1.operator", updatedRules["conditions"].([]interface{})[1].(map[string]interface{})["operator"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.1.value", updatedRules["conditions"].([]interface{})[1].(map[string]interface{})["value"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.2.fact", updatedRules["conditions"].([]interface{})[2].(map[string]interface{})["fact"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.2.operator", updatedRules["conditions"].([]interface{})[2].(map[string]interface{})["operator"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.2.value", updatedRules["conditions"].([]interface{})[2].(map[string]interface{})["value"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.3.fact", updatedRules["conditions"].([]interface{})[3].(map[string]interface{})["fact"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.3.operator", updatedRules["conditions"].([]interface{})[3].(map[string]interface{})["operator"].(string)),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "rules.0.conditions.3.value", updatedRules["conditions"].([]interface{})[3].(map[string]interface{})["value"].(string)),
				),
			},
			{
				ResourceName:      "conformity_custom_rule.example",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "remediation_notes", notes),
					resource.TestCheckResourceAttr("conformity_custom_rule.example", "enabled", "false"),
				),
			},
		},
	})
}

func transform(val string) string {
	return fmt.Sprintf("updated %s", val)
}

func stringifyMap(m map[string]interface{}) string {
	ret := ""
	for k, v := range m {
		if reflect.TypeOf(v).Kind().String() == "string" {
			ret += fmt.Sprintf("%s = \"%s\"\n", k, v)
		} else if reflect.TypeOf(v).Kind().String() == "map" {
			ret += fmt.Sprintf("%s {\n", k)
			ret += stringifyMap(v.(map[string]interface{}))
			ret += "}\n"
		} else if reflect.TypeOf(v).Kind().String() == "slice" {
			for _, obj := range v.([]interface{}) {
				ret += fmt.Sprintf("%s {\n", k)
				ret += stringifyMap(obj.(map[string]interface{}))
				ret += "}\n"
			}
		}
	}
	return ret
}

func stringifySlice(k string, s []interface{}) string {
	ret := ""
	for _, e := range s {
		ret += fmt.Sprintf("%s {\n", k)
		ret += stringifyMap(e.(map[string]interface{}))
		ret += "}\n"
	}
	return ret
}

func testAccCheckConformityCustomRuleConfigBasic(name string, description string, resourceType string, service string, categories []string, severity string, cloud string, enabled bool, attributes []interface{}, rules map[string]interface{}) string {
	return fmt.Sprintf(
		`resource "conformity_custom_rule" "example" {
				  name             = "%s"
				  description      = "%s"
				  service          = "%s"
				  resource_type     = "%s"
				  categories       = ["%s"]
				  severity         = "%s"
				  cloud_provider         = "%s"
				  enabled          = %t
				  %s
				  rules {
					%s
				  }
			}`,
		name, description, service, resourceType, strings.Join(categories, "\",\""), severity, cloud, enabled,
		stringifySlice("attributes", attributes), stringifyMap(rules),
	)
}

func testAccCheckConformityCustomRuleConfigFull(name string, description string, resourceType string, service string, categories []string, severity string, remediationNotes string, cloud string, enabled bool, attributes []interface{}, rules map[string]interface{}) string {
	return fmt.Sprintf(
		`resource "conformity_custom_rule" "example" {
				  name             = "%s"
				  description      = "%s"
				  remediation_notes = "%s"
				  service          = "%s"
				  resource_type     = "%s"
				  categories       = ["%s"]
				  severity         = "%s"
				  cloud_provider         = "%s"
				  enabled          = %t
				  %s
				  rules {
					%s
				  }
				}`,
		name, description, remediationNotes, service, resourceType, strings.Join(categories, "\",\""), severity, cloud,
		enabled, stringifySlice("attributes", attributes), stringifyMap(rules),
	)
}

func testAccCheckConformityCustomRuleConfigMissingName() string {
	return `resource "conformity_custom_rule" "example" {
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingDescription() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingService() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingResourceType() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingCategories() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = []
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingEnabled() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingAttributes() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigMissingConditions() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigInvalidProvider() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "invalid"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigInvalidCategory() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "invalid", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "any"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigInvalidOperation() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "HIGH"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "invalid"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}

func testAccCheckConformityCustomRuleConfigInvalidSeverity() string {
	return `resource "conformity_custom_rule" "example" {
      name             = "S3 Bucket Custom Rule update name"
	  description      = "This custom rule ensures S3 buckets follow our best practice updated"
	  remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	  service          = "S3"
	  resource_type     = "s3-bucket"
	  categories       = ["security", "sustainability", "performance-efficiency", "operational-excellence"]
	  severity         = "invalid"
	  cloud_provider         = "azure"
	  enabled          = true
	  attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
		operation = "all"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
	}`
}
