package bot

import (
	"github.com/Liki4/KaiheilaBot/internal/handlers"
	"github.com/Liki4/KaiheilaBot/internal/handlers/ncm"
	"github.com/lonelyevil/khl"
)

func RegisterHandlers(s *khl.Session) {
	s.AddHandler(handlers.PingHandler)
	s.AddHandler(ncm.LoginHandler)
	s.AddHandler(ncm.SearchHandler)
}
