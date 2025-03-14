// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a group (batch) of assets from an IoT SiteWise Monitor project.
func (c *Client) BatchDisassociateProjectAssets(ctx context.Context, params *BatchDisassociateProjectAssetsInput, optFns ...func(*Options)) (*BatchDisassociateProjectAssetsOutput, error) {
	if params == nil {
		params = &BatchDisassociateProjectAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateProjectAssets", params, optFns, c.addOperationBatchDisassociateProjectAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateProjectAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateProjectAssetsInput struct {

	// The IDs of the assets to be disassociated from the project.
	//
	// This member is required.
	AssetIds []string

	// The ID of the project from which to disassociate the assets.
	//
	// This member is required.
	ProjectId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	noSmithyDocumentSerde
}

type BatchDisassociateProjectAssetsOutput struct {

	// A list of associated error information, if any.
	Errors []types.AssetErrorDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDisassociateProjectAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisassociateProjectAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisassociateProjectAssets{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchDisassociateProjectAssetsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opBatchDisassociateProjectAssetsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBatchDisassociateProjectAssetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateProjectAssets(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchDisassociateProjectAssetsMiddleware struct {
}

func (*endpointPrefix_opBatchDisassociateProjectAssetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchDisassociateProjectAssetsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchDisassociateProjectAssetsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchDisassociateProjectAssetsMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpBatchDisassociateProjectAssets struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpBatchDisassociateProjectAssets) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpBatchDisassociateProjectAssets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*BatchDisassociateProjectAssetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *BatchDisassociateProjectAssetsInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opBatchDisassociateProjectAssetsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpBatchDisassociateProjectAssets{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opBatchDisassociateProjectAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchDisassociateProjectAssets",
	}
}
