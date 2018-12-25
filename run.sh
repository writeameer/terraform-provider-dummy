#!/usr/bin/env bash

export TF_LOG=DEBUG
terraform init
terraform plan
terraform apply -auto-approve

