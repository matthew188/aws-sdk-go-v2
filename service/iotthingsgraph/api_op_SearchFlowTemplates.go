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

// Searches for summary information about workflows.
//
// Deprecated: since: 2022-08-30
func (c *Client) SearchFlowTemplates(ctx context.Context, params *SearchFlowTemplatesInput, optFns ...func(*Options)) (*SearchFlowTemplatesOutput, error) {
	if params == nil {
		params = &SearchFlowTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchFlowTemplates", params, optFns, c.addOperationSearchFlowTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchFlowTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFlowTemplatesInput struct {

	// An array of objects that limit the result set. The only valid filter is
	// DEVICE_MODEL_ID.
	Filters []types.FlowTemplateFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchFlowTemplatesOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary information about each workflow in the
	// result set.
	Summaries []types.FlowTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchFlowTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchFlowTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchFlowTemplates{}, middleware.After)
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
	if err = addOpSearchFlowTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFlowTemplates(options.Region), middleware.Before); err != nil {
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

// SearchFlowTemplatesAPIClient is a client that implements the SearchFlowTemplates
// operation.
type SearchFlowTemplatesAPIClient interface {
	SearchFlowTemplates(context.Context, *SearchFlowTemplatesInput, ...func(*Options)) (*SearchFlowTemplatesOutput, error)
}

var _ SearchFlowTemplatesAPIClient = (*Client)(nil)

// SearchFlowTemplatesPaginatorOptions is the paginator options for
// SearchFlowTemplates
type SearchFlowTemplatesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchFlowTemplatesPaginator is a paginator for SearchFlowTemplates
type SearchFlowTemplatesPaginator struct {
	options   SearchFlowTemplatesPaginatorOptions
	client    SearchFlowTemplatesAPIClient
	params    *SearchFlowTemplatesInput
	nextToken *string
	firstPage bool
}

// NewSearchFlowTemplatesPaginator returns a new SearchFlowTemplatesPaginator
func NewSearchFlowTemplatesPaginator(client SearchFlowTemplatesAPIClient, params *SearchFlowTemplatesInput, optFns ...func(*SearchFlowTemplatesPaginatorOptions)) *SearchFlowTemplatesPaginator {
	if params == nil {
		params = &SearchFlowTemplatesInput{}
	}

	options := SearchFlowTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchFlowTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchFlowTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchFlowTemplates page.
func (p *SearchFlowTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchFlowTemplatesOutput, error) {
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

	result, err := p.client.SearchFlowTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchFlowTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchFlowTemplates",
	}
}
