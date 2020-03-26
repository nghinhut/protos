// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: n13t/idm/v0/idm.proto

package n13t_idm_v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _idm_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	return nil
}

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// HealthCheckResponseValidationError is the validation error returned by
// HealthCheckResponse.Validate if the designated constraints aren't met.
type HealthCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckResponseValidationError) ErrorName() string {
	return "HealthCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckResponseValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Search

	return nil
}

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on SumRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SumRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	return nil
}

// SumRequestValidationError is the validation error returned by
// SumRequest.Validate if the designated constraints aren't met.
type SumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumRequestValidationError) ErrorName() string { return "SumRequestValidationError" }

// Error satisfies the builtin error interface
func (e SumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumRequestValidationError{}

// Validate checks the field values on SumReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SumReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for V

	// no validation rules for Err

	return nil
}

// SumReplyValidationError is the validation error returned by
// SumReply.Validate if the designated constraints aren't met.
type SumReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumReplyValidationError) ErrorName() string { return "SumReplyValidationError" }

// Error satisfies the builtin error interface
func (e SumReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumReplyValidationError{}

// Validate checks the field values on ConcatRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConcatRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	return nil
}

// ConcatRequestValidationError is the validation error returned by
// ConcatRequest.Validate if the designated constraints aren't met.
type ConcatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcatRequestValidationError) ErrorName() string { return "ConcatRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConcatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcatRequestValidationError{}

// Validate checks the field values on ConcatReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConcatReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for V

	// no validation rules for Err

	return nil
}

// ConcatReplyValidationError is the validation error returned by
// ConcatReply.Validate if the designated constraints aren't met.
type ConcatReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcatReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcatReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcatReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcatReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcatReplyValidationError) ErrorName() string { return "ConcatReplyValidationError" }

// Error satisfies the builtin error interface
func (e ConcatReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcatReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcatReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcatReplyValidationError{}

// Validate checks the field values on IsValidRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *IsValidRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Password

	return nil
}

// IsValidRequestValidationError is the validation error returned by
// IsValidRequest.Validate if the designated constraints aren't met.
type IsValidRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsValidRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsValidRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsValidRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsValidRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsValidRequestValidationError) ErrorName() string { return "IsValidRequestValidationError" }

// Error satisfies the builtin error interface
func (e IsValidRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsValidRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsValidRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsValidRequestValidationError{}

// Validate checks the field values on IsValidResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *IsValidResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for V

	// no validation rules for Err

	return nil
}

// IsValidResponseValidationError is the validation error returned by
// IsValidResponse.Validate if the designated constraints aren't met.
type IsValidResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsValidResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsValidResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsValidResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsValidResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsValidResponseValidationError) ErrorName() string { return "IsValidResponseValidationError" }

// Error satisfies the builtin error interface
func (e IsValidResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsValidResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsValidResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsValidResponseValidationError{}
