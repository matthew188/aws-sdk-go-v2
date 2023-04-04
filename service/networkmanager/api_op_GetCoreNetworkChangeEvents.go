// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a core network change event.
func (c *Client) GetCoreNetworkChangeEvents(ctx context.Context, params *GetCoreNetworkChangeEventsInput, optFns ...func(*Options)) (*GetCoreNetworkChangeEventsOutput, error) {
	if params == nil {
		params = &GetCoreNetworkChangeEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoreNetworkChangeEvents", params, optFns, c.addOperationGetCoreNetworkChangeEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoreNetworkChangeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoreNetworkChangeEventsInput struct {

	// The ID of a core network.
	//
	// This member is required.
	CoreNetworkId *string

	// The ID of the policy version.
	//
	// This member is required.
	PolicyVersionId *int32

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCoreNetworkChangeEventsOutput struct {

	// The response to GetCoreNetworkChangeEventsRequest.
	CoreNetworkChangeEvents []types.CoreNetworkChangeEvent

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCoreNetworkChangeEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCoreNetworkChangeEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCoreNetworkChangeEvents{}, middleware.After)
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
	if err = addOpGetCoreNetworkChangeEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoreNetworkChangeEvents(options.Region), middleware.Before); err != nil {
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

// GetCoreNetworkChangeEventsAPIClient is a client that implements the
// GetCoreNetworkChangeEvents operation.
type GetCoreNetworkChangeEventsAPIClient interface {
	GetCoreNetworkChangeEvents(context.Context, *GetCoreNetworkChangeEventsInput, ...func(*Options)) (*GetCoreNetworkChangeEventsOutput, error)
}

var _ GetCoreNetworkChangeEventsAPIClient = (*Client)(nil)

// GetCoreNetworkChangeEventsPaginatorOptions is the paginator options for
// GetCoreNetworkChangeEvents
type GetCoreNetworkChangeEventsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCoreNetworkChangeEventsPaginator is a paginator for
// GetCoreNetworkChangeEvents
type GetCoreNetworkChangeEventsPaginator struct {
	options   GetCoreNetworkChangeEventsPaginatorOptions
	client    GetCoreNetworkChangeEventsAPIClient
	params    *GetCoreNetworkChangeEventsInput
	nextToken *string
	firstPage bool
}

// NewGetCoreNetworkChangeEventsPaginator returns a new
// GetCoreNetworkChangeEventsPaginator
func NewGetCoreNetworkChangeEventsPaginator(client GetCoreNetworkChangeEventsAPIClient, params *GetCoreNetworkChangeEventsInput, optFns ...func(*GetCoreNetworkChangeEventsPaginatorOptions)) *GetCoreNetworkChangeEventsPaginator {
	if params == nil {
		params = &GetCoreNetworkChangeEventsInput{}
	}

	options := GetCoreNetworkChangeEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCoreNetworkChangeEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCoreNetworkChangeEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCoreNetworkChangeEvents page.
func (p *GetCoreNetworkChangeEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCoreNetworkChangeEventsOutput, error) {
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

	result, err := p.client.GetCoreNetworkChangeEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCoreNetworkChangeEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetCoreNetworkChangeEvents",
	}
}
