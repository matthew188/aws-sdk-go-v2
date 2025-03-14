// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all entity types or a specific entity type if a name is specified. This is
// a paginated API. If you provide a null maxResults, this action retrieves a
// maximum of 10 records per page. If you provide a maxResults, the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetEntityTypesResponse as part of your request. A null pagination token
// fetches the records from the beginning.
func (c *Client) GetEntityTypes(ctx context.Context, params *GetEntityTypesInput, optFns ...func(*Options)) (*GetEntityTypesOutput, error) {
	if params == nil {
		params = &GetEntityTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEntityTypes", params, optFns, c.addOperationGetEntityTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEntityTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEntityTypesInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name.
	Name *string

	// The next token for the subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEntityTypesOutput struct {

	// An array of entity types.
	EntityTypes []types.EntityType

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEntityTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEntityTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEntityTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEntityTypes(options.Region), middleware.Before); err != nil {
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

// GetEntityTypesAPIClient is a client that implements the GetEntityTypes
// operation.
type GetEntityTypesAPIClient interface {
	GetEntityTypes(context.Context, *GetEntityTypesInput, ...func(*Options)) (*GetEntityTypesOutput, error)
}

var _ GetEntityTypesAPIClient = (*Client)(nil)

// GetEntityTypesPaginatorOptions is the paginator options for GetEntityTypes
type GetEntityTypesPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEntityTypesPaginator is a paginator for GetEntityTypes
type GetEntityTypesPaginator struct {
	options   GetEntityTypesPaginatorOptions
	client    GetEntityTypesAPIClient
	params    *GetEntityTypesInput
	nextToken *string
	firstPage bool
}

// NewGetEntityTypesPaginator returns a new GetEntityTypesPaginator
func NewGetEntityTypesPaginator(client GetEntityTypesAPIClient, params *GetEntityTypesInput, optFns ...func(*GetEntityTypesPaginatorOptions)) *GetEntityTypesPaginator {
	if params == nil {
		params = &GetEntityTypesInput{}
	}

	options := GetEntityTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEntityTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEntityTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetEntityTypes page.
func (p *GetEntityTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEntityTypesOutput, error) {
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

	result, err := p.client.GetEntityTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetEntityTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetEntityTypes",
	}
}
