// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/amplifyuibuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of components for a specified Amplify app and backend
// environment.
func (c *Client) ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) {
	if params == nil {
		params = &ListComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponents", params, optFns, c.addOperationListComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentsInput struct {

	// The unique ID for the Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The maximum number of components to retrieve.
	MaxResults int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentsOutput struct {

	// The list of components for the Amplify app.
	//
	// This member is required.
	Entities []types.ComponentSummary

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponents{}, middleware.After)
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
	if err = addOpListComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponents(options.Region), middleware.Before); err != nil {
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

// ListComponentsAPIClient is a client that implements the ListComponents
// operation.
type ListComponentsAPIClient interface {
	ListComponents(context.Context, *ListComponentsInput, ...func(*Options)) (*ListComponentsOutput, error)
}

var _ ListComponentsAPIClient = (*Client)(nil)

// ListComponentsPaginatorOptions is the paginator options for ListComponents
type ListComponentsPaginatorOptions struct {
	// The maximum number of components to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentsPaginator is a paginator for ListComponents
type ListComponentsPaginator struct {
	options   ListComponentsPaginatorOptions
	client    ListComponentsAPIClient
	params    *ListComponentsInput
	nextToken *string
	firstPage bool
}

// NewListComponentsPaginator returns a new ListComponentsPaginator
func NewListComponentsPaginator(client ListComponentsAPIClient, params *ListComponentsInput, optFns ...func(*ListComponentsPaginatorOptions)) *ListComponentsPaginator {
	if params == nil {
		params = &ListComponentsInput{}
	}

	options := ListComponentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponents page.
func (p *ListComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifyuibuilder",
		OperationName: "ListComponents",
	}
}
