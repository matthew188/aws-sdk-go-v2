// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes all of the budget actions for an account.
func (c *Client) DescribeBudgetActionsForAccount(ctx context.Context, params *DescribeBudgetActionsForAccountInput, optFns ...func(*Options)) (*DescribeBudgetActionsForAccountOutput, error) {
	if params == nil {
		params = &DescribeBudgetActionsForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetActionsForAccount", params, optFns, c.addOperationDescribeBudgetActionsForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetActionsForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetActionsForAccountInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeBudgetActionsForAccountOutput struct {

	// A list of the budget action resources information.
	//
	// This member is required.
	Actions []types.Action

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetActionsForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetActionsForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetActionsForAccount{}, middleware.After)
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
	if err = addOpDescribeBudgetActionsForAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetActionsForAccount(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetActionsForAccountAPIClient is a client that implements the
// DescribeBudgetActionsForAccount operation.
type DescribeBudgetActionsForAccountAPIClient interface {
	DescribeBudgetActionsForAccount(context.Context, *DescribeBudgetActionsForAccountInput, ...func(*Options)) (*DescribeBudgetActionsForAccountOutput, error)
}

var _ DescribeBudgetActionsForAccountAPIClient = (*Client)(nil)

// DescribeBudgetActionsForAccountPaginatorOptions is the paginator options for
// DescribeBudgetActionsForAccount
type DescribeBudgetActionsForAccountPaginatorOptions struct {
	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetActionsForAccountPaginator is a paginator for
// DescribeBudgetActionsForAccount
type DescribeBudgetActionsForAccountPaginator struct {
	options   DescribeBudgetActionsForAccountPaginatorOptions
	client    DescribeBudgetActionsForAccountAPIClient
	params    *DescribeBudgetActionsForAccountInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetActionsForAccountPaginator returns a new
// DescribeBudgetActionsForAccountPaginator
func NewDescribeBudgetActionsForAccountPaginator(client DescribeBudgetActionsForAccountAPIClient, params *DescribeBudgetActionsForAccountInput, optFns ...func(*DescribeBudgetActionsForAccountPaginatorOptions)) *DescribeBudgetActionsForAccountPaginator {
	if params == nil {
		params = &DescribeBudgetActionsForAccountInput{}
	}

	options := DescribeBudgetActionsForAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetActionsForAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetActionsForAccountPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgetActionsForAccount page.
func (p *DescribeBudgetActionsForAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetActionsForAccountOutput, error) {
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

	result, err := p.client.DescribeBudgetActionsForAccount(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgetActionsForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetActionsForAccount",
	}
}
