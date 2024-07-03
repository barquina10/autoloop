terraform {
  required_providers {
    autoloop = {
      source  = "barquina10/autoloop"
      version = "0.1.0"
    }
  }
}

provider "autoloop" {}

resource "autoloop_command" "example" {
  command = "aws s3 ls"
}

output "aws_command_output" {
  value = autoloop_command.example.output
}