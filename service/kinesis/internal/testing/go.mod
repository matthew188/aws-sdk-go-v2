module github.com/matthew188/aws-sdk-go-v2/service/kinesis/internal/testing

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10
	github.com/matthew188/aws-sdk-go-v2/service/internal/eventstreamtesting v1.0.57
	github.com/matthew188/aws-sdk-go-v2/service/kinesis v1.17.8
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../../../

replace github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream => ../../../../aws/protocol/eventstream/

replace github.com/matthew188/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/matthew188/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/eventstreamtesting => ../../../../service/internal/eventstreamtesting/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/matthew188/aws-sdk-go-v2/service/kinesis => ../../../../service/kinesis/

replace github.com/matthew188/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/matthew188/aws-sdk-go-v2/service/ssooidc => ../../../../service/ssooidc/

replace github.com/matthew188/aws-sdk-go-v2/service/sts => ../../../../service/sts/
