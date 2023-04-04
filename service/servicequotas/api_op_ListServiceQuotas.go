// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the applied quota values for the specified AWS service. For some quotas,
// only the default values are available. If the applied quota value is not
// available for a quota, the quota is not retrieved.
func (c *Client) ListServiceQuotas(ctx context.Context, params *ListServiceQuotasInput, optFns ...func(*Options)) (*ListServiceQuotasOutput, error) {
	if params == nil {
		params = &ListServiceQuotasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceQuotas", params, optFns, c.addOperationListServiceQuotasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceQuotasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceQuotasInput struct {

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceQuotasOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the quotas.
	Quotas []types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceQuotasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceQuotas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceQuotas{}, middleware.After)
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
	if err = addOpListServiceQuotasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceQuotas(options.Region), middleware.Before); err != nil {
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

// ListServiceQuotasAPIClient is a client that implements the ListServiceQuotas
// operation.
type ListServiceQuotasAPIClient interface {
	ListServiceQuotas(context.Context, *ListServiceQuotasInput, ...func(*Options)) (*ListServiceQuotasOutput, error)
}

var _ ListServiceQuotasAPIClient = (*Client)(nil)

// ListServiceQuotasPaginatorOptions is the paginator options for ListServiceQuotas
type ListServiceQuotasPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceQuotasPaginator is a paginator for ListServiceQuotas
type ListServiceQuotasPaginator struct {
	options   ListServiceQuotasPaginatorOptions
	client    ListServiceQuotasAPIClient
	params    *ListServiceQuotasInput
	nextToken *string
	firstPage bool
}

// NewListServiceQuotasPaginator returns a new ListServiceQuotasPaginator
func NewListServiceQuotasPaginator(client ListServiceQuotasAPIClient, params *ListServiceQuotasInput, optFns ...func(*ListServiceQuotasPaginatorOptions)) *ListServiceQuotasPaginator {
	if params == nil {
		params = &ListServiceQuotasInput{}
	}

	options := ListServiceQuotasPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceQuotasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceQuotasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceQuotas page.
func (p *ListServiceQuotasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceQuotasOutput, error) {
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

	result, err := p.client.ListServiceQuotas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceQuotas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListServiceQuotas",
	}
}
