// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/matthew188/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/matthew188/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a metrics configuration (specified by the metrics configuration ID) from
// the bucket. Note that this doesn't include the daily storage metrics. To use
// this operation, you must have permissions to perform the
// s3:GetMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
// For information about CloudWatch request metrics for Amazon S3, see Monitoring
// Metrics with Amazon CloudWatch
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
// The following operations are related to GetBucketMetricsConfiguration:
//
// *
// PutBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html)
//
// *
// DeleteBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html)
//
// *
// ListBucketMetricsConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html)
//
// *
// Monitoring Metrics with Amazon CloudWatch
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html)
func (c *Client) GetBucketMetricsConfiguration(ctx context.Context, params *GetBucketMetricsConfigurationInput, optFns ...func(*Options)) (*GetBucketMetricsConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketMetricsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketMetricsConfiguration", params, optFns, c.addOperationGetBucketMetricsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketMetricsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketMetricsConfigurationInput struct {

	// The name of the bucket containing the metrics configuration to retrieve.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the metrics configuration. The ID has a 64 character
	// limit and can only contain letters, numbers, periods, dashes, and underscores.
	//
	// This member is required.
	Id *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type GetBucketMetricsConfigurationOutput struct {

	// Specifies the metrics configuration.
	MetricsConfiguration *types.MetricsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketMetricsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketMetricsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketMetricsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketMetricsConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketMetricsConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetBucketMetricsConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketMetricsConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketMetricsConfiguration",
	}
}

// getGetBucketMetricsConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getGetBucketMetricsConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketMetricsConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketMetricsConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketMetricsConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
