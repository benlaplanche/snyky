package main

ALLOW_LIST := [
 "express"
]

contains_package(arr, elem) = true {
  arr[_] = elem
} else = false { true }

deny[msg] {
	version := input.dependencies[name]
    
	not contains_package(ALLOW_LIST,name)
 	msg := sprintf("Package %v version %v is not on the allow list", [name, version])
}