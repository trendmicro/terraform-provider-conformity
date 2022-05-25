package cloudconformity

type externalIdData struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
}
type apiKeyList struct {
	Data []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
}
type AccountKeys struct {
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}
type accountAccess struct {
	Keys              *AccountKeys `json:"keys,omitempty"`
	Type              string       `json:"type,omitempty"`
	SubscriptionId    string       `json:"subscriptionId,omitempty"`
	ActiveDirectoryId string       `json:"activeDirectoryId,omitempty"`
}
type AccountConfiguration struct {
	ExternalId string `json:"externalId,omitempty"`
	RoleArn    string `json:"roleArn,omitempty"`
}
type CloudData struct {
	Azure struct {
		SubscriptionId string `json:"subscriptionId"`
	} `json:"azure"`
}

type GetRuleSettings struct {
	Enabled       bool                   `json:"enabled"`
	Id            string                 `json:"id"`
	RiskLevel     string                 `json:"riskLevel"`
	RuleExists    bool                   `json:"ruleExists"`
	ExtraSettings []*RuleSettingExtra    `json:"extraSettings"`
	Exceptions    *RuleSettingExceptions `json:"exceptions,omitempty"`
}
type GetAccountRuleSettings struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Settings struct {
				Rules []GetRuleSettings `json:"rules"`
			} `json:"settings"`
		} `json:"attributes"`
	} `json:"data"`
}
type RuleSettingExceptions struct {
	FilterTags []string `json:"filterTags,omitempty"`
	Resources  []string `json:"resources,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}
type MappingValues struct {
	Type   string      `json:"type,omitempty"`
	Name   string      `json:"name,omitempty"`
	Value  string      `json:"value,omitempty"`
	Values interface{} `json:"values,omitempty"`
}
type RuleSettingMapping struct {
	Values []*MappingValues `json:"values"`
}
type RuleSettingValues struct {
	Label   string `json:"label,omitempty"`
	Value   string `json:"value,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}
type RuleSettingMultipleObject struct {
	Value struct {
		EventName        string `json:"eventName,omitempty"`
		EventSource      string `json:"eventSource,omitempty"`
		UserIdentityType string `json:"userIdentityType,omitempty"`
	} `json:"value,omitempty"`
}
type RuleSettingExtra struct {
	Name      string      `json:"name,omitempty"`
	Type      string      `json:"type,omitempty"`
	Regions   *bool       `json:"regions,omitempty"`
	ValueKeys *[]string   `json:"valueKeys,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Values    interface{} `json:"values,omitempty"`
	Mappings  interface{} `json:"mappings,omitempty"`
}

type RuleSetting struct {
	// possible duplicated struct in ProfileSettings
	// for now will use different struct
	Enabled       bool                  `json:"enabled"`
	Exceptions    RuleSettingExceptions `json:"exceptions"`
	Id            string                `json:"id"`
	Provider      string                `json:"provider"`
	RiskLevel     string                `json:"riskLevel"`
	RuleExists    bool                  `json:"ruleExists"`
	ExtraSettings []RuleSettingExtra    `json:"extraSettings"`
}
type RuleSettingAttributes struct {
	Note        string      `json:"note"`
	RuleSetting RuleSetting `json:"ruleSetting"`
}
type AccountRuleSettings struct {
	Data struct {
		Id         string                `json:"id,omitempty"`
		Attributes RuleSettingAttributes `json:"attributes"`
	} `json:"data"`
}

type BotDisabledRegions struct {
	AfSouth1     bool `json:"af-south-1,omitempty"`
	ApSouth1     bool `json:"ap-south-1,omitempty"`
	EuWest3      bool `json:"eu-west-3,omitempty"`
	EuNorth1     bool `json:"eu-north-1,omitempty"`
	EuWest2      bool `json:"eu-west-2,omitempty"`
	EuSouth1     bool `json:"eu-south-1,omitempty"`
	EuWest1      bool `json:"eu-west-1,omitempty"`
	ApNorthEast3 bool `json:"ap-northeast-3,omitempty"`
	ApNorthEast2 bool `json:"ap-northeast-2,omitempty"`
	ApNorthEast1 bool `json:"ap-northeast-1,omitempty"`
	MeSouth1     bool `json:"me-south-1,omitempty"`
	SaEast1      bool `json:"sa-east-1,omitempty"`
	CaCentral1   bool `json:"ca-central-1,omitempty"`
	ApEast1      bool `json:"ap-east-1,omitempty"`
	ApSouthEast1 bool `json:"ap-southeast-1,omitempty"`
	ApSouthEast2 bool `json:"ap-southeast-2,omitempty"`
	EuCentral1   bool `json:"eu-central-1,omitempty"`
	UsEast1      bool `json:"us-east-1,omitempty"`
	UsEast2      bool `json:"us-east-2,omitempty"`
	UsWest1      bool `json:"us-west-1,omitempty"`
	UsWest2      bool `json:"us-west-2,omitempty"`
}

type AccountBot struct {
	Disabled        *bool               `json:"disabled"`
	Delay           *int                `json:"delay,omitempty"`
	DisabledRegions *BotDisabledRegions `json:"disabledRegions,omitempty"`
}

type AccountBotSettingsData struct {
	Attributes accountAtrributes `json:"attributes"`
	Type       string            `json:"type"`
	Id         string            `json:"id"`
}
type AccountBotSettingsReponse struct {
	Data []AccountBotSettingsData `json:"data"`
}

type AccountBotSettingsRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Settings AccountSettings `json:"settings,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

type AccountSettings struct {
	Bot *AccountBot `json:"bot,omitempty"`
}

type accountAtrributes struct {
	Name          string                `json:"name"`
	Environment   string                `json:"environment"`
	Tags          []string              `json:"tags"`
	//Access        accountAccess         `json:"access,omitempty"`
	Configuration *AccountConfiguration `json:"configuration,omitempty"`
	CoudType      string                `json:"cloud-type,omitempty"`
	CloudData     *CloudData            `json:"cloud-data,omitempty"`
	Settings      *AccountSettings      `json:"settings,omitempty"`
}

type createAccountAtrributes struct {
	Name          string                `json:"name"`
	Environment   string                `json:"environment"`
	Tags          []string              `json:"tags"`
	Access        accountAccess         `json:"access,omitempty"`
	Configuration *AccountConfiguration `json:"configuration,omitempty"`
	CoudType      string                `json:"cloud-type,omitempty"`
	CloudData     *CloudData            `json:"cloud-data,omitempty"`
	Settings      *AccountSettings      `json:"settings,omitempty"`
}

type accountData struct {
	Type       string            `json:"type,omitempty"`
	Attributes accountAtrributes `json:"attributes"`
}

type createAccountData struct {
	Type       string            `json:"type,omitempty"`
	Attributes createAccountAtrributes `json:"attributes"`
}

type AccountPayload struct {
	Data createAccountData `json:"data"`
}

type AccountResponse struct {
	Data struct {
		ID         string            `json:"id"`
		Attributes accountAtrributes `json:"attributes"`
	} `json:"data"`
}
type accountDetails struct {
	Data accountData `json:"data"`
	Type string      `json:"type"`
	Id   string      `json:"id"`
}
type accountAccessAndDetails struct {
	AccountDetails accountDetails         `json:"accountDetails"`
	AccessSettings accountData            `json:"accessDetails"`
	RuleSettings   GetAccountRuleSettings `json:"ruleSettings"`
}

type deleteResponse struct {
	Meta struct {
		Status string `json:"status"`
	} `json:"meta"`
}

type groupData struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Attributes struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	} `json:"attributes"`
}

