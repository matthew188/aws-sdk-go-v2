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

// Gets IPAM discovered accounts. A discovered account is an Amazon Web Services
// account that is monitored under a resource discovery. If you have integrated
// IPAM with Amazon Web Services Organizations, all accounts in the organization
// are discovered accounts. Only the IPAM account can get all discovered accounts
// in the organization.
func (c *Client) GetIpamDiscoveredAccounts(ctx context.Context, params *GetIpamDiscoveredAccountsInput, optFns ...func(*Options)) (*GetIpamDiscoveredAccountsOutput, error) {
	if params == nil {
		params = &GetIpamDiscoveredAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIpamDiscoveredAccounts", params, optFns, c.addOperationGetIpamDiscoveredAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIpamDiscoveredAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIpamDiscoveredAccountsInput struct {

	// The Amazon Web Services Region that the account information is returned from.
	//
	// This member is required.
	DiscoveryRegion *string

	// A resource discovery ID.
	//
	// This member is required.
	IpamResourceDiscoveryId *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Discovered account filters.
	Filters []types.Filter

	// The maximum number of discovered accounts to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIpamDiscoveredAccountsOutput struct {

	// Discovered accounts.
	IpamDiscoveredAccounts []types.IpamDiscoveredAccount

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIpamDiscoveredAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetIpamDiscoveredAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetIpamDiscoveredAccounts{}, middleware.After)
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
	if err = addOpGetIpamDiscoveredAccountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIpamDiscoveredAccounts(options.Region), middleware.Before); err != nil {
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

// GetIpamDiscoveredAccountsAPIClient is a client that implements the
// GetIpamDiscoveredAccounts operation.
type GetIpamDiscoveredAccountsAPIClient interface {
	GetIpamDiscoveredAccounts(context.Context, *GetIpamDiscoveredAccountsInput, ...func(*Options)) (*GetIpamDiscoveredAccountsOutput, error)
}

var _ GetIpamDiscoveredAccountsAPIClient = (*Client)(nil)

// GetIpamDiscoveredAccountsPaginatorOptions is the paginator options for
// GetIpamDiscoveredAccounts
type GetIpamDiscoveredAccountsPaginatorOptions struct {
	// The maximum number of discovered accounts to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIpamDiscoveredAccountsPaginator is a paginator for GetIpamDiscoveredAccounts
type GetIpamDiscoveredAccountsPaginator struct {
	options   GetIpamDiscoveredAccountsPaginatorOptions
	client    GetIpamDiscoveredAccountsAPIClient
	params    *GetIpamDiscoveredAccountsInput
	nextToken *string
	firstPage bool
}

// NewGetIpamDiscoveredAccountsPaginator returns a new
// GetIpamDiscoveredAccountsPaginator
func NewGetIpamDiscoveredAccountsPaginator(client GetIpamDiscoveredAccountsAPIClient, params *GetIpamDiscoveredAccountsInput, optFns ...func(*GetIpamDiscoveredAccountsPaginatorOptions)) *GetIpamDiscoveredAccountsPaginator {
	if params == nil {
		params = &GetIpamDiscoveredAccountsInput{}
	}

	options := GetIpamDiscoveredAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIpamDiscoveredAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIpamDiscoveredAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIpamDiscoveredAccounts page.
func (p *GetIpamDiscoveredAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIpamDiscoveredAccountsOutput, error) {
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

	result, err := p.client.GetIpamDiscoveredAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIpamDiscoveredAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetIpamDiscoveredAccounts",
	}
}
