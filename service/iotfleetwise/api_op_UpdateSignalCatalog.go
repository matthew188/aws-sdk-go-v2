// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a signal catalog.
func (c *Client) UpdateSignalCatalog(ctx context.Context, params *UpdateSignalCatalogInput, optFns ...func(*Options)) (*UpdateSignalCatalogOutput, error) {
	if params == nil {
		params = &UpdateSignalCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSignalCatalog", params, optFns, c.addOperationUpdateSignalCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSignalCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSignalCatalogInput struct {

	// The name of the signal catalog to update.
	//
	// This member is required.
	Name *string

	// A brief description of the signal catalog to update.
	Description *string

	// A list of information about nodes to add to the signal catalog.
	NodesToAdd []types.Node

	// A list of fullyQualifiedName of nodes to remove from the signal catalog.
	NodesToRemove []string

	// A list of information about nodes to update in the signal catalog.
	NodesToUpdate []types.Node

	noSmithyDocumentSerde
}

type UpdateSignalCatalogOutput struct {

	// The ARN of the updated signal catalog.
	//
	// This member is required.
	Arn *string

	// The name of the updated signal catalog.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSignalCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateSignalCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateSignalCatalog{}, middleware.After)
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
	if err = addOpUpdateSignalCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSignalCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSignalCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "UpdateSignalCatalog",
	}
}
