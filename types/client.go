package types

import (
	"bufio"
	"io"
	"net"

	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"github.com/docker/go-units"
)

// CheckpointCreateOptions holds parameters to create a checkpoint from a container
type CheckpointCreateOptions struct {
	CheckpointID string
	Exit         bool
}

// ContainerAttachOptions holds parameters to attach to a container.
// swagger:parameters postContainersAttach wsContainersAttach
type ContainerAttachOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
	// Return logs
	//
	// in: query
	// default: false
	Logs bool `json:"logs"`
	// Return a stream
	//
	// in: query
	// default: false
	Stream bool `json:"stream"`
	// If stream=true, attach to stdin
	//
	// in: query
	// default: false
	Stdin bool `json:"stdin"`
	// If logs=true, return stdout. If stream=true, attach to stdout.
	//
	// in: query
	// default: false
	Stdout bool `json:"stdout"`
	// If logs=true, return stderr. If stream=true, attach to stderr.
	//
	// in: query
	// default: false
	Stderr bool `json:"stderr"`
	// Override the key sequence for detaching a container. Format is a single
	// character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, ,
	// or _.
	DetachKeys string `json:"detachKeys"`
}

// ContainerCommitOptions holds parameters to commit changes into a container.
// TODO: dedupe with ContainerCommitConfig
// swagger:parameters postCommit
type ContainerCommitOptions struct {
	// TODO: reference is passed as separate repo/tag parameters
	Reference string
	// Commit message
	// in: query
	Comment string `json:"comment"`
	// Author of image (e.g., "John Hannibal Smith <hannibal@a-team.com>")
	// in: query
	Author string `json:"author"`
	// Dockerfile instructions to apply while committing
	// in: query
	Changes []string `json:"changes"`
	// Whether to pause the container before committing
	// in: query
	// default: false
	Pause bool
	// The container's configuration
	// in: body
	Config *container.Config
}

// ContainerExecInspect holds information returned by exec inspect.
type ContainerExecInspect struct {
	ExecID      string
	ContainerID string
	Running     bool
	ExitCode    int
}

// ContainerListOptions holds parameters to list containers with.
// swagger:parameters getContainers
type ContainerListOptions struct {
	// Show all containers.
	//
	// in: query
	All bool `json:"all"`

	// Show the size of containers
	Size bool `json:"size"`

	// Show only containers created since given ID, include non-running ones.
	//
	// in: query
	Since string `json:"since"`

	// Show only containers created before given ID, include non-running ones.
	//
	// in: query
	Before string `json:"before"`

	// Show <limit> last created containers, include non-running ones.
	//
	// in: query
	Limit int `json:"limit"`

	// A JSON encoded value of the filters (a map[string][]string) to process
	// on the containers list.
	//
	// in: query
	// type: string
	Filter filters.Args `json:"filters,string"`
}

// ContainerStatsOptions holds parameters for container stats
// swagger:parameters getContainersStats
type ContainerStatsOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Stream stats
	//
	// in: query
	// default: true
	Stream string `json:"stream"`
}

// ContainerLogsOptions holds parameters to filter logs with.
// swagger:parameters getContainersLogs
type ContainerLogsOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
	// Show stdout log
	//
	// in: query
	// default: false
	ShowStdout bool `json:"stdout"`
	// Show stderr log
	//
	// in: query
	// default: false
	ShowStderr bool `json:"stdin"`
	// Only return logs after this time, as a UNIX timestamp
	//
	// in: query
	// default: 0
	Since string `json:"since"`
	// Print timestamps for every line
	//
	// in: query
	// default: false
	Timestamps bool `json:"timestamps"`
	// Return a stream
	//
	// in: query
	// default: false
	Follow bool `json:"follow"`
	// Output specified number of lines at the end of logs, as "all" or a number
	// of lines
	//
	// in: query
	// default: all
	Tail string `json:"tail"`
	// Show extra details provided to logs
	//
	// in: query
	// default: false
	Details bool `json:"details"`
}

