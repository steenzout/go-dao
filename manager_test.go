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
	"github.com/steenzout/go-dao/mock"
)

// ManagerSuite test suite for the Manager struct.
type ManagerTestSuite struct {
	DatabaseTestSuite
}

func (s *ManagerTestSuite) Test() {
	ctx, err := s.manager.StartTransaction()
	if err != nil {
		s.Fail(err.Error())
	}
	if s.NotNil(ctx) {
		defer s.manager.EndTransaction(ctx)

		dao, err := s.manager.CreateDAO(ctx, mock.DAOMock)
		if err != nil {
			s.Fail(err.Error())
		}

		if s.NotNil(dao) {
			mockDAO := dao.(mock.MockDAO)
			mockDAO.MockSomething()

			err = s.manager.CommitTransaction(ctx)
			if err != nil {
				s.Fail(err.Error())
			}
		}
	}
}
