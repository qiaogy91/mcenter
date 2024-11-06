package permission

func NewMeta(m map[string]any) *Meta {
	return &Meta{data: m}
}

type Meta struct {
	data map[string]any
}

func (m *Meta) GetBool(k string) bool {
	v, ok := m.data[k]
	if !ok {
		return false
	}

	return v.(bool)
}

func (m *Meta) GetString(k string) string {
	v, ok := m.data[k]
	if !ok {
		return ""
	}

	return v.(string)
}
