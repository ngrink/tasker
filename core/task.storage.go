package core

import (
	"encoding/json"
	"errors"
	"os"
)

func (t *TaskContainer) Load() error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Init()
			t.Save()
			return nil
		}
		return err
	}

	if len(file) == 0 {
		t.Init()
		t.Save()
		return nil
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskContainer) Save() error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}
