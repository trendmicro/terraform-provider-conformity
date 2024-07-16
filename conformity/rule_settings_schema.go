package conformity

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ExceptionsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// need to verify which one it depricated both filter_tag and tags
				"filter_tags": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"resources": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"tags": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func ExtraSettingSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{"multiple-string-values", "multiple-number-values", "multiple-aws-account-values",
						"choice-multiple-value", "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "tags",
						"countries", "multiple-ip-values", "regions", "ignored-regions", "multiple-object-values", "multiple-vpc-gateway-mappings"}, true),
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"regions": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
						// add validation here
						// region should follow the correct syntax
					},
				},
				"tags": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"multiple_object_values": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"event_name": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"event_source": {
								Type:     schema.TypeString,
								Required: true,
							},
							"user_identity_type": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
				"values":   valuesSchema(),
				"mappings": mappingSchema(),
			},
		},
	}
}

func mappingSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"values": {
					Type:     schema.TypeSet,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"type": {
								Type:     schema.TypeString,
								Required: true,
							},
							"name": {
								Type:     schema.TypeString,
								Required: true,
							},
							"value": {
								Type:     schema.TypeString,
								Optional: true,
								ValidateFunc: validation.StringMatch(
									regexp.MustCompile("^vpc-"), `value should start with "vpc-"`),
							},
							"values": {
								Type:     schema.TypeSet,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"value": {
											Type:     schema.TypeString,
											Required: true,
											ValidateFunc: validation.StringMatch(
												regexp.MustCompile("^nat-|igw-"), `value should start with "nat-" or "igw-"`),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func valuesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// attributes `value` is commonly use for all the types therefore It should be required
				// eg. [{"value": "ELBSecurityPolicy-2016-08"}]
				"label": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"customised": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"settings": valueSettingSchema(),
			},
		},
	}
}

func valueSettingSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": {
					Type:     schema.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{"multiple-string-values", "multiple-number-values", "multiple-aws-account-values",
						"choice-multiple-value", "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "tags",
						"countries", "multiple-ip-values", "regions", "ignored-regions", "multiple-object-values", "multiple-vpc-gateway-mappings"}, true),
				},
				"values": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"value": {
								Type:     schema.TypeString,
								Required: true,
							},
							"default": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}
