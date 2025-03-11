package integrator

type IntegrateTask struct {
	Begin    float64
	End      float64
	terminal bool
}

func GetTerminalTask() IntegrateTask {
	return IntegrateTask{
		Begin:    0,
		End:      0,
		terminal: true,
	}
}

func GetIntegrateTask(a, b float64) IntegrateTask {
	return IntegrateTask{
		Begin:    a,
		End:      b,
		terminal: false,
	}
}

func (task IntegrateTask) IsTerminal() bool {
	return task.terminal
}

func (task IntegrateTask) GetMiddle() float64 {
	return (task.Begin + task.End) / 2
}

func (task IntegrateTask) SplitTask() (IntegrateTask, IntegrateTask) {
	middle := task.GetMiddle()

	task1 := IntegrateTask{
		Begin:    task.Begin,
		End:      middle,
		terminal: false,
	}

	task2 := IntegrateTask{
		Begin:    middle,
		End:      task.End,
		terminal: false,
	}

	return task1, task2
}
