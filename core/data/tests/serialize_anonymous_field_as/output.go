// Code generated by typeshare 1.0.0. DO NOT EDIT.
package proto

import "encoding/json"

type SomeEnumTypes string
const (
	// The associated String contains some opaque context
	SomeEnumTypeVariantContext SomeEnumTypes = "Context"
	SomeEnumTypeVariantOther SomeEnumTypes = "Other"
)
type SomeEnum struct{ 
	Type SomeEnumTypes `json:"type"`
	content interface{}
}

func (s *SomeEnum) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    SomeEnumTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	s.Type = enum.Tag
	switch s.Type {
	case SomeEnumTypeVariantContext:
		var res string
		s.content = &res
	case SomeEnumTypeVariantOther:
		var res int
		s.content = &res

	}
	if err := json.Unmarshal(enum.Content, &s.content); err != nil {
		return err
	}

	return nil
}

func (s SomeEnum) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    SomeEnumTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = s.Type
    enum.Content = s.content
    return json.Marshal(enum)
}

func (s SomeEnum) Context() string {
	res, _ := s.content.(*string)
	return *res
}
func (s SomeEnum) Other() int {
	res, _ := s.content.(*int)
	return *res
}

func NewSomeEnumTypeVariantContext(content string) SomeEnum {
    return SomeEnum{
        Type: SomeEnumTypeVariantContext,
        content: &content,
    }
}
func NewSomeEnumTypeVariantOther(content int) SomeEnum {
    return SomeEnum{
        Type: SomeEnumTypeVariantOther,
        content: &content,
    }
}

