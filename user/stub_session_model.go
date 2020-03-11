package user

type StubSessionModel struct {
    createCalls []int
    getIdCalls []string
}

func (s *StubSessionModel) Create(userId int) error {
    s.createCalls = append(s.createCalls, userId)
    return nil
}

func (s *StubSessionModel) GetId() int {
    s.getIdCalls = append(s.getIdCalls, "")
    return 333
}
