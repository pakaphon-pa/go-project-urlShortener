package test

func (s *TestSuite) Test_GetAll() {
	result := s.storage.Read("test")

	s.Assert().NotEmpty(result)
}
