// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx file systems, if a FileSystemIds
// value is provided for that file system. Otherwise, it returns descriptions of
// all file systems owned by your Amazon Web Services account in the Amazon Web
// Services Region of the endpoint that you're calling. When retrieving all file
// system descriptions, you can optionally specify the MaxResults parameter to
// limit the number of descriptions in a response. If more file system descriptions
// remain, Amazon FSx returns a NextToken value in the response. In this case, send
// a later request with the NextToken request parameter set to the value of
// NextToken from the last response. This operation is used in an iterative process
// to retrieve a list of your file system descriptions. DescribeFileSystems is
// called first without a NextTokenvalue. Then the operation continues to be called
// with the NextToken parameter set to the value of the last NextToken value until
// a response has no NextToken. When using this operation, keep the following in
// mind:
//
// * The implementation might return fewer than MaxResults file system
// descriptions while still including a NextToken value.
//
// * The order of file
// systems returned in the response of one DescribeFileSystems call and the order
// of file systems returned across the responses of a multicall iteration is
// unspecified.
func (c *Client) DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
	if params == nil {
		params = &DescribeFileSystemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFileSystems", params, optFns, c.addOperationDescribeFileSystemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFileSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DescribeFileSystems operation.
type DescribeFileSystemsInput struct {

	// IDs of the file systems whose descriptions you want to retrieve (String).
	FileSystemIds []string

	// Maximum number of file systems to return in the response (integer). This
	// parameter value must be greater than 0. The number of items that Amazon FSx
	// returns is the minimum of the MaxResults parameter specified in the request and
	// the service's internal maximum number of items per page.
	MaxResults *int32

	// Opaque pagination token returned from a previous DescribeFileSystems operation
	// (String). If a token present, the operation continues the list from where the
	// returning call left off.
	NextToken *string

	noSmithyDocumentSerde
}

// The response object for DescribeFileSystems operation.
type DescribeFileSystemsOutput struct {

	// An array of file system descriptions.
	FileSystems []types.FileSystem

	// Present if there are more file systems than returned in the response (String).
	// You can use the NextToken value in the later request to fetch the descriptions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFileSystemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFileSystems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFileSystems{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFileSystems(options.Region), middleware.Before); err != nil {
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

// DescribeFileSystemsAPIClient is a client that implements the DescribeFileSystems
// operation.
type DescribeFileSystemsAPIClient interface {
	DescribeFileSystems(context.Context, *DescribeFileSystemsInput, ...func(*Options)) (*DescribeFileSystemsOutput, error)
}

var _ DescribeFileSystemsAPIClient = (*Client)(nil)

// DescribeFileSystemsPaginatorOptions is the paginator options for
// DescribeFileSystems
type DescribeFileSystemsPaginatorOptions struct {
	// Maximum number of file systems to return in the response (integer). This
	// parameter value must be greater than 0. The number of items that Amazon FSx
	// returns is the minimum of the MaxResults parameter specified in the request and
	// the service's internal maximum number of items per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFileSystemsPaginator is a paginator for DescribeFileSystems
type DescribeFileSystemsPaginator struct {
	options   DescribeFileSystemsPaginatorOptions
	client    DescribeFileSystemsAPIClient
	params    *DescribeFileSystemsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFileSystemsPaginator returns a new DescribeFileSystemsPaginator
func NewDescribeFileSystemsPaginator(client DescribeFileSystemsAPIClient, params *DescribeFileSystemsInput, optFns ...func(*DescribeFileSystemsPaginatorOptions)) *DescribeFileSystemsPaginator {
	if params == nil {
		params = &DescribeFileSystemsInput{}
	}

	options := DescribeFileSystemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFileSystemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFileSystemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFileSystems page.
func (p *DescribeFileSystemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
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

	result, err := p.client.DescribeFileSystems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFileSystems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeFileSystems",
	}
}
