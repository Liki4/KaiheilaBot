package bot

import (
	"github.com/Liki4/KaiheilaBot/internal/Handlers"
	"github.com/lonelyevil/khl"
)

func RegisterHandlers(s *khl.Session) {
	s.AddHandler(Handlers.PingHandler)
	s.AddHandler(Handlers.RobotHandler)
}