// ContainerExportOptions holds parameters for exporting containers.
// swagger:parameters getContainersExport
type ContainerExportOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// ContainerUpdateOptions holds parameters for updating containers.
type ContainerUpdateOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// ContainerCreateOptions holds parameters for creating containers.
// TODO: dedupe with ContainerCreateConfig
// swagger:parameters postContainersCreate
type ContainerCreateOptions struct {
	// Assign the specified name to the container
	//
	// in: query
	// pattern: /?[a-zA-Z0-9_-]+
	Name string `json:"name"`
}

// ContainerRemoveOptions holds parameters to remove containers.
// swagger:parameters deleteContainers
type ContainerRemoveOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Remove volumes associated with this container
	//
	// in: query
	// default: false
	RemoveVolumes bool `json:"v"`
	// Remove links associated with this container
	//
	// in: query
	// default: false
	RemoveLinks bool `json:"link"`
	// Kill the container before removing it
	//
	// in: query
	// default: false
	Force bool `json:"force"`
}

// ContainerStartOptions holds parameters to start containers.
// swagger:parameters postContainersStart
type ContainerStartOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	CheckpointID string
}

// ContainerStopOptions holds parameters for POST /containers/{id}/stop
// swagger:parameters postContainersStop
type ContainerStopOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _.
	//
	// in:query
	DetachKeys string `json:"detachKeys"`
}

// ContainerKillOptions holds parameters for POST /containers/{id}/kill
// swagger:parameters postContainersKill
type ContainerKillOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Signal to send to the container: integer or string like SIGINT. When not
	// set, SIGKILL is assumed and the call waits for the container to exit.
	//
	// in:query
	Signal string `json:"signal"`
}

// ContainerRestartOptions holds parameters for POST /containers/{id}/restart
// swagger:parameters postContainersRestart
type ContainerRestartOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Number of seconds to wait before killing container
	//
	// in:query
	Timeout int `json:"t"`
}

// ContainerPauseOptions holds parameters for POST /containers/{id}/pause
// swagger:parameters postContainersPause
type ContainerPauseOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// ContainerUnpauseOptions holds parameters for POST /containers/{id}/unpause
// swagger:parameters postContainersUnpause
type ContainerUnpauseOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// ContainerTopOptions holds parameters for GET /containers/{id}/top
// swagger:parameters getContainersTop
type ContainerTopOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// Arguments to pass to ps (e.g., aux)
	//
	// in: query
	// default: -ef
	PsArgs string `json:"ps_args"`
}

// ContainerRenameOptions holds parameters for POST /containers/{id}/rename
// swagger:parameters getContainersRename
type ContainerRenameOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// The new name for the container
	//
	// in: query
	// required: true
	Name string `json:"name"`
}

// CopyToContainerOptions holds information
// about files to copy into a container
type CopyToContainerOptions struct {
	AllowOverwriteDirWithFile bool
}

// EventsOptions hold parameters to filter events with.
// swagger:parameters getEvents
type EventsOptions struct {
	// Show events since timestamp, for polling
	// in: query
	Since string `json:"since"`
	// Show events until timestamp, for polling
	// in: query
	Until string `json:"until"`
	// Filters to process the events by.
	Filters filters.Args `json:"filters"`
}

// NetworkListOptions holds parameters to filter the list of networks with.
// swagger:parameters getNetworksList
type NetworkListOptions struct {
	// JSON encoded network list filter. The filter value is one of:
	// - driver=<driver-name> Matches a network’s driver.
	// - id=<network-id> Matches all or part of a network id.
	// - label=<key> or label=<key>=<value> of a network label.
	// - name=<network-name> Matches all or part of a network name.
	// - type=["custom"|"builtin"] Filters networks by type. The custom keyword
	//   returns all user-defined networks.
	// in: query
	Filters filters.Args `json:"filters"`
}

