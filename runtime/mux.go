package runtime

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/textproto"
	"regexp"
	"strings"

	"github.com/davidhong1/grpc-gateway/v2/internal/httprule"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// UnescapingMode defines the behavior of ServeMux when unescaping path parameters.
type UnescapingMode int

const (
	// UnescapingModeLegacy is the default V2 behavior, which escapes the entire
	// path string before doing any routing.
	UnescapingModeLegacy UnescapingMode = iota

	// UnescapingModeAllExceptReserved unescapes all path parameters except RFC 6570
	// reserved characters.
	UnescapingModeAllExceptReserved

	// UnescapingModeAllExceptSlash unescapes URL path parameters except path
	// separators, which will be left as "%2F".
	UnescapingModeAllExceptSlash

	// UnescapingModeAllCharacters unescapes all URL path parameters.
	UnescapingModeAllCharacters

	// UnescapingModeDefault is the default escaping type.
	// TODO(v3): default this to UnescapingModeAllExceptReserved per grpc-httpjson-transcoding's
	// reference implementation
	UnescapingModeDefault = UnescapingModeLegacy
)

var encodedPathSplitter = regexp.MustCompile("(/|%2F)")

// A HandlerFunc handles a specific pair of path pattern and HTTP method.
type HandlerFunc func(w http.ResponseWriter, r *http.Request, pathParams map[string]string)

// ServeMux is a request multiplexer for grpc-gateway.
// It matches http requests to patterns and invokes the corresponding handler.
type ServeMux struct {
	// handlers maps HTTP method to a list of handlers.
	handlers                  map[string][]handler
	forwardResponseOptions    []func(context.Context, http.ResponseWriter, proto.Message) error
	marshalers                marshalerRegistry
	incomingHeaderMatcher     HeaderMatcherFunc
	outgoingHeaderMatcher     HeaderMatcherFunc
	outgoingTrailerMatcher    HeaderMatcherFunc
	metadataAnnotators        []func(context.Context, *http.Request) metadata.MD
	errorHandler              ErrorHandlerFunc
	streamErrorHandler        StreamErrorHandlerFunc
	routingErrorHandler       RoutingErrorHandlerFunc
	disablePathLengthFallback bool
	unescapingMode            UnescapingMode
}

// ServeMuxOption is an option that can be given to a ServeMux on construction.
type ServeMuxOption func(*ServeMux)

// WithForwardResponseOption returns a ServeMuxOption representing the forwardResponseOption.
//
// forwardResponseOption is an option that will be called on the relevant context.Context,
// http.ResponseWriter, and proto.Message before every forwarded response.
//
// The message may be nil in the case where just a header is being sent.
func WithForwardResponseOption(forwardResponseOption func(context.Context, http.ResponseWriter, proto.Message) error) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.forwardResponseOptions = append(serveMux.forwardResponseOptions, forwardResponseOption)
	}
}

// WithUnescapingMode sets the escaping type. See the definitions of UnescapingMode
// for more information.
func WithUnescapingMode(mode UnescapingMode) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.unescapingMode = mode
	}
}

// SetQueryParameterParser sets the query parameter parser, used to populate message from query parameters.
// Configuring this will mean the generated OpenAPI output is no longer correct, and it should be
// done with careful consideration.
func SetQueryParameterParser(queryParameterParser QueryParameterParser) ServeMuxOption {
	return func(serveMux *ServeMux) {
		currentQueryParser = queryParameterParser
	}
}

// HeaderMatcherFunc checks whether a header key should be forwarded to/from gRPC context.
type HeaderMatcherFunc func(string) (string, bool)

// DefaultHeaderMatcher is used to pass http request headers to/from gRPC context. This adds permanent HTTP header
// keys (as specified by the IANA, e.g: Accept, Cookie, Host) to the gRPC metadata with the grpcgateway- prefix. If you want to know which headers are considered permanent, you can view the isPermanentHTTPHeader function.
// HTTP headers that start with 'Grpc-Metadata-' are mapped to gRPC metadata after removing the prefix 'Grpc-Metadata-'.
// Other headers are not added to the gRPC metadata.
func DefaultHeaderMatcher(key string) (string, bool) {
	switch key = textproto.CanonicalMIMEHeaderKey(key); {
	case isPermanentHTTPHeader(key):
		return MetadataPrefix + key, true
	case strings.HasPrefix(key, MetadataHeaderPrefix):
		return key[len(MetadataHeaderPrefix):], true
	}
	return "", false
}

func defaultOutgoingHeaderMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", MetadataHeaderPrefix, key), true
}

func defaultOutgoingTrailerMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", MetadataTrailerPrefix, key), true
}

// WithIncomingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for incoming request to gateway.
//
// This matcher will be called with each header in http.Request. If matcher returns true, that header will be
// passed to gRPC context. To transform the header before passing to gRPC context, matcher should return the modified header.
func WithIncomingHeaderMatcher(fn HeaderMatcherFunc) ServeMuxOption {
	for _, header := range fn.matchedMalformedHeaders() {
		grpclog.Warningf("The configured forwarding filter would allow %q to be sent to the gRPC server, which will likely cause errors. See https://github.com/grpc/grpc-go/pull/4803#issuecomment-986093310 for more information.", header)
	}

	return func(mux *ServeMux) {
		mux.incomingHeaderMatcher = fn
	}
}

// matchedMalformedHeaders returns the malformed headers that would be forwarded to gRPC server.
func (fn HeaderMatcherFunc) matchedMalformedHeaders() []string {
	if fn == nil {
		return nil
	}
	headers := make([]string, 0)
	for header := range malformedHTTPHeaders {
		out, accept := fn(header)
		if accept && isMalformedHTTPHeader(out) {
			headers = append(headers, out)
		}
	}
	return headers
}

// WithOutgoingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for outgoing response from gateway.
//
// This matcher will be called with each header in response header metadata. If matcher returns true, that header will be
// passed to http response returned from gateway. To transform the header before passing to response,
// matcher should return the modified header.
func WithOutgoingHeaderMatcher(fn HeaderMatcherFunc) ServeMuxOption {
	return func(mux *ServeMux) {
		mux.outgoingHeaderMatcher = fn
	}
}

// WithOutgoingTrailerMatcher returns a ServeMuxOption representing a headerMatcher for outgoing response from gateway.
//
// This matcher will be called with each header in response trailer metadata. If matcher returns true, that header will be
// passed to http response returned from gateway. To transform the header before passing to response,
// matcher should return the modified header.
func WithOutgoingTrailerMatcher(fn HeaderMatcherFunc) ServeMuxOption {
	return func(mux *ServeMux) {
		mux.outgoingTrailerMatcher = fn
	}
}

// WithMetadata returns a ServeMuxOption for passing metadata to a gRPC context.
//
// This can be used by services that need to read from http.Request and modify gRPC context. A common use case
// is reading token from cookie and adding it in gRPC context.
func WithMetadata(annotator func(context.Context, *http.Request) metadata.MD) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.metadataAnnotators = append(serveMux.metadataAnnotators, annotator)
	}
}

// WithErrorHandler returns a ServeMuxOption for configuring a custom error handler.
//
// This can be used to configure a custom error response.
func WithErrorHandler(fn ErrorHandlerFunc) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.errorHandler = fn
	}
}

// WithStreamErrorHandler returns a ServeMuxOption that will use the given custom stream
// error handler, which allows for customizing the error trailer for server-streaming
// calls.
//
// For stream errors that occur before any response has been written, the mux's
// ErrorHandler will be invoked. However, once data has been written, the errors must
// be handled differently: they must be included in the response body. The response body's
// final message will include the error details returned by the stream error handler.
func WithStreamErrorHandler(fn StreamErrorHandlerFunc) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.streamErrorHandler = fn
	}
}

// WithRoutingErrorHandler returns a ServeMuxOption for configuring a custom error handler to  handle http routing errors.
//
// Method called for errors which can happen before gRPC route selected or executed.
// The following error codes: StatusMethodNotAllowed StatusNotFound StatusBadRequest
func WithRoutingErrorHandler(fn RoutingErrorHandlerFunc) ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.routingErrorHandler = fn
	}
}

// WithDisablePathLengthFallback returns a ServeMuxOption for disable path length fallback.
func WithDisablePathLengthFallback() ServeMuxOption {
	return func(serveMux *ServeMux) {
		serveMux.disablePathLengthFallback = true
	}
}

// WithHealthEndpointAt returns a ServeMuxOption that will add an endpoint to the created ServeMux at the path specified by endpointPath.
// When called the handler will forward the request to the upstream grpc service health check (defined in the
// gRPC Health Checking Protocol).
//
// See here https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/health_check/ for more information on how
// to setup the protocol in the grpc server.
//
// If you define a service as query parameter, this will also be forwarded as service in the HealthCheckRequest.
func WithHealthEndpointAt(healthCheckClient grpc_health_v1.HealthClient, endpointPath string) ServeMuxOption {
	return func(s *ServeMux) {
		// error can be ignored since pattern is definitely valid
		_ = s.HandlePath(
			http.MethodGet, endpointPath, func(w http.ResponseWriter, r *http.Request, _ map[string]string,
			) {
				_, outboundMarshaler := MarshalerForRequest(s, r)

				resp, err := healthCheckClient.Check(r.Context(), &grpc_health_v1.HealthCheckRequest{
					Service: r.URL.Query().Get("service"),
				})
				if err != nil {
					s.errorHandler(r.Context(), s, outboundMarshaler, w, r, err)
					return
				}

				w.Header().Set("Content-Type", "application/json")

				if resp.GetStatus() != grpc_health_v1.HealthCheckResponse_SERVING {
					switch resp.GetStatus() {
					case grpc_health_v1.HealthCheckResponse_NOT_SERVING, grpc_health_v1.HealthCheckResponse_UNKNOWN:
						err = status.Error(codes.Unavailable, resp.String())
					case grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN:
						err = status.Error(codes.NotFound, resp.String())
					}

					s.errorHandler(r.Context(), s, outboundMarshaler, w, r, err)
					return
				}

				_ = outboundMarshaler.NewEncoder(w).Encode(resp)
			})
	}
}

