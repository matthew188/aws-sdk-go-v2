// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new behavior graph for the calling account, and sets that account as
// the administrator account. This operation is called by the account that is
// enabling Detective. Before you try to enable Detective, make sure that your
// account has been enrolled in Amazon GuardDuty for at least 48 hours. If you do
// not meet this requirement, you cannot enable Detective. If you do meet the
// GuardDuty prerequisite, then when you make the request to enable Detective, it
// checks whether your data volume is within the Detective quota. If it exceeds the
// quota, then you cannot enable Detective. The operation also enables Detective
// for the calling account in the currently selected Region. It returns the ARN of
// the new behavior graph. CreateGraph triggers a process to create the
// corresponding data tables for the new behavior graph. An account can only be the
// administrator account for one behavior graph within a Region. If the same
// account calls CreateGraph with the same administrator account, it always returns
// the same behavior graph ARN. It does not create a new behavior graph.
func (c *Client) CreateGraph(ctx context.Context, params *CreateGraphInput, optFns ...func(*Options)) (*CreateGraphOutput, error) {
	if params == nil {
		params = &CreateGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGraph", params, optFns, c.addOperationCreateGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGraphInput struct {

	// The tags to assign to the new behavior graph. You can add up to 50 tags. For
	// each tag, you provide the tag key and the tag value. Each tag key can contain up
	// to 128 characters. Each tag value can contain up to 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateGraphOutput struct {

	// The ARN of the new behavior graph.
	GraphArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGraph{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGraph(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "CreateGraph",
	}
}
