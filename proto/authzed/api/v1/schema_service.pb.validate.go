// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/v1/schema_service.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ReadSchemaRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReadSchemaRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadSchemaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadSchemaRequestMultiError, or nil if none found.
func (m *ReadSchemaRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadSchemaRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReadSchemaRequestMultiError(errors)
	}

	return nil
}

// ReadSchemaRequestMultiError is an error wrapping multiple validation errors
// returned by ReadSchemaRequest.ValidateAll() if the designated constraints
// aren't met.
type ReadSchemaRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadSchemaRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadSchemaRequestMultiError) AllErrors() []error { return m }

// ReadSchemaRequestValidationError is the validation error returned by
// ReadSchemaRequest.Validate if the designated constraints aren't met.
type ReadSchemaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadSchemaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadSchemaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadSchemaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadSchemaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadSchemaRequestValidationError) ErrorName() string {
	return "ReadSchemaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadSchemaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadSchemaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadSchemaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadSchemaRequestValidationError{}

// Validate checks the field values on ReadSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadSchemaResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadSchemaResponseMultiError, or nil if none found.
func (m *ReadSchemaResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadSchemaResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SchemaText

	if len(errors) > 0 {
		return ReadSchemaResponseMultiError(errors)
	}

	return nil
}

// ReadSchemaResponseMultiError is an error wrapping multiple validation errors
// returned by ReadSchemaResponse.ValidateAll() if the designated constraints
// aren't met.
type ReadSchemaResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadSchemaResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadSchemaResponseMultiError) AllErrors() []error { return m }

// ReadSchemaResponseValidationError is the validation error returned by
// ReadSchemaResponse.Validate if the designated constraints aren't met.
type ReadSchemaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadSchemaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadSchemaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadSchemaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadSchemaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadSchemaResponseValidationError) ErrorName() string {
	return "ReadSchemaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadSchemaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadSchemaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadSchemaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadSchemaResponseValidationError{}

// Validate checks the field values on WriteSchemaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WriteSchemaRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteSchemaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WriteSchemaRequestMultiError, or nil if none found.
func (m *WriteSchemaRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteSchemaRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetSchema()) > 262144 {
		err := WriteSchemaRequestValidationError{
			field:  "Schema",
			reason: "value length must be at most 262144 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return WriteSchemaRequestMultiError(errors)
	}

	return nil
}

// WriteSchemaRequestMultiError is an error wrapping multiple validation errors
// returned by WriteSchemaRequest.ValidateAll() if the designated constraints
// aren't met.
type WriteSchemaRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteSchemaRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteSchemaRequestMultiError) AllErrors() []error { return m }

// WriteSchemaRequestValidationError is the validation error returned by
// WriteSchemaRequest.Validate if the designated constraints aren't met.
type WriteSchemaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteSchemaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteSchemaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteSchemaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteSchemaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteSchemaRequestValidationError) ErrorName() string {
	return "WriteSchemaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WriteSchemaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteSchemaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteSchemaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteSchemaRequestValidationError{}

// Validate checks the field values on WriteSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WriteSchemaResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WriteSchemaResponseMultiError, or nil if none found.
func (m *WriteSchemaResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteSchemaResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WriteSchemaResponseMultiError(errors)
	}

	return nil
}

// WriteSchemaResponseMultiError is an error wrapping multiple validation
// errors returned by WriteSchemaResponse.ValidateAll() if the designated
// constraints aren't met.
type WriteSchemaResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteSchemaResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteSchemaResponseMultiError) AllErrors() []error { return m }

// WriteSchemaResponseValidationError is the validation error returned by
// WriteSchemaResponse.Validate if the designated constraints aren't met.
type WriteSchemaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteSchemaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteSchemaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteSchemaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteSchemaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteSchemaResponseValidationError) ErrorName() string {
	return "WriteSchemaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WriteSchemaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteSchemaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteSchemaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteSchemaResponseValidationError{}
