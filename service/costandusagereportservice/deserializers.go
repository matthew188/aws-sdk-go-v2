// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/matthew188/aws-sdk-go-v2/service/costandusagereportservice/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
)

type awsAwsjson11_deserializeOpDeleteReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpDeleteReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDeleteReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteReportDefinition(response, &metadata)
	}
	output := &DeleteReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson11_deserializeOpDocumentDeleteReportDefinitionOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorDeleteReportDefinition(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpDescribeReportDefinitions struct {
}

func (*awsAwsjson11_deserializeOpDescribeReportDefinitions) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDescribeReportDefinitions) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeReportDefinitions(response, &metadata)
	}
	output := &DescribeReportDefinitionsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson11_deserializeOpDocumentDescribeReportDefinitionsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorDescribeReportDefinitions(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpModifyReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpModifyReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpModifyReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorModifyReportDefinition(response, &metadata)
	}
	output := &ModifyReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson11_deserializeOpDocumentModifyReportDefinitionOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorModifyReportDefinition(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpPutReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpPutReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpPutReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson11_deserializeOpErrorPutReportDefinition(response, &metadata)
	}
	output := &PutReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson11_deserializeOpDocumentPutReportDefinitionOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorPutReportDefinition(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("DuplicateReportNameException", errorCode):
		return awsAwsjson11_deserializeErrorDuplicateReportNameException(response, errorBody)

	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ReportLimitReachedException", errorCode):
		return awsAwsjson11_deserializeErrorReportLimitReachedException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson11_deserializeErrorDuplicateReportNameException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.DuplicateReportNameException{}
	err := awsAwsjson11_deserializeDocumentDuplicateReportNameException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeErrorInternalErrorException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.InternalErrorException{}
	err := awsAwsjson11_deserializeDocumentInternalErrorException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeErrorReportLimitReachedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ReportLimitReachedException{}
	err := awsAwsjson11_deserializeDocumentReportLimitReachedException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ValidationException{}
	err := awsAwsjson11_deserializeDocumentValidationException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson11_deserializeDocumentAdditionalArtifactList(v *[]types.AdditionalArtifact, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.AdditionalArtifact
	if *v == nil {
		cv = []types.AdditionalArtifact{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.AdditionalArtifact
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected AdditionalArtifact to be of type string, got %T instead", value)
			}
			col = types.AdditionalArtifact(jtv)
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentDuplicateReportNameException(v **types.DuplicateReportNameException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.DuplicateReportNameException
	if *v == nil {
		sv = &types.DuplicateReportNameException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentInternalErrorException(v **types.InternalErrorException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalErrorException
	if *v == nil {
		sv = &types.InternalErrorException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentReportDefinition(v **types.ReportDefinition, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ReportDefinition
	if *v == nil {
		sv = &types.ReportDefinition{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AdditionalArtifacts":
			if err := awsAwsjson11_deserializeDocumentAdditionalArtifactList(&sv.AdditionalArtifacts, value); err != nil {
				return err
			}

		case "AdditionalSchemaElements":
			if err := awsAwsjson11_deserializeDocumentSchemaElementList(&sv.AdditionalSchemaElements, value); err != nil {
				return err
			}

		case "BillingViewArn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected BillingViewArn to be of type string, got %T instead", value)
				}
				sv.BillingViewArn = ptr.String(jtv)
			}

		case "Compression":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected CompressionFormat to be of type string, got %T instead", value)
				}
				sv.Compression = types.CompressionFormat(jtv)
			}

		case "Format":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ReportFormat to be of type string, got %T instead", value)
				}
				sv.Format = types.ReportFormat(jtv)
			}

		case "RefreshClosedReports":
			if value != nil {
				jtv, ok := value.(bool)
				if !ok {
					return fmt.Errorf("expected RefreshClosedReports to be of type *bool, got %T instead", value)
				}
				sv.RefreshClosedReports = ptr.Bool(jtv)
			}

		case "ReportName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ReportName to be of type string, got %T instead", value)
				}
				sv.ReportName = ptr.String(jtv)
			}

		case "ReportVersioning":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ReportVersioning to be of type string, got %T instead", value)
				}
				sv.ReportVersioning = types.ReportVersioning(jtv)
			}

		case "S3Bucket":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected S3Bucket to be of type string, got %T instead", value)
				}
				sv.S3Bucket = ptr.String(jtv)
			}

		case "S3Prefix":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected S3Prefix to be of type string, got %T instead", value)
				}
				sv.S3Prefix = ptr.String(jtv)
			}

		case "S3Region":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected AWSRegion to be of type string, got %T instead", value)
				}
				sv.S3Region = types.AWSRegion(jtv)
			}

		case "TimeUnit":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected TimeUnit to be of type string, got %T instead", value)
				}
				sv.TimeUnit = types.TimeUnit(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentReportDefinitionList(v *[]types.ReportDefinition, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.ReportDefinition
	if *v == nil {
		cv = []types.ReportDefinition{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.ReportDefinition
		destAddr := &col
		if err := awsAwsjson11_deserializeDocumentReportDefinition(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentReportLimitReachedException(v **types.ReportLimitReachedException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ReportLimitReachedException
	if *v == nil {
		sv = &types.ReportLimitReachedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentSchemaElementList(v *[]types.SchemaElement, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.SchemaElement
	if *v == nil {
		cv = []types.SchemaElement{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.SchemaElement
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected SchemaElement to be of type string, got %T instead", value)
			}
			col = types.SchemaElement(jtv)
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentDeleteReportDefinitionOutput(v **DeleteReportDefinitionOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *DeleteReportDefinitionOutput
	if *v == nil {
		sv = &DeleteReportDefinitionOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ResponseMessage":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected DeleteResponseMessage to be of type string, got %T instead", value)
				}
				sv.ResponseMessage = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentDescribeReportDefinitionsOutput(v **DescribeReportDefinitionsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *DescribeReportDefinitionsOutput
	if *v == nil {
		sv = &DescribeReportDefinitionsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		case "ReportDefinitions":
			if err := awsAwsjson11_deserializeDocumentReportDefinitionList(&sv.ReportDefinitions, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentModifyReportDefinitionOutput(v **ModifyReportDefinitionOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *ModifyReportDefinitionOutput
	if *v == nil {
		sv = &ModifyReportDefinitionOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentPutReportDefinitionOutput(v **PutReportDefinitionOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *PutReportDefinitionOutput
	if *v == nil {
		sv = &PutReportDefinitionOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
