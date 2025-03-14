// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the problems with your application.
func (c *Client) ListProblems(ctx context.Context, params *ListProblemsInput, optFns ...func(*Options)) (*ListProblemsOutput, error) {
	if params == nil {
		params = &ListProblemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProblems", params, optFns, c.addOperationListProblemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProblemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProblemsInput struct {

	// The name of the component.
	ComponentName *string

	// The time when the problem ended, in epoch seconds. If not specified, problems
	// within the past seven days are returned.
	EndTime *time.Time

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// The time when the problem was detected, in epoch seconds. If you don't specify a
	// time frame for the request, problems within the past seven days are returned.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListProblemsOutput struct {

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The list of problems.
	ProblemList []types.Problem

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProblemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProblems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProblems{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProblems(options.Region), middleware.Before); err != nil {
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

// ListProblemsAPIClient is a client that implements the ListProblems operation.
type ListProblemsAPIClient interface {
	ListProblems(context.Context, *ListProblemsInput, ...func(*Options)) (*ListProblemsOutput, error)
}

var _ ListProblemsAPIClient = (*Client)(nil)

// ListProblemsPaginatorOptions is the paginator options for ListProblems
type ListProblemsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProblemsPaginator is a paginator for ListProblems
type ListProblemsPaginator struct {
	options   ListProblemsPaginatorOptions
	client    ListProblemsAPIClient
	params    *ListProblemsInput
	nextToken *string
	firstPage bool
}

// NewListProblemsPaginator returns a new ListProblemsPaginator
func NewListProblemsPaginator(client ListProblemsAPIClient, params *ListProblemsInput, optFns ...func(*ListProblemsPaginatorOptions)) *ListProblemsPaginator {
	if params == nil {
		params = &ListProblemsInput{}
	}

	options := ListProblemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProblemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProblemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProblems page.
func (p *ListProblemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProblemsOutput, error) {
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

	result, err := p.client.ListProblems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProblems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListProblems",
	}
}
