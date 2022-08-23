package requestmeta

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// RequestMetadataHeaderKey defines a key in the request metadata header.
type RequestMetadataHeaderKey string

// BoolRequestMetadataHeaderKey defines a key for a boolean vlaue in the request metadata header.
type BoolRequestMetadataHeaderKey RequestMetadataHeaderKey

const (
	// RequestServerVersion, if specified in a request header, asks SpiceDB to return its
	// server version in the response header (if supported).
	// Value: `1`
	RequestServerVersion BoolRequestMetadataHeaderKey = "io.spicedb.requestversion"

	// RequestDebugInformation, if specified in a request header, asks SpiceDB to return debug information
	// for the API call (if applicable and supported).
	// Value: `1`
	RequestDebugInformation BoolRequestMetadataHeaderKey = "io.spicedb.requestdebuginfo"
)

// AddRequestHeaders returns a new context with the given values as request headers.
func AddRequestHeaders(ctx context.Context, keys ...BoolRequestMetadataHeaderKey) context.Context {
	values := make(map[RequestMetadataHeaderKey]string, len(keys))
	for _, key := range keys {
		values[RequestMetadataHeaderKey(key)] = "1"
	}
	return SetRequestHeaders(ctx, values)
}

// SetRequestHeaders returns a new context with the given values as request headers.
func SetRequestHeaders(ctx context.Context, values map[RequestMetadataHeaderKey]string) context.Context {
	pairs := make([]string, 0, len(values)*2)
	for key, value := range values {
		pairs = append(pairs, string(key))
		pairs = append(pairs, value)
	}
	return metadata.AppendToOutgoingContext(ctx, pairs...)
}