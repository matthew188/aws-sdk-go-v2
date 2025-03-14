// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the membership details for the specified room in an Amazon Chime
// Enterprise account, such as the members' IDs, email addresses, and names.
func (c *Client) ListRoomMemberships(ctx context.Context, params *ListRoomMembershipsInput, optFns ...func(*Options)) (*ListRoomMembershipsOutput, error) {
	if params == nil {
		params = &ListRoomMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoomMemberships", params, optFns, c.addOperationListRoomMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoomMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoomMembershipsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoomMembershipsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The room membership details.
	RoomMemberships []types.RoomMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoomMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoomMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoomMemberships{}, middleware.After)
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
	if err = addOpListRoomMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoomMemberships(options.Region), middleware.Before); err != nil {
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

// ListRoomMembershipsAPIClient is a client that implements the ListRoomMemberships
// operation.
type ListRoomMembershipsAPIClient interface {
	ListRoomMemberships(context.Context, *ListRoomMembershipsInput, ...func(*Options)) (*ListRoomMembershipsOutput, error)
}

var _ ListRoomMembershipsAPIClient = (*Client)(nil)

// ListRoomMembershipsPaginatorOptions is the paginator options for
// ListRoomMemberships
type ListRoomMembershipsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoomMembershipsPaginator is a paginator for ListRoomMemberships
type ListRoomMembershipsPaginator struct {
	options   ListRoomMembershipsPaginatorOptions
	client    ListRoomMembershipsAPIClient
	params    *ListRoomMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListRoomMembershipsPaginator returns a new ListRoomMembershipsPaginator
func NewListRoomMembershipsPaginator(client ListRoomMembershipsAPIClient, params *ListRoomMembershipsInput, optFns ...func(*ListRoomMembershipsPaginatorOptions)) *ListRoomMembershipsPaginator {
	if params == nil {
		params = &ListRoomMembershipsInput{}
	}

	options := ListRoomMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoomMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoomMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoomMemberships page.
func (p *ListRoomMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoomMembershipsOutput, error) {
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

	result, err := p.client.ListRoomMemberships(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRoomMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListRoomMemberships",
	}
}
