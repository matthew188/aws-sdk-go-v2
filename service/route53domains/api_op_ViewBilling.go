// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns all the domain-related billing records for the current Amazon Web
// Services account for a specified period
func (c *Client) ViewBilling(ctx context.Context, params *ViewBillingInput, optFns ...func(*Options)) (*ViewBillingOutput, error) {
	if params == nil {
		params = &ViewBillingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ViewBilling", params, optFns, c.addOperationViewBillingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ViewBillingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ViewBilling request includes the following elements.
type ViewBillingInput struct {

	// The end date and time for the time period for which you want a list of billing
	// records. Specify the date and time in Unix time format and Coordinated Universal
	// time (UTC).
	End *time.Time

	// For an initial request for a list of billing records, omit this element. If the
	// number of billing records that are associated with the current Amazon Web
	// Services account during the specified period is greater than the value that you
	// specified for MaxItems, you can use Marker to return additional billing records.
	// Get the value of NextPageMarker from the previous response, and submit another
	// request that includes the value of NextPageMarker in the Marker element.
	// Constraints: The marker must match the value of NextPageMarker that was returned
	// in the previous response.
	Marker *string

	// The number of billing records to be returned. Default: 20
	MaxItems *int32

	// The beginning date and time for the time period for which you want a list of
	// billing records. Specify the date and time in Unix time format and Coordinated
	// Universal time (UTC).
	Start *time.Time

	noSmithyDocumentSerde
}

// The ViewBilling response includes the following elements.
type ViewBillingOutput struct {

	// A summary of billing records.
	BillingRecords []types.BillingRecord

	// If there are more billing records than you specified for MaxItems in the
	// request, submit another request and include the value of NextPageMarker in the
	// value of Marker.
	NextPageMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationViewBillingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpViewBilling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpViewBilling{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opViewBilling(options.Region), middleware.Before); err != nil {
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

// ViewBillingAPIClient is a client that implements the ViewBilling operation.
type ViewBillingAPIClient interface {
	ViewBilling(context.Context, *ViewBillingInput, ...func(*Options)) (*ViewBillingOutput, error)
}

var _ ViewBillingAPIClient = (*Client)(nil)

// ViewBillingPaginatorOptions is the paginator options for ViewBilling
type ViewBillingPaginatorOptions struct {
	// The number of billing records to be returned. Default: 20
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ViewBillingPaginator is a paginator for ViewBilling
type ViewBillingPaginator struct {
	options   ViewBillingPaginatorOptions
	client    ViewBillingAPIClient
	params    *ViewBillingInput
	nextToken *string
	firstPage bool
}

// NewViewBillingPaginator returns a new ViewBillingPaginator
func NewViewBillingPaginator(client ViewBillingAPIClient, params *ViewBillingInput, optFns ...func(*ViewBillingPaginatorOptions)) *ViewBillingPaginator {
	if params == nil {
		params = &ViewBillingInput{}
	}

	options := ViewBillingPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ViewBillingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ViewBillingPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ViewBilling page.
func (p *ViewBillingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ViewBillingOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ViewBilling(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opViewBilling(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ViewBilling",
	}
}
