// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the versions of a secret. Secrets Manager uses staging labels to indicate
// the different versions of a secret. For more information, see  Secrets Manager
// concepts: Versions
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version).
// To list the secrets in the account, use ListSecrets. Secrets Manager generates a
// CloudTrail log entry when you call this action. Do not include sensitive
// information in request parameters because it might be logged. For more
// information, see Logging Secrets Manager events with CloudTrail
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html).
// Required permissions: secretsmanager:ListSecretVersionIds. For more information,
// see  IAM policy actions for Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html).
func (c *Client) ListSecretVersionIds(ctx context.Context, params *ListSecretVersionIdsInput, optFns ...func(*Options)) (*ListSecretVersionIdsOutput, error) {
	if params == nil {
		params = &ListSecretVersionIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecretVersionIds", params, optFns, c.addOperationListSecretVersionIdsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecretVersionIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecretVersionIdsInput struct {

	// The ARN or name of the secret whose versions you want to list. For an ARN, we
	// recommend that you specify a complete ARN rather than a partial ARN. See Finding
	// a secret from a partial ARN
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen).
	//
	// This member is required.
	SecretId *string

	// Specifies whether to include versions of secrets that don't have any staging
	// labels attached to them. Versions without staging labels are considered
	// deprecated and are subject to deletion by Secrets Manager.
	IncludeDeprecated *bool

	// The number of results to include in the response. If there are more results
	// available, in the response, Secrets Manager includes NextToken. To get the next
	// results, call ListSecretVersionIds again with the value from NextToken.
	MaxResults *int32

	// A token that indicates where the output should continue from, if a previous call
	// did not show all results. To get the next results, call ListSecretVersionIds
	// again with this value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecretVersionIdsOutput struct {

	// The ARN of the secret.
	ARN *string

	// The name of the secret.
	Name *string

	// Secrets Manager includes this value if there's more output available than what
	// is included in the current response. This can occur even when the response
	// includes no values at all, such as when you ask for a filtered view of a long
	// list. To get the next results, call ListSecretVersionIds again with this value.
	NextToken *string

	// A list of the versions of the secret.
	Versions []types.SecretVersionsListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecretVersionIdsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSecretVersionIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecretVersionIds{}, middleware.After)
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
	if err = addOpListSecretVersionIdsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecretVersionIds(options.Region), middleware.Before); err != nil {
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

// ListSecretVersionIdsAPIClient is a client that implements the
// ListSecretVersionIds operation.
type ListSecretVersionIdsAPIClient interface {
	ListSecretVersionIds(context.Context, *ListSecretVersionIdsInput, ...func(*Options)) (*ListSecretVersionIdsOutput, error)
}

var _ ListSecretVersionIdsAPIClient = (*Client)(nil)

// ListSecretVersionIdsPaginatorOptions is the paginator options for
// ListSecretVersionIds
type ListSecretVersionIdsPaginatorOptions struct {
	// The number of results to include in the response. If there are more results
	// available, in the response, Secrets Manager includes NextToken. To get the next
	// results, call ListSecretVersionIds again with the value from NextToken.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecretVersionIdsPaginator is a paginator for ListSecretVersionIds
type ListSecretVersionIdsPaginator struct {
	options   ListSecretVersionIdsPaginatorOptions
	client    ListSecretVersionIdsAPIClient
	params    *ListSecretVersionIdsInput
	nextToken *string
	firstPage bool
}

// NewListSecretVersionIdsPaginator returns a new ListSecretVersionIdsPaginator
func NewListSecretVersionIdsPaginator(client ListSecretVersionIdsAPIClient, params *ListSecretVersionIdsInput, optFns ...func(*ListSecretVersionIdsPaginatorOptions)) *ListSecretVersionIdsPaginator {
	if params == nil {
		params = &ListSecretVersionIdsInput{}
	}

	options := ListSecretVersionIdsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecretVersionIdsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecretVersionIdsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecretVersionIds page.
func (p *ListSecretVersionIdsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecretVersionIdsOutput, error) {
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

	result, err := p.client.ListSecretVersionIds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecretVersionIds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "ListSecretVersionIds",
	}
}
