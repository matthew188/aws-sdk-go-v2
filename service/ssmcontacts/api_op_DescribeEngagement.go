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

// Incident Manager uses engagements to engage contacts and escalation plans during
// an incident. Use this command to describe the engagement that occurred during an
// incident.
func (c *Client) DescribeEngagement(ctx context.Context, params *DescribeEngagementInput, optFns ...func(*Options)) (*DescribeEngagementOutput, error) {
	if params == nil {
		params = &DescribeEngagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEngagement", params, optFns, c.addOperationDescribeEngagementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEngagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEngagementInput struct {

	// The Amazon Resource Name (ARN) of the engagement you want the details of.
	//
	// This member is required.
	EngagementId *string

	noSmithyDocumentSerde
}

type DescribeEngagementOutput struct {

	// The ARN of the escalation plan or contacts involved in the engagement.
	//
	// This member is required.
	ContactArn *string

	// The secure content of the message that was sent to the contact. Use this field
	// for engagements to VOICE and EMAIL.
	//
	// This member is required.
	Content *string

	// The ARN of the engagement.
	//
	// This member is required.
	EngagementArn *string

	// The user that started the engagement.
	//
	// This member is required.
	Sender *string

	// The secure subject of the message that was sent to the contact. Use this field
	// for engagements to VOICE and EMAIL.
	//
	// This member is required.
	Subject *string

	// The ARN of the incident in which the engagement occurred.
	IncidentId *string

	// The insecure content of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicContent *string

	// The insecure subject of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicSubject *string

	// The time that the engagement started.
	StartTime *time.Time

	// The time that the engagement ended.
	StopTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEngagementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEngagement{}, middleware.After)
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
	if err = addOpDescribeEngagementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngagement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEngagement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "DescribeEngagement",
	}
}
