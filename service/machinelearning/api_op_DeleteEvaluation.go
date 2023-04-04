// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns the DELETED status to an Evaluation, rendering it unusable. After
// invoking the DeleteEvaluation operation, you can use the GetEvaluation operation
// to verify that the status of the Evaluation changed to DELETED. Caution: The
// results of the DeleteEvaluation operation are irreversible.
func (c *Client) DeleteEvaluation(ctx context.Context, params *DeleteEvaluationInput, optFns ...func(*Options)) (*DeleteEvaluationOutput, error) {
	if params == nil {
		params = &DeleteEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEvaluation", params, optFns, c.addOperationDeleteEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEvaluationInput struct {

	// A user-supplied ID that uniquely identifies the Evaluation to delete.
	//
	// This member is required.
	EvaluationId *string

	noSmithyDocumentSerde
}

// Represents the output of a DeleteEvaluation operation. The output indicates that
// Amazon Machine Learning (Amazon ML) received the request. You can use the
// GetEvaluation operation and check the value of the Status parameter to see
// whether an Evaluation is marked as DELETED.
type DeleteEvaluationOutput struct {

	// A user-supplied ID that uniquely identifies the Evaluation. This value should be
	// identical to the value of the EvaluationId in the request.
	EvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEvaluation{}, middleware.After)
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
	if err = addOpDeleteEvaluationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEvaluation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEvaluation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteEvaluation",
	}
}
