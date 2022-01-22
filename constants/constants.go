package constants

import "time"

// request related
const (
	HTTP_METHOD_GET  = "GET"
	HTTP_METHOD_POST = "POST"
)

// api status
const (
	API_FAILED_STATUS  = "Fail"
	API_SUCCESS_STATUS = "Success"
)

// DB constants
const (
	DB_READER = "reader"
	DB_WRITER = "writer"
)

// caching constants durations are in nano sec
var (
	CACHE_TTL_VERY_SHORT time.Duration = 60 * 1_000_000_000
	CACHE_TTL_SHORT      time.Duration = 300 * 1_000_000_000
	CACHE_TTL_MEDIUM     time.Duration = 1_800 * 1_000_000_000
	CACHE_TTL_LONG       time.Duration = 3_600 * 1_000_000_000
	CACHE_TTL_VERY_LONG  time.Duration = 86_400 * 1_000_000_000
)

// info messages
const (
	INFO_CACHE_DISABLED = "cache disabled"
)

// code alphabets
const (
	CODE_ALPHABET_SHORT = "abcdef0123456789"
	CODE_ALPHABET_LONG  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// error messages
const (
	INVALID_PASSWORD                  = "incorrect password"
	INVALID_USER_ID                   = "invalid userId"
	ERROR_USER_IDS_NOT_FOUND          = "userIds not found"
	INVALID_REQUEST                   = "invalid json request body"
	INVALID_MAIL_ID                   = "invalid email address provided"
	ERROR_IN_HASHING_PASSWORD         = "error while hashing password"
	ERROR_IN_STORING_UNIQUE_USER      = "the user already exists"
	ERROR_IN_AUTHENTICATING_USER      = "the password is incorrect"
	USER_ID_NOT_ARRAY_OF_INT          = "userIds must be an array of integers"
	ERROR_NO_USER_EXIST               = "User does not exist"
	ERROR_CONTENT_NOT_FOUND           = "content not found"
	ERROR_INVALID_INTERACTION_DETAILS = "invalid interaction details"
	ALREADY_LIKED                     = "user has already liked the post"
)

const (
	MIN_LENGTH_OF_FIRSTNAME = 2
	MAX_LENGTH_OF_FIRSTNAME = 50
	MIN_LENGTH_OF_LASTNAME  = 2
	MAX_LENGTH_OF_LASTNAME  = 50
	MIN_LENGTH_OF_PASSWORD  = 6
	MAX_LENGTH_OF_PASSWORD  = 100
)

// response messages
const (
	LIKE_MESSAGE = "post liked"
)

// required request body fields
var (
	LIKE_POST_REQUIRED_FIELDS = []string{"userId", "contentId"}
	LIKE_POST_OPTIONAL_FIELDS = []string{}
)

const (
	API_TYPE_LIKE = "likes"
	API_TYPE_READ = "reads"
)
