package model

func (n *Note) Id() string {
	return n.ID
}

func (n *Note) Nature() string {
	return "note"
}

func (n *Note) SetId(id string) {
	n.ID = id
}

func (r *Rendering) Id() string {
	return r.ID
}

func (r *Rendering) Nature() string {
	return "rendering"
}

func (r *Rendering) SetId(id string) {
	r.ID = id
}
