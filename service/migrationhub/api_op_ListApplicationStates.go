// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the migration statuses for your applications. If you use the optional
// ApplicationIds parameter, only the migration statuses for those applications
// will be returned.
func (c *Client) ListApplicationStates(ctx context.Context, params *ListApplicationStatesInput, optFns ...func(*Options)) (*ListApplicationStatesOutput, error) {
	if params == nil {
		params = &ListApplicationStatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationStates", params, optFns, c.addOperationListApplicationStatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationStatesInput struct {

	// The configurationIds from the Application Discovery Service that uniquely
	// identifies your applications.
	ApplicationIds []string

	// Maximum number of results to be returned per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationStatesOutput struct {

	// A list of Applications that exist in Application Discovery Service.
	ApplicationStateList []types.ApplicationState

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationStatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationStates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationStates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationStates(options.Region), middleware.Before); err != nil {
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

// ListApplicationStatesAPIClient is a client that implements the
// ListApplicationStates operation.
type ListApplicationStatesAPIClient interface {
	ListApplicationStates(context.Context, *ListApplicationStatesInput, ...func(*Options)) (*ListApplicationStatesOutput, error)
}

var _ ListApplicationStatesAPIClient = (*Client)(nil)

// ListApplicationStatesPaginatorOptions is the paginator options for
// ListApplicationStates
type ListApplicationStatesPaginatorOptions struct {
	// Maximum number of results to be returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationStatesPaginator is a paginator for ListApplicationStates
type ListApplicationStatesPaginator struct {
	options   ListApplicationStatesPaginatorOptions
	client    ListApplicationStatesAPIClient
	params    *ListApplicationStatesInput
	nextToken *string
	firstPage bool
}

// NewListApplicationStatesPaginator returns a new ListApplicationStatesPaginator
func NewListApplicationStatesPaginator(client ListApplicationStatesAPIClient, params *ListApplicationStatesInput, optFns ...func(*ListApplicationStatesPaginatorOptions)) *ListApplicationStatesPaginator {
	if params == nil {
		params = &ListApplicationStatesInput{}
	}

	options := ListApplicationStatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationStatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationStatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationStates page.
func (p *ListApplicationStatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationStatesOutput, error) {
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

	result, err := p.client.ListApplicationStates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationStates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListApplicationStates",
	}
}
