package errors

// Below are predefined keys to errors
const (
	ErrBind                     = "BQ0000001"
	ErrResourceNotFound         = "BQ0000002"
	ErrUndefined                = "BQ0000003"
	ErrLDAPServerFatal          = "BQ0000004"
	ErrAuthentication           = "BQ0000005"
	ErrPrivateKey               = "BQ0000006"
	ErrPublicKey                = "BQ0000007"
	ErrSign                     = "BQ0000008"
	ErrUnexpectedContentType    = "BQ0000009"
	ErrMySQLFatal               = "BQ0000010"
	ErrUnsupportedAuthAlgorithm = "BQ0000011"
	ErrResourceConflict         = "BQ0000012"
)

// CustomError wraps any error message
type CustomError interface {
	Error() string
}

// ResourceNotFoundError wraps 404 errors
type ResourceNotFoundError struct {
	err       error
	message   string
	requestID string
}

func NewResourceNotFoundError(err error, message, requestId string) *ResourceNotFoundError {
	return &ResourceNotFoundError{err, message, requestId}
}

func (resourceNotFoundError *ResourceNotFoundError) Error() string {
	return ErrResourceNotFound + "|" + resourceNotFoundError.message + "|" + resourceNotFoundError.requestID
}

// UndefinedError wraps all unknown errors
type UndefinedError struct {
	err error
}

func NewUndefinedError(err error) *UndefinedError {
	return &UndefinedError{err: err}
}

func (undefinedError *UndefinedError) Error() string {
	return ErrUndefined + "|" + undefinedError.err.Error()
}

// BindError wraps JSON struct binding errors
type BindError struct {
	err       error
	message   string
	requestID string
}

func NewBindError(err error, message, requestId string) *BindError {
	return &BindError{err, message, requestId}
}

func (bindError *BindError) Error() string {
	return ErrBind + "|" + bindError.err.Error() + ". " + bindError.message + "|" + bindError.requestID
}

// LDAPServerFatalError wraps common ldap/ad server errors
type LdapServerFatalError struct {
	err       error
	message   string
	requestID string
}

func NewldapServerFatalError(err error, message, requestId string) *LdapServerFatalError {
	return &LdapServerFatalError{err, message, requestId}
}

func (ldapServerFatalError *LdapServerFatalError) Error() string {
	return ErrLDAPServerFatal + "|" + ldapServerFatalError.err.Error() + ". " + ldapServerFatalError.message + "|" + ldapServerFatalError.requestID
}

// AuthenticationError wraps all failed auth attempts (both client and server)
type AuthenticationError struct {
	err       error
	message   string
	requestID string
}

func NewAuthenticationError(err error, message, requestId string) *AuthenticationError {
	return &AuthenticationError{err, message, requestId}
}

func (authenticationError *AuthenticationError) Error() string {
	return ErrAuthentication + "|" + authenticationError.err.Error() + ". " + authenticationError.message + "|" + authenticationError.requestID
}

// PrivateKeyError wraps all errors regarding the private key used for signin
// the JWT tokens
type PrivateKeyError struct {
	err       error
	message   string
	requestID string
}

func NewPrivateKeyError(err error, message, requestId string) *PrivateKeyError {
	return &PrivateKeyError{err, message, requestId}
}

func (privateKeyError *PrivateKeyError) Error() string {
	return ErrPrivateKey + "|" + privateKeyError.err.Error() + ". " + privateKeyError.message + "|" + privateKeyError.requestID
}

// PublicKeyError wraps all errors regarding the private key used for signin
// the JWT tokens
type PublicKeyError struct {
	err       error
	message   string
	requestID string
}

func NewPublicKeyError(err error, message, requestId string) *PublicKeyError {
	return &PublicKeyError{err, message, requestId}
}

func (publicKeyError *PublicKeyError) Error() string {
	return ErrPublicKey + "|" + publicKeyError.err.Error() + ". " + publicKeyError.message + "|" + publicKeyError.requestID
}

// SignError wraps all failed token sign attempts
type SignError struct {
	err       error
	message   string
	requestID string
}

func NewSignError(err error, message, requestId string) *SignError {
	return &SignError{err, message, requestId}
}

func (signError *SignError) Error() string {
	return ErrSign + "|" + signError.err.Error() + ". " + signError.message + "|" + signError.requestID
}

// ContentTypeError wraps all missing or wrong content type headers in requests
type ContentTypeError struct {
	err       error
	requestID string
}

func NewContentTypeError(err error, requestId string) *ContentTypeError {
	return &ContentTypeError{err, requestId}
}

func (contentTypeError *ContentTypeError) Error() string {
	return ErrUnexpectedContentType + "|" + contentTypeError.err.Error() + "|" + contentTypeError.requestID
}

// MySQLFatalError wraps all mysql common errors
type MySQLFatalError struct {
	err       error
	requestID string
}

func NewMySQLFatalError(err error, requestId string) *MySQLFatalError {
	return &MySQLFatalError{err, requestId}
}

func (mysqlFatalError *MySQLFatalError) Error() string {
	return ErrMySQLFatal + "|" + mysqlFatalError.err.Error() + "|" + mysqlFatalError.requestID
}

// UnsupportedAlgorithmError wraps unknown algorithm requests
type UnsupportedAlgorithmError struct {
	err       error
	requestID string
}

func NewUnsupportedAlgorithmError(err error, requestId string) *UnsupportedAlgorithmError {
	return &UnsupportedAlgorithmError{err, requestId}
}

func (unsuportedAlgorithmError *UnsupportedAlgorithmError) Error() string {
	return ErrUnsupportedAuthAlgorithm + "|" + unsuportedAlgorithmError.err.Error() + "|" + unsuportedAlgorithmError.requestID
}

// UnsupportedAlgorithmError wraps unknown algorithm requests
type ResourceConflictError struct {
	err       error
	requestID string
}

func NewResourceConflictError(err error, requestId string) *ResourceConflictError {
	return &ResourceConflictError{err, requestId}
}

func (resourceConflictError *ResourceConflictError) Error() string {
	return ErrResourceConflict + "|" + resourceConflictError.err.Error() + "|" + resourceConflictError.requestID
}
