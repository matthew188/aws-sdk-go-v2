// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the details of a single query execution or a list of up to 50 query
// executions, which you provide as an array of query execution ID strings.
// Requires you to have access to the workgroup in which the queries ran. To get a
// list of query execution IDs, use ListQueryExecutionsInput$WorkGroup. Query
// executions differ from named (saved) queries. Use BatchGetNamedQueryInput to get
// details about named queries.
func (c *Client) BatchGetQueryExecution(ctx context.Context, params *BatchGetQueryExecutionInput, optFns ...func(*Options)) (*BatchGetQueryExecutionOutput, error) {
	if params == nil {
		params = &BatchGetQueryExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetQueryExecution", params, optFns, c.addOperationBatchGetQueryExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetQueryExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains an array of query execution IDs.
type BatchGetQueryExecutionInput struct {

	// An array of query execution IDs.
	//
	// This member is required.
	QueryExecutionIds []string

	noSmithyDocumentSerde
}

type BatchGetQueryExecutionOutput struct {

	// Information about a query execution.
	QueryExecutions []types.QueryExecution

	// Information about the query executions that failed to run.
	UnprocessedQueryExecutionIds []types.UnprocessedQueryExecutionId

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetQueryExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetQueryExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetQueryExecution{}, middleware.After)
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
	if err = addOpBatchGetQueryExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetQueryExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetQueryExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "BatchGetQueryExecution",
	}
}