// WithHealthzEndpoint returns a ServeMuxOption that will add a /healthz endpoint to the created ServeMux.
//
// See WithHealthEndpointAt for the general implementation.
func WithHealthzEndpoint(healthCheckClient grpc_health_v1.HealthClient) ServeMuxOption {
	return WithHealthEndpointAt(healthCheckClient, "/healthz")
}

// NewServeMux returns a new ServeMux whose internal mapping is empty.
func NewServeMux(opts ...ServeMuxOption) *ServeMux {
	serveMux := &ServeMux{
		handlers:               make(map[string][]handler),
		forwardResponseOptions: make([]func(context.Context, http.ResponseWriter, proto.Message) error, 0),
		marshalers:             makeMarshalerMIMERegistry(),
		errorHandler:           DefaultHTTPErrorHandler,
		streamErrorHandler:     DefaultStreamErrorHandler,
		routingErrorHandler:    DefaultRoutingErrorHandler,
		unescapingMode:         UnescapingModeDefault,
	}

	for _, opt := range opts {
		opt(serveMux)
	}

	if serveMux.incomingHeaderMatcher == nil {
		serveMux.incomingHeaderMatcher = DefaultHeaderMatcher
	}
	if serveMux.outgoingHeaderMatcher == nil {
		serveMux.outgoingHeaderMatcher = defaultOutgoingHeaderMatcher
	}
	if serveMux.outgoingTrailerMatcher == nil {
		serveMux.outgoingTrailerMatcher = defaultOutgoingTrailerMatcher
	}

	return serveMux
}

// Handle associates "h" to the pair of HTTP method and path pattern.
func (s *ServeMux) Handle(meth string, pat Pattern, h HandlerFunc) {
	s.handlers[meth] = append([]handler{{pat: pat, h: h}}, s.handlers[meth]...)
}

// HandlePath allows users to configure custom path handlers.
// refer: https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/inject_router/
func (s *ServeMux) HandlePath(meth string, pathPattern string, h HandlerFunc) error {
	compiler, err := httprule.Parse(pathPattern)
	if err != nil {
		return fmt.Errorf("parsing path pattern: %w", err)
	}
	tp := compiler.Compile()
	pattern, err := NewPattern(tp.Version, tp.OpCodes, tp.Pool, tp.Verb)
	if err != nil {
		return fmt.Errorf("creating new pattern: %w", err)
	}
	s.Handle(meth, pattern, h)
	return nil
}

