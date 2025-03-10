// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update an existing pipe. When you call UpdatePipe, only the fields that are
// included in the request are changed, the rest are unchanged. The exception to
// this is if you modify any Amazon Web Services-service specific fields in the
// SourceParameters, EnrichmentParameters, or TargetParameters objects. The fields
// in these objects are updated atomically as one and override existing values.
// This is by design and means that if you don't specify an optional field in one
// of these Parameters objects, that field will be set to its system-default value
// after the update. For more information about pipes, see  Amazon EventBridge
// Pipes (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html)
// in the Amazon EventBridge User Guide.
func (c *Client) UpdatePipe(ctx context.Context, params *UpdatePipeInput, optFns ...func(*Options)) (*UpdatePipeOutput, error) {
	if params == nil {
		params = &UpdatePipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipe", params, optFns, c.addOperationUpdatePipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePipeInput struct {

	// The name of the pipe.
	//
	// This member is required.
	Name *string

	// The ARN of the role that allows the pipe to send data to the target.
	//
	// This member is required.
	RoleArn *string

	// A description of the pipe.
	Description *string

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// The ARN of the enrichment resource.
	Enrichment *string

	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters *types.PipeEnrichmentParameters

	// The parameters required to set up a source for your pipe.
	SourceParameters *types.UpdatePipeSourceParameters

	// The ARN of the target resource.
	Target *string

	// The parameters required to set up a target for your pipe.
	TargetParameters *types.PipeTargetParameters

	noSmithyDocumentSerde
}

type UpdatePipeOutput struct {

	// The ARN of the pipe.
	Arn *string

	// The time the pipe was created.
	CreationTime *time.Time

	// The state the pipe is in.
	CurrentState types.PipeState

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// When the pipe was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModifiedTime *time.Time

	// The name of the pipe.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipe{}, middleware.After)
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
	if err = addOpUpdatePipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pipes",
		OperationName: "UpdatePipe",
	}
}
