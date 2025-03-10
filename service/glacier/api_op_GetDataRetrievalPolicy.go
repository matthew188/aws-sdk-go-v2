// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/matthew188/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/matthew188/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation returns the current data retrieval policy for the account and
// region specified in the GET request. For more information about data retrieval
// policies, see Amazon Glacier Data Retrieval Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/data-retrieval-policy.html).
func (c *Client) GetDataRetrievalPolicy(ctx context.Context, params *GetDataRetrievalPolicyInput, optFns ...func(*Options)) (*GetDataRetrievalPolicyOutput, error) {
	if params == nil {
		params = &GetDataRetrievalPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataRetrievalPolicy", params, optFns, c.addOperationGetDataRetrievalPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataRetrievalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetDataRetrievalPolicy.
type GetDataRetrievalPolicyInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to the GetDataRetrievalPolicy request.
type GetDataRetrievalPolicyOutput struct {

	// Contains the returned data retrieval policy in JSON format.
	Policy *types.DataRetrievalPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataRetrievalPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataRetrievalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataRetrievalPolicy{}, middleware.After)
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
	if err = addOpGetDataRetrievalPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataRetrievalPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetDataRetrievalPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "GetDataRetrievalPolicy",
	}
}
