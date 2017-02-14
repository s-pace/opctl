package model

// static call graph; see https://en.wikipedia.org/wiki/Call_graph
type Scg struct {
	Container *ScgContainerCall `yaml:"container,omitempty"`
	Op        *ScgOpCall        `yaml:"op,omitempty"`
	Parallel  []*Scg            `yaml:"parallel,omitempty"`
	Serial    []*Scg            `yaml:"serial,omitempty"`
}

type ScgContainerCall struct {
	// each entry of cmd will be interpolated
	Cmd  []string               `yaml:"cmd,omitempty"`
	Dirs map[string]*ScgBinding `yaml:"dirs,omitempty"`

	// each env var value will be interpolated
	EnvVars map[string]string      `yaml:"envVars,omitempty"`
	Files   map[string]*ScgBinding `yaml:"files,omitempty"`
	Image   string                 `yaml:"image"`
	Sockets map[string]*ScgBinding `yaml:"sockets,omitempty"`
	StdErr  map[string]*ScgBinding `yaml:"stdErr,omitempty"`
	StdOut  map[string]*ScgBinding `yaml:"stdOut,omitempty"`
	WorkDir string                 `yaml:"workDir,omitempty"`
}

type ScgBinding struct {
	Bind string `yaml:"bind,omitempty"`
}

type ScgOpCall struct {
	Ref string `yaml:"ref"`
	// binds in scope variables to inputs of referenced op
	Inputs map[string]string `yaml:"inputs,omitempty"`
	// binds in scope variables to outputs of referenced op
	Outputs map[string]string `yaml:"outputs,omitempty"`
}
