// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new preset. For information about job templates see the User Guide at
// http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func (c *Client) CreatePreset(ctx context.Context, params *CreatePresetInput, optFns ...func(*Options)) (*CreatePresetOutput, error) {
	if params == nil {
		params = &CreatePresetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePreset", params, optFns, c.addOperationCreatePresetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePresetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresetInput struct {

	// The name of the preset you are creating.
	//
	// This member is required.
	Name *string

	// Settings for preset
	//
	// This member is required.
	Settings *types.PresetSettings

	// Optional. A category for the preset you are creating.
	Category *string

	// Optional. A description of the preset you are creating.
	Description *string

	// The tags that you want to add to the resource. You can tag resources with a
	// key-value pair or with only a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreatePresetOutput struct {

	// A preset is a collection of preconfigured media conversion settings that you
	// want MediaConvert to apply to the output during the conversion process.
	Preset *types.Preset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePresetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePreset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePreset{}, middleware.After)
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
	if err = addOpCreatePresetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePreset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePreset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "CreatePreset",
	}
}
