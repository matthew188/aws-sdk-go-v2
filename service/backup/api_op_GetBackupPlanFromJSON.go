// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a valid JSON document specifying a backup plan or an error.
func (c *Client) GetBackupPlanFromJSON(ctx context.Context, params *GetBackupPlanFromJSONInput, optFns ...func(*Options)) (*GetBackupPlanFromJSONOutput, error) {
	if params == nil {
		params = &GetBackupPlanFromJSONInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBackupPlanFromJSON", params, optFns, c.addOperationGetBackupPlanFromJSONMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBackupPlanFromJSONOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBackupPlanFromJSONInput struct {

	// A customer-supplied backup plan document in JSON format.
	//
	// This member is required.
	BackupPlanTemplateJson *string

	noSmithyDocumentSerde
}

type GetBackupPlanFromJSONOutput struct {

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or more
	// sets of Rules.
	BackupPlan *types.BackupPlan

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBackupPlanFromJSONMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBackupPlanFromJSON{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBackupPlanFromJSON{}, middleware.After)
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
	if err = addOpGetBackupPlanFromJSONValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBackupPlanFromJSON(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBackupPlanFromJSON(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "GetBackupPlanFromJSON",
	}
}
