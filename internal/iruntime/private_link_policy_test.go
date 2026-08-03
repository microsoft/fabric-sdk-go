package iruntime

import (
	"testing"
)

func TestWorkspaceIDFromPath(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		expectedID string
		expectedOK bool
	}{
		// --- valid workspace paths ---
		{
			name:       "with sub-resource",
			path:       "/v1/workspaces/12345678-1234-1234-1234-123456789abc/items",
			expectedID: "12345678-1234-1234-1234-123456789abc",
			expectedOK: true,
		},
		{
			name:       "without sub-resource",
			path:       "/v1/workspaces/12345678-1234-1234-1234-123456789abc",
			expectedID: "12345678-1234-1234-1234-123456789abc",
			expectedOK: true,
		},
		{
			name:       "nested sub-resource",
			path:       "/v1/workspaces/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/lakehouses/11111111-2222-3333-4444-555555555555",
			expectedID: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			expectedOK: true,
		},
		{
			name:       "uppercase GUID",
			path:       "/v1/workspaces/AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE/items",
			expectedID: "AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE",
			expectedOK: true,
		},
		{
			name:       "mixed-case GUID",
			path:       "/v1/workspaces/aAbBcCdD-1234-5678-9012-EeFf00112233/items",
			expectedID: "aAbBcCdD-1234-5678-9012-EeFf00112233",
			expectedOK: true,
		},
		{
			name:       "trailing slash after GUID",
			path:       "/v1/workspaces/12345678-1234-1234-1234-123456789abc/",
			expectedID: "12345678-1234-1234-1234-123456789abc",
			expectedOK: true,
		},

		// --- communication policy exclusions ---
		{
			name:       "communicationPolicy excluded",
			path:       "/v1/workspaces/12345678-1234-1234-1234-123456789abc/networking/communicationPolicy",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "communicationPolicy with sub-path is not excluded",
			path:       "/v1/workspaces/12345678-1234-1234-1234-123456789abc/networking/communicationPolicy/settings",
			expectedID: "12345678-1234-1234-1234-123456789abc",
			expectedOK: true,
		},

		// --- non-matching paths ---
		{
			name:       "admin path does not match",
			path:       "/v1/admin/workspaces/12345678-1234-1234-1234-123456789abc",
			expectedID: "",
			expectedOK: false,
		},
		{
			// GET https://api.fabric.microsoft.com/v1/admin/workspaces/networking/communicationpolicies
			name:       "admin communicationpolicies list path does not match",
			path:       "/v1/admin/workspaces/networking/communicationpolicies",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "no workspace segment",
			path:       "/v1/capacities",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "empty path",
			path:       "",
			expectedID: "",
			expectedOK: false,
		},

		// --- invalid GUID formats ---
		{
			name:       "non-hex text",
			path:       "/v1/workspaces/not-a-valid-guid/items",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "GUID without dashes (32 chars)",
			path:       "/v1/workspaces/12345678123412341234123456789abc/items",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "double slash before GUID",
			path:       "/v1/workspaces//12345678-1234-1234-1234-123456789abc/items",
			expectedID: "",
			expectedOK: false,
		},
		{
			name:       "trailing slash with no GUID",
			path:       "/v1/workspaces/",
			expectedID: "",
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotID, gotOK := workspaceIDFromPath(tc.path)
			if gotOK != tc.expectedOK {
				t.Fatalf("workspaceIDFromPath(%q) ok = %v, want %v", tc.path, gotOK, tc.expectedOK)
			}
			if gotID != tc.expectedID {
				t.Errorf("workspaceIDFromPath(%q) id = %q, want %q", tc.path, gotID, tc.expectedID)
			}
		})
	}
}

func TestIsCommunicationPolicyPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{name: "exact communicationPolicy", path: "/networking/communicationPolicy", expected: true},
		{name: "unrelated path", path: "/items", expected: false},
		{name: "empty string", path: "", expected: false},
		{name: "extra segment after policy", path: "/networking/communicationPolicy/extra", expected: false},
		{name: "trailing slash", path: "/networking/communicationPolicy/", expected: false},
		{name: "different casing", path: "/networking/CommunicationPolicy", expected: true},
		{name: "suffix match (policyExtra)", path: "/networking/communicationPolicyExtra", expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isCommunicationPolicyPath(tc.path)
			if got != tc.expected {
				t.Errorf("isCommunicationPolicyPath(%q) = %v, want %v", tc.path, got, tc.expected)
			}
		})
	}
}
