package config

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	MarshalIndent() []byte
	Marshal() []byte
}
