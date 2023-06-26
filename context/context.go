package context

import (
	"context"
	"reflect"
)

// Get - Get the value associated to the key
func Get(ctx context.Context, key interface{}) (value interface{}) {

	value = ctx.Value(key)

	if value == nil {
		value = reflect.New(reflect.TypeOf(key).Elem()).Interface()
		ctx = context.WithValue(ctx, key, value)
	}

	return value
}

// Set - Put key and value attached to context
func Set(ctx context.Context, key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}
