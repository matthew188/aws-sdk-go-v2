// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all repositories associated with the specified approval rule template.
func (c *Client) ListRepositoriesForApprovalRuleTemplate(ctx context.Context, params *ListRepositoriesForApprovalRuleTemplateInput, optFns ...func(*Options)) (*ListRepositoriesForApprovalRuleTemplateOutput, error) {
	if params == nil {
		params = &ListRepositoriesForApprovalRuleTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositoriesForApprovalRuleTemplate", params, optFns, c.addOperationListRepositoriesForApprovalRuleTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoriesForApprovalRuleTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoriesForApprovalRuleTemplateInput struct {

	// The name of the approval rule template for which you want to list repositories
	// that are associated with that template.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRepositoriesForApprovalRuleTemplateOutput struct {

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// A list of repository names that are associated with the specified approval rule
	// template.
	RepositoryNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositoriesForApprovalRuleTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRepositoriesForApprovalRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRepositoriesForApprovalRuleTemplate{}, middleware.After)
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
	if err = addOpListRepositoriesForApprovalRuleTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositoriesForApprovalRuleTemplate(options.Region), middleware.Before); err != nil {
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

// ListRepositoriesForApprovalRuleTemplateAPIClient is a client that implements the
// ListRepositoriesForApprovalRuleTemplate operation.
type ListRepositoriesForApprovalRuleTemplateAPIClient interface {
	ListRepositoriesForApprovalRuleTemplate(context.Context, *ListRepositoriesForApprovalRuleTemplateInput, ...func(*Options)) (*ListRepositoriesForApprovalRuleTemplateOutput, error)
}

var _ ListRepositoriesForApprovalRuleTemplateAPIClient = (*Client)(nil)

// ListRepositoriesForApprovalRuleTemplatePaginatorOptions is the paginator options
// for ListRepositoriesForApprovalRuleTemplate
type ListRepositoriesForApprovalRuleTemplatePaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositoriesForApprovalRuleTemplatePaginator is a paginator for
// ListRepositoriesForApprovalRuleTemplate
type ListRepositoriesForApprovalRuleTemplatePaginator struct {
	options   ListRepositoriesForApprovalRuleTemplatePaginatorOptions
	client    ListRepositoriesForApprovalRuleTemplateAPIClient
	params    *ListRepositoriesForApprovalRuleTemplateInput
	nextToken *string
	firstPage bool
}

// NewListRepositoriesForApprovalRuleTemplatePaginator returns a new
// ListRepositoriesForApprovalRuleTemplatePaginator
func NewListRepositoriesForApprovalRuleTemplatePaginator(client ListRepositoriesForApprovalRuleTemplateAPIClient, params *ListRepositoriesForApprovalRuleTemplateInput, optFns ...func(*ListRepositoriesForApprovalRuleTemplatePaginatorOptions)) *ListRepositoriesForApprovalRuleTemplatePaginator {
	if params == nil {
		params = &ListRepositoriesForApprovalRuleTemplateInput{}
	}

	options := ListRepositoriesForApprovalRuleTemplatePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRepositoriesForApprovalRuleTemplatePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositoriesForApprovalRuleTemplatePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRepositoriesForApprovalRuleTemplate page.
func (p *ListRepositoriesForApprovalRuleTemplatePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositoriesForApprovalRuleTemplateOutput, error) {
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

	result, err := p.client.ListRepositoriesForApprovalRuleTemplate(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRepositoriesForApprovalRuleTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListRepositoriesForApprovalRuleTemplate",
	}
}
