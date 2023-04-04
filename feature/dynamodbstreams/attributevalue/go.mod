module github.com/matthew188/aws-sdk-go-v2/feature/dynamodbstreams/attributevalue

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/service/dynamodb v1.19.2
	github.com/matthew188/aws-sdk-go-v2/service/dynamodbstreams v1.14.7
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../../

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/matthew188/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
