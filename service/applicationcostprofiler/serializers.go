// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationcostprofiler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/applicationcostprofiler/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteReportDefinition struct {
}

func (*awsRestjson1_serializeOpDeleteReportDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteReportDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/reportDefinition/{reportId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteReportDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteReportDefinitionInput(v *DeleteReportDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ReportId == nil || len(*v.ReportId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member reportId must not be empty")}
	}
	if v.ReportId != nil {
		if err := encoder.SetURI("reportId").String(*v.ReportId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetReportDefinition struct {
}

func (*awsRestjson1_serializeOpGetReportDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetReportDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/reportDefinition/{reportId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetReportDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetReportDefinitionInput(v *GetReportDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ReportId == nil || len(*v.ReportId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member reportId must not be empty")}
	}
	if v.ReportId != nil {
		if err := encoder.SetURI("reportId").String(*v.ReportId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpImportApplicationUsage struct {
}

func (*awsRestjson1_serializeOpImportApplicationUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpImportApplicationUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ImportApplicationUsageInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/importApplicationUsage")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentImportApplicationUsageInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsImportApplicationUsageInput(v *ImportApplicationUsageInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentImportApplicationUsageInput(v *ImportApplicationUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SourceS3Location != nil {
		ok := object.Key("sourceS3Location")
		if err := awsRestjson1_serializeDocumentSourceS3Location(v.SourceS3Location, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListReportDefinitions struct {
}

func (*awsRestjson1_serializeOpListReportDefinitions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListReportDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListReportDefinitionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/reportDefinition")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListReportDefinitionsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListReportDefinitionsInput(v *ListReportDefinitionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpPutReportDefinition struct {
}

func (*awsRestjson1_serializeOpPutReportDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutReportDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/reportDefinition")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutReportDefinitionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsPutReportDefinitionInput(v *PutReportDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutReportDefinitionInput(v *PutReportDefinitionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DestinationS3Location != nil {
		ok := object.Key("destinationS3Location")
		if err := awsRestjson1_serializeDocumentS3Location(v.DestinationS3Location, ok); err != nil {
			return err
		}
	}

	if len(v.Format) > 0 {
		ok := object.Key("format")
		ok.String(string(v.Format))
	}

	if v.ReportDescription != nil {
		ok := object.Key("reportDescription")
		ok.String(*v.ReportDescription)
	}

	if len(v.ReportFrequency) > 0 {
		ok := object.Key("reportFrequency")
		ok.String(string(v.ReportFrequency))
	}

	if v.ReportId != nil {
		ok := object.Key("reportId")
		ok.String(*v.ReportId)
	}

	return nil
}

type awsRestjson1_serializeOpUpdateReportDefinition struct {
}

func (*awsRestjson1_serializeOpUpdateReportDefinition) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateReportDefinitionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/reportDefinition/{reportId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateReportDefinitionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateReportDefinitionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateReportDefinitionInput(v *UpdateReportDefinitionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ReportId == nil || len(*v.ReportId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member reportId must not be empty")}
	}
	if v.ReportId != nil {
		if err := encoder.SetURI("reportId").String(*v.ReportId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateReportDefinitionInput(v *UpdateReportDefinitionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DestinationS3Location != nil {
		ok := object.Key("destinationS3Location")
		if err := awsRestjson1_serializeDocumentS3Location(v.DestinationS3Location, ok); err != nil {
			return err
		}
	}

	if len(v.Format) > 0 {
		ok := object.Key("format")
		ok.String(string(v.Format))
	}

	if v.ReportDescription != nil {
		ok := object.Key("reportDescription")
		ok.String(*v.ReportDescription)
	}

	if len(v.ReportFrequency) > 0 {
		ok := object.Key("reportFrequency")
		ok.String(string(v.ReportFrequency))
	}

	return nil
}

func awsRestjson1_serializeDocumentS3Location(v *types.S3Location, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Bucket != nil {
		ok := object.Key("bucket")
		ok.String(*v.Bucket)
	}

	if v.Prefix != nil {
		ok := object.Key("prefix")
		ok.String(*v.Prefix)
	}

	return nil
}

func awsRestjson1_serializeDocumentSourceS3Location(v *types.SourceS3Location, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Bucket != nil {
		ok := object.Key("bucket")
		ok.String(*v.Bucket)
	}

	if v.Key != nil {
		ok := object.Key("key")
		ok.String(*v.Key)
	}

	if len(v.Region) > 0 {
		ok := object.Key("region")
		ok.String(string(v.Region))
	}

	return nil
}
