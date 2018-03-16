/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package base

import "github.com/nebulaim/telegramd/mtproto"

/*
	Error handling

	There will be errors when working with the API, and they must be correctly handled on the client.

	An error is characterized by several parameters:

	Error Code
	Similar to HTTP status. Contains information on the type of error that occurred: for example,
	a data input error, privacy error, or server error. This is a required parameter.

	Error Type
	A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED.
	This is an optional parameter.

	Error Description
	May contain more detailed information on the error and how to resolve it,
	for example: authorization required, use auth.* methods. Please note that the description text is subject to change,
	one should avoid tying application logic to these messages. This is an optional parameter.

	Error Constructors
	There should be a way to handle errors that are returned in rpc_error constructors.

	If an error constructor does not differentiate between type and description
	but instead contains a single field called error_message (as in the example above),
	it must be split into 2 components, for example, using the following regular expression: /^([A-Z_0-9]+)(: (.+))?/.
 */
type RpcErrorCode uint32

const (
	// The request must be repeated, but directed to a different data center.
	//
	// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
	// PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
	// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
	// USER_MIGRATE_X: the user whose identity is being used to execute
	// 				   queries is associated with a different data center (for registration)
	//
	// In all these cases, the error description’s string literal
	// 		contains the number of the data center (instead of the X) to which the repeated query must be sent.
	// More information about redirects between data centers »
	ERROR_SEE_OTHER RpcErrorCode = 303

	// The query contains errors. In the event that a request was created using a form
	// and contains user generated data,
	// the user should be notified that the data must be corrected before the query is repeated.
	//
	//
	// Examples of Errors:
	//	FIRSTNAME_INVALID: The first name is invalid
	//	LASTNAME_INVALID: The last name is invalid
	//	PHONE_NUMBER_INVALID: The phone number is invalid
	//	PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
	//	PHONE_CODE_EMPTY: phone_code is missing
	//	PHONE_CODE_EXPIRED: The confirmation code has expired
	//	API_ID_INVALID: The api_id/api_hash combination is invalid
	//	PHONE_NUMBER_OCCUPIED: The phone number is already in use
	//	PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
	//	USERS_TOO_FEW: Not enough users (to create a chat, for example)
	//	USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
	//	TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
	//	FILE_PART_INVALID: The file part number is invalid
	//	FILE_PARTS_INVALID: The number of file parts is invalid
	//	FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage
	//	MD5_CHECKSUM_INVALID: The MD5 checksums do not match
	//	PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
	//	FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
	//	FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
	//	MSG_WAIT_FAILED: A waiting call returned an error
	BAD_REQUEST RpcErrorCode = 400

	// There was an unauthorized attempt to use functionality available only to authorized users.
	//
	// Examples of Errors:
	//	AUTH_KEY_UNREGISTERED: The key is not registered in the system
	//	AUTH_KEY_INVALID: The key is invalid
	//	USER_DEACTIVATED: The user has been deleted/deactivated
	//	SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	//	SESSION_EXPIRED: The authorization has expired
	//	ACTIVE_USER_REQUIRED: The method is only available to already activated users
	//	AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	UNAUTHORIZED RpcErrorCode = 401

	// Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	FORBIDDEN RpcErrorCode = 403

	// The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded.
	// For example, in an attempt to request a large number of text messages (SMS) for the same phone number.
	//
	// Error Example:
	// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	FLOOD RpcErrorCode = 420


	// An internal server error occurred while a request was being processed;
	// for example, there was a disruption while accessing a database or file storage.
	//
	// If a client receives a 500 error, or you believe this error should not have occurred,
	// please collect as much information as possible about the query and error and send it to the developers.
	INTERNAL RpcErrorCode = 500

	// If a server returns an error with a code other than the ones listed above,
	// it may be considered the same as a 500 error and treated as an internal server error.
	//
	OTHER RpcErrorCode = 501
)

/*
	auth.signIn Errors

	Code	Type	Description
	400 BAD_REQUEST	PHONE_NUMBER_INVALID	Invalid phone number
	400 BAD_REQUEST	PHONE_CODE_EMPTY	phone_code from the SMS is empty
	400 BAD_REQUEST	PHONE_CODE_EXPIRED	SMS expired
	400 BAD_REQUEST	PHONE_CODE_INVALID	Invalid SMS code was sent
	400 BAD_REQUEST	PHONE_NUMBER_UNOCCUPIED	The code is valid but no user with the given number is registered
 */

// func (i RpcErrorCode) String() string {
// }
func NewRpcError(errorCode RpcErrorCode, errorMessage string) *mtproto.TLRpcError {
	return &mtproto.TLRpcError{
		ErrorCode: int32(errorCode),
		ErrorMessage: errorMessage,
	}
}

/*
	Only a small portion of the API methods are available to unauthorized users:

	auth.sendCode
	auth.sendCall
	auth.checkPhone
	auth.signUp
	auth.signIn
	auth.importAuthorization
	help.getConfig
	help.getNearestDc

	Other methods will result in an error: 401 UNAUTHORIZED.
 */