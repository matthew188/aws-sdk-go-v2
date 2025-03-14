// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a mitigation action. Requires permission to access the
// DescribeMitigationAction
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeMitigationAction(ctx context.Context, params *DescribeMitigationActionInput, optFns ...func(*Options)) (*DescribeMitigationActionOutput, error) {
	if params == nil {
		params = &DescribeMitigationActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMitigationAction", params, optFns, c.addOperationDescribeMitigationActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMitigationActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMitigationActionInput struct {

	// The friendly name that uniquely identifies the mitigation action.
	//
	// This member is required.
	ActionName *string

	noSmithyDocumentSerde
}

type DescribeMitigationActionOutput struct {

	// The ARN that identifies this migration action.
	ActionArn *string

	// A unique identifier for this action.
	ActionId *string

	// The friendly name that uniquely identifies the mitigation action.
	ActionName *string

	// Parameters that control how the mitigation action is applied, specific to the
	// type of mitigation action.
	ActionParams *types.MitigationActionParams

	// The type of mitigation action.
	ActionType types.MitigationActionType

	// The date and time when the mitigation action was added to your Amazon Web
	// Services accounts.
	CreationDate *time.Time

	// The date and time when the mitigation action was last changed.
	LastModifiedDate *time.Time

	// The ARN of the IAM role used to apply this action.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMitigationActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMitigationAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMitigationAction{}, middleware.After)
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
	if err = addOpDescribeMitigationActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMitigationAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMitigationAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeMitigationAction",
	}
}
