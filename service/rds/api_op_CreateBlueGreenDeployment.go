// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a blue/green deployment. A blue/green deployment creates a staging
// environment that copies the production environment. In a blue/green deployment,
// the blue environment is the current production environment. The green
// environment is the staging environment. The staging environment stays in sync
// with the current production environment using logical replication. You can make
// changes to the databases in the green environment without affecting production
// workloads. For example, you can upgrade the major or minor DB engine version,
// change database parameters, or make schema changes in the staging environment.
// You can thoroughly test changes in the green environment. When ready, you can
// switch over the environments to promote the green environment to be the new
// production environment. The switchover typically takes under a minute. For more
// information, see Using Amazon RDS Blue/Green Deployments for database updates
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
// in the Amazon RDS User Guide and  Using Amazon RDS Blue/Green Deployments for
// database updates
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
// in the Amazon Aurora User Guide.
func (c *Client) CreateBlueGreenDeployment(ctx context.Context, params *CreateBlueGreenDeploymentInput, optFns ...func(*Options)) (*CreateBlueGreenDeploymentOutput, error) {
	if params == nil {
		params = &CreateBlueGreenDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBlueGreenDeployment", params, optFns, c.addOperationCreateBlueGreenDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBlueGreenDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBlueGreenDeploymentInput struct {

	// The name of the blue/green deployment. Constraints:
	//
	// * Can't be the same as an
	// existing blue/green deployment name in the same account and Amazon Web Services
	// Region.
	//
	// This member is required.
	BlueGreenDeploymentName *string

	// The Amazon Resource Name (ARN) of the source production database. Specify the
	// database that you want to clone. The blue/green deployment creates this database
	// in the green environment. You can make updates to the database in the green
	// environment, such as an engine version upgrade. When you are ready, you can
	// switch the database in the green environment to be the production database.
	//
	// This member is required.
	Source *string

	// Tags to assign to the blue/green deployment.
	Tags []types.Tag

	// The DB cluster parameter group associated with the Aurora DB cluster in the
	// green environment. To test parameter changes, specify a DB cluster parameter
	// group that is different from the one associated with the source DB cluster.
	TargetDBClusterParameterGroupName *string

	// The DB parameter group associated with the DB instance in the green environment.
	// To test parameter changes, specify a DB parameter group that is different from
	// the one associated with the source DB instance.
	TargetDBParameterGroupName *string

	// The engine version of the database in the green environment. Specify the engine
	// version to upgrade to in the green environment.
	TargetEngineVersion *string

	noSmithyDocumentSerde
}

type CreateBlueGreenDeploymentOutput struct {

	// Contains the details about a blue/green deployment. For more information, see
	// Using Amazon RDS Blue/Green Deployments for database updates
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
	// in the Amazon RDS User Guide and  Using Amazon RDS Blue/Green Deployments for
	// database updates
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
	// in the Amazon Aurora User Guide.
	BlueGreenDeployment *types.BlueGreenDeployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBlueGreenDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateBlueGreenDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateBlueGreenDeployment{}, middleware.After)
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
	if err = addOpCreateBlueGreenDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBlueGreenDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBlueGreenDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateBlueGreenDeployment",
	}
}
