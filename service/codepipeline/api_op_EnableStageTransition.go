// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables artifacts in a pipeline to transition to a stage in a pipeline.
func (c *Client) EnableStageTransition(ctx context.Context, params *EnableStageTransitionInput, optFns ...func(*Options)) (*EnableStageTransitionOutput, error) {
	if params == nil {
		params = &EnableStageTransitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableStageTransition", params, optFns, c.addOperationEnableStageTransitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableStageTransitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an EnableStageTransition action.
type EnableStageTransitionInput struct {

	// The name of the pipeline in which you want to enable the flow of artifacts from
	// one stage to another.
	//
	// This member is required.
	PipelineName *string

	// The name of the stage where you want to enable the transition of artifacts,
	// either into the stage (inbound) or from that stage to the next stage (outbound).
	//
	// This member is required.
	StageName *string

	// Specifies whether artifacts are allowed to enter the stage and be processed by
	// the actions in that stage (inbound) or whether already processed artifacts are
	// allowed to transition to the next stage (outbound).
	//
	// This member is required.
	TransitionType types.StageTransitionType

	noSmithyDocumentSerde
}

type EnableStageTransitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableStageTransitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableStageTransition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableStageTransition{}, middleware.After)
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
	if err = addOpEnableStageTransitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableStageTransition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableStageTransition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "EnableStageTransition",
	}
}
