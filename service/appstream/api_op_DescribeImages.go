// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes one or more specified images, if the image names
// or image ARNs are provided. Otherwise, all images in the account are described.
func (c *Client) DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
	if params == nil {
		params = &DescribeImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImages", params, optFns, c.addOperationDescribeImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagesInput struct {

	// The ARNs of the public, private, and shared images to describe.
	Arns []string

	// The maximum size of each page of results.
	MaxResults *int32

	// The names of the public or private images to describe.
	Names []string

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The type of image (public, private, or shared) to describe.
	Type types.VisibilityType

	noSmithyDocumentSerde
}

type DescribeImagesOutput struct {

	// Information about the images.
	Images []types.Image

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImages(options.Region), middleware.Before); err != nil {
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

// DescribeImagesAPIClient is a client that implements the DescribeImages
// operation.
type DescribeImagesAPIClient interface {
	DescribeImages(context.Context, *DescribeImagesInput, ...func(*Options)) (*DescribeImagesOutput, error)
}

var _ DescribeImagesAPIClient = (*Client)(nil)

// DescribeImagesPaginatorOptions is the paginator options for DescribeImages
type DescribeImagesPaginatorOptions struct {
	// The maximum size of each page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImagesPaginator is a paginator for DescribeImages
type DescribeImagesPaginator struct {
	options   DescribeImagesPaginatorOptions
	client    DescribeImagesAPIClient
	params    *DescribeImagesInput
	nextToken *string
	firstPage bool
}

// NewDescribeImagesPaginator returns a new DescribeImagesPaginator
func NewDescribeImagesPaginator(client DescribeImagesAPIClient, params *DescribeImagesInput, optFns ...func(*DescribeImagesPaginatorOptions)) *DescribeImagesPaginator {
	if params == nil {
		params = &DescribeImagesInput{}
	}

	options := DescribeImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeImages page.
func (p *DescribeImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
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

	result, err := p.client.DescribeImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeImages",
	}
}
