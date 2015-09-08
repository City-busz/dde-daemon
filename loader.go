package loader

import (
	"pkg.linuxdeepin.com/lib/log"
)

var logger = log.NewLogger("dde.daemon.loader")

type Module struct {
	Name   string
	Start  func()
	Stop   func()
	Enable bool
}

var modules = make([]*Module, 0)

func StartPlugin(plugin string) {
	for _, m := range modules {
		if m.Name == plugin {
			m.Start()
			return
		}
	}
}

func StopPlugin(plugin string) {
	for _, m := range modules {
		if m.Name == plugin {
			m.Stop()
			return
		}
	}
}

func Enable(name string, enable bool) {
	for _, m := range modules {
		if m.Name == name {
			m.Enable = enable
		}
	}
}

func Register(newModule *Module) {
	for _, m := range modules {
		if m.Name == newModule.Name {
			return
		}
	}
	if newModule.Start == nil || newModule.Stop == nil {
		logger.Error("can't register an incomplete module " + newModule.Name)
		return
	}
	modules = append(modules, newModule)
}

func Start() {
	for _, m := range modules {
		if !m.Enable {
			continue
		}

		func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("Start module", m.Name, "failed:", err)
				}
			}()
			m.Start()
		}()
	}
}
func Stop() {
	for _, m := range modules {
		func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("Stop module", m.Name, "failed:", err)
				}
			}()
			m.Stop()
		}()
	}
}
