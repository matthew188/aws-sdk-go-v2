// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a DB security group. The specified DB security group must not be
// associated with any DB instances. EC2-Classic was retired on August 15, 2022. If
// you haven't migrated from EC2-Classic to a VPC, we recommend that you migrate as
// soon as possible. For more information, see Migrate from EC2-Classic to a VPC
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in the
// Amazon EC2 User Guide, the blog EC2-Classic Networking is Retiring – Here’s How
// to Prepare
// (http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/),
// and Moving a DB instance not in a VPC into a VPC
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html)
// in the Amazon RDS User Guide.
func (c *Client) DeleteDBSecurityGroup(ctx context.Context, params *DeleteDBSecurityGroupInput, optFns ...func(*Options)) (*DeleteDBSecurityGroupOutput, error) {
	if params == nil {
		params = &DeleteDBSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBSecurityGroup", params, optFns, c.addOperationDeleteDBSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBSecurityGroupInput struct {

	// The name of the DB security group to delete. You can't delete the default DB
	// security group. Constraints:
	//
	// * Must be 1 to 255 letters, numbers, or
	// hyphens.
	//
	// * First character must be a letter
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens
	//
	// * Must not be "Default"
	//
	// This member is required.
	DBSecurityGroupName *string

	noSmithyDocumentSerde
}

type DeleteDBSecurityGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSecurityGroup{}, middleware.After)
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
	if err = addOpDeleteDBSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSecurityGroup",
	}
}
