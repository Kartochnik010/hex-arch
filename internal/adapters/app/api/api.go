package api

import "hex/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (api Adapter) GetAddition(a int32, b int32) (int32, error) {
	res, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(res, "addition")
	if err != nil {
		return 0, err
	}

	return res, nil
}
func (api Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	return api.arith.Subtraction(a, b)
}
func (api Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	return api.arith.Multiplication(a, b)
}
func (api Adapter) GetDivision(a int32, b int32) (int32, error) {
	return api.arith.Division(a, b)
}
