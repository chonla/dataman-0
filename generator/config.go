package generator

type config struct {
	Datasets string       `yml:"datasets"`
	Export   exportConfig `yml:"export"`
}

type exportConfig struct {
	Output  outputConfig   `yml:"output"`
	Columns []columnConfig `yml:"columns"`
}

type outputConfig struct {
	Count    int    `yml:"count"`
	To       string `yml:"to"`
	FileName string `yml:"filename"`
}

type columnConfig struct {
	Name  string `yml:"name"`
	Value string `yml:"value"`
	Type  string `yml:"type,omitempty"`
}