type GroupDataList struct {
	Data []groupData `json:"data"`
}

type GroupDetails struct {
	Data groupData `json:"data"`
}

type UserAccountAccessList struct {
	Account string `json:"account"`
	Level   string `json:"level"`
}
type userAttributes struct {
	FirstName    string                  `json:"firstName,omitempty"`
	LastName     string                  `json:"lastName,omitempty"`
	ResFirstName string                  `json:"first-name,omitempty"`
	ResLastName  string                  `json:"last-name,omitempty"`
	Mfa          bool                    `json:"mfa,omitempty"`
	LastLogIn    int                     `json:"last-login-date,omitempty"`
	Email        string                  `json:"email"`
	Role         string                  `json:"role"`
	AccessList   []UserAccountAccessList `json:"accessList,omitempty"`
}
type userRelationships struct {
	AccountAccessList []UserAccountAccessList `json:"accountAccessList"`
}
type UserDetails struct {
	Data struct {
		Type          string            `json:"type,omitempty"`
		ID            string            `json:"id,omitempty"`
		Attributes    userAttributes    `json:"attributes,omitempty"`
		Relationships userRelationships `json:"relationships,omitempty"`
	} `json:"data"`
}

type UserAccessDetails struct {
	Data struct {
		Role       string                  `json:"role"`
		AccessList []UserAccountAccessList `json:"accessList"`
	} `json:"data"`
}

