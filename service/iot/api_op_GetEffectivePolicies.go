// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the policies that have an effect on the authorization behavior of
// the specified device when it connects to the IoT device gateway. Requires
// permission to access the GetEffectivePolicies
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) GetEffectivePolicies(ctx context.Context, params *GetEffectivePoliciesInput, optFns ...func(*Options)) (*GetEffectivePoliciesOutput, error) {
	if params == nil {
		params = &GetEffectivePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEffectivePolicies", params, optFns, c.addOperationGetEffectivePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEffectivePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectivePoliciesInput struct {

	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string

	// The principal. Valid principals are CertificateArn
	// (arn:aws:iot:region:accountId:cert/certificateId), thingGroupArn
	// (arn:aws:iot:region:accountId:thinggroup/groupName) and CognitoId (region:id).
	Principal *string

	// The thing name.
	ThingName *string

	noSmithyDocumentSerde
}

type GetEffectivePoliciesOutput struct {

	// The effective policies.
	EffectivePolicies []types.EffectivePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEffectivePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEffectivePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEffectivePolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectivePolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEffectivePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetEffectivePolicies",
	}
}
