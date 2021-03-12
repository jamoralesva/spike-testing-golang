package demogomock

//Definimos una interfaz
// CRUD
type Data interface {
  Send(key string, value []byte) error
  Get(key string) ([]byte, error)
}

type DataImpl struct{
  //some dependencies
}

func NewData() Data {
  return &DataImpl{}
}
func (d DataImpl) Send(key string, value []byte) error {
  panic("implement me")
}
func (d DataImpl) Get(key string) ([]byte, error) {
  panic("implement me")
}