// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkmeetings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Chime SDK meeting in the specified media Region with no
// initial attendees. For more information about specifying media Regions, see
// Amazon Chime SDK Media Regions
// (https://docs.aws.amazon.com/chime/latest/dg/chime-sdk-meetings-regions.html) in
// the Amazon Chime Developer Guide. For more information about the Amazon Chime
// SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) CreateMeeting(ctx context.Context, params *CreateMeetingInput, optFns ...func(*Options)) (*CreateMeetingOutput, error) {
	if params == nil {
		params = &CreateMeetingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMeeting", params, optFns, c.addOperationCreateMeetingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeetingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingInput struct {

	// The unique identifier for the client request. Use a different token for
	// different meetings.
	//
	// This member is required.
	ClientRequestToken *string

	// The external meeting ID. Pattern: [-_&@+=,(){}\[\]\/«».:|'"#a-zA-Z0-9À-ÿ\s]*
	// Values that begin with aws: are reserved. You can't configure a value that uses
	// this prefix. Case insensitive.
	//
	// This member is required.
	ExternalMeetingId *string

	// The Region in which to create the meeting. Available values: af-south-1,
	// ap-northeast-1, ap-northeast-2, ap-south-1, ap-southeast-1, ap-southeast-2,
	// ca-central-1, eu-central-1, eu-north-1, eu-south-1, eu-west-1, eu-west-2,
	// eu-west-3, sa-east-1, us-east-1, us-east-2, us-west-1, us-west-2. Available
	// values in AWS GovCloud (US) Regions: us-gov-east-1, us-gov-west-1.
	//
	// This member is required.
	MediaRegion *string

	// Lists the audio and video features enabled for a meeting, such as echo
	// reduction.
	MeetingFeatures *types.MeetingFeaturesConfiguration

	// Reserved.
	MeetingHostId *string

	// The configuration for resource targets to receive notifications when meeting and
	// attendee events occur.
	NotificationsConfiguration *types.NotificationsConfiguration

	// When specified, replicates the media from the primary meeting to the new
	// meeting.
	PrimaryMeetingId *string

	// Applies one or more tags to an Amazon Chime SDK meeting. Note the following:
	//
	// *
	// Not all resources have tags. For a list of services with resources that support
	// tagging using this operation, see Services that support the Resource Groups
	// Tagging API
	// (https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/supported-services.html).
	// If the resource doesn't yet support this operation, the resource's service might
	// support tagging using its own API operations. For more information, refer to the
	// documentation for that service.
	//
	// * Each resource can have up to 50 tags. For
	// other limits, see Tag Naming and Usage Conventions
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html#tag-conventions)
	// in the AWS General Reference.
	//
	// * You can only tag resources that are located in
	// the specified AWS Region for the AWS account.
	//
	// * To add tags to a resource, you
	// need the necessary permissions for the service that the resource belongs to as
	// well as permissions for adding tags. For more information, see the documentation
	// for each service.
	//
	// Do not store personally identifiable information (PII) or
	// other confidential or sensitive information in tags. We use tags to provide you
	// with billing and administration services. Tags are not intended to be used for
	// private or sensitive data. Minimum permissions In addition to the
	// tag:TagResources permission required by this operation, you must also have the
	// tagging permission defined by the service that created the resource. For
	// example, to tag a ChimeSDKMeetings instance using the TagResources operation,
	// you must have both of the following permissions:
	// tag:TagResourcesChimeSDKMeetings:CreateTags Some services might have specific
	// requirements for tagging some resources. For example, to tag an Amazon S3
	// bucket, you must also have the s3:GetBucketTagging permission. If the expected
	// minimum permissions don't work, check the documentation for that service's
	// tagging APIs for more information.
	Tags []types.Tag

	// A consistent and opaque identifier, created and maintained by the builder to
	// represent a segment of their users.
	TenantIds []string

	noSmithyDocumentSerde
}

type CreateMeetingOutput struct {

	// The meeting information, including the meeting ID and MediaPlacement.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMeetingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeeting{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateMeetingMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMeetingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeeting(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMeeting struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMeeting) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMeeting) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMeetingInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMeetingInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMeetingMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMeeting{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMeeting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMeeting",
	}
}
