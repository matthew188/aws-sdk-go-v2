// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a managed endpoint. A managed endpoint is a gateway that connects Amazon
// EMR Studio to Amazon EMR on EKS so that Amazon EMR Studio can communicate with
// your virtual cluster.
func (c *Client) DeleteManagedEndpoint(ctx context.Context, params *DeleteManagedEndpointInput, optFns ...func(*Options)) (*DeleteManagedEndpointOutput, error) {
	if params == nil {
		params = &DeleteManagedEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteManagedEndpoint", params, optFns, c.addOperationDeleteManagedEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteManagedEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteManagedEndpointInput struct {

	// The ID of the managed endpoint.
	//
	// This member is required.
	Id *string

	// The ID of the endpoint's virtual cluster.
	//
	// This member is required.
	VirtualClusterId *string

	noSmithyDocumentSerde
}

type DeleteManagedEndpointOutput struct {

	// The output displays the ID of the managed endpoint.
	Id *string

	// The output displays the ID of the endpoint's virtual cluster.
	VirtualClusterId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteManagedEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteManagedEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteManagedEndpoint{}, middleware.After)
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
	if err = addOpDeleteManagedEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteManagedEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteManagedEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "DeleteManagedEndpoint",
	}
}
