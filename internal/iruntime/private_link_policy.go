package iruntime

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/google/uuid"
)

// PrivateLinkPolicy modifies request URLs to support workspace-level private links
type PrivateLinkPolicy struct{}

// workspaceIdRegex matches the workspaceId in the URI as a GUID
// Admin APIs are excluded intentionally since they are blocked through workspace PE
var workspaceIdRegex = regexp.MustCompile(`v1/workspaces/([0-9a-fA-F\-]{36})(/.*)?$`)

// NewPrivateLinkPolicy creates a new instance of PrivateLinkPolicy
func NewPrivateLinkPolicy() *PrivateLinkPolicy {
	return &PrivateLinkPolicy{}
}

// Do implements the policy.Policy interface
func (p *PrivateLinkPolicy) Do(req *policy.Request) (*http.Response, error) {
	// Get the original URL
	originalURL := req.Raw().URL

	// Check if the URI matches the pattern
	matches := workspaceIdRegex.FindStringSubmatch(originalURL.Path)
	if len(matches) > 1 {
		// Extract the workspaceId as a string without the dashes
		workspaceIDStr := matches[1]
		workspaceID, err := uuid.Parse(workspaceIDStr)
		if err != nil {
			return req.Next()
		}

		// Convert to string without dashes and construct prefix
		workspaceIDNoDashes := strings.ReplaceAll(workspaceID.String(), "-", "")
		uriPrefix := workspaceIDNoDashes + ".z" + workspaceIDNoDashes[:2] + ".w."

		// Check if the host is already prefixed - if so, skip modification
		if strings.HasPrefix(originalURL.Host, uriPrefix) {
			return req.Next()
		}

		// Construct the new host with the workspaceId prefix
		newHost := uriPrefix + originalURL.Host

		// Update the request URL
		req.Raw().URL.Host = newHost
	}

	return req.Next()
}
