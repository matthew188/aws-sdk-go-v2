// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a pricing rule that can be associated to a pricing plan, or set of
// pricing plans.
func (c *Client) ListPricingRules(ctx context.Context, params *ListPricingRulesInput, optFns ...func(*Options)) (*ListPricingRulesOutput, error) {
	if params == nil {
		params = &ListPricingRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPricingRules", params, optFns, c.addOperationListPricingRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPricingRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPricingRulesInput struct {

	// The preferred billing period to get the pricing plan.
	BillingPeriod *string

	// A DescribePricingRuleFilter that specifies the Amazon Resource Name (ARNs) of
	// pricing rules to retrieve pricing rules information.
	Filters *types.ListPricingRulesFilter

	// The maximum number of pricing rules to retrieve.
	MaxResults *int32

	// The pagination token that's used on subsequent call to get pricing rules.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPricingRulesOutput struct {

	// The billing period for which the described pricing rules are applicable.
	BillingPeriod *string

	// The pagination token that's used on subsequent calls to get pricing rules.
	NextToken *string

	// A list containing the described pricing rules.
	PricingRules []types.PricingRuleListElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPricingRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPricingRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPricingRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPricingRules(options.Region), middleware.Before); err != nil {
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

// ListPricingRulesAPIClient is a client that implements the ListPricingRules
// operation.
type ListPricingRulesAPIClient interface {
	ListPricingRules(context.Context, *ListPricingRulesInput, ...func(*Options)) (*ListPricingRulesOutput, error)
}

var _ ListPricingRulesAPIClient = (*Client)(nil)

// ListPricingRulesPaginatorOptions is the paginator options for ListPricingRules
type ListPricingRulesPaginatorOptions struct {
	// The maximum number of pricing rules to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPricingRulesPaginator is a paginator for ListPricingRules
type ListPricingRulesPaginator struct {
	options   ListPricingRulesPaginatorOptions
	client    ListPricingRulesAPIClient
	params    *ListPricingRulesInput
	nextToken *string
	firstPage bool
}

// NewListPricingRulesPaginator returns a new ListPricingRulesPaginator
func NewListPricingRulesPaginator(client ListPricingRulesAPIClient, params *ListPricingRulesInput, optFns ...func(*ListPricingRulesPaginatorOptions)) *ListPricingRulesPaginator {
	if params == nil {
		params = &ListPricingRulesInput{}
	}

	options := ListPricingRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPricingRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPricingRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPricingRules page.
func (p *ListPricingRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPricingRulesOutput, error) {
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

	result, err := p.client.ListPricingRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPricingRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListPricingRules",
	}
}
