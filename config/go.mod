module github.com/matthew188/aws-sdk-go-v2/config

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/credentials v1.13.18
	github.com/matthew188/aws-sdk-go-v2/feature/ec2/imds v1.13.1
	github.com/matthew188/aws-sdk-go-v2/internal/ini v1.3.32
	github.com/matthew188/aws-sdk-go-v2/service/sso v1.12.6
	github.com/matthew188/aws-sdk-go-v2/service/ssooidc v1.14.6
	github.com/matthew188/aws-sdk-go-v2/service/sts v1.18.7
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/matthew188/aws-sdk-go-v2 => ../

replace github.com/matthew188/aws-sdk-go-v2/credentials => ../credentials/

replace github.com/matthew188/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/internal/ini => ../internal/ini/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url => ../service/internal/presigned-url/

replace github.com/matthew188/aws-sdk-go-v2/service/sso => ../service/sso/

replace github.com/matthew188/aws-sdk-go-v2/service/ssooidc => ../service/ssooidc/

replace github.com/matthew188/aws-sdk-go-v2/service/sts => ../service/sts/
