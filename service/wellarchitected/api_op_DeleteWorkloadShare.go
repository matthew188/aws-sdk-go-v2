// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a workload share.
func (c *Client) DeleteWorkloadShare(ctx context.Context, params *DeleteWorkloadShareInput, optFns ...func(*Options)) (*DeleteWorkloadShareOutput, error) {
	if params == nil {
		params = &DeleteWorkloadShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkloadShare", params, optFns, c.addOperationDeleteWorkloadShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkloadShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Delete Workload Share
type DeleteWorkloadShareInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The ID associated with the workload share.
	//
	// This member is required.
	ShareId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	//
	// This member is required.
	WorkloadId *string

	noSmithyDocumentSerde
}

type DeleteWorkloadShareOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWorkloadShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteWorkloadShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteWorkloadShare{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteWorkloadShareMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteWorkloadShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkloadShare(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteWorkloadShare struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteWorkloadShare) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteWorkloadShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteWorkloadShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteWorkloadShareInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteWorkloadShareMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteWorkloadShare{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteWorkloadShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "DeleteWorkloadShare",
	}
}
