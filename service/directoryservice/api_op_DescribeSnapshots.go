// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains information about the directory snapshots that belong to this account.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeSnapshots.NextToken member contains a token that you pass in the next
// call to DescribeSnapshots to retrieve the next set of items. You can also
// specify a maximum number of return results with the Limit parameter.
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshots", params, optFns, c.addOperationDescribeSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeSnapshots operation.
type DescribeSnapshotsInput struct {

	// The identifier of the directory for which to retrieve snapshot information.
	DirectoryId *string

	// The maximum number of objects to return.
	Limit *int32

	// The DescribeSnapshotsResult.NextToken value from a previous call to
	// DescribeSnapshots. Pass null if this is the first call.
	NextToken *string

	// A list of identifiers of the snapshots to obtain the information for. If this
	// member is null or empty, all snapshots are returned using the Limit and
	// NextToken members.
	SnapshotIds []string

	noSmithyDocumentSerde
}

// Contains the results of the DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {

	// If not null, more results are available. Pass this value in the NextToken member
	// of a subsequent call to DescribeSnapshots.
	NextToken *string

	// The list of Snapshot objects that were retrieved. It is possible that this list
	// contains less than the number of items specified in the Limit member of the
	// request. This occurs if there are less than the requested number of items left
	// to retrieve, or if the limitations of the operation have been exceeded.
	Snapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSnapshots{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeSnapshotsAPIClient is a client that implements the DescribeSnapshots
// operation.
type DescribeSnapshotsAPIClient interface {
	DescribeSnapshots(context.Context, *DescribeSnapshotsInput, ...func(*Options)) (*DescribeSnapshotsOutput, error)
}

var _ DescribeSnapshotsAPIClient = (*Client)(nil)

// DescribeSnapshotsPaginatorOptions is the paginator options for DescribeSnapshots
type DescribeSnapshotsPaginatorOptions struct {
	// The maximum number of objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSnapshotsPaginator is a paginator for DescribeSnapshots
type DescribeSnapshotsPaginator struct {
	options   DescribeSnapshotsPaginatorOptions
	client    DescribeSnapshotsAPIClient
	params    *DescribeSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSnapshotsPaginator returns a new DescribeSnapshotsPaginator
func NewDescribeSnapshotsPaginator(client DescribeSnapshotsAPIClient, params *DescribeSnapshotsInput, optFns ...func(*DescribeSnapshotsPaginatorOptions)) *DescribeSnapshotsPaginator {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	options := DescribeSnapshotsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSnapshots page.
func (p *DescribeSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeSnapshots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSnapshots",
	}
}
