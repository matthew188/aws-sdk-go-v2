// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to find which experiments or launches are using a specified
// segment.
func (c *Client) ListSegmentReferences(ctx context.Context, params *ListSegmentReferencesInput, optFns ...func(*Options)) (*ListSegmentReferencesOutput, error) {
	if params == nil {
		params = &ListSegmentReferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSegmentReferences", params, optFns, c.addOperationListSegmentReferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSegmentReferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSegmentReferencesInput struct {

	// The ARN of the segment that you want to view information for.
	//
	// This member is required.
	Segment *string

	// Specifies whether to return information about launches or experiments that use
	// this segment.
	//
	// This member is required.
	Type types.SegmentReferenceResourceType

	// The maximum number of results to include in the response. If you omit this, the
	// default of 50 is used.
	MaxResults *int32

	// The token to use when requesting the next set of results. You received this
	// token from a previous ListSegmentReferences operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSegmentReferencesOutput struct {

	// The token to use in a subsequent ListSegmentReferences operation to return the
	// next set of results.
	NextToken *string

	// An array of structures, where each structure contains information about one
	// experiment or launch that uses this segment.
	ReferencedBy []types.RefResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSegmentReferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSegmentReferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSegmentReferences{}, middleware.After)
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
	if err = addOpListSegmentReferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSegmentReferences(options.Region), middleware.Before); err != nil {
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

// ListSegmentReferencesAPIClient is a client that implements the
// ListSegmentReferences operation.
type ListSegmentReferencesAPIClient interface {
	ListSegmentReferences(context.Context, *ListSegmentReferencesInput, ...func(*Options)) (*ListSegmentReferencesOutput, error)
}

var _ ListSegmentReferencesAPIClient = (*Client)(nil)

// ListSegmentReferencesPaginatorOptions is the paginator options for
// ListSegmentReferences
type ListSegmentReferencesPaginatorOptions struct {
	// The maximum number of results to include in the response. If you omit this, the
	// default of 50 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSegmentReferencesPaginator is a paginator for ListSegmentReferences
type ListSegmentReferencesPaginator struct {
	options   ListSegmentReferencesPaginatorOptions
	client    ListSegmentReferencesAPIClient
	params    *ListSegmentReferencesInput
	nextToken *string
	firstPage bool
}

// NewListSegmentReferencesPaginator returns a new ListSegmentReferencesPaginator
func NewListSegmentReferencesPaginator(client ListSegmentReferencesAPIClient, params *ListSegmentReferencesInput, optFns ...func(*ListSegmentReferencesPaginatorOptions)) *ListSegmentReferencesPaginator {
	if params == nil {
		params = &ListSegmentReferencesInput{}
	}

	options := ListSegmentReferencesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSegmentReferencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSegmentReferencesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSegmentReferences page.
func (p *ListSegmentReferencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSegmentReferencesOutput, error) {
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

	result, err := p.client.ListSegmentReferences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSegmentReferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "ListSegmentReferences",
	}
}
