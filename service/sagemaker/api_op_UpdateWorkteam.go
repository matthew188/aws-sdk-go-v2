// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing work team with new member definitions or description.
func (c *Client) UpdateWorkteam(ctx context.Context, params *UpdateWorkteamInput, optFns ...func(*Options)) (*UpdateWorkteamOutput, error) {
	if params == nil {
		params = &UpdateWorkteamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkteam", params, optFns, c.addOperationUpdateWorkteamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkteamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkteamInput struct {

	// The name of the work team to update.
	//
	// This member is required.
	WorkteamName *string

	// An updated description for the work team.
	Description *string

	// A list of MemberDefinition objects that contains objects that identify the
	// workers that make up the work team. Workforces can be created using Amazon
	// Cognito or your own OIDC Identity Provider (IdP). For private workforces created
	// using Amazon Cognito use CognitoMemberDefinition. For workforces created using
	// your own OIDC identity provider (IdP) use OidcMemberDefinition. You should not
	// provide input for both of these parameters in a single request. For workforces
	// created using Amazon Cognito, private work teams correspond to Amazon Cognito
	// user groups within the user pool used to create a workforce. All of the
	// CognitoMemberDefinition objects that make up the member definition must have the
	// same ClientId and UserPool values. To add a Amazon Cognito user group to an
	// existing worker pool, see Adding groups to a User Pool. For more information
	// about user pools, see Amazon Cognito User Pools
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html).
	// For workforces created using your own OIDC IdP, specify the user groups that you
	// want to include in your private work team in OidcMemberDefinition by listing
	// those groups in Groups. Be aware that user groups that are already in the work
	// team must also be listed in Groups when you make this request to remain on the
	// work team. If you do not include these user groups, they will no longer be
	// associated with the work team you update.
	MemberDefinitions []types.MemberDefinition

	// Configures SNS topic notifications for available or expiring work items
	NotificationConfiguration *types.NotificationConfiguration

	noSmithyDocumentSerde
}

type UpdateWorkteamOutput struct {

	// A Workteam object that describes the updated work team.
	//
	// This member is required.
	Workteam *types.Workteam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkteamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkteam{}, middleware.After)
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
	if err = addOpUpdateWorkteamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkteam(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkteam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateWorkteam",
	}
}
