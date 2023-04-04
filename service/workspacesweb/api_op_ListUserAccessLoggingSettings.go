// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of user access logging settings.
func (c *Client) ListUserAccessLoggingSettings(ctx context.Context, params *ListUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*ListUserAccessLoggingSettingsOutput, error) {
	if params == nil {
		params = &ListUserAccessLoggingSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserAccessLoggingSettings", params, optFns, c.addOperationListUserAccessLoggingSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserAccessLoggingSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserAccessLoggingSettingsInput struct {

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUserAccessLoggingSettingsOutput struct {

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// The user access logging settings.
	UserAccessLoggingSettings []types.UserAccessLoggingSettingsSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserAccessLoggingSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUserAccessLoggingSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserAccessLoggingSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserAccessLoggingSettings(options.Region), middleware.Before); err != nil {
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

// ListUserAccessLoggingSettingsAPIClient is a client that implements the
// ListUserAccessLoggingSettings operation.
type ListUserAccessLoggingSettingsAPIClient interface {
	ListUserAccessLoggingSettings(context.Context, *ListUserAccessLoggingSettingsInput, ...func(*Options)) (*ListUserAccessLoggingSettingsOutput, error)
}

var _ ListUserAccessLoggingSettingsAPIClient = (*Client)(nil)

// ListUserAccessLoggingSettingsPaginatorOptions is the paginator options for
// ListUserAccessLoggingSettings
type ListUserAccessLoggingSettingsPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserAccessLoggingSettingsPaginator is a paginator for
// ListUserAccessLoggingSettings
type ListUserAccessLoggingSettingsPaginator struct {
	options   ListUserAccessLoggingSettingsPaginatorOptions
	client    ListUserAccessLoggingSettingsAPIClient
	params    *ListUserAccessLoggingSettingsInput
	nextToken *string
	firstPage bool
}

// NewListUserAccessLoggingSettingsPaginator returns a new
// ListUserAccessLoggingSettingsPaginator
func NewListUserAccessLoggingSettingsPaginator(client ListUserAccessLoggingSettingsAPIClient, params *ListUserAccessLoggingSettingsInput, optFns ...func(*ListUserAccessLoggingSettingsPaginatorOptions)) *ListUserAccessLoggingSettingsPaginator {
	if params == nil {
		params = &ListUserAccessLoggingSettingsInput{}
	}

	options := ListUserAccessLoggingSettingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserAccessLoggingSettingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserAccessLoggingSettingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserAccessLoggingSettings page.
func (p *ListUserAccessLoggingSettingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserAccessLoggingSettingsOutput, error) {
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

	result, err := p.client.ListUserAccessLoggingSettings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserAccessLoggingSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "ListUserAccessLoggingSettings",
	}
}
