package cloudconformity

import (
	"fmt"
	"strings"
)

//get both Account access settings and details
func (c *Client) GetAccount(accountId string) (*accountAccessAndDetails, error) {
	accountAccessAndDetails := &accountAccessAndDetails{}
	accountDetails, err := c.GetAccountDetails(accountId)
	if err != nil {
		return nil, err
	}
	accessSettings, err := c.GetAccountAccessSettings(accountId)
	if err != nil {
		return nil, err
	}
	ruleSettings, err := c.GetAccountRuleSettings(accountId)
	if err != nil {
		return nil, err
	}

	accountAccessAndDetails.AccountDetails = *accountDetails
	accountAccessAndDetails.AccessSettings = *accessSettings
	accountAccessAndDetails.RuleSettings = *ruleSettings

	return accountAccessAndDetails, nil

}

//get GCP Account settings and details
func (c *Client) GetGCPAccount(accountId string) (*accountAccessAndDetails, error) {
	accountAccessAndDetails := &accountAccessAndDetails{}
	accountDetails, err := c.GetAccountDetails(accountId)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	ruleSettings, err := c.GetAccountRuleSettings(accountId)
	if err != nil {
		return nil, err
	}

	accountAccessAndDetails.AccountDetails = *accountDetails
	accountAccessAndDetails.RuleSettings = *ruleSettings

	return accountAccessAndDetails, nil

}

// allows ADMIN users to get the current setting Conformity uses to access the specified account
// return the role arn and externalId
func (c *Client) GetAccountAccessSettings(accountId string) (*accountData, error) {
	accountData := &accountData{}
	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/accounts/%s/access", accountId), nil, "", accountData)
	if err != nil {
		return nil, err
	}
	return accountData, nil
}

// allows you to get the details of the specified account
// return the account name and evironment
func (c *Client) GetAccountDetails(accountId string) (*accountDetails, error) {
	accountDetails := &accountDetails{}
	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/accounts/%s", accountId), nil, "", accountDetails)
	if err != nil {
		return nil, err
	}
	return accountDetails, nil
}

// allows you to get rule settings for all configured rules of the specified account
func (c *Client) GetAccountRuleSettings(accountId string) (*GetAccountRuleSettings, error) {
	ruleSettings := &GetAccountRuleSettings{}
	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/accounts/%s/settings/rules", accountId), nil, "", ruleSettings)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return nil, err
	}
	return ruleSettings, nil
}
