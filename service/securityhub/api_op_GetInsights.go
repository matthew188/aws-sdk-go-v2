// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists and describes insights for the specified insight ARNs.
func (c *Client) GetInsights(ctx context.Context, params *GetInsightsInput, optFns ...func(*Options)) (*GetInsightsOutput, error) {
	if params == nil {
		params = &GetInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsights", params, optFns, c.addOperationGetInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightsInput struct {

	// The ARNs of the insights to describe. If you do not provide any insight ARNs,
	// then GetInsights returns all of your custom insights. It does not return any
	// managed insights.
	InsightArns []string

	// The maximum number of items to return in the response.
	MaxResults int32

	// The token that is required for pagination. On your first call to the GetInsights
	// operation, set the value of this parameter to NULL. For subsequent calls to the
	// operation, to continue listing data, set the value of this parameter to the
	// value returned from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetInsightsOutput struct {

	// The insights returned by the operation.
	//
	// This member is required.
	Insights []types.Insight

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsights{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsights(options.Region), middleware.Before); err != nil {
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

// GetInsightsAPIClient is a client that implements the GetInsights operation.
type GetInsightsAPIClient interface {
	GetInsights(context.Context, *GetInsightsInput, ...func(*Options)) (*GetInsightsOutput, error)
}

var _ GetInsightsAPIClient = (*Client)(nil)

// GetInsightsPaginatorOptions is the paginator options for GetInsights
type GetInsightsPaginatorOptions struct {
	// The maximum number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInsightsPaginator is a paginator for GetInsights
type GetInsightsPaginator struct {
	options   GetInsightsPaginatorOptions
	client    GetInsightsAPIClient
	params    *GetInsightsInput
	nextToken *string
	firstPage bool
}

// NewGetInsightsPaginator returns a new GetInsightsPaginator
func NewGetInsightsPaginator(client GetInsightsAPIClient, params *GetInsightsInput, optFns ...func(*GetInsightsPaginatorOptions)) *GetInsightsPaginator {
	if params == nil {
		params = &GetInsightsInput{}
	}

	options := GetInsightsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInsightsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInsightsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInsights page.
func (p *GetInsightsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInsightsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetInsights(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetInsights",
	}
}
