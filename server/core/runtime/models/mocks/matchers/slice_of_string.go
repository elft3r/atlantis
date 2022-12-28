// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnySliceOfString() []string {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]string))(nil)).Elem()))
	var nullValue []string
	return nullValue
}

func EqSliceOfString(value []string) []string {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []string
	return nullValue
}

func NotEqSliceOfString(value []string) []string {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue []string
	return nullValue
}

func SliceOfStringThat(matcher pegomock.ArgumentMatcher) []string {
	pegomock.RegisterMatcher(matcher)
	var nullValue []string
	return nullValue
}
