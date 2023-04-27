package conversation

type Conversation interface {
	ID() string
	ConversionID() string
	AddMessage(message string)
	GetMessage() interface{}
	SetLastMessage(msg string)
	StrategyName() string
	SetStrategy(strategyName string)
}
