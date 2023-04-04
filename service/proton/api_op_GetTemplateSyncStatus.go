// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the status of a template sync.
func (c *Client) GetTemplateSyncStatus(ctx context.Context, params *GetTemplateSyncStatusInput, optFns ...func(*Options)) (*GetTemplateSyncStatusOutput, error) {
	if params == nil {
		params = &GetTemplateSyncStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplateSyncStatus", params, optFns, c.addOperationGetTemplateSyncStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateSyncStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateSyncStatusInput struct {

	// The template name.
	//
	// This member is required.
	TemplateName *string

	// The template type.
	//
	// This member is required.
	TemplateType types.TemplateType

	// The template major version.
	//
	// This member is required.
	TemplateVersion *string

	noSmithyDocumentSerde
}

type GetTemplateSyncStatusOutput struct {

	// The template sync desired state that's returned by Proton.
	DesiredState *types.Revision

	// The details of the last successful sync that's returned by Proton.
	LatestSuccessfulSync *types.ResourceSyncAttempt

	// The details of the last sync that's returned by Proton.
	LatestSync *types.ResourceSyncAttempt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateSyncStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetTemplateSyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetTemplateSyncStatus{}, middleware.After)
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
	if err = addOpGetTemplateSyncStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplateSyncStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemplateSyncStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetTemplateSyncStatus",
	}
}
