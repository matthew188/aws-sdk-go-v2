// Code generated by smithy-go-codegen DO NOT EDIT.

package qldbsession

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/qldbsession/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson10_serializeOpSendCommand struct {
}

func (*awsAwsjson10_serializeOpSendCommand) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpSendCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendCommandInput)
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
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("QLDBSession.SendCommand")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentSendCommandInput(input, jsonEncoder.Value); err != nil {
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
func awsAwsjson10_serializeDocumentAbortTransactionRequest(v *types.AbortTransactionRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson10_serializeDocumentCommitTransactionRequest(v *types.CommitTransactionRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CommitDigest != nil {
		ok := object.Key("CommitDigest")
		ok.Base64EncodeBytes(v.CommitDigest)
	}

	if v.TransactionId != nil {
		ok := object.Key("TransactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

func awsAwsjson10_serializeDocumentEndSessionRequest(v *types.EndSessionRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson10_serializeDocumentExecuteStatementRequest(v *types.ExecuteStatementRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Parameters != nil {
		ok := object.Key("Parameters")
		if err := awsAwsjson10_serializeDocumentStatementParameters(v.Parameters, ok); err != nil {
			return err
		}
	}

	if v.Statement != nil {
		ok := object.Key("Statement")
		ok.String(*v.Statement)
	}

	if v.TransactionId != nil {
		ok := object.Key("TransactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

func awsAwsjson10_serializeDocumentFetchPageRequest(v *types.FetchPageRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.NextPageToken != nil {
		ok := object.Key("NextPageToken")
		ok.String(*v.NextPageToken)
	}

	if v.TransactionId != nil {
		ok := object.Key("TransactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

func awsAwsjson10_serializeDocumentStartSessionRequest(v *types.StartSessionRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.LedgerName != nil {
		ok := object.Key("LedgerName")
		ok.String(*v.LedgerName)
	}

	return nil
}

func awsAwsjson10_serializeDocumentStartTransactionRequest(v *types.StartTransactionRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson10_serializeDocumentStatementParameters(v []types.ValueHolder, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentValueHolder(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeDocumentValueHolder(v *types.ValueHolder, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.IonBinary != nil {
		ok := object.Key("IonBinary")
		ok.Base64EncodeBytes(v.IonBinary)
	}

	if v.IonText != nil {
		ok := object.Key("IonText")
		ok.String(*v.IonText)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentSendCommandInput(v *SendCommandInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AbortTransaction != nil {
		ok := object.Key("AbortTransaction")
		if err := awsAwsjson10_serializeDocumentAbortTransactionRequest(v.AbortTransaction, ok); err != nil {
			return err
		}
	}

	if v.CommitTransaction != nil {
		ok := object.Key("CommitTransaction")
		if err := awsAwsjson10_serializeDocumentCommitTransactionRequest(v.CommitTransaction, ok); err != nil {
			return err
		}
	}

	if v.EndSession != nil {
		ok := object.Key("EndSession")
		if err := awsAwsjson10_serializeDocumentEndSessionRequest(v.EndSession, ok); err != nil {
			return err
		}
	}

	if v.ExecuteStatement != nil {
		ok := object.Key("ExecuteStatement")
		if err := awsAwsjson10_serializeDocumentExecuteStatementRequest(v.ExecuteStatement, ok); err != nil {
			return err
		}
	}

	if v.FetchPage != nil {
		ok := object.Key("FetchPage")
		if err := awsAwsjson10_serializeDocumentFetchPageRequest(v.FetchPage, ok); err != nil {
			return err
		}
	}

	if v.SessionToken != nil {
		ok := object.Key("SessionToken")
		ok.String(*v.SessionToken)
	}

	if v.StartSession != nil {
		ok := object.Key("StartSession")
		if err := awsAwsjson10_serializeDocumentStartSessionRequest(v.StartSession, ok); err != nil {
			return err
		}
	}

	if v.StartTransaction != nil {
		ok := object.Key("StartTransaction")
		if err := awsAwsjson10_serializeDocumentStartTransactionRequest(v.StartTransaction, ok); err != nil {
			return err
		}
	}

	return nil
}
