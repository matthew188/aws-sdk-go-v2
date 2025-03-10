// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Amazon GuardDuty detector specified by the detectorId. There might
// be regional differences because some data sources might not be available in all
// the Amazon Web Services Regions where GuardDuty is presently supported. For more
// information, see Regions and endpoints
// (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html).
func (c *Client) UpdateDetector(ctx context.Context, params *UpdateDetectorInput, optFns ...func(*Options)) (*UpdateDetectorOutput, error) {
	if params == nil {
		params = &UpdateDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDetector", params, optFns, c.addOperationUpdateDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorInput struct {

	// The unique ID of the detector to update.
	//
	// This member is required.
	DetectorId *string

	// Describes which data sources will be updated. There might be regional
	// differences because some data sources might not be available in all the Amazon
	// Web Services Regions where GuardDuty is presently supported. For more
	// information, see Regions and endpoints
	// (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html).
	//
	// Deprecated: This parameter is deprecated, use Features instead
	DataSources *types.DataSourceConfigurations

	// Specifies whether the detector is enabled or not enabled.
	Enable bool

	// Provides the features that will be updated for the detector.
	Features []types.DetectorFeatureConfiguration

	// An enum value that specifies how frequently findings are exported, such as to
	// CloudWatch Events.
	FindingPublishingFrequency types.FindingPublishingFrequency

	noSmithyDocumentSerde
}

type UpdateDetectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDetector{}, middleware.After)
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
	if err = addOpUpdateDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateDetector",
	}
}
