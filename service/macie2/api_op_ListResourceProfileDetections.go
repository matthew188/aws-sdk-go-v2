// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the types and amount of sensitive data that Amazon
// Macie found in an S3 bucket.
func (c *Client) ListResourceProfileDetections(ctx context.Context, params *ListResourceProfileDetectionsInput, optFns ...func(*Options)) (*ListResourceProfileDetectionsOutput, error) {
	if params == nil {
		params = &ListResourceProfileDetectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceProfileDetections", params, optFns, c.addOperationListResourceProfileDetectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceProfileDetectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceProfileDetectionsInput struct {

	// The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.
	//
	// This member is required.
	ResourceArn *string

	// The maximum number of items to include in each page of a paginated response.
	MaxResults int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceProfileDetectionsOutput struct {

	// An array of objects, one for each type of sensitive data that Amazon Macie found
	// in the bucket. Each object reports the number of occurrences of the specified
	// type and provides information about the custom data identifier or managed data
	// identifier that detected the data.
	Detections []types.Detection

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceProfileDetectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResourceProfileDetections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResourceProfileDetections{}, middleware.After)
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
	if err = addOpListResourceProfileDetectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceProfileDetections(options.Region), middleware.Before); err != nil {
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

// ListResourceProfileDetectionsAPIClient is a client that implements the
// ListResourceProfileDetections operation.
type ListResourceProfileDetectionsAPIClient interface {
	ListResourceProfileDetections(context.Context, *ListResourceProfileDetectionsInput, ...func(*Options)) (*ListResourceProfileDetectionsOutput, error)
}

var _ ListResourceProfileDetectionsAPIClient = (*Client)(nil)

// ListResourceProfileDetectionsPaginatorOptions is the paginator options for
// ListResourceProfileDetections
type ListResourceProfileDetectionsPaginatorOptions struct {
	// The maximum number of items to include in each page of a paginated response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceProfileDetectionsPaginator is a paginator for
// ListResourceProfileDetections
type ListResourceProfileDetectionsPaginator struct {
	options   ListResourceProfileDetectionsPaginatorOptions
	client    ListResourceProfileDetectionsAPIClient
	params    *ListResourceProfileDetectionsInput
	nextToken *string
	firstPage bool
}

// NewListResourceProfileDetectionsPaginator returns a new
// ListResourceProfileDetectionsPaginator
func NewListResourceProfileDetectionsPaginator(client ListResourceProfileDetectionsAPIClient, params *ListResourceProfileDetectionsInput, optFns ...func(*ListResourceProfileDetectionsPaginatorOptions)) *ListResourceProfileDetectionsPaginator {
	if params == nil {
		params = &ListResourceProfileDetectionsInput{}
	}

	options := ListResourceProfileDetectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceProfileDetectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceProfileDetectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceProfileDetections page.
func (p *ListResourceProfileDetectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceProfileDetectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListResourceProfileDetections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceProfileDetections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListResourceProfileDetections",
	}
}
