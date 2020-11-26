package main

deny[msg] {
not input.kind == "apiVersion"
  msg := "apiVersion must be specified" 
}