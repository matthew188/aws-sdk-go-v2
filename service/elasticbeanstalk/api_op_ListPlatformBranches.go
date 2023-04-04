// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the platform branches available for your account in an AWS Region.
// Provides summary information about each platform branch. For definitions of
// platform branch and other platform-related terms, see AWS Elastic Beanstalk
// Platforms Glossary
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
func (c *Client) ListPlatformBranches(ctx context.Context, params *ListPlatformBranchesInput, optFns ...func(*Options)) (*ListPlatformBranchesOutput, error) {
	if params == nil {
		params = &ListPlatformBranchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlatformBranches", params, optFns, c.addOperationListPlatformBranchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlatformBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlatformBranchesInput struct {

	// Criteria for restricting the resulting list of platform branches. The filter is
	// evaluated as a logical conjunction (AND) of the separate SearchFilter terms. The
	// following list shows valid attribute values for each of the SearchFilter terms.
	// Most operators take a single value. The in and not_in operators can take
	// multiple values.
	//
	// * Attribute = BranchName:
	//
	// * Operator: = | != | begins_with |
	// ends_with | contains | in | not_in
	//
	// * Attribute = LifecycleState:
	//
	// * Operator: =
	// | != | in | not_in
	//
	// * Values: beta | supported | deprecated | retired
	//
	// *
	// Attribute = PlatformName:
	//
	// * Operator: = | != | begins_with | ends_with |
	// contains | in | not_in
	//
	// * Attribute = TierType:
	//
	// * Operator: = | !=
	//
	// * Values:
	// WebServer/Standard | Worker/SQS/HTTP
	//
	// Array size: limited to 10 SearchFilter
	// objects. Within each SearchFilter item, the Values array is limited to 10 items.
	Filters []types.SearchFilter

	// The maximum number of platform branch values returned in one call.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlatformBranchesOutput struct {

	// In a paginated request, if this value isn't null, it's the token that you can
	// pass in a subsequent request to get the next response page.
	NextToken *string

	// Summary information about the platform branches.
	PlatformBranchSummaryList []types.PlatformBranchSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlatformBranchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformBranches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformBranches{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformBranches(options.Region), middleware.Before); err != nil {
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

// ListPlatformBranchesAPIClient is a client that implements the
// ListPlatformBranches operation.
type ListPlatformBranchesAPIClient interface {
	ListPlatformBranches(context.Context, *ListPlatformBranchesInput, ...func(*Options)) (*ListPlatformBranchesOutput, error)
}

var _ ListPlatformBranchesAPIClient = (*Client)(nil)

// ListPlatformBranchesPaginatorOptions is the paginator options for
// ListPlatformBranches
type ListPlatformBranchesPaginatorOptions struct {
	// The maximum number of platform branch values returned in one call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlatformBranchesPaginator is a paginator for ListPlatformBranches
type ListPlatformBranchesPaginator struct {
	options   ListPlatformBranchesPaginatorOptions
	client    ListPlatformBranchesAPIClient
	params    *ListPlatformBranchesInput
	nextToken *string
	firstPage bool
}

// NewListPlatformBranchesPaginator returns a new ListPlatformBranchesPaginator
func NewListPlatformBranchesPaginator(client ListPlatformBranchesAPIClient, params *ListPlatformBranchesInput, optFns ...func(*ListPlatformBranchesPaginatorOptions)) *ListPlatformBranchesPaginator {
	if params == nil {
		params = &ListPlatformBranchesInput{}
	}

	options := ListPlatformBranchesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlatformBranchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlatformBranchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlatformBranches page.
func (p *ListPlatformBranchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlatformBranchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.ListPlatformBranches(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlatformBranches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ListPlatformBranches",
	}
}
