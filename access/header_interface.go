package access

type H interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	TraceID() string
}
