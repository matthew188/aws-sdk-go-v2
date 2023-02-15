// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the metadata of either all the lists under the account or the specified
// list.
func (c *Client) GetListsMetadata(ctx context.Context, params *GetListsMetadataInput, optFns ...func(*Options)) (*GetListsMetadataOutput, error) {
	if params == nil {
		params = &GetListsMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetListsMetadata", params, optFns, c.addOperationGetListsMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetListsMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetListsMetadataInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name of the list.
	Name *string

	// The next token for the subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetListsMetadataOutput struct {

	// The metadata of the specified list or all lists under the account.
	Lists []types.AllowDenyList

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetListsMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetListsMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetListsMetadata{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetListsMetadata(options.Region), middleware.Before); err != nil {
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

// GetListsMetadataAPIClient is a client that implements the GetListsMetadata
// operation.
type GetListsMetadataAPIClient interface {
	GetListsMetadata(context.Context, *GetListsMetadataInput, ...func(*Options)) (*GetListsMetadataOutput, error)
}

var _ GetListsMetadataAPIClient = (*Client)(nil)

// GetListsMetadataPaginatorOptions is the paginator options for GetListsMetadata
type GetListsMetadataPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetListsMetadataPaginator is a paginator for GetListsMetadata
type GetListsMetadataPaginator struct {
	options   GetListsMetadataPaginatorOptions
	client    GetListsMetadataAPIClient
	params    *GetListsMetadataInput
	nextToken *string
	firstPage bool
}

// NewGetListsMetadataPaginator returns a new GetListsMetadataPaginator
func NewGetListsMetadataPaginator(client GetListsMetadataAPIClient, params *GetListsMetadataInput, optFns ...func(*GetListsMetadataPaginatorOptions)) *GetListsMetadataPaginator {
	if params == nil {
		params = &GetListsMetadataInput{}
	}

	options := GetListsMetadataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetListsMetadataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetListsMetadataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetListsMetadata page.
func (p *GetListsMetadataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetListsMetadataOutput, error) {
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

	result, err := p.client.GetListsMetadata(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetListsMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetListsMetadata",
	}
}
