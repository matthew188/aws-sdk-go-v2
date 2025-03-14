// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a replication group's shards (node groups) by allowing you to add
// shards, remove shards, or rebalance the keyspaces among existing shards.
func (c *Client) ModifyReplicationGroupShardConfiguration(ctx context.Context, params *ModifyReplicationGroupShardConfigurationInput, optFns ...func(*Options)) (*ModifyReplicationGroupShardConfigurationOutput, error) {
	if params == nil {
		params = &ModifyReplicationGroupShardConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationGroupShardConfiguration", params, optFns, c.addOperationModifyReplicationGroupShardConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationGroupShardConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a ModifyReplicationGroupShardConfiguration operation.
type ModifyReplicationGroupShardConfigurationInput struct {

	// Indicates that the shard reconfiguration process begins immediately. At present,
	// the only permitted value for this parameter is true. Value: true
	//
	// This member is required.
	ApplyImmediately bool

	// The number of node groups (shards) that results from the modification of the
	// shard configuration.
	//
	// This member is required.
	NodeGroupCount int32

	// The name of the Redis (cluster mode enabled) cluster (replication group) on
	// which the shards are to be configured.
	//
	// This member is required.
	ReplicationGroupId *string

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by
	// NodeGroupsToRemove from the cluster.
	NodeGroupsToRemove []string

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRetain is a list of NodeGroupIds to retain in the cluster.
	// ElastiCache for Redis will attempt to remove all node groups except those listed
	// by NodeGroupsToRetain from the cluster.
	NodeGroupsToRetain []string

	// Specifies the preferred availability zones for each node group in the cluster.
	// If the value of NodeGroupCount is greater than the current number of node groups
	// (shards), you can use this parameter to specify the preferred availability zones
	// of the cluster's shards. If you omit this parameter ElastiCache selects
	// availability zones for you. You can specify this parameter only if the value of
	// NodeGroupCount is greater than the current number of node groups (shards).
	ReshardingConfiguration []types.ReshardingConfiguration

	noSmithyDocumentSerde
}

type ModifyReplicationGroupShardConfigurationOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyReplicationGroupShardConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyReplicationGroupShardConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyReplicationGroupShardConfiguration{}, middleware.After)
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
	if err = addOpModifyReplicationGroupShardConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationGroupShardConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReplicationGroupShardConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyReplicationGroupShardConfiguration",
	}
}
