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

// Searches for vocabularies within a specific Amazon Connect instance using State,
// NameStartsWith, and LanguageCode.
func (c *Client) SearchVocabularies(ctx context.Context, params *SearchVocabulariesInput, optFns ...func(*Options)) (*SearchVocabulariesOutput, error) {
	if params == nil {
		params = &SearchVocabulariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchVocabularies", params, optFns, c.addOperationSearchVocabulariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchVocabulariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchVocabulariesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The language code of the vocabulary entries. For a list of languages and their
	// corresponding language codes, see What is Amazon Transcribe?
	// (https://docs.aws.amazon.com/transcribe/latest/dg/transcribe-whatis.html)
	LanguageCode types.VocabularyLanguageCode

	// The maximum number of results to return per page.
	MaxResults int32

	// The starting pattern of the name of the vocabulary.
	NameStartsWith *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The current state of the custom vocabulary.
	State types.VocabularyState

	noSmithyDocumentSerde
}

type SearchVocabulariesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The list of the available custom vocabularies.
	VocabularySummaryList []types.VocabularySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchVocabulariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchVocabularies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchVocabularies{}, middleware.After)
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
	if err = addOpSearchVocabulariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchVocabularies(options.Region), middleware.Before); err != nil {
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

// SearchVocabulariesAPIClient is a client that implements the SearchVocabularies
// operation.
type SearchVocabulariesAPIClient interface {
	SearchVocabularies(context.Context, *SearchVocabulariesInput, ...func(*Options)) (*SearchVocabulariesOutput, error)
}

var _ SearchVocabulariesAPIClient = (*Client)(nil)

// SearchVocabulariesPaginatorOptions is the paginator options for
// SearchVocabularies
type SearchVocabulariesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchVocabulariesPaginator is a paginator for SearchVocabularies
type SearchVocabulariesPaginator struct {
	options   SearchVocabulariesPaginatorOptions
	client    SearchVocabulariesAPIClient
	params    *SearchVocabulariesInput
	nextToken *string
	firstPage bool
}

// NewSearchVocabulariesPaginator returns a new SearchVocabulariesPaginator
func NewSearchVocabulariesPaginator(client SearchVocabulariesAPIClient, params *SearchVocabulariesInput, optFns ...func(*SearchVocabulariesPaginatorOptions)) *SearchVocabulariesPaginator {
	if params == nil {
		params = &SearchVocabulariesInput{}
	}

	options := SearchVocabulariesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchVocabulariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchVocabulariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchVocabularies page.
func (p *SearchVocabulariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchVocabulariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.SearchVocabularies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchVocabularies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "SearchVocabularies",
	}
}
