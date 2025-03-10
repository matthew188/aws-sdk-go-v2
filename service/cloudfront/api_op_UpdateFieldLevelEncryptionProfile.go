// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a field-level encryption profile.
func (c *Client) UpdateFieldLevelEncryptionProfile(ctx context.Context, params *UpdateFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*UpdateFieldLevelEncryptionProfileOutput, error) {
	if params == nil {
		params = &UpdateFieldLevelEncryptionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFieldLevelEncryptionProfile", params, optFns, c.addOperationUpdateFieldLevelEncryptionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFieldLevelEncryptionProfileInput struct {

	// Request to update a field-level encryption profile.
	//
	// This member is required.
	FieldLevelEncryptionProfileConfig *types.FieldLevelEncryptionProfileConfig

	// The ID of the field-level encryption profile request.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the profile
	// identity to update. For example: E2QWRUHAPOMQZL.
	IfMatch *string

	noSmithyDocumentSerde
}

type UpdateFieldLevelEncryptionProfileOutput struct {

	// The result of the field-level encryption profile request.
	ETag *string

	// Return the results of updating the profile.
	FieldLevelEncryptionProfile *types.FieldLevelEncryptionProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateFieldLevelEncryptionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateFieldLevelEncryptionProfile{}, middleware.After)
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
	if err = addOpUpdateFieldLevelEncryptionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateFieldLevelEncryptionProfile",
	}
}
