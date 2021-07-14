package models

type Command struct {
	Command *string  `json:"command"`
	Cwd     *string  `json:"cwd"`
	Args    []string `json:"args"`
}

type CmdResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}
