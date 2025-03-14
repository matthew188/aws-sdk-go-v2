// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops an entity recognizer training job while in progress. If the training job
// state is TRAINING, the job is marked for termination and put into the
// STOP_REQUESTED state. If the training job completes before it can be stopped, it
// is put into the TRAINED; otherwise the training job is stopped and putted into
// the STOPPED state and the service sends back an HTTP 200 response with an empty
// HTTP body.
func (c *Client) StopTrainingEntityRecognizer(ctx context.Context, params *StopTrainingEntityRecognizerInput, optFns ...func(*Options)) (*StopTrainingEntityRecognizerOutput, error) {
	if params == nil {
		params = &StopTrainingEntityRecognizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopTrainingEntityRecognizer", params, optFns, c.addOperationStopTrainingEntityRecognizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopTrainingEntityRecognizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTrainingEntityRecognizerInput struct {

	// The Amazon Resource Name (ARN) that identifies the entity recognizer currently
	// being trained.
	//
	// This member is required.
	EntityRecognizerArn *string

	noSmithyDocumentSerde
}

type StopTrainingEntityRecognizerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopTrainingEntityRecognizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopTrainingEntityRecognizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTrainingEntityRecognizer{}, middleware.After)
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
	if err = addOpStopTrainingEntityRecognizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopTrainingEntityRecognizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopTrainingEntityRecognizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StopTrainingEntityRecognizer",
	}
}
