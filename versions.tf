terraform {
  required_version = ">= 1.0.0"

  required_providers {
    autoloop = {
      source  = "barquina10/terraform-provider-autoloop"
      version = "0.1.0"
    }
  }
}