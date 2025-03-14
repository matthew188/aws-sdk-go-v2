// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/service/sso/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the STS short-term credentials for a given role name that is assigned to
// the user.
func (c *Client) GetRoleCredentials(ctx context.Context, params *GetRoleCredentialsInput, optFns ...func(*Options)) (*GetRoleCredentialsOutput, error) {
	if params == nil {
		params = &GetRoleCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoleCredentials", params, optFns, c.addOperationGetRoleCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRoleCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoleCredentialsInput struct {

	// The token issued by the CreateToken API call. For more information, see
	// CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the IAM Identity Center OIDC API Reference Guide.
	//
	// This member is required.
	AccessToken *string

	// The identifier for the AWS account that is assigned to the user.
	//
	// This member is required.
	AccountId *string

	// The friendly name of the role that is assigned to the user.
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type GetRoleCredentialsOutput struct {

	// The credentials for the role that is assigned to the user.
	RoleCredentials *types.RoleCredentials

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRoleCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRoleCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRoleCredentials{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpGetRoleCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoleCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoleCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRoleCredentials",
	}
}
