package dao_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/steenzout/go-dao"
)

func TestProcess(t *testing.T) {
	f := func(m dao.Manager, ctx *dao.Context) error {
		tm := m.(*TestManager)
		dao, err := tm.CreateMockDAO(ctx)
		if err != nil {
			return err
		}
		dao.MockSomething()

		return nil
	}
	assert.Nil(t, dao.Process(&manager, f))
}
