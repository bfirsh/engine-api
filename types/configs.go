package types

import (
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
)

// configs holds structs used for internal communication between the
// frontend (such as an http server) and the backend (such as the
// docker daemon).

// ContainerCreateConfig holds parameters for creating containers
type ContainerCreateConfig struct {
	// Assign the specified name to the container
	//
	// in: query
	// pattern: /?[a-zA-Z0-9_-]+
	Name             string `json:"name"`
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
	AdjustCPUShares  bool
}

// ContainerRmConfig holds arguments for the container remove
// swagger:parameters deleteContainers
// operation. This struct is used to tell the backend what operations
// to perform.
type ContainerRmConfig struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Kill the container before removing it
	//
	// in: query
	// default: false
	ForceRemove bool `json:"force"`
	// Remove volumes associated with this container
	//
	// in: query
	// default: false
	RemoveVolume bool `json:"v"`
	// Remove links associated with this container
	//
	// in: query
	// default: false
	RemoveLink bool `json:"link"`
}

// ContainerCommitConfig contains build configs for commit operation,
// and is used when making a commit with the current state of the container.
type ContainerCommitConfig struct {
	Pause   bool
	Repo    string
	Tag     string
	Author  string
	Comment string
	// merge container config into commit config before commit
	MergeConfigs bool
	Config       *container.Config
}

// ExecConfig is a small subset of the Config struct that holds the configuration
// for the exec feature of docker.
type ExecConfig struct {
	User         string   // User that will run the command
	Privileged   bool     // Is the container in privileged mode
	Tty          bool     // Attach standard streams to a tty.
	AttachStdin  bool     // Attach the standard input, makes possible user interaction
	AttachStderr bool     // Attach the standard error
	AttachStdout bool     // Attach the standard output
	Detach       bool     // Execute in detach mode
	DetachKeys   string   // Escape keys for detach
	Cmd          []string // Execution commands and args
}
