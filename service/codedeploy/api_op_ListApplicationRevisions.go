// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about revisions for an application.
func (c *Client) ListApplicationRevisions(ctx context.Context, params *ListApplicationRevisionsInput, optFns ...func(*Options)) (*ListApplicationRevisionsOutput, error) {
	if params == nil {
		params = &ListApplicationRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationRevisions", params, optFns, c.addOperationListApplicationRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListApplicationRevisions operation.
type ListApplicationRevisionsInput struct {

	// The name of an CodeDeploy application associated with the IAM user or Amazon Web
	// Services account.
	//
	// This member is required.
	ApplicationName *string

	// Whether to list revisions based on whether the revision is the target revision
	// of a deployment group:
	//
	// * include: List revisions that are target revisions of a
	// deployment group.
	//
	// * exclude: Do not list revisions that are target revisions of
	// a deployment group.
	//
	// * ignore: List all revisions.
	Deployed types.ListStateFilterAction

	// An identifier returned from the previous ListApplicationRevisions call. It can
	// be used to return the next set of applications in the list.
	NextToken *string

	// An Amazon S3 bucket name to limit the search for revisions. If set to null, all
	// of the user's buckets are searched.
	S3Bucket *string

	// A key prefix for the set of Amazon S3 objects to limit the search for revisions.
	S3KeyPrefix *string

	// The column name to use to sort the list results:
	//
	// * registerTime: Sort by the
	// time the revisions were registered with CodeDeploy.
	//
	// * firstUsedTime: Sort by
	// the time the revisions were first used in a deployment.
	//
	// * lastUsedTime: Sort by
	// the time the revisions were last used in a deployment.
	//
	// If not specified or set
	// to null, the results are returned in an arbitrary order.
	SortBy types.ApplicationRevisionSortBy

	// The order in which to sort the list results:
	//
	// * ascending: ascending order.
	//
	// *
	// descending: descending order.
	//
	// If not specified, the results are sorted in
	// ascending order. If set to null, the results are sorted in an arbitrary order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

// Represents the output of a ListApplicationRevisions operation.
type ListApplicationRevisionsOutput struct {

	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list application revisions call to return the next
	// set of application revisions in the list.
	NextToken *string

	// A list of locations that contain the matching revisions.
	Revisions []types.RevisionLocation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationRevisions{}, middleware.After)
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
	if err = addOpListApplicationRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationRevisions(options.Region), middleware.Before); err != nil {
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

// ListApplicationRevisionsAPIClient is a client that implements the
// ListApplicationRevisions operation.
type ListApplicationRevisionsAPIClient interface {
	ListApplicationRevisions(context.Context, *ListApplicationRevisionsInput, ...func(*Options)) (*ListApplicationRevisionsOutput, error)
}

var _ ListApplicationRevisionsAPIClient = (*Client)(nil)

// ListApplicationRevisionsPaginatorOptions is the paginator options for
// ListApplicationRevisions
type ListApplicationRevisionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationRevisionsPaginator is a paginator for ListApplicationRevisions
type ListApplicationRevisionsPaginator struct {
	options   ListApplicationRevisionsPaginatorOptions
	client    ListApplicationRevisionsAPIClient
	params    *ListApplicationRevisionsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationRevisionsPaginator returns a new
// ListApplicationRevisionsPaginator
func NewListApplicationRevisionsPaginator(client ListApplicationRevisionsAPIClient, params *ListApplicationRevisionsInput, optFns ...func(*ListApplicationRevisionsPaginatorOptions)) *ListApplicationRevisionsPaginator {
	if params == nil {
		params = &ListApplicationRevisionsInput{}
	}

	options := ListApplicationRevisionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationRevisions page.
func (p *ListApplicationRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationRevisionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListApplicationRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListApplicationRevisions",
	}
}
