// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops an InferSNOMEDCT inference job in progress.
func (c *Client) StopSNOMEDCTInferenceJob(ctx context.Context, params *StopSNOMEDCTInferenceJobInput, optFns ...func(*Options)) (*StopSNOMEDCTInferenceJobOutput, error) {
	if params == nil {
		params = &StopSNOMEDCTInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopSNOMEDCTInferenceJob", params, optFns, c.addOperationStopSNOMEDCTInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopSNOMEDCTInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopSNOMEDCTInferenceJobInput struct {

	// The job id of the asynchronous InferSNOMEDCT job to be stopped.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type StopSNOMEDCTInferenceJobOutput struct {

	// The identifier generated for the job. To get the status of job, use this
	// identifier with the DescribeSNOMEDCTInferenceJob operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopSNOMEDCTInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopSNOMEDCTInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopSNOMEDCTInferenceJob{}, middleware.After)
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
	if err = addOpStopSNOMEDCTInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopSNOMEDCTInferenceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopSNOMEDCTInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StopSNOMEDCTInferenceJob",
	}
}
