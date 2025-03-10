// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the specific Connect instance config.
func (c *Client) GetConnectInstanceConfig(ctx context.Context, params *GetConnectInstanceConfigInput, optFns ...func(*Options)) (*GetConnectInstanceConfigOutput, error) {
	if params == nil {
		params = &GetConnectInstanceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnectInstanceConfig", params, optFns, c.addOperationGetConnectInstanceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectInstanceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetConnectInstanceConfigRequest
type GetConnectInstanceConfigInput struct {

	// Amazon Connect Instance Id
	//
	// This member is required.
	ConnectInstanceId *string

	noSmithyDocumentSerde
}

// GetConnectInstanceConfigResponse
type GetConnectInstanceConfigOutput struct {

	// Instance config object
	ConnectInstanceConfig *types.InstanceConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectInstanceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConnectInstanceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConnectInstanceConfig{}, middleware.After)
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
	if err = addOpGetConnectInstanceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnectInstanceConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConnectInstanceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect-campaigns",
		OperationName: "GetConnectInstanceConfig",
	}
}
