// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new MediaPackage VOD Asset resource.
func (c *Client) CreateAsset(ctx context.Context, params *CreateAssetInput, optFns ...func(*Options)) (*CreateAssetOutput, error) {
	if params == nil {
		params = &CreateAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAsset", params, optFns, c.addOperationCreateAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A new MediaPackage VOD Asset configuration.
type CreateAssetInput struct {

	// The unique identifier for the Asset.
	//
	// This member is required.
	Id *string

	// The ID of the PackagingGroup for the Asset.
	//
	// This member is required.
	PackagingGroupId *string

	// ARN of the source object in S3.
	//
	// This member is required.
	SourceArn *string

	// The IAM role ARN used to access the source S3 bucket.
	//
	// This member is required.
	SourceRoleArn *string

	// The resource ID to include in SPEKE key requests.
	ResourceId *string

	// A collection of tags associated with a resource
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssetOutput struct {

	// The ARN of the Asset.
	Arn *string

	// The time the Asset was initially submitted for Ingest.
	CreatedAt *string

	// The list of egress endpoints available for the Asset.
	EgressEndpoints []types.EgressEndpoint

	// The unique identifier for the Asset.
	Id *string

	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string

	// The resource ID to include in SPEKE key requests.
	ResourceId *string

	// ARN of the source object in S3.
	SourceArn *string

	// The IAM role_arn used to access the source S3 bucket.
	SourceRoleArn *string

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAsset{}, middleware.After)
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
	if err = addOpCreateAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAsset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "CreateAsset",
	}
}
