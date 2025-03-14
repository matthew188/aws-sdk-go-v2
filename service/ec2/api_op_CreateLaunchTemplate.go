// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch template. A launch template contains the parameters to launch
// an instance. When you launch an instance using RunInstances, you can specify a
// launch template instead of providing the launch parameters in the request. For
// more information, see Launch an instance from a launch template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// in the Amazon Elastic Compute Cloud User Guide. If you want to clone an existing
// launch template as the basis for creating a new launch template, you can use the
// Amazon EC2 console. The API, SDKs, and CLI do not support cloning a template.
// For more information, see Create a launch template from an existing launch
// template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template-from-existing-launch-template)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateLaunchTemplate(ctx context.Context, params *CreateLaunchTemplateInput, optFns ...func(*Options)) (*CreateLaunchTemplateOutput, error) {
	if params == nil {
		params = &CreateLaunchTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchTemplate", params, optFns, c.addOperationCreateLaunchTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchTemplateInput struct {

	// The information for the launch template.
	//
	// This member is required.
	LaunchTemplateData *types.RequestLaunchTemplateData

	// A name for the launch template.
	//
	// This member is required.
	LaunchTemplateName *string

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. For more information, see Ensuring idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to apply to the launch template on creation. To tag the launch
	// template, the resource type must be launch-template. To specify the tags for the
	// resources that are created when an instance is launched, you must use the
	// TagSpecifications parameter in the launch template data
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestLaunchTemplateData.html)
	// structure.
	TagSpecifications []types.TagSpecification

	// A description for the first version of the launch template.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateLaunchTemplateOutput struct {

	// Information about the launch template.
	LaunchTemplate *types.LaunchTemplate

	// If the launch template contains parameters or parameter combinations that are
	// not valid, an error code and an error message are returned for each issue that's
	// found.
	Warning *types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateLaunchTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLaunchTemplate{}, middleware.After)
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
	if err = addOpCreateLaunchTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLaunchTemplate",
	}
}
