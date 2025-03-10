// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists collaborations the caller owns, is active in, or has been invited to.
func (c *Client) ListCollaborations(ctx context.Context, params *ListCollaborationsInput, optFns ...func(*Options)) (*ListCollaborationsOutput, error) {
	if params == nil {
		params = &ListCollaborationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollaborations", params, optFns, c.addOperationListCollaborationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollaborationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollaborationsInput struct {

	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	MaxResults *int32

	// The caller's status in a collaboration.
	MemberStatus types.FilterableMemberStatus

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollaborationsOutput struct {

	// The list of collaborations.
	//
	// This member is required.
	CollaborationList []types.CollaborationSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollaborationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCollaborations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCollaborations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollaborations(options.Region), middleware.Before); err != nil {
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

// ListCollaborationsAPIClient is a client that implements the ListCollaborations
// operation.
type ListCollaborationsAPIClient interface {
	ListCollaborations(context.Context, *ListCollaborationsInput, ...func(*Options)) (*ListCollaborationsOutput, error)
}

var _ ListCollaborationsAPIClient = (*Client)(nil)

// ListCollaborationsPaginatorOptions is the paginator options for
// ListCollaborations
type ListCollaborationsPaginatorOptions struct {
	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollaborationsPaginator is a paginator for ListCollaborations
type ListCollaborationsPaginator struct {
	options   ListCollaborationsPaginatorOptions
	client    ListCollaborationsAPIClient
	params    *ListCollaborationsInput
	nextToken *string
	firstPage bool
}

// NewListCollaborationsPaginator returns a new ListCollaborationsPaginator
func NewListCollaborationsPaginator(client ListCollaborationsAPIClient, params *ListCollaborationsInput, optFns ...func(*ListCollaborationsPaginatorOptions)) *ListCollaborationsPaginator {
	if params == nil {
		params = &ListCollaborationsInput{}
	}

	options := ListCollaborationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollaborationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollaborationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollaborations page.
func (p *ListCollaborationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollaborationsOutput, error) {
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

	result, err := p.client.ListCollaborations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCollaborations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "ListCollaborations",
	}
}
