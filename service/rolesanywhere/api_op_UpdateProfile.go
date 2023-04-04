// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the profile. A profile is configuration resource to list the roles that
// RolesAnywhere service is trusted to assume. In addition, by applying a profile
// you can scope-down permissions with IAM managed policies. Required permissions:
// rolesanywhere:UpdateProfile.
func (c *Client) UpdateProfile(ctx context.Context, params *UpdateProfileInput, optFns ...func(*Options)) (*UpdateProfileOutput, error) {
	if params == nil {
		params = &UpdateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProfile", params, optFns, c.addOperationUpdateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProfileInput struct {

	// The unique identifier of the profile.
	//
	// This member is required.
	ProfileId *string

	// The number of seconds the vended session credentials are valid for.
	DurationSeconds *int32

	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns []string

	// The name of the profile.
	Name *string

	// A list of IAM roles that this profile can assume in a CreateSession
	// (https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html)
	// operation.
	RoleArns []string

	// A session policy that applies to the trust boundary of the vended session
	// credentials.
	SessionPolicy *string

	noSmithyDocumentSerde
}

type UpdateProfileOutput struct {

	// The state of the profile after a read or write operation.
	Profile *types.ProfileDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProfile{}, middleware.After)
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
	if err = addOpUpdateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "UpdateProfile",
	}
}
