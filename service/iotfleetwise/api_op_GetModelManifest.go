// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a vehicle model (model manifest).
func (c *Client) GetModelManifest(ctx context.Context, params *GetModelManifestInput, optFns ...func(*Options)) (*GetModelManifestOutput, error) {
	if params == nil {
		params = &GetModelManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModelManifest", params, optFns, c.addOperationGetModelManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelManifestInput struct {

	// The name of the vehicle model to retrieve information about.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetModelManifestOutput struct {

	// The Amazon Resource Name (ARN) of the vehicle model.
	//
	// This member is required.
	Arn *string

	// The time the vehicle model was created, in seconds since epoch (January 1, 1970
	// at midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	// The last time the vehicle model was modified.
	//
	// This member is required.
	LastModificationTime *time.Time

	// The name of the vehicle model.
	//
	// This member is required.
	Name *string

	// A brief description of the vehicle model.
	Description *string

	// The ARN of the signal catalog associated with the vehicle model.
	SignalCatalogArn *string

	// The state of the vehicle model. If the status is ACTIVE, the vehicle model can't
	// be edited. You can edit the vehicle model if the status is marked DRAFT.
	Status types.ManifestStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetModelManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetModelManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetModelManifest{}, middleware.After)
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
	if err = addOpGetModelManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModelManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "GetModelManifest",
	}
}
