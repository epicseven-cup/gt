package cache

import (
	"encoding/json"
	"os"
)

const BasePath = "~/.gt/cache"

type Cache struct {
	ProjectName  string   `json:"project_name"`
	CacheHeaders []string `json:"cache_headers"`
	CacheTags    []string `json:"cache_tags"`
	CacheIssues  []string `json:"cache_issues"`
	Path         string   `json:"path,omitempty"`
}

func NewCache(projectName string) (*Cache, error) {
	//_, err := os.Stat(BasePath)
	//if errors.Is(err, os.ErrNotExist) {
	//	err = os.Mkdir(BasePath, 0755)
	//}
	//if err != nil {
	//	return nil, err
	//}
	//configPath := filepath.Join(BasePath, projectName+".json")
	//data, err := os.ReadFile(configPath)
	//if err != nil {
	//	return nil, err
	//}
	//cache := &Cache{}
	//// Since the Path in the json is optional, this will just gives it a default value, just in case the path does not exist in the json
	//cache.Path = configPath
	//err = json.Unmarshal(data, cache)
	//if err != nil {
	//	return nil, err
	//}
	//return cache, nil

	return nil, nil
}

// Save writes into cache of the project
func (cache *Cache) Save() error {
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return os.WriteFile(cache.Path, data, 0644)
}
