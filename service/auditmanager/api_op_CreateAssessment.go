// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an assessment in Audit Manager.
func (c *Client) CreateAssessment(ctx context.Context, params *CreateAssessmentInput, optFns ...func(*Options)) (*CreateAssessmentOutput, error) {
	if params == nil {
		params = &CreateAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssessment", params, optFns, c.addOperationCreateAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentInput struct {

	// The assessment report storage destination for the assessment that's being
	// created.
	//
	// This member is required.
	AssessmentReportsDestination *types.AssessmentReportsDestination

	// The identifier for the framework that the assessment will be created from.
	//
	// This member is required.
	FrameworkId *string

	// The name of the assessment to be created.
	//
	// This member is required.
	Name *string

	// The list of roles for the assessment.
	//
	// This member is required.
	Roles []types.Role

	// The wrapper that contains the Amazon Web Services accounts and services that are
	// in scope for the assessment.
	//
	// This member is required.
	Scope *types.Scope

	// The optional description of the assessment to be created.
	Description *string

	// The tags that are associated with the assessment.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssessmentOutput struct {

	// An entity that defines the scope of audit evidence collected by Audit Manager.
	// An Audit Manager assessment is an implementation of an Audit Manager framework.
	Assessment *types.Assessment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssessment{}, middleware.After)
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
	if err = addOpCreateAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "CreateAssessment",
	}
}
