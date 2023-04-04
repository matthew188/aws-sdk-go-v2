// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of your historical recommendation generations within the past
// 30 days.
func (c *Client) ListSavingsPlansPurchaseRecommendationGeneration(ctx context.Context, params *ListSavingsPlansPurchaseRecommendationGenerationInput, optFns ...func(*Options)) (*ListSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	if params == nil {
		params = &ListSavingsPlansPurchaseRecommendationGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSavingsPlansPurchaseRecommendationGeneration", params, optFns, c.addOperationListSavingsPlansPurchaseRecommendationGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSavingsPlansPurchaseRecommendationGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSavingsPlansPurchaseRecommendationGenerationInput struct {

	// The status of the recommendation generation.
	GenerationStatus types.GenerationStatus

	// The token to retrieve the next set of results.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	// The IDs for each specific recommendation.
	RecommendationIds []string

	noSmithyDocumentSerde
}

type ListSavingsPlansPurchaseRecommendationGenerationOutput struct {

	// The list of historical recommendation generations.
	GenerationSummaryList []types.GenerationSummary

	// The token to retrieve the next set of results.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSavingsPlansPurchaseRecommendationGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSavingsPlansPurchaseRecommendationGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSavingsPlansPurchaseRecommendationGeneration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSavingsPlansPurchaseRecommendationGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSavingsPlansPurchaseRecommendationGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "ListSavingsPlansPurchaseRecommendationGeneration",
	}
}
