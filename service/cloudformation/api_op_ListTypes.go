// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns summary information about extension that have been registered with
// CloudFormation.
func (c *Client) ListTypes(ctx context.Context, params *ListTypesInput, optFns ...func(*Options)) (*ListTypesOutput, error) {
	if params == nil {
		params = &ListTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypes", params, optFns, c.addOperationListTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypesInput struct {

	// The deprecation status of the extension that you want to get summary information
	// about. Valid values include:
	//
	// * LIVE: The extension is registered for use in
	// CloudFormation operations.
	//
	// * DEPRECATED: The extension has been deregistered
	// and can no longer be used in CloudFormation operations.
	DeprecatedStatus types.DeprecatedStatus

	// Filter criteria to use in determining which extensions to return. Filters must
	// be compatible with Visibility to return valid results. For example, specifying
	// AWS_TYPES for Category and PRIVATE for Visibility returns an empty list of
	// types, but specifying PUBLIC for Visibility returns the desired list.
	Filters *types.TypeFilters

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous paginated request didn't return all the remaining results, the
	// response object's NextToken parameter value is set to a token. To retrieve the
	// next set of results, call this action again and assign that token to the request
	// object's NextToken parameter. If there are no remaining results, the previous
	// response object's NextToken parameter is set to null.
	NextToken *string

	// For resource types, the provisioning behavior of the resource type.
	// CloudFormation determines the provisioning type during registration, based on
	// the types of handlers in the schema handler package submitted. Valid values
	// include:
	//
	// * FULLY_MUTABLE: The resource type includes an update handler to
	// process updates to the type during stack update operations.
	//
	// * IMMUTABLE: The
	// resource type doesn't include an update handler, so the type can't be updated
	// and must instead be replaced during stack update operations.
	//
	// *
	// NON_PROVISIONABLE: The resource type doesn't include create, read, and delete
	// handlers, and therefore can't actually be provisioned.
	//
	// The default is
	// FULLY_MUTABLE.
	ProvisioningType types.ProvisioningType

	// The type of extension.
	Type types.RegistryType

	// The scope at which the extensions are visible and usable in CloudFormation
	// operations. Valid values include:
	//
	// * PRIVATE: Extensions that are visible and
	// usable within this account and region. This includes:
	//
	// * Private extensions you
	// have registered in this account and region.
	//
	// * Public extensions that you have
	// activated in this account and region.
	//
	// * PUBLIC: Extensions that are publicly
	// visible and available to be activated within any Amazon Web Services account.
	// This includes extensions from Amazon Web Services, in addition to third-party
	// publishers.
	//
	// The default is PRIVATE.
	Visibility types.Visibility

	noSmithyDocumentSerde
}

type ListTypesOutput struct {

	// If the request doesn't return all the remaining results, NextToken is set to a
	// token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null.
	NextToken *string

	// A list of TypeSummary structures that contain information about the specified
	// extensions.
	TypeSummaries []types.TypeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypes(options.Region), middleware.Before); err != nil {
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

// ListTypesAPIClient is a client that implements the ListTypes operation.
type ListTypesAPIClient interface {
	ListTypes(context.Context, *ListTypesInput, ...func(*Options)) (*ListTypesOutput, error)
}

var _ ListTypesAPIClient = (*Client)(nil)

// ListTypesPaginatorOptions is the paginator options for ListTypes
type ListTypesPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTypesPaginator is a paginator for ListTypes
type ListTypesPaginator struct {
	options   ListTypesPaginatorOptions
	client    ListTypesAPIClient
	params    *ListTypesInput
	nextToken *string
	firstPage bool
}

// NewListTypesPaginator returns a new ListTypesPaginator
func NewListTypesPaginator(client ListTypesAPIClient, params *ListTypesInput, optFns ...func(*ListTypesPaginatorOptions)) *ListTypesPaginator {
	if params == nil {
		params = &ListTypesInput{}
	}

	options := ListTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTypes page.
func (p *ListTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTypesOutput, error) {
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

	result, err := p.client.ListTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypes",
	}
}
