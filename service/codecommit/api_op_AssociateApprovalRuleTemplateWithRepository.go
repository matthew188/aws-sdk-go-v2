// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between an approval rule template and a specified
// repository. Then, the next time a pull request is created in the repository
// where the destination reference (if specified) matches the destination reference
// (branch) for the pull request, an approval rule that matches the template
// conditions is automatically created for that pull request. If no destination
// references are specified in the template, an approval rule that matches the
// template contents is created for all pull requests in that repository.
func (c *Client) AssociateApprovalRuleTemplateWithRepository(ctx context.Context, params *AssociateApprovalRuleTemplateWithRepositoryInput, optFns ...func(*Options)) (*AssociateApprovalRuleTemplateWithRepositoryOutput, error) {
	if params == nil {
		params = &AssociateApprovalRuleTemplateWithRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateApprovalRuleTemplateWithRepository", params, optFns, c.addOperationAssociateApprovalRuleTemplateWithRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateApprovalRuleTemplateWithRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateApprovalRuleTemplateWithRepositoryInput struct {

	// The name for the approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The name of the repository that you want to associate with the template.
	//
	// This member is required.
	RepositoryName *string

	noSmithyDocumentSerde
}

type AssociateApprovalRuleTemplateWithRepositoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateApprovalRuleTemplateWithRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateApprovalRuleTemplateWithRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateApprovalRuleTemplateWithRepository{}, middleware.After)
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
	if err = addOpAssociateApprovalRuleTemplateWithRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateApprovalRuleTemplateWithRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateApprovalRuleTemplateWithRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "AssociateApprovalRuleTemplateWithRepository",
	}
}
