#!/usr/bin/env bash

export TF_LOG=INFO
terraform init
terraform plan
terraform apply -auto-approve

