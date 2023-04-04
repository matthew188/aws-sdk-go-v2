// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the agreements for the server that's identified by the
// ServerId that you supply. If you want to limit the results to a certain number,
// supply a value for the MaxResults parameter. If you ran the command previously
// and received a value for NextToken, you can supply that value to continue
// listing agreements from where you left off.
func (c *Client) ListAgreements(ctx context.Context, params *ListAgreementsInput, optFns ...func(*Options)) (*ListAgreementsOutput, error) {
	if params == nil {
		params = &ListAgreementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgreements", params, optFns, c.addOperationListAgreementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgreementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAgreementsInput struct {

	// The identifier of the server for which you want a list of agreements.
	//
	// This member is required.
	ServerId *string

	// The maximum number of agreements to return.
	MaxResults *int32

	// When you can get additional results from the ListAgreements call, a NextToken
	// parameter is returned in the output. You can then pass in a subsequent command
	// to the NextToken parameter to continue listing additional agreements.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAgreementsOutput struct {

	// Returns an array, where each item contains the details of an agreement.
	//
	// This member is required.
	Agreements []types.ListedAgreement

	// Returns a token that you can use to call ListAgreements again and receive
	// additional results, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgreementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAgreements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAgreements{}, middleware.After)
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
	if err = addOpListAgreementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgreements(options.Region), middleware.Before); err != nil {
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

// ListAgreementsAPIClient is a client that implements the ListAgreements
// operation.
type ListAgreementsAPIClient interface {
	ListAgreements(context.Context, *ListAgreementsInput, ...func(*Options)) (*ListAgreementsOutput, error)
}

var _ ListAgreementsAPIClient = (*Client)(nil)

// ListAgreementsPaginatorOptions is the paginator options for ListAgreements
type ListAgreementsPaginatorOptions struct {
	// The maximum number of agreements to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgreementsPaginator is a paginator for ListAgreements
type ListAgreementsPaginator struct {
	options   ListAgreementsPaginatorOptions
	client    ListAgreementsAPIClient
	params    *ListAgreementsInput
	nextToken *string
	firstPage bool
}

// NewListAgreementsPaginator returns a new ListAgreementsPaginator
func NewListAgreementsPaginator(client ListAgreementsAPIClient, params *ListAgreementsInput, optFns ...func(*ListAgreementsPaginatorOptions)) *ListAgreementsPaginator {
	if params == nil {
		params = &ListAgreementsInput{}
	}

	options := ListAgreementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgreementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgreementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgreements page.
func (p *ListAgreementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgreementsOutput, error) {
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

	result, err := p.client.ListAgreements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAgreements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListAgreements",
	}
}
