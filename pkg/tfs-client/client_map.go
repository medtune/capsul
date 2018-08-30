package tfsclient

type Map map[string]TFSClient

func NewMap() Map {
	return make(Map, 5)
}

func (m Map) Register(name string, c TFSClient) {

}
