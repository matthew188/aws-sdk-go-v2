// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a description of custom domain names that are associated with an App
// Runner service.
func (c *Client) DescribeCustomDomains(ctx context.Context, params *DescribeCustomDomainsInput, optFns ...func(*Options)) (*DescribeCustomDomainsOutput, error) {
	if params == nil {
		params = &DescribeCustomDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomDomains", params, optFns, c.addOperationDescribeCustomDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomDomainsInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want
	// associated custom domain names to be described for.
	//
	// This member is required.
	ServiceArn *string

	// The maximum number of results that each response (result page) can include. It's
	// used for a paginated request. If you don't specify MaxResults, the request
	// retrieves all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. It's used for a paginated request. The
	// request retrieves the next result page. All other parameter values must be
	// identical to the ones that are specified in the initial request. If you don't
	// specify NextToken, the request retrieves the first result page.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeCustomDomainsOutput struct {

	// A list of descriptions of custom domain names that are associated with the
	// service. In a paginated request, the request returns up to MaxResults records
	// per call.
	//
	// This member is required.
	CustomDomains []types.CustomDomain

	// The App Runner subdomain of the App Runner service. The associated custom domain
	// names are mapped to this target name.
	//
	// This member is required.
	DNSTarget *string

	// The Amazon Resource Name (ARN) of the App Runner service whose associated custom
	// domain names you want to describe.
	//
	// This member is required.
	ServiceArn *string

	// DNS Target records for the custom domains of this Amazon VPC.
	//
	// This member is required.
	VpcDNSTargets []types.VpcDNSTarget

	// The token that you can pass in a subsequent request to get the next result page.
	// It's returned in a paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeCustomDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeCustomDomains{}, middleware.After)
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
	if err = addOpDescribeCustomDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomDomains(options.Region), middleware.Before); err != nil {
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

// DescribeCustomDomainsAPIClient is a client that implements the
// DescribeCustomDomains operation.
type DescribeCustomDomainsAPIClient interface {
	DescribeCustomDomains(context.Context, *DescribeCustomDomainsInput, ...func(*Options)) (*DescribeCustomDomainsOutput, error)
}

var _ DescribeCustomDomainsAPIClient = (*Client)(nil)

// DescribeCustomDomainsPaginatorOptions is the paginator options for
// DescribeCustomDomains
type DescribeCustomDomainsPaginatorOptions struct {
	// The maximum number of results that each response (result page) can include. It's
	// used for a paginated request. If you don't specify MaxResults, the request
	// retrieves all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCustomDomainsPaginator is a paginator for DescribeCustomDomains
type DescribeCustomDomainsPaginator struct {
	options   DescribeCustomDomainsPaginatorOptions
	client    DescribeCustomDomainsAPIClient
	params    *DescribeCustomDomainsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCustomDomainsPaginator returns a new DescribeCustomDomainsPaginator
func NewDescribeCustomDomainsPaginator(client DescribeCustomDomainsAPIClient, params *DescribeCustomDomainsInput, optFns ...func(*DescribeCustomDomainsPaginatorOptions)) *DescribeCustomDomainsPaginator {
	if params == nil {
		params = &DescribeCustomDomainsInput{}
	}

	options := DescribeCustomDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCustomDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCustomDomainsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCustomDomains page.
func (p *DescribeCustomDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCustomDomainsOutput, error) {
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

	result, err := p.client.DescribeCustomDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeCustomDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DescribeCustomDomains",
	}
}
