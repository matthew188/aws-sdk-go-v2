// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about stream keys for the specified channel.
func (c *Client) ListStreamKeys(ctx context.Context, params *ListStreamKeysInput, optFns ...func(*Options)) (*ListStreamKeysOutput, error) {
	if params == nil {
		params = &ListStreamKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamKeys", params, optFns, c.addOperationListStreamKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamKeysInput struct {

	// Channel ARN used to filter the list.
	//
	// This member is required.
	ChannelArn *string

	// Maximum number of streamKeys to return. Default: 1.
	MaxResults int32

	// The first stream key to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStreamKeysOutput struct {

	// List of stream keys.
	//
	// This member is required.
	StreamKeys []types.StreamKeySummary

	// If there are more stream keys than maxResults, use nextToken in the request to
	// get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamKeys{}, middleware.After)
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
	if err = addOpListStreamKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamKeys(options.Region), middleware.Before); err != nil {
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

// ListStreamKeysAPIClient is a client that implements the ListStreamKeys
// operation.
type ListStreamKeysAPIClient interface {
	ListStreamKeys(context.Context, *ListStreamKeysInput, ...func(*Options)) (*ListStreamKeysOutput, error)
}

var _ ListStreamKeysAPIClient = (*Client)(nil)

// ListStreamKeysPaginatorOptions is the paginator options for ListStreamKeys
type ListStreamKeysPaginatorOptions struct {
	// Maximum number of streamKeys to return. Default: 1.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamKeysPaginator is a paginator for ListStreamKeys
type ListStreamKeysPaginator struct {
	options   ListStreamKeysPaginatorOptions
	client    ListStreamKeysAPIClient
	params    *ListStreamKeysInput
	nextToken *string
	firstPage bool
}

// NewListStreamKeysPaginator returns a new ListStreamKeysPaginator
func NewListStreamKeysPaginator(client ListStreamKeysAPIClient, params *ListStreamKeysInput, optFns ...func(*ListStreamKeysPaginatorOptions)) *ListStreamKeysPaginator {
	if params == nil {
		params = &ListStreamKeysInput{}
	}

	options := ListStreamKeysPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamKeys page.
func (p *ListStreamKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListStreamKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreamKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "ListStreamKeys",
	}
}
