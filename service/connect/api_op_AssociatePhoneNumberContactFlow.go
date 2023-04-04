// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a flow with a phone number claimed to your Amazon Connect instance.
// If the number is claimed to a traffic distribution group, and you are calling
// this API using an instance in the Amazon Web Services Region where the traffic
// distribution group was created, you can use either a full phone number ARN or
// UUID value for the PhoneNumberId URI request parameter. However, if the number
// is claimed to a traffic distribution group and you are calling this API using an
// instance in the alternate Amazon Web Services Region associated with the traffic
// distribution group, you must provide a full phone number ARN. If a UUID is
// provided in this scenario, you will receive a ResourceNotFoundException.
func (c *Client) AssociatePhoneNumberContactFlow(ctx context.Context, params *AssociatePhoneNumberContactFlowInput, optFns ...func(*Options)) (*AssociatePhoneNumberContactFlowOutput, error) {
	if params == nil {
		params = &AssociatePhoneNumberContactFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePhoneNumberContactFlow", params, optFns, c.addOperationAssociatePhoneNumberContactFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePhoneNumberContactFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePhoneNumberContactFlowInput struct {

	// The identifier of the flow.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// A unique identifier for the phone number.
	//
	// This member is required.
	PhoneNumberId *string

	noSmithyDocumentSerde
}

type AssociatePhoneNumberContactFlowOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePhoneNumberContactFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociatePhoneNumberContactFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociatePhoneNumberContactFlow{}, middleware.After)
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
	if err = addOpAssociatePhoneNumberContactFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePhoneNumberContactFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePhoneNumberContactFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "AssociatePhoneNumberContactFlow",
	}
}
