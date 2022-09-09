package cloudconformity

import (
	"encoding/json"

	"strings"
)

func (c *Client) CreateAzureActiveDirectory(AzureActiveDirectory ActiveAzureDirectory) (string, error) {

	azureActiveDirectoryDetails := AzureActiveDirectoryResponse{}
	rb, err := json.Marshal(AzureActiveDirectory)
	if err != nil {
		return "", err
	}
	_, err = c.ClientRequest(Post{}, "/azure/active-directories", strings.NewReader(string(rb)), "", &azureActiveDirectoryDetails)
	if err != nil {

		return "", err
	}

	return azureActiveDirectoryDetails.Data.ID, nil

}
