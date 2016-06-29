package mock

import "github.com/steenzout/go-dao"

func NewDataSource() *dao.DataSource {
	return &dao.DataSource{
		DB: nil,
		Name: "mock",
	}
}
