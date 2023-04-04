// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the authorization rules for a specified Client VPN endpoint.
func (c *Client) DescribeClientVpnAuthorizationRules(ctx context.Context, params *DescribeClientVpnAuthorizationRulesInput, optFns ...func(*Options)) (*DescribeClientVpnAuthorizationRulesOutput, error) {
	if params == nil {
		params = &DescribeClientVpnAuthorizationRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClientVpnAuthorizationRules", params, optFns, c.addOperationDescribeClientVpnAuthorizationRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClientVpnAuthorizationRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnAuthorizationRulesInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	// * description
	// - The description of the authorization rule.
	//
	// * destination-cidr - The CIDR of
	// the network to which the authorization rule applies.
	//
	// * group-id - The ID of the
	// Active Directory group to which the authorization rule grants access.
	Filters []types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeClientVpnAuthorizationRulesOutput struct {

	// Information about the authorization rules.
	AuthorizationRules []types.AuthorizationRule

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClientVpnAuthorizationRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnAuthorizationRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnAuthorizationRules{}, middleware.After)
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
	if err = addOpDescribeClientVpnAuthorizationRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnAuthorizationRules(options.Region), middleware.Before); err != nil {
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

// DescribeClientVpnAuthorizationRulesAPIClient is a client that implements the
// DescribeClientVpnAuthorizationRules operation.
type DescribeClientVpnAuthorizationRulesAPIClient interface {
	DescribeClientVpnAuthorizationRules(context.Context, *DescribeClientVpnAuthorizationRulesInput, ...func(*Options)) (*DescribeClientVpnAuthorizationRulesOutput, error)
}

var _ DescribeClientVpnAuthorizationRulesAPIClient = (*Client)(nil)

// DescribeClientVpnAuthorizationRulesPaginatorOptions is the paginator options for
// DescribeClientVpnAuthorizationRules
type DescribeClientVpnAuthorizationRulesPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClientVpnAuthorizationRulesPaginator is a paginator for
// DescribeClientVpnAuthorizationRules
type DescribeClientVpnAuthorizationRulesPaginator struct {
	options   DescribeClientVpnAuthorizationRulesPaginatorOptions
	client    DescribeClientVpnAuthorizationRulesAPIClient
	params    *DescribeClientVpnAuthorizationRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeClientVpnAuthorizationRulesPaginator returns a new
// DescribeClientVpnAuthorizationRulesPaginator
func NewDescribeClientVpnAuthorizationRulesPaginator(client DescribeClientVpnAuthorizationRulesAPIClient, params *DescribeClientVpnAuthorizationRulesInput, optFns ...func(*DescribeClientVpnAuthorizationRulesPaginatorOptions)) *DescribeClientVpnAuthorizationRulesPaginator {
	if params == nil {
		params = &DescribeClientVpnAuthorizationRulesInput{}
	}

	options := DescribeClientVpnAuthorizationRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClientVpnAuthorizationRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClientVpnAuthorizationRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClientVpnAuthorizationRules page.
func (p *DescribeClientVpnAuthorizationRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClientVpnAuthorizationRulesOutput, error) {
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

	result, err := p.client.DescribeClientVpnAuthorizationRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeClientVpnAuthorizationRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClientVpnAuthorizationRules",
	}
}
