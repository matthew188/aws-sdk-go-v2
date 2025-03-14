// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants entitlements to an existing flow.
func (c *Client) GrantFlowEntitlements(ctx context.Context, params *GrantFlowEntitlementsInput, optFns ...func(*Options)) (*GrantFlowEntitlementsOutput, error) {
	if params == nil {
		params = &GrantFlowEntitlementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GrantFlowEntitlements", params, optFns, c.addOperationGrantFlowEntitlementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GrantFlowEntitlementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to grant entitlements on a flow.
type GrantFlowEntitlementsInput struct {

	// The list of entitlements that you want to grant.
	//
	// This member is required.
	Entitlements []types.GrantEntitlementRequest

	// The flow that you want to grant entitlements on.
	//
	// This member is required.
	FlowArn *string

	noSmithyDocumentSerde
}

type GrantFlowEntitlementsOutput struct {

	// The entitlements that were just granted.
	Entitlements []types.Entitlement

	// The ARN of the flow that these entitlements were granted to.
	FlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGrantFlowEntitlementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGrantFlowEntitlements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGrantFlowEntitlements{}, middleware.After)
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
	if err = addOpGrantFlowEntitlementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGrantFlowEntitlements(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGrantFlowEntitlements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "GrantFlowEntitlements",
	}
}
