package user

const correctSessionId = 54321

type StubSessionModel struct {
    createCalls []int
    getIdCalls []string
    getUserIdCalls []string
    getCalls []int
}

func (s *StubSessionModel) Create(userId int) error {
    s.createCalls = append(s.createCalls, userId)
    return nil
}

func (s *StubSessionModel) GetId() int {
    s.getIdCalls = append(s.getIdCalls, "")
    return correctSessionId
}

func (s *StubSessionModel) Get(sessionId int) error {
    s.getCalls = append(s.getCalls, sessionId)
    return nil
}

func (s *StubSessionModel) GetUserId() int {
    s.getUserIdCalls = append(s.getUserIdCalls, "")
    return 123
}
