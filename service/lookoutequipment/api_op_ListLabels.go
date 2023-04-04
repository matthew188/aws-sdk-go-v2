// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a list of labels.
func (c *Client) ListLabels(ctx context.Context, params *ListLabelsInput, optFns ...func(*Options)) (*ListLabelsOutput, error) {
	if params == nil {
		params = &ListLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLabels", params, optFns, c.addOperationListLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLabelsInput struct {

	// Retruns the name of the label group.
	//
	// This member is required.
	LabelGroupName *string

	// Lists the labels that pertain to a particular piece of equipment.
	Equipment *string

	// Returns labels with a particular fault code.
	FaultCode *string

	// Returns all labels with a start time earlier than the end time given.
	IntervalEndTime *time.Time

	// Returns all the labels with a end time equal to or later than the start time
	// given.
	IntervalStartTime *time.Time

	// Specifies the maximum number of labels to list.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of label
	// groups.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLabelsOutput struct {

	// A summary of the items in the label group.
	LabelSummaries []types.LabelSummary

	// An opaque pagination token indicating where to continue the listing of datasets.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListLabels{}, middleware.After)
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
	if err = addOpListLabelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLabels(options.Region), middleware.Before); err != nil {
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

// ListLabelsAPIClient is a client that implements the ListLabels operation.
type ListLabelsAPIClient interface {
	ListLabels(context.Context, *ListLabelsInput, ...func(*Options)) (*ListLabelsOutput, error)
}

var _ ListLabelsAPIClient = (*Client)(nil)

// ListLabelsPaginatorOptions is the paginator options for ListLabels
type ListLabelsPaginatorOptions struct {
	// Specifies the maximum number of labels to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLabelsPaginator is a paginator for ListLabels
type ListLabelsPaginator struct {
	options   ListLabelsPaginatorOptions
	client    ListLabelsAPIClient
	params    *ListLabelsInput
	nextToken *string
	firstPage bool
}

// NewListLabelsPaginator returns a new ListLabelsPaginator
func NewListLabelsPaginator(client ListLabelsAPIClient, params *ListLabelsInput, optFns ...func(*ListLabelsPaginatorOptions)) *ListLabelsPaginator {
	if params == nil {
		params = &ListLabelsInput{}
	}

	options := ListLabelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLabelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLabelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLabels page.
func (p *ListLabelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLabelsOutput, error) {
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

	result, err := p.client.ListLabels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListLabels",
	}
}
