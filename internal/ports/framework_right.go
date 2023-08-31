package ports

type DbPort interface {
	CloseDBConnection() error
	AddToHistory(answer int32, operation string) error
}
