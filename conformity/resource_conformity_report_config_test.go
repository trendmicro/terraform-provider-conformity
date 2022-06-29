package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type filter struct {
	Categories                 []string
	ComplianceStandards        []string
	FilterTags                 []string
	Message                    string // bool converted to string
	NewerThanDays              string // int converted to string
	OlderThanDays              string // init converted to string
	Providers                  []string
	Regions                    []string
	ReportComplianceStandardId string
	Resource                   string
	ResourceSearchMode         string
	ResourceTypes              []string
	RiskLevels                 []string
	RuleIds                    []string
	Services                   []string
	Statuses                   []string
	Suppressed                 string // bool converted to string
	SuppressedFilterMode       string
	Tags                       []string
	Text                       string
	WithChecks                 string // bool converted to string
	WithoutChecks              string // bool converted to string
}

type configuration struct {
	Emails                []string
	Frequency             string
	GenerateReportType    string
	IncludeChecks         string // bool converted to string
	Scheduled             string // bool converted to string
	SendEmail             string // bool converted to string
	ShouldEmailIncludeCsv string // bool converted to string
	ShouldEmailIncludePdf string // bool converted to string
	Title                 string
	Tz                    string
}

func TestAccResourceconformityReportConfig(t *testing.T) {
	filter := filter{
		Categories:                 []string{"security"},
		ComplianceStandards:        []string{"AWAF"},
		FilterTags:                 []string{"staging"},
		Message:                    "true",
		NewerThanDays:              "5",
		OlderThanDays:              "5",
		Providers:                  []string{"AWS"},
		Regions:                    []string{"us-west-2"},
		ReportComplianceStandardId: "ISO27001",
		Resource:                   "joh?Smh",
		ResourceSearchMode:         "text",
		ResourceTypes:              []string{"some_resource"},
		RiskLevels:                 []string{"VERY_HIGH"},
		RuleIds:                    []string{"EC2"},
		Services:                   []string{"IAM"},
		Statuses:                   []string{"SUCCESS"},
		Suppressed:                 "true",
		SuppressedFilterMode:       "v1",
		Tags:                       []string{"some_tag"},
		Text:                       "some_text",
		WithChecks:                 "true",
		WithoutChecks:              "false",
	}
	config := configuration{
		Emails:                []string{"youremail@somecompany.com"},
		Frequency:             "* * *",
		GenerateReportType:    "GENERIC",
		IncludeChecks:         "true",
		Scheduled:             "true",
		SendEmail:             "true",
		ShouldEmailIncludeCsv: "true",
		ShouldEmailIncludePdf: "true",
		Title:                 "Daily report of IAM",
		Tz:                    "Asia/Manila",
	}

	// setting for update
	UpdatedConfig := config
	UpdatedConfig.Emails = []string{"joe@somecompany.com"}
	UpdatedConfig.Frequency = "*/3 * *"
	UpdatedConfig.GenerateReportType = "COMPLIANCE-STANDARD"
	UpdatedConfig.IncludeChecks = "false"
	UpdatedConfig.ShouldEmailIncludeCsv = "false"
	UpdatedConfig.ShouldEmailIncludePdf = "false"
	UpdatedConfig.Title = "Daily report of AppService"
	UpdatedConfig.Tz = "Asia/Singapore"

	UpdatedFilter := filter
	UpdatedFilter.Categories = []string{"cost-optimisation"}
	UpdatedFilter.ComplianceStandards = []string{"CISAWSF"}
	UpdatedFilter.FilterTags = []string{"prod"}
	UpdatedFilter.NewerThanDays = "6"
	UpdatedFilter.OlderThanDays = "6"
	UpdatedFilter.Providers = []string{"azure"}
	UpdatedFilter.Regions = []string{"eu-west-1"}
	UpdatedFilter.ReportComplianceStandardId = "CISAWSF"
	UpdatedFilter.Resource = "some_resource"
	UpdatedFilter.ResourceTypes = []string{"some_resource_2"}
	UpdatedFilter.RiskLevels = []string{"LOW"}
	UpdatedFilter.RuleIds = []string{"AppService"}
	UpdatedFilter.Services = []string{"AppService"}
	UpdatedFilter.Statuses = []string{"FAILURE"}
	UpdatedFilter.Suppressed = "false"
	UpdatedFilter.SuppressedFilterMode = "v2"
	UpdatedFilter.Tags = []string{"some_tag_2"}
	UpdatedFilter.Text = "some_text_2"

	// setting for error if scheduled is false and frequency got value
	checkScheduledFalseFrequencyConfig := config
	checkScheduledFalseFrequencyConfig.Scheduled = "false"

	// setting for error if scheduled is false and tz got value
	checkScheduledFalseTzConfig := config
	checkScheduledFalseTzConfig.Scheduled = "false"
	checkScheduledFalseTzConfig.Frequency = ""

	// setting for error if scheduled is false and send_email got value
	checkScheduledFalseSendEmailConfig := config
	checkScheduledFalseSendEmailConfig.Scheduled = "false"
	checkScheduledFalseSendEmailConfig.Frequency = ""
	checkScheduledFalseSendEmailConfig.Tz = ""

	// setting for error if scheduled is true but frequency not set
	checkScheduledTrueFrequencyConfig := config
	checkScheduledTrueFrequencyConfig.Scheduled = "true"
	checkScheduledTrueFrequencyConfig.Frequency = ""

	// setting for error if scheduled is true but tz not set
	checkScheduledTrueTzConfig := config
	checkScheduledTrueTzConfig.Scheduled = "true"
	checkScheduledTrueTzConfig.Tz = ""

	// setting for frequency validation
	checkFrequencyConfig := config
	checkFrequencyConfig.Frequency = "sample_cron"

	// setting for tz validation
	checkTzConfig := config
	checkTzConfig.Tz = "Nowhere/City"

	accountId := "account_id"
	groupId := "group_id"
	updatedAccountId := "updated_account_id"
	updatedGroupId := "updated_group_id"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckConformityReportConfigDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			// check create
			{
				Config: testAccCheckConformityReportConfigBasic(filter, config, groupId, accountId),
				Check: resource.ComposeTestCheckFunc(

					// ids
					resource.TestCheckResourceAttr("conformity_report_config.report", "account_id", "account_id"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "group_id", "group_id"),

					//configuration
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.emails.0", "youremail@somecompany.com"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.frequency", "* * *"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.generate_report_type", "GENERIC"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.include_checks", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.scheduled", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.send_email", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.should_email_include_csv", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.should_email_include_pdf", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.title", "Daily report of IAM"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.tz", "Asia/Manila"),

					//filter
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.categories.0", "security"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.compliance_standards.0", "AWAF"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.filter_tags.0", "staging"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.message", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.newer_than_days", "5"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.older_than_days", "5"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.providers.0", "AWS"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.regions.0", "us-west-2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.report_compliance_standard_id", "ISO27001"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource", "joh?Smh"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource_search_mode", "text"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource_types.0", "some_resource"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.risk_levels.0", "VERY_HIGH"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.rule_ids.0", "EC2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.services.0", "IAM"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.statuses.0", "SUCCESS"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.suppressed", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.suppressed_filter_mode", "v1"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.tags.0", "some_tag"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.text", "some_text"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.with_checks", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.without_checks", "false"),
				),
			},
			// check update
			{
				Config: testAccCheckConformityReportConfigBasic(UpdatedFilter, UpdatedConfig, groupId, accountId),
				Check: resource.ComposeTestCheckFunc(

					// ids
					resource.TestCheckResourceAttr("conformity_report_config.report", "account_id", "account_id"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "group_id", "group_id"),

					//configuration
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.emails.0", "joe@somecompany.com"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.frequency", "*/3 * *"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.generate_report_type", "COMPLIANCE-STANDARD"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.include_checks", "false"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.scheduled", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.send_email", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.should_email_include_csv", "false"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.should_email_include_pdf", "false"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.title", "Daily report of AppService"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "configuration.0.tz", "Asia/Singapore"),

					//filter
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.categories.0", "cost-optimisation"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.compliance_standards.0", "CISAWSF"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.filter_tags.0", "prod"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.message", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.newer_than_days", "6"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.older_than_days", "6"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.providers.0", "azure"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.regions.0", "eu-west-1"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.report_compliance_standard_id", "CISAWSF"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource", "some_resource"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource_search_mode", "text"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.resource_types.0", "some_resource_2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.risk_levels.0", "LOW"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.rule_ids.0", "AppService"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.services.0", "AppService"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.statuses.0", "FAILURE"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.suppressed", "false"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.suppressed_filter_mode", "v2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.tags.0", "some_tag_2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.text", "some_text_2"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.with_checks", "true"),
					resource.TestCheckResourceAttr("conformity_report_config.report", "filter.0.without_checks", "false"),
				),
			},
			{
				// check: when scheduled is false but frequency got value
				Config:      testAccCheckConformityReportConfigBasic(filter, checkScheduledFalseFrequencyConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"scheduled" must be true in order to use "frequency"`),
			},
			{
				// check: when scheduled is false but tz got value
				Config:      testAccCheckConformityReportConfigBasic(filter, checkScheduledFalseTzConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"scheduled" must be true in order to use "tz"`),
			},
			{
				// check: when scheduled is false but send_email got value
				Config:      testAccCheckConformityReportConfigBasic(filter, checkScheduledFalseSendEmailConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"scheduled" must be true in order to use "send_email"`),
			},
			{
				// check: when scheduled is true but frequency not set
				Config:      testAccCheckConformityReportConfigBasic(filter, checkScheduledTrueFrequencyConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"frequency" must be set in order to use "scheduled"`),
			},
			{
				// check: when scheduled is true but tz not set
				Config:      testAccCheckConformityReportConfigBasic(filter, checkScheduledTrueTzConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"tz" must be set in order to use "scheduled"`),
			},
			{
				// check: frequency validation
				Config:      testAccCheckConformityReportConfigBasic(filter, checkFrequencyConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`"configuration.0.frequency" expected cron expression, got:`),
			},
			{
				// check: tz validation
				Config:      testAccCheckConformityReportConfigBasic(filter, checkTzConfig, groupId, accountId),
				ExpectError: regexp.MustCompile(`given value of "tz" is not supported, got: ` + checkTzConfig.Tz),
			},
			{
				// check: account_id and group_id update
				Config: testAccCheckConformityReportConfigBasic(filter, config, updatedGroupId, updatedAccountId),
				// No check function is given because we expect this configuration
				// to fail before any infrastructure is created
				ExpectError: regexp.MustCompile("'account_id', and 'group_id' cannot be changed"),
			},
		},
	})
}

func testAccCheckConformityReportConfigBasic(f filter, c configuration, groupId string, accountId string) string {
	var (
		frequency string
		tz        string
	)
	if c.Frequency != "" {
		frequency = `frequency = "` + c.Frequency + `"`
	} else {
		frequency = ""
	}
	if c.Tz != "" {
		tz = `tz = "` + c.Tz + `"`
	} else {
		tz = ""
	}

	return `
	resource "conformity_report_config" "report" {
		account_id = "` + accountId + `"
		group_id = "` + groupId + `"

		configuration {
			emails = ["` + c.Emails[0] + `"]
			` + frequency + `
			generate_report_type = "` + c.GenerateReportType + `"
			include_checks = "` + c.IncludeChecks + `"
			scheduled = "` + c.Scheduled + `"
			send_email = "` + c.SendEmail + `"
			should_email_include_csv = "` + c.ShouldEmailIncludeCsv + `"
			should_email_include_pdf = "` + c.ShouldEmailIncludePdf + `"
			title = "` + c.Title + `"
			` + tz + `
		}
		filter {
			categories = ["` + f.Categories[0] + `"]
			compliance_standards = ["` + f.ComplianceStandards[0] + `"]
			filter_tags = ["` + f.FilterTags[0] + `"]
			message = ` + f.Message + `
			newer_than_days = ` + f.NewerThanDays + `
			older_than_days = ` + f.OlderThanDays + `
			providers = ["` + f.Providers[0] + `"]
			regions = ["` + f.Regions[0] + `"]
			report_compliance_standard_id = "` + f.ReportComplianceStandardId + `"
			resource  = "` + f.Resource + `"
			resource_search_mode = "` + f.ResourceSearchMode + `"
			resource_types = ["` + f.ResourceTypes[0] + `"]
			risk_levels = ["` + f.RiskLevels[0] + `"]
			rule_ids = ["` + f.RuleIds[0] + `"]
			services = ["` + f.Services[0] + `"]
			statuses = ["` + f.Statuses[0] + `"]
			suppressed = ` + f.Suppressed + `
			suppressed_filter_mode = "` + f.SuppressedFilterMode + `"
			tags = ["` + f.Tags[0] + `"]
			text =  "` + f.Text + `"
			with_checks = "` + f.WithChecks + `"
			without_checks = "` + f.WithoutChecks + `"

		}	
	}
	`
}

func testAccCheckConformityReportConfigDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_report_config" {
			continue
		}
		reportId := rs.Primary.ID

		deleteReportConfig, err := c.DeleteReportConfig(reportId)
		if deleteReportConfig.Meta.Status != "deleted" {
			return fmt.Errorf("Conformity Report Config not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
