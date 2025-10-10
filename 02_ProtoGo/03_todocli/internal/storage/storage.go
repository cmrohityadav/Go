package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cmrohityadav/go/02_protogo/03_todocli/internal/types"
)

var GlobalTodo = &types.Todos{}

func Load(path string) (*types.Todos, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error: Faild to load file %w", err)
	}
	defer f.Close()

	var localTodos types.Todos
	err = json.NewDecoder(f).Decode(&localTodos);

	if err != nil {
		return nil, fmt.Errorf("error: Faild to load file %w", err)
	}

	return &localTodos, nil

}

func Save(path string) (error) {

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("faild to open JSON file for saving: %w", err)
	}
	defer f.Close() 

	
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(GlobalTodo); err != nil {
		return fmt.Errorf("faild to Encode JSON file for saving: %w", err)
	}

	return nil
}
