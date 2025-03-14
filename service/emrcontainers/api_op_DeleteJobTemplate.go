// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a job template. Job template stores values of StartJobRun API request in
// a template and can be used to start a job run. Job template allows two use
// cases: avoid repeating recurring StartJobRun API request values, enforcing
// certain values in StartJobRun API request.
func (c *Client) DeleteJobTemplate(ctx context.Context, params *DeleteJobTemplateInput, optFns ...func(*Options)) (*DeleteJobTemplateOutput, error) {
	if params == nil {
		params = &DeleteJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJobTemplate", params, optFns, c.addOperationDeleteJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobTemplateInput struct {

	// The ID of the job template that will be deleted.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DeleteJobTemplateOutput struct {

	// This output contains the ID of the job template that was deleted.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJobTemplate{}, middleware.After)
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
	if err = addOpDeleteJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "DeleteJobTemplate",
	}
}
