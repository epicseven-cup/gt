package cache

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

const BASE_PATH = "/.gt/cache"

const HEADER_REGEX = `\w+:`
const ISSUE_REGEX = `#\d+`

type Cache struct {
	ProjectName  string          `json:"project_name"`
	CacheHeaders map[string]bool `json:"cache_headers"`
	CacheTags    map[string]bool `json:"cache_tags"`
	CacheIssues  map[string]bool `json:"cache_issues"`
	Path         string          `json:"path,omitempty"`
}

func CreateDefaultConfigFile(projectName string, path string) (*Cache, error) {
	defaultConfig := &Cache{
		ProjectName: projectName,
		Path:        path,
	}
	jsonConfig, err := json.Marshal(defaultConfig)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, jsonConfig, 0644)
	if err != nil {
		return nil, err
	}
	return defaultConfig, nil
}

func addCache(regex string, storage map[string]bool, content string) (map[string]bool, error) {
	match, err := regexp.MatchString(regex, content)
	if err != nil {
		return storage, err
	}
	if match {
		r, err := regexp.Compile(regex)
		if err != nil {
			return storage, err
		}
		for _, i := range r.FindAll([]byte(content), -1) {
			storage[string(i)] = true
		}
	}
	return storage, nil
}

// GetCache gets the cache information of the current project, if it doesn't exist it will create one
func GetCache(projectName string) (*Cache, error) {
	home, err := os.UserHomeDir()
	headerPath := filepath.Join(home, BASE_PATH)
	_, err = os.Stat(headerPath)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(headerPath, 0755)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(headerPath, projectName+".json")

	_, err = os.Stat(configPath)

	// Create Default config file
	if errors.Is(err, os.ErrNotExist) {
		c, err := CreateDefaultConfigFile(projectName, configPath)
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	// Get config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	cache := &Cache{}

	err = json.Unmarshal(data, cache)
	if err != nil {
		return nil, err
	}
	return cache, nil
}

// Save writes into cache of the project
func (c *Cache) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(c.Path, data, 0644)
}

func (c *Cache) Update(content string) error {
	headers, err := addCache(HEADER_REGEX, c.CacheHeaders, content)
	if err != nil {
		return err
	}
	c.CacheHeaders = headers

	issues, err := addCache(ISSUE_REGEX, c.CacheIssues, content)
	if err != nil {
		return err
	}
	c.CacheIssues = issues
	return nil
}
