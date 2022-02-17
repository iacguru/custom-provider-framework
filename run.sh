cwd=`pwd`
rm -rf providers/terraform-provider-github/examples/*terraform*
rm -rf providers/terraform-provider-github/examples/.terraform*
cd providers
make install
cd terraform-provider-github/examples
terraform init && terraform apply --auto-approve
cd $cwd
rm -rf examples/*terraform*
rm -rf examples/.terraform*
cd providers/examples
terraform init && terraform apply --auto-approve
cd $cwd