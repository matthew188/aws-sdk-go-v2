// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the answer to a specific question in a workload review.
func (c *Client) UpdateAnswer(ctx context.Context, params *UpdateAnswerInput, optFns ...func(*Options)) (*UpdateAnswerOutput, error) {
	if params == nil {
		params = &UpdateAnswerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnswer", params, optFns, c.addOperationUpdateAnswerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnswerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to update answer.
type UpdateAnswerInput struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless, or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless. Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef.
	// Each lens is identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID of the question.
	//
	// This member is required.
	QuestionId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	//
	// This member is required.
	WorkloadId *string

	// A list of choices to update on a question in your workload. The String key
	// corresponds to the choice ID to be updated.
	ChoiceUpdates map[string]types.ChoiceUpdate

	// Defines whether this question is applicable to a lens review.
	IsApplicable bool

	// The notes associated with the workload.
	Notes *string

	// The reason why a question is not applicable to your workload.
	Reason types.AnswerReason

	// List of selected choice IDs in a question answer. The values entered replace the
	// previously selected choices.
	SelectedChoices []string

	noSmithyDocumentSerde
}

// Output of a update answer call.
type UpdateAnswerOutput struct {

	// An answer of the question.
	Answer *types.Answer

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless, or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless. Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef.
	// Each lens is identified by its LensSummary$LensAlias.
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnswerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnswer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnswer{}, middleware.After)
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
	if err = addOpUpdateAnswerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnswer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnswer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpdateAnswer",
	}
}
