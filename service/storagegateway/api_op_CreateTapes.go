// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one or more virtual tapes. You write data to the virtual tapes and then
// archive the tapes. This operation is only supported in the tape gateway type.
// Cache storage must be allocated to the gateway before you can create virtual
// tapes. Use the AddCache operation to add cache storage to a gateway.
func (c *Client) CreateTapes(ctx context.Context, params *CreateTapesInput, optFns ...func(*Options)) (*CreateTapesOutput, error) {
	if params == nil {
		params = &CreateTapesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTapes", params, optFns, c.addOperationCreateTapesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateTapesInput
type CreateTapesInput struct {

	// A unique identifier that you use to retry a request. If you retry a request, use
	// the same ClientToken you specified in the initial request. Using the same
	// ClientToken prevents creating the tape multiple times.
	//
	// This member is required.
	ClientToken *string

	// The unique Amazon Resource Name (ARN) that represents the gateway to associate
	// the virtual tapes with. Use the ListGateways operation to return a list of
	// gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The number of virtual tapes that you want to create.
	//
	// This member is required.
	NumTapesToCreate *int32

	// A prefix that you append to the barcode of the virtual tape you are creating.
	// This prefix makes the barcode unique. The prefix must be 1-4 characters in
	// length and must be one of the uppercase letters from A to Z.
	//
	// This member is required.
	TapeBarcodePrefix *string

	// The size, in bytes, of the virtual tapes that you want to create. The size must
	// be aligned by gigabyte (102410241024 bytes).
	//
	// This member is required.
	TapeSizeInBytes *int64

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The ID of the pool that you want to add your tape to for archiving. The tape in
	// this pool is archived in the S3 storage class that is associated with the pool.
	// When you use your backup application to eject the tape, the tape is archived
	// directly into the storage class (S3 Glacier or S3 Glacier Deep Archive) that
	// corresponds to the pool.
	PoolId *string

	// A list of up to 50 tags that can be assigned to a virtual tape. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag

	// Set to TRUE if the tape you are creating is to be configured as a
	// write-once-read-many (WORM) tape.
	Worm bool

	noSmithyDocumentSerde
}

// CreateTapeOutput
type CreateTapesOutput struct {

	// A list of unique Amazon Resource Names (ARNs) that represents the virtual tapes
	// that were created.
	TapeARNs []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTapesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTapes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTapes{}, middleware.After)
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
	if err = addOpCreateTapesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTapes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTapes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateTapes",
	}
}
