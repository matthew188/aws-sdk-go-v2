module github.com/matthew188/aws-sdk-go-v2/feature/s3/manager

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/config v1.18.19
	github.com/matthew188/aws-sdk-go-v2/service/s3 v1.31.1
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../../

replace github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream => ../../../aws/protocol/eventstream/

replace github.com/matthew188/aws-sdk-go-v2/config => ../../../config/

replace github.com/matthew188/aws-sdk-go-v2/credentials => ../../../credentials/

replace github.com/matthew188/aws-sdk-go-v2/feature/ec2/imds => ../../../feature/ec2/imds/

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/internal/ini => ../../../internal/ini/

replace github.com/matthew188/aws-sdk-go-v2/internal/v4a => ../../../internal/v4a/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/checksum => ../../../service/internal/checksum/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/s3shared => ../../../service/internal/s3shared/

replace github.com/matthew188/aws-sdk-go-v2/service/s3 => ../../../service/s3/

replace github.com/matthew188/aws-sdk-go-v2/service/sso => ../../../service/sso/

replace github.com/matthew188/aws-sdk-go-v2/service/ssooidc => ../../../service/ssooidc/

replace github.com/matthew188/aws-sdk-go-v2/service/sts => ../../../service/sts/
