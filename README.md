# Conformity Terraform Provider

## How to set up local machine:

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-conformity
```
#### 2. Install dependencies:
```sh
go mod tidy
go mod vendor
```
#### 3. Create the Artifact:
```sh
make install
```
#### 4. Now, you can test terraform code:
```sh
cd example/path-to-main/
terraform init
terraform apply
```
Notes:<br> 
* for your own config, create a file name `terraform.tfvars`
* add the following:
```sh
region  = "region"
apikey  = "apikey"
```


 Turn on debug:
```sh
export TF_LOG_CORE=TRACE
export TF_LOG_PROVIDER=TRACE
```

## How to protect API keys

#### 1. with file

Create a file name `terraform.tfvars` and add all necessary variables here

Ensure `terraform.tfvars` is included in `.gitignore` so these secrets are not accidentally pushed to a remote git repository.

#### 2. with environment variables

Terraform provides a way of reading variables from the environment: https://www.terraform.io/docs/cli/config/environment-variables.html#tf_var_name

## Updating documentation
Use the [Doc Preview Tool](https://registry.terraform.io/tools/doc-preview) to understand how the markdown will look once released. The [Provider Documentation](https://developer.hashicorp.com/terraform/registry/providers/docs) can also provide further guidance.

## How to release
### Steps
#### 1. Go to terraform provider GitHub: https://github.com/trendmicro/terraform-provider-conformity/releases

#### 2. Click "Draft a new release" button

#### 3. Click "Choose a Tag" dropdown, provide tag with value “xxx”, then select "+ Create new Tag : xxx on publish" popup item below.

#### 4. Choose the main branch as "Target"

#### 5. Fill the release title “xxx”

#### 6. Add the released changes to the description. *Do avoid Jira Ticket's IDs as those are not publicly visible.*

#### 7. Click "Publish release" button at the bottom.

### Check the release
After releasing, a webhook will be sent to Terraform registry automatically.
Within about 10 minutes https://registry.terraform.io/providers/trendmicro/conformity/latest/docs  should be updated with the new release from Github.