// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of resources of the specified type that have been analyzed by
// the specified analyzer..
func (c *Client) ListAnalyzedResources(ctx context.Context, params *ListAnalyzedResourcesInput, optFns ...func(*Options)) (*ListAnalyzedResourcesOutput, error) {
	if params == nil {
		params = &ListAnalyzedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalyzedResources", params, optFns, c.addOperationListAnalyzedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalyzedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves a list of resources that have been analyzed.
type ListAnalyzedResourcesInput struct {

	// The ARN of the analyzer
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources)
	// to retrieve a list of analyzed resources from.
	//
	// This member is required.
	AnalyzerArn *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	// The type of resource.
	ResourceType types.ResourceType

	noSmithyDocumentSerde
}

// The response to the request.
type ListAnalyzedResourcesOutput struct {

	// A list of resources that were analyzed.
	//
	// This member is required.
	AnalyzedResources []types.AnalyzedResourceSummary

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnalyzedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalyzedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalyzedResources{}, middleware.After)
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
	if err = addOpListAnalyzedResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnalyzedResources(options.Region), middleware.Before); err != nil {
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

// ListAnalyzedResourcesAPIClient is a client that implements the
// ListAnalyzedResources operation.
type ListAnalyzedResourcesAPIClient interface {
	ListAnalyzedResources(context.Context, *ListAnalyzedResourcesInput, ...func(*Options)) (*ListAnalyzedResourcesOutput, error)
}

var _ ListAnalyzedResourcesAPIClient = (*Client)(nil)

// ListAnalyzedResourcesPaginatorOptions is the paginator options for
// ListAnalyzedResources
type ListAnalyzedResourcesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnalyzedResourcesPaginator is a paginator for ListAnalyzedResources
type ListAnalyzedResourcesPaginator struct {
	options   ListAnalyzedResourcesPaginatorOptions
	client    ListAnalyzedResourcesAPIClient
	params    *ListAnalyzedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListAnalyzedResourcesPaginator returns a new ListAnalyzedResourcesPaginator
func NewListAnalyzedResourcesPaginator(client ListAnalyzedResourcesAPIClient, params *ListAnalyzedResourcesInput, optFns ...func(*ListAnalyzedResourcesPaginatorOptions)) *ListAnalyzedResourcesPaginator {
	if params == nil {
		params = &ListAnalyzedResourcesInput{}
	}

	options := ListAnalyzedResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnalyzedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnalyzedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnalyzedResources page.
func (p *ListAnalyzedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnalyzedResourcesOutput, error) {
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

	result, err := p.client.ListAnalyzedResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnalyzedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "ListAnalyzedResources",
	}
}
