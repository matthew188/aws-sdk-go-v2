// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all active backup plans for an authenticated account. The list
// contains information such as Amazon Resource Names (ARNs), plan IDs, creation
// and deletion dates, version IDs, plan names, and creator request IDs.
func (c *Client) ListBackupPlans(ctx context.Context, params *ListBackupPlansInput, optFns ...func(*Options)) (*ListBackupPlansOutput, error) {
	if params == nil {
		params = &ListBackupPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupPlans", params, optFns, c.addOperationListBackupPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupPlansInput struct {

	// A Boolean value with a default value of FALSE that returns deleted backup plans
	// when set to TRUE.
	IncludeDeleted *bool

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBackupPlansOutput struct {

	// An array of backup plan list items containing metadata about your saved backup
	// plans.
	BackupPlansList []types.BackupPlansListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupPlans{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupPlans(options.Region), middleware.Before); err != nil {
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

// ListBackupPlansAPIClient is a client that implements the ListBackupPlans
// operation.
type ListBackupPlansAPIClient interface {
	ListBackupPlans(context.Context, *ListBackupPlansInput, ...func(*Options)) (*ListBackupPlansOutput, error)
}

var _ ListBackupPlansAPIClient = (*Client)(nil)

// ListBackupPlansPaginatorOptions is the paginator options for ListBackupPlans
type ListBackupPlansPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupPlansPaginator is a paginator for ListBackupPlans
type ListBackupPlansPaginator struct {
	options   ListBackupPlansPaginatorOptions
	client    ListBackupPlansAPIClient
	params    *ListBackupPlansInput
	nextToken *string
	firstPage bool
}

// NewListBackupPlansPaginator returns a new ListBackupPlansPaginator
func NewListBackupPlansPaginator(client ListBackupPlansAPIClient, params *ListBackupPlansInput, optFns ...func(*ListBackupPlansPaginatorOptions)) *ListBackupPlansPaginator {
	if params == nil {
		params = &ListBackupPlansInput{}
	}

	options := ListBackupPlansPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupPlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupPlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupPlans page.
func (p *ListBackupPlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupPlansOutput, error) {
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

	result, err := p.client.ListBackupPlans(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBackupPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupPlans",
	}
}
