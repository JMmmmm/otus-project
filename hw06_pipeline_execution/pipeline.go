package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	result := in
	for _, stage := range stages {
		if stage == nil {
			continue
		}
		result = stage(wrap(result, done))
	}

	return result
}

func wrap(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}

				out <- value
			}
		}
	}()

	return out
}
