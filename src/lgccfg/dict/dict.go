package dict

import (
	"sync"
)

// Type declaration
type Dictionary struct {
	mu sync.Mutex
	ma map[string]string
}

// Allocate new dictionary
func NewDictionary() *Dictionary {
	d := new(Dictionary)
	d.ma = make(map[string]string)
	return d
}

// Create unique entry
func (d *Dictionary) Create(k string, v string) bool {
	if _, ok := d.ma[k]; ok {
		return false
	}
	d.mu.Lock()
	d.ma[k] = v
	d.mu.Unlock()
	return true
}

// Read entry
func (d *Dictionary) Read(k string) (string, bool) {
	v, ok := d.ma[k]
	return v, ok
}

// Update existing entry
func (d *Dictionary) Update(k string, v string) bool {
	if _, ok := d.ma[k]; !ok {
		return false
	}
	d.mu.Lock()
	d.ma[k] = v
	d.mu.Unlock()
	return true
}

// Delete entry
func (d *Dictionary) Delete(k string, v string) bool {
	if _, ok := d.ma[k]; !ok {
		return false
	}
	d.mu.Lock()
	delete(d.ma, k)
	d.mu.Unlock()
	return true
}
