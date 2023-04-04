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

// Updates a vehicle model (model manifest). If created vehicles are associated
// with a vehicle model, it can't be updated.
func (c *Client) UpdateModelManifest(ctx context.Context, params *UpdateModelManifestInput, optFns ...func(*Options)) (*UpdateModelManifestOutput, error) {
	if params == nil {
		params = &UpdateModelManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelManifest", params, optFns, c.addOperationUpdateModelManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelManifestInput struct {

	// The name of the vehicle model to update.
	//
	// This member is required.
	Name *string

	// A brief description of the vehicle model.
	Description *string

	// A list of fullyQualifiedName of nodes, which are a general abstraction of
	// signals, to add to the vehicle model.
	NodesToAdd []string

	// A list of fullyQualifiedName of nodes, which are a general abstraction of
	// signals, to remove from the vehicle model.
	NodesToRemove []string

	// The state of the vehicle model. If the status is ACTIVE, the vehicle model can't
	// be edited. If the status is DRAFT, you can edit the vehicle model.
	Status types.ManifestStatus

	noSmithyDocumentSerde
}

type UpdateModelManifestOutput struct {

	// The Amazon Resource Name (ARN) of the updated vehicle model.
	//
	// This member is required.
	Arn *string

	// The name of the updated vehicle model.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateModelManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateModelManifest{}, middleware.After)
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
	if err = addOpUpdateModelManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "UpdateModelManifest",
	}
}
