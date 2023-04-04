// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a package from the specified Amazon OpenSearch Service domain. The
// package can't be in use with any OpenSearch index for the dissociation to
// succeed. The package is still available in OpenSearch Service for association
// later. For more information, see Custom packages for Amazon OpenSearch Service
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html).
func (c *Client) DissociatePackage(ctx context.Context, params *DissociatePackageInput, optFns ...func(*Options)) (*DissociatePackageOutput, error) {
	if params == nil {
		params = &DissociatePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DissociatePackage", params, optFns, c.addOperationDissociatePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DissociatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the DissociatePackage operation.
type DissociatePackageInput struct {

	// Name of the domain to dissociate the package from.
	//
	// This member is required.
	DomainName *string

	// Internal ID of the package to dissociate from the domain. Use
	// ListPackagesForDomain to find this value.
	//
	// This member is required.
	PackageID *string

	noSmithyDocumentSerde
}

// Container for the response returned by an DissociatePackage operation.
type DissociatePackageOutput struct {

	// Information about a package that has been dissociated from the domain.
	DomainPackageDetails *types.DomainPackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDissociatePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDissociatePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDissociatePackage{}, middleware.After)
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
	if err = addOpDissociatePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDissociatePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDissociatePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DissociatePackage",
	}
}
