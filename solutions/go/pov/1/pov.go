package pov

type Tree struct {
	root  string
	leafs []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		root:  value,
		leafs: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.root
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.leafs
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	if tr.Value() == from {
		return tr
	}

	if len(tr.leafs) == 0 {
		return nil
	}

	path := tr.findPath(from)
	if path == nil {
		return nil
	}

	var pov *Tree

	var prev *Tree
	var curr *Tree

	for i := len(path) - 1; i >= 0; i-- {
		prev = curr
		curr = path[i]

		if i == len(path)-1 {
			pov = curr
		}

		children := []*Tree{}
		if i-1 >= 0 {
			children = []*Tree{path[i-1]}
		}

		for _, ch := range curr.Children() {
			if ch == prev {
				continue
			}

			children = append(children, ch)
		}

		curr.leafs = children
	}

	return pov
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	pathFrom := tr.findPath(from)
	pathTo := tr.findPath(to)

	if pathFrom == nil || pathTo == nil {
		return []string{}
	}

	pathToMap := make(map[*Tree]int, len(pathTo))

	for i := 0; i < len(pathTo); i++ {
		pathToMap[pathTo[i]] = i
	}

	search := []string{}
	for i := len(pathFrom) - 1; i >= 0; i-- {
		node := pathFrom[i]
		search = append(search, node.Value())
		pointer, intersection := pathToMap[node]

		if intersection {
			for i := pointer + 1; i < len(pathTo); i++ {
				search = append(search, pathTo[i].Value())
			}

			break
		}
	}

	return search
}

func (tr *Tree) findPath(target string) []*Tree {
	if tr.Value() == target {
		return []*Tree{tr}
	}

	for _, ch := range tr.Children() {
		search := ch.findPath(target)

		if search != nil {
			return append([]*Tree{tr}, search...)
		}
	}

	return nil
}
