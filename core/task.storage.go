package core

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/ngrink/tasker/lib"
)

var (
	tasksPath = lib.GetDataBasepath() + "tasks.json"
)

func (t *TaskContainer) Load() error {
	file, err := os.ReadFile(tasksPath)
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

func (t *TaskContainer) Save() {
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll(filepath.Dir(tasksPath), 0755)
	err = os.WriteFile(tasksPath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
