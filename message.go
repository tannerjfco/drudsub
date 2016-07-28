package drudsub

const (
	SeverityFatal   = "FATAL"
	SeverityError   = "ERROR"
	SeverityWarning = "WARNING"
	SeverityInfo    = "INFO"
	SeverityDebug   = "DEBUG"
)

// Message represents a Pub/Sub/NATS message.
type Message struct {
	Data       []byte
	Attributes map[string]string
}
