package cube

import (
	"time"
	"encoding/json"
)

type Message struct {
	Version string          `json:"version"`
	Id      string          `json:"id"`
	From    string          `json:"from"`
	To      string          `json:"to"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type LogMessage struct {
	Version    string `json:"version"`
	Id         string `json:"id"`
	Class      string `json:"class"`
	InstanceId string `json:"instanceId"`
	Level      string `json:"level"`
	Text       string `json:"text"`
}

type Cube interface {
	GetParams() map[string]string
	GetClass() string
	GetInstanceId() string
	PublishMessage(toChannel string, message Message)
	MakeRequest(channel string, message Message, timeout time.Duration)

	LogDebug(text string)
	LogError(text string)
	LogFatal(text string)
	LogInfo(text string)
	LogWarning(text string)
	LogTrace(text string)
}

type HandlerInterface interface {
	OnStart(instance Cube)
	OnStop(instance Cube)
	OnReceiveMessage(instance Cube, message Message)
	OnReceiveRequest(instance Cube, message Message, replyToRequest func(Message) error)
}
