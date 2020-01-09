// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: n13t/gitlab/v4/webhook-event.proto

package n13t_gitlab_v4

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
var _webhook_event_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on WebhookEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WebhookEvent) Validate() error {
	if m == nil {
		return nil
	}

	if !_WebhookEvent_ObjectKind_Pattern.MatchString(m.GetObjectKind()) {
		return WebhookEventValidationError{
			field:  "ObjectKind",
			reason: "value does not match regex pattern \"^(push|tag_push|issue|note|merge_request|wiki_page|pipeline|build)$\"",
		}
	}

	if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WebhookEventValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WebhookEventValidationError is the validation error returned by
// WebhookEvent.Validate if the designated constraints aren't met.
type WebhookEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WebhookEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WebhookEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WebhookEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WebhookEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WebhookEventValidationError) ErrorName() string { return "WebhookEventValidationError" }

// Error satisfies the builtin error interface
func (e WebhookEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWebhookEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WebhookEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WebhookEventValidationError{}

var _WebhookEvent_ObjectKind_Pattern = regexp.MustCompile("^(push|tag_push|issue|note|merge_request|wiki_page|pipeline|build)$")