package user

import "errors"

var sessionId = 5000
var sessionIdToUserId = map[int]int{}

type StubSessionModel struct {
    sessionId int
    userId int

    createCalls []int
    getIdCalls []string
    getUserIdCalls []string
    getCalls []int
    deleteCalls []string
}

func (s *StubSessionModel) Create(userId int) error {
    s.createCalls = append(s.createCalls, userId)
    sessionId += 1
    s.sessionId = sessionId
    s.userId = userId
    sessionIdToUserId[sessionId] = userId

    return nil
}

func (s *StubSessionModel) GetId() int {
    s.getIdCalls = append(s.getIdCalls, "")

    return s.sessionId
}

func (s *StubSessionModel) Get(sessionId int) error {
    s.getCalls = append(s.getCalls, sessionId)

    userId, ok := sessionIdToUserId[sessionId]
    if !ok {
        return errors.New("session not exists")
    }
    s.sessionId = sessionId
    s.userId = userId
    return nil
}

func (s *StubSessionModel) GetUserId() int {
    s.getUserIdCalls = append(s.getUserIdCalls, "")

    return s.userId
}

func (s *StubSessionModel) Delete() {
    s.deleteCalls = append(s.deleteCalls, "")
}
