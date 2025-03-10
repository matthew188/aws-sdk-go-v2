// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/cloud9/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpCreateEnvironmentEC2 struct {
}

func (*awsAwsjson11_serializeOpCreateEnvironmentEC2) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateEnvironmentEC2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateEnvironmentEC2Input)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.CreateEnvironmentEC2")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCreateEnvironmentEC2Input(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpCreateEnvironmentMembership struct {
}

func (*awsAwsjson11_serializeOpCreateEnvironmentMembership) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateEnvironmentMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateEnvironmentMembershipInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.CreateEnvironmentMembership")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCreateEnvironmentMembershipInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeleteEnvironment struct {
}

func (*awsAwsjson11_serializeOpDeleteEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.DeleteEnvironment")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteEnvironmentInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeleteEnvironmentMembership struct {
}

func (*awsAwsjson11_serializeOpDeleteEnvironmentMembership) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteEnvironmentMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteEnvironmentMembershipInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.DeleteEnvironmentMembership")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteEnvironmentMembershipInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeEnvironmentMemberships struct {
}

func (*awsAwsjson11_serializeOpDescribeEnvironmentMemberships) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeEnvironmentMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeEnvironmentMembershipsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.DescribeEnvironmentMemberships")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeEnvironmentMembershipsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeEnvironments struct {
}

func (*awsAwsjson11_serializeOpDescribeEnvironments) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeEnvironmentsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.DescribeEnvironments")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeEnvironmentsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeEnvironmentStatus struct {
}

func (*awsAwsjson11_serializeOpDescribeEnvironmentStatus) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeEnvironmentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeEnvironmentStatusInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.DescribeEnvironmentStatus")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeEnvironmentStatusInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListEnvironments struct {
}

func (*awsAwsjson11_serializeOpListEnvironments) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListEnvironmentsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.ListEnvironments")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListEnvironmentsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListTagsForResource struct {
}

func (*awsAwsjson11_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpTagResource struct {
}

func (*awsAwsjson11_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUntagResource struct {
}

func (*awsAwsjson11_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUpdateEnvironment struct {
}

func (*awsAwsjson11_serializeOpUpdateEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUpdateEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.UpdateEnvironment")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUpdateEnvironmentInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUpdateEnvironmentMembership struct {
}

func (*awsAwsjson11_serializeOpUpdateEnvironmentMembership) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUpdateEnvironmentMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateEnvironmentMembershipInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSCloud9WorkspaceManagementService.UpdateEnvironmentMembership")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUpdateEnvironmentMembershipInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentBoundedEnvironmentIdList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentPermissionsList(v []types.Permissions, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsAwsjson11_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentCreateEnvironmentEC2Input(v *CreateEnvironmentEC2Input, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AutomaticStopTimeMinutes != nil {
		ok := object.Key("automaticStopTimeMinutes")
		ok.Integer(*v.AutomaticStopTimeMinutes)
	}

	if v.ClientRequestToken != nil {
		ok := object.Key("clientRequestToken")
		ok.String(*v.ClientRequestToken)
	}

	if len(v.ConnectionType) > 0 {
		ok := object.Key("connectionType")
		ok.String(string(v.ConnectionType))
	}

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if v.DryRun != nil {
		ok := object.Key("dryRun")
		ok.Boolean(*v.DryRun)
	}

	if v.ImageId != nil {
		ok := object.Key("imageId")
		ok.String(*v.ImageId)
	}

	if v.InstanceType != nil {
		ok := object.Key("instanceType")
		ok.String(*v.InstanceType)
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.OwnerArn != nil {
		ok := object.Key("ownerArn")
		ok.String(*v.OwnerArn)
	}

	if v.SubnetId != nil {
		ok := object.Key("subnetId")
		ok.String(*v.SubnetId)
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsAwsjson11_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentCreateEnvironmentMembershipInput(v *CreateEnvironmentMembershipInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	if len(v.Permissions) > 0 {
		ok := object.Key("permissions")
		ok.String(string(v.Permissions))
	}

	if v.UserArn != nil {
		ok := object.Key("userArn")
		ok.String(*v.UserArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteEnvironmentInput(v *DeleteEnvironmentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteEnvironmentMembershipInput(v *DeleteEnvironmentMembershipInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	if v.UserArn != nil {
		ok := object.Key("userArn")
		ok.String(*v.UserArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeEnvironmentMembershipsInput(v *DescribeEnvironmentMembershipsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.Permissions != nil {
		ok := object.Key("permissions")
		if err := awsAwsjson11_serializeDocumentPermissionsList(v.Permissions, ok); err != nil {
			return err
		}
	}

	if v.UserArn != nil {
		ok := object.Key("userArn")
		ok.String(*v.UserArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeEnvironmentsInput(v *DescribeEnvironmentsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentIds != nil {
		ok := object.Key("environmentIds")
		if err := awsAwsjson11_serializeDocumentBoundedEnvironmentIdList(v.EnvironmentIds, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeEnvironmentStatusInput(v *DescribeEnvironmentStatusInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListEnvironmentsInput(v *ListEnvironmentsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson11_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	if v.TagKeys != nil {
		ok := object.Key("TagKeys")
		if err := awsAwsjson11_serializeDocumentTagKeyList(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUpdateEnvironmentInput(v *UpdateEnvironmentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	if len(v.ManagedCredentialsAction) > 0 {
		ok := object.Key("managedCredentialsAction")
		ok.String(string(v.ManagedCredentialsAction))
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUpdateEnvironmentMembershipInput(v *UpdateEnvironmentMembershipInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EnvironmentId != nil {
		ok := object.Key("environmentId")
		ok.String(*v.EnvironmentId)
	}

	if len(v.Permissions) > 0 {
		ok := object.Key("permissions")
		ok.String(string(v.Permissions))
	}

	if v.UserArn != nil {
		ok := object.Key("userArn")
		ok.String(*v.UserArn)
	}

	return nil
}
