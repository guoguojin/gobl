package property

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Properties is a map of properties
type Properties map[string]Property

// NewProperties creates a new map of properties
func NewProperties() Properties {
	return make(Properties)
}

// Set adds or updates the property with the given property key
func (p Properties) Set(name string, property Property) Properties {
	p[name] = property
	return p
}

// Add inserts a new property to the properties map with the given property key
// If the key already exists, the property will be updated
func (p Properties) Add(property Property) Properties {
	p[property.Name] = property
	return p
}

// Get retrieves the property with the given key if it exists and a bool indicating the existence of a key in the map
func (p Properties) Get(key string) (Property, bool) {
	v, ok := p[key]
	return v, ok
}

// GetOrElse will retrieve the value at the given key if it exists,
// or inserts the other value into the map with the given key and returns
// that value instead
func (p Properties) GetOrElse(key string, other Property) Property {
	v, ok := p[key]
	if !ok {
		p[key] = other
		return other
	}
	return v
}

func (p Properties) Properties() []Property {
	properties := make([]Property, 0, len(p))
	for _, v := range p {
		properties = append(properties, v)
	}
	return properties
}

func (p Properties) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Properties) UnmarshalBinary(data []byte) error {
	properties := NewProperties()
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&properties); err != nil {
		return err
	}

	// now we have to convert the data to their intended types
	// because the json types are limited
	for k := range properties {
		pr := properties[k]
		switch pr.Value.(type) {
		case json.Number:
			num, ok := pr.Value.(json.Number)
			if !ok {
				return fmt.Errorf("value %v is not a json.Number", pr.Value)
			}
			if err := convertJSONNumber(&pr, num); err != nil {
				return err
			}
			p.Set(k, pr)
		case []json.Number:
			nums, ok := pr.Value.([]json.Number)
			if !ok {
				return fmt.Errorf("value %v is not a []json.Number", pr.Value)
			}
			if err := convertJSONNumberSlice(&pr, nums); err != nil {
				return err
			}
			p.Set(k, pr)
		case []interface{}:
			vs, ok := pr.Value.([]interface{})
			if !ok {
				return fmt.Errorf("value %v is not a []interface{}", pr.Value)
			}

			if len(vs) == 0 {
				p.Set(k, pr)
				continue
			}
			switch vs[0].(type) {
			case json.Number:
				nums := make([]json.Number, len(vs))
				for i, v := range vs {
					nums[i], ok = v.(json.Number)
					if !ok {
						return fmt.Errorf("value %v is not a json.Number", pr.Value)
					}
				}
				if err := convertJSONNumberSlice(&pr, nums); err != nil {
					return err
				}
				p.Set(k, pr)
			case string:
				if err := convertInterfaceSliceOfStrings(&pr, vs); err != nil {
					return err
				}
				p.Set(k, pr)
			case bool:
				if err := convertInterfaceSliceOfBools(&pr, vs); err != nil {
					return err
				}
				p.Set(k, pr)
			default:
				p.Set(k, pr)
			}
		case interface{}:
			if err := convertInterface(&pr); err != nil {
				return err
			}
			p.Set(k, pr)
		default:
			p.Set(k, pr)
		}
	}
	return nil
}

func (p Properties) Value() (driver.Value, error) {
	return p.MarshalBinary()
}

func (p *Properties) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("cannot scan nil into Properties")
	}

	switch v := value.(type) {
	case []byte:
		return p.UnmarshalBinary(v)
	case string:
		return p.UnmarshalBinary([]byte(v))
	default:
		return fmt.Errorf("cannot scan type %T into Properties", value)
	}
}
