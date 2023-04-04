// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all engagements that have happened in an incident.
func (c *Client) ListEngagements(ctx context.Context, params *ListEngagementsInput, optFns ...func(*Options)) (*ListEngagementsOutput, error) {
	if params == nil {
		params = &ListEngagementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEngagements", params, optFns, c.addOperationListEngagementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEngagementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEngagementsInput struct {

	// The Amazon Resource Name (ARN) of the incident you're listing engagements for.
	IncidentId *string

	// The maximum number of engagements per page of results.
	MaxResults *int32

	// The pagination token to continue to the next page of results.
	NextToken *string

	// The time range to lists engagements for an incident.
	TimeRangeValue *types.TimeRange

	noSmithyDocumentSerde
}

type ListEngagementsOutput struct {

	// A list of each engagement that occurred during the specified time range of an
	// incident.
	//
	// This member is required.
	Engagements []types.Engagement

	// The pagination token to continue to the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEngagementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEngagements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEngagements{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEngagements(options.Region), middleware.Before); err != nil {
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

// ListEngagementsAPIClient is a client that implements the ListEngagements
// operation.
type ListEngagementsAPIClient interface {
	ListEngagements(context.Context, *ListEngagementsInput, ...func(*Options)) (*ListEngagementsOutput, error)
}

var _ ListEngagementsAPIClient = (*Client)(nil)

// ListEngagementsPaginatorOptions is the paginator options for ListEngagements
type ListEngagementsPaginatorOptions struct {
	// The maximum number of engagements per page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEngagementsPaginator is a paginator for ListEngagements
type ListEngagementsPaginator struct {
	options   ListEngagementsPaginatorOptions
	client    ListEngagementsAPIClient
	params    *ListEngagementsInput
	nextToken *string
	firstPage bool
}

// NewListEngagementsPaginator returns a new ListEngagementsPaginator
func NewListEngagementsPaginator(client ListEngagementsAPIClient, params *ListEngagementsInput, optFns ...func(*ListEngagementsPaginatorOptions)) *ListEngagementsPaginator {
	if params == nil {
		params = &ListEngagementsInput{}
	}

	options := ListEngagementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEngagementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEngagementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEngagements page.
func (p *ListEngagementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEngagementsOutput, error) {
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

	result, err := p.client.ListEngagements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEngagements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "ListEngagements",
	}
}
