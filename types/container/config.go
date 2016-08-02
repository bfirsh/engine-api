package container

import (
	"time"

	"github.com/docker/engine-api/types/strslice"
	"github.com/docker/go-connections/nat"
)

// HealthConfig holds configuration settings for the HEALTHCHECK feature.
type HealthConfig struct {
	// Test is the test to perform to check that the container is healthy.
	// An empty slice means to inherit the default.
	// The options are:
	// {} : inherit healthcheck
	// {"NONE"} : disable healthcheck
	// {"CMD", args...} : exec arguments directly
	// {"CMD-SHELL", command} : run command with system's default shell
	Test []string `json:",omitempty"`

	// Zero means to inherit. Durations are expressed as integer nanoseconds.
	Interval time.Duration `json:",omitempty"` // Interval is the time to wait between checks.
	Timeout  time.Duration `json:",omitempty"` // Timeout is the time to wait before considering the check to have hung.

	// Retries is the number of consecutive failures needed to consider a container as unhealthy.
	// Zero means inherit.
	Retries int `json:",omitempty"`
}

// Config contains the configuration for a container.
//
// It should hold only portable information about the container.
// Here, "portable" means "independent from the host we are running on".
// Non-portable information *should* appear in HostConfig.
//
// All fields added to this struct must be marked `omitempty` to keep getting
// predictable hashes from the old `v1Compatibility` configuration.
//
type Config struct {
	// The hostname for this container
	Hostname string
	// The domain name for this container
	Domainname string
	// The user that commands will run as inside the container
	User string
	// Attach standard input to allow user interaction
	AttachStdin bool
	// Attach standard output
	AttachStdout bool
	// Attach standard error
	AttachStderr bool
	// List of exposed ports
	ExposedPorts map[nat.Port]struct{} `json:",omitempty"`
	// Attach standard streams to a tty, including stdin if it is not closed.
	Tty bool
	// Open stdin
	OpenStdin bool
	// If true, close stdin after one attached client disconnects
	StdinOnce bool
	// List of environment variables to set in the container, in the format key=value
	Env []string
	// Command to run when starting the container
	Cmd strslice.StrSlice
	// Healthcheck describes how to check the container is healthy
	Healthcheck *HealthConfig `json:",omitempty"`
	// True if command is already escaped (Windows specific)
	ArgsEscaped bool `json:",omitempty"`
	// Name of the image as it was passed by the operator (eg. could be symbolic)
	Image string
	// List of volumes (mounts) used for the container
	Volumes map[string]struct{}
	// Current directory (PWD) in the command will be launched
	WorkingDir string
	// Entrypoint to run when starting the container
	Entrypoint strslice.StrSlice
	// Disable networking for this container
	NetworkDisabled bool `json:",omitempty"`
	// MAC address of the container
	MacAddress string `json:",omitempty"`
	// ONBUILD metadata that were defined on the image Dockerfile
	OnBuild []string
	// List of labels set to this container
	Labels map[string]string
	// Signal to stop a container
	StopSignal string `json:",omitempty"`
	// Timeout (in seconds) to stop a container
	StopTimeout *int `json:",omitempty"`
	// Shell for shell-form of RUN, CMD, ENTRYPOINT
	Shell strslice.StrSlice `json:",omitempty"`
}
