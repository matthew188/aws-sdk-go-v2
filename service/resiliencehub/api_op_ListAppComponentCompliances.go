// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the compliances for an Resilience Hub Application Component.
func (c *Client) ListAppComponentCompliances(ctx context.Context, params *ListAppComponentCompliancesInput, optFns ...func(*Options)) (*ListAppComponentCompliancesOutput, error) {
	if params == nil {
		params = &ListAppComponentCompliancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppComponentCompliances", params, optFns, c.addOperationListAppComponentCompliancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppComponentCompliancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppComponentCompliancesInput struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app-assessment/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AssessmentArn *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppComponentCompliancesOutput struct {

	// The compliances for an Resilience Hub Application Component, returned as an
	// object. This object contains the names of the Application Components,
	// compliances, costs, resiliency scores, outage scores, and more.
	//
	// This member is required.
	ComponentCompliances []types.AppComponentCompliance

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppComponentCompliancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppComponentCompliances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppComponentCompliances{}, middleware.After)
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
	if err = addOpListAppComponentCompliancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppComponentCompliances(options.Region), middleware.Before); err != nil {
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

// ListAppComponentCompliancesAPIClient is a client that implements the
// ListAppComponentCompliances operation.
type ListAppComponentCompliancesAPIClient interface {
	ListAppComponentCompliances(context.Context, *ListAppComponentCompliancesInput, ...func(*Options)) (*ListAppComponentCompliancesOutput, error)
}

var _ ListAppComponentCompliancesAPIClient = (*Client)(nil)

// ListAppComponentCompliancesPaginatorOptions is the paginator options for
// ListAppComponentCompliances
type ListAppComponentCompliancesPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppComponentCompliancesPaginator is a paginator for
// ListAppComponentCompliances
type ListAppComponentCompliancesPaginator struct {
	options   ListAppComponentCompliancesPaginatorOptions
	client    ListAppComponentCompliancesAPIClient
	params    *ListAppComponentCompliancesInput
	nextToken *string
	firstPage bool
}

// NewListAppComponentCompliancesPaginator returns a new
// ListAppComponentCompliancesPaginator
func NewListAppComponentCompliancesPaginator(client ListAppComponentCompliancesAPIClient, params *ListAppComponentCompliancesInput, optFns ...func(*ListAppComponentCompliancesPaginatorOptions)) *ListAppComponentCompliancesPaginator {
	if params == nil {
		params = &ListAppComponentCompliancesInput{}
	}

	options := ListAppComponentCompliancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppComponentCompliancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppComponentCompliancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppComponentCompliances page.
func (p *ListAppComponentCompliancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppComponentCompliancesOutput, error) {
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

	result, err := p.client.ListAppComponentCompliances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppComponentCompliances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "ListAppComponentCompliances",
	}
}
