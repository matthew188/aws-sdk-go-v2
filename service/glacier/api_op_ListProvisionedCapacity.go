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

// This operation lists the provisioned capacity units for the specified AWS
// account.
func (c *Client) ListProvisionedCapacity(ctx context.Context, params *ListProvisionedCapacityInput, optFns ...func(*Options)) (*ListProvisionedCapacityOutput, error) {
	if params == nil {
		params = &ListProvisionedCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisionedCapacity", params, optFns, c.addOperationListProvisionedCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisionedCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisionedCapacityInput struct {

	// The AWS account ID of the account that owns the vault. You can either specify an
	// AWS account ID or optionally a single '-' (hyphen), in which case Amazon S3
	// Glacier uses the AWS account ID associated with the credentials used to sign the
	// request. If you use an account ID, don't include any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

type ListProvisionedCapacityOutput struct {

	// The response body contains the following JSON fields.
	ProvisionedCapacityList []types.ProvisionedCapacityDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProvisionedCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProvisionedCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisionedCapacity{}, middleware.After)
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
	if err = addOpListProvisionedCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisionedCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListProvisionedCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListProvisionedCapacity",
	}
}
