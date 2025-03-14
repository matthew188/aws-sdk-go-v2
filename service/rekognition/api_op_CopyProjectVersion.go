// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies a version of an Amazon Rekognition Custom Labels model from a source
// project to a destination project. The source and destination projects can be in
// different AWS accounts but must be in the same AWS Region. You can't copy a
// model to another AWS service. To copy a model version to a different AWS
// account, you need to create a resource-based policy known as a project policy.
// You attach the project policy to the source project by calling PutProjectPolicy.
// The project policy gives permission to copy the model version from a trusting
// AWS account to a trusted account. For more information creating and attaching a
// project policy, see Attaching a project policy (SDK) in the Amazon Rekognition
// Custom Labels Developer Guide. If you are copying a model version to a project
// in the same AWS account, you don't need to create a project policy. To copy a
// model, the destination project, source project, and source model version must
// already exist. Copying a model version takes a while to complete. To get the
// current status, call DescribeProjectVersions and check the value of Status in
// the ProjectVersionDescription object. The copy operation has finished when the
// value of Status is COPYING_COMPLETED.
func (c *Client) CopyProjectVersion(ctx context.Context, params *CopyProjectVersionInput, optFns ...func(*Options)) (*CopyProjectVersionOutput, error) {
	if params == nil {
		params = &CopyProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyProjectVersion", params, optFns, c.addOperationCopyProjectVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyProjectVersionInput struct {

	// The ARN of the project in the trusted AWS account that you want to copy the
	// model version to.
	//
	// This member is required.
	DestinationProjectArn *string

	// The S3 bucket and folder location where the training output for the source model
	// version is placed.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The ARN of the source project in the trusting AWS account.
	//
	// This member is required.
	SourceProjectArn *string

	// The ARN of the model version in the source project that you want to copy to a
	// destination project.
	//
	// This member is required.
	SourceProjectVersionArn *string

	// A name for the version of the model that's copied to the destination project.
	//
	// This member is required.
	VersionName *string

	// The identifier for your AWS Key Management Service key (AWS KMS key). You can
	// supply the Amazon Resource Name (ARN) of your KMS key, the ID of your KMS key,
	// an alias for your KMS key, or an alias ARN. The key is used to encrypt training
	// results and manifest files written to the output Amazon S3 bucket
	// (OutputConfig). If you choose to use your own KMS key, you need the following
	// permissions on the KMS key.
	//
	// * kms:CreateGrant
	//
	// * kms:DescribeKey
	//
	// *
	// kms:GenerateDataKey
	//
	// * kms:Decrypt
	//
	// If you don't specify a value for KmsKeyId,
	// images copied into the service are encrypted using a key that AWS owns and
	// manages.
	KmsKeyId *string

	// The key-value tags to assign to the model version.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CopyProjectVersionOutput struct {

	// The ARN of the copied model version in the destination project.
	ProjectVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyProjectVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopyProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyProjectVersion{}, middleware.After)
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
	if err = addOpCopyProjectVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyProjectVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyProjectVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CopyProjectVersion",
	}
}
