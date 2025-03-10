// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of satellites.
func (c *Client) ListSatellites(ctx context.Context, params *ListSatellitesInput, optFns ...func(*Options)) (*ListSatellitesOutput, error) {
	if params == nil {
		params = &ListSatellitesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSatellites", params, optFns, c.addOperationListSatellitesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSatellitesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSatellitesInput struct {

	// Maximum number of satellites returned.
	MaxResults *int32

	// Next token that can be supplied in the next call to get the next page of
	// satellites.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSatellitesOutput struct {

	// Next token that can be supplied in the next call to get the next page of
	// satellites.
	NextToken *string

	// List of satellites.
	Satellites []types.SatelliteListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSatellitesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSatellites{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSatellites{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSatellites(options.Region), middleware.Before); err != nil {
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

// ListSatellitesAPIClient is a client that implements the ListSatellites
// operation.
type ListSatellitesAPIClient interface {
	ListSatellites(context.Context, *ListSatellitesInput, ...func(*Options)) (*ListSatellitesOutput, error)
}

var _ ListSatellitesAPIClient = (*Client)(nil)

// ListSatellitesPaginatorOptions is the paginator options for ListSatellites
type ListSatellitesPaginatorOptions struct {
	// Maximum number of satellites returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSatellitesPaginator is a paginator for ListSatellites
type ListSatellitesPaginator struct {
	options   ListSatellitesPaginatorOptions
	client    ListSatellitesAPIClient
	params    *ListSatellitesInput
	nextToken *string
	firstPage bool
}

// NewListSatellitesPaginator returns a new ListSatellitesPaginator
func NewListSatellitesPaginator(client ListSatellitesAPIClient, params *ListSatellitesInput, optFns ...func(*ListSatellitesPaginatorOptions)) *ListSatellitesPaginator {
	if params == nil {
		params = &ListSatellitesInput{}
	}

	options := ListSatellitesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSatellitesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSatellitesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSatellites page.
func (p *ListSatellitesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSatellitesOutput, error) {
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

	result, err := p.client.ListSatellites(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSatellites(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "ListSatellites",
	}
}
