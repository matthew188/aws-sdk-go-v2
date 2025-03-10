// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcontactlens

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connectcontactlens/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of analysis segments for a real-time analysis session.
func (c *Client) ListRealtimeContactAnalysisSegments(ctx context.Context, params *ListRealtimeContactAnalysisSegmentsInput, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsOutput, error) {
	if params == nil {
		params = &ListRealtimeContactAnalysisSegmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRealtimeContactAnalysisSegments", params, optFns, c.addOperationListRealtimeContactAnalysisSegmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRealtimeContactAnalysisSegmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRealtimeContactAnalysisSegmentsInput struct {

	// The identifier of the contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRealtimeContactAnalysisSegmentsOutput struct {

	// An analyzed transcript or category.
	//
	// This member is required.
	Segments []types.RealtimeContactAnalysisSegment

	// If there are additional results, this is the token for the next set of results.
	// If response includes nextToken there are two possible scenarios:
	//
	// * There are
	// more segments so another call is required to get them.
	//
	// * There are no more
	// segments at this time, but more may be available later (real-time analysis is in
	// progress) so the client should call the operation again to get new segments.
	//
	// If
	// response does not include nextToken, the analysis is completed (successfully or
	// failed) and there are no more segments to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRealtimeContactAnalysisSegmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRealtimeContactAnalysisSegments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRealtimeContactAnalysisSegments{}, middleware.After)
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
	if err = addOpListRealtimeContactAnalysisSegmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRealtimeContactAnalysisSegments(options.Region), middleware.Before); err != nil {
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

// ListRealtimeContactAnalysisSegmentsAPIClient is a client that implements the
// ListRealtimeContactAnalysisSegments operation.
type ListRealtimeContactAnalysisSegmentsAPIClient interface {
	ListRealtimeContactAnalysisSegments(context.Context, *ListRealtimeContactAnalysisSegmentsInput, ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsOutput, error)
}

var _ ListRealtimeContactAnalysisSegmentsAPIClient = (*Client)(nil)

// ListRealtimeContactAnalysisSegmentsPaginatorOptions is the paginator options for
// ListRealtimeContactAnalysisSegments
type ListRealtimeContactAnalysisSegmentsPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRealtimeContactAnalysisSegmentsPaginator is a paginator for
// ListRealtimeContactAnalysisSegments
type ListRealtimeContactAnalysisSegmentsPaginator struct {
	options   ListRealtimeContactAnalysisSegmentsPaginatorOptions
	client    ListRealtimeContactAnalysisSegmentsAPIClient
	params    *ListRealtimeContactAnalysisSegmentsInput
	nextToken *string
	firstPage bool
}

// NewListRealtimeContactAnalysisSegmentsPaginator returns a new
// ListRealtimeContactAnalysisSegmentsPaginator
func NewListRealtimeContactAnalysisSegmentsPaginator(client ListRealtimeContactAnalysisSegmentsAPIClient, params *ListRealtimeContactAnalysisSegmentsInput, optFns ...func(*ListRealtimeContactAnalysisSegmentsPaginatorOptions)) *ListRealtimeContactAnalysisSegmentsPaginator {
	if params == nil {
		params = &ListRealtimeContactAnalysisSegmentsInput{}
	}

	options := ListRealtimeContactAnalysisSegmentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRealtimeContactAnalysisSegmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRealtimeContactAnalysisSegmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRealtimeContactAnalysisSegments page.
func (p *ListRealtimeContactAnalysisSegmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRealtimeContactAnalysisSegmentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRealtimeContactAnalysisSegments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRealtimeContactAnalysisSegments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListRealtimeContactAnalysisSegments",
	}
}
