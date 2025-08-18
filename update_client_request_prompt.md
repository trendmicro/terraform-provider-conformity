# Update ClientRequest Function Invocations

**Update all `ClientRequest` function invocations in the `pkg/cloudconformity` folder to use the correct signature. The `url_params` parameter should be an array where:**
- **First element**: A string key that corresponds to a key in the `LegacyFunctions` map to retrieve the URL template
- **Remaining elements**: Parameters used to format the URL template using `fmt.Sprintf`

**Ensure all calls follow this pattern:**
```go
client.ClientRequest(method, []interface{}{"function_key", param1, param2, ...}, payload, rawQuery, result)
```

## Current ClientRequest Signature
```go
func (client *Client) ClientRequest(m method, url_params []interface{}, payload io.Reader, rawQuery string, result interface{}) ([]byte, error)
```

## Available Function Keys
The function keys are defined in the `LegacyFunctions` and `V1Functions` maps in `/pkg/cloudconformity/conformity.go`. Examples include:
- `"get_account"` - requires 1 parameter (account ID)
- `"update_account_rule_settings"` - requires 2 parameters (account ID, rule ID)
- `"get_current_user"` - requires 0 parameters
- `"create_aws_account"` - requires 0 parameters

## Example Transformations
```go
// Before (incorrect)
client.ClientRequest(Get{}, fmt.Sprintf("/accounts/%s", accountId), nil, "", &result)

// After (correct)
client.ClientRequest(Get{}, []interface{}{"get_account", accountId}, nil, "", &result)
```

```go
// Before (incorrect)
client.ClientRequest(Put{}, "/accounts/", payload, "", &result)

// After (correct)
client.ClientRequest(Put{}, []interface{}{"create_aws_account"}, payload, "", &result)