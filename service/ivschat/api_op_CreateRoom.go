// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a room that allows clients to connect and pass messages.
func (c *Client) CreateRoom(ctx context.Context, params *CreateRoomInput, optFns ...func(*Options)) (*CreateRoomOutput, error) {
	if params == nil {
		params = &CreateRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoom", params, optFns, c.addOperationCreateRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRoomInput struct {

	// Array of logging-configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers []string

	// Maximum number of characters in a single message. Messages are expected to be
	// UTF-8 encoded and this limit applies specifically to rune/code-point count, not
	// number of bytes. Default: 500.
	MaximumMessageLength int32

	// Maximum number of messages per second that can be sent to the room (by all
	// clients). Default: 10.
	MaximumMessageRatePerSecond int32

	// Configuration information for optional review of messages.
	MessageReviewHandler *types.MessageReviewHandler

	// Room name. The value does not need to be unique.
	Name *string

	// Tags to attach to the resource. Array of maps, each of the form string:string
	// (key:value). See Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS Chat has no constraints beyond what is documented
	// there.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRoomOutput struct {

	// Room ARN, assigned by the system.
	Arn *string

	// Time when the room was created. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	CreateTime *time.Time

	// Room ID, generated by the system. This is a relative identifier, the part of the
	// ARN that uniquely identifies the room.
	Id *string

	// Array of logging configurations attached to the room, from the request (if
	// specified).
	LoggingConfigurationIdentifiers []string

	// Maximum number of characters in a single message, from the request (if
	// specified).
	MaximumMessageLength int32

	// Maximum number of messages per second that can be sent to the room (by all
	// clients), from the request (if specified).
	MaximumMessageRatePerSecond int32

	// Configuration information for optional review of messages.
	MessageReviewHandler *types.MessageReviewHandler

	// Room name, from the request (if specified).
	Name *string

	// Tags attached to the resource, from the request (if specified).
	Tags map[string]string

	// Time of the room’s last update. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoom{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoom(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRoom(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "CreateRoom",
	}
}
