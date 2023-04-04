// Code generated by smithy-go-codegen DO NOT EDIT.

package document

import (
	"bytes"
	"encoding/json"
	smithydocument "github.com/aws/smithy-go/document"
	smithydocumentjson "github.com/aws/smithy-go/document/json"
)

// github.com/matthew188/aws-sdk-go-v2/service/opensearchserverless/internal/document.smithyDocument
// is an interface which is used to bind a document type to its service client.
type smithyDocument interface {
	isSmithyDocument()
}

// github.com/matthew188/aws-sdk-go-v2/service/opensearchserverless/internal/document.Interface
// is a JSON-like data model type that is protocol agnostic and is usedto send
// open-content to a service.
type Interface interface {
	smithyDocument
	smithydocument.Marshaler
	smithydocument.Unmarshaler
}

type documentMarshaler struct {
	value interface{}
}

func (m *documentMarshaler) UnmarshalSmithyDocument(v interface{}) error {
	mBytes, err := m.MarshalSmithyDocument()
	if err != nil {
		return err
	}

	jDecoder := json.NewDecoder(bytes.NewReader(mBytes))
	jDecoder.UseNumber()

	var jv interface{}
	if err := jDecoder.Decode(&v); err != nil {
		return err
	}

	return NewDocumentUnmarshaler(v).UnmarshalSmithyDocument(&jv)
}

func (m *documentMarshaler) MarshalSmithyDocument() ([]byte, error) {
	return smithydocumentjson.NewEncoder().Encode(m.value)
}

func (m *documentMarshaler) isSmithyDocument() {}

var _ Interface = (*documentMarshaler)(nil)

type documentUnmarshaler struct {
	value interface{}
}

func (m *documentUnmarshaler) UnmarshalSmithyDocument(v interface{}) error {
	decoder := smithydocumentjson.NewDecoder()
	return decoder.DecodeJSONInterface(m.value, v)
}

func (m *documentUnmarshaler) MarshalSmithyDocument() ([]byte, error) {
	return json.Marshal(m.value)
}

func (m *documentUnmarshaler) isSmithyDocument() {}

var _ Interface = (*documentUnmarshaler)(nil)

// NewDocumentMarshaler creates a new document marshaler for the given input type
func NewDocumentMarshaler(v interface{}) Interface {
	return &documentMarshaler{
		value: v,
	}
}

// NewDocumentUnmarshaler creates a new document unmarshaler for the given service
// response
func NewDocumentUnmarshaler(v interface{}) Interface {
	return &documentUnmarshaler{
		value: v,
	}
}

// github.com/matthew188/aws-sdk-go-v2/service/opensearchserverless/internal/document.IsInterface
// returns whether the given Interface implementation is a valid client
// implementation
func IsInterface(v Interface) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	v.isSmithyDocument()
	return true
}
