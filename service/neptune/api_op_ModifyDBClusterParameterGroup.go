// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a DB cluster parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue, and
// ApplyMethod. A maximum of 20 parameters can be modified in a single request.
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB cluster associated with
// the parameter group before the change can take effect. After you create a DB
// cluster parameter group, you should wait at least 5 minutes before creating your
// first DB cluster that uses that DB cluster parameter group as the default
// parameter group. This allows Amazon Neptune to fully complete the create action
// before the parameter group is used as the default for a new DB cluster. This is
// especially important for parameters that are critical when creating the default
// database for a DB cluster, such as the character set for the default database
// defined by the character_set_database parameter. You can use the Parameter
// Groups option of the Amazon Neptune console or the DescribeDBClusterParameters
// command to verify that your DB cluster parameter group has been created or
// modified.
func (c *Client) ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterParameterGroup", params, optFns, c.addOperationModifyDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterParameterGroupInput struct {

	// The name of the DB cluster parameter group to modify.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameters in the DB cluster parameter group to modify.
	//
	// This member is required.
	Parameters []types.Parameter

	noSmithyDocumentSerde
}

type ModifyDBClusterParameterGroupOutput struct {

	// The name of the DB cluster parameter group. Constraints:
	//
	// * Must be 1 to 255
	// letters or numbers.
	//
	// * First character must be a letter
	//
	// * Cannot end with a
	// hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase
	// string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterParameterGroup{}, middleware.After)
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
	if err = addOpModifyDBClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterParameterGroup",
	}
}
