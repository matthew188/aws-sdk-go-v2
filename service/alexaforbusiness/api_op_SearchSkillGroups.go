// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches skill groups and lists the ones that meet a set of filter and sort
// criteria.
func (c *Client) SearchSkillGroups(ctx context.Context, params *SearchSkillGroupsInput, optFns ...func(*Options)) (*SearchSkillGroupsOutput, error) {
	if params == nil {
		params = &SearchSkillGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSkillGroups", params, optFns, c.addOperationSearchSkillGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSkillGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSkillGroupsInput struct {

	// The filters to use to list a specified set of skill groups. The supported filter
	// key is SkillGroupName.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string

	// The sort order to use in listing the specified set of skill groups. The
	// supported sort key is SkillGroupName.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchSkillGroupsOutput struct {

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The skill groups that meet the filter criteria, in sort order.
	SkillGroups []types.SkillGroupData

	// The total number of skill groups returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchSkillGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSkillGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSkillGroups{}, middleware.After)
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
	if err = addOpSearchSkillGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSkillGroups(options.Region), middleware.Before); err != nil {
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

// SearchSkillGroupsAPIClient is a client that implements the SearchSkillGroups
// operation.
type SearchSkillGroupsAPIClient interface {
	SearchSkillGroups(context.Context, *SearchSkillGroupsInput, ...func(*Options)) (*SearchSkillGroupsOutput, error)
}

var _ SearchSkillGroupsAPIClient = (*Client)(nil)

// SearchSkillGroupsPaginatorOptions is the paginator options for SearchSkillGroups
type SearchSkillGroupsPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchSkillGroupsPaginator is a paginator for SearchSkillGroups
type SearchSkillGroupsPaginator struct {
	options   SearchSkillGroupsPaginatorOptions
	client    SearchSkillGroupsAPIClient
	params    *SearchSkillGroupsInput
	nextToken *string
	firstPage bool
}

// NewSearchSkillGroupsPaginator returns a new SearchSkillGroupsPaginator
func NewSearchSkillGroupsPaginator(client SearchSkillGroupsAPIClient, params *SearchSkillGroupsInput, optFns ...func(*SearchSkillGroupsPaginatorOptions)) *SearchSkillGroupsPaginator {
	if params == nil {
		params = &SearchSkillGroupsInput{}
	}

	options := SearchSkillGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchSkillGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchSkillGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchSkillGroups page.
func (p *SearchSkillGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchSkillGroupsOutput, error) {
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

	result, err := p.client.SearchSkillGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchSkillGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchSkillGroups",
	}
}
