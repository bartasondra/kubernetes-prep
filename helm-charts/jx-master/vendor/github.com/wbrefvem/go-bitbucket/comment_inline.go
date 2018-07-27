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

type CommentInline struct {

	// The comment's anchor line in the new version of the file.
	To int32 `json:"to,omitempty"`

	// The comment's anchor line in the old version of the file.
	From int32 `json:"from,omitempty"`

	// The path of the file this comment is anchored to.
	Path string `json:"path"`
}
