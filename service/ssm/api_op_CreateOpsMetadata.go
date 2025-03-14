// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If you create a new application in Application Manager, Amazon Web Services
// Systems Manager calls this API operation to specify information about the new
// application, including the application type.
func (c *Client) CreateOpsMetadata(ctx context.Context, params *CreateOpsMetadataInput, optFns ...func(*Options)) (*CreateOpsMetadataOutput, error) {
	if params == nil {
		params = &CreateOpsMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOpsMetadata", params, optFns, c.addOperationCreateOpsMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOpsMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOpsMetadataInput struct {

	// A resource ID for a new Application Manager application.
	//
	// This member is required.
	ResourceId *string

	// Metadata for a new Application Manager application.
	Metadata map[string]types.MetadataValue

	// Optional metadata that you assign to a resource. You can specify a maximum of
	// five tags for an OpsMetadata object. Tags enable you to categorize a resource in
	// different ways, such as by purpose, owner, or environment. For example, you
	// might want to tag an OpsMetadata object to identify an environment or target
	// Amazon Web Services Region. In this case, you could specify the following
	// key-value pairs:
	//
	// * Key=Environment,Value=Production
	//
	// *
	// Key=Region,Value=us-east-2
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateOpsMetadataOutput struct {

	// The Amazon Resource Name (ARN) of the OpsMetadata Object or blob created by the
	// call.
	OpsMetadataArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOpsMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOpsMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOpsMetadata{}, middleware.After)
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
	if err = addOpCreateOpsMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOpsMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOpsMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateOpsMetadata",
	}
}
