// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the SIP rules under the administrator's AWS account.
func (c *Client) ListSipRules(ctx context.Context, params *ListSipRulesInput, optFns ...func(*Options)) (*ListSipRulesOutput, error) {
	if params == nil {
		params = &ListSipRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSipRules", params, optFns, c.addOperationListSipRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSipRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSipRulesInput struct {

	// The maximum number of results to return in a single call. Defaults to 100.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	// The SIP media application ID.
	SipMediaApplicationId *string

	noSmithyDocumentSerde
}

type ListSipRulesOutput struct {

	// The token used to return the next page of results.
	NextToken *string

	// The list of SIP rules and details.
	SipRules []types.SipRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSipRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSipRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSipRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSipRules(options.Region), middleware.Before); err != nil {
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

// ListSipRulesAPIClient is a client that implements the ListSipRules operation.
type ListSipRulesAPIClient interface {
	ListSipRules(context.Context, *ListSipRulesInput, ...func(*Options)) (*ListSipRulesOutput, error)
}

var _ ListSipRulesAPIClient = (*Client)(nil)

// ListSipRulesPaginatorOptions is the paginator options for ListSipRules
type ListSipRulesPaginatorOptions struct {
	// The maximum number of results to return in a single call. Defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSipRulesPaginator is a paginator for ListSipRules
type ListSipRulesPaginator struct {
	options   ListSipRulesPaginatorOptions
	client    ListSipRulesAPIClient
	params    *ListSipRulesInput
	nextToken *string
	firstPage bool
}

// NewListSipRulesPaginator returns a new ListSipRulesPaginator
func NewListSipRulesPaginator(client ListSipRulesAPIClient, params *ListSipRulesInput, optFns ...func(*ListSipRulesPaginatorOptions)) *ListSipRulesPaginator {
	if params == nil {
		params = &ListSipRulesInput{}
	}

	options := ListSipRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSipRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSipRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSipRules page.
func (p *ListSipRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSipRulesOutput, error) {
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

	result, err := p.client.ListSipRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSipRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListSipRules",
	}
}
