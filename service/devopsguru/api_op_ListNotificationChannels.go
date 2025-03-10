// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of notification channels configured for DevOps Guru. Each
// notification channel is used to notify you when DevOps Guru generates an insight
// that contains information about how to improve your operations. The one
// supported notification channel is Amazon Simple Notification Service (Amazon
// SNS).
func (c *Client) ListNotificationChannels(ctx context.Context, params *ListNotificationChannelsInput, optFns ...func(*Options)) (*ListNotificationChannelsOutput, error) {
	if params == nil {
		params = &ListNotificationChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotificationChannels", params, optFns, c.addOperationListNotificationChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotificationChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotificationChannelsInput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNotificationChannelsOutput struct {

	// An array that contains the requested notification channels.
	Channels []types.NotificationChannel

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotificationChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNotificationChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNotificationChannels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotificationChannels(options.Region), middleware.Before); err != nil {
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

// ListNotificationChannelsAPIClient is a client that implements the
// ListNotificationChannels operation.
type ListNotificationChannelsAPIClient interface {
	ListNotificationChannels(context.Context, *ListNotificationChannelsInput, ...func(*Options)) (*ListNotificationChannelsOutput, error)
}

var _ ListNotificationChannelsAPIClient = (*Client)(nil)

// ListNotificationChannelsPaginatorOptions is the paginator options for
// ListNotificationChannels
type ListNotificationChannelsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotificationChannelsPaginator is a paginator for ListNotificationChannels
type ListNotificationChannelsPaginator struct {
	options   ListNotificationChannelsPaginatorOptions
	client    ListNotificationChannelsAPIClient
	params    *ListNotificationChannelsInput
	nextToken *string
	firstPage bool
}

// NewListNotificationChannelsPaginator returns a new
// ListNotificationChannelsPaginator
func NewListNotificationChannelsPaginator(client ListNotificationChannelsAPIClient, params *ListNotificationChannelsInput, optFns ...func(*ListNotificationChannelsPaginatorOptions)) *ListNotificationChannelsPaginator {
	if params == nil {
		params = &ListNotificationChannelsInput{}
	}

	options := ListNotificationChannelsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotificationChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotificationChannelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotificationChannels page.
func (p *ListNotificationChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotificationChannelsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListNotificationChannels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotificationChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "ListNotificationChannels",
	}
}
