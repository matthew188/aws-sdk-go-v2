// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new subnet group. subnet groups must contain at least one subnet in at
// least two Availability Zones in the Amazon Web Services Region.
func (c *Client) CreateDBSubnetGroup(ctx context.Context, params *CreateDBSubnetGroupInput, optFns ...func(*Options)) (*CreateDBSubnetGroupOutput, error) {
	if params == nil {
		params = &CreateDBSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBSubnetGroup", params, optFns, c.addOperationCreateDBSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBSubnetGroup.
type CreateDBSubnetGroupInput struct {

	// The description for the subnet group.
	//
	// This member is required.
	DBSubnetGroupDescription *string

	// The name for the subnet group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 letters, numbers, periods,
	// underscores, spaces, or hyphens. Must not be default. Example: mySubnetgroup
	//
	// This member is required.
	DBSubnetGroupName *string

	// The Amazon EC2 subnet IDs for the subnet group.
	//
	// This member is required.
	SubnetIds []string

	// The tags to be assigned to the subnet group.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDBSubnetGroupOutput struct {

	// Detailed information about a subnet group.
	DBSubnetGroup *types.DBSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBSubnetGroup{}, middleware.After)
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
	if err = addOpCreateDBSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBSubnetGroup",
	}
}
