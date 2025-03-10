// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the instance types for the specified Outpost.
func (c *Client) GetOutpostInstanceTypes(ctx context.Context, params *GetOutpostInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error) {
	if params == nil {
		params = &GetOutpostInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOutpostInstanceTypes", params, optFns, c.addOperationGetOutpostInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOutpostInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutpostInstanceTypesInput struct {

	// The ID or the Amazon Resource Name (ARN) of the Outpost.
	//
	// This member is required.
	OutpostId *string

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type GetOutpostInstanceTypesOutput struct {

	// Information about the instance types.
	InstanceTypes []types.InstanceTypeItem

	// The pagination token.
	NextToken *string

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The ID of the Outpost.
	OutpostId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOutpostInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetOutpostInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOutpostInstanceTypes{}, middleware.After)
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
	if err = addOpGetOutpostInstanceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOutpostInstanceTypes(options.Region), middleware.Before); err != nil {
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

// GetOutpostInstanceTypesAPIClient is a client that implements the
// GetOutpostInstanceTypes operation.
type GetOutpostInstanceTypesAPIClient interface {
	GetOutpostInstanceTypes(context.Context, *GetOutpostInstanceTypesInput, ...func(*Options)) (*GetOutpostInstanceTypesOutput, error)
}

var _ GetOutpostInstanceTypesAPIClient = (*Client)(nil)

// GetOutpostInstanceTypesPaginatorOptions is the paginator options for
// GetOutpostInstanceTypes
type GetOutpostInstanceTypesPaginatorOptions struct {
	// The maximum page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOutpostInstanceTypesPaginator is a paginator for GetOutpostInstanceTypes
type GetOutpostInstanceTypesPaginator struct {
	options   GetOutpostInstanceTypesPaginatorOptions
	client    GetOutpostInstanceTypesAPIClient
	params    *GetOutpostInstanceTypesInput
	nextToken *string
	firstPage bool
}

// NewGetOutpostInstanceTypesPaginator returns a new
// GetOutpostInstanceTypesPaginator
func NewGetOutpostInstanceTypesPaginator(client GetOutpostInstanceTypesAPIClient, params *GetOutpostInstanceTypesInput, optFns ...func(*GetOutpostInstanceTypesPaginatorOptions)) *GetOutpostInstanceTypesPaginator {
	if params == nil {
		params = &GetOutpostInstanceTypesInput{}
	}

	options := GetOutpostInstanceTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOutpostInstanceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOutpostInstanceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOutpostInstanceTypes page.
func (p *GetOutpostInstanceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error) {
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

	result, err := p.client.GetOutpostInstanceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetOutpostInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "GetOutpostInstanceTypes",
	}
}
