// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeevents

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/personalizeevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records user interaction event data. For more information see Recording Events
// (https://docs.aws.amazon.com/personalize/latest/dg/recording-events.html).
func (c *Client) PutEvents(ctx context.Context, params *PutEventsInput, optFns ...func(*Options)) (*PutEventsOutput, error) {
	if params == nil {
		params = &PutEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEvents", params, optFns, c.addOperationPutEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventsInput struct {

	// A list of event data from the session.
	//
	// This member is required.
	EventList []types.Event

	// The session ID associated with the user's visit. Your application generates the
	// sessionId when a user first visits your website or uses your application. Amazon
	// Personalize uses the sessionId to associate events with the user before they log
	// in. For more information, see Recording Events
	// (https://docs.aws.amazon.com/personalize/latest/dg/recording-events.html).
	//
	// This member is required.
	SessionId *string

	// The tracking ID for the event. The ID is generated by a call to the
	// CreateEventTracker
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html)
	// API.
	//
	// This member is required.
	TrackingId *string

	// The user associated with the event.
	UserId *string

	noSmithyDocumentSerde
}

type PutEventsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEvents{}, middleware.After)
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
	if err = addOpPutEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "PutEvents",
	}
}
