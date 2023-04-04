// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the settings for the Insights event selectors that you configured for
// your trail. GetInsightSelectors shows if CloudTrail Insights event logging is
// enabled on the trail, and if it is, which insight types are enabled. If you run
// GetInsightSelectors on a trail that does not have Insights events enabled, the
// operation throws the exception InsightNotEnabledException For more information,
// see Logging CloudTrail Insights Events for Trails
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html)
// in the CloudTrail User Guide.
func (c *Client) GetInsightSelectors(ctx context.Context, params *GetInsightSelectorsInput, optFns ...func(*Options)) (*GetInsightSelectorsOutput, error) {
	if params == nil {
		params = &GetInsightSelectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightSelectors", params, optFns, c.addOperationGetInsightSelectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightSelectorsInput struct {

	// Specifies the name of the trail or trail ARN. If you specify a trail name, the
	// string must meet the following requirements:
	//
	// * Contain only ASCII letters (a-z,
	// A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	//
	// * Start with a
	// letter or number, and end with a letter or number
	//
	// * Be between 3 and 128
	// characters
	//
	// * Have no adjacent periods, underscores or dashes. Names like
	// my-_namespace and my--namespace are not valid.
	//
	// * Not be in IP address format
	// (for example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must be in the
	// format: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	TrailName *string

	noSmithyDocumentSerde
}

type GetInsightSelectorsOutput struct {

	// A JSON string that contains the insight types you want to log on a trail. In
	// this release, ApiErrorRateInsight and ApiCallRateInsight are supported as
	// insight types.
	InsightSelectors []types.InsightSelector

	// The Amazon Resource Name (ARN) of a trail for which you want to get Insights
	// selectors.
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightSelectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInsightSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInsightSelectors{}, middleware.After)
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
	if err = addOpGetInsightSelectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightSelectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInsightSelectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetInsightSelectors",
	}
}
