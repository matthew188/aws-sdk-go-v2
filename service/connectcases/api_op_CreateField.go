// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a field in the Cases domain. This field is used to define the case
// object model (that is, defines what data can be captured on cases) in a Cases
// domain.
func (c *Client) CreateField(ctx context.Context, params *CreateFieldInput, optFns ...func(*Options)) (*CreateFieldOutput, error) {
	if params == nil {
		params = &CreateFieldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateField", params, optFns, c.addOperationCreateFieldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFieldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFieldInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The name of the field.
	//
	// This member is required.
	Name *string

	// Defines the data type, some system constraints, and default display of the
	// field.
	//
	// This member is required.
	Type types.FieldType

	// The description of the field.
	Description *string

	noSmithyDocumentSerde
}

type CreateFieldOutput struct {

	// The Amazon Resource Name (ARN) of the field.
	//
	// This member is required.
	FieldArn *string

	// The unique identifier of a field.
	//
	// This member is required.
	FieldId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFieldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateField{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateField{}, middleware.After)
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
	if err = addOpCreateFieldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateField(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateField(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cases",
		OperationName: "CreateField",
	}
}
