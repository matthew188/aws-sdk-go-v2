module github.com/matthew188/aws-sdk-go-v2/internal/configsources/configtesting

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2/config v1.18.19
	github.com/matthew188/aws-sdk-go-v2/internal/configsources v1.1.31
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../../

replace github.com/matthew188/aws-sdk-go-v2/config => ../../../config/

replace github.com/matthew188/aws-sdk-go-v2/credentials => ../../../credentials/

replace github.com/matthew188/aws-sdk-go-v2/feature/ec2/imds => ../../../feature/ec2/imds/

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/internal/ini => ../../../internal/ini/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/matthew188/aws-sdk-go-v2/service/sso => ../../../service/sso/

replace github.com/matthew188/aws-sdk-go-v2/service/ssooidc => ../../../service/ssooidc/

replace github.com/matthew188/aws-sdk-go-v2/service/sts => ../../../service/sts/
