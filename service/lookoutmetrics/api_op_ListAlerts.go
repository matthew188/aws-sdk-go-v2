// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the alerts attached to a detector. Amazon Lookout for Metrics API actions
// are eventually consistent. If you do a read operation on a resource immediately
// after creating or modifying it, use retries to allow time for the write
// operation to complete.
func (c *Client) ListAlerts(ctx context.Context, params *ListAlertsInput, optFns ...func(*Options)) (*ListAlertsOutput, error) {
	if params == nil {
		params = &ListAlertsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAlerts", params, optFns, c.addOperationListAlertsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAlertsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlertsInput struct {

	// The ARN of the alert's detector.
	AnomalyDetectorArn *string

	// The maximum number of results that will be displayed by the request.
	MaxResults *int32

	// If the result of the previous request is truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAlertsOutput struct {

	// Contains information about an alert.
	AlertSummaryList []types.AlertSummary

	// If the response is truncated, the service returns this token. To retrieve the
	// next set of results, use this token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAlertsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAlerts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAlerts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAlerts(options.Region), middleware.Before); err != nil {
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

// ListAlertsAPIClient is a client that implements the ListAlerts operation.
type ListAlertsAPIClient interface {
	ListAlerts(context.Context, *ListAlertsInput, ...func(*Options)) (*ListAlertsOutput, error)
}

var _ ListAlertsAPIClient = (*Client)(nil)

// ListAlertsPaginatorOptions is the paginator options for ListAlerts
type ListAlertsPaginatorOptions struct {
	// The maximum number of results that will be displayed by the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAlertsPaginator is a paginator for ListAlerts
type ListAlertsPaginator struct {
	options   ListAlertsPaginatorOptions
	client    ListAlertsAPIClient
	params    *ListAlertsInput
	nextToken *string
	firstPage bool
}

// NewListAlertsPaginator returns a new ListAlertsPaginator
func NewListAlertsPaginator(client ListAlertsAPIClient, params *ListAlertsInput, optFns ...func(*ListAlertsPaginatorOptions)) *ListAlertsPaginator {
	if params == nil {
		params = &ListAlertsInput{}
	}

	options := ListAlertsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAlertsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAlertsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAlerts page.
func (p *ListAlertsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAlertsOutput, error) {
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

	result, err := p.client.ListAlerts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAlerts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "ListAlerts",
	}
}
