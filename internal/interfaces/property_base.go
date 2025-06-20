package interfaces

import (
	l "github.com/faelmori/logz"
	"github.com/google/uuid"
	"reflect"
)

// IPropertyValBase is an interface that defines the methods for a property value.
type IPropertyValBase[T any] interface {
	GetLogger() l.Logger
	GetID() uuid.UUID
	GetName() string
	Value() *T
	StartCtl() <-chan string
	Type() reflect.Type
	Get(async bool) any
	Set(t *T) bool
	Clear() bool
	IsNil() bool
	Serialize(format, filePath string) ([]byte, error)
	Deserialize(data []byte, format, filePath string) error
}