// NetworkGetOptions contains the options for GET /networks/{id}
// swagger:parameters getNetwork
type NetworkGetOptions struct {
	// The ID of the network
	// in: path
	ID string `json:"id"`
}

// HijackedResponse holds connection information for a hijacked request.
type HijackedResponse struct {
	Conn   net.Conn
	Reader *bufio.Reader
}

// Close closes the hijacked connection and reader.
func (h *HijackedResponse) Close() {
	h.Conn.Close()
}

// CloseWriter is an interface that implements structs
// that close input streams to prevent from writing.
type CloseWriter interface {
	CloseWrite() error
}

// CloseWrite closes a readWriter for writing.
func (h *HijackedResponse) CloseWrite() error {
	if conn, ok := h.Conn.(CloseWriter); ok {
		return conn.CloseWrite()
	}
	return nil
}

// ImageBuildOptions holds the information
// necessary to build images.
// TODO parameters postBuild
type ImageBuildOptions struct {
	// Path within the build context to the Dockerfile. This is ignored if remote
	// is specified and points to an individual filename.
	// in: query
	Dockerfile string `json:"dockerfile"`
	// A name and optional tag to apply to the image in the name:tag format. If
	// you omit the tag the default latest value is assumed.
	// in: query
	Tags []string `json:"t"`
	// Suppress verbose build output.
	// in: query
	SuppressOutput bool `json:"q"`
	// A Git repository URI or HTTP/HTTPS URI build source. If the URI specifies
	// a filename, the file’s contents are placed into a file called Dockerfile.
	// in: query
	RemoteContext string `json:"remote"`
	// Do not use the cache when building the image.
	// in: query
	NoCache bool `json:"nocache"`
	// Remove intermediate containers after a successful build.
	// in: query
	// default: true
	Remove bool `json:"rm"`
	// Always remove intermediate containers. Enabling this also enables rm.
	// in: query
	ForceRemove bool `json:"forcerm"`
	// Attempt to pull the image even if an older image exists locally.
	// in: query
	PullParent bool `json:"pull"`
	// TODO
	Isolation container.Isolation
	// CPUs in which to allow execution (e.g., "0-3", "0,1").
	// in: query
	CPUSetCPUs string `json:"cpusetcpus"`
	// in: query
	CPUSetMems string `json:"cpusetmems"`
	// CPU shares (relative weight).
	// in: query
	CPUShares int64 `json:"cpushares"`
	// Microseconds of CPU time that the container can get in a CPU period.
	// in: query
	CPUQuota int64 `json:"cpuquota"`
	// The length of a CPU period in microseconds.
	// in: query
	CPUPeriod int64 `json:"cpuperiod"`
	// Set memory limit for build.
	// in: query
	Memory int64 `json:"memory"`
	// Total memory (memory + swap), -1 to enable unlimited swap.
	// in: query
	MemorySwap int64 `json:"memoryswap"`
	// TODO
	CgroupParent string
	// Size of /dev/shm in bytes. The size must be greater than 0. If omitted the
	// system uses 64MB.
	// in: query
	ShmSize int64 `json:"shmsize"`
	// TODO
	Ulimits []*units.Ulimit
	// JSON map of string pairs for build-time variables. Users pass these values
	// at build-time. Docker uses the buildargs as the environment context for
	// command(s) run via the Dockerfile’s RUN instruction or for variable
	// expansion in other Dockerfile instructions. This is not meant for passing
	// secret values.
	// TODO
	BuildArgs map[string]string `json:"buildargs"`
	// TODO
	AuthConfigs map[string]AuthConfig
	Context     io.Reader
	// JSON map of string pairs for labels to set on the image.
	// in: query
	Labels map[string]string `json:"labels"`
}

// ImageBuildResponse holds information
// returned by a server after building
// an image.
type ImageBuildResponse struct {
	Body   io.ReadCloser
	OSType string
}

