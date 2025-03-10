// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a campaign by either deploying a new solution or changing the value of
// the campaign's minProvisionedTPS parameter. To update a campaign, the campaign
// status must be ACTIVE or CREATE FAILED. Check the campaign status using the
// DescribeCampaign
// (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeCampaign.html)
// operation. You can still get recommendations from a campaign while an update is
// in progress. The campaign will use the previous solution version and campaign
// configuration to generate recommendations until the latest campaign update
// status is Active. For more information on campaigns, see CreateCampaign
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html).
func (c *Client) UpdateCampaign(ctx context.Context, params *UpdateCampaignInput, optFns ...func(*Options)) (*UpdateCampaignOutput, error) {
	if params == nil {
		params = &UpdateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCampaign", params, optFns, c.addOperationUpdateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCampaignInput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// This member is required.
	CampaignArn *string

	// The configuration details of a campaign.
	CampaignConfig *types.CampaignConfig

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second that Amazon Personalize will support.
	MinProvisionedTPS *int32

	// The ARN of a new solution version to deploy.
	SolutionVersionArn *string

	noSmithyDocumentSerde
}

type UpdateCampaignOutput struct {

	// The same campaign ARN as given in the request.
	CampaignArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCampaign{}, middleware.After)
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
	if err = addOpUpdateCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "UpdateCampaign",
	}
}
