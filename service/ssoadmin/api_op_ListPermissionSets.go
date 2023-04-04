// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the PermissionSets in an IAM Identity Center instance.
func (c *Client) ListPermissionSets(ctx context.Context, params *ListPermissionSetsInput, optFns ...func(*Options)) (*ListPermissionSetsOutput, error) {
	if params == nil {
		params = &ListPermissionSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionSets", params, optFns, c.addOperationListPermissionSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionSetsInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPermissionSetsOutput struct {

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Defines the level of access on an AWS account.
	PermissionSets []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissionSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissionSets{}, middleware.After)
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
	if err = addOpListPermissionSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionSets(options.Region), middleware.Before); err != nil {
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

// ListPermissionSetsAPIClient is a client that implements the ListPermissionSets
// operation.
type ListPermissionSetsAPIClient interface {
	ListPermissionSets(context.Context, *ListPermissionSetsInput, ...func(*Options)) (*ListPermissionSetsOutput, error)
}

var _ ListPermissionSetsAPIClient = (*Client)(nil)

// ListPermissionSetsPaginatorOptions is the paginator options for
// ListPermissionSets
type ListPermissionSetsPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionSetsPaginator is a paginator for ListPermissionSets
type ListPermissionSetsPaginator struct {
	options   ListPermissionSetsPaginatorOptions
	client    ListPermissionSetsAPIClient
	params    *ListPermissionSetsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionSetsPaginator returns a new ListPermissionSetsPaginator
func NewListPermissionSetsPaginator(client ListPermissionSetsAPIClient, params *ListPermissionSetsInput, optFns ...func(*ListPermissionSetsPaginatorOptions)) *ListPermissionSetsPaginator {
	if params == nil {
		params = &ListPermissionSetsInput{}
	}

	options := ListPermissionSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissionSets page.
func (p *ListPermissionSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionSetsOutput, error) {
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

	result, err := p.client.ListPermissionSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPermissionSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListPermissionSets",
	}
}
