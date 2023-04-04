// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all rulesets available in the current account or rulesets associated with a
// specific resource (dataset).
func (c *Client) ListRulesets(ctx context.Context, params *ListRulesetsInput, optFns ...func(*Options)) (*ListRulesetsOutput, error) {
	if params == nil {
		params = &ListRulesetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRulesets", params, optFns, c.addOperationListRulesetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRulesetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRulesetsInput struct {

	// The maximum number of results to return in this request.
	MaxResults *int32

	// A token generated by DataBrew that specifies where to continue pagination if a
	// previous request was truncated. To get the next set of pages, pass in the
	// NextToken value from the response object of the previous page call.
	NextToken *string

	// The Amazon Resource Name (ARN) of a resource (dataset). Using this parameter
	// indicates to return only those rulesets that are associated with the specified
	// resource.
	TargetArn *string

	noSmithyDocumentSerde
}

type ListRulesetsOutput struct {

	// A list of RulesetItem. RulesetItem contains meta data of a ruleset.
	//
	// This member is required.
	Rulesets []types.RulesetItem

	// A token that you can use in a subsequent call to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRulesetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRulesets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRulesets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRulesets(options.Region), middleware.Before); err != nil {
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

// ListRulesetsAPIClient is a client that implements the ListRulesets operation.
type ListRulesetsAPIClient interface {
	ListRulesets(context.Context, *ListRulesetsInput, ...func(*Options)) (*ListRulesetsOutput, error)
}

var _ ListRulesetsAPIClient = (*Client)(nil)

// ListRulesetsPaginatorOptions is the paginator options for ListRulesets
type ListRulesetsPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRulesetsPaginator is a paginator for ListRulesets
type ListRulesetsPaginator struct {
	options   ListRulesetsPaginatorOptions
	client    ListRulesetsAPIClient
	params    *ListRulesetsInput
	nextToken *string
	firstPage bool
}

// NewListRulesetsPaginator returns a new ListRulesetsPaginator
func NewListRulesetsPaginator(client ListRulesetsAPIClient, params *ListRulesetsInput, optFns ...func(*ListRulesetsPaginatorOptions)) *ListRulesetsPaginator {
	if params == nil {
		params = &ListRulesetsInput{}
	}

	options := ListRulesetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRulesetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRulesetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRulesets page.
func (p *ListRulesetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRulesetsOutput, error) {
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

	result, err := p.client.ListRulesets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRulesets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "ListRulesets",
	}
}
