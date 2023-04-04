// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of migrations between Amazon Lex V1 and Amazon Lex V2.
func (c *Client) GetMigrations(ctx context.Context, params *GetMigrationsInput, optFns ...func(*Options)) (*GetMigrationsOutput, error) {
	if params == nil {
		params = &GetMigrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMigrations", params, optFns, c.addOperationGetMigrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMigrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMigrationsInput struct {

	// The maximum number of migrations to return in the response. The default is 10.
	MaxResults *int32

	// Filters the list to contain only migrations in the specified state.
	MigrationStatusEquals types.MigrationStatus

	// A pagination token that fetches the next page of migrations. If the response to
	// this operation is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of migrations, specify the pagination token in
	// the request.
	NextToken *string

	// The field to sort the list of migrations by. You can sort by the Amazon Lex V1
	// bot name or the date and time that the migration was started.
	SortByAttribute types.MigrationSortAttribute

	// The order so sort the list.
	SortByOrder types.SortOrder

	// Filters the list to contain only bots whose name contains the specified string.
	// The string is matched anywhere in bot name.
	V1BotNameContains *string

	noSmithyDocumentSerde
}

type GetMigrationsOutput struct {

	// An array of summaries for migrations from Amazon Lex V1 to Amazon Lex V2. To see
	// details of the migration, use the migrationId from the summary in a call to the
	// operation.
	MigrationSummaries []types.MigrationSummary

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of migrations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMigrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMigrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMigrations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMigrations(options.Region), middleware.Before); err != nil {
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

// GetMigrationsAPIClient is a client that implements the GetMigrations operation.
type GetMigrationsAPIClient interface {
	GetMigrations(context.Context, *GetMigrationsInput, ...func(*Options)) (*GetMigrationsOutput, error)
}

var _ GetMigrationsAPIClient = (*Client)(nil)

// GetMigrationsPaginatorOptions is the paginator options for GetMigrations
type GetMigrationsPaginatorOptions struct {
	// The maximum number of migrations to return in the response. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMigrationsPaginator is a paginator for GetMigrations
type GetMigrationsPaginator struct {
	options   GetMigrationsPaginatorOptions
	client    GetMigrationsAPIClient
	params    *GetMigrationsInput
	nextToken *string
	firstPage bool
}

// NewGetMigrationsPaginator returns a new GetMigrationsPaginator
func NewGetMigrationsPaginator(client GetMigrationsAPIClient, params *GetMigrationsInput, optFns ...func(*GetMigrationsPaginatorOptions)) *GetMigrationsPaginator {
	if params == nil {
		params = &GetMigrationsInput{}
	}

	options := GetMigrationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMigrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMigrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMigrations page.
func (p *GetMigrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMigrationsOutput, error) {
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

	result, err := p.client.GetMigrations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetMigrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetMigrations",
	}
}
