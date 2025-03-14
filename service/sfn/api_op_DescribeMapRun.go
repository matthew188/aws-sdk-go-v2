// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a Map Run's configuration, progress, and results. For
// more information, see Examining Map Run
// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-examine-map-run.html)
// in the Step Functions Developer Guide.
func (c *Client) DescribeMapRun(ctx context.Context, params *DescribeMapRunInput, optFns ...func(*Options)) (*DescribeMapRunOutput, error) {
	if params == nil {
		params = &DescribeMapRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMapRun", params, optFns, c.addOperationDescribeMapRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMapRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMapRunInput struct {

	// The Amazon Resource Name (ARN) that identifies a Map Run.
	//
	// This member is required.
	MapRunArn *string

	noSmithyDocumentSerde
}

type DescribeMapRunOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution in which the Map
	// Run was started.
	//
	// This member is required.
	ExecutionArn *string

	// A JSON object that contains information about the total number of child workflow
	// executions for the Map Run, and the count of child workflow executions for each
	// status, such as failed and succeeded.
	//
	// This member is required.
	ExecutionCounts *types.MapRunExecutionCounts

	// A JSON object that contains information about the total number of items, and the
	// item count for each processing status, such as pending and failed.
	//
	// This member is required.
	ItemCounts *types.MapRunItemCounts

	// The Amazon Resource Name (ARN) that identifies a Map Run.
	//
	// This member is required.
	MapRunArn *string

	// The maximum number of child workflow executions configured to run in parallel
	// for the Map Run at the same time.
	//
	// This member is required.
	MaxConcurrency int32

	// The date when the Map Run was started.
	//
	// This member is required.
	StartDate *time.Time

	// The current status of the Map Run.
	//
	// This member is required.
	Status types.MapRunStatus

	// The maximum number of failed child workflow executions before the Map Run fails.
	//
	// This member is required.
	ToleratedFailureCount int64

	// The maximum percentage of failed child workflow executions before the Map Run
	// fails.
	//
	// This member is required.
	ToleratedFailurePercentage float32

	// The date when the Map Run was stopped.
	StopDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMapRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeMapRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeMapRun{}, middleware.After)
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
	if err = addOpDescribeMapRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMapRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMapRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeMapRun",
	}
}
