// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// X-Ray reevaluates insights periodically until they're resolved, and records each
// intermediate state as an event. You can review an insight's events in the Impact
// Timeline on the Inspect page in the X-Ray console.
func (c *Client) GetInsightEvents(ctx context.Context, params *GetInsightEventsInput, optFns ...func(*Options)) (*GetInsightEventsOutput, error) {
	if params == nil {
		params = &GetInsightEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightEvents", params, optFns, c.addOperationGetInsightEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightEventsInput struct {

	// The insight's unique identifier. Use the GetInsightSummaries action to retrieve
	// an InsightId.
	//
	// This member is required.
	InsightId *string

	// Used to retrieve at most the specified value of events.
	MaxResults *int32

	// Specify the pagination token returned by a previous request to retrieve the next
	// page of events.
	NextToken *string

	noSmithyDocumentSerde
}

type GetInsightEventsOutput struct {

	// A detailed description of the event. This includes the time of the event, client
	// and root cause impact statistics, and the top anomalous service at the time of
	// the event.
	InsightEvents []types.InsightEvent

	// Use this token to retrieve the next page of insight events.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsightEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsightEvents{}, middleware.After)
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
	if err = addOpGetInsightEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightEvents(options.Region), middleware.Before); err != nil {
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

// GetInsightEventsAPIClient is a client that implements the GetInsightEvents
// operation.
type GetInsightEventsAPIClient interface {
	GetInsightEvents(context.Context, *GetInsightEventsInput, ...func(*Options)) (*GetInsightEventsOutput, error)
}

var _ GetInsightEventsAPIClient = (*Client)(nil)

// GetInsightEventsPaginatorOptions is the paginator options for GetInsightEvents
type GetInsightEventsPaginatorOptions struct {
	// Used to retrieve at most the specified value of events.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInsightEventsPaginator is a paginator for GetInsightEvents
type GetInsightEventsPaginator struct {
	options   GetInsightEventsPaginatorOptions
	client    GetInsightEventsAPIClient
	params    *GetInsightEventsInput
	nextToken *string
	firstPage bool
}

// NewGetInsightEventsPaginator returns a new GetInsightEventsPaginator
func NewGetInsightEventsPaginator(client GetInsightEventsAPIClient, params *GetInsightEventsInput, optFns ...func(*GetInsightEventsPaginatorOptions)) *GetInsightEventsPaginator {
	if params == nil {
		params = &GetInsightEventsInput{}
	}

	options := GetInsightEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInsightEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInsightEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInsightEvents page.
func (p *GetInsightEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInsightEventsOutput, error) {
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

	result, err := p.client.GetInsightEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInsightEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetInsightEvents",
	}
}
