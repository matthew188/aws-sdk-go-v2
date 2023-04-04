// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates Amazon QuickSight customizations for the current Amazon Web Services
// Region. Currently, the only customization that you can use is a theme. You can
// use customizations for your Amazon Web Services account or, if you specify a
// namespace, for a Amazon QuickSight namespace instead. Customizations that apply
// to a namespace override customizations that apply to an Amazon Web Services
// account. To find out which customizations apply, use the
// DescribeAccountCustomization API operation.
func (c *Client) UpdateAccountCustomization(ctx context.Context, params *UpdateAccountCustomizationInput, optFns ...func(*Options)) (*UpdateAccountCustomizationOutput, error) {
	if params == nil {
		params = &UpdateAccountCustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountCustomization", params, optFns, c.addOperationUpdateAccountCustomizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountCustomizationInput struct {

	// The Amazon QuickSight customizations you're updating in the current Amazon Web
	// Services Region.
	//
	// This member is required.
	AccountCustomization *types.AccountCustomization

	// The ID for the Amazon Web Services account that you want to update Amazon
	// QuickSight customizations for.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that you want to update Amazon QuickSight customizations for.
	Namespace *string

	noSmithyDocumentSerde
}

type UpdateAccountCustomizationOutput struct {

	// The Amazon QuickSight customizations you're updating in the current Amazon Web
	// Services Region.
	AccountCustomization *types.AccountCustomization

	// The Amazon Resource Name (ARN) for the updated customization for this Amazon Web
	// Services account.
	Arn *string

	// The ID for the Amazon Web Services account that you want to update Amazon
	// QuickSight customizations for.
	AwsAccountId *string

	// The namespace associated with the customization that you're updating.
	Namespace *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccountCustomizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccountCustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccountCustomization{}, middleware.After)
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
	if err = addOpUpdateAccountCustomizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountCustomization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccountCustomization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateAccountCustomization",
	}
}
