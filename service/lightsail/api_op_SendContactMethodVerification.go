// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a verification request to an email contact method to ensure it's owned by
// the requester. SMS contact methods don't need to be verified. A contact method
// is used to send you notifications about your Amazon Lightsail resources. You can
// add one email address and one mobile phone number contact method in each Amazon
// Web Services Region. However, SMS text messaging is not supported in some Amazon
// Web Services Regions, and SMS text messages cannot be sent to some
// countries/regions. For more information, see Notifications in Amazon Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications).
// A verification request is sent to the contact method when you initially create
// it. Use this action to send another verification request if a previous
// verification request was deleted, or has expired. Notifications are not sent to
// an email contact method until after it is verified, and confirmed as valid.
func (c *Client) SendContactMethodVerification(ctx context.Context, params *SendContactMethodVerificationInput, optFns ...func(*Options)) (*SendContactMethodVerificationOutput, error) {
	if params == nil {
		params = &SendContactMethodVerificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendContactMethodVerification", params, optFns, c.addOperationSendContactMethodVerificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendContactMethodVerificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendContactMethodVerificationInput struct {

	// The protocol to verify, such as Email or SMS (text messaging).
	//
	// This member is required.
	Protocol types.ContactMethodVerificationProtocol

	noSmithyDocumentSerde
}

type SendContactMethodVerificationOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendContactMethodVerificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendContactMethodVerification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendContactMethodVerification{}, middleware.After)
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
	if err = addOpSendContactMethodVerificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendContactMethodVerification(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendContactMethodVerification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "SendContactMethodVerification",
	}
}
