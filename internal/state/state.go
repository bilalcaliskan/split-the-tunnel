package state

import (
	"encoding/json"
	"os"
)

type RouteEntry struct {
	Domain     string `json:"domain"`
	ResolvedIP string `json:"resolvedIP"`
	Gateway    string `json:"gateway"`
}

type State struct {
	Entries []RouteEntry `json:"entries"`
}

func (s *State) AddEntry(entry RouteEntry) {
	s.Entries = append(s.Entries, entry)
}

func (s *State) RemoveEntry(domain string) {
	for i, entry := range s.Entries {
		if entry.Domain == domain {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			break
		}
	}
}

func (s *State) GetEntry(domain string) *RouteEntry {
	for _, entry := range s.Entries {
		if entry.Domain == domain {
			return &entry
		}
	}

	return nil
}

func (s *State) Read(path string) error {
	// Attempt to get the file status
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create an empty state and write to new file
			return s.Write(path)
		}
		// Some other error occurred
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, s)
	return err
}

func (s *State) Write(path string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
