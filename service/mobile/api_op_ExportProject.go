// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports project configuration to a snapshot which can be downloaded and shared.
// Note that mobile app push credentials are encrypted in exported projects, so
// they can only be shared successfully within the same AWS account.
func (c *Client) ExportProject(ctx context.Context, params *ExportProjectInput, optFns ...func(*Options)) (*ExportProjectOutput, error) {
	if params == nil {
		params = &ExportProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportProject", params, optFns, c.addOperationExportProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used in requests to export project configuration details.
type ExportProjectInput struct {

	// Unique project identifier.
	//
	// This member is required.
	ProjectId *string

	noSmithyDocumentSerde
}

// Result structure used for requests to export project configuration details.
type ExportProjectOutput struct {

	// URL which can be used to download the exported project configuation file(s).
	DownloadUrl *string

	// URL which can be shared to allow other AWS users to create their own project in
	// AWS Mobile Hub with the same configuration as the specified project. This URL
	// pertains to a snapshot in time of the project configuration that is created when
	// this API is called. If you want to share additional changes to your project
	// configuration, then you will need to create and share a new snapshot by calling
	// this method again.
	ShareUrl *string

	// Unique identifier for the exported snapshot of the project configuration. This
	// snapshot identifier is included in the share URL.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportProject{}, middleware.After)
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
	if err = addOpExportProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "AWSMobileHubService",
		OperationName: "ExportProject",
	}
}
