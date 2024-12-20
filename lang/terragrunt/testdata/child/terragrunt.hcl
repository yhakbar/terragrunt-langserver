include "root" {
  path = find_in_parent_folders()
  expose = true
}

locals {
  name = "dummy"
}

terraform {
  source = "tfr:///terraform-aws-modules/dynamodb-table/aws?version=3.3.0"
}

inputs = {
  meaning = include.root.locals.meaning
}