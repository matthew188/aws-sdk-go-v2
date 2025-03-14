// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists phone numbers claimed to your Amazon Connect instance or traffic
// distribution group. If the provided TargetArn is a traffic distribution group,
// you can call this API in both Amazon Web Services Regions associated with
// traffic distribution group. For more information about phone numbers, see Set Up
// Phone Numbers for Your Contact Center
// (https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListPhoneNumbersV2(ctx context.Context, params *ListPhoneNumbersV2Input, optFns ...func(*Options)) (*ListPhoneNumbersV2Output, error) {
	if params == nil {
		params = &ListPhoneNumbersV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumbersV2", params, optFns, c.addOperationListPhoneNumbersV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumbersV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumbersV2Input struct {

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The ISO country code.
	PhoneNumberCountryCodes []types.PhoneNumberCountryCode

	// The prefix of the phone number. If provided, it must contain + as part of the
	// country code.
	PhoneNumberPrefix *string

	// The type of phone number.
	PhoneNumberTypes []types.PhoneNumberType

	// The Amazon Resource Name (ARN) for Amazon Connect instances or traffic
	// distribution groups that phone numbers are claimed to. If TargetArn input is not
	// provided, this API lists numbers claimed to all the Amazon Connect instances
	// belonging to your account in the same Amazon Web Services Region as the request.
	TargetArn *string

	noSmithyDocumentSerde
}

type ListPhoneNumbersV2Output struct {

	// Information about phone numbers that have been claimed to your Amazon Connect
	// instances or traffic distribution groups.
	ListPhoneNumbersSummaryList []types.ListPhoneNumbersSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPhoneNumbersV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumbersV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumbersV2{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbersV2(options.Region), middleware.Before); err != nil {
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

// ListPhoneNumbersV2APIClient is a client that implements the ListPhoneNumbersV2
// operation.
type ListPhoneNumbersV2APIClient interface {
	ListPhoneNumbersV2(context.Context, *ListPhoneNumbersV2Input, ...func(*Options)) (*ListPhoneNumbersV2Output, error)
}

var _ ListPhoneNumbersV2APIClient = (*Client)(nil)

// ListPhoneNumbersV2PaginatorOptions is the paginator options for
// ListPhoneNumbersV2
type ListPhoneNumbersV2PaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPhoneNumbersV2Paginator is a paginator for ListPhoneNumbersV2
type ListPhoneNumbersV2Paginator struct {
	options   ListPhoneNumbersV2PaginatorOptions
	client    ListPhoneNumbersV2APIClient
	params    *ListPhoneNumbersV2Input
	nextToken *string
	firstPage bool
}

// NewListPhoneNumbersV2Paginator returns a new ListPhoneNumbersV2Paginator
func NewListPhoneNumbersV2Paginator(client ListPhoneNumbersV2APIClient, params *ListPhoneNumbersV2Input, optFns ...func(*ListPhoneNumbersV2PaginatorOptions)) *ListPhoneNumbersV2Paginator {
	if params == nil {
		params = &ListPhoneNumbersV2Input{}
	}

	options := ListPhoneNumbersV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPhoneNumbersV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPhoneNumbersV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPhoneNumbersV2 page.
func (p *ListPhoneNumbersV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPhoneNumbersV2Output, error) {
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

	result, err := p.client.ListPhoneNumbersV2(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPhoneNumbersV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListPhoneNumbersV2",
	}
}
