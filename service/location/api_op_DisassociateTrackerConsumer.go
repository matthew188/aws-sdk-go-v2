// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between a tracker resource and a geofence collection.
// Once you unlink a tracker resource from a geofence collection, the tracker
// positions will no longer be automatically evaluated against geofences.
func (c *Client) DisassociateTrackerConsumer(ctx context.Context, params *DisassociateTrackerConsumerInput, optFns ...func(*Options)) (*DisassociateTrackerConsumerOutput, error) {
	if params == nil {
		params = &DisassociateTrackerConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTrackerConsumer", params, optFns, c.addOperationDisassociateTrackerConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTrackerConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTrackerConsumerInput struct {

	// The Amazon Resource Name (ARN) for the geofence collection to be disassociated
	// from the tracker resource. Used when you need to specify a resource across all
	// Amazon Web Services.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer
	//
	// This member is required.
	ConsumerArn *string

	// The name of the tracker resource to be dissociated from the consumer.
	//
	// This member is required.
	TrackerName *string

	noSmithyDocumentSerde
}

type DisassociateTrackerConsumerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateTrackerConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateTrackerConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateTrackerConsumer{}, middleware.After)
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
	if err = addEndpointPrefix_opDisassociateTrackerConsumerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDisassociateTrackerConsumerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTrackerConsumer(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDisassociateTrackerConsumerMiddleware struct {
}

func (*endpointPrefix_opDisassociateTrackerConsumerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDisassociateTrackerConsumerMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDisassociateTrackerConsumerMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDisassociateTrackerConsumerMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateTrackerConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "DisassociateTrackerConsumer",
	}
}
