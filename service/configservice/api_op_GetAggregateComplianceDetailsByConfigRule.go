// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the evaluation results for the specified Config rule for a specific
// resource in a rule. The results indicate which Amazon Web Services resources
// were evaluated by the rule, when each resource was last evaluated, and whether
// each resource complies with the rule. The results can return an empty result
// page. But if you have a nextToken, the results are displayed on the next page.
func (c *Client) GetAggregateComplianceDetailsByConfigRule(ctx context.Context, params *GetAggregateComplianceDetailsByConfigRuleInput, optFns ...func(*Options)) (*GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	if params == nil {
		params = &GetAggregateComplianceDetailsByConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAggregateComplianceDetailsByConfigRule", params, optFns, c.addOperationGetAggregateComplianceDetailsByConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAggregateComplianceDetailsByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateComplianceDetailsByConfigRuleInput struct {

	// The 12-digit account ID of the source account.
	//
	// This member is required.
	AccountId *string

	// The source region from where the data is aggregated.
	//
	// This member is required.
	AwsRegion *string

	// The name of the Config rule for which you want compliance information.
	//
	// This member is required.
	ConfigRuleName *string

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The resource compliance status. For the
	// GetAggregateComplianceDetailsByConfigRuleRequest data type, Config supports only
	// the COMPLIANT and NON_COMPLIANT. Config does not support the NOT_APPLICABLE and
	// INSUFFICIENT_DATA values.
	ComplianceType types.ComplianceType

	// The maximum number of evaluation results returned on each page. The default is
	// 50. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetAggregateComplianceDetailsByConfigRuleOutput struct {

	// Returns an AggregateEvaluationResults object.
	AggregateEvaluationResults []types.AggregateEvaluationResult

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAggregateComplianceDetailsByConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateComplianceDetailsByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateComplianceDetailsByConfigRule{}, middleware.After)
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
	if err = addOpGetAggregateComplianceDetailsByConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateComplianceDetailsByConfigRule(options.Region), middleware.Before); err != nil {
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

// GetAggregateComplianceDetailsByConfigRuleAPIClient is a client that implements
// the GetAggregateComplianceDetailsByConfigRule operation.
type GetAggregateComplianceDetailsByConfigRuleAPIClient interface {
	GetAggregateComplianceDetailsByConfigRule(context.Context, *GetAggregateComplianceDetailsByConfigRuleInput, ...func(*Options)) (*GetAggregateComplianceDetailsByConfigRuleOutput, error)
}

var _ GetAggregateComplianceDetailsByConfigRuleAPIClient = (*Client)(nil)

// GetAggregateComplianceDetailsByConfigRulePaginatorOptions is the paginator
// options for GetAggregateComplianceDetailsByConfigRule
type GetAggregateComplianceDetailsByConfigRulePaginatorOptions struct {
	// The maximum number of evaluation results returned on each page. The default is
	// 50. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAggregateComplianceDetailsByConfigRulePaginator is a paginator for
// GetAggregateComplianceDetailsByConfigRule
type GetAggregateComplianceDetailsByConfigRulePaginator struct {
	options   GetAggregateComplianceDetailsByConfigRulePaginatorOptions
	client    GetAggregateComplianceDetailsByConfigRuleAPIClient
	params    *GetAggregateComplianceDetailsByConfigRuleInput
	nextToken *string
	firstPage bool
}

// NewGetAggregateComplianceDetailsByConfigRulePaginator returns a new
// GetAggregateComplianceDetailsByConfigRulePaginator
func NewGetAggregateComplianceDetailsByConfigRulePaginator(client GetAggregateComplianceDetailsByConfigRuleAPIClient, params *GetAggregateComplianceDetailsByConfigRuleInput, optFns ...func(*GetAggregateComplianceDetailsByConfigRulePaginatorOptions)) *GetAggregateComplianceDetailsByConfigRulePaginator {
	if params == nil {
		params = &GetAggregateComplianceDetailsByConfigRuleInput{}
	}

	options := GetAggregateComplianceDetailsByConfigRulePaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAggregateComplianceDetailsByConfigRulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAggregateComplianceDetailsByConfigRulePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAggregateComplianceDetailsByConfigRule page.
func (p *GetAggregateComplianceDetailsByConfigRulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetAggregateComplianceDetailsByConfigRule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAggregateComplianceDetailsByConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateComplianceDetailsByConfigRule",
	}
}
