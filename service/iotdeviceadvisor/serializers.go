// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateSuiteDefinition struct {
}

func (*awsRestjson1_serializeOpCreateSuiteDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateSuiteDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateSuiteDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateSuiteDefinitionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateSuiteDefinitionInput(v *CreateSuiteDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateSuiteDefinitionInput(v *CreateSuiteDefinitionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SuiteDefinitionConfiguration != nil {
		ok := object.Key("suiteDefinitionConfiguration")
		if err := awsRestjson1_serializeDocumentSuiteDefinitionConfiguration(v.SuiteDefinitionConfiguration, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteSuiteDefinition struct {
}

func (*awsRestjson1_serializeOpDeleteSuiteDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteSuiteDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteSuiteDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteSuiteDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteSuiteDefinitionInput(v *DeleteSuiteDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetEndpoint struct {
}

func (*awsRestjson1_serializeOpGetEndpoint) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetEndpointInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/endpoint")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetEndpointInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetEndpointInput(v *GetEndpointInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.CertificateArn != nil {
		encoder.SetQuery("certificateArn").String(*v.CertificateArn)
	}

	if v.ThingArn != nil {
		encoder.SetQuery("thingArn").String(*v.ThingArn)
	}

	return nil
}

type awsRestjson1_serializeOpGetSuiteDefinition struct {
}

func (*awsRestjson1_serializeOpGetSuiteDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSuiteDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSuiteDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSuiteDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSuiteDefinitionInput(v *GetSuiteDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	if v.SuiteDefinitionVersion != nil {
		encoder.SetQuery("suiteDefinitionVersion").String(*v.SuiteDefinitionVersion)
	}

	return nil
}

type awsRestjson1_serializeOpGetSuiteRun struct {
}

func (*awsRestjson1_serializeOpGetSuiteRun) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSuiteRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSuiteRunInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}/suiteRuns/{suiteRunId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSuiteRunInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSuiteRunInput(v *GetSuiteRunInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	if v.SuiteRunId == nil || len(*v.SuiteRunId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteRunId must not be empty")}
	}
	if v.SuiteRunId != nil {
		if err := encoder.SetURI("suiteRunId").String(*v.SuiteRunId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetSuiteRunReport struct {
}

func (*awsRestjson1_serializeOpGetSuiteRunReport) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSuiteRunReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSuiteRunReportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}/suiteRuns/{suiteRunId}/report")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSuiteRunReportInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSuiteRunReportInput(v *GetSuiteRunReportInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	if v.SuiteRunId == nil || len(*v.SuiteRunId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteRunId must not be empty")}
	}
	if v.SuiteRunId != nil {
		if err := encoder.SetURI("suiteRunId").String(*v.SuiteRunId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListSuiteDefinitions struct {
}

func (*awsRestjson1_serializeOpListSuiteDefinitions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSuiteDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSuiteDefinitionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSuiteDefinitionsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSuiteDefinitionsInput(v *ListSuiteDefinitionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != 0 {
		encoder.SetQuery("maxResults").Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListSuiteRuns struct {
}

func (*awsRestjson1_serializeOpListSuiteRuns) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSuiteRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSuiteRunsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteRuns")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSuiteRunsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSuiteRunsInput(v *ListSuiteRunsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != 0 {
		encoder.SetQuery("maxResults").Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.SuiteDefinitionId != nil {
		encoder.SetQuery("suiteDefinitionId").String(*v.SuiteDefinitionId)
	}

	if v.SuiteDefinitionVersion != nil {
		encoder.SetQuery("suiteDefinitionVersion").String(*v.SuiteDefinitionVersion)
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartSuiteRun struct {
}

func (*awsRestjson1_serializeOpStartSuiteRun) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartSuiteRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartSuiteRunInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}/suiteRuns")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStartSuiteRunInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartSuiteRunInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartSuiteRunInput(v *StartSuiteRunInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartSuiteRunInput(v *StartSuiteRunInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SuiteDefinitionVersion != nil {
		ok := object.Key("suiteDefinitionVersion")
		ok.String(*v.SuiteDefinitionVersion)
	}

	if v.SuiteRunConfiguration != nil {
		ok := object.Key("suiteRunConfiguration")
		if err := awsRestjson1_serializeDocumentSuiteRunConfiguration(v.SuiteRunConfiguration, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStopSuiteRun struct {
}

func (*awsRestjson1_serializeOpStopSuiteRun) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStopSuiteRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StopSuiteRunInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}/suiteRuns/{suiteRunId}/stop")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStopSuiteRunInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStopSuiteRunInput(v *StopSuiteRunInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	if v.SuiteRunId == nil || len(*v.SuiteRunId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteRunId must not be empty")}
	}
	if v.SuiteRunId != nil {
		if err := encoder.SetURI("suiteRunId").String(*v.SuiteRunId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsTagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUntagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		for i := range v.TagKeys {
			encoder.AddQuery("tagKeys").String(v.TagKeys[i])
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateSuiteDefinition struct {
}

func (*awsRestjson1_serializeOpUpdateSuiteDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateSuiteDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateSuiteDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/suiteDefinitions/{suiteDefinitionId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PATCH"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateSuiteDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateSuiteDefinitionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUpdateSuiteDefinitionInput(v *UpdateSuiteDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.SuiteDefinitionId == nil || len(*v.SuiteDefinitionId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member suiteDefinitionId must not be empty")}
	}
	if v.SuiteDefinitionId != nil {
		if err := encoder.SetURI("suiteDefinitionId").String(*v.SuiteDefinitionId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateSuiteDefinitionInput(v *UpdateSuiteDefinitionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SuiteDefinitionConfiguration != nil {
		ok := object.Key("suiteDefinitionConfiguration")
		if err := awsRestjson1_serializeDocumentSuiteDefinitionConfiguration(v.SuiteDefinitionConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentDeviceUnderTest(v *types.DeviceUnderTest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CertificateArn != nil {
		ok := object.Key("certificateArn")
		ok.String(*v.CertificateArn)
	}

	if v.ThingArn != nil {
		ok := object.Key("thingArn")
		ok.String(*v.ThingArn)
	}

	return nil
}

func awsRestjson1_serializeDocumentDeviceUnderTestList(v []types.DeviceUnderTest, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentDeviceUnderTest(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSelectedTestList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSuiteDefinitionConfiguration(v *types.SuiteDefinitionConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DevicePermissionRoleArn != nil {
		ok := object.Key("devicePermissionRoleArn")
		ok.String(*v.DevicePermissionRoleArn)
	}

	if v.Devices != nil {
		ok := object.Key("devices")
		if err := awsRestjson1_serializeDocumentDeviceUnderTestList(v.Devices, ok); err != nil {
			return err
		}
	}

	if v.IntendedForQualification {
		ok := object.Key("intendedForQualification")
		ok.Boolean(v.IntendedForQualification)
	}

	if v.IsLongDurationTest {
		ok := object.Key("isLongDurationTest")
		ok.Boolean(v.IsLongDurationTest)
	}

	if len(v.Protocol) > 0 {
		ok := object.Key("protocol")
		ok.String(string(v.Protocol))
	}

	if v.RootGroup != nil {
		ok := object.Key("rootGroup")
		ok.String(*v.RootGroup)
	}

	if v.SuiteDefinitionName != nil {
		ok := object.Key("suiteDefinitionName")
		ok.String(*v.SuiteDefinitionName)
	}

	return nil
}

func awsRestjson1_serializeDocumentSuiteRunConfiguration(v *types.SuiteRunConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ParallelRun {
		ok := object.Key("parallelRun")
		ok.Boolean(v.ParallelRun)
	}

	if v.PrimaryDevice != nil {
		ok := object.Key("primaryDevice")
		if err := awsRestjson1_serializeDocumentDeviceUnderTest(v.PrimaryDevice, ok); err != nil {
			return err
		}
	}

	if v.SelectedTestList != nil {
		ok := object.Key("selectedTestList")
		if err := awsRestjson1_serializeDocumentSelectedTestList(v.SelectedTestList, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}