// ImageCreateOptions holds information to create images.
// swagger:parameters postImageCreate
type ImageCreateOptions struct {
	// Name of the image to pull. The name may include a tag or digest. This
	// parameter may only be used when pulling an image.
	// in: query
	FromImage string `json:"fromImage"`
	// Source to import. The value may be a URL from which the image can be
	// retrieved or - to read the image from the request body. This parameter may
	// only be used when importing an image.
	// in: query
	FromSrc string `json:"fromSrc"`
	// Repository name given to an image when it is imported. The repo may
	// include a tag. This parameter may only be used when importing an image.
	// in: query
	Repo string `json:"repo"`
	// Tag or digest
	// in: query
	Tag string `json:"tag"`
	// The base64 encoded credentials for the registry
	// in: header
	RegistryAuth string `json:"X-Registry-Auth"`
}

// ImageImportSource holds source information for ImageImport
type ImageImportSource struct {
	Source     io.Reader // Source is the data to send to the server to create this image from (mutually exclusive with SourceName)
	SourceName string    // SourceName is the name of the image to pull (mutually exclusive with Source)
}

// ImageImportOptions holds information to import images from the client host.
type ImageImportOptions struct {
	Tag     string   // Tag is the name to tag this image with. This attribute is deprecated.
	Message string   // Message is the message to tag the image with
	Changes []string // Changes are the raw changes to apply to this image
}

// ImageListOptions holds parameters to filter the list of images with.
// swagger:parameters getImages
type ImageListOptions struct {
	// TODO
	MatchName string
	// Include images without tags
	// in: query
	// default: false
	All bool `json:"all"`
	// Filters to process the image list by. Available filters:
	// - dangling=true
	// - before=<image-name>[:<tag>], <image id> or <image@digest>
	// - since=<image-name>[:<tag>], <image id> or <image@digest>
	Filters filters.Args `json:"filters"`
}

// ImageHistoryOptions holds parameters for GET /images/{name}/history
// swagger:parameters getImagesHistory
type ImageHistoryOptions struct {
	// Name of the image
	// in: path
	Name string `json:"name"`
}

// ImageTagOptions holds paramters for POST /images/{name}/tag
// swagger:parameters postImagesTag
type ImageTagOptions struct {
	// Name of the image
	// in: path
	Name string `json:"name"`
	// The repository to tag in
	// in: query
	Repo string `json:"repo"`
	// The new tag
	// in: query
	Tag string `json:"tag"`
}

// ImageLoadResponse returns information to the client about a load process.
type ImageLoadResponse struct {
	// Body must be closed to avoid a resource leak
	Body io.ReadCloser
	JSON bool
}

// ImagePullOptions holds information to pull images.
type ImagePullOptions struct {
	All           bool
	RegistryAuth  string // RegistryAuth is the base64 encoded credentials for the registry
	PrivilegeFunc RequestPrivilegeFunc
}

// RequestPrivilegeFunc is a function interface that
// clients can supply to retry operations after
// getting an authorization error.
// This function returns the registry authentication
// header value in base 64 format, or an error
// if the privilege request fails.
type RequestPrivilegeFunc func() (string, error)

//ImagePushOptions holds information to push images.
type ImagePushOptions ImagePullOptions

// ImageRemoveOptions holds parameters to remove images.
// swagger:parameters deleteImages
type ImageRemoveOptions struct {
	// The name of the image
	// in: path
	Name string `json:"name"`
	// in: query
	// default: false
	Force bool `json:"force"`
	// in: query
	// default: false
	PruneChildren bool `json:"noprune"`
}

// ImageSearchOptions holds parameters to search images with.
// TODO:parameters getImagesSearch
type ImageSearchOptions struct {
	// The base64 encoded credentials for the registry
	// in: header
	RegistryAuth string `json:"X-Registry-Auth"`
	// Filters to process on the image list
	// in: query
	Filters filters.Args `json:"filters"`
	// Maximum returned search results
	// in: query
	Limit         int `json:"limit"`
	PrivilegeFunc RequestPrivilegeFunc
}

