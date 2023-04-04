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

// Provides a detailed description of the definition of an analysis. If you do not
// need to know details about the content of an Analysis, for instance if you are
// trying to check the status of a recently created or updated Analysis, use the
// DescribeAnalysis
// (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeAnalysis.html)
// instead.
func (c *Client) DescribeAnalysisDefinition(ctx context.Context, params *DescribeAnalysisDefinitionInput, optFns ...func(*Options)) (*DescribeAnalysisDefinitionOutput, error) {
	if params == nil {
		params = &DescribeAnalysisDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnalysisDefinition", params, optFns, c.addOperationDescribeAnalysisDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnalysisDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAnalysisDefinitionInput struct {

	// The ID of the analysis that you're describing. The ID is part of the URL of the
	// analysis.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the Amazon Web Services account that contains the analysis. You must
	// be using the Amazon Web Services account that the analysis is in.
	//
	// This member is required.
	AwsAccountId *string

	noSmithyDocumentSerde
}

type DescribeAnalysisDefinitionOutput struct {

	// The ID of the analysis described.
	AnalysisId *string

	// The definition of an analysis. A definition is the data model of all features in
	// a Dashboard, Template, or Analysis.
	Definition *types.AnalysisDefinition

	// Errors associated with the analysis.
	Errors []types.AnalysisError

	// The descriptive name of the analysis.
	Name *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// Status associated with the analysis.
	//
	// * CREATION_IN_PROGRESS
	//
	// *
	// CREATION_SUCCESSFUL
	//
	// * CREATION_FAILED
	//
	// * UPDATE_IN_PROGRESS
	//
	// *
	// UPDATE_SUCCESSFUL
	//
	// * UPDATE_FAILED
	//
	// * DELETED
	ResourceStatus types.ResourceStatus

	// The HTTP status of the request.
	Status int32

	// The ARN of the theme of the analysis.
	ThemeArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAnalysisDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAnalysisDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAnalysisDefinition{}, middleware.After)
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
	if err = addOpDescribeAnalysisDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnalysisDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAnalysisDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAnalysisDefinition",
	}
}
