// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an EBS volume that can be attached to an instance in the same
// Availability Zone. You can create a new empty volume or restore a volume from an
// EBS snapshot. Any Amazon Web Services Marketplace product codes from the
// snapshot are propagated to the volume. You can create encrypted volumes.
// Encrypted volumes must be attached to instances that support Amazon EBS
// encryption. Volumes that are created from encrypted snapshots are also
// automatically encrypted. For more information, see Amazon EBS encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. You can tag your volumes during
// creation. For more information, see Tag your Amazon EC2 resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
// Amazon Elastic Compute Cloud User Guide. For more information, see Create an
// Amazon EBS volume
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateVolume(ctx context.Context, params *CreateVolumeInput, optFns ...func(*Options)) (*CreateVolumeOutput, error) {
	if params == nil {
		params = &CreateVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVolume", params, optFns, c.addOperationCreateVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVolumeInput struct {

	// The Availability Zone in which to create the volume.
	//
	// This member is required.
	AvailabilityZone *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Indicates whether the volume should be encrypted. The effect of setting the
	// encryption state to true depends on the volume origin (new or from a snapshot),
	// starting encryption state, ownership, and whether encryption by default is
	// enabled. For more information, see Encryption by default
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default)
	// in the Amazon Elastic Compute Cloud User Guide. Encrypted Amazon EBS volumes
	// must be attached to instances that support Amazon EBS encryption. For more
	// information, see Supported instance types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
	Encrypted *bool

	// The number of I/O operations per second (IOPS). For gp3, io1, and io2 volumes,
	// this represents the number of IOPS that are provisioned for the volume. For gp2
	// volumes, this represents the baseline performance of the volume and the rate at
	// which the volume accumulates I/O credits for bursting. The following are the
	// supported values for each volume type:
	//
	// * gp3: 3,000-16,000 IOPS
	//
	// * io1:
	// 100-64,000 IOPS
	//
	// * io2: 100-64,000 IOPS
	//
	// io1 and io2 volumes support up to
	// 64,000 IOPS only on Instances built on the Nitro System
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families support performance up to 32,000 IOPS. This parameter is
	// required for io1 and io2 volumes. The default for gp3 volumes is 3,000 IOPS.
	// This parameter is not supported for gp2, st1, sc1, or standard volumes.
	Iops *int32

	// The identifier of the Key Management Service (KMS) KMS key to use for Amazon EBS
	// encryption. If this parameter is not specified, your KMS key for Amazon EBS is
	// used. If KmsKeyId is specified, the encrypted state must be true. You can
	// specify the KMS key using any of the following:
	//
	// * Key ID. For example,
	// 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Key alias. For example,
	// alias/ExampleAlias.
	//
	// * Key ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// *
	// Alias ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// Amazon Web Services
	// authenticates the KMS key asynchronously. Therefore, if you specify an ID,
	// alias, or ARN that is not valid, the action can appear to complete, but
	// eventually fails.
	KmsKeyId *string

	// Indicates whether to enable Amazon EBS Multi-Attach. If you enable Multi-Attach,
	// you can attach the volume to up to 16 Instances built on the Nitro System
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances)
	// in the same Availability Zone. This parameter is supported with io1 and io2
	// volumes only. For more information, see  Amazon EBS Multi-Attach
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volumes-multi.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	MultiAttachEnabled *bool

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The size of the volume, in GiBs. You must specify either a snapshot ID or a
	// volume size. If you specify a snapshot, the default is the snapshot size. You
	// can specify a volume size that is equal to or larger than the snapshot size. The
	// following are the supported volumes sizes for each volume type:
	//
	// * gp2 and gp3:
	// 1-16,384
	//
	// * io1 and io2: 4-16,384
	//
	// * st1 and sc1: 125-16,384
	//
	// * standard:
	// 1-1,024
	Size *int32

	// The snapshot from which to create the volume. You must specify either a snapshot
	// ID or a volume size.
	SnapshotId *string

	// The tags to apply to the volume during creation.
	TagSpecifications []types.TagSpecification

	// The throughput to provision for a volume, with a maximum of 1,000 MiB/s. This
	// parameter is valid only for gp3 volumes. Valid Range: Minimum value of 125.
	// Maximum value of 1000.
	Throughput *int32

	// The volume type. This parameter can be one of the following values:
	//
	// * General
	// Purpose SSD: gp2 | gp3
	//
	// * Provisioned IOPS SSD: io1 | io2
	//
	// * Throughput
	// Optimized HDD: st1
	//
	// * Cold HDD: sc1
	//
	// * Magnetic: standard
	//
	// Throughput Optimized
	// HDD (st1) and Cold HDD (sc1) volumes can't be used as boot volumes. For more
	// information, see Amazon EBS volume types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
	// Amazon Elastic Compute Cloud User Guide. Default: gp2
	VolumeType types.VolumeType

	noSmithyDocumentSerde
}

// Describes a volume.
type CreateVolumeOutput struct {

	// Information about the volume attachments.
	Attachments []types.VolumeAttachment

	// The Availability Zone for the volume.
	AvailabilityZone *string

	// The time stamp when volume creation was initiated.
	CreateTime *time.Time

	// Indicates whether the volume is encrypted.
	Encrypted *bool

	// Indicates whether the volume was created using fast snapshot restore.
	FastRestored *bool

	// The number of I/O operations per second (IOPS). For gp3, io1, and io2 volumes,
	// this represents the number of IOPS that are provisioned for the volume. For gp2
	// volumes, this represents the baseline performance of the volume and the rate at
	// which the volume accumulates I/O credits for bursting.
	Iops *int32

	// The Amazon Resource Name (ARN) of the Key Management Service (KMS) KMS key that
	// was used to protect the volume encryption key for the volume.
	KmsKeyId *string

	// Indicates whether Amazon EBS Multi-Attach is enabled.
	MultiAttachEnabled *bool

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The size of the volume, in GiBs.
	Size *int32

	// The snapshot from which the volume was created, if applicable.
	SnapshotId *string

	// The volume state.
	State types.VolumeState

	// Any tags assigned to the volume.
	Tags []types.Tag

	// The throughput that the volume supports, in MiB/s.
	Throughput *int32

	// The ID of the volume.
	VolumeId *string

	// The volume type.
	VolumeType types.VolumeType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVolume{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateVolumeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVolume(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateVolume struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVolume) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVolume) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVolumeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVolumeInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateVolumeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVolume{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVolume",
	}
}
