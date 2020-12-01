package runtime

import (
	"log"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	p := &processor{}
	p.tryThisFunc()
}

type processor struct {
}

func (p *processor) tryThisFunc() {
	programCounter, fileName, lineNumber, ok := runtime.Caller(0)
	log.Printf("programCounter: %v, fileName: %v lineNumber: %v, ok: %v", programCounter, fileName, lineNumber, ok)
	funcDetails := runtime.FuncForPC(programCounter)
	log.Printf("funcName: %v funcEntry: %v", funcDetails.Name(), funcDetails.Entry())
}
