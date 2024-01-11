package transaction

import "time"

type CampaignTrasactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTrasactionFormatter {
	formatter := CampaignTrasactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter

}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTrasactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTrasactionFormatter{}
	}
	var transactionsFormatter []CampaignTrasactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}
