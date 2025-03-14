// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a resource group with the specified name and description. You can
// optionally include either a resource query or a service configuration. For more
// information about constructing a resource query, see Build queries and groups in
// Resource Groups
// (https://docs.aws.amazon.com/ARG/latest/userguide/getting_started-query.html) in
// the Resource Groups User Guide. For more information about service-linked groups
// and service configurations, see Service configurations for Resource Groups
// (https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html). Minimum
// permissions To run this command, you must have the following permissions:
//
// *
// resource-groups:CreateGroup
func (c *Client) CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) {
	if params == nil {
		params = &CreateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGroup", params, optFns, c.addOperationCreateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupInput struct {

	// The name of the group, which is the identifier of the group in other operations.
	// You can't change the name of a resource group after you create it. A resource
	// group name can consist of letters, numbers, hyphens, periods, and underscores.
	// The name cannot start with AWS or aws; these are reserved. A resource group name
	// must be unique within each Amazon Web Services Region in your Amazon Web
	// Services account.
	//
	// This member is required.
	Name *string

	// A configuration associates the resource group with an Amazon Web Services
	// service and specifies how the service can interact with the resources in the
	// group. A configuration is an array of GroupConfigurationItem elements. For
	// details about the syntax of service configurations, see Service configurations
	// for Resource Groups
	// (https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html). A resource
	// group can contain either a Configuration or a ResourceQuery, but not both.
	Configuration []types.GroupConfigurationItem

	// The description of the resource group. Descriptions can consist of letters,
	// numbers, hyphens, underscores, periods, and spaces.
	Description *string

	// The resource query that determines which Amazon Web Services resources are
	// members of this group. For more information about resource queries, see Create a
	// tag-based group in Resource Groups
	// (https://docs.aws.amazon.com/ARG/latest/userguide/gettingstarted-query.html#gettingstarted-query-cli-tag).
	// A resource group can contain either a ResourceQuery or a Configuration, but not
	// both.
	ResourceQuery *types.ResourceQuery

	// The tags to add to the group. A tag is key-value pair string.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateGroupOutput struct {

	// The description of the resource group.
	Group *types.Group

	// The service configuration associated with the resource group. For details about
	// the syntax of a service configuration, see Service configurations for Resource
	// Groups (https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html).
	GroupConfiguration *types.GroupConfiguration

	// The resource query associated with the group. For more information about
	// resource queries, see Create a tag-based group in Resource Groups
	// (https://docs.aws.amazon.com/ARG/latest/userguide/gettingstarted-query.html#gettingstarted-query-cli-tag).
	ResourceQuery *types.ResourceQuery

	// The tags associated with the group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGroup{}, middleware.After)
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
	if err = addOpCreateGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "CreateGroup",
	}
}
