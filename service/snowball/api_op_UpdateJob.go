// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// While a job's JobState value is New, you can update some of the information
// associated with a job. Once the job changes to a different job state, usually
// within 60 minutes of the job being created, this action is no longer available.
func (c *Client) UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) {
	if params == nil {
		params = &UpdateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJob", params, optFns, c.addOperationUpdateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobInput struct {

	// The job ID of the job that you want to update, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string

	// The ID of the updated Address object.
	AddressId *string

	// The updated description of this job's JobMetadata object.
	Description *string

	// The updated ID for the forwarding address for a job. This field is not supported
	// in most regions.
	ForwardingAddressId *string

	// The new or updated Notification object.
	Notification *types.Notification

	// Specifies the service or services on the Snow Family device that your
	// transferred data will be exported from or imported into. Amazon Web Services
	// Snow Family supports Amazon S3 and NFS (Network File System) and the Amazon Web
	// Services Storage Gateway service Tape Gateway type.
	OnDeviceServiceConfiguration *types.OnDeviceServiceConfiguration

	// The updated JobResource object, or the updated JobResource object.
	Resources *types.JobResource

	// The new role Amazon Resource Name (ARN) that you want to associate with this
	// job. To create a role ARN, use the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)Identity
	// and Access Management (IAM) API action.
	RoleARN *string

	// The updated shipping option value of this job's ShippingDetails object.
	ShippingOption types.ShippingOption

	// The updated SnowballCapacityPreference of this job's JobMetadata object. The 50
	// TB Snowballs are only available in the US regions. For more information, see
	// "https://docs.aws.amazon.com/snowball/latest/snowcone-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide or
	// "https://docs.aws.amazon.com/snowball/latest/developer-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide.
	SnowballCapacityPreference types.SnowballCapacity

	noSmithyDocumentSerde
}

type UpdateJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateJob{}, middleware.After)
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
	if err = addOpUpdateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "UpdateJob",
	}
}
