// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the associations that were created between Resolver rules and VPCs using
// the current Amazon Web Services account.
func (c *Client) ListResolverRuleAssociations(ctx context.Context, params *ListResolverRuleAssociationsInput, optFns ...func(*Options)) (*ListResolverRuleAssociationsOutput, error) {
	if params == nil {
		params = &ListResolverRuleAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverRuleAssociations", params, optFns, c.addOperationListResolverRuleAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverRuleAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverRuleAssociationsInput struct {

	// An optional specification to return a subset of Resolver rules, such as Resolver
	// rules that are associated with the same VPC ID. If you submit a second or
	// subsequent ListResolverRuleAssociations request and specify the NextToken
	// parameter, you must use the same values for Filters, if any, as in the previous
	// request.
	Filters []types.Filter

	// The maximum number of rule associations that you want to return in the response
	// to a ListResolverRuleAssociations request. If you don't specify a value for
	// MaxResults, Resolver returns up to 100 rule associations.
	MaxResults *int32

	// For the first ListResolverRuleAssociation request, omit this value. If you have
	// more than MaxResults rule associations, you can submit another
	// ListResolverRuleAssociation request to get the next group of rule associations.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverRuleAssociationsOutput struct {

	// The value that you specified for MaxResults in the request.
	MaxResults *int32

	// If more than MaxResults rule associations match the specified criteria, you can
	// submit another ListResolverRuleAssociation request to get the next group of
	// results. In the next request, specify the value of NextToken from the previous
	// response.
	NextToken *string

	// The associations that were created between Resolver rules and VPCs using the
	// current Amazon Web Services account, and that match the specified filters, if
	// any.
	ResolverRuleAssociations []types.ResolverRuleAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverRuleAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverRuleAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverRuleAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverRuleAssociations(options.Region), middleware.Before); err != nil {
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

// ListResolverRuleAssociationsAPIClient is a client that implements the
// ListResolverRuleAssociations operation.
type ListResolverRuleAssociationsAPIClient interface {
	ListResolverRuleAssociations(context.Context, *ListResolverRuleAssociationsInput, ...func(*Options)) (*ListResolverRuleAssociationsOutput, error)
}

var _ ListResolverRuleAssociationsAPIClient = (*Client)(nil)

// ListResolverRuleAssociationsPaginatorOptions is the paginator options for
// ListResolverRuleAssociations
type ListResolverRuleAssociationsPaginatorOptions struct {
	// The maximum number of rule associations that you want to return in the response
	// to a ListResolverRuleAssociations request. If you don't specify a value for
	// MaxResults, Resolver returns up to 100 rule associations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverRuleAssociationsPaginator is a paginator for
// ListResolverRuleAssociations
type ListResolverRuleAssociationsPaginator struct {
	options   ListResolverRuleAssociationsPaginatorOptions
	client    ListResolverRuleAssociationsAPIClient
	params    *ListResolverRuleAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListResolverRuleAssociationsPaginator returns a new
// ListResolverRuleAssociationsPaginator
func NewListResolverRuleAssociationsPaginator(client ListResolverRuleAssociationsAPIClient, params *ListResolverRuleAssociationsInput, optFns ...func(*ListResolverRuleAssociationsPaginatorOptions)) *ListResolverRuleAssociationsPaginator {
	if params == nil {
		params = &ListResolverRuleAssociationsInput{}
	}

	options := ListResolverRuleAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverRuleAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverRuleAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverRuleAssociations page.
func (p *ListResolverRuleAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverRuleAssociationsOutput, error) {
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

	result, err := p.client.ListResolverRuleAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResolverRuleAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverRuleAssociations",
	}
}
