// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a Dev Environment for a source repository in a
// project. Dev Environments are specific to the user who creates them.
func (c *Client) GetDevEnvironment(ctx context.Context, params *GetDevEnvironmentInput, optFns ...func(*Options)) (*GetDevEnvironmentOutput, error) {
	if params == nil {
		params = &GetDevEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevEnvironment", params, optFns, c.addOperationGetDevEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDevEnvironmentInput struct {

	// The system-generated unique ID of the Dev Environment for which you want to view
	// information. To retrieve a list of Dev Environment IDs, use ListDevEnvironments.
	//
	// This member is required.
	Id *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	noSmithyDocumentSerde
}

type GetDevEnvironmentOutput struct {

	// The system-generated unique ID of the user who created the Dev Environment.
	//
	// This member is required.
	CreatorId *string

	// The system-generated unique ID of the Dev Environment.
	//
	// This member is required.
	Id *string

	// The amount of time the Dev Environment will run without any activity detected
	// before stopping, in minutes.
	//
	// This member is required.
	InactivityTimeoutMinutes int32

	// The Amazon EC2 instace type to use for the Dev Environment.
	//
	// This member is required.
	InstanceType types.InstanceType

	// The time when the Dev Environment was last updated, in coordinated universal
	// time (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// Information about the amount of storage allocated to the Dev Environment. By
	// default, a Dev Environment is configured to have 16GB of persistent storage.
	//
	// This member is required.
	PersistentStorage *types.PersistentStorage

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The source repository that contains the branch cloned into the Dev Environment.
	//
	// This member is required.
	Repositories []types.DevEnvironmentRepositorySummary

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The current status of the Dev Environment.
	//
	// This member is required.
	Status types.DevEnvironmentStatus

	// The user-specified alias for the Dev Environment.
	Alias *string

	// Information about the integrated development environment (IDE) configured for
	// the Dev Environment.
	Ides []types.Ide

	// The reason for the status.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDevEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDevEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDevEnvironment{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetDevEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDevEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDevEnvironment",
	}
}
