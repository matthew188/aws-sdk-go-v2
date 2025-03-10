// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of trust store certificates.
func (c *Client) ListTrustStoreCertificates(ctx context.Context, params *ListTrustStoreCertificatesInput, optFns ...func(*Options)) (*ListTrustStoreCertificatesOutput, error) {
	if params == nil {
		params = &ListTrustStoreCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrustStoreCertificates", params, optFns, c.addOperationListTrustStoreCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrustStoreCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrustStoreCertificatesInput struct {

	// The ARN of the trust store
	//
	// This member is required.
	TrustStoreArn *string

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrustStoreCertificatesOutput struct {

	// The certificate list.
	CertificateList []types.CertificateSummary

	// The pagination token used to retrieve the next page of results for this
	// operation.>
	NextToken *string

	// The ARN of the trust store.
	TrustStoreArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrustStoreCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrustStoreCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrustStoreCertificates{}, middleware.After)
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
	if err = addOpListTrustStoreCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrustStoreCertificates(options.Region), middleware.Before); err != nil {
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

// ListTrustStoreCertificatesAPIClient is a client that implements the
// ListTrustStoreCertificates operation.
type ListTrustStoreCertificatesAPIClient interface {
	ListTrustStoreCertificates(context.Context, *ListTrustStoreCertificatesInput, ...func(*Options)) (*ListTrustStoreCertificatesOutput, error)
}

var _ ListTrustStoreCertificatesAPIClient = (*Client)(nil)

// ListTrustStoreCertificatesPaginatorOptions is the paginator options for
// ListTrustStoreCertificates
type ListTrustStoreCertificatesPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrustStoreCertificatesPaginator is a paginator for
// ListTrustStoreCertificates
type ListTrustStoreCertificatesPaginator struct {
	options   ListTrustStoreCertificatesPaginatorOptions
	client    ListTrustStoreCertificatesAPIClient
	params    *ListTrustStoreCertificatesInput
	nextToken *string
	firstPage bool
}

// NewListTrustStoreCertificatesPaginator returns a new
// ListTrustStoreCertificatesPaginator
func NewListTrustStoreCertificatesPaginator(client ListTrustStoreCertificatesAPIClient, params *ListTrustStoreCertificatesInput, optFns ...func(*ListTrustStoreCertificatesPaginatorOptions)) *ListTrustStoreCertificatesPaginator {
	if params == nil {
		params = &ListTrustStoreCertificatesInput{}
	}

	options := ListTrustStoreCertificatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrustStoreCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrustStoreCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrustStoreCertificates page.
func (p *ListTrustStoreCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrustStoreCertificatesOutput, error) {
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

	result, err := p.client.ListTrustStoreCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrustStoreCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "ListTrustStoreCertificates",
	}
}