type ReportConfigFilter struct {
	Categories                 []string `json:"categories,omitempty"`
	ComplianceStandards        []string `json:"complianceStandards,omitempty"`
	FilterTags                 []string `json:"filterTags,omitempty"`
	Message                    bool     `json:"message,omitempty"`
	NewerThanDays              int      `json:"newerThanDays,omitempty"`
	OlderThanDays              int      `json:"olderThanDays,omitempty"`
	Providers                  []string `json:"providers,omitempty"`
	Regions                    []string `json:"regions,omitempty"`
	ReportComplianceStandardId string   `json:"reportComplianceStandardId,omitempty"`
	Resource                   string   `json:"resource,omitempty"`
	ResourceSearchMode         string   `json:"resourceSearchMode,omitempty"`
	ResourceTypes              []string `json:"resourceTypes,omitempty"`
	RiskLevels                 string   `json:"riskLevels,omitempty"`
	RuleIds                    []string `json:"ruleIds,omitempty"`
	Services                   []string `json:"services,omitempty"`
	Statuses                   []string `json:"statuses,omitempty"`
	Suppressed                 bool     `json:"suppressed"`
	SuppressedFilterMode       string   `json:"suppressedFilterMode"`
	Tags                       []string `json:"tags,omitempty"`
	Text                       string   `json:"text,omitempty"`
	WithChecks                 bool     `json:"withChecks"`
	WithoutChecks              bool     `json:"withoutChecks"`
}
type ReportConfiguration struct {
	Emails                []string           `json:"emails,omitempty"`
	Filter                ReportConfigFilter `json:"filter,omitempty"`
	Frequency             string             `json:"frequency,omitempty"`
	GenerateReportType    string             `json:"generateReportType"`
	IncludeChecks         bool               `json:"includeChecks"`
	Scheduled             bool               `json:"scheduled"`
	SendEmail             bool               `json:"sendEmail"`
	ShouldEmailIncludeCsv bool               `json:"shouldEmailIncludeCsv"`
	ShouldEmailIncludePdf bool               `json:"shouldEmailIncludePdf"`
	Title                 string             `json:"title,omitempty"`
	Tz                    string             `json:"tz,omitempty"`
}

type reportConfigRelationships struct {
	Organization struct {
		Data struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data,omitempty"`
	} `json:"organisation,omitempty"`

	Account struct {
		Data struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data,omitempty"`
	} `json:"account,omitempty"`

	Group struct {
		Data struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data,omitempty"`
	} `json:"group,omitempty"`
}

