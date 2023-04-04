// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an input.
func (c *Client) UpdateInput(ctx context.Context, params *UpdateInputInput, optFns ...func(*Options)) (*UpdateInputOutput, error) {
	if params == nil {
		params = &UpdateInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInput", params, optFns, c.addOperationUpdateInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update an input.
type UpdateInputInput struct {

	// Unique ID of the input.
	//
	// This member is required.
	InputId *string

	// Destination settings for PUSH type inputs.
	Destinations []types.InputDestinationRequest

	// Settings for the devices.
	InputDevices []types.InputDeviceRequest

	// A list of security groups referenced by IDs to attach to the input.
	InputSecurityGroups []string

	// A list of the MediaConnect Flow ARNs that you want to use as the source of the
	// input. You can specify as few as one Flow and presently, as many as two. The
	// only requirement is when you have more than one is that each Flow is in a
	// separate Availability Zone as this ensures your EML input is redundant to AZ
	// issues.
	MediaConnectFlows []types.MediaConnectFlowRequest

	// Name of the input.
	Name *string

	// The Amazon Resource Name (ARN) of the role this input assumes during and after
	// creation.
	RoleArn *string

	// The source URLs for a PULL-type input. Every PULL type input needs exactly two
	// source URLs for redundancy. Only specify sources for PULL type Inputs. Leave
	// Destinations empty.
	Sources []types.InputSourceRequest

	noSmithyDocumentSerde
}

// Placeholder documentation for UpdateInputResponse
type UpdateInputOutput struct {

	// Placeholder documentation for Input
	Input *types.Input

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInput{}, middleware.After)
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
	if err = addOpUpdateInputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateInput",
	}
}
