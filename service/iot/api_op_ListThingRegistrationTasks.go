// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List bulk thing provisioning tasks. Requires permission to access the
// ListThingRegistrationTasks
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListThingRegistrationTasks(ctx context.Context, params *ListThingRegistrationTasksInput, optFns ...func(*Options)) (*ListThingRegistrationTasksOutput, error) {
	if params == nil {
		params = &ListThingRegistrationTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingRegistrationTasks", params, optFns, c.addOperationListThingRegistrationTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingRegistrationTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingRegistrationTasksInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// The status of the bulk thing provisioning task.
	Status types.Status

	noSmithyDocumentSerde
}

type ListThingRegistrationTasksOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// A list of bulk thing provisioning task IDs.
	TaskIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingRegistrationTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingRegistrationTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingRegistrationTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingRegistrationTasks(options.Region), middleware.Before); err != nil {
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

// ListThingRegistrationTasksAPIClient is a client that implements the
// ListThingRegistrationTasks operation.
type ListThingRegistrationTasksAPIClient interface {
	ListThingRegistrationTasks(context.Context, *ListThingRegistrationTasksInput, ...func(*Options)) (*ListThingRegistrationTasksOutput, error)
}

var _ ListThingRegistrationTasksAPIClient = (*Client)(nil)

// ListThingRegistrationTasksPaginatorOptions is the paginator options for
// ListThingRegistrationTasks
type ListThingRegistrationTasksPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingRegistrationTasksPaginator is a paginator for
// ListThingRegistrationTasks
type ListThingRegistrationTasksPaginator struct {
	options   ListThingRegistrationTasksPaginatorOptions
	client    ListThingRegistrationTasksAPIClient
	params    *ListThingRegistrationTasksInput
	nextToken *string
	firstPage bool
}

// NewListThingRegistrationTasksPaginator returns a new
// ListThingRegistrationTasksPaginator
func NewListThingRegistrationTasksPaginator(client ListThingRegistrationTasksAPIClient, params *ListThingRegistrationTasksInput, optFns ...func(*ListThingRegistrationTasksPaginatorOptions)) *ListThingRegistrationTasksPaginator {
	if params == nil {
		params = &ListThingRegistrationTasksInput{}
	}

	options := ListThingRegistrationTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingRegistrationTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingRegistrationTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingRegistrationTasks page.
func (p *ListThingRegistrationTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingRegistrationTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListThingRegistrationTasks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListThingRegistrationTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingRegistrationTasks",
	}
}
