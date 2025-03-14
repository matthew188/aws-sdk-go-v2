// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata, such as the authentication information about the Hadoop
// Distributed File System (HDFS) location.
func (c *Client) DescribeLocationHdfs(ctx context.Context, params *DescribeLocationHdfsInput, optFns ...func(*Options)) (*DescribeLocationHdfsOutput, error) {
	if params == nil {
		params = &DescribeLocationHdfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationHdfs", params, optFns, c.addOperationDescribeLocationHdfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationHdfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocationHdfsInput struct {

	// The Amazon Resource Name (ARN) of the HDFS cluster location to describe.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

type DescribeLocationHdfsOutput struct {

	// The ARNs of the agents that are used to connect to the HDFS cluster.
	AgentArns []string

	// The type of authentication used to determine the identity of the user.
	AuthenticationType types.HdfsAuthenticationType

	// The size of the data blocks to write into the HDFS cluster.
	BlockSize *int32

	// The time that the HDFS location was created.
	CreationTime *time.Time

	// The Kerberos principal with access to the files and folders on the HDFS cluster.
	// This parameter is used if the AuthenticationType is defined as KERBEROS.
	KerberosPrincipal *string

	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string

	// The ARN of the HDFS cluster location.
	LocationArn *string

	// The URI of the HDFS cluster location.
	LocationUri *string

	// The NameNode that manage the HDFS namespace.
	NameNodes []types.HdfsNameNode

	// The Quality of Protection (QOP) configuration specifies the Remote Procedure
	// Call (RPC) and data transfer protection settings configured on the Hadoop
	// Distributed File System (HDFS) cluster.
	QopConfiguration *types.QopConfiguration

	// The number of DataNodes to replicate the data to when writing to the HDFS
	// cluster.
	ReplicationFactor *int32

	// The user name used to identify the client on the host operating system. This
	// parameter is used if the AuthenticationType is defined as SIMPLE.
	SimpleUser *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationHdfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationHdfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationHdfs{}, middleware.After)
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
	if err = addOpDescribeLocationHdfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationHdfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationHdfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationHdfs",
	}
}
