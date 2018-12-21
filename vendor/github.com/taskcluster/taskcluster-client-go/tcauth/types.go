// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcauth

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Response for a request to get access to an S3 bucket.
	//
	// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#
	AWSS3CredentialsResponse struct {

		// Temporary STS credentials for use when operating on S3
		//
		// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/credentials
		Credentials TemporarySecurityCredentials `json:"credentials"`

		// Date and time of when the temporary credentials expires.
		//
		// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`
	}

	// Response from a request to authenticate a hawk request.
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[1]
	AuthenticationFailedResponse struct {

		// Message saying why the authentication failed.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[1]/properties/message
		Message string `json:"message"`

		// The kind of response, `auth-failed` or `auth-success`.
		//
		// Possible values:
		//   * "auth-failed"
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[1]/properties/status
		Status string `json:"status"`
	}

	// Response from a request to authenticate a hawk request.
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]
	AuthenticationSuccessfulResponse struct {

		// The `clientId` that made this request.  This may be the `id` supplied in
		// the Authorization header, or in the case of a named temporary credential
		// may be embedded in the payload.  In any case, this clientId can be used
		// for logging, auditing, and identifying the credential but **must** not be
		// used for access control.  That's what scopes are for.
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/clientId
		ClientID string `json:"clientId"`

		// The expiration time for the credentials used to make this request.
		// This should be treated as the latest time at which the authorization
		// is valid.  For most cases, where the access being authorized occurs
		// immediately, this field can be ignored, as the value will always be
		// in the future if the status is `auth-success`.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/expires
		Expires tcclient.Time `json:"expires"`

		// Payload as extracted from `Authentication` header. This property is
		// only present if a hash is available. You are not required to validate
		// this hash, but if you do, please check `scheme` to ensure that it's
		// on a scheme you support.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/hash
		Hash json.RawMessage `json:"hash,omitempty"`

		// Authentication scheme the client used. Generally, you don't need to
		// read this property unless `hash` is provided and you want to validate
		// the payload hash. Additional values may be added in the future.
		//
		// Possible values:
		//   * "hawk"
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/scheme
		Scheme string `json:"scheme"`

		// List of scopes the client is authorized to access.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/scopes
		Scopes []string `json:"scopes"`

		// The kind of response, `auth-failed` or `auth-success`.
		//
		// Possible values:
		//   * "auth-success"
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#/oneOf[0]/properties/status
		Status string `json:"status"`
	}

	// Response to a request for an Shared-Access-Signature to access an Azure
	// Blob Storage container.
	//
	// See https://schemas.taskcluster.net/auth/v1/azure-container-response.json#
	AzureBlobSharedAccessSignature struct {

		// Date and time of when the Shared-Access-Signature expires.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-container-response.json#/properties/expiry
		Expiry tcclient.Time `json:"expiry"`

		// Shared-Access-Signature string. This is the querystring parameters to
		// be appened after `?` or `&` depending on whether or not a querystring is
		// already present in the URL.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-container-response.json#/properties/sas
		Sas string `json:"sas"`
	}

	// A list of Azure accounts managed by taskcluster-auth
	//
	// See https://schemas.taskcluster.net/auth/v1/azure-account-list-response.json#
	AzureListAccountResponse struct {

		// A list of accountIds that are managed by auth. These are
		// the accounts that can have SAS credentials fetched for tables
		// within them.
		//
		// Array items:
		// See https://schemas.taskcluster.net/auth/v1/azure-account-list-response.json#/properties/accounts/items
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-account-list-response.json#/properties/accounts
		Accounts []string `json:"accounts"`
	}

	// A list of Azure containers in an account
	//
	// See https://schemas.taskcluster.net/auth/v1/azure-container-list-response.json#
	AzureListContainersResponse struct {

		// A list of containers that are in an account.  Credentials are available for
		// these containers from the `azureBlobSAS` method.
		//
		// Array items:
		// See https://schemas.taskcluster.net/auth/v1/azure-container-list-response.json#/properties/containers/items
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-container-list-response.json#/properties/containers
		Containers []string `json:"containers"`

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of containers.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called this method with `continuationToken`
		// until you get a result without a `continuationToken`.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-container-list-response.json#/properties/continuationToken
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// A list of Azure tables in an account
	//
	// See https://schemas.taskcluster.net/auth/v1/azure-table-list-response.json#
	AzureListTableResponse struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of tables.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `azureAccountTables` with `continuationToken`
		// until you get a result without a `continuationToken`.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-table-list-response.json#/properties/continuationToken
		ContinuationToken string `json:"continuationToken,omitempty"`

		// A list of tables that are in an account. These are
		// the tables that can have SAS credentials fetched for them.
		//
		// Array items:
		// See https://schemas.taskcluster.net/auth/v1/azure-table-list-response.json#/properties/tables/items
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-table-list-response.json#/properties/tables
		Tables []string `json:"tables"`
	}

	// Response to a request for an Shared-Access-Signature to access and Azure
	// Table Storage table.
	//
	// See https://schemas.taskcluster.net/auth/v1/azure-table-access-response.json#
	AzureTableSharedAccessSignature struct {

		// Date and time of when the Shared-Access-Signature expires.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-table-access-response.json#/properties/expiry
		Expiry tcclient.Time `json:"expiry"`

		// Shared-Access-Signature string. This is the querystring parameters to
		// be appened after `?` or `&` depending on whether or not a querystring is
		// already present in the URL.
		//
		// See https://schemas.taskcluster.net/auth/v1/azure-table-access-response.json#/properties/sas
		Sas string `json:"sas"`
	}

	// Properties to create a client.
	//
	// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#
	CreateClientRequest struct {

		// If `true`, the service may delete this client after it has expired.  If
		// `false` (the default), the client will remain after expiration, although
		// it cannot be used for authentication in that state.
		//
		// Default:    false
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#/properties/deleteOnExpiration
		DeleteOnExpiration bool `json:"deleteOnExpiration,omitempty"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#/properties/description
		Description string `json:"description"`

		// Date and time where the clients access is set to expire
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// List of scopes the client has (unexpanded).
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-request.json#/properties/scopes
		Scopes []string `json:"scopes,omitempty"`
	}

	// All details about a client including the `accessToken`
	//
	// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#
	CreateClientResponse struct {

		// AccessToken used for authenticating requests, you should store this
		// you won't be able to retrive it again!
		//
		// Syntax:     ^[a-zA-Z0-9_-]{22,66}$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/accessToken
		AccessToken string `json:"accessToken"`

		// ClientId of the client
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/clientId
		ClientID string `json:"clientId"`

		// Date and time when this client was created
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/created
		Created tcclient.Time `json:"created"`

		// If `true`, the service may delete this client after it has expired.  If
		// `false`, the client will remain after expiration, although it cannot be
		// used for authentication in that state.
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/deleteOnExpiration
		DeleteOnExpiration bool `json:"deleteOnExpiration"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/description
		Description string `json:"description"`

		// If true, this client is disabled and cannot be used.  This usually occurs when the
		// scopes available to the user owning the client no longer satisfy the client.
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/disabled
		Disabled bool `json:"disabled"`

		// List of scopes granted to this client by matching roles, including the
		// client's scopes and the implicit role `client-id:<clientId>`.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/expandedScopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/expandedScopes
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time where the clients access is set to expire
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// Date of last time this client was used. Will only be updated every 6 hours
		// or so this may be off by up-to 6 hours. But it still gives a solid hint
		// as to whether or not this client is in use.
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/lastDateUsed
		LastDateUsed tcclient.Time `json:"lastDateUsed"`

		// Date and time of last modification
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/lastModified
		LastModified tcclient.Time `json:"lastModified"`

		// Date and time of when the `accessToken` was reset last time.
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/lastRotated
		LastRotated tcclient.Time `json:"lastRotated"`

		// List of scopes the client has (unexpanded).
		//
		// Default:    []
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/create-client-response.json#/properties/scopes
		Scopes []string `json:"scopes"`
	}

	// Data to create or update a role.
	//
	// See https://schemas.taskcluster.net/auth/v1/create-role-request.json#
	CreateRoleRequest struct {

		// Description of what this role is used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		//
		// See https://schemas.taskcluster.net/auth/v1/create-role-request.json#/properties/description
		Description string `json:"description"`

		// List of scopes the role grants access to.  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/create-role-request.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/create-role-request.json#/properties/scopes
		Scopes []string `json:"scopes"`
	}

	// Access credentials and urls for the Sentry project.
	// Credentials will expire in 24-48 hours, you should refresh them within
	// 24 hours.
	//
	// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/dsn
	Dsn struct {

		// Access credential and URL for public error reports.
		// These credentials can be used for up-to 24 hours.
		// This is for use in client-side applications only.
		//
		// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/dsn/properties/public
		Public string `json:"public"`

		// Access credential and URL for private error reports.
		// These credentials can be used for up-to 24 hours.
		// This is for use in serser-side applications and should **not** be
		// leaked.
		//
		// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/dsn/properties/secret
		Secret string `json:"secret"`
	}

	// Get all details about a client, useful for tools modifying a client
	//
	// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#
	GetClientResponse struct {

		// ClientId of the client scopes is requested about
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/clientId
		ClientID string `json:"clientId"`

		// Date and time when this client was created
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/created
		Created tcclient.Time `json:"created"`

		// If `true`, the service may delete this client after it has expired.  If
		// `false`, the client will remain after expiration, although it cannot be
		// used for authentication in that state.
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/deleteOnExpiration
		DeleteOnExpiration bool `json:"deleteOnExpiration"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/description
		Description string `json:"description"`

		// If true, this client is disabled and cannot be used.  This usually occurs when the
		// scopes available to the user owning the client no longer satisfy the client.
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/disabled
		Disabled bool `json:"disabled"`

		// List of scopes granted to this client by matching roles.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Scope that client is granted by a role
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/expandedScopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/expandedScopes
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time where the clients access is set to expire
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// Date of last time this client was used. Will only be updated every 6 hours
		// or so this may be off by up-to 6 hours. But it still gives a solid hint
		// as to whether or not this client is in use.
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/lastDateUsed
		LastDateUsed tcclient.Time `json:"lastDateUsed"`

		// Date and time of last modification
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/lastModified
		LastModified tcclient.Time `json:"lastModified"`

		// Date and time of when the `accessToken` was reset last time.
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/lastRotated
		LastRotated tcclient.Time `json:"lastRotated"`

		// List of scopes the client has (unexpanded).  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/get-client-response.json#/properties/scopes
		Scopes []string `json:"scopes"`
	}

	// Get all details about a role
	//
	// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#
	GetRoleResponse struct {

		// Date and time when this role was created
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/created
		Created tcclient.Time `json:"created"`

		// Description of what this role is used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/description
		Description string `json:"description"`

		// List of scopes granted anyone who assumes this role, including anything
		// granted by roles that can be assumed when you have this role.
		// Hence, this includes any scopes in-directly granted as well.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/expandedScopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/expandedScopes
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time of last modification
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/lastModified
		LastModified tcclient.Time `json:"lastModified"`

		// roleId of the role requested
		//
		// Syntax:     ^[\x20-\x7e]+$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/roleId
		RoleID string `json:"roleId"`

		// List of scopes the role grants access to.  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*(?<!\*\*)$
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/get-role-response.json#/properties/scopes
		Scopes []string `json:"scopes"`
	}

	// Request to authenticate a hawk request.
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#
	HawkSignatureAuthenticationRequest struct {

		// Authorization header, **must** only be specified if request being
		// authenticated has a `Authorization` header.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/authorization
		Authorization string `json:"authorization,omitempty"`

		// Host for which the request came in, this is typically the `Host` header
		// excluding the port if any.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/host
		Host string `json:"host"`

		// HTTP method of the request being authenticated.
		//
		// Possible values:
		//   * "get"
		//   * "post"
		//   * "put"
		//   * "head"
		//   * "delete"
		//   * "options"
		//   * "trace"
		//   * "copy"
		//   * "lock"
		//   * "mkcol"
		//   * "move"
		//   * "purge"
		//   * "propfind"
		//   * "proppatch"
		//   * "unlock"
		//   * "report"
		//   * "mkactivity"
		//   * "checkout"
		//   * "merge"
		//   * "m-search"
		//   * "notify"
		//   * "subscribe"
		//   * "unsubscribe"
		//   * "patch"
		//   * "search"
		//   * "connect"
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/method
		Method string `json:"method"`

		// Port on which the request came in, this is typically `80` or `443`.
		// If you are running behind a reverse proxy look for the `x-forwarded-port`
		// header.
		//
		// Mininum:    0
		// Maximum:    65535
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/port
		Port int64 `json:"port"`

		// Resource the request operates on including querystring. This is the
		// string that follows the HTTP method.
		// **Note,** order of querystring elements is important.
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/resource
		Resource string `json:"resource"`

		// Source IP of the authentication request or request that requires
		// authentication. This is only used for audit logging.
		//
		// One of:
		//   * SourceIP
		//   * SourceIP1
		//
		// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/sourceIp
		SourceIP string `json:"sourceIp,omitempty"`
	}

	// Response from a request to authenticate a hawk request.
	//
	// One of:
	//   * AuthenticationSuccessfulResponse
	//   * AuthenticationFailedResponse
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-response.json#
	HawkSignatureAuthenticationResponse json.RawMessage

	// List of clients
	//
	// See https://schemas.taskcluster.net/auth/v1/list-clients-response.json#
	ListClientResponse struct {

		// See https://schemas.taskcluster.net/auth/v1/list-clients-response.json#/properties/clients
		Clients []GetClientResponse `json:"clients"`

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		//
		// See https://schemas.taskcluster.net/auth/v1/list-clients-response.json#/properties/continuationToken
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// List of roles
	//
	// See https://schemas.taskcluster.net/auth/v1/list-roles-response.json#
	ListRolesResponse []GetRoleResponse

	// Sentry DSN for submitting errors.
	//
	// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#
	SentryDSNResponse struct {

		// Access credentials and urls for the Sentry project.
		// Credentials will expire in 24-48 hours, you should refresh them within
		// 24 hours.
		//
		// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/dsn
		Dsn Dsn `json:"dsn"`

		// Expiration time for the credentials. The credentials should not be used
		// after this time. They might not be revoked immediately, but will be at
		// some arbitrary point after this date-time.
		//
		// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// Project name that the DSN grants access to.
		//
		// See https://schemas.taskcluster.net/auth/v1/sentry-dsn-response.json#/properties/project
		Project string `json:"project"`
	}

	// A set of scopes
	//
	// See https://schemas.taskcluster.net/auth/v1/scopeset.json#
	SetOfScopes struct {

		// List of scopes.  Scopes must be composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/scopeset.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/scopeset.json#/properties/scopes
		Scopes []string `json:"scopes,omitempty"`
	}

	// Source IP of the authentication request or request that requires
	// authentication. This is only used for audit logging.
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/sourceIp/oneOf[0]
	SourceIP string

	// Source IP of the authentication request or request that requires
	// authentication. This is only used for audit logging.
	//
	// See https://schemas.taskcluster.net/auth/v1/authenticate-hawk-request.json#/properties/sourceIp/oneOf[1]
	SourceIP1 string

	// Token for submitting statistics to statsum.
	//
	// See https://schemas.taskcluster.net/auth/v1/statsum-token-response.json#
	StatsumTokenResponse struct {

		// Base URL for the statsum server this project is allocated on.
		//
		// See https://schemas.taskcluster.net/auth/v1/statsum-token-response.json#/properties/baseUrl
		BaseURL string `json:"baseUrl"`

		// Time at which the token expires and should not be used anymore.
		//
		// See https://schemas.taskcluster.net/auth/v1/statsum-token-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// Project name that the token grants access to.
		//
		// See https://schemas.taskcluster.net/auth/v1/statsum-token-response.json#/properties/project
		Project string `json:"project"`

		// JWT token to be used as `Bearer <token>` when submitting data to statsum.
		//
		// See https://schemas.taskcluster.net/auth/v1/statsum-token-response.json#/properties/token
		Token string `json:"token"`
	}

	// Temporary STS credentials for use when operating on S3
	//
	// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/credentials
	TemporarySecurityCredentials struct {

		// Access key identifier that identifies the temporary security
		// credentials.
		//
		// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/credentials/properties/accessKeyId
		AccessKeyID string `json:"accessKeyId"`

		// Secret access key used to sign requests
		//
		// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/credentials/properties/secretAccessKey
		SecretAccessKey string `json:"secretAccessKey"`

		// A token that must passed with request to use the temporary
		// security credentials.
		//
		// See https://schemas.taskcluster.net/auth/v1/aws-s3-credentials-response.json#/properties/credentials/properties/sessionToken
		SessionToken string `json:"sessionToken"`
	}

	// Details on how the test request should be authenticated.
	//
	// See https://schemas.taskcluster.net/auth/v1/test-authenticate-request.json#
	TestAuthenticateRequest struct {

		// List of scopes that should be client used should be given.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-request.json#/properties/clientScopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-request.json#/properties/clientScopes
		ClientScopes []string `json:"clientScopes,omitempty"`

		// List of scopes the request should require.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-request.json#/properties/requiredScopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-request.json#/properties/requiredScopes
		RequiredScopes []string `json:"requiredScopes,omitempty"`
	}

	// Details on how the test request was authenticated.
	//
	// See https://schemas.taskcluster.net/auth/v1/test-authenticate-response.json#
	TestAuthenticateResponse struct {

		// ClientId from the request as it will be logged
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-response.json#/properties/clientId
		ClientID string `json:"clientId,omitempty"`

		// List of scopes the request was authorized.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-response.json#/properties/scopes/items
		//
		// See https://schemas.taskcluster.net/auth/v1/test-authenticate-response.json#/properties/scopes
		Scopes []string `json:"scopes,omitempty"`
	}

	// Token for connecting a worker to webhooktunnel proxy
	//
	// See https://schemas.taskcluster.net/auth/v1/webhooktunnel-token-response.json#
	WebhooktunnelTokenResponse struct {

		// websocket url at which proxy is hosted
		//
		// See https://schemas.taskcluster.net/auth/v1/webhooktunnel-token-response.json#/properties/proxyUrl
		ProxyURL string `json:"proxyUrl"`

		// jwt token to be used as `Bearer <token>` when connecting to proxy.
		//
		// See https://schemas.taskcluster.net/auth/v1/webhooktunnel-token-response.json#/properties/token
		Token string `json:"token"`

		// id for proxy connection
		//
		// See https://schemas.taskcluster.net/auth/v1/webhooktunnel-token-response.json#/properties/tunnelId
		TunnelID string `json:"tunnelId"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// HawkSignatureAuthenticationResponse is of type json.RawMessage...
func (this *HawkSignatureAuthenticationResponse) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *HawkSignatureAuthenticationResponse) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("HawkSignatureAuthenticationResponse: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
