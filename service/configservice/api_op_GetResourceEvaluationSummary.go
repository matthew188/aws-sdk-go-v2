// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a summary of resource evaluation for the specified resource evaluation
// ID from the proactive rules that were run. The results indicate which evaluation
// context was used to evaluate the rules, which resource details were evaluated,
// the evaluation mode that was run, and whether the resource details comply with
// the configuration of the proactive rules. To see additional information about
// the evaluation result, such as which rule flagged a resource as NON_COMPLIANT,
// use the GetComplianceDetailsByResource
// (https://docs.aws.amazon.com/config/latest/APIReference/API_GetComplianceDetailsByResource.html)
// API. For more information, see the Examples
// (https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceEvaluationSummary.html#API_GetResourceEvaluationSummary_Examples)
// section.
func (c *Client) GetResourceEvaluationSummary(ctx context.Context, params *GetResourceEvaluationSummaryInput, optFns ...func(*Options)) (*GetResourceEvaluationSummaryOutput, error) {
	if params == nil {
		params = &GetResourceEvaluationSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceEvaluationSummary", params, optFns, c.addOperationGetResourceEvaluationSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceEvaluationSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceEvaluationSummaryInput struct {

	// The unique ResourceEvaluationId of Amazon Web Services resource execution for
	// which you want to retrieve the evaluation summary.
	//
	// This member is required.
	ResourceEvaluationId *string

	noSmithyDocumentSerde
}

type GetResourceEvaluationSummaryOutput struct {

	// The compliance status of the resource evaluation summary.
	Compliance types.ComplianceType

	// Returns an EvaluationContext object.
	EvaluationContext *types.EvaluationContext

	// Lists results of the mode that you requested to retrieve the resource evaluation
	// summary. The valid values are Detective or Proactive.
	EvaluationMode types.EvaluationMode

	// The start timestamp when Config rule starts evaluating compliance for the
	// provided resource details.
	EvaluationStartTimestamp *time.Time

	// Returns an EvaluationStatus object.
	EvaluationStatus *types.EvaluationStatus

	// Returns a ResourceDetails object.
	ResourceDetails *types.ResourceDetails

	// The unique ResourceEvaluationId of Amazon Web Services resource execution for
	// which you want to retrieve the evaluation summary.
	ResourceEvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceEvaluationSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourceEvaluationSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourceEvaluationSummary{}, middleware.After)
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
	if err = addOpGetResourceEvaluationSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceEvaluationSummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceEvaluationSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetResourceEvaluationSummary",
	}
}
