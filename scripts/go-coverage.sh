#!/bin/sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "Total coverage: $COVERAGE"
[ $(echo "$COVERAGE >= 80.0" | bc) -eq 1 ] # Ensure coverage is above 80%
