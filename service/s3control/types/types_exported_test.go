// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/s3control/types"
)

func ExampleJobManifestGenerator_outputUsage() {
	var union types.JobManifestGenerator
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.JobManifestGeneratorMemberS3JobManifestGenerator:
		_ = v.Value // Value is types.S3JobManifestGenerator

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3JobManifestGenerator

func ExampleObjectLambdaContentTransformation_outputUsage() {
	var union types.ObjectLambdaContentTransformation
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ObjectLambdaContentTransformationMemberAwsLambda:
		_ = v.Value // Value is types.AwsLambdaTransformation

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AwsLambdaTransformation
