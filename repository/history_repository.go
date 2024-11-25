package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transaction_app/entities"
)

type HistoryRepository interface {
	ReadHistory() ([]entities.History, error)
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

func (h *historyRepository) ReadHistory() ([]entities.History, error) {
	var histories []entities.History
	file, err := os.Open(h.historyFile)
	if err != nil {
		return nil, errors.New("history file not found")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&histories); err != nil {
		return nil, errors.New("failed to unmarshal history file")
	}

	return histories, nil
}

func (h *historyRepository) LogHistory(history entities.History) error {
	var histories []entities.History

	data, err := os.ReadFile(h.historyFile)
	if err != nil {
		return errors.New("history file not found")
	}

	err = json.Unmarshal(data, &histories)
	if err != nil {
		return errors.New("failed to unmarshal history file")
	}

	histories = append(histories, history)
	data, err = json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return errors.New("failed to marshal history file")
	}

	return os.WriteFile(h.historyFile, data, os.ModePerm)
}
