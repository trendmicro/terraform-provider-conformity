package cloudconformity

import (
	"fmt"
)

//get both Account access settings and details
func (c *Client) GetAccount(accountId string) (*accountAccessAndDetails, error) {
	accountAccessAndDetails := &accountAccessAndDetails{}
	accountDetails, err := c.GetAccountDetails(accountId)
	if err != nil {
		return nil, err
	}
	AccessSettings, err := c.GetAccountAccessSettings(accountId)
	if err != nil {
		return nil, err
	}
	accountAccessAndDetails.AccountDetails = *accountDetails
	accountAccessAndDetails.AccessSettings = *AccessSettings
	return accountAccessAndDetails, nil

}

// allows ADMIN users to get the current setting Cloud Conformity uses to access the specified account
// return the role arn and externalId
func (c *Client) GetAccountAccessSettings(accountId string) (*accountData, error) {
	accountData := &accountData{}
	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/v1/accounts/%s/access", accountId), nil, accountData)
	if err != nil {
		return nil, err
	}
	return accountData, nil
}

// allows you to get the details of the specified account
// return the account name and evironment
func (c *Client) GetAccountDetails(accountId string) (*accountDetails, error) {
	accountDetails := &accountDetails{}
	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/v1/accounts/%s", accountId), nil, accountDetails)
	if err != nil {
		return nil, err
	}
	return accountDetails, nil
}
