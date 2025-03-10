// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a particular detector version.
func (c *Client) GetDetectorVersion(ctx context.Context, params *GetDetectorVersionInput, optFns ...func(*Options)) (*GetDetectorVersionOutput, error) {
	if params == nil {
		params = &GetDetectorVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDetectorVersion", params, optFns, c.addOperationGetDetectorVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorVersionInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	noSmithyDocumentSerde
}

type GetDetectorVersionOutput struct {

	// The detector version ARN.
	Arn *string

	// The timestamp when the detector version was created.
	CreatedTime *string

	// The detector version description.
	Description *string

	// The detector ID.
	DetectorId *string

	// The detector version ID.
	DetectorVersionId *string

	// The Amazon SageMaker model endpoints included in the detector version.
	ExternalModelEndpoints []string

	// The timestamp when the detector version was last updated.
	LastUpdatedTime *string

	// The model versions included in the detector version.
	ModelVersions []types.ModelVersion

	// The execution mode of the rule in the dectector FIRST_MATCHED indicates that
	// Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at
	// the first matched rule. Amazon Fraud dectector then provides the outcomes for
	// that single rule. ALL_MATCHED indicates that Amazon Fraud Detector evaluates all
	// rules and returns the outcomes for all matched rules. You can define and edit
	// the rule mode at the detector version level, when it is in draft status.
	RuleExecutionMode types.RuleExecutionMode

	// The rules included in the detector version.
	Rules []types.Rule

	// The status of the detector version.
	Status types.DetectorVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDetectorVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDetectorVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDetectorVersion{}, middleware.After)
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
	if err = addOpGetDetectorVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetectorVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDetectorVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetDetectorVersion",
	}
}
