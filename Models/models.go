package Models

type Transaction struct {
	Amount           string
	CardShortNo      string
	CardFullNo       string
	Place            string
	Date             string
	Time             string
	BalanceAvailable string
}

type Instapay struct {
	Amount      string
	CardShortNo string
	CardFullNo  string
	Date        string
	Time        string
}

type MessageReceived struct {
	DateTime string `json:"dateTime"`
	Card     string `json:"card"`
	Amount   string `json:"amount"`
	Notes    string `json:"notes"`
	Bank     string `json:"bank"`
}
