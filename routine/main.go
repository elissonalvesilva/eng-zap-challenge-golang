package routine

import (
	coleta "github.com/elissonalvesilva/eng-zap-challenge-golang/routine/coleta"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/routine/parser"
)

func Run() {
	coleta.InitColeta()
	coleta.Run()
	parser.Run()
}
