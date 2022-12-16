resource "conformity_report_config" "report" {
  // required | type: string
  account_id = "" 
  // optional | type: string
  group_id   = ""   
  configuration { 
    // optional | type: array of strings
    email = [] 
    // optional | type: string
    frequency = ""
    // optional | type: string | default: GENERIC
    // value can be: "GENERIC", "COMPLIANCE-STANDARD"
    generate_report_type = "" 
    // optional | type: bool | default: true
    include_account_names=""
    // optional | type: bool | default: false
    include_checks = bool
    // optional | type: bool | default: false
    scheduled = bool
    //optional | type: bool | default: false
    send_email = bool
    // optional | type: bool | default: false
    should_email_include_csv = bool
    // optional | type: bool | default: true
    should_email_include_pdf = bool
    // required | type: string 
    title = ""
    // optional | type: string
    tz = ""
  }
  filter {
    // optional | type: array of strings 
    # value can be : "security", "cost-optimisation", "reliability", "performance-efficiency", "operational-excellence" 
    categories = []
    // optional | type: array of strings
    # value can be: "AWAF" "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2", "NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"
    compliance_standards = []
    // optional | type: array of strings 
    filter_tags = []
    // optional | type: bool
    message = bool
    // optional | type: number
    newer_than_days = 0
    // optional | type: number
    older_than_days = 0
    // optional | typel: array of strings
    // value can be: "aws", "azure"
    providers = []
    // optional | typel: array of strings
    // value can be: array of AWS region
    regions = []
    // optional | stype: string
    // value can be: "AWAF", "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2", "NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"
    report_compliance_standard_id = ""
    // optional | type: string
    resource = ""
    // optional | type: string 
    # value can be a text or regex
    resource_search_mode = ""
    // optional | type: array of strings
    resource_types = []
    // optional | type: string
    # value can be: 
    # "LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"
    risk_levels = ""
    // optional | type: array of strings
    rule_ids = []
    // optional | type: array of strings
    # value can be: 
    # "EC2" "ELB" "EBS" "VPC" "S3" "CloudTrail" "Route53" "RDS" "IAM" "RTM" "KMS" "SNS" "SQS" "CloudFormation"  "Config" "CloudFront" "AutoScaling" 
    # "Redshift" "CloudWatch" "CloudWatchEvents" "CloudWatchLogs" "ResourceGroup" "SES" "DynamoDB" "ElastiCache" "Elasticsearch" "WorkSpaces" "ACM" "Budgets" 
    # "Inspector" "TrustedAdvisor" "Shield" "EMR" "WAF" "Lambda" "Support" "Kinesis" "Organizations" "EFS" "ElasticBeanstalk" "Macie" "ELBv2" "CloudConformity" 
    # "APIGateway" "GuardDuty" "Health" "ConfigService" "MQ" "Firehose" "SSM" "Route53Domains" "SageMaker" "DAX" "Neptune" "ECR" "Glue" "XRay" "SecretsManager" 
    # "DocumentDB" "DMS" "Miscellaneous" "EKS" "Backup" "StorageGateway" "ECS" "SecurityHub" "Comprehend" "WellArchitected" "AccessAnalyzer" "StorageAccounts" 
    # "SecurityCenter" "ActiveDirectory" "MySQL" "PostgreSQL" "Sql" "Monitor" "AppService" "Network" "ActivityLog" "VirtualMachines" "AKS" "KeyVault" 
    # "Locks" "AccessControl" "Advisor" "RecoveryServices" "Resources" "Subscriptions" "CosmosDB" "RedisCache" "Search" "AppInsights"
    services = []
    // optional | type: array of strings
    # value can be: "SUCCESS", "FAILURE"
    statuses = []
    // optional | type: bool
    suppressed = bool
    // optional | type: string | default: v1
    # value can be: "v1", "v2"
    suppressed_filter_mode = ""
    // optional | type: array of strings
    // value can be: metadata tags for aws resources
    tags = []
    //optional | type: string
    text = ""
    // optional | type: bool | default: false
    with_checks = bool
    // optional | type: bool | default: false
    without_checks = bool
  }
}