// ServeHTTP dispatches the request to the first handler whose pattern matches to r.Method and r.URL.Path.
func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	path := r.URL.Path
	if !strings.HasPrefix(path, "/") {
		_, outboundMarshaler := MarshalerForRequest(s, r)
		s.routingErrorHandler(ctx, s, outboundMarshaler, w, r, http.StatusBadRequest)
		return
	}

	// TODO(v3): remove UnescapingModeLegacy
	if s.unescapingMode != UnescapingModeLegacy && r.URL.RawPath != "" {
		path = r.URL.RawPath
	}

	if override := r.Header.Get("X-HTTP-Method-Override"); override != "" && s.isPathLengthFallback(r) {
		if err := r.ParseForm(); err != nil {
			_, outboundMarshaler := MarshalerForRequest(s, r)
			sterr := status.Error(codes.InvalidArgument, err.Error())
			s.errorHandler(ctx, s, outboundMarshaler, w, r, sterr)
			return
		}
		r.Method = strings.ToUpper(override)
	}

	var pathComponents []string
	// since in UnescapeModeLegacy, the URL will already have been fully unescaped, if we also split on "%2F"
	// in this escaping mode we would be double unescaping but in UnescapingModeAllCharacters, we still do as the
	// path is the RawPath (i.e. unescaped). That does mean that the behavior of this function will change its default
	// behavior when the UnescapingModeDefault gets changed from UnescapingModeLegacy to UnescapingModeAllExceptReserved
	if s.unescapingMode == UnescapingModeAllCharacters {
		pathComponents = encodedPathSplitter.Split(path[1:], -1)
	} else {
		pathComponents = strings.Split(path[1:], "/")
	}

	lastPathComponent := pathComponents[len(pathComponents)-1]

	for _, h := range s.handlers[r.Method] {
		// If the pattern has a verb, explicitly look for a suffix in the last
		// component that matches a colon plus the verb. This allows us to
		// handle some cases that otherwise can't be correctly handled by the
		// former LastIndex case, such as when the verb literal itself contains
		// a colon. This should work for all cases that have run through the
		// parser because we know what verb we're looking for, however, there
		// are still some cases that the parser itself cannot disambiguate. See
		// the comment there if interested.

		var verb string
		patVerb := h.pat.Verb()

		idx := -1
		if patVerb != "" && strings.HasSuffix(lastPathComponent, ":"+patVerb) {
			idx = len(lastPathComponent) - len(patVerb) - 1
		}
		if idx == 0 {
			_, outboundMarshaler := MarshalerForRequest(s, r)
			s.routingErrorHandler(ctx, s, outboundMarshaler, w, r, http.StatusNotFound)
			return
		}

		comps := make([]string, len(pathComponents))
		copy(comps, pathComponents)

		if idx > 0 {
			comps[len(comps)-1], verb = lastPathComponent[:idx], lastPathComponent[idx+1:]
		}

		pathParams, err := h.pat.MatchAndEscape(comps, verb, s.unescapingMode)
		if err != nil {
			var mse MalformedSequenceError
			if ok := errors.As(err, &mse); ok {
				_, outboundMarshaler := MarshalerForRequest(s, r)
				s.errorHandler(ctx, s, outboundMarshaler, w, r, &HTTPStatusError{
					HTTPStatus: http.StatusBadRequest,
					Err:        mse,
				})
			}
			continue
		}
		h.h(w, r, pathParams)
		return
	}

	// if no handler has found for the request, lookup for other methods
	// to handle POST -> GET fallback if the request is subject to path
	// length fallback.
	// Note we are not eagerly checking the request here as we want to return the
	// right HTTP status code, and we need to process the fallback candidates in
	// order to do that.
	for m, handlers := range s.handlers {
		if m == r.Method {
			continue
		}
		for _, h := range handlers {
			var verb string
			patVerb := h.pat.Verb()

			idx := -1
			if patVerb != "" && strings.HasSuffix(lastPathComponent, ":"+patVerb) {
				idx = len(lastPathComponent) - len(patVerb) - 1
			}

			comps := make([]string, len(pathComponents))
			copy(comps, pathComponents)

			if idx > 0 {
				comps[len(comps)-1], verb = lastPathComponent[:idx], lastPathComponent[idx+1:]
			}

			pathParams, err := h.pat.MatchAndEscape(comps, verb, s.unescapingMode)
			if err != nil {
				var mse MalformedSequenceError
				if ok := errors.As(err, &mse); ok {
					_, outboundMarshaler := MarshalerForRequest(s, r)
					s.errorHandler(ctx, s, outboundMarshaler, w, r, &HTTPStatusError{
						HTTPStatus: http.StatusBadRequest,
						Err:        mse,
					})
				}
				continue
			}

			// X-HTTP-Method-Override is optional. Always allow fallback to POST.
			// Also, only consider POST -> GET fallbacks, and avoid falling back to
			// potentially dangerous operations like DELETE.
			if s.isPathLengthFallback(r) && m == http.MethodGet {
				if err := r.ParseForm(); err != nil {
					_, outboundMarshaler := MarshalerForRequest(s, r)
					sterr := status.Error(codes.InvalidArgument, err.Error())
					s.errorHandler(ctx, s, outboundMarshaler, w, r, sterr)
					return
				}
				h.h(w, r, pathParams)
				return
			}
			_, outboundMarshaler := MarshalerForRequest(s, r)
			s.routingErrorHandler(ctx, s, outboundMarshaler, w, r, http.StatusMethodNotAllowed)
			return
		}
	}

	_, outboundMarshaler := MarshalerForRequest(s, r)
	s.routingErrorHandler(ctx, s, outboundMarshaler, w, r, http.StatusNotFound)
}

// GetForwardResponseOptions returns the ForwardResponseOptions associated with this ServeMux.
func (s *ServeMux) GetForwardResponseOptions() []func(context.Context, http.ResponseWriter, proto.Message) error {
	return s.forwardResponseOptions
}

func (s *ServeMux) isPathLengthFallback(r *http.Request) bool {
	return !s.disablePathLengthFallback && r.Method == "POST" && r.Header.Get("Content-Type") == "application/x-www-form-urlencoded"
}

type handler struct {
	pat Pattern
	h   HandlerFunc
}
