#!/bin/bash
for n in {1..500}; do
  cp terraform.tf 500/terraform-${n}.tf
done