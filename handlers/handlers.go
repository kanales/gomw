package handlers

import (
	"encoding/json"
	"gomw/models"
	"net/http"
	"os/exec"
	"runtime"
)

var RunCmd = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	response := models.CmdResponse{}
	if runtime.GOOS == "windows" {
		rw.WriteHeader(500)
		response.Error = "Can't execute on Windows"
		json.NewEncoder(rw).Encode(response)
		return
	}

	cmd := models.Command{}
	err := json.NewDecoder(r.Body).Decode(&cmd)

	if err != nil {
		rw.WriteHeader(400)
		response.Error = err.Error()

	} else if cmd.Command == nil {
		rw.WriteHeader(400)
		response.Error = "Missing field \"command\""

	} else {
		command := exec.Command(*cmd.Command, cmd.Args...)
		if cmd.Cwd != nil {
			command.Dir = *cmd.Cwd
		}

		out, err := command.Output()

		if err != nil {
			rw.WriteHeader(400)
			response.Error = err.Error()

			json.NewEncoder(rw).Encode(response)
			return
		}

		response.Output = string(out)
	}

	json.NewEncoder(rw).Encode(response)
})
