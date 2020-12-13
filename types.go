// @Author: abbeymart | Abi Akindele | @Created: 2020-11-29 | @Updated: 2020-11-29
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcdata

// simple functional types
// PredicateType
type PredicateType func(val interface{}) bool

// BinaryPredicateType
type BinaryPredicateType func(p, v interface{}) bool

// IntPredicateType
type IntPredicateType func(val int) bool

// FloatPredicateType
type FloatPredicateType func(val float64) bool

// StringPredicateType
type StringPredicateType func(val string) bool

// UnaryOperatorType
type UnaryOperatorType func(val interface{}) interface{}

// BiOperatorType
type BiOperatorType func(val1, val2 interface{}) interface{}

// FunctionType
type FunctionType func(val interface{}) interface{}

// BiFunctionType
type BiFunctionType func(val1, val2 interface{}) interface{}

// ConsumerType
type ConsumerType func(val interface{})

// BiConsumerType
type BiConsumerType func(val1, val2 interface{})

// SupplierType
type SupplierType func() interface{}

// ComparatorType
type ComparatorType func(val1, val2 interface{}) int

// interface method types
// Predicate: sort
type Predicate interface {
	test(p interface{}) bool
}

// BinaryPredicate
type BinaryPredicate interface {
	test(p, v interface{}) bool
}
// IntPredicate
type IntPredicate interface {
	test(p int) bool
}

// IntPredicate
type FloatPredicate interface {
	test(p float64) bool
}

// StringPredicate
type StringPredicate interface {
	test(p string) bool
}

// Consumer interface use case: ForEach function
type Consumer interface {
	accept(p interface{})
}

// BiConsumer interface use case: ForEach function
type BiConsumer interface {
	accept(p, v interface{})
}

// Supplier interface use case: Generator function
type Supplier interface {
	supply() interface{}
}

// Comparator: interface use case: Sort function
type Comparator interface {
	compare(p1, p2 interface{}) int
}

// UnaryOperator interface use case: Map function
type UnaryOperator interface {
	apply(p interface{}) interface{}
}

// BiOperator
type BiOperator interface {
	apply(p, v interface{}) interface{}
}

// Function interface use case: Map function
type Function interface {
	apply(p interface{}) interface{}
}

// BiFunction
type BiFunction interface {
	apply(p, v interface{}) interface{}
}
