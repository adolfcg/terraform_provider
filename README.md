
### 1. Configuration File   $HOME/.terraformrc 

---

```
plugin_cache_dir   = "$HOME/.terraform.d/plugin-cache" 
disable_checkpoint = true
```

#### Mocking resource structure sample

1. style = americano/capucchino/latte
2. size = large/medium/small
3. place = togo/here


### 2. Create the required Files


```
cd $HOME/go/src
mkdir sirena_provider
```

Required source files for custom provider are:

- main.go
- provider.go
- resource_server.go

The code layout looks like this:

```
% pwd

/Users/acabr/dev/golang/src/sirena_provider

tree .
.
├── main.go
├── provider.go
├── resource_server.go
```


### 3. Build the Custom Provider code

```
# creating new go.mod: module sirena_provider
go mod init 

#formats Go source code
go fmt  

#Download all the dependencies that are required in your source files
go mod tidy  

#builds the provider			
go build -o  terraform-provider-order
```




### 4. Copy the provider to the plugin directory 

```
mkdir -p ~/.terraform.d/plugins/sirena.com/tutorial/order/1.0.0/darwin_arm64 
cp terraform-provider-order ~/.terraform.d/plugins/sirena.com/tutorial/order/1.0.0/darwin_arm64 
```


#### Reset comands [if required]

```
rm -r .terraform
rm -r .terraform.lock.hcl
rm -r terraform.tfstate*
```


### 5. Create the terraform files 

- main.tf
- version.tf (optional)

### 6. Test the Provider and output values

```
terraform init 
terraform plan 
terraform apply -auto-approve
terraform show -json | jq .
```