// ImageGetOptions holds parameters for GET /images/{name}/get
// swagger:parameters getImagesGet
type ImageGetOptions struct {
	// The image name
	// in: query
	Name string `json:"name"`
}

// ResizeOptions holds parameters to resize a tty.
// It can be used to resize container ttys and
// exec process ttys too.
// swagger:parameters postContainersResize
type ResizeOptions struct {
	// The container ID or name
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// The height of the tty
	//
	// in: query
	Height int `json:"h"`

	// The width of the tty
	//
	// in: query
	Width int `json:"w"`
}

// VersionResponse holds version information for the client and the server
type VersionResponse struct {
	Client *Version
	Server *Version
}

// ServerOK returns true when the client could connect to the docker server
// and parse the information received. It returns false otherwise.
func (v VersionResponse) ServerOK() bool {
	return v.Server != nil
}

// NodeListOptions holds parameters to list nodes with.
// swagger:parameters getNodes
type NodeListOptions struct {
	// A JSON encoded value of the filters (a map[string][]string) to process
	// on the containers list.
	//
	// in: query
	// type: string
	Filter filters.Args `json:"filters,string"`
}

// NodeUpdateOptions contains the parameters for updating nodes
// TODO:parameters updateNode
type NodeUpdateOptions struct {
	// The ID of the node
	// in: path
	ID string `json:"id"`
	// in: body
	Body *swarm.NodeSpec
}

// NodeRemoveOptions holds parameters to remove nodes with.
// swagger:parameters removeNode
type NodeRemoveOptions struct {
	// The ID of the node
	// in: path
	ID string `json:"id"`
}

// ServiceCreateOptions contains the options to use when creating a service.
// TODO:parameters createService
type ServiceCreateOptions struct {
	// The base64 encoded registry authorization credentials to use when updating the service.
	// in: header
	EncodedRegistryAuth string `json:"X-Registry-Auth"`
	// in: body
	Body *swarm.ServiceSpec
}

// ServiceCreateResponse contains the information returned to a client
// on the  creation of a new service.
// swagger:model
type ServiceCreateResponse struct {
	// The ID of the service created
	ID string `json:"Id"`
}

// ServiceUpdateOptions contains the options to be used for updating services.
// TODO:parameters updateService
type ServiceUpdateOptions struct {
	// The ID or name of the service
	// in: path
	ID string `json:"id"`
	// The base64 encoded registry authorization credentials to use when updating the service.
	// in: header
	EncodedRegistryAuth string `json:"X-Registry-Auth"`
	// in: body
	Body *swarm.ServiceSpec
	// TODO(stevvooe): Consider moving the version parameter of ServiceUpdate
	// into this field. While it does open API users up to racy writes, most
	// users may not need that level of consistency in practice.
}

// ServiceRemoveParameters are parameters for removing services
// swagger:parameters removeService
type ServiceRemoveParameters struct {
	// The ID or name of the service
	// in: path
	ID string `json:"id"`
}

// ServiceListOptions are parameters for listing services
// swagger:parameters getServices
type ServiceListOptions struct {
	// A JSON encoded value of the filters (a map[string][]string) to process
	// on the service list.
	//
	// in: query
	// type: string
	Filter filters.Args `json:"filters,string"`
}

// TaskListOptions holds parameters for listing tasks
// swagger:parameters getTasks
type TaskListOptions struct {
	// A JSON encoded value of the filters (a map[string][]string) to process
	// on the tasks
	//
	// in: query
	// type: string
	Filter filters.Args `json:"filters,string"`
}

// VolumeListOptions holds parameters for listing volmes
// swagger:parameters getVolumesList
type VolumeListOptions struct {
	// A JSON encoded value of the filters (a map[string][]string) to process
	// on the tasks
	//
	// in: query
	// type: string
	Filter filters.Args `json:"filters,string"`
}
