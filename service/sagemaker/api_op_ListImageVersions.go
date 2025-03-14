// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the versions of a specified image and their properties. The list can be
// filtered by creation time or modified time.
func (c *Client) ListImageVersions(ctx context.Context, params *ListImageVersionsInput, optFns ...func(*Options)) (*ListImageVersionsOutput, error) {
	if params == nil {
		params = &ListImageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImageVersions", params, optFns, c.addOperationListImageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageVersionsInput struct {

	// The name of the image to list the versions of.
	//
	// This member is required.
	ImageName *string

	// A filter that returns only versions created on or after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only versions created on or before the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only versions modified on or after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only versions modified on or before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of versions to return in the response. The default value is
	// 10.
	MaxResults *int32

	// If the previous call to ListImageVersions didn't return the full set of
	// versions, the call returns a token for getting the next set of versions.
	NextToken *string

	// The property used to sort results. The default value is CREATION_TIME.
	SortBy types.ImageVersionSortBy

	// The sort order. The default value is DESCENDING.
	SortOrder types.ImageVersionSortOrder

	noSmithyDocumentSerde
}

type ListImageVersionsOutput struct {

	// A list of versions and their properties.
	ImageVersions []types.ImageVersion

	// A token for getting the next set of versions, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListImageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListImageVersions{}, middleware.After)
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
	if err = addOpListImageVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImageVersions(options.Region), middleware.Before); err != nil {
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

// ListImageVersionsAPIClient is a client that implements the ListImageVersions
// operation.
type ListImageVersionsAPIClient interface {
	ListImageVersions(context.Context, *ListImageVersionsInput, ...func(*Options)) (*ListImageVersionsOutput, error)
}

var _ ListImageVersionsAPIClient = (*Client)(nil)

// ListImageVersionsPaginatorOptions is the paginator options for ListImageVersions
type ListImageVersionsPaginatorOptions struct {
	// The maximum number of versions to return in the response. The default value is
	// 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImageVersionsPaginator is a paginator for ListImageVersions
type ListImageVersionsPaginator struct {
	options   ListImageVersionsPaginatorOptions
	client    ListImageVersionsAPIClient
	params    *ListImageVersionsInput
	nextToken *string
	firstPage bool
}

// NewListImageVersionsPaginator returns a new ListImageVersionsPaginator
func NewListImageVersionsPaginator(client ListImageVersionsAPIClient, params *ListImageVersionsInput, optFns ...func(*ListImageVersionsPaginatorOptions)) *ListImageVersionsPaginator {
	if params == nil {
		params = &ListImageVersionsInput{}
	}

	options := ListImageVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImageVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImageVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImageVersions page.
func (p *ListImageVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImageVersionsOutput, error) {
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

	result, err := p.client.ListImageVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImageVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListImageVersions",
	}
}
