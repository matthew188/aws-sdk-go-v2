// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Terminates the ML compute instance. Before terminating the instance, SageMaker
// disconnects the ML storage volume from it. SageMaker preserves the ML storage
// volume. SageMaker stops charging you for the ML compute instance when you call
// StopNotebookInstance. To access data on the ML storage volume for a notebook
// instance that has been terminated, call the StartNotebookInstance API.
// StartNotebookInstance launches another ML compute instance, configures it, and
// attaches the preserved ML storage volume so you can continue your work.
func (c *Client) StopNotebookInstance(ctx context.Context, params *StopNotebookInstanceInput, optFns ...func(*Options)) (*StopNotebookInstanceOutput, error) {
	if params == nil {
		params = &StopNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopNotebookInstance", params, optFns, c.addOperationStopNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopNotebookInstanceInput struct {

	// The name of the notebook instance to terminate.
	//
	// This member is required.
	NotebookInstanceName *string

	noSmithyDocumentSerde
}

type StopNotebookInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopNotebookInstance{}, middleware.After)
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
	if err = addOpStopNotebookInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopNotebookInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopNotebookInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopNotebookInstance",
	}
}
