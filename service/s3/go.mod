module github.com/matthew188/aws-sdk-go-v2/service/s3

go 1.15

require (
	github.com/matthew188/aws-sdk-go-v2 v1.17.7
	github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10
	github.com/matthew188/aws-sdk-go-v2/internal/configsources v1.1.31
	github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 v2.4.25
	github.com/matthew188/aws-sdk-go-v2/internal/v4a v1.0.23
	github.com/matthew188/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11
	github.com/matthew188/aws-sdk-go-v2/service/internal/checksum v1.1.26
	github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url v1.9.25
	github.com/matthew188/aws-sdk-go-v2/service/internal/s3shared v1.14.0
	github.com/matthew188/aws-sdk-go-v2/service/s3 v1.31.1
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/matthew188/aws-sdk-go-v2 => ../../

replace github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/matthew188/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/matthew188/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/matthew188/aws-sdk-go-v2/internal/v4a => ../../internal/v4a/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/checksum => ../../service/internal/checksum/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/

replace github.com/matthew188/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
