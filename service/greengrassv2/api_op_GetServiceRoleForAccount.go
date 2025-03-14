// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the service role associated with IoT Greengrass for your Amazon Web
// Services account in this Amazon Web Services Region. IoT Greengrass uses this
// role to verify the identity of client devices and manage core device
// connectivity information. For more information, see Greengrass service role
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html)
// in the IoT Greengrass Version 2 Developer Guide.
func (c *Client) GetServiceRoleForAccount(ctx context.Context, params *GetServiceRoleForAccountInput, optFns ...func(*Options)) (*GetServiceRoleForAccountOutput, error) {
	if params == nil {
		params = &GetServiceRoleForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceRoleForAccount", params, optFns, c.addOperationGetServiceRoleForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceRoleForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceRoleForAccountInput struct {
	noSmithyDocumentSerde
}

type GetServiceRoleForAccountOutput struct {

	// The time when the service role was associated with IoT Greengrass for your
	// Amazon Web Services account in this Amazon Web Services Region.
	AssociatedAt *string

	// The ARN of the service role that is associated with IoT Greengrass for your
	// Amazon Web Services account in this Amazon Web Services Region.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceRoleForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceRoleForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceRoleForAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceRoleForAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceRoleForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetServiceRoleForAccount",
	}
}
