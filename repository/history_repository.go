package repository

import (
	"encoding/json"
	"os"
	"transaction_app/entities"
)

type HistoryRepository interface {
	LogHistory(history entities.History) error
}

type historyRepository struct {
	historyFile string
}

func NewHistoryRepository(historyFile string) HistoryRepository {
	return &historyRepository{
		historyFile: historyFile,
	}
}

func (h *historyRepository) LogHistory(history entities.History) error {
	var histories []entities.History

	data, err := os.ReadFile(h.historyFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &histories)
	if err != nil {
		return err
	}

	histories = append(histories, history)
	data, err = json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(h.historyFile, data, os.ModePerm)
}
