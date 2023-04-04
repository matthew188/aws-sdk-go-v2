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

// Returns the details of a single named query or a list of up to 50 queries, which
// you provide as an array of query ID strings. Requires you to have access to the
// workgroup in which the queries were saved. Use ListNamedQueriesInput to get the
// list of named query IDs in the specified workgroup. If information could not be
// retrieved for a submitted query ID, information about the query ID submitted is
// listed under UnprocessedNamedQueryId. Named queries differ from executed
// queries. Use BatchGetQueryExecutionInput to get details about each unique query
// execution, and ListQueryExecutionsInput to get a list of query execution IDs.
func (c *Client) BatchGetNamedQuery(ctx context.Context, params *BatchGetNamedQueryInput, optFns ...func(*Options)) (*BatchGetNamedQueryOutput, error) {
	if params == nil {
		params = &BatchGetNamedQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetNamedQuery", params, optFns, c.addOperationBatchGetNamedQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetNamedQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains an array of named query IDs.
type BatchGetNamedQueryInput struct {

	// An array of query IDs.
	//
	// This member is required.
	NamedQueryIds []string

	noSmithyDocumentSerde
}

type BatchGetNamedQueryOutput struct {

	// Information about the named query IDs submitted.
	NamedQueries []types.NamedQuery

	// Information about provided query IDs.
	UnprocessedNamedQueryIds []types.UnprocessedNamedQueryId

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetNamedQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetNamedQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetNamedQuery{}, middleware.After)
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
	if err = addOpBatchGetNamedQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetNamedQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetNamedQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "BatchGetNamedQuery",
	}
}
