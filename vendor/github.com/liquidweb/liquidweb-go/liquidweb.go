package liquidweb

import (
	"fmt"

	"github.com/liquidweb/liquidweb-go/types"
)

// Backend is an interface for calls against Liquid Web's API.
type Backend interface {
	Call(string, interface{}, interface{}) error
	CallRaw(string, interface{}) ([]byte, error)
}

// LWAPIRes is a convenient interface used (for example) by Call to ensure a passed
// struct knows how to indicate whether or not it had an error.
type LWAPIRes interface {
	Error() string
	HasError() bool
}

// ListMeta handles Liquid Web's pagination in HTTP responses.
type ListMeta struct {
	ItemCount types.FlexInt `json:"item_count,omitempty"`
	ItemTotal types.FlexInt `json:"item_total,omitempty"`
	PageNum   types.FlexInt `json:"page_num,omitempty"`
	PageSize  types.FlexInt `json:"page_size,omitempty"`
	PageTotal types.FlexInt `json:"page_total,omitempty"`
}

// PageParams support pagination parameters in parameter types.
type PageParams struct {
	PageNum  types.FlexInt `json:"page_num,omitempty"`
	PageSize types.FlexInt `json:"page_size,omitempty"`
}

// A LWAPIError is used to identify error responses when JSON unmarshalling json from a
// byte slice.
type LWAPIError struct {
	ErrorMsg     string `json:"error,omitempty"`
	ErrorClass   string `json:"error_class,omitempty"`
	ErrorFullMsg string `json:"full_message,omitempty"`
}

// Given a LWAPIError, returns a string containing the ErrorClass and ErrorFullMsg.
func (e LWAPIError) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorClass, e.ErrorFullMsg)
}

// HasError returns boolean if ErrorClass was present or not. You can
// use this function to determine if a LWAPIRes response indicates an error or not.
func (e LWAPIError) HasError() bool {
	return e.ErrorClass != ""
}
