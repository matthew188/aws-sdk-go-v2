// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists one or more access control configurations for an index. This includes user
// and group access information for your documents. This is useful for user context
// filtering, where search results are filtered based on the user or their group
// access to documents.
func (c *Client) ListAccessControlConfigurations(ctx context.Context, params *ListAccessControlConfigurationsInput, optFns ...func(*Options)) (*ListAccessControlConfigurationsOutput, error) {
	if params == nil {
		params = &ListAccessControlConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessControlConfigurations", params, optFns, c.addOperationListAccessControlConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessControlConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessControlConfigurationsInput struct {

	// The identifier of the index for the access control configuration.
	//
	// This member is required.
	IndexId *string

	// The maximum number of access control configurations to return.
	MaxResults *int32

	// If the previous response was incomplete (because there's more data to retrieve),
	// Amazon Kendra returns a pagination token in the response. You can use this
	// pagination token to retrieve the next set of access control configurations.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccessControlConfigurationsOutput struct {

	// The details of your access control configurations.
	//
	// This member is required.
	AccessControlConfigurations []types.AccessControlConfigurationSummary

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in the subsequent request to retrieve the next set of access control
	// configurations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessControlConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccessControlConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccessControlConfigurations{}, middleware.After)
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
	if err = addOpListAccessControlConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessControlConfigurations(options.Region), middleware.Before); err != nil {
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

// ListAccessControlConfigurationsAPIClient is a client that implements the
// ListAccessControlConfigurations operation.
type ListAccessControlConfigurationsAPIClient interface {
	ListAccessControlConfigurations(context.Context, *ListAccessControlConfigurationsInput, ...func(*Options)) (*ListAccessControlConfigurationsOutput, error)
}

var _ ListAccessControlConfigurationsAPIClient = (*Client)(nil)

// ListAccessControlConfigurationsPaginatorOptions is the paginator options for
// ListAccessControlConfigurations
type ListAccessControlConfigurationsPaginatorOptions struct {
	// The maximum number of access control configurations to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessControlConfigurationsPaginator is a paginator for
// ListAccessControlConfigurations
type ListAccessControlConfigurationsPaginator struct {
	options   ListAccessControlConfigurationsPaginatorOptions
	client    ListAccessControlConfigurationsAPIClient
	params    *ListAccessControlConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListAccessControlConfigurationsPaginator returns a new
// ListAccessControlConfigurationsPaginator
func NewListAccessControlConfigurationsPaginator(client ListAccessControlConfigurationsAPIClient, params *ListAccessControlConfigurationsInput, optFns ...func(*ListAccessControlConfigurationsPaginatorOptions)) *ListAccessControlConfigurationsPaginator {
	if params == nil {
		params = &ListAccessControlConfigurationsInput{}
	}

	options := ListAccessControlConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessControlConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessControlConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessControlConfigurations page.
func (p *ListAccessControlConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessControlConfigurationsOutput, error) {
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

	result, err := p.client.ListAccessControlConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessControlConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListAccessControlConfigurations",
	}
}
