// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes whether the Amazon Web Services account is opted in to cross-account
// backup. Returns an error if the account is not a member of an Organizations
// organization. Example: describe-global-settings --region us-west-2
func (c *Client) DescribeGlobalSettings(ctx context.Context, params *DescribeGlobalSettingsInput, optFns ...func(*Options)) (*DescribeGlobalSettingsOutput, error) {
	if params == nil {
		params = &DescribeGlobalSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalSettings", params, optFns, c.addOperationDescribeGlobalSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalSettingsInput struct {
	noSmithyDocumentSerde
}

type DescribeGlobalSettingsOutput struct {

	// The status of the flag isCrossAccountBackupEnabled.
	GlobalSettings map[string]string

	// The date and time that the flag isCrossAccountBackupEnabled was last updated.
	// This update is in Unix format and Coordinated Universal Time (UTC). The value of
	// LastUpdateTime is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	LastUpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGlobalSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGlobalSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGlobalSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGlobalSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeGlobalSettings",
	}
}
