// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates multiple ReplicationConfigurationTemplates by ID.
func (c *Client) UpdateReplicationConfigurationTemplate(ctx context.Context, params *UpdateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationTemplateOutput, error) {
	if params == nil {
		params = &UpdateReplicationConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationConfigurationTemplate", params, optFns, c.addOperationUpdateReplicationConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationConfigurationTemplateInput struct {

	// Update replication configuration template template ID request.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// Update replication configuration template ARN request.
	Arn *string

	// Update replication configuration template associate default Application
	// Migration Service Security group request.
	AssociateDefaultSecurityGroup *bool

	// Update replication configuration template bandwidth throttling request.
	BandwidthThrottling int64

	// Update replication configuration template create Public IP request.
	CreatePublicIP *bool

	// Update replication configuration template data plane routing request.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Update replication configuration template use default large Staging Disk type
	// request.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Update replication configuration template EBS encryption request.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Update replication configuration template EBS encryption key ARN request.
	EbsEncryptionKeyArn *string

	// Update replication configuration template Replication Server instance type
	// request.
	ReplicationServerInstanceType *string

	// Update replication configuration template Replication Server Security groups IDs
	// request.
	ReplicationServersSecurityGroupsIDs []string

	// Update replication configuration template Staging Area subnet ID request.
	StagingAreaSubnetId *string

	// Update replication configuration template Staging Area Tags request.
	StagingAreaTags map[string]string

	// Update replication configuration template use dedicated Replication Server
	// request.
	UseDedicatedReplicationServer *bool

	noSmithyDocumentSerde
}

type UpdateReplicationConfigurationTemplateOutput struct {

	// Replication Configuration template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// Replication Configuration template ARN.
	Arn *string

	// Replication Configuration template associate default Application Migration
	// Service Security group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration template bandwidth throttling.
	BandwidthThrottling int64

	// Replication Configuration template create Public IP.
	CreatePublicIP *bool

	// Replication Configuration template data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Replication Configuration template use default large Staging Disk type.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration template EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Replication Configuration template EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration template server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration template server Security Groups IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration template Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration template Staging Area Tags.
	StagingAreaTags map[string]string

	// Replication Configuration template Tags.
	Tags map[string]string

	// Replication Configuration template use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
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
	if err = addOpUpdateReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "UpdateReplicationConfigurationTemplate",
	}
}
