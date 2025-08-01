package types

import (
	"database/sql/driver"
	"encoding/json"
	"os"
	"reflect"

	ci "github.com/rafa-mori/gobe/internal/interfaces"
	gl "github.com/rafa-mori/gobe/logger"
	l "github.com/rafa-mori/logz"

	"github.com/google/uuid"
)

type JsonB map[string]interface{}

// Serializer manual para o GORM
func (m JsonB) Value() (driver.Value, error) { return json.Marshal(m) }

func (m *JsonB) Scan(vl any) error {
	if vl == nil {
		*m = JsonB{}
		return nil
	}
	return json.Unmarshal(vl.([]byte), m)
}

// Property is a struct that holds the properties of the GoLife instance.
type Property[T any] struct {
	// Telemetry is the telemetry for this GoLife instance.
	metrics *Telemetry
	// Prop is the property for this GoLife instance.
	prop ci.IPropertyValBase[T]
	// Cb is the callback function for this GoLife instance.
	cb func(any) (bool, error)
}

// NewProperty creates a new IProperty[T] with the given value and Reference.
func NewProperty[T any](name string, v *T, withMetrics bool, cb func(any) (bool, error)) ci.IProperty[T] {
	p := &Property[T]{
		prop: newVal[T](name, v),
		cb:   cb,
	}
	if withMetrics {
		p.metrics = NewTelemetry()
	}
	return p
}

// GetName returns the name of the property.
func (p *Property[T]) GetName() string {
	return p.prop.GetName()
}

// GetValue returns the value of the property.
func (p *Property[T]) GetValue() T {
	value := p.prop.Get(false)
	if value == nil {
		return *new(T)
	}
	return *value.(*T)
}

// SetValue sets the value of the property.
func (p *Property[T]) SetValue(v *T) {
	p.prop.Set(v)
	if p.cb != nil {
		if _, err := p.cb(v); err != nil {
			//p.metrics.Log("error", "Error in callback function: "+err.Error())
		}
	}
}

// GetReference returns the reference of the property.
func (p *Property[T]) GetReference() (uuid.UUID, string) {
	return p.prop.GetID(), p.prop.GetName()
}

// Prop is a struct that holds the properties of the GoLife instance.
func (p *Property[T]) Prop() ci.IPropertyValBase[T] {
	return p.prop
}

// GetLogger returns the logger of the property.
func (p *Property[T]) GetLogger() l.Logger {

	return p.Prop().GetLogger()

}

// Serialize serializes the ProcessInput instance to the specified format.
func (p *Property[T]) Serialize(format, filePath string) ([]byte, error) {
	value := p.GetValue()
	mapper := NewMapper[T](&value, filePath)
	return mapper.Serialize(format)
}

// Deserialize deserializes the data into the ProcessInput instance.
func (p *Property[T]) Deserialize(data []byte, format, filePath string) error {

	if len(data) == 0 {
		return nil
	}
	value := p.GetValue()
	if !reflect.ValueOf(value).IsValid() {
		p.SetValue(new(T))
	}
	mapper := NewMapper[T](&value, filePath)
	if v, vErr := mapper.Deserialize(data, format); vErr != nil {
		gl.Log("error", "Failed to deserialize data:", vErr.Error())
		return vErr
	} else {
		p.SetValue(v)
	}
	return nil
}

func (p *Property[T]) SaveToFile(filePath string, format string) error {
	if data, err := p.Serialize(format, filePath); err != nil {
		gl.Log("error", "Failed to serialize data:", err.Error())
		return err
	} else {
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			gl.Log("error", "Failed to write to file:", err.Error())
			return err
		}
	}
	return nil
}

func (p *Property[T]) LoadFromFile(filename, format string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return p.Deserialize(data, format, filename)
}
