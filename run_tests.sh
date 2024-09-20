#!/bin/bash

# Run tests for add_numbers_test.go
echo "Running tests for add_numbers_test.go"
go test -run TestAddNumbers

# Run tests for main_test.go
echo "Running tests for main_test.go"
go test -run TestMultiplyNumbers