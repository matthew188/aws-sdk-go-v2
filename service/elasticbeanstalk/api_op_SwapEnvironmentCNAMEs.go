// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Swaps the CNAMEs of two environments.
func (c *Client) SwapEnvironmentCNAMEs(ctx context.Context, params *SwapEnvironmentCNAMEsInput, optFns ...func(*Options)) (*SwapEnvironmentCNAMEsOutput, error) {
	if params == nil {
		params = &SwapEnvironmentCNAMEsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SwapEnvironmentCNAMEs", params, optFns, c.addOperationSwapEnvironmentCNAMEsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SwapEnvironmentCNAMEsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Swaps the CNAMEs of two environments.
type SwapEnvironmentCNAMEsInput struct {

	// The ID of the destination environment. Condition: You must specify at least the
	// DestinationEnvironmentID or the DestinationEnvironmentName. You may also specify
	// both. You must specify the SourceEnvironmentId with the
	// DestinationEnvironmentId.
	DestinationEnvironmentId *string

	// The name of the destination environment. Condition: You must specify at least
	// the DestinationEnvironmentID or the DestinationEnvironmentName. You may also
	// specify both. You must specify the SourceEnvironmentName with the
	// DestinationEnvironmentName.
	DestinationEnvironmentName *string

	// The ID of the source environment. Condition: You must specify at least the
	// SourceEnvironmentID or the SourceEnvironmentName. You may also specify both. If
	// you specify the SourceEnvironmentId, you must specify the
	// DestinationEnvironmentId.
	SourceEnvironmentId *string

	// The name of the source environment. Condition: You must specify at least the
	// SourceEnvironmentID or the SourceEnvironmentName. You may also specify both. If
	// you specify the SourceEnvironmentName, you must specify the
	// DestinationEnvironmentName.
	SourceEnvironmentName *string

	noSmithyDocumentSerde
}

type SwapEnvironmentCNAMEsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSwapEnvironmentCNAMEsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSwapEnvironmentCNAMEs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSwapEnvironmentCNAMEs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSwapEnvironmentCNAMEs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSwapEnvironmentCNAMEs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "SwapEnvironmentCNAMEs",
	}
}
