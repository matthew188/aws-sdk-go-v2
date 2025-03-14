// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon Redshift Serverless managed endpoint.
func (c *Client) UpdateEndpointAccess(ctx context.Context, params *UpdateEndpointAccessInput, optFns ...func(*Options)) (*UpdateEndpointAccessOutput, error) {
	if params == nil {
		params = &UpdateEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpointAccess", params, optFns, c.addOperationUpdateEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointAccessInput struct {

	// The name of the VPC endpoint to update.
	//
	// This member is required.
	EndpointName *string

	// The list of VPC security groups associated with the endpoint after the endpoint
	// is modified.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type UpdateEndpointAccessOutput struct {

	// The updated VPC endpoint.
	Endpoint *types.EndpointAccess

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpointAccess{}, middleware.After)
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
	if err = addOpUpdateEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "UpdateEndpointAccess",
	}
}
