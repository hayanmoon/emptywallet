module github.com/vanclp/emptywallet/api

go 1.17

require (
	github.com/aws/aws-lambda-go v1.6.0
	github.com/aws/aws-sdk-go v1.40.54
	github.com/emptywallet/api/sevice v0.0.0-unpublished
)

replace github.com/emptywallet/api/service v0.0.0-unpublished => /Users/vanclp/emptywallet/api/hello/service/

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)
