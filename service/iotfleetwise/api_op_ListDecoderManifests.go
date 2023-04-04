// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists decoder manifests. This API operation uses pagination. Specify the
// nextToken parameter in the request to return more results.
func (c *Client) ListDecoderManifests(ctx context.Context, params *ListDecoderManifestsInput, optFns ...func(*Options)) (*ListDecoderManifestsOutput, error) {
	if params == nil {
		params = &ListDecoderManifestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDecoderManifests", params, optFns, c.addOperationListDecoderManifestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDecoderManifestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDecoderManifestsInput struct {

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated
	// with the decoder manifest.
	ModelManifestArn *string

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDecoderManifestsOutput struct {

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// A list of information about each decoder manifest.
	Summaries []types.DecoderManifestSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDecoderManifestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDecoderManifests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDecoderManifests{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDecoderManifests(options.Region), middleware.Before); err != nil {
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

// ListDecoderManifestsAPIClient is a client that implements the
// ListDecoderManifests operation.
type ListDecoderManifestsAPIClient interface {
	ListDecoderManifests(context.Context, *ListDecoderManifestsInput, ...func(*Options)) (*ListDecoderManifestsOutput, error)
}

var _ ListDecoderManifestsAPIClient = (*Client)(nil)

// ListDecoderManifestsPaginatorOptions is the paginator options for
// ListDecoderManifests
type ListDecoderManifestsPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDecoderManifestsPaginator is a paginator for ListDecoderManifests
type ListDecoderManifestsPaginator struct {
	options   ListDecoderManifestsPaginatorOptions
	client    ListDecoderManifestsAPIClient
	params    *ListDecoderManifestsInput
	nextToken *string
	firstPage bool
}

// NewListDecoderManifestsPaginator returns a new ListDecoderManifestsPaginator
func NewListDecoderManifestsPaginator(client ListDecoderManifestsAPIClient, params *ListDecoderManifestsInput, optFns ...func(*ListDecoderManifestsPaginatorOptions)) *ListDecoderManifestsPaginator {
	if params == nil {
		params = &ListDecoderManifestsInput{}
	}

	options := ListDecoderManifestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDecoderManifestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDecoderManifestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDecoderManifests page.
func (p *ListDecoderManifestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDecoderManifestsOutput, error) {
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

	result, err := p.client.ListDecoderManifests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDecoderManifests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "ListDecoderManifests",
	}
}
