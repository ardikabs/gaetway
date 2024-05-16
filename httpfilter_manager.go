package gonvoy

import (
	"fmt"

	"github.com/ardikabs/gonvoy/pkg/errs"
	"github.com/ardikabs/gonvoy/pkg/util"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

// HttpFilterAction represents the action to be taken on each phase of an HTTP filter.
type HttpFilterAction uint

const (
	// ActionSkip is basically a no-op action, which mean the filter phase will be skipped.
	// From Envoy perspective, this is equivalent to a Continue status.
	ActionSkip HttpFilterAction = iota
	// ActionContinue is operation action that continues the current filter phase.
	// From Envoy perspective, this is equivalent to a Continue status.
	// Although the status is identical to ActionSkip, the purpose of ActionContinue
	// is to indicate that the current phase is operating correctly and needs to advance to the subsequent filter phase.
	ActionContinue
	// ActionPause is an operation action that pause the current filter phase.
	// This pause could be essential as future filter phases may rely on the processes of preceding phases.
	// For example, the EncodeData phase may require header alterations, which are established in the EncodeHeaders phase.
	// Therefore, EncodeHeaders should return with ActionPause, and the header changes may occur later in the EncodeData phase.
	// From Envoy perspective, this is equivalent to a StopAndBuffer status.
	// Be aware, when employing this action, the subsequent filter phase must return with ActionContinue,
	// otherwise the filter chain will be hanging.
	ActionPause
)

// HttpFilterDecoderFunc is a function type that represents a decoder for HTTP filters.
// It is used during the decoder phases, which is when the filter processes the incoming request.
type HttpFilterDecoderFunc func(Context, HttpFilterDecodeProcessor) (HttpFilterAction, error)

// HttpFilterEncoderFunc is a function type that represents an encoder for HTTP filters.
// It is used during the encoder phases, which is when the filter processes the upstream response.
type HttpFilterEncoderFunc func(Context, HttpFilterEncodeProcessor) (HttpFilterAction, error)

// HttpFilterServer represents an interface for an HTTP filter server.
// It provides methods to serve the decode and encode phases of an HTTP filter.
type HttpFilterServer interface {
	// ServeDecodeFilter serves decode phase of an HTTP filter.
	// This method is designed for internal use as it is used during the decoding phase only.
	// Decode phase is when the filter processes the incoming request, which consist of processing headers and body.
	//
	ServeDecodeFilter(HttpFilterDecoderFunc) *HttpFilterResult

	// ServeEncodeFilter serves encode phase of an HTTP filter.
	// This method is designed for internal use as it is used during the encoding phase only.
	// Encode phase is when the filter processes the upstream response, which consist of processing headers and body.
	//
	ServeEncodeFilter(HttpFilterEncoderFunc) *HttpFilterResult
}

// HttpFilterManager represents an interface for managing HTTP filters.
type HttpFilterManager interface {
	// SetErrorHandler sets a custom error handler for an HTTP Filter.
	//
	SetErrorHandler(ErrorHandler)

	// AddHTTPFilterHandler adds an HTTP Filter Handler to the chain,
	// which should be run during filter startup (HttpFilter.OnBegin).
	// It's important to note the order when adding filter handlers.
	// While HTTP requests follow FIFO sequences, HTTP responses follow LIFO sequences.
	//
	// Example usage:
	//	func (f *UserFilter) OnBegin(c RuntimeContext) error {
	//		...
	//		c.AddHTTPFilterHandler(handlerA)
	//		c.AddHTTPFilterHandler(handlerB)
	//		c.AddHTTPFilterHandler(handlerC)
	//		c.AddHTTPFilterHandler(handlerD)
	//	}
	//
	// During HTTP requests, traffic flows from `handlerA -> handlerB -> handlerC -> handlerD`.
	// During HTTP responses, traffic flows in reverse: `handlerD -> handlerC -> handlerB -> handlerA`.
	//
	AddHTTPFilterHandler(HttpFilterHandler)
}

func newHttpFilterManager(c Context) *httpFilterManager {
	return &httpFilterManager{
		ctx:          c,
		errorHandler: DefaultErrorHandler,
	}
}

type httpFilterManager struct {
	ctx          Context
	errorHandler ErrorHandler
	first        HttpFilterProcessor
	last         HttpFilterProcessor
}

func (h *httpFilterManager) SetErrorHandler(handler ErrorHandler) {
	if util.IsNil(handler) {
		return
	}

	h.errorHandler = handler
}

func (h *httpFilterManager) AddHTTPFilterHandler(handler HttpFilterHandler) {
	if util.IsNil(handler) || handler.Disable() {
		return
	}

	proc := newHttpFilterProcessor(handler)
	if h.first == nil {
		h.first = proc
		h.last = proc
		return
	}

	proc.SetPrevious(h.last)
	h.last.SetNext(proc)
	h.last = proc
}

func (h *httpFilterManager) ServeDecodeFilter(decoder HttpFilterDecoderFunc) (res *HttpFilterResult) {
	res = initHttpFilterResult(h.ctx, h.errorHandler)
	defer res.Finalize()
	if h.first == nil {
		return
	}

	res.Action, res.Err = decoder(h.ctx, h.first)
	return
}

func (h *httpFilterManager) ServeEncodeFilter(encoder HttpFilterEncoderFunc) (res *HttpFilterResult) {
	res = initHttpFilterResult(h.ctx, h.errorHandler)
	defer res.Finalize()
	if h.last == nil {
		return
	}

	res.Action, res.Err = encoder(h.ctx, h.last)
	return
}

func initHttpFilterResult(c Context, errorHandler ErrorHandler) *HttpFilterResult {
	return &HttpFilterResult{
		ctx:          c,
		errorHandler: errorHandler,
		Action:       ActionSkip,
		Status:       api.Continue,
	}
}

// HttpFilterResult represents the result of an HTTP filter operation.
type HttpFilterResult struct {
	ctx          Context
	errorHandler ErrorHandler

	// Action represents the action to be taken based on the filter result.
	Action HttpFilterAction
	// Err represents the error encountered during the filter operation, if any.
	Err error
	// Status represents the final status of the filter processing.
	Status api.StatusType
}

// Finalize performs the finalization logic for the HttpFilterResult.
// It recovers from any panics, sets the appropriate error and status,
// and determines the action to be taken based on the result.
func (res *HttpFilterResult) Finalize() {
	if r := recover(); r != nil {
		res.Err = fmt.Errorf("%w; %v", errs.ErrPanic, r)
	}

	if res.Err != nil {
		res.Status = res.errorHandler(res.ctx, res.Err)
		return
	}

	switch res.Action {
	case ActionContinue:
		res.Status = res.ctx.StatusType()
	case ActionPause:
		res.Status = api.StopAndBuffer
	default:
		res.Status = api.Continue
	}
}
