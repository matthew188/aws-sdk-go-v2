// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a keyword configuration on an origination phone number or
// pool. A keyword is a word that you can search for on a particular phone number
// or pool. It is also a specific word or phrase that an end user can send to your
// number to elicit a response, such as an informational message or a special
// offer. When your number receives a message that begins with a keyword, Amazon
// Pinpoint responds with a customizable message. If you specify a keyword that
// isn't valid, an Error is returned.
func (c *Client) PutKeyword(ctx context.Context, params *PutKeywordInput, optFns ...func(*Options)) (*PutKeywordOutput, error) {
	if params == nil {
		params = &PutKeywordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutKeyword", params, optFns, c.addOperationPutKeywordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutKeywordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutKeywordInput struct {

	// The new keyword to add.
	//
	// This member is required.
	Keyword *string

	// The message associated with the keyword.
	//
	// * AUTOMATIC_RESPONSE: A message is
	// sent to the recipient.
	//
	// * OPT_OUT: Keeps the recipient from receiving future
	// messages.
	//
	// * OPT_IN: The recipient wants to receive future messages.
	//
	// This member is required.
	KeywordMessage *string

	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn,
	// SenderId or SenderIdArn. You can use DescribePhoneNumbers get the values for
	// PhoneNumberId and PhoneNumberArn while DescribeSenderIds can be used to get the
	// values for SenderId and SenderIdArn.
	//
	// This member is required.
	OriginationIdentity *string

	// The action to perform for the new keyword when it is received.
	KeywordAction types.KeywordAction

	noSmithyDocumentSerde
}

type PutKeywordOutput struct {

	// The keyword that was added.
	Keyword *string

	// The action to perform when the keyword is used.
	KeywordAction types.KeywordAction

	// The message associated with the keyword.
	KeywordMessage *string

	// The PhoneNumberId or PoolId that the keyword was associated with.
	OriginationIdentity *string

	// The PhoneNumberArn or PoolArn that the keyword was associated with.
	OriginationIdentityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutKeywordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutKeyword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutKeyword{}, middleware.After)
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
	if err = addOpPutKeywordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutKeyword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutKeyword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "PutKeyword",
	}
}
