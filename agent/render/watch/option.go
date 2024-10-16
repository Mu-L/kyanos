package watch

import "strings"

type WatchOptions struct {
	WideOutput   bool
	StaticRecord bool
	Opts         string
	DebugOutput  bool
}

func (w *WatchOptions) Init() {
	if w.Opts != "" {
		if strings.Contains(w.Opts, "wide") {
			w.WideOutput = true
		}
	}
}
