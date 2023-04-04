// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about all activities registered in the specified domain that
// match the specified name and registration status. The result includes
// information like creation date, current status of the activity, etc. The results
// may be split into multiple pages. To retrieve subsequent pages, make the call
// again using the nextPageToken returned by the initial call. Access Control You
// can use IAM policies to control this action's access to Amazon SWF resources as
// follows:
//
// * Use a Resource element with the domain name to limit the action to
// only specified domains.
//
// * Use an Action element to allow or deny permission to
// call this action.
//
// * You cannot use an IAM policy to constrain this action's
// parameters.
//
// If the caller doesn't have sufficient permissions to invoke the
// action, or the parameter values fall outside the specified constraints, the
// action fails. The associated event attribute's cause parameter is set to
// OPERATION_NOT_PERMITTED. For details and example IAM policies, see Using IAM to
// Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) ListActivityTypes(ctx context.Context, params *ListActivityTypesInput, optFns ...func(*Options)) (*ListActivityTypesOutput, error) {
	if params == nil {
		params = &ListActivityTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActivityTypes", params, optFns, c.addOperationListActivityTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActivityTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActivityTypesInput struct {

	// The name of the domain in which the activity types have been registered.
	//
	// This member is required.
	Domain *string

	// Specifies the registration status of the activity types to list.
	//
	// This member is required.
	RegistrationStatus types.RegistrationStatus

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize int32

	// If specified, only lists the activity types that have this name.
	Name *string

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime". The configured maximumPageSize determines how many results
	// can be returned in a single call.
	NextPageToken *string

	// When set to true, returns the results in reverse order. By default, the results
	// are returned in ascending alphabetical order by name of the activity types.
	ReverseOrder bool

	noSmithyDocumentSerde
}

// Contains a paginated list of activity type information structures.
type ListActivityTypesOutput struct {

	// List of activity type information.
	//
	// This member is required.
	TypeInfos []types.ActivityTypeInfo

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListActivityTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListActivityTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListActivityTypes{}, middleware.After)
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
	if err = addOpListActivityTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActivityTypes(options.Region), middleware.Before); err != nil {
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

// ListActivityTypesAPIClient is a client that implements the ListActivityTypes
// operation.
type ListActivityTypesAPIClient interface {
	ListActivityTypes(context.Context, *ListActivityTypesInput, ...func(*Options)) (*ListActivityTypesOutput, error)
}

var _ ListActivityTypesAPIClient = (*Client)(nil)

// ListActivityTypesPaginatorOptions is the paginator options for ListActivityTypes
type ListActivityTypesPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListActivityTypesPaginator is a paginator for ListActivityTypes
type ListActivityTypesPaginator struct {
	options   ListActivityTypesPaginatorOptions
	client    ListActivityTypesAPIClient
	params    *ListActivityTypesInput
	nextToken *string
	firstPage bool
}

// NewListActivityTypesPaginator returns a new ListActivityTypesPaginator
func NewListActivityTypesPaginator(client ListActivityTypesAPIClient, params *ListActivityTypesInput, optFns ...func(*ListActivityTypesPaginatorOptions)) *ListActivityTypesPaginator {
	if params == nil {
		params = &ListActivityTypesInput{}
	}

	options := ListActivityTypesPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListActivityTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextPageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListActivityTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListActivityTypes page.
func (p *ListActivityTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListActivityTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	result, err := p.client.ListActivityTypes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListActivityTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "ListActivityTypes",
	}
}
