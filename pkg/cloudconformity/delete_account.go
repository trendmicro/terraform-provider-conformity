package cloudconformity

// allows an ADMIN to delete the specified account
func (c *Client) DeleteAccount(accountId string) (*deleteResponse, error) {

	deleteAccountResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_account", accountId}, nil, "", &deleteAccountResponse)
	if err != nil {
		return nil, err
	}

	return &deleteAccountResponse, nil
}
