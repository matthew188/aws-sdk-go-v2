// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom model that replicates a source custom model that you
// import. The source model can be in your Amazon Web Services account or another
// one. If the source model is in another Amazon Web Services account, then it must
// have a resource-based policy that authorizes you to import it. The source model
// must be in the same Amazon Web Services Region that you're using when you
// import. You can't import a model that's in a different Region.
func (c *Client) ImportModel(ctx context.Context, params *ImportModelInput, optFns ...func(*Options)) (*ImportModelOutput, error) {
	if params == nil {
		params = &ImportModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportModel", params, optFns, c.addOperationImportModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportModelInput struct {

	// The Amazon Resource Name (ARN) of the custom model to import.
	//
	// This member is required.
	SourceModelArn *string

	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend
	// permission to use Amazon Key Management Service (KMS) to encrypt or decrypt the
	// custom model.
	DataAccessRoleArn *string

	// ID for the KMS key that Amazon Comprehend uses to encrypt trained custom models.
	// The ModelKmsKeyId can be either of the following formats:
	//
	// * KMS Key ID:
	// "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * Amazon Resource Name (ARN) of a KMS
	// Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	ModelKmsKeyId *string

	// The name to assign to the custom model that is created in Amazon Comprehend by
	// this import.
	ModelName *string

	// Tags to associate with the custom model that is created by this import. A tag is
	// a key-value pair that adds as a metadata to a resource used by Amazon
	// Comprehend. For example, a tag with "Sales" as the key might be added to a
	// resource to indicate its use by the sales department.
	Tags []types.Tag

	// The version name given to the custom model that is created by this import.
	// Version names can have a maximum of 256 characters. Alphanumeric characters,
	// hyphens (-) and underscores (_) are allowed. The version name must be unique
	// among all models with the same classifier name in the account/Region.
	VersionName *string

	noSmithyDocumentSerde
}

type ImportModelOutput struct {

	// The Amazon Resource Name (ARN) of the custom model being imported.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportModel{}, middleware.After)
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
	if err = addOpImportModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ImportModel",
	}
}
