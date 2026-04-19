package handler

import "reflect"

type OptionalParam[T any] struct {
	Value T
	IsSet bool
}

func (o *OptionalParam[T]) Receiver() reflect.Value {
	return reflect.ValueOf(o).Elem().Field(0)
}

// React to request param being parsed to update internal state
// MUST have pointer receiver
func (o *OptionalParam[T]) OnParamSet(isSet bool, parsed any) {
	o.IsSet = isSet
}
