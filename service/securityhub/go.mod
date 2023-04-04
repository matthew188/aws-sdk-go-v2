module github.com/matthew188/aws-sdk-go-v2/service/securityhub

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/internal/configsources v1.1.31
	github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 v2.4.25
	github.com/aws/smithy-go v1.13.5
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
