// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of DC2 ReservedNodeOfferings that matches the payment type,
// term, and usage price of the given DC1 reserved node.
func (c *Client) GetReservedNodeExchangeOfferings(ctx context.Context, params *GetReservedNodeExchangeOfferingsInput, optFns ...func(*Options)) (*GetReservedNodeExchangeOfferingsOutput, error) {
	if params == nil {
		params = &GetReservedNodeExchangeOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservedNodeExchangeOfferings", params, optFns, c.addOperationGetReservedNodeExchangeOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservedNodeExchangeOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservedNodeExchangeOfferingsInput struct {

	// A string representing the node identifier for the DC1 Reserved Node to be
	// exchanged.
	//
	// This member is required.
	ReservedNodeId *string

	// A value that indicates the starting point for the next set of
	// ReservedNodeOfferings.
	Marker *string

	// An integer setting the maximum number of ReservedNodeOfferings to retrieve.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type GetReservedNodeExchangeOfferingsOutput struct {

	// An optional parameter that specifies the starting point for returning a set of
	// response records. When the results of a GetReservedNodeExchangeOfferings request
	// exceed the value specified in MaxRecords, Amazon Redshift returns a value in the
	// marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the marker parameter and retrying the
	// request.
	Marker *string

	// Returns an array of ReservedNodeOffering objects.
	ReservedNodeOfferings []types.ReservedNodeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservedNodeExchangeOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetReservedNodeExchangeOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetReservedNodeExchangeOfferings{}, middleware.After)
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
	if err = addOpGetReservedNodeExchangeOfferingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservedNodeExchangeOfferings(options.Region), middleware.Before); err != nil {
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

// GetReservedNodeExchangeOfferingsAPIClient is a client that implements the
// GetReservedNodeExchangeOfferings operation.
type GetReservedNodeExchangeOfferingsAPIClient interface {
	GetReservedNodeExchangeOfferings(context.Context, *GetReservedNodeExchangeOfferingsInput, ...func(*Options)) (*GetReservedNodeExchangeOfferingsOutput, error)
}

var _ GetReservedNodeExchangeOfferingsAPIClient = (*Client)(nil)

// GetReservedNodeExchangeOfferingsPaginatorOptions is the paginator options for
// GetReservedNodeExchangeOfferings
type GetReservedNodeExchangeOfferingsPaginatorOptions struct {
	// An integer setting the maximum number of ReservedNodeOfferings to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetReservedNodeExchangeOfferingsPaginator is a paginator for
// GetReservedNodeExchangeOfferings
type GetReservedNodeExchangeOfferingsPaginator struct {
	options   GetReservedNodeExchangeOfferingsPaginatorOptions
	client    GetReservedNodeExchangeOfferingsAPIClient
	params    *GetReservedNodeExchangeOfferingsInput
	nextToken *string
	firstPage bool
}

// NewGetReservedNodeExchangeOfferingsPaginator returns a new
// GetReservedNodeExchangeOfferingsPaginator
func NewGetReservedNodeExchangeOfferingsPaginator(client GetReservedNodeExchangeOfferingsAPIClient, params *GetReservedNodeExchangeOfferingsInput, optFns ...func(*GetReservedNodeExchangeOfferingsPaginatorOptions)) *GetReservedNodeExchangeOfferingsPaginator {
	if params == nil {
		params = &GetReservedNodeExchangeOfferingsInput{}
	}

	options := GetReservedNodeExchangeOfferingsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetReservedNodeExchangeOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetReservedNodeExchangeOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetReservedNodeExchangeOfferings page.
func (p *GetReservedNodeExchangeOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetReservedNodeExchangeOfferingsOutput, error) {
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

	result, err := p.client.GetReservedNodeExchangeOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetReservedNodeExchangeOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "GetReservedNodeExchangeOfferings",
	}
}
