// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of schemas with minimal details. Schemas in Deleting status will
// not be included in the results. Empty results will be returned if there are no
// schemas available. When the RegistryId is not provided, all the schemas across
// registries will be part of the API response.
func (c *Client) ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) {
	if params == nil {
		params = &ListSchemasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemas", params, optFns, c.addOperationListSchemasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemasInput struct {

	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	// A wrapper structure that may contain the registry name and Amazon Resource Name
	// (ARN).
	RegistryId *types.RegistryId

	noSmithyDocumentSerde
}

type ListSchemasOutput struct {

	// A continuation token for paginating the returned list of tokens, returned if the
	// current segment of the list is not the last.
	NextToken *string

	// An array of SchemaListItem objects containing details of each schema.
	Schemas []types.SchemaListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemas{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemas(options.Region), middleware.Before); err != nil {
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

// ListSchemasAPIClient is a client that implements the ListSchemas operation.
type ListSchemasAPIClient interface {
	ListSchemas(context.Context, *ListSchemasInput, ...func(*Options)) (*ListSchemasOutput, error)
}

var _ ListSchemasAPIClient = (*Client)(nil)

// ListSchemasPaginatorOptions is the paginator options for ListSchemas
type ListSchemasPaginatorOptions struct {
	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchemasPaginator is a paginator for ListSchemas
type ListSchemasPaginator struct {
	options   ListSchemasPaginatorOptions
	client    ListSchemasAPIClient
	params    *ListSchemasInput
	nextToken *string
	firstPage bool
}

// NewListSchemasPaginator returns a new ListSchemasPaginator
func NewListSchemasPaginator(client ListSchemasAPIClient, params *ListSchemasInput, optFns ...func(*ListSchemasPaginatorOptions)) *ListSchemasPaginator {
	if params == nil {
		params = &ListSchemasInput{}
	}

	options := ListSchemasPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchemasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchemasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchemas page.
func (p *ListSchemasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchemasOutput, error) {
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

	result, err := p.client.ListSchemas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchemas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListSchemas",
	}
}
