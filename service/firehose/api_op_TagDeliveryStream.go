// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates tags for the specified delivery stream. A tag is a key-value
// pair that you can define and assign to Amazon Web Services resources. If you
// specify a tag that already exists, the tag value is replaced with the value that
// you specify in the request. Tags are metadata. For example, you can add friendly
// names and descriptions or other types of information that can help you
// distinguish the delivery stream. For more information about tags, see Using Cost
// Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the Amazon Web Services Billing and Cost Management User Guide. Each delivery
// stream can have up to 50 tags. This operation has a limit of five transactions
// per second per account.
func (c *Client) TagDeliveryStream(ctx context.Context, params *TagDeliveryStreamInput, optFns ...func(*Options)) (*TagDeliveryStreamOutput, error) {
	if params == nil {
		params = &TagDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagDeliveryStream", params, optFns, c.addOperationTagDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagDeliveryStreamInput struct {

	// The name of the delivery stream to which you want to add the tags.
	//
	// This member is required.
	DeliveryStreamName *string

	// A set of key-value pairs to use to create the tags.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagDeliveryStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagDeliveryStream{}, middleware.After)
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
	if err = addOpTagDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagDeliveryStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "TagDeliveryStream",
	}
}
