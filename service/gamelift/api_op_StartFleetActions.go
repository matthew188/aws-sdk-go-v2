// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resumes certain types of activity on fleet instances that were suspended with
// StopFleetActions
// (https://docs.aws.amazon.com/gamelift/latest/apireference/API_StopFleetActions.html).
// For multi-location fleets, fleet actions are managed separately for each
// location. Currently, this operation is used to restart a fleet's auto-scaling
// activity. This operation can be used in the following ways:
//
// * To restart
// actions on instances in the fleet's home Region, provide a fleet ID and the type
// of actions to resume.
//
// * To restart actions on instances in one of the fleet's
// remote locations, provide a fleet ID, a location name, and the type of actions
// to resume.
//
// If successful, GameLift once again initiates scaling events as
// triggered by the fleet's scaling policies. If actions on the fleet location were
// never stopped, this operation will have no effect. Learn more Setting up
// GameLift fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) StartFleetActions(ctx context.Context, params *StartFleetActionsInput, optFns ...func(*Options)) (*StartFleetActionsOutput, error) {
	if params == nil {
		params = &StartFleetActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFleetActions", params, optFns, c.addOperationStartFleetActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFleetActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFleetActionsInput struct {

	// List of actions to restart on the fleet.
	//
	// This member is required.
	Actions []types.FleetAction

	// A unique identifier for the fleet to restart actions on. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The fleet location to restart fleet actions for. Specify a location in the form
	// of an Amazon Web Services Region code, such as us-west-2.
	Location *string

	noSmithyDocumentSerde
}

type StartFleetActionsOutput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to a GameLift fleet resource and uniquely identifies it. ARNs are
	// unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912.
	FleetArn *string

	// A unique identifier for the fleet to restart actions on.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFleetActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartFleetActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFleetActions{}, middleware.After)
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
	if err = addOpStartFleetActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFleetActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartFleetActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StartFleetActions",
	}
}
