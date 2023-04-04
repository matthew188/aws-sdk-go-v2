// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the aliases of a template.
func (c *Client) ListTemplateAliases(ctx context.Context, params *ListTemplateAliasesInput, optFns ...func(*Options)) (*ListTemplateAliasesOutput, error) {
	if params == nil {
		params = &ListTemplateAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplateAliases", params, optFns, c.addOperationListTemplateAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplateAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplateAliasesInput struct {

	// The ID of the Amazon Web Services account that contains the template aliases
	// that you're listing.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTemplateAliasesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// A structure containing the list of the template's aliases.
	TemplateAliasList []types.TemplateAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplateAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplateAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplateAliases{}, middleware.After)
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
	if err = addOpListTemplateAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplateAliases(options.Region), middleware.Before); err != nil {
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

// ListTemplateAliasesAPIClient is a client that implements the ListTemplateAliases
// operation.
type ListTemplateAliasesAPIClient interface {
	ListTemplateAliases(context.Context, *ListTemplateAliasesInput, ...func(*Options)) (*ListTemplateAliasesOutput, error)
}

var _ ListTemplateAliasesAPIClient = (*Client)(nil)

// ListTemplateAliasesPaginatorOptions is the paginator options for
// ListTemplateAliases
type ListTemplateAliasesPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTemplateAliasesPaginator is a paginator for ListTemplateAliases
type ListTemplateAliasesPaginator struct {
	options   ListTemplateAliasesPaginatorOptions
	client    ListTemplateAliasesAPIClient
	params    *ListTemplateAliasesInput
	nextToken *string
	firstPage bool
}

// NewListTemplateAliasesPaginator returns a new ListTemplateAliasesPaginator
func NewListTemplateAliasesPaginator(client ListTemplateAliasesAPIClient, params *ListTemplateAliasesInput, optFns ...func(*ListTemplateAliasesPaginatorOptions)) *ListTemplateAliasesPaginator {
	if params == nil {
		params = &ListTemplateAliasesInput{}
	}

	options := ListTemplateAliasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTemplateAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTemplateAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTemplateAliases page.
func (p *ListTemplateAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTemplateAliasesOutput, error) {
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

	result, err := p.client.ListTemplateAliases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTemplateAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListTemplateAliases",
	}
}
