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

// Returns status information for each of your Config managed rules. The status
// includes information such as the last time Config invoked the rule, the last
// time Config failed to invoke the rule, and the related error for the last
// failure.
func (c *Client) DescribeConfigRuleEvaluationStatus(ctx context.Context, params *DescribeConfigRuleEvaluationStatusInput, optFns ...func(*Options)) (*DescribeConfigRuleEvaluationStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigRuleEvaluationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigRuleEvaluationStatus", params, optFns, c.addOperationDescribeConfigRuleEvaluationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigRuleEvaluationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigRuleEvaluationStatusInput struct {

	// The name of the Config managed rules for which you want status information. If
	// you do not specify any names, Config returns status information for all Config
	// managed rules that you use.
	ConfigRuleNames []string

	// The number of rule evaluation results that you want returned. This parameter is
	// required if the rule limit for your account is more than the default of 150
	// rules. For information about requesting a rule limit increase, see Config Limits
	// (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config)
	// in the Amazon Web Services General Reference Guide.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConfigRuleEvaluationStatusOutput struct {

	// Status information about your Config managed rules.
	ConfigRulesEvaluationStatus []types.ConfigRuleEvaluationStatus

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigRuleEvaluationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigRuleEvaluationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigRuleEvaluationStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigRuleEvaluationStatus(options.Region), middleware.Before); err != nil {
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

// DescribeConfigRuleEvaluationStatusAPIClient is a client that implements the
// DescribeConfigRuleEvaluationStatus operation.
type DescribeConfigRuleEvaluationStatusAPIClient interface {
	DescribeConfigRuleEvaluationStatus(context.Context, *DescribeConfigRuleEvaluationStatusInput, ...func(*Options)) (*DescribeConfigRuleEvaluationStatusOutput, error)
}

var _ DescribeConfigRuleEvaluationStatusAPIClient = (*Client)(nil)

// DescribeConfigRuleEvaluationStatusPaginatorOptions is the paginator options for
// DescribeConfigRuleEvaluationStatus
type DescribeConfigRuleEvaluationStatusPaginatorOptions struct {
	// The number of rule evaluation results that you want returned. This parameter is
	// required if the rule limit for your account is more than the default of 150
	// rules. For information about requesting a rule limit increase, see Config Limits
	// (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config)
	// in the Amazon Web Services General Reference Guide.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConfigRuleEvaluationStatusPaginator is a paginator for
// DescribeConfigRuleEvaluationStatus
type DescribeConfigRuleEvaluationStatusPaginator struct {
	options   DescribeConfigRuleEvaluationStatusPaginatorOptions
	client    DescribeConfigRuleEvaluationStatusAPIClient
	params    *DescribeConfigRuleEvaluationStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeConfigRuleEvaluationStatusPaginator returns a new
// DescribeConfigRuleEvaluationStatusPaginator
func NewDescribeConfigRuleEvaluationStatusPaginator(client DescribeConfigRuleEvaluationStatusAPIClient, params *DescribeConfigRuleEvaluationStatusInput, optFns ...func(*DescribeConfigRuleEvaluationStatusPaginatorOptions)) *DescribeConfigRuleEvaluationStatusPaginator {
	if params == nil {
		params = &DescribeConfigRuleEvaluationStatusInput{}
	}

	options := DescribeConfigRuleEvaluationStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConfigRuleEvaluationStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConfigRuleEvaluationStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConfigRuleEvaluationStatus page.
func (p *DescribeConfigRuleEvaluationStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConfigRuleEvaluationStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConfigRuleEvaluationStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConfigRuleEvaluationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigRuleEvaluationStatus",
	}
}
