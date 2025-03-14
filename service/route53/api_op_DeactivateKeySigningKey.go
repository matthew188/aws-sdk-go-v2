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

// Deactivates a key-signing key (KSK) so that it will not be used for signing by
// DNSSEC. This operation changes the KSK status to INACTIVE.
func (c *Client) DeactivateKeySigningKey(ctx context.Context, params *DeactivateKeySigningKeyInput, optFns ...func(*Options)) (*DeactivateKeySigningKeyOutput, error) {
	if params == nil {
		params = &DeactivateKeySigningKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivateKeySigningKey", params, optFns, c.addOperationDeactivateKeySigningKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivateKeySigningKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateKeySigningKeyInput struct {

	// A unique string used to identify a hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// A string used to identify a key-signing key (KSK).
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeactivateKeySigningKeyOutput struct {

	// A complex type that describes change information about changes made to your
	// hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeactivateKeySigningKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeactivateKeySigningKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeactivateKeySigningKey{}, middleware.After)
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
	if err = addOpDeactivateKeySigningKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateKeySigningKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeactivateKeySigningKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeactivateKeySigningKey",
	}
}
