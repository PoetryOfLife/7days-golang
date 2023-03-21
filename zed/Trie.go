package zed

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}
