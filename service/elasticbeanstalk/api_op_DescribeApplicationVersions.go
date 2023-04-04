// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a list of application versions.
func (c *Client) DescribeApplicationVersions(ctx context.Context, params *DescribeApplicationVersionsInput, optFns ...func(*Options)) (*DescribeApplicationVersionsOutput, error) {
	if params == nil {
		params = &DescribeApplicationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationVersions", params, optFns, c.addOperationDescribeApplicationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe application versions.
type DescribeApplicationVersionsInput struct {

	// Specify an application name to show only application versions for that
	// application.
	ApplicationName *string

	// For a paginated request. Specify a maximum number of application versions to
	// include in each response. If no MaxRecords is specified, all available
	// application versions are retrieved in a single response.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	// Specify a version label to show a specific application version.
	VersionLabels []string

	noSmithyDocumentSerde
}

// Result message wrapping a list of application version descriptions.
type DescribeApplicationVersionsOutput struct {

	// List of ApplicationVersionDescription objects sorted in order of creation.
	ApplicationVersions []types.ApplicationVersionDescription

	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeApplicationVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeApplicationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeApplicationVersions",
	}
}
