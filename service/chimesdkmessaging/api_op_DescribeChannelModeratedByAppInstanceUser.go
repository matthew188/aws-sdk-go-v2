// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the full details of a channel moderated by the specified AppInstanceUser
// or AppInstanceBot. The x-amz-chime-bearer request header is mandatory. Use the
// ARN of the AppInstanceUser or AppInstanceBot that makes the API call as the
// value in the header.
func (c *Client) DescribeChannelModeratedByAppInstanceUser(ctx context.Context, params *DescribeChannelModeratedByAppInstanceUserInput, optFns ...func(*Options)) (*DescribeChannelModeratedByAppInstanceUserOutput, error) {
	if params == nil {
		params = &DescribeChannelModeratedByAppInstanceUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChannelModeratedByAppInstanceUser", params, optFns, c.addOperationDescribeChannelModeratedByAppInstanceUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChannelModeratedByAppInstanceUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChannelModeratedByAppInstanceUserInput struct {

	// The ARN of the user or bot in the moderated channel.
	//
	// This member is required.
	AppInstanceUserArn *string

	// The ARN of the moderated channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type DescribeChannelModeratedByAppInstanceUserOutput struct {

	// The moderated channel.
	Channel *types.ChannelModeratedByAppInstanceUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChannelModeratedByAppInstanceUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannelModeratedByAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannelModeratedByAppInstanceUser{}, middleware.After)
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
	if err = addOpDescribeChannelModeratedByAppInstanceUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannelModeratedByAppInstanceUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChannelModeratedByAppInstanceUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DescribeChannelModeratedByAppInstanceUser",
	}
}
