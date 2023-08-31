package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ar Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (ar Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (ar Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (ar Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
