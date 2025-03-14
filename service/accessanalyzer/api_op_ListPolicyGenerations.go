// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the policy generations requested in the last seven days.
func (c *Client) ListPolicyGenerations(ctx context.Context, params *ListPolicyGenerationsInput, optFns ...func(*Options)) (*ListPolicyGenerationsOutput, error) {
	if params == nil {
		params = &ListPolicyGenerationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyGenerations", params, optFns, c.addOperationListPolicyGenerationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyGenerationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyGenerationsInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	// The ARN of the IAM entity (user or role) for which you are generating a policy.
	// Use this with ListGeneratedPolicies to filter the results to only include
	// results for a specific principal.
	PrincipalArn *string

	noSmithyDocumentSerde
}

type ListPolicyGenerationsOutput struct {

	// A PolicyGeneration object that contains details about the generated policy.
	//
	// This member is required.
	PolicyGenerations []types.PolicyGeneration

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPolicyGenerationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyGenerations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyGenerations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyGenerations(options.Region), middleware.Before); err != nil {
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

// ListPolicyGenerationsAPIClient is a client that implements the
// ListPolicyGenerations operation.
type ListPolicyGenerationsAPIClient interface {
	ListPolicyGenerations(context.Context, *ListPolicyGenerationsInput, ...func(*Options)) (*ListPolicyGenerationsOutput, error)
}

var _ ListPolicyGenerationsAPIClient = (*Client)(nil)

// ListPolicyGenerationsPaginatorOptions is the paginator options for
// ListPolicyGenerations
type ListPolicyGenerationsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPolicyGenerationsPaginator is a paginator for ListPolicyGenerations
type ListPolicyGenerationsPaginator struct {
	options   ListPolicyGenerationsPaginatorOptions
	client    ListPolicyGenerationsAPIClient
	params    *ListPolicyGenerationsInput
	nextToken *string
	firstPage bool
}

// NewListPolicyGenerationsPaginator returns a new ListPolicyGenerationsPaginator
func NewListPolicyGenerationsPaginator(client ListPolicyGenerationsAPIClient, params *ListPolicyGenerationsInput, optFns ...func(*ListPolicyGenerationsPaginatorOptions)) *ListPolicyGenerationsPaginator {
	if params == nil {
		params = &ListPolicyGenerationsInput{}
	}

	options := ListPolicyGenerationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPolicyGenerationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPolicyGenerationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicyGenerations page.
func (p *ListPolicyGenerationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPolicyGenerationsOutput, error) {
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

	result, err := p.client.ListPolicyGenerations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPolicyGenerations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "ListPolicyGenerations",
	}
}
