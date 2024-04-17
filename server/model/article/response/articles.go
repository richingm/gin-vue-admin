package response

// Mind represents the whole mind map.
type Mind struct {
	Meta   Meta       `json:"meta"`
	Format string     `json:"format"`
	Data   []MindNode `json:"data"`
}

// Meta contains metadata for the mind map.
type Meta struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
}

// MindNode represents a node in the mind map.
type MindNode struct {
	ID        string `json:"id"`
	IsRoot    bool   `json:"isroot,omitempty"`
	Topic     string `json:"topic"`
	ParentId  string `json:"parentid"`
	Direction string `json:"direction,omitempty"` // omitempty as it may not always be present
}
