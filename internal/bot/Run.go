package bot

import (
	"fmt"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"os"
	"os/signal"
	"syscall"

	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
)

func Run() {
	l := log.Logger{
		Level:  log.ErrorLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := khl.New(conf.Get().KhlBot.Token, plog.NewLogger(&l))

	RegisterHandlers(s)

	_ = s.Open()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the KHL session.
	_ = s.Close()
}
