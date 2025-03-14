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

// Searches network profiles and lists the ones that meet a set of filter and sort
// criteria.
func (c *Client) SearchNetworkProfiles(ctx context.Context, params *SearchNetworkProfilesInput, optFns ...func(*Options)) (*SearchNetworkProfilesOutput, error) {
	if params == nil {
		params = &SearchNetworkProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchNetworkProfiles", params, optFns, c.addOperationSearchNetworkProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchNetworkProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchNetworkProfilesInput struct {

	// The filters to use to list a specified set of network profiles. Valid filters
	// are NetworkProfileName, Ssid, and SecurityType.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use to list the specified set of network profiles. Valid sort
	// criteria includes NetworkProfileName, Ssid, and SecurityType.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchNetworkProfilesOutput struct {

	// The network profiles that meet the specified set of filter criteria, in sort
	// order. It is a list of NetworkProfileData objects.
	NetworkProfiles []types.NetworkProfileData

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The total number of network profiles returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchNetworkProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchNetworkProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchNetworkProfiles{}, middleware.After)
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
	if err = addOpSearchNetworkProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchNetworkProfiles(options.Region), middleware.Before); err != nil {
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

// SearchNetworkProfilesAPIClient is a client that implements the
// SearchNetworkProfiles operation.
type SearchNetworkProfilesAPIClient interface {
	SearchNetworkProfiles(context.Context, *SearchNetworkProfilesInput, ...func(*Options)) (*SearchNetworkProfilesOutput, error)
}

var _ SearchNetworkProfilesAPIClient = (*Client)(nil)

// SearchNetworkProfilesPaginatorOptions is the paginator options for
// SearchNetworkProfiles
type SearchNetworkProfilesPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchNetworkProfilesPaginator is a paginator for SearchNetworkProfiles
type SearchNetworkProfilesPaginator struct {
	options   SearchNetworkProfilesPaginatorOptions
	client    SearchNetworkProfilesAPIClient
	params    *SearchNetworkProfilesInput
	nextToken *string
	firstPage bool
}

// NewSearchNetworkProfilesPaginator returns a new SearchNetworkProfilesPaginator
func NewSearchNetworkProfilesPaginator(client SearchNetworkProfilesAPIClient, params *SearchNetworkProfilesInput, optFns ...func(*SearchNetworkProfilesPaginatorOptions)) *SearchNetworkProfilesPaginator {
	if params == nil {
		params = &SearchNetworkProfilesInput{}
	}

	options := SearchNetworkProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchNetworkProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchNetworkProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchNetworkProfiles page.
func (p *SearchNetworkProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchNetworkProfilesOutput, error) {
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

	result, err := p.client.SearchNetworkProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchNetworkProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchNetworkProfiles",
	}
}
