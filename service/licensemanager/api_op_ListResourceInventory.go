// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists resources managed using Systems Manager inventory.
func (c *Client) ListResourceInventory(ctx context.Context, params *ListResourceInventoryInput, optFns ...func(*Options)) (*ListResourceInventoryOutput, error) {
	if params == nil {
		params = &ListResourceInventoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceInventory", params, optFns, c.addOperationListResourceInventoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceInventoryInput struct {

	// Filters to scope the results. The following filters and logical operators are
	// supported:
	//
	// * account_id - The ID of the Amazon Web Services account that owns
	// the resource. Logical operators are EQUALS | NOT_EQUALS.
	//
	// * application_name -
	// The name of the application. Logical operators are EQUALS | BEGINS_WITH.
	//
	// *
	// license_included - The type of license included. Logical operators are EQUALS |
	// NOT_EQUALS. Possible values are sql-server-enterprise | sql-server-standard |
	// sql-server-web | windows-server-datacenter.
	//
	// * platform - The platform of the
	// resource. Logical operators are EQUALS | BEGINS_WITH.
	//
	// * resource_id - The ID of
	// the resource. Logical operators are EQUALS | NOT_EQUALS.
	//
	// * tag: - The key/value
	// combination of a tag assigned to the resource. Logical operators are EQUALS
	// (single account) or EQUALS | NOT_EQUALS (cross account).
	Filters []types.InventoryFilter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceInventoryOutput struct {

	// Token for the next set of results.
	NextToken *string

	// Information about the resources.
	ResourceInventoryList []types.ResourceInventory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceInventoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceInventory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceInventory{}, middleware.After)
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
	if err = addOpListResourceInventoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceInventory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListResourceInventory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListResourceInventory",
	}
}
