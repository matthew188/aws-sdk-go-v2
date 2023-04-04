// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a dataset import job created using the CreateDatasetImportJob
// (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html)
// operation. You can delete only dataset import jobs that have a status of ACTIVE
// or CREATE_FAILED. To get the status, use the DescribeDatasetImportJob
// (https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetImportJob.html)
// operation.
func (c *Client) DeleteDatasetImportJob(ctx context.Context, params *DeleteDatasetImportJobInput, optFns ...func(*Options)) (*DeleteDatasetImportJobOutput, error) {
	if params == nil {
		params = &DeleteDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDatasetImportJob", params, optFns, c.addOperationDeleteDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatasetImportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset import job to delete.
	//
	// This member is required.
	DatasetImportJobArn *string

	noSmithyDocumentSerde
}

type DeleteDatasetImportJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDatasetImportJob{}, middleware.After)
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
	if err = addOpDeleteDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeleteDatasetImportJob",
	}
}
