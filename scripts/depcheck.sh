#!/bin/sh
go mod tidy
git diff --exit-code go.mod go.sum
