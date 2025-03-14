// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates a key-signing key (KSK) so that it can be used for signing by DNSSEC.
// This operation changes the KSK status to ACTIVE.
func (c *Client) ActivateKeySigningKey(ctx context.Context, params *ActivateKeySigningKeyInput, optFns ...func(*Options)) (*ActivateKeySigningKeyOutput, error) {
	if params == nil {
		params = &ActivateKeySigningKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivateKeySigningKey", params, optFns, c.addOperationActivateKeySigningKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivateKeySigningKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ActivateKeySigningKeyInput struct {

	// A unique string used to identify a hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// A string used to identify a key-signing key (KSK). Name can include numbers,
	// letters, and underscores (_). Name must be unique for each key-signing key in
	// the same hosted zone.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type ActivateKeySigningKeyOutput struct {

	// A complex type that describes change information about changes made to your
	// hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivateKeySigningKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpActivateKeySigningKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpActivateKeySigningKey{}, middleware.After)
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
	if err = addOpActivateKeySigningKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivateKeySigningKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opActivateKeySigningKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ActivateKeySigningKey",
	}
}
