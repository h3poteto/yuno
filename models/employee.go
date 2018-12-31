package models

type Employee struct {
	Code      string `json:"code`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Key       string `json:"key"`
	TypeName  string `json:"type_name"`
	SlackID   string `json:"slack_id"`
	db        *database
}

func NewEmployee(code, lastName, firstName, key, typeName, slackID string) (*Employee, error) {
	db, err := Initialize()
	if err != nil {
		return nil, err
	}
	return &Employee{
		code,
		lastName,
		firstName,
		key,
		typeName,
		slackID,
		db,
	}, nil
}

func (e *Employee) Save() error {
	_, err := e.db.db.Exec(
		`INSERT INTO employees (code, last_name, first_name, key, type_name, slack_id) VALUES (?, ?, ?, ?, ?, ?)`,
		e.Code,
		e.LastName,
		e.FirstName,
		e.Key,
		e.TypeName,
		e.SlackID,
	)
	return err
}

func GetEmployeeFromSlackID(slackID string) (*Employee, error) {
	db, err := Initialize()
	if err != nil {
		return nil, err
	}
	row := db.db.QueryRow(`SELECT * FROM employees WHERE slack_id = ?`, slackID)
	var code, lastName, firstName, key, typeName string
	err = row.Scan(&code, &lastName, &firstName, &key, &typeName)
	if err != nil {
		return nil, err
	}
	return &Employee{
		code,
		lastName,
		firstName,
		key,
		typeName,
		slackID,
		db,
	}, nil
}
