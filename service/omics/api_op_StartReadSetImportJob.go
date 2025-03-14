// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a read set import job.
func (c *Client) StartReadSetImportJob(ctx context.Context, params *StartReadSetImportJobInput, optFns ...func(*Options)) (*StartReadSetImportJobOutput, error) {
	if params == nil {
		params = &StartReadSetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReadSetImportJob", params, optFns, c.addOperationStartReadSetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReadSetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReadSetImportJobInput struct {

	// A service role for the job.
	//
	// This member is required.
	RoleArn *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's source files.
	//
	// This member is required.
	Sources []types.StartReadSetImportJobSourceItem

	// To ensure that jobs don't run multiple times, specify a unique token for each
	// job.
	ClientToken *string

	noSmithyDocumentSerde
}

type StartReadSetImportJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status types.ReadSetImportJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReadSetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartReadSetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartReadSetImportJob{}, middleware.After)
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
	if err = addEndpointPrefix_opStartReadSetImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartReadSetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReadSetImportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartReadSetImportJobMiddleware struct {
}

func (*endpointPrefix_opStartReadSetImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartReadSetImportJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opStartReadSetImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opStartReadSetImportJobMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opStartReadSetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "StartReadSetImportJob",
	}
}
