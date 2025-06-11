package errors

import "runtime"

type FrameError interface {
	Frame() uintptr
}

func (e *fundamentalError) Frame() uintptr {
	return e.frame
}

func (e *msgError) Frame() uintptr {
	return e.frame
}

func (e *msgJoinedError) Frame() uintptr {
	return e.frame
}

func (e *noMsgError) Frame() uintptr {
	return e.frame
}

func (e *causeError) Frame() uintptr {
	return e.frame
}

func (e *wrapError) Frame() uintptr {
	return e.frame
}

func frames(extraSkip int) uintptr {
	var pcs [1]uintptr
	_ = runtime.Callers(3+extraSkip, pcs[:])
	return pcs[0]
}
