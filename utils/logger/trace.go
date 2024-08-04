package logger

import (
	"context"

	"github.com/google/uuid"
)

func GetMetadataKey() string {
	return "trace_metadata"
}

func GetTraceIdKey() string {
	return "x_trace_id"
}

// InitTraceNode
func GenLogTraceMetadata() *TraceNode {
	t := NewTraceNode()
	t.Set("x_trace_id", uuid.New().String())
	return t
}

// TraceNode add other kv
func InjectMetadata(ctx context.Context, mapPtr *map[string]string) bool {
	meta := ExtractTraceNodeFromContext(ctx)
	for k, v := range meta.ForkMap() {
		(*mapPtr)[k] = v
	}
	return true
}

// Get TraceNode
func ExtractTraceNodeFromContext(ctx context.Context) *TraceNode {
	meta := ctx.Value(GetMetadataKey())
	if meta == nil {
		return NewTraceNode()
	} else {
		if val, ok := meta.(*TraceNode); ok {
			return val
		} else {
			return NewTraceNode()
		}
	}
}
