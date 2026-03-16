/*
Copyright 2026 Fabien Dupont.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//go:build e2e

package e2e

import (
	"fmt"
)

// createCloudConfigSecret creates a YAML cloud-config for the CCM.
func createCloudConfigSecret(endpoint, orgName, token, siteID, tenantID string) string {
	return fmt.Sprintf(`endpoint: %q
orgName: %q
token: %q
siteId: %q
tenantId: %q
`, endpoint, orgName, token, siteID, tenantID)
}
