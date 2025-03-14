// Code generated by smithy-go-codegen DO NOT EDIT.

package gamesparks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/gamesparks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a paginated list of stage deployment summaries from the game.
func (c *Client) ListStageDeployments(ctx context.Context, params *ListStageDeploymentsInput, optFns ...func(*Options)) (*ListStageDeploymentsOutput, error) {
	if params == nil {
		params = &ListStageDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStageDeployments", params, optFns, c.addOperationListStageDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStageDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStageDeploymentsInput struct {

	// The name of the game.
	//
	// This member is required.
	GameName *string

	// The name of the stage.
	//
	// This member is required.
	StageName *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStageDeploymentsOutput struct {

	// The token that indicates the start of the next sequential page of results. Use
	// this value when making the next call to this operation to continue where the
	// last one finished.
	NextToken *string

	// A list of stage deployment summaries. You can use the deployment IDs in the
	// UpdateStageDeployment and GetStageDeployment actions.
	StageDeployments []types.StageDeploymentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStageDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStageDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStageDeployments{}, middleware.After)
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
	if err = addOpListStageDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStageDeployments(options.Region), middleware.Before); err != nil {
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

// ListStageDeploymentsAPIClient is a client that implements the
// ListStageDeployments operation.
type ListStageDeploymentsAPIClient interface {
	ListStageDeployments(context.Context, *ListStageDeploymentsInput, ...func(*Options)) (*ListStageDeploymentsOutput, error)
}

var _ ListStageDeploymentsAPIClient = (*Client)(nil)

// ListStageDeploymentsPaginatorOptions is the paginator options for
// ListStageDeployments
type ListStageDeploymentsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStageDeploymentsPaginator is a paginator for ListStageDeployments
type ListStageDeploymentsPaginator struct {
	options   ListStageDeploymentsPaginatorOptions
	client    ListStageDeploymentsAPIClient
	params    *ListStageDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewListStageDeploymentsPaginator returns a new ListStageDeploymentsPaginator
func NewListStageDeploymentsPaginator(client ListStageDeploymentsAPIClient, params *ListStageDeploymentsInput, optFns ...func(*ListStageDeploymentsPaginatorOptions)) *ListStageDeploymentsPaginator {
	if params == nil {
		params = &ListStageDeploymentsInput{}
	}

	options := ListStageDeploymentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStageDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStageDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStageDeployments page.
func (p *ListStageDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStageDeploymentsOutput, error) {
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

	result, err := p.client.ListStageDeployments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStageDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamesparks",
		OperationName: "ListStageDeployments",
	}
}