type ReportConfigDetails struct {
	Data struct {
		ID            string                    `json:"id,omitempty"`
		Type          string                    `json:"type,omitempty"`
		Relationships reportConfigRelationships `json:"relationships,omitempty"`
		Attributes    struct {
			AccountId     string              `json:"accountId,omitempty"`
			GroupId       string              `json:"groupId,omitempty"`
			Configuration ReportConfiguration `json:"configuration,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

type CommunicationConfiguration struct {
	Channel             string   `json:"channel,omitempty"`
	ChannelName         string   `json:"channelName,omitempty"`
	Users               []string `json:"users,omitempty"`
	DisplayExtraData    bool     `json:"displayExtraData,omitempty"`
	DisplayResource     bool     `json:"displayResource,omitempty"`
	DisplayTags         bool     `json:"displayTags,omitempty"`
	Url                 string   `json:"url,omitempty"`
	DisplayIntroducedBy bool     `json:"displayIntroducedBy,omitempty"`
	Arn                 string   `json:"arn,omitempty"`
	ServiceKey          string   `json:"serviceKey,omitempty"`
	ServiceName         string   `json:"serviceName,omitempty"`
	SecurityToken       string   `json:"securityToken,omitempty"`
}

type CommunicationFilter struct {
	Categories  []string `json:"categories,omitempty"`
	Compliances []string `json:"compliances,omitempty"`
	FilterTags  []string `json:"filterTags,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	RiskLevels  []string `json:"riskLevels,omitempty"`
	RuleIds     []string `json:"ruleIds,omitempty"`
	Services    []string `json:"services,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type CommunicaitonRelationshipsData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type CommunicaitonRelationships struct {
	Organisation struct {
		Data CommunicaitonRelationshipsData `json:"data,omitempty"`
	} `json:"organisation,omitempty"`

	Account struct {
		Data CommunicaitonRelationshipsData `json:"data,omitempty"`
	} `json:"account"`
}

type communicationData struct {
	ID            string                     `json:"id,omitempty"`
	Type          string                     `json:"type,omitempty"`
	Relationships CommunicaitonRelationships `json:"relationships,omitempty"`
	Attributes    struct {
		Channel       string                      `json:"channel,omitempty"`
		Enabled       bool                        `json:"enabled,omitempty"`
		Type          string                      `json:"type,omitempty"`
		Configuration *CommunicationConfiguration `json:"configuration,omitempty"`
		Filter        *CommunicationFilter        `json:"filter,omitempty"`
	} `json:"attributes"`
}
type CommunicationSettings struct {
	Data communicationData `json:"data"`
}

type CommunicationResponse struct {
	Data []communicationData `json:"data"`
}

type profileAttributes struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}
type RuleSettingsData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
type ProfileRelationships struct {
	RuleSettings struct {
		Data []RuleSettingsData `json:"data"`
	} `json:"ruleSettings,omitempty"`
}

type IncludedExceptions struct {
	FilterTags []string `json:"filterTags,omitempty"`
	Resources  []string `json:"resources,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}
type ProfileValues struct {
	Label   string `json:"label,omitempty"`
	Value   string `json:"value,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}
type IncludedExtraSettings struct {
	Countries bool             `json:"countries,omitempty"`
	Multiple  bool             `json:"multiple,omitempty"`
	Name      string           `json:"name,omitempty"`
	Regions   bool             `json:"regions,omitempty"`
	Type      string           `json:"type,omitempty"`
	Value     interface{}      `json:"value,omitempty"`
	Values    []*ProfileValues `json:"values,omitempty"`
}
type IncludedAttributes struct {
	Enabled       bool                    `json:"enabled"`
	Provider      string                  `json:"provider"`
	RiskLevel     string                  `json:"riskLevel"`
	Exceptions    *IncludedExceptions     `json:"exceptions,omitempty"`
	ExtraSettings []IncludedExtraSettings `json:"extraSettings,omitempty"`
}

type ProfileIncluded struct {
	ID         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Attributes IncludedAttributes `json:"attributes,omitempty"`
}

type ProfileSettings struct {
	Data struct {
		Attributes    profileAttributes    `json:"attributes,omitempty"`
		Relationships ProfileRelationships `json:"relationships,omitempty"`
		Type          string               `json:"type,omitempty"`
		ID            string               `json:"id,omitempty"`
	} `json:"data"`
	Included []ProfileIncluded `json:"included,omitempty"`
}
type ApplyProfileInclude struct {
	Exceptions bool `json:"exceptions"`
}
type ApplyProfileSettings struct {
	Meta struct {
		AccountIds []string             `json:"accountIds"`
		Mode       string               `json:"mode"`
		Notes      string               `json:"notes"`
		Types      []string             `json:"types"`
		Include    *ApplyProfileInclude `json:"include,omitempty"`
	} `json:"meta"`
}

type ApplyProfileResponse struct {
	Meta struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"meta"`
}
