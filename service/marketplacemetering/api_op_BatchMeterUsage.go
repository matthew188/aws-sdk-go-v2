// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/marketplacemetering/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// BatchMeterUsage is called from a SaaS application listed on AWS Marketplace to
// post metering records for a set of customers. For identical requests, the API is
// idempotent; requests can be retried with the same records or a subset of the
// input records. Every request to BatchMeterUsage is for one product. If you need
// to meter usage for multiple products, you must make multiple calls to
// BatchMeterUsage. Usage records are expected to be submitted as quickly as
// possible after the event that is being recorded, and are not accepted more than
// 6 hours after the event. BatchMeterUsage can process up to 25 UsageRecords at a
// time. A UsageRecord can optionally include multiple usage allocations, to
// provide customers with usage data split into buckets by tags that you define (or
// allow the customer to define). BatchMeterUsage returns a list of
// UsageRecordResult objects, showing the result for each UsageRecord, as well as a
// list of UnprocessedRecords, indicating errors in the service side that you
// should retry. BatchMeterUsage requests must be less than 1MB in size. For an
// example of using BatchMeterUsage, see  BatchMeterUsage code example
// (https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-batchmeterusage-example)
// in the AWS Marketplace Seller Guide.
func (c *Client) BatchMeterUsage(ctx context.Context, params *BatchMeterUsageInput, optFns ...func(*Options)) (*BatchMeterUsageOutput, error) {
	if params == nil {
		params = &BatchMeterUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchMeterUsage", params, optFns, c.addOperationBatchMeterUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchMeterUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A BatchMeterUsageRequest contains UsageRecords, which indicate quantities of
// usage within your application.
type BatchMeterUsageInput struct {

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of a new
	// product.
	//
	// This member is required.
	ProductCode *string

	// The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords
	// at a time.
	//
	// This member is required.
	UsageRecords []types.UsageRecord

	noSmithyDocumentSerde
}

// Contains the UsageRecords processed by BatchMeterUsage and any records that have
// failed due to transient error.
type BatchMeterUsageOutput struct {

	// Contains all UsageRecords processed by BatchMeterUsage. These records were
	// either honored by AWS Marketplace Metering Service or were invalid. Invalid
	// records should be fixed before being resubmitted.
	Results []types.UsageRecordResult

	// Contains all UsageRecords that were not processed by BatchMeterUsage. This is a
	// list of UsageRecords. You can retry the failed request by making another
	// BatchMeterUsage call with this list as input in the BatchMeterUsageRequest.
	UnprocessedRecords []types.UsageRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchMeterUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchMeterUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchMeterUsage{}, middleware.After)
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
	if err = addOpBatchMeterUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchMeterUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchMeterUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "BatchMeterUsage",
	}
}
