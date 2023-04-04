// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List provisioned resources for a component with details. For more information
// about components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html) in the
// Proton User Guide.
func (c *Client) ListComponentProvisionedResources(ctx context.Context, params *ListComponentProvisionedResourcesInput, optFns ...func(*Options)) (*ListComponentProvisionedResourcesOutput, error) {
	if params == nil {
		params = &ListComponentProvisionedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentProvisionedResources", params, optFns, c.addOperationListComponentProvisionedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentProvisionedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentProvisionedResourcesInput struct {

	// The name of the component whose provisioned resources you want.
	//
	// This member is required.
	ComponentName *string

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the list of provisioned resources that was
	// previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentProvisionedResourcesOutput struct {

	// An array of provisioned resources for a component.
	//
	// This member is required.
	ProvisionedResources []types.ProvisionedResource

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the current requested list of provisioned
	// resources.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentProvisionedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListComponentProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListComponentProvisionedResources{}, middleware.After)
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
	if err = addOpListComponentProvisionedResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentProvisionedResources(options.Region), middleware.Before); err != nil {
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

// ListComponentProvisionedResourcesAPIClient is a client that implements the
// ListComponentProvisionedResources operation.
type ListComponentProvisionedResourcesAPIClient interface {
	ListComponentProvisionedResources(context.Context, *ListComponentProvisionedResourcesInput, ...func(*Options)) (*ListComponentProvisionedResourcesOutput, error)
}

var _ ListComponentProvisionedResourcesAPIClient = (*Client)(nil)

// ListComponentProvisionedResourcesPaginatorOptions is the paginator options for
// ListComponentProvisionedResources
type ListComponentProvisionedResourcesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentProvisionedResourcesPaginator is a paginator for
// ListComponentProvisionedResources
type ListComponentProvisionedResourcesPaginator struct {
	options   ListComponentProvisionedResourcesPaginatorOptions
	client    ListComponentProvisionedResourcesAPIClient
	params    *ListComponentProvisionedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListComponentProvisionedResourcesPaginator returns a new
// ListComponentProvisionedResourcesPaginator
func NewListComponentProvisionedResourcesPaginator(client ListComponentProvisionedResourcesAPIClient, params *ListComponentProvisionedResourcesInput, optFns ...func(*ListComponentProvisionedResourcesPaginatorOptions)) *ListComponentProvisionedResourcesPaginator {
	if params == nil {
		params = &ListComponentProvisionedResourcesInput{}
	}

	options := ListComponentProvisionedResourcesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentProvisionedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentProvisionedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponentProvisionedResources page.
func (p *ListComponentProvisionedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentProvisionedResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListComponentProvisionedResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponentProvisionedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListComponentProvisionedResources",
	}
}
