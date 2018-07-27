/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

// A reference to a range of bytes in a log file (does not contain the actual bytes).
type PipelineLogRange struct {

	// The position of the last byte of the range in the log.
	LastBytePosition int32 `json:"last_byte_position,omitempty"`

	// The position of the first byte of the range in the log.
	FirstBytePosition int32 `json:"first_byte_position,omitempty"`
}
