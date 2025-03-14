// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a file import.
func (c *Client) StartImportFileTask(ctx context.Context, params *StartImportFileTaskInput, optFns ...func(*Options)) (*StartImportFileTaskOutput, error) {
	if params == nil {
		params = &StartImportFileTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImportFileTask", params, optFns, c.addOperationStartImportFileTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportFileTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportFileTaskInput struct {

	// A descriptive name for the request.
	//
	// This member is required.
	Name *string

	// The S3 bucket where the import file is located. The bucket name is required to
	// begin with migrationhub-strategy-.
	//
	// This member is required.
	S3Bucket *string

	// The Amazon S3 key name of the import file.
	//
	// This member is required.
	S3key *string

	// Specifies the source that the servers are coming from. By default, Strategy
	// Recommendations assumes that the servers specified in the import file are
	// available in AWS Application Discovery Service.
	DataSourceType types.DataSourceType

	// Groups the resources in the import file together with a unique name. This ID can
	// be as filter in ListApplicationComponents and ListServers.
	GroupId []types.Group

	// The S3 bucket where Strategy Recommendations uploads import results. The bucket
	// name is required to begin with migrationhub-strategy-.
	S3bucketForReportData *string

	noSmithyDocumentSerde
}

type StartImportFileTaskOutput struct {

	// The ID for a specific import task. The ID is unique within an AWS account.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportFileTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartImportFileTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartImportFileTask{}, middleware.After)
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
	if err = addOpStartImportFileTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImportFileTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImportFileTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "StartImportFileTask",
	}
}
