locals {
#  environment_vars = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  region = "us-west-2"
  meaning = 42
}

#include "root" {
#  path = find_in_parent_folders()
#}