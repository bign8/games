package room

// MsgHandler is a function that can handle messages
type MsgHandler func([]byte)

// Manager creates a new set of rooms for a communication channel
type Manager struct {
	rooms map[byte]MsgHandler
}

// NewManager creates a new room manager
func NewManager() *Manager {
	return &Manager{
		rooms: make(map[byte]MsgHandler),
	}
}

// Send broadcasts a room message
func (m *Manager) Send(byte, []byte) error {
	return nil
}

// Listen adds a MsgHandler to aparticular room
func (m *Manager) Listen(byte, MsgHandler) error {
	return nil
}
