package capsul

func New() {

}

type Model struct {
	Name    string
	Version string
}

type Dataset struct {
	Name    string
	Version string
}

type Capsul struct {
	ID           string
	Name         string
	Dataset      *Dataset
	TrainedModel *Model
}
