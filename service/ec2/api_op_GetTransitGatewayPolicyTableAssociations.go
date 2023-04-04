// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the transit gateway policy table associations.
func (c *Client) GetTransitGatewayPolicyTableAssociations(ctx context.Context, params *GetTransitGatewayPolicyTableAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayPolicyTableAssociationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayPolicyTableAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayPolicyTableAssociations", params, optFns, c.addOperationGetTransitGatewayPolicyTableAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayPolicyTableAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayPolicyTableAssociationsInput struct {

	// The ID of the transit gateway policy table.
	//
	// This member is required.
	TransitGatewayPolicyTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters associated with the transit gateway policy table.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTransitGatewayPolicyTableAssociationsOutput struct {

	// Returns details about the transit gateway policy table association.
	Associations []types.TransitGatewayPolicyTableAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransitGatewayPolicyTableAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayPolicyTableAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayPolicyTableAssociations{}, middleware.After)
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
	if err = addOpGetTransitGatewayPolicyTableAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayPolicyTableAssociations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayPolicyTableAssociationsAPIClient is a client that implements
// the GetTransitGatewayPolicyTableAssociations operation.
type GetTransitGatewayPolicyTableAssociationsAPIClient interface {
	GetTransitGatewayPolicyTableAssociations(context.Context, *GetTransitGatewayPolicyTableAssociationsInput, ...func(*Options)) (*GetTransitGatewayPolicyTableAssociationsOutput, error)
}

var _ GetTransitGatewayPolicyTableAssociationsAPIClient = (*Client)(nil)

// GetTransitGatewayPolicyTableAssociationsPaginatorOptions is the paginator
// options for GetTransitGatewayPolicyTableAssociations
type GetTransitGatewayPolicyTableAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayPolicyTableAssociationsPaginator is a paginator for
// GetTransitGatewayPolicyTableAssociations
type GetTransitGatewayPolicyTableAssociationsPaginator struct {
	options   GetTransitGatewayPolicyTableAssociationsPaginatorOptions
	client    GetTransitGatewayPolicyTableAssociationsAPIClient
	params    *GetTransitGatewayPolicyTableAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayPolicyTableAssociationsPaginator returns a new
// GetTransitGatewayPolicyTableAssociationsPaginator
func NewGetTransitGatewayPolicyTableAssociationsPaginator(client GetTransitGatewayPolicyTableAssociationsAPIClient, params *GetTransitGatewayPolicyTableAssociationsInput, optFns ...func(*GetTransitGatewayPolicyTableAssociationsPaginatorOptions)) *GetTransitGatewayPolicyTableAssociationsPaginator {
	if params == nil {
		params = &GetTransitGatewayPolicyTableAssociationsInput{}
	}

	options := GetTransitGatewayPolicyTableAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayPolicyTableAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayPolicyTableAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTransitGatewayPolicyTableAssociations page.
func (p *GetTransitGatewayPolicyTableAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayPolicyTableAssociationsOutput, error) {
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

	result, err := p.client.GetTransitGatewayPolicyTableAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayPolicyTableAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayPolicyTableAssociations",
	}
}
