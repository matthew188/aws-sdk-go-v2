// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes all versions of a solution and the Solution object itself. Before
// deleting a solution, you must delete all campaigns based on the solution. To
// determine what campaigns are using the solution, call ListCampaigns
// (https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html) and
// supply the Amazon Resource Name (ARN) of the solution. You can't delete a
// solution if an associated SolutionVersion is in the CREATE PENDING or IN
// PROGRESS state. For more information on solutions, see CreateSolution
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html).
func (c *Client) DeleteSolution(ctx context.Context, params *DeleteSolutionInput, optFns ...func(*Options)) (*DeleteSolutionOutput, error) {
	if params == nil {
		params = &DeleteSolutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSolution", params, optFns, c.addOperationDeleteSolutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSolutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSolutionInput struct {

	// The ARN of the solution to delete.
	//
	// This member is required.
	SolutionArn *string

	noSmithyDocumentSerde
}

type DeleteSolutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSolutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSolution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSolution{}, middleware.After)
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
	if err = addOpDeleteSolutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSolution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSolution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DeleteSolution",
	}
}
