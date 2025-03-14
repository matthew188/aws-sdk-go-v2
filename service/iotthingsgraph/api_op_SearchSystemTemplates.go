// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for summary information about systems in the user's account. You can
// filter by the ID of a workflow to return only systems that use the specified
// workflow.
//
// Deprecated: since: 2022-08-30
func (c *Client) SearchSystemTemplates(ctx context.Context, params *SearchSystemTemplatesInput, optFns ...func(*Options)) (*SearchSystemTemplatesOutput, error) {
	if params == nil {
		params = &SearchSystemTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSystemTemplates", params, optFns, c.addOperationSearchSystemTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSystemTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSystemTemplatesInput struct {

	// An array of filters that limit the result set. The only valid filter is
	// FLOW_TEMPLATE_ID.
	Filters []types.SystemTemplateFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchSystemTemplatesOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary information about each system
	// deployment in the result set.
	Summaries []types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchSystemTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSystemTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSystemTemplates{}, middleware.After)
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
	if err = addOpSearchSystemTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSystemTemplates(options.Region), middleware.Before); err != nil {
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

// SearchSystemTemplatesAPIClient is a client that implements the
// SearchSystemTemplates operation.
type SearchSystemTemplatesAPIClient interface {
	SearchSystemTemplates(context.Context, *SearchSystemTemplatesInput, ...func(*Options)) (*SearchSystemTemplatesOutput, error)
}

var _ SearchSystemTemplatesAPIClient = (*Client)(nil)

// SearchSystemTemplatesPaginatorOptions is the paginator options for
// SearchSystemTemplates
type SearchSystemTemplatesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchSystemTemplatesPaginator is a paginator for SearchSystemTemplates
type SearchSystemTemplatesPaginator struct {
	options   SearchSystemTemplatesPaginatorOptions
	client    SearchSystemTemplatesAPIClient
	params    *SearchSystemTemplatesInput
	nextToken *string
	firstPage bool
}

// NewSearchSystemTemplatesPaginator returns a new SearchSystemTemplatesPaginator
func NewSearchSystemTemplatesPaginator(client SearchSystemTemplatesAPIClient, params *SearchSystemTemplatesInput, optFns ...func(*SearchSystemTemplatesPaginatorOptions)) *SearchSystemTemplatesPaginator {
	if params == nil {
		params = &SearchSystemTemplatesInput{}
	}

	options := SearchSystemTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchSystemTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchSystemTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchSystemTemplates page.
func (p *SearchSystemTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchSystemTemplatesOutput, error) {
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

	result, err := p.client.SearchSystemTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchSystemTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchSystemTemplates",
	}
}
