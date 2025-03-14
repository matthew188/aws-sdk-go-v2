// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a member from a channel. The x-amz-chime-bearer request header is
// mandatory. Use the AppInstanceUserArn of the user that makes the API call as the
// value in the header.
func (c *Client) DeleteChannelMembership(ctx context.Context, params *DeleteChannelMembershipInput, optFns ...func(*Options)) (*DeleteChannelMembershipOutput, error) {
	if params == nil {
		params = &DeleteChannelMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChannelMembership", params, optFns, c.addOperationDeleteChannelMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChannelMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteChannelMembershipInput struct {

	// The ARN of the channel from which you want to remove the user.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The AppInstanceUserArn of the member that you're removing from the channel.
	//
	// This member is required.
	MemberArn *string

	// The ID of the SubChannel in the request. Only for use by moderators.
	SubChannelId *string

	noSmithyDocumentSerde
}

type DeleteChannelMembershipOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteChannelMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteChannelMembership{}, middleware.After)
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
	if err = addOpDeleteChannelMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChannelMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteChannelMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteChannelMembership",
	}
}
