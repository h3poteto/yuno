package kingtime

import (
	"encoding/json"
	"time"
)

type RequestTimerecord struct {
	Date string    `json:"date"`
	Time time.Time `json:"time"`
	Code int       `json:"code"`
}

type ResponseTimerecord struct {
	Date        string `json:"date"`
	EmployeeKey string `json:employeeKey`
}

func (k *Kingtime) Attendance(employeeKey string) (*ResponseTimerecord, error) {
	current := time.Now()
	date := current.Format("2006-01-02")
	input := &RequestTimerecord{
		Code: 1,
		Date: date,
		Time: current,
	}
	body, err := json.Marshal(input)
	output, err := k.post("/daily-workings/timerecord/"+employeeKey, body)
	if err != nil {
		return nil, err
	}
	var response ResponseTimerecord

	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
