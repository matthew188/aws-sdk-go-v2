// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves recommendations for the specified session. To avoid retrieving the
// same recommendations in subsequent calls, use NotifyRecommendationsReceived
// (https://docs.aws.amazon.com/wisdom/latest/APIReference/API_NotifyRecommendationsReceived.html).
// This API supports long-polling behavior with the waitTimeSeconds parameter.
// Short poll is the default behavior and only returns recommendations already
// available. To perform a manual query against an assistant, use QueryAssistant
// (https://docs.aws.amazon.com/wisdom/latest/APIReference/API_QueryAssistant.html).
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	if params == nil {
		params = &GetRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendations", params, optFns, c.addOperationGetRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationsInput struct {

	// The identifier of the Wisdom assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The identifier of the session. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	SessionId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The duration (in seconds) for which the call waits for a recommendation to be
	// made available before returning. If a recommendation is available, the call
	// returns sooner than WaitTimeSeconds. If no messages are available and the wait
	// time expires, the call returns successfully with an empty list.
	WaitTimeSeconds int32

	noSmithyDocumentSerde
}

type GetRecommendationsOutput struct {

	// The recommendations.
	//
	// This member is required.
	Recommendations []types.RecommendationData

	// The triggers corresponding to recommendations.
	Triggers []types.RecommendationTrigger

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
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
	if err = addOpGetRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "GetRecommendations",
	}
}
