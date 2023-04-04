// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Lambda function URL. When you delete a function URL, you can't recover
// it. Creating a new function URL results in a different URL address.
func (c *Client) DeleteFunctionUrlConfig(ctx context.Context, params *DeleteFunctionUrlConfigInput, optFns ...func(*Options)) (*DeleteFunctionUrlConfigOutput, error) {
	if params == nil {
		params = &DeleteFunctionUrlConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFunctionUrlConfig", params, optFns, c.addOperationDeleteFunctionUrlConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFunctionUrlConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFunctionUrlConfigInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name – my-function.
	//
	// *
	// Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN – 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The alias name.
	Qualifier *string

	noSmithyDocumentSerde
}

type DeleteFunctionUrlConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFunctionUrlConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFunctionUrlConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFunctionUrlConfig{}, middleware.After)
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
	if err = addOpDeleteFunctionUrlConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFunctionUrlConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFunctionUrlConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "DeleteFunctionUrlConfig",
	}
}
