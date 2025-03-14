// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of the specified resource in the specified stack. For
// deleted stacks, DescribeStackResource returns resource information for up to 90
// days after the stack has been deleted.
func (c *Client) DescribeStackResource(ctx context.Context, params *DescribeStackResourceInput, optFns ...func(*Options)) (*DescribeStackResourceOutput, error) {
	if params == nil {
		params = &DescribeStackResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackResource", params, optFns, c.addOperationDescribeStackResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStackResource action.
type DescribeStackResourceInput struct {

	// The logical name of the resource as specified in the template. Default: There is
	// no default value.
	//
	// This member is required.
	LogicalResourceId *string

	// The name or the unique stack ID that's associated with the stack, which aren't
	// always interchangeable:
	//
	// * Running stacks: You can specify either the stack's
	// name or its unique stack ID.
	//
	// * Deleted stacks: You must specify the unique
	// stack ID.
	//
	// Default: There is no default value.
	//
	// This member is required.
	StackName *string

	noSmithyDocumentSerde
}

// The output for a DescribeStackResource action.
type DescribeStackResourceOutput struct {

	// A StackResourceDetail structure containing the description of the specified
	// resource in the specified stack.
	StackResourceDetail *types.StackResourceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStackResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResource{}, middleware.After)
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
	if err = addOpDescribeStackResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStackResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackResource",
	}
}
