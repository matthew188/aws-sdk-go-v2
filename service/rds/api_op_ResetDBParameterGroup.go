// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a DB parameter group to the engine/system default
// value. To reset specific parameters, provide a list of the following:
// ParameterName and ApplyMethod. To reset the entire DB parameter group, specify
// the DBParameterGroup name and ResetAllParameters parameters. When resetting the
// entire group, dynamic parameters are updated immediately and static parameters
// are set to pending-reboot to take effect on the next DB instance restart or
// RebootDBInstance request.
func (c *Client) ResetDBParameterGroup(ctx context.Context, params *ResetDBParameterGroupInput, optFns ...func(*Options)) (*ResetDBParameterGroupOutput, error) {
	if params == nil {
		params = &ResetDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetDBParameterGroup", params, optFns, c.addOperationResetDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetDBParameterGroupInput struct {

	// The name of the DB parameter group. Constraints:
	//
	// * Must match the name of an
	// existing DBParameterGroup.
	//
	// This member is required.
	DBParameterGroupName *string

	// To reset the entire DB parameter group, specify the DBParameterGroup name and
	// ResetAllParameters parameters. To reset specific parameters, provide a list of
	// the following: ParameterName and ApplyMethod. A maximum of 20 parameters can be
	// modified in a single request. MySQL Valid Values (for Apply method): immediate |
	// pending-reboot You can use the immediate value with dynamic parameters only. You
	// can use the pending-reboot value for both dynamic and static parameters, and
	// changes are applied when DB instance reboots. MariaDB Valid Values (for Apply
	// method): immediate | pending-reboot You can use the immediate value with dynamic
	// parameters only. You can use the pending-reboot value for both dynamic and
	// static parameters, and changes are applied when DB instance reboots. Oracle
	// Valid Values (for Apply method): pending-reboot
	Parameters []types.Parameter

	// A value that indicates whether to reset all parameters in the DB parameter group
	// to default values. By default, all parameters in the DB parameter group are
	// reset to default values.
	ResetAllParameters bool

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the ModifyDBParameterGroup or
// ResetDBParameterGroup action.
type ResetDBParameterGroupOutput struct {

	// The name of the DB parameter group.
	DBParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResetDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResetDBParameterGroup{}, middleware.After)
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
	if err = addOpResetDBParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetDBParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ResetDBParameterGroup",
	}
}
