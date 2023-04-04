// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List position configurations for a given resource, such as positioning solvers.
// This action is no longer supported. Calls to retrieve position information
// should use the GetResourcePosition
// (https://docs.aws.amazon.com/iot-wireless/2020-11-22/apireference/API_GetResourcePosition.html)
// API operation instead.
//
// Deprecated: This operation is no longer supported.
func (c *Client) ListPositionConfigurations(ctx context.Context, params *ListPositionConfigurationsInput, optFns ...func(*Options)) (*ListPositionConfigurationsOutput, error) {
	if params == nil {
		params = &ListPositionConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPositionConfigurations", params, optFns, c.addOperationListPositionConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPositionConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPositionConfigurationsInput struct {

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// Resource type for which position configurations are listed.
	ResourceType types.PositionResourceType

	noSmithyDocumentSerde
}

type ListPositionConfigurationsOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// A list of position configurations.
	PositionConfigurationList []types.PositionConfigurationItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPositionConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPositionConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPositionConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPositionConfigurations(options.Region), middleware.Before); err != nil {
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

// ListPositionConfigurationsAPIClient is a client that implements the
// ListPositionConfigurations operation.
type ListPositionConfigurationsAPIClient interface {
	ListPositionConfigurations(context.Context, *ListPositionConfigurationsInput, ...func(*Options)) (*ListPositionConfigurationsOutput, error)
}

var _ ListPositionConfigurationsAPIClient = (*Client)(nil)

// ListPositionConfigurationsPaginatorOptions is the paginator options for
// ListPositionConfigurations
type ListPositionConfigurationsPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPositionConfigurationsPaginator is a paginator for
// ListPositionConfigurations
type ListPositionConfigurationsPaginator struct {
	options   ListPositionConfigurationsPaginatorOptions
	client    ListPositionConfigurationsAPIClient
	params    *ListPositionConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListPositionConfigurationsPaginator returns a new
// ListPositionConfigurationsPaginator
func NewListPositionConfigurationsPaginator(client ListPositionConfigurationsAPIClient, params *ListPositionConfigurationsInput, optFns ...func(*ListPositionConfigurationsPaginatorOptions)) *ListPositionConfigurationsPaginator {
	if params == nil {
		params = &ListPositionConfigurationsInput{}
	}

	options := ListPositionConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPositionConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPositionConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPositionConfigurations page.
func (p *ListPositionConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPositionConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPositionConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPositionConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListPositionConfigurations",
	}
}
