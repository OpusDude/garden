package connection

import "code.cloudfoundry.org/garden"

type process struct {
	id string

	processInputStream *processStream
	status             chan garden.ProcessStatus
}

func newProcess(id string, processInputStream *processStream) *process {
	return &process{
		id:                 id,
		processInputStream: processInputStream,
		status:             make(chan garden.ProcessStatus, 1),
	}
}

func (p *process) ID() string {
	return p.id
}

func (p *process) ExitStatus() chan garden.ProcessStatus {
	return p.status
}

func (p *process) Wait() (int, error) {
	ret := <-p.status
	return ret.Code, ret.Err
}

func (p *process) SetTTY(tty garden.TTYSpec) error {
	return p.processInputStream.SetTTY(tty)
}

func (p *process) Signal(signal garden.Signal) error {
	return p.processInputStream.Signal(signal)
}

func (p *process) exited(exitStatus garden.ProcessStatus) {
	//the exited function should only be called once otherwise the
	//line below will block
	p.status <- exitStatus
}
