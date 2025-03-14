// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the available media insights pipeline configurations.
func (c *Client) ListMediaInsightsPipelineConfigurations(ctx context.Context, params *ListMediaInsightsPipelineConfigurationsInput, optFns ...func(*Options)) (*ListMediaInsightsPipelineConfigurationsOutput, error) {
	if params == nil {
		params = &ListMediaInsightsPipelineConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMediaInsightsPipelineConfigurations", params, optFns, c.addOperationListMediaInsightsPipelineConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMediaInsightsPipelineConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMediaInsightsPipelineConfigurationsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMediaInsightsPipelineConfigurationsOutput struct {

	// The requested list of media insights pipeline configurations.
	MediaInsightsPipelineConfigurations []types.MediaInsightsPipelineConfigurationSummary

	// The token used to return the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMediaInsightsPipelineConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMediaInsightsPipelineConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMediaInsightsPipelineConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMediaInsightsPipelineConfigurations(options.Region), middleware.Before); err != nil {
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

// ListMediaInsightsPipelineConfigurationsAPIClient is a client that implements the
// ListMediaInsightsPipelineConfigurations operation.
type ListMediaInsightsPipelineConfigurationsAPIClient interface {
	ListMediaInsightsPipelineConfigurations(context.Context, *ListMediaInsightsPipelineConfigurationsInput, ...func(*Options)) (*ListMediaInsightsPipelineConfigurationsOutput, error)
}

var _ ListMediaInsightsPipelineConfigurationsAPIClient = (*Client)(nil)

// ListMediaInsightsPipelineConfigurationsPaginatorOptions is the paginator options
// for ListMediaInsightsPipelineConfigurations
type ListMediaInsightsPipelineConfigurationsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMediaInsightsPipelineConfigurationsPaginator is a paginator for
// ListMediaInsightsPipelineConfigurations
type ListMediaInsightsPipelineConfigurationsPaginator struct {
	options   ListMediaInsightsPipelineConfigurationsPaginatorOptions
	client    ListMediaInsightsPipelineConfigurationsAPIClient
	params    *ListMediaInsightsPipelineConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListMediaInsightsPipelineConfigurationsPaginator returns a new
// ListMediaInsightsPipelineConfigurationsPaginator
func NewListMediaInsightsPipelineConfigurationsPaginator(client ListMediaInsightsPipelineConfigurationsAPIClient, params *ListMediaInsightsPipelineConfigurationsInput, optFns ...func(*ListMediaInsightsPipelineConfigurationsPaginatorOptions)) *ListMediaInsightsPipelineConfigurationsPaginator {
	if params == nil {
		params = &ListMediaInsightsPipelineConfigurationsInput{}
	}

	options := ListMediaInsightsPipelineConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMediaInsightsPipelineConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMediaInsightsPipelineConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMediaInsightsPipelineConfigurations page.
func (p *ListMediaInsightsPipelineConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMediaInsightsPipelineConfigurationsOutput, error) {
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

	result, err := p.client.ListMediaInsightsPipelineConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMediaInsightsPipelineConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListMediaInsightsPipelineConfigurations",
	}
}
