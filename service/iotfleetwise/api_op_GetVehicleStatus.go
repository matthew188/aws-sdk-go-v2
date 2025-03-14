// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the status of a vehicle with any associated
// campaigns.
func (c *Client) GetVehicleStatus(ctx context.Context, params *GetVehicleStatusInput, optFns ...func(*Options)) (*GetVehicleStatusOutput, error) {
	if params == nil {
		params = &GetVehicleStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVehicleStatus", params, optFns, c.addOperationGetVehicleStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVehicleStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVehicleStatusInput struct {

	// The ID of the vehicle to retrieve information about.
	//
	// This member is required.
	VehicleName *string

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type GetVehicleStatusOutput struct {

	// Lists information about the state of the vehicle with deployed campaigns.
	Campaigns []types.VehicleStatus

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVehicleStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetVehicleStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetVehicleStatus{}, middleware.After)
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
	if err = addOpGetVehicleStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVehicleStatus(options.Region), middleware.Before); err != nil {
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

// GetVehicleStatusAPIClient is a client that implements the GetVehicleStatus
// operation.
type GetVehicleStatusAPIClient interface {
	GetVehicleStatus(context.Context, *GetVehicleStatusInput, ...func(*Options)) (*GetVehicleStatusOutput, error)
}

var _ GetVehicleStatusAPIClient = (*Client)(nil)

// GetVehicleStatusPaginatorOptions is the paginator options for GetVehicleStatus
type GetVehicleStatusPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetVehicleStatusPaginator is a paginator for GetVehicleStatus
type GetVehicleStatusPaginator struct {
	options   GetVehicleStatusPaginatorOptions
	client    GetVehicleStatusAPIClient
	params    *GetVehicleStatusInput
	nextToken *string
	firstPage bool
}

// NewGetVehicleStatusPaginator returns a new GetVehicleStatusPaginator
func NewGetVehicleStatusPaginator(client GetVehicleStatusAPIClient, params *GetVehicleStatusInput, optFns ...func(*GetVehicleStatusPaginatorOptions)) *GetVehicleStatusPaginator {
	if params == nil {
		params = &GetVehicleStatusInput{}
	}

	options := GetVehicleStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetVehicleStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetVehicleStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetVehicleStatus page.
func (p *GetVehicleStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetVehicleStatusOutput, error) {
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

	result, err := p.client.GetVehicleStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetVehicleStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "GetVehicleStatus",
	}
}
