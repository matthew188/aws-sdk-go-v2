// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) aggregated usage data for an account.
func (c *Client) GetUsageTotals(ctx context.Context, params *GetUsageTotalsInput, optFns ...func(*Options)) (*GetUsageTotalsOutput, error) {
	if params == nil {
		params = &GetUsageTotalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageTotals", params, optFns, c.addOperationGetUsageTotalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageTotalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageTotalsInput struct {

	// The inclusive time period to retrieve the data for. Valid values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days. If you don't specify a value for this parameter, Amazon
	// Macie provides aggregated usage data for the preceding 30 days.
	TimeRange *string

	noSmithyDocumentSerde
}

type GetUsageTotalsOutput struct {

	// The inclusive time period that the usage data applies to. Possible values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days.
	TimeRange types.TimeRange

	// An array of objects that contains the results of the query. Each object contains
	// the data for a specific usage metric.
	UsageTotals []types.UsageTotal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsageTotalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageTotals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageTotals{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageTotals(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUsageTotals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetUsageTotals",
	}
}
