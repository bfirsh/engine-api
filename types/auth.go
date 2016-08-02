package types

// AuthConfig contains authorization information for connecting to a Registry
// swagger:parameters postAuth
type AuthConfig struct {
	// Username
	// in: body
	Username string `json:"username,omitempty"`
	// Password
	// in: body
	Password string `json:"password,omitempty"`
	// in: body
	Auth string `json:"auth,omitempty"`

	// Email is an optional value associated with the username.
	// This field is deprecated and will be removed in a later
	// version of docker.
	// in: body
	Email string `json:"email,omitempty"`

	// The registry to authenticate with
	// in: body
	ServerAddress string `json:"serveraddress,omitempty"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	// in:body
	IdentityToken string `json:"identitytoken,omitempty"`

	// RegistryToken is a bearer token to be sent to a registry
	// in:body
	RegistryToken string `json:"registrytoken,omitempty"`
}
