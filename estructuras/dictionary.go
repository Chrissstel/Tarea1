package estructuras

type Dictionary struct {
	items map[string]int
}

// Constructor
func NewDictionary() *Dictionary {
	return &Dictionary{
		items: make(map[string]int),
	}
}

// Put
func (d *Dictionary) Put(key string, value int) {
	d.items[key] = value
}

// Get
func (d *Dictionary) Get(key string) (int, bool) {
	value, ok := d.items[key]
	return value, ok
}

// Remove
func (d *Dictionary) Remove(key string) {
	delete(d.items, key)
}

// Contains
func (d *Dictionary) Contains(key string) bool {
	_, ok := d.items[key]
	return ok
}

// Size
func (d *Dictionary) Size() int {
	return len(d.items)
}
