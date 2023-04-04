// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists details of the engagement to a contact channel.
func (c *Client) DescribePage(ctx context.Context, params *DescribePageInput, optFns ...func(*Options)) (*DescribePageOutput, error) {
	if params == nil {
		params = &DescribePageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePage", params, optFns, c.addOperationDescribePageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePageInput struct {

	// The ID of the engagement to a contact channel.
	//
	// This member is required.
	PageId *string

	noSmithyDocumentSerde
}

type DescribePageOutput struct {

	// The ARN of the contact that was engaged.
	//
	// This member is required.
	ContactArn *string

	// The secure content of the message that was sent to the contact. Use this field
	// for engagements to VOICE and EMAIL.
	//
	// This member is required.
	Content *string

	// The ARN of the engagement that engaged the contact channel.
	//
	// This member is required.
	EngagementArn *string

	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	//
	// This member is required.
	PageArn *string

	// The user that started the engagement.
	//
	// This member is required.
	Sender *string

	// The secure subject of the message that was sent to the contact. Use this field
	// for engagements to VOICE and EMAIL.
	//
	// This member is required.
	Subject *string

	// The time that the contact channel received the engagement.
	DeliveryTime *time.Time

	// The ARN of the incident that engaged the contact channel.
	IncidentId *string

	// The insecure content of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicContent *string

	// The insecure subject of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicSubject *string

	// The time that the contact channel acknowledged the engagement.
	ReadTime *time.Time

	// The time the engagement was sent to the contact channel.
	SentTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePage{}, middleware.After)
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
	if err = addOpDescribePageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "DescribePage",
	}
}
