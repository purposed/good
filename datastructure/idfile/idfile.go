package idfile

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// IDFile represents a persistent string <==> uint32 mapping.
type IDFile struct {
	filepath string

	values []string
	idMap  map[string]uint32
	nextID uint32

	lock sync.RWMutex
}

// New initializes the idfile, reading the existing values if applicable.
func New(filepath string) (*IDFile, error) {
	m := &IDFile{
		filepath: filepath,
		idMap:    make(map[string]uint32),
	}

	if err := m.initialize(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *IDFile) initialize() error {
	if _, err := os.Stat(m.filepath); err == nil {
		raw, err := ioutil.ReadFile(m.filepath)
		if err != nil {
			return err
		}

		for i, value := range strings.Split(string(raw), "\n") {
			m.values = append(m.values, value)
			m.idMap[value] = uint32(i)
			m.nextID++
		}
	}
	return nil
}

// AddValue adds a new value to the IDFile (if it doesn't already exist), returning its assigned ID.
func (m *IDFile) AddValue(value string) uint32 {
	m.lock.Lock()
	defer m.lock.Unlock()

	if valID, ok := m.idMap[value]; ok {
		return valID
	}

	myID := m.nextID

	m.nextID++
	m.idMap[value] = myID
	m.values = append(m.values, value)

	return myID
}

// ResolveValue returns the ID corresponding to the specified value.
func (m *IDFile) ResolveValue(value string) (uint32, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if valID, ok := m.idMap[value]; ok {
		return valID, true
	}
	return 0, false
}

// ResolveID returns the value of a given ID.
func (m *IDFile) ResolveID(valID uint32) (string, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if valID < 0 || valID >= m.nextID {
		return "", errors.New("invalid ID")
	}
	return m.values[valID], nil
}

// Commit writes the IDFile to disk.
func (m *IDFile) Commit() error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	raw := strings.Join(m.values, "\n")
	if err := ioutil.WriteFile(m.filepath, []byte(raw), 0600); err != nil {
		return err
	}
	return nil
}

// NextID returns the next ID that will be assigned by the idfile.
func (m *IDFile) NextID() uint32 {
	return m.nextID
}
