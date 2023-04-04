// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List targets for the specified policy. Requires permission to access the
// ListTargetsForPolicy
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListTargetsForPolicy(ctx context.Context, params *ListTargetsForPolicyInput, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
	if params == nil {
		params = &ListTargetsForPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargetsForPolicy", params, optFns, c.addOperationListTargetsForPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsForPolicyInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// A marker used to get the next set of results.
	Marker *string

	// The maximum number of results to return at one time.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListTargetsForPolicyOutput struct {

	// A marker used to get the next set of results.
	NextMarker *string

	// The policy targets.
	Targets []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTargetsForPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargetsForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargetsForPolicy{}, middleware.After)
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
	if err = addOpListTargetsForPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetsForPolicy(options.Region), middleware.Before); err != nil {
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

// ListTargetsForPolicyAPIClient is a client that implements the
// ListTargetsForPolicy operation.
type ListTargetsForPolicyAPIClient interface {
	ListTargetsForPolicy(context.Context, *ListTargetsForPolicyInput, ...func(*Options)) (*ListTargetsForPolicyOutput, error)
}

var _ ListTargetsForPolicyAPIClient = (*Client)(nil)

// ListTargetsForPolicyPaginatorOptions is the paginator options for
// ListTargetsForPolicy
type ListTargetsForPolicyPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTargetsForPolicyPaginator is a paginator for ListTargetsForPolicy
type ListTargetsForPolicyPaginator struct {
	options   ListTargetsForPolicyPaginatorOptions
	client    ListTargetsForPolicyAPIClient
	params    *ListTargetsForPolicyInput
	nextToken *string
	firstPage bool
}

// NewListTargetsForPolicyPaginator returns a new ListTargetsForPolicyPaginator
func NewListTargetsForPolicyPaginator(client ListTargetsForPolicyAPIClient, params *ListTargetsForPolicyInput, optFns ...func(*ListTargetsForPolicyPaginatorOptions)) *ListTargetsForPolicyPaginator {
	if params == nil {
		params = &ListTargetsForPolicyInput{}
	}

	options := ListTargetsForPolicyPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTargetsForPolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTargetsForPolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTargetsForPolicy page.
func (p *ListTargetsForPolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListTargetsForPolicy(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTargetsForPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListTargetsForPolicy",
	}
}
