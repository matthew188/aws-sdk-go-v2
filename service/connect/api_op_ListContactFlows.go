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

// Provides information about the flows for the specified Amazon Connect instance.
// You can also create and update flows using the Amazon Connect Flow language
// (https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html).
// For more information about flows, see Flows
// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-contact-flows.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListContactFlows(ctx context.Context, params *ListContactFlowsInput, optFns ...func(*Options)) (*ListContactFlowsOutput, error) {
	if params == nil {
		params = &ListContactFlowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContactFlows", params, optFns, c.addOperationListContactFlowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactFlowsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The type of flow.
	ContactFlowTypes []types.ContactFlowType

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListContactFlowsOutput struct {

	// Information about the flows.
	ContactFlowSummaryList []types.ContactFlowSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContactFlowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContactFlows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactFlows{}, middleware.After)
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
	if err = addOpListContactFlowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContactFlows(options.Region), middleware.Before); err != nil {
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

// ListContactFlowsAPIClient is a client that implements the ListContactFlows
// operation.
type ListContactFlowsAPIClient interface {
	ListContactFlows(context.Context, *ListContactFlowsInput, ...func(*Options)) (*ListContactFlowsOutput, error)
}

var _ ListContactFlowsAPIClient = (*Client)(nil)

// ListContactFlowsPaginatorOptions is the paginator options for ListContactFlows
type ListContactFlowsPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactFlowsPaginator is a paginator for ListContactFlows
type ListContactFlowsPaginator struct {
	options   ListContactFlowsPaginatorOptions
	client    ListContactFlowsAPIClient
	params    *ListContactFlowsInput
	nextToken *string
	firstPage bool
}

// NewListContactFlowsPaginator returns a new ListContactFlowsPaginator
func NewListContactFlowsPaginator(client ListContactFlowsAPIClient, params *ListContactFlowsInput, optFns ...func(*ListContactFlowsPaginatorOptions)) *ListContactFlowsPaginator {
	if params == nil {
		params = &ListContactFlowsInput{}
	}

	options := ListContactFlowsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactFlowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactFlowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContactFlows page.
func (p *ListContactFlowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactFlowsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListContactFlows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContactFlows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListContactFlows",
	}
}
