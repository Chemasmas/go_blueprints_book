package trace

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T) {
	//t.Error("No hemos escrito ningun test")
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("E retorno no debe de ser Nil")
	} else {
		tracer.Trace("Hola desde el Tracer")
			if buf.String()!="Hola desde el Tracer\n"  {
				t.Errorf("Trace No debe de escribir '%s'", buf.String() )
			}
	}
}

func TestOff(t *testing.T){
	var silentTracer Tracer = Off()
	silentTracer.Trace("algo")
}
