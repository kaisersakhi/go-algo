package shared

const (
    LessThan int = -1
    EqualTo = 0
    GreaterThan = 1
)
type Comparable[T any] interface {
    Compare(T) int
}