// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an Identity and Access Management (IAM) role with an Neptune DB
// cluster.
func (c *Client) AddRoleToDBCluster(ctx context.Context, params *AddRoleToDBClusterInput, optFns ...func(*Options)) (*AddRoleToDBClusterOutput, error) {
	if params == nil {
		params = &AddRoleToDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddRoleToDBCluster", params, optFns, c.addOperationAddRoleToDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddRoleToDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddRoleToDBClusterInput struct {

	// The name of the DB cluster to associate the IAM role with.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) of the IAM role to associate with the Neptune DB
	// cluster, for example arn:aws:iam::123456789012:role/NeptuneAccessRole.
	//
	// This member is required.
	RoleArn *string

	// The name of the feature for the Neptune DB cluster that the IAM role is to be
	// associated with. For the list of supported feature names, see DBEngineVersion.
	FeatureName *string

	noSmithyDocumentSerde
}

type AddRoleToDBClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddRoleToDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddRoleToDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddRoleToDBCluster{}, middleware.After)
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
	if err = addOpAddRoleToDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddRoleToDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddRoleToDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AddRoleToDBCluster",
	}
}
