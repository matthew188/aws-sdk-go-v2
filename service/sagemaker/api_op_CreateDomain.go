// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Domain used by Amazon SageMaker Studio. A domain consists of an
// associated Amazon Elastic File System (EFS) volume, a list of authorized users,
// and a variety of security, application, policy, and Amazon Virtual Private Cloud
// (VPC) configurations. Users within a domain can share notebook files and other
// artifacts with each other. EFS storage When a domain is created, an EFS volume
// is created for use by all of the users within the domain. Each user receives a
// private home directory within the EFS volume for notebooks, Git repositories,
// and data files. SageMaker uses the Amazon Web Services Key Management Service
// (Amazon Web Services KMS) to encrypt the EFS volume attached to the domain with
// an Amazon Web Services managed key by default. For more control, you can specify
// a customer managed key. For more information, see Protect Data at Rest Using
// Encryption
// (https://docs.aws.amazon.com/sagemaker/latest/dg/encryption-at-rest.html). VPC
// configuration All SageMaker Studio traffic between the domain and the EFS volume
// is through the specified VPC and subnets. For other Studio traffic, you can
// specify the AppNetworkAccessType parameter. AppNetworkAccessType corresponds to
// the network access type that you choose when you onboard to Studio. The
// following options are available:
//
// * PublicInternetOnly - Non-EFS traffic goes
// through a VPC managed by Amazon SageMaker, which allows internet access. This is
// the default value.
//
// * VpcOnly - All Studio traffic is through the specified VPC
// and subnets. Internet access is disabled by default. To allow internet access,
// you must specify a NAT gateway. When internet access is disabled, you won't be
// able to run a Studio notebook or to train or host models unless your VPC has an
// interface endpoint to the SageMaker API and runtime or a NAT gateway and your
// security groups allow outbound connections.
//
// NFS traffic over TCP on port 2049
// needs to be allowed in both inbound and outbound rules in order to launch a
// SageMaker Studio app successfully. For more information, see Connect SageMaker
// Studio Notebooks to Resources in a VPC
// (https://docs.aws.amazon.com/sagemaker/latest/dg/studio-notebooks-and-internet-access.html).
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, c.addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The mode of authentication that members use to access the domain.
	//
	// This member is required.
	AuthMode types.AuthMode

	// The default settings to use to create a user profile when UserSettings isn't
	// specified in the call to the CreateUserProfile API. SecurityGroups is aggregated
	// when specified in both calls. For all other settings in UserSettings, the values
	// specified in CreateUserProfile take precedence over those specified in
	// CreateDomain.
	//
	// This member is required.
	DefaultUserSettings *types.UserSettings

	// A name for the domain.
	//
	// This member is required.
	DomainName *string

	// The VPC subnets that Studio uses for communication.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for
	// communication.
	//
	// This member is required.
	VpcId *string

	// Specifies the VPC used for non-EFS traffic. The default value is
	// PublicInternetOnly.
	//
	// * PublicInternetOnly - Non-EFS traffic is through a VPC
	// managed by Amazon SageMaker, which allows direct internet access
	//
	// * VpcOnly -
	// All Studio traffic is through the specified VPC and subnets
	AppNetworkAccessType types.AppNetworkAccessType

	// The entity that creates and manages the required security groups for inter-app
	// communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType
	// is VPCOnly and
	// DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is
	// provided.
	AppSecurityGroupManagement types.AppSecurityGroupManagement

	// The default settings used to create a space.
	DefaultSpaceSettings *types.DefaultSpaceSettings

	// A collection of Domain settings.
	DomainSettings *types.DomainSettings

	// Use KmsKeyId.
	//
	// Deprecated: This property is deprecated, use KmsKeyId instead.
	HomeEfsFileSystemKmsKeyId *string

	// SageMaker uses Amazon Web Services KMS to encrypt the EFS volume attached to the
	// domain with an Amazon Web Services managed key by default. For more control,
	// specify a customer managed key.
	KmsKeyId *string

	// Tags to associated with the Domain. Each tag consists of a key and an optional
	// value. Tag keys must be unique per resource. Tags are searchable using the
	// Search API. Tags that you specify for the Domain are also added to all Apps that
	// the Domain launches.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDomainOutput struct {

	// The Amazon Resource Name (ARN) of the created domain.
	DomainArn *string

	// The URL to the created domain.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDomain{}, middleware.After)
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
	if err = addOpCreateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateDomain",
	}
}
