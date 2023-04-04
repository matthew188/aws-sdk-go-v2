// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the template body for a specified stack. You can get the template for
// running or deleted stacks. For deleted stacks, GetTemplate returns the template
// for up to 90 days after the stack has been deleted. If the template doesn't
// exist, a ValidationError is returned.
func (c *Client) GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) {
	if params == nil {
		params = &GetTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplate", params, optFns, c.addOperationGetTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for a GetTemplate action.
type GetTemplateInput struct {

	// The name or Amazon Resource Name (ARN) of a change set for which CloudFormation
	// returns the associated template. If you specify a name, you must also specify
	// the StackName.
	ChangeSetName *string

	// The name or the unique stack ID that's associated with the stack, which aren't
	// always interchangeable:
	//
	// * Running stacks: You can specify either the stack's
	// name or its unique stack ID.
	//
	// * Deleted stacks: You must specify the unique
	// stack ID.
	//
	// Default: There is no default value.
	StackName *string

	// For templates that include transforms, the stage of the template that
	// CloudFormation returns. To get the user-submitted template, specify Original. To
	// get the template after CloudFormation has processed all transforms, specify
	// Processed. If the template doesn't include transforms, Original and Processed
	// return the same template. By default, CloudFormation specifies Processed.
	TemplateStage types.TemplateStage

	noSmithyDocumentSerde
}

// The output for GetTemplate action.
type GetTemplateOutput struct {

	// The stage of the template that you can retrieve. For stacks, the Original and
	// Processed templates are always available. For change sets, the Original template
	// is always available. After CloudFormation finishes creating the change set, the
	// Processed template becomes available.
	StagesAvailable []types.TemplateStage

	// Structure containing the template body. (For more information, go to Template
	// Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the CloudFormation User Guide.) CloudFormation returns the same template that
	// was used when the stack was created.
	TemplateBody *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTemplate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "GetTemplate",
	}
}
