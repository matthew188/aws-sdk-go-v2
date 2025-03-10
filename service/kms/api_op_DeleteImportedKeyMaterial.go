// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes key material that you previously imported. This operation makes the
// specified KMS key unusable. For more information about importing key material
// into KMS, see Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
// the Key Management Service Developer Guide. When the specified KMS key is in the
// PendingDeletion state, this operation does not change the KMS key's state.
// Otherwise, it changes the KMS key's state to PendingImport. After you delete key
// material, you can use ImportKeyMaterial to reimport the same key material into
// the KMS key. The KMS key that you use for this operation must be in a compatible
// key state. For details, see Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:DeleteImportedKeyMaterial
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * GetParametersForImport
//
// * ImportKeyMaterial
func (c *Client) DeleteImportedKeyMaterial(ctx context.Context, params *DeleteImportedKeyMaterialInput, optFns ...func(*Options)) (*DeleteImportedKeyMaterialOutput, error) {
	if params == nil {
		params = &DeleteImportedKeyMaterialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteImportedKeyMaterial", params, optFns, c.addOperationDeleteImportedKeyMaterialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteImportedKeyMaterialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteImportedKeyMaterialInput struct {

	// Identifies the KMS key from which you are deleting imported key material. The
	// Origin of the KMS key must be EXTERNAL. Specify the key ID or key ARN of the KMS
	// key. For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	noSmithyDocumentSerde
}

type DeleteImportedKeyMaterialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteImportedKeyMaterialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteImportedKeyMaterial{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteImportedKeyMaterial{}, middleware.After)
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
	if err = addOpDeleteImportedKeyMaterialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteImportedKeyMaterial(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteImportedKeyMaterial(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DeleteImportedKeyMaterial",
	}
}
