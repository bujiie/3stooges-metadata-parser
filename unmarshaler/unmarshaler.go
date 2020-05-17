package unmarshaler

type UnmarshalerIFace interface {
	Unmarshal(data []byte) (interface{}, error)
}
