package structs

type Data struct {
	Timestamp string `json:"timestamp"`
	Amount    string `json:"amount"`
	Delegator string `json:"delegator"`
	Level     string `json:"level"`
}

type Sender struct {
	Address string `json:"address"`
}

type Delegation struct {
	Timestamp string `json:"timestamp"`
	Sender    Sender `json:"sender"`
	Level     int    `json:"level"`
	Amount    int    `json:"amount"`
}
