// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// To get the description of a registered consumer, provide the ARN of the
// consumer. Alternatively, you can provide the ARN of the data stream and the name
// you gave the consumer when you registered it. You may also provide all three
// parameters, as long as they don't conflict with each other. If you don't know
// the name or ARN of the consumer that you want to describe, you can use the
// ListStreamConsumers operation to get a list of the descriptions of all the
// consumers that are currently registered with a given data stream. This operation
// has a limit of 20 transactions per second per stream.
func (c *Client) DescribeStreamConsumer(ctx context.Context, params *DescribeStreamConsumerInput, optFns ...func(*Options)) (*DescribeStreamConsumerOutput, error) {
	if params == nil {
		params = &DescribeStreamConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStreamConsumer", params, optFns, c.addOperationDescribeStreamConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamConsumerInput struct {

	// The ARN returned by Kinesis Data Streams when you registered the consumer.
	ConsumerARN *string

	// The name that you gave to the consumer.
	ConsumerName *string

	// The ARN of the Kinesis data stream that the consumer is registered with. For
	// more information, see Amazon Resource Names (ARNs) and Amazon Web Services
	// Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	StreamARN *string

	noSmithyDocumentSerde
}

type DescribeStreamConsumerOutput struct {

	// An object that represents the details of the consumer.
	//
	// This member is required.
	ConsumerDescription *types.ConsumerDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamConsumer{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamConsumer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStreamConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DescribeStreamConsumer",
	}
}
