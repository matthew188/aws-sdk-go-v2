// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of predictors created using the CreateAutoPredictor or
// CreatePredictor operations. For each predictor, this operation returns a summary
// of its properties, including its Amazon Resource Name (ARN). You can retrieve
// the complete set of properties by using the ARN with the DescribeAutoPredictor
// and DescribePredictor operations. You can filter the list using an array of
// Filter objects.
func (c *Client) ListPredictors(ctx context.Context, params *ListPredictorsInput, optFns ...func(*Options)) (*ListPredictorsOutput, error) {
	if params == nil {
		params = &ListPredictorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPredictors", params, optFns, c.addOperationListPredictorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPredictorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPredictorsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the predictors that match the statement from the list,
	// respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	// * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the predictors that match the statement, specify IS. To
	// exclude matching predictors, specify IS_NOT.
	//
	// * Key - The name of the parameter
	// to filter on. Valid values are DatasetGroupArn and Status.
	//
	// * Value - The value
	// to match.
	//
	// For example, to list all predictors whose status is ACTIVE, you would
	// specify: "Filters": [ { "Condition": "IS", "Key": "Status", "Value": "ACTIVE" }
	// ]
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPredictorsOutput struct {

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// An array of objects that summarize each predictor's properties.
	Predictors []types.PredictorSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPredictorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPredictors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPredictors{}, middleware.After)
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
	if err = addOpListPredictorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPredictors(options.Region), middleware.Before); err != nil {
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

// ListPredictorsAPIClient is a client that implements the ListPredictors
// operation.
type ListPredictorsAPIClient interface {
	ListPredictors(context.Context, *ListPredictorsInput, ...func(*Options)) (*ListPredictorsOutput, error)
}

var _ ListPredictorsAPIClient = (*Client)(nil)

// ListPredictorsPaginatorOptions is the paginator options for ListPredictors
type ListPredictorsPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPredictorsPaginator is a paginator for ListPredictors
type ListPredictorsPaginator struct {
	options   ListPredictorsPaginatorOptions
	client    ListPredictorsAPIClient
	params    *ListPredictorsInput
	nextToken *string
	firstPage bool
}

// NewListPredictorsPaginator returns a new ListPredictorsPaginator
func NewListPredictorsPaginator(client ListPredictorsAPIClient, params *ListPredictorsInput, optFns ...func(*ListPredictorsPaginatorOptions)) *ListPredictorsPaginator {
	if params == nil {
		params = &ListPredictorsInput{}
	}

	options := ListPredictorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPredictorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPredictorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPredictors page.
func (p *ListPredictorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPredictorsOutput, error) {
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

	result, err := p.client.ListPredictors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPredictors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListPredictors",
	}
}
