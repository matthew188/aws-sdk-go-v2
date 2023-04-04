// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a firewall rule group association, which enables DNS filtering for a
// VPC with one rule group. A VPC can have more than one firewall rule group
// association, and a rule group can be associated with more than one VPC.
func (c *Client) GetFirewallRuleGroupAssociation(ctx context.Context, params *GetFirewallRuleGroupAssociationInput, optFns ...func(*Options)) (*GetFirewallRuleGroupAssociationOutput, error) {
	if params == nil {
		params = &GetFirewallRuleGroupAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFirewallRuleGroupAssociation", params, optFns, c.addOperationGetFirewallRuleGroupAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFirewallRuleGroupAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFirewallRuleGroupAssociationInput struct {

	// The identifier of the FirewallRuleGroupAssociation.
	//
	// This member is required.
	FirewallRuleGroupAssociationId *string

	noSmithyDocumentSerde
}

type GetFirewallRuleGroupAssociationOutput struct {

	// The association that you requested.
	FirewallRuleGroupAssociation *types.FirewallRuleGroupAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFirewallRuleGroupAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFirewallRuleGroupAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFirewallRuleGroupAssociation{}, middleware.After)
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
	if err = addOpGetFirewallRuleGroupAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFirewallRuleGroupAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFirewallRuleGroupAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetFirewallRuleGroupAssociation",
	}
}
