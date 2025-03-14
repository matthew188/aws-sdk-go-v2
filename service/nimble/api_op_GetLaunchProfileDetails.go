// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Launch profile details include the launch profile resource and summary
// information of resources that are used by, or available to, the launch profile.
// This includes the name and description of all studio components used by the
// launch profiles, and the name and description of streaming images that can be
// used with this launch profile.
func (c *Client) GetLaunchProfileDetails(ctx context.Context, params *GetLaunchProfileDetailsInput, optFns ...func(*Options)) (*GetLaunchProfileDetailsOutput, error) {
	if params == nil {
		params = &GetLaunchProfileDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLaunchProfileDetails", params, optFns, c.addOperationGetLaunchProfileDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLaunchProfileDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLaunchProfileDetailsInput struct {

	// The ID of the launch profile used to control access from the streaming session.
	//
	// This member is required.
	LaunchProfileId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	noSmithyDocumentSerde
}

type GetLaunchProfileDetailsOutput struct {

	// The launch profile.
	LaunchProfile *types.LaunchProfile

	// A collection of streaming images.
	StreamingImages []types.StreamingImage

	// A collection of studio component summaries.
	StudioComponentSummaries []types.StudioComponentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLaunchProfileDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLaunchProfileDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLaunchProfileDetails{}, middleware.After)
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
	if err = addOpGetLaunchProfileDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLaunchProfileDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLaunchProfileDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "GetLaunchProfileDetails",
	}
}
