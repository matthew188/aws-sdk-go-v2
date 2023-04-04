// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the accelerators for an Amazon Web Services account.
func (c *Client) ListAccelerators(ctx context.Context, params *ListAcceleratorsInput, optFns ...func(*Options)) (*ListAcceleratorsOutput, error) {
	if params == nil {
		params = &ListAcceleratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccelerators", params, optFns, c.addOperationListAcceleratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAcceleratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAcceleratorsInput struct {

	// The number of Global Accelerator objects that you want to return with this call.
	// The default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAcceleratorsOutput struct {

	// The list of accelerators for a customer account.
	Accelerators []types.Accelerator

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAcceleratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccelerators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccelerators{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccelerators(options.Region), middleware.Before); err != nil {
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

// ListAcceleratorsAPIClient is a client that implements the ListAccelerators
// operation.
type ListAcceleratorsAPIClient interface {
	ListAccelerators(context.Context, *ListAcceleratorsInput, ...func(*Options)) (*ListAcceleratorsOutput, error)
}

var _ ListAcceleratorsAPIClient = (*Client)(nil)

// ListAcceleratorsPaginatorOptions is the paginator options for ListAccelerators
type ListAcceleratorsPaginatorOptions struct {
	// The number of Global Accelerator objects that you want to return with this call.
	// The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAcceleratorsPaginator is a paginator for ListAccelerators
type ListAcceleratorsPaginator struct {
	options   ListAcceleratorsPaginatorOptions
	client    ListAcceleratorsAPIClient
	params    *ListAcceleratorsInput
	nextToken *string
	firstPage bool
}

// NewListAcceleratorsPaginator returns a new ListAcceleratorsPaginator
func NewListAcceleratorsPaginator(client ListAcceleratorsAPIClient, params *ListAcceleratorsInput, optFns ...func(*ListAcceleratorsPaginatorOptions)) *ListAcceleratorsPaginator {
	if params == nil {
		params = &ListAcceleratorsInput{}
	}

	options := ListAcceleratorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAcceleratorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAcceleratorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccelerators page.
func (p *ListAcceleratorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAcceleratorsOutput, error) {
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

	result, err := p.client.ListAccelerators(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccelerators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListAccelerators",
	}
}
