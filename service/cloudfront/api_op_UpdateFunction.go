// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a CloudFront function. You can update a function's code or the comment
// that describes the function. You cannot update a function's name. To update a
// function, you provide the function's name and version (ETag value) along with
// the updated function code. To get the name and version, you can use
// ListFunctions and DescribeFunction.
func (c *Client) UpdateFunction(ctx context.Context, params *UpdateFunctionInput, optFns ...func(*Options)) (*UpdateFunctionOutput, error) {
	if params == nil {
		params = &UpdateFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunction", params, optFns, c.addOperationUpdateFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionInput struct {

	// The function code. For more information about writing a CloudFront function, see
	// Writing function code for CloudFront Functions
	// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html)
	// in the Amazon CloudFront Developer Guide.
	//
	// This member is required.
	FunctionCode []byte

	// Configuration information about the function.
	//
	// This member is required.
	FunctionConfig *types.FunctionConfig

	// The current version (ETag value) of the function that you are updating, which
	// you can get using DescribeFunction.
	//
	// This member is required.
	IfMatch *string

	// The name of the function that you are updating.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type UpdateFunctionOutput struct {

	// The version identifier for the current version of the CloudFront function.
	ETag *string

	// Contains configuration information and metadata about a CloudFront function.
	FunctionSummary *types.FunctionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateFunction{}, middleware.After)
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
	if err = addOpUpdateFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateFunction",
	}
}
