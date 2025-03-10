// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the backend environments for an Amplify app.
func (c *Client) ListBackendEnvironments(ctx context.Context, params *ListBackendEnvironmentsInput, optFns ...func(*Options)) (*ListBackendEnvironmentsOutput, error) {
	if params == nil {
		params = &ListBackendEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackendEnvironments", params, optFns, c.addOperationListBackendEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackendEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list backend environments request.
type ListBackendEnvironmentsInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment
	EnvironmentName *string

	// The maximum number of records to list in a single response.
	MaxResults int32

	// A pagination token. Set to null to start listing backend environments from the
	// start. If a non-null pagination token is returned in a result, pass its value in
	// here to list more backend environments.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list backend environments result.
type ListBackendEnvironmentsOutput struct {

	// The list of backend environments for an Amplify app.
	//
	// This member is required.
	BackendEnvironments []types.BackendEnvironment

	// A pagination token. If a non-null pagination token is returned in a result, pass
	// its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackendEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackendEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackendEnvironments{}, middleware.After)
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
	if err = addOpListBackendEnvironmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackendEnvironments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListBackendEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListBackendEnvironments",
	}
}
