// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the permissions an account has to configure Amazon Inspector.
func (c *Client) ListAccountPermissions(ctx context.Context, params *ListAccountPermissionsInput, optFns ...func(*Options)) (*ListAccountPermissionsOutput, error) {
	if params == nil {
		params = &ListAccountPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountPermissions", params, optFns, c.addOperationListAccountPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountPermissionsInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// The service scan type to check permissions for.
	Service types.Service

	noSmithyDocumentSerde
}

type ListAccountPermissionsOutput struct {

	// Contains details on the permissions an account has to configure Amazon
	// Inspector.
	//
	// This member is required.
	Permissions []types.Permission

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccountPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccountPermissions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountPermissions(options.Region), middleware.Before); err != nil {
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

// ListAccountPermissionsAPIClient is a client that implements the
// ListAccountPermissions operation.
type ListAccountPermissionsAPIClient interface {
	ListAccountPermissions(context.Context, *ListAccountPermissionsInput, ...func(*Options)) (*ListAccountPermissionsOutput, error)
}

var _ ListAccountPermissionsAPIClient = (*Client)(nil)

// ListAccountPermissionsPaginatorOptions is the paginator options for
// ListAccountPermissions
type ListAccountPermissionsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountPermissionsPaginator is a paginator for ListAccountPermissions
type ListAccountPermissionsPaginator struct {
	options   ListAccountPermissionsPaginatorOptions
	client    ListAccountPermissionsAPIClient
	params    *ListAccountPermissionsInput
	nextToken *string
	firstPage bool
}

// NewListAccountPermissionsPaginator returns a new ListAccountPermissionsPaginator
func NewListAccountPermissionsPaginator(client ListAccountPermissionsAPIClient, params *ListAccountPermissionsInput, optFns ...func(*ListAccountPermissionsPaginatorOptions)) *ListAccountPermissionsPaginator {
	if params == nil {
		params = &ListAccountPermissionsInput{}
	}

	options := ListAccountPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountPermissions page.
func (p *ListAccountPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountPermissionsOutput, error) {
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

	result, err := p.client.ListAccountPermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccountPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "ListAccountPermissions",
	}
}
