package FileReader

func Get_Test_Data(file_name string) (srcNodes []string, dstNodes []string) {
	return
}

type Edge struct {
	U int
	V int
	W int
}

func Get_All_Edges(file_name string) (edges []Edge) {
	return
}

type Node struct {
	Id   int
	Name string
}

func Get_All_Nodes(file_name string) (nodes []Node) {
	return
}

// Define results structure
type Result struct {
	Src    string
	Dst    string
	Weight int
	Path   string
}

func Read_Results_Data(file_name string) (results []Result) {
	// Convert path array to a string such as: [Node4, Node3, Node1] => "Node1 -> Node3 -> Node4"
	return
}
