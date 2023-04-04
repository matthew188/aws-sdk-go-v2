// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a detector. After activation, you can only change a detector's ingestion
// delay and description.
func (c *Client) UpdateAnomalyDetector(ctx context.Context, params *UpdateAnomalyDetectorInput, optFns ...func(*Options)) (*UpdateAnomalyDetectorOutput, error) {
	if params == nil {
		params = &UpdateAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnomalyDetector", params, optFns, c.addOperationUpdateAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnomalyDetectorInput struct {

	// The ARN of the detector to update.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// Contains information about the configuration to which the detector will be
	// updated.
	AnomalyDetectorConfig *types.AnomalyDetectorConfig

	// The updated detector description.
	AnomalyDetectorDescription *string

	// The Amazon Resource Name (ARN) of an AWS KMS encryption key.
	KmsKeyArn *string

	noSmithyDocumentSerde
}

type UpdateAnomalyDetectorOutput struct {

	// The ARN of the updated detector.
	AnomalyDetectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnomalyDetector{}, middleware.After)
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
	if err = addOpUpdateAnomalyDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnomalyDetector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnomalyDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "UpdateAnomalyDetector",
	}
}
