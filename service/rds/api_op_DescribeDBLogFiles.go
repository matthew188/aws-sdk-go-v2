// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DB log files for the DB instance. This command doesn't apply
// to RDS Custom.
func (c *Client) DescribeDBLogFiles(ctx context.Context, params *DescribeDBLogFilesInput, optFns ...func(*Options)) (*DescribeDBLogFilesOutput, error) {
	if params == nil {
		params = &DescribeDBLogFilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBLogFiles", params, optFns, c.addOperationDescribeDBLogFilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBLogFilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBLogFilesInput struct {

	// The customer-assigned name of the DB instance that contains the log files you
	// want to list. Constraints:
	//
	// * Must match the identifier of an existing
	// DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Filters the available log files for files written since the specified date, in
	// POSIX timestamp format with milliseconds.
	FileLastWritten int64

	// Filters the available log files for files larger than the specified size.
	FileSize int64

	// Filters the available log files for log file names that contain the specified
	// string.
	FilenameContains *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// The pagination token provided in the previous request. If this parameter is
	// specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// The response from a call to DescribeDBLogFiles.
type DescribeDBLogFilesOutput struct {

	// The DB log files returned.
	DescribeDBLogFiles []types.DescribeDBLogFilesDetails

	// A pagination token that can be used in a later DescribeDBLogFiles request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBLogFilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBLogFiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBLogFiles{}, middleware.After)
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
	if err = addOpDescribeDBLogFilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBLogFiles(options.Region), middleware.Before); err != nil {
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

// DescribeDBLogFilesAPIClient is a client that implements the DescribeDBLogFiles
// operation.
type DescribeDBLogFilesAPIClient interface {
	DescribeDBLogFiles(context.Context, *DescribeDBLogFilesInput, ...func(*Options)) (*DescribeDBLogFilesOutput, error)
}

var _ DescribeDBLogFilesAPIClient = (*Client)(nil)

// DescribeDBLogFilesPaginatorOptions is the paginator options for
// DescribeDBLogFiles
type DescribeDBLogFilesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBLogFilesPaginator is a paginator for DescribeDBLogFiles
type DescribeDBLogFilesPaginator struct {
	options   DescribeDBLogFilesPaginatorOptions
	client    DescribeDBLogFilesAPIClient
	params    *DescribeDBLogFilesInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBLogFilesPaginator returns a new DescribeDBLogFilesPaginator
func NewDescribeDBLogFilesPaginator(client DescribeDBLogFilesAPIClient, params *DescribeDBLogFilesInput, optFns ...func(*DescribeDBLogFilesPaginatorOptions)) *DescribeDBLogFilesPaginator {
	if params == nil {
		params = &DescribeDBLogFilesInput{}
	}

	options := DescribeDBLogFilesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBLogFilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBLogFilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBLogFiles page.
func (p *DescribeDBLogFilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBLogFilesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBLogFiles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBLogFiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBLogFiles",
	}
}
