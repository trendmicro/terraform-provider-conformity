## How to set up local machine:

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-cloudconformity
```
#### 2. Install dependencies:
```sh
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
