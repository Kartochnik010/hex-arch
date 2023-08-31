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
	res, err := api.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(res, "subtraction")
	if err != nil {
		return 0, err
	}

	return res, nil
}
func (api Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	res, err := api.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(res, "multiplication")
	if err != nil {
		return 0, err
	}

	return res, nil
}
func (api Adapter) GetDivision(a int32, b int32) (int32, error) {
	res, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(res, "division")
	if err != nil {
		return 0, err
	}

	return res, nil
}
