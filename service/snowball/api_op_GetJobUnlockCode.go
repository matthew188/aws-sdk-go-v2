// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 360 days after the associated job has been
// created. The UnlockCode value is a 29-character code with 25 alphanumeric
// characters and 4 hyphens. This code is used to decrypt the manifest file when it
// is passed along with the manifest to the Snow device through the Snowball client
// when the client is started for the first time. The only valid status for calling
// this API is WithCustomer as the manifest and Unlock code values are used for
// securing your device and should only be used when you have the device. As a best
// practice, we recommend that you don't save a copy of the UnlockCode in the same
// location as the manifest file for that job. Saving these separately helps
// prevent unauthorized parties from gaining access to the Snow device associated
// with that job.
func (c *Client) GetJobUnlockCode(ctx context.Context, params *GetJobUnlockCodeInput, optFns ...func(*Options)) (*GetJobUnlockCodeOutput, error) {
	if params == nil {
		params = &GetJobUnlockCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJobUnlockCode", params, optFns, c.addOperationGetJobUnlockCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobUnlockCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobUnlockCodeInput struct {

	// The ID for the job that you want to get the UnlockCode value for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetJobUnlockCodeOutput struct {

	// The UnlockCode value for the specified job. The UnlockCode value can be accessed
	// for up to 360 days after the job has been created.
	UnlockCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetJobUnlockCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobUnlockCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobUnlockCode{}, middleware.After)
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
	if err = addOpGetJobUnlockCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobUnlockCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetJobUnlockCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetJobUnlockCode",
	}
}
