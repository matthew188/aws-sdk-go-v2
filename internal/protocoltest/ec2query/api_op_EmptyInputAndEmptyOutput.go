// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response members. While this should be rare, code generators must
// support this.
func (c *Client) EmptyInputAndEmptyOutput(ctx context.Context, params *EmptyInputAndEmptyOutputInput, optFns ...func(*Options)) (*EmptyInputAndEmptyOutputOutput, error) {
	if params == nil {
		params = &EmptyInputAndEmptyOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EmptyInputAndEmptyOutput", params, optFns, c.addOperationEmptyInputAndEmptyOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EmptyInputAndEmptyOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EmptyInputAndEmptyOutputInput struct {
	noSmithyDocumentSerde
}

type EmptyInputAndEmptyOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEmptyInputAndEmptyOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEmptyInputAndEmptyOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEmptyInputAndEmptyOutput{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EmptyInputAndEmptyOutput",
	}
}
