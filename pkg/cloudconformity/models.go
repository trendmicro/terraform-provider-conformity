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
type accountKeys struct {
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}
type accountAccess struct {
	Keys              accountKeys `json:"keys,omitempty"`
	Type              string      `json:"type,omitempty"`
	SubscriptionId    string      `json:"subscriptionId,omitempty"`
	ActiveDirectoryId string      `json:"activeDirectoryId,omitempty"`
}
type accountConfiguration struct {
	ExternalId string `json:"externalId,omitempty"`
	RoleArn    string `json:"roleArn,omitempty"`
}
type cloudData struct {
	Azure struct {
		SubscriptionId string `json:"subscriptionId"`
	} `json:"azure"`
}
type accountAtrributes struct {
	Name          string               `json:"name"`
	Environment   string               `json:"environment"`
	Tags          []string             `json:"tags"`
	Access        accountAccess        `json:"access,omitempty"`
	Configuration accountConfiguration `json:"configuration,omitempty"`
	CoudType      string               `json:"cloud-type,omitempty"`
	CloudData     cloudData            `json:"cloud-data,omitempty"`
}

type accountData struct {
	Type       string            `json:"type,omitempty"`
	Attributes accountAtrributes `json:"attributes"`
}

type AccountPayload struct {
	Data accountData `json:"data"`
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
	AccountDetails accountDetails `json:"accountDetails"`
	AccessSettings accountData    `json:"accessDetails"`
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
