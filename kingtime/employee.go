package kingtime

import "encoding/json"

type ResponseEmployee struct {
	Code      string `json:"code`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Key       string `json:"key"`
	TypeName  string `json:"type_name"`
}

func (k *Kingtime) GetEmployee(code string) (*ResponseEmployee, error) {
	output, err := k.get("/employees/" + code)
	if err != nil {
		return nil, err
	}
	var response ResponseEmployee
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
