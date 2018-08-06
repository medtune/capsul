package testutil

type Dataframe struct {
	FilePath string
	Type     string
	Label    string
}

type TestGrid struct {
	Testdata []Dataframe
}

var GlobalGrid = map[string]Dataframe{
	"inception": {"", "", ""},
	"mnist":     {"", "", ""},
}
