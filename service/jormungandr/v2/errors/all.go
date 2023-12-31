package errors

type RuntimeError string

func (e RuntimeError) Error() string {
	return string(e)
}

const (
	// 定义各类运行时错误
	EntityIDError          RuntimeError = "Entity ID is wrong."
	RequestTickTooBigError RuntimeError = "Request too many tick. Consider use multiple requests instead."
	ValueError             RuntimeError = "Value error."
	RequestError           RuntimeError = "Request value error."
	NotImplementError      RuntimeError = "Not implement."
	TickError              RuntimeError = "Error while running tick."
	BaseException          RuntimeError = "Unknown error."
)
