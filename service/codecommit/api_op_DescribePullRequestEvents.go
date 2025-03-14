// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about one or more pull request events.
func (c *Client) DescribePullRequestEvents(ctx context.Context, params *DescribePullRequestEventsInput, optFns ...func(*Options)) (*DescribePullRequestEventsOutput, error) {
	if params == nil {
		params = &DescribePullRequestEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePullRequestEvents", params, optFns, c.addOperationDescribePullRequestEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePullRequestEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePullRequestEventsInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests.
	//
	// This member is required.
	PullRequestId *string

	// The Amazon Resource Name (ARN) of the user whose actions resulted in the event.
	// Examples include updating the pull request with more commits or changing the
	// status of a pull request.
	ActorArn *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 events, which is also the maximum number of events that can
	// be returned in a result.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// Optional. The pull request event type about which you want to return
	// information.
	PullRequestEventType types.PullRequestEventType

	noSmithyDocumentSerde
}

type DescribePullRequestEventsOutput struct {

	// Information about the pull request events.
	//
	// This member is required.
	PullRequestEvents []types.PullRequestEvent

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePullRequestEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePullRequestEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePullRequestEvents{}, middleware.After)
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
	if err = addOpDescribePullRequestEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePullRequestEvents(options.Region), middleware.Before); err != nil {
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

// DescribePullRequestEventsAPIClient is a client that implements the
// DescribePullRequestEvents operation.
type DescribePullRequestEventsAPIClient interface {
	DescribePullRequestEvents(context.Context, *DescribePullRequestEventsInput, ...func(*Options)) (*DescribePullRequestEventsOutput, error)
}

var _ DescribePullRequestEventsAPIClient = (*Client)(nil)

// DescribePullRequestEventsPaginatorOptions is the paginator options for
// DescribePullRequestEvents
type DescribePullRequestEventsPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 events, which is also the maximum number of events that can
	// be returned in a result.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePullRequestEventsPaginator is a paginator for DescribePullRequestEvents
type DescribePullRequestEventsPaginator struct {
	options   DescribePullRequestEventsPaginatorOptions
	client    DescribePullRequestEventsAPIClient
	params    *DescribePullRequestEventsInput
	nextToken *string
	firstPage bool
}

// NewDescribePullRequestEventsPaginator returns a new
// DescribePullRequestEventsPaginator
func NewDescribePullRequestEventsPaginator(client DescribePullRequestEventsAPIClient, params *DescribePullRequestEventsInput, optFns ...func(*DescribePullRequestEventsPaginatorOptions)) *DescribePullRequestEventsPaginator {
	if params == nil {
		params = &DescribePullRequestEventsInput{}
	}

	options := DescribePullRequestEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePullRequestEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePullRequestEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePullRequestEvents page.
func (p *DescribePullRequestEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePullRequestEventsOutput, error) {
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

	result, err := p.client.DescribePullRequestEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePullRequestEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DescribePullRequestEvents",
	}
}
