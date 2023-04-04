// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For the specified directory, lists all the certificates registered for a secure
// LDAP or client certificate authentication.
func (c *Client) ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificates", params, optFns, c.addOperationListCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificatesInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	// The number of items that should show up on one page
	Limit *int32

	// A token for requesting another page of certificates if the NextToken response
	// element indicates that more certificates are available. Use the value of the
	// returned NextToken element in your request until the token comes back as null.
	// Pass null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCertificatesOutput struct {

	// A list of certificates with basic details including certificate ID, certificate
	// common name, certificate state.
	CertificatesInfo []types.CertificateInfo

	// Indicates whether another page of certificates is available when the number of
	// available certificates exceeds the page limit.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificates{}, middleware.After)
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
	if err = addOpListCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificates(options.Region), middleware.Before); err != nil {
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

// ListCertificatesAPIClient is a client that implements the ListCertificates
// operation.
type ListCertificatesAPIClient interface {
	ListCertificates(context.Context, *ListCertificatesInput, ...func(*Options)) (*ListCertificatesOutput, error)
}

var _ ListCertificatesAPIClient = (*Client)(nil)

// ListCertificatesPaginatorOptions is the paginator options for ListCertificates
type ListCertificatesPaginatorOptions struct {
	// The number of items that should show up on one page
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCertificatesPaginator is a paginator for ListCertificates
type ListCertificatesPaginator struct {
	options   ListCertificatesPaginatorOptions
	client    ListCertificatesAPIClient
	params    *ListCertificatesInput
	nextToken *string
	firstPage bool
}

// NewListCertificatesPaginator returns a new ListCertificatesPaginator
func NewListCertificatesPaginator(client ListCertificatesAPIClient, params *ListCertificatesInput, optFns ...func(*ListCertificatesPaginatorOptions)) *ListCertificatesPaginator {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	options := ListCertificatesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCertificates page.
func (p *ListCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListCertificates",
	}
}
