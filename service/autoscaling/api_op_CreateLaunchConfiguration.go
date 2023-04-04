// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch configuration. If you exceed your maximum limit of launch
// configurations, the call fails. To query this limit, call the
// DescribeAccountLimits API. For information about updating this limit, see Quotas
// for Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Launch
// configurations
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchConfiguration.html)
// in the Amazon EC2 Auto Scaling User Guide. Amazon EC2 Auto Scaling configures
// instances launched as part of an Auto Scaling group using either a launch
// template or a launch configuration. We strongly recommend that you do not use
// launch configurations. They do not provide full functionality for Amazon EC2
// Auto Scaling or Amazon EC2. For information about using launch templates, see
// Launch templates
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-templates.html) in
// the Amazon EC2 Auto Scaling User Guide.
func (c *Client) CreateLaunchConfiguration(ctx context.Context, params *CreateLaunchConfigurationInput, optFns ...func(*Options)) (*CreateLaunchConfigurationOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfiguration", params, optFns, c.addOperationCreateLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationInput struct {

	// The name of the launch configuration. This name must be unique per Region per
	// account.
	//
	// This member is required.
	LaunchConfigurationName *string

	// Specifies whether to assign a public IPv4 address to the group's instances. If
	// the instance is launched into a default subnet, the default is to assign a
	// public IPv4 address, unless you disabled the option to assign a public IPv4
	// address on the subnet. If the instance is launched into a nondefault subnet, the
	// default is not to assign a public IPv4 address, unless you enabled the option to
	// assign a public IPv4 address on the subnet. If you specify true, each instance
	// in the Auto Scaling group receives a unique public IPv4 address. For more
	// information, see Launching Auto Scaling instances in a VPC
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the
	// Amazon EC2 Auto Scaling User Guide. If you specify this property, you must
	// specify at least one subnet for VPCZoneIdentifier when you create your group.
	AssociatePublicIpAddress *bool

	// The block device mapping entries that define the block devices to attach to the
	// instances at launch. By default, the block devices specified in the block device
	// mapping for the AMI are used. For more information, see Block device mappings
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	BlockDeviceMappings []types.BlockDeviceMapping

	// Available for backward compatibility.
	ClassicLinkVPCId *string

	// Available for backward compatibility.
	ClassicLinkVPCSecurityGroups []string

	// Specifies whether the launch configuration is optimized for EBS I/O (true) or
	// not (false). The optimization provides dedicated throughput to Amazon EBS and an
	// optimized configuration stack to provide optimal I/O performance. This
	// optimization is not available with all instance types. Additional fees are
	// incurred when you enable EBS optimization for an instance type that is not
	// EBS-optimized by default. For more information, see Amazon EBS-optimized
	// instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the
	// Amazon EC2 User Guide for Linux Instances. The default value is false.
	EbsOptimized *bool

	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM role.
	// For more information, see IAM role for applications that run on Amazon EC2
	// instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	IamInstanceProfile *string

	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	// For more information, see Finding a Linux AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the
	// Amazon EC2 User Guide for Linux Instances. If you specify InstanceId, an ImageId
	// is not required.
	ImageId *string

	// The ID of the instance to use to create the launch configuration. The new launch
	// configuration derives attributes from the instance, except for the block device
	// mapping. To create a launch configuration with a block device mapping or
	// override any other instance attributes, specify them as part of the same
	// request. For more information, see Creating a launch configuration using an EC2
	// instance
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-lc-with-instanceID.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceId *string

	// Controls whether instances in this group are launched with detailed (true) or
	// basic (false) monitoring. The default value is true (enabled). When detailed
	// monitoring is enabled, Amazon CloudWatch generates metrics every minute and your
	// account is charged a fee. When you disable detailed monitoring, CloudWatch
	// generates metrics every 5 minutes. For more information, see Configure
	// Monitoring for Auto Scaling Instances
	// (https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceMonitoring *types.InstanceMonitoring

	// Specifies the instance type of the EC2 instance. For information about available
	// instance types, see Available instance types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes)
	// in the Amazon EC2 User Guide for Linux Instances. If you specify InstanceId, an
	// InstanceType is not required.
	InstanceType *string

	// The ID of the kernel associated with the AMI. We recommend that you use PV-GRUB
	// instead of kernels and RAM disks. For more information, see User provided
	// kernels
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	KernelId *string

	// The name of the key pair. For more information, see Amazon EC2 key pairs and
	// Linux instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon EC2 User Guide for Linux Instances.
	KeyName *string

	// The metadata options for the instances. For more information, see Configuring
	// the Instance Metadata Options
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds)
	// in the Amazon EC2 Auto Scaling User Guide.
	MetadataOptions *types.InstanceMetadataOptions

	// The tenancy of the instance, either default or dedicated. An instance with
	// dedicated tenancy runs on isolated, single-tenant hardware and can only be
	// launched into a VPC. To launch dedicated instances into a shared tenancy VPC (a
	// VPC with the instance placement tenancy attribute set to default), you must set
	// the value of this property to dedicated. For more information, see Configuring
	// instance tenancy with Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide. If you specify PlacementTenancy, you
	// must specify at least one subnet for VPCZoneIdentifier when you create your
	// group. Valid values: default | dedicated
	PlacementTenancy *string

	// The ID of the RAM disk to select. We recommend that you use PV-GRUB instead of
	// kernels and RAM disks. For more information, see User provided kernels
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	RamdiskId *string

	// A list that contains the security group IDs to assign to the instances in the
	// Auto Scaling group. For more information, see Control traffic to resources using
	// security groups
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide.
	SecurityGroups []string

	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds the
	// current Spot price. For more information, see Request Spot Instances for
	// fault-tolerant and flexible applications
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-template-spot-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide. Valid Range: Minimum value of 0.001
	// When you change your maximum price by creating a new launch configuration,
	// running instances will continue to run as long as the maximum price for those
	// running instances is higher than the current Spot price.
	SpotPrice *string

	// The user data to make available to the launched EC2 instances. For more
	// information, see Instance metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// (Linux) and Instance metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html)
	// (Windows). If you are using a command line tool, base64-encoding is performed
	// for you, and you can load the text from a file. Otherwise, you must provide
	// base64-encoded text. User data is limited to 16 KB.
	UserData *string

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLaunchConfiguration{}, middleware.After)
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
	if err = addOpCreateLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateLaunchConfiguration",
	}
}
