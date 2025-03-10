// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a batch job and returns the unique identifier of this execution of the
// batch job. The associated application must be running in order to start the
// batch job.
func (c *Client) StartBatchJob(ctx context.Context, params *StartBatchJobInput, optFns ...func(*Options)) (*StartBatchJobOutput, error) {
	if params == nil {
		params = &StartBatchJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBatchJob", params, optFns, c.addOperationStartBatchJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBatchJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBatchJobInput struct {

	// The unique identifier of the application associated with this batch job.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier of the batch job.
	//
	// This member is required.
	BatchJobIdentifier types.BatchJobIdentifier

	// The collection of batch job parameters. For details about limits for keys and
	// values, see Coding variables in JCL
	// (https://www.ibm.com/docs/en/workload-automation/9.3.0?topic=zos-coding-variables-in-jcl).
	JobParams map[string]string

	noSmithyDocumentSerde
}

type StartBatchJobOutput struct {

	// The unique identifier of this execution of the batch job.
	//
	// This member is required.
	ExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBatchJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBatchJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBatchJob{}, middleware.After)
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
	if err = addOpStartBatchJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBatchJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBatchJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "StartBatchJob",
	}
}
