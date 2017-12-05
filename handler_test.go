//
// Copyright 2016 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package dao_test

import (
	"errors"

	"github.com/stretchr/testify/assert"

	"github.com/steenzout/go-dao"
	"github.com/steenzout/go-dao/mock"
)

// HandlerTestSuite test suite for the Manager struct.
type HandlerTestSuite struct {
	DatabaseTestSuite
}

// TestProcess test for process function.
func (s *HandlerTestSuite) TestProcess() {
	f := func(ctx *dao.Context) error {
		dao, err := s.manager.CreateDAO(ctx, mock.DAOMock)
		if err != nil {
			return err
		}
		return dao.(mock.MockDAO).MockSomething()
	}
	err := dao.Process(s.manager, f)
	s.Nil(err)
}

// TestProcessWhenPanic test for process function that causes panic.
func (s *HandlerTestSuite) TestProcessWhenPanic() {
	f := func(ctx *dao.Context) error {
		panic(errors.New("something happened"))
	}
	err := dao.Process(s.manager, f)
	if s.NotNil(err) {
		expected := errors.New("panic: something happened")
		assert.Equal(s.T(), expected, err)
	}
}
