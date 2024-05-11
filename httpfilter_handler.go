package gonvoy

import (
	"errors"
	"net/http"

	"github.com/ardikabs/gonvoy/pkg/errs"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

// HttpFilterHandler represents an interface for an HTTP filter handler.
// In a typical HTTP flow, the sequence of events can be as follows:
// OnRequestHeader -> OnRequestBody -> <Any number of intermediate Envoy processes> -> OnResponseHeader -> OnResponseBody
type HttpFilterHandler interface {
	// Disable disables the HTTP filter handler.
	//
	Disable() bool

	// OnRequestHeader is called when processing the HTTP request header during the OnRequestBodyPhase.
	//
	OnRequestHeader(c Context, header http.Header) error

	// OnRequestBody is called when processing the HTTP request body during the OnRequestBodyPhase.
	//
	OnRequestBody(c Context, body []byte) error

	// OnResponseHeader is called when processing the HTTP response header during the OnResponseHeaderPhase.
	//
	OnResponseHeader(c Context, header http.Header) error

	// OnResponseBody is called when processing the HTTP response body during the OnResponseBodyPhase.
	//
	OnResponseBody(c Context, body []byte) error
}

var (
	ResponseUnauthorized        = NewMinimalJSONResponse("UNAUTHORIZED", "UNAUTHORIZED")
	ResponseForbidden           = NewMinimalJSONResponse("FORBIDDEN", "FORBIDDEN")
	ResponseTooManyRequest      = NewMinimalJSONResponse("TOO_MANY_REQUEST", "TOO_MANY_REQUEST")
	ResponseInternalServerError = NewMinimalJSONResponse("RUNTIME_ERROR", "RUNTIME_ERROR")
	ResponseServiceUnavailable  = NewMinimalJSONResponse("SERVICE_UNAVAILABLE", "SERVICE_UNAVAILABLE")
)

type ErrorHandler func(Context, error) api.StatusType

func DefaultErrorHandler(ctx Context, err error) api.StatusType {
	unwrapErr := errs.Unwrap(err)
	if unwrapErr == nil {
		return api.Continue
	}

	switch unwrapErr {
	case errs.ErrUnauthorized:
		err = ctx.JSON(
			http.StatusUnauthorized,
			ResponseUnauthorized,
			NewGatewayHeaders(),
			WithResponseCodeDetails(DefaultResponseCodeDetailUnauthorized.Wrap(err.Error())))

	case errs.ErrAccessDenied:
		err = ctx.JSON(
			http.StatusForbidden,
			ResponseForbidden,
			NewGatewayHeaders(),
			WithResponseCodeDetails(DefaultResponseCodeDetailAccessDenied.Wrap(err.Error())))

	case errs.ErrOperationNotPermitted:
		err = ctx.JSON(
			http.StatusBadGateway,
			NewMinimalJSONResponse("BAD_GATEWAY", "BAD_GATEWAY", err),
			NewGatewayHeaders(),
			WithResponseCodeDetails(DefaultResponseCodeDetailError.Wrap(err.Error())))

	default:
		log := ctx.Log().WithCallDepth(3)
		if errors.Is(err, errs.ErrPanic) {
			log = log.WithCallDepth(1)
		}

		// hide internal error to end user
		// but printed out the error details to envoy log
		host := MustGetProperty(ctx, "request.host", "-")
		method := MustGetProperty(ctx, "request.method", "-")
		path := MustGetProperty(ctx, "request.path", "-")
		log.Error(err, "unidentified error", "host", host, "method", method, "path", path)
		err = ctx.JSON(
			http.StatusInternalServerError,
			ResponseInternalServerError,
			NewGatewayHeaders(),
			WithResponseCodeDetails(DefaultResponseCodeDetailError.Wrap(err.Error())))
	}

	if err != nil {
		// if we encounter another error, we will ignore the error
		// and allowing the request/response to proceed to the next Envoy filter.
		// Though, this condition is expected to be highly unlikely.
		return api.Continue
	}

	return api.LocalReply
}

var _ HttpFilterHandler = PassthroughHttpFilterHandler{}

type PassthroughHttpFilterHandler struct{}

func (PassthroughHttpFilterHandler) Disable() bool                                        { return false }
func (PassthroughHttpFilterHandler) OnRequestHeader(c Context, header http.Header) error  { return nil }
func (PassthroughHttpFilterHandler) OnRequestBody(c Context, body []byte) error           { return nil }
func (PassthroughHttpFilterHandler) OnResponseHeader(c Context, header http.Header) error { return nil }
func (PassthroughHttpFilterHandler) OnResponseBody(c Context, body []byte) error          { return nil }
