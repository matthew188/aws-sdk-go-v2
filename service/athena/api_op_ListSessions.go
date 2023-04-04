// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the sessions in a workgroup that are in an active state like CREATING,
// CREATED, IDLE, or BUSY. Newer sessions are listed first; older sessions are
// listed later.
func (c *Client) ListSessions(ctx context.Context, params *ListSessionsInput, optFns ...func(*Options)) (*ListSessionsOutput, error) {
	if params == nil {
		params = &ListSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSessions", params, optFns, c.addOperationListSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSessionsInput struct {

	// The workgroup to which the session belongs.
	//
	// This member is required.
	WorkGroup *string

	// The maximum number of sessions to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// A filter for a specific session state. A description of each state follows.
	// CREATING - The session is being started, including acquiring resources. CREATED
	// - The session has been started. IDLE - The session is able to accept a
	// calculation. BUSY - The session is processing another task and is unable to
	// accept a calculation. TERMINATING - The session is in the process of shutting
	// down. TERMINATED - The session and its resources are no longer running. DEGRADED
	// - The session has no healthy coordinators. FAILED - Due to a failure, the
	// session and its resources are no longer running.
	StateFilter types.SessionState

	noSmithyDocumentSerde
}

type ListSessionsOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// A list of sessions.
	Sessions []types.SessionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSessions{}, middleware.After)
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
	if err = addOpListSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSessions(options.Region), middleware.Before); err != nil {
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

// ListSessionsAPIClient is a client that implements the ListSessions operation.
type ListSessionsAPIClient interface {
	ListSessions(context.Context, *ListSessionsInput, ...func(*Options)) (*ListSessionsOutput, error)
}

var _ ListSessionsAPIClient = (*Client)(nil)

// ListSessionsPaginatorOptions is the paginator options for ListSessions
type ListSessionsPaginatorOptions struct {
	// The maximum number of sessions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSessionsPaginator is a paginator for ListSessions
type ListSessionsPaginator struct {
	options   ListSessionsPaginatorOptions
	client    ListSessionsAPIClient
	params    *ListSessionsInput
	nextToken *string
	firstPage bool
}

// NewListSessionsPaginator returns a new ListSessionsPaginator
func NewListSessionsPaginator(client ListSessionsAPIClient, params *ListSessionsInput, optFns ...func(*ListSessionsPaginatorOptions)) *ListSessionsPaginator {
	if params == nil {
		params = &ListSessionsInput{}
	}

	options := ListSessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSessions page.
func (p *ListSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSessionsOutput, error) {
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

	result, err := p.client.ListSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ListSessions",
	}
}
