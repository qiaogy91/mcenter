package endpoint

func NewMetaData(d map[string]any) *MetaData { return &MetaData{data: d} }

type MetaData struct {
	data map[string]any
}

func (md *MetaData) GetBool(k string) bool {
	v, ok := md.data[k]
	if !ok {
		return false
	}
	return v.(bool)
}

func (md *MetaData) GetString(k string) string {
	v, ok := md.data[k]
	if !ok {
		return ""
	}
	return v.(string)
}
