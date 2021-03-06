/*
Copyright 2014 Rohith All rights reserved.

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

package marathon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeployments(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	deployments, err := endpoint.Client.Deployments()
	assert.NoError(t, err)
	assert.NotNil(t, deployments)
	assert.Equal(t, len(deployments), 1)
	deployment := deployments[0]
	assert.NotNil(t, deployment)
	assert.Equal(t, deployment.ID, "867ed450-f6a8-4d33-9b0e-e11c5513990b")
	assert.NotNil(t, deployment.Steps)
	assert.Equal(t, len(deployment.Steps), 1)
}

func TestDeploymentsV1(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, &ConfigContainer{
		server: &ServerConfig{
			Version: "1.1.1",
		},
	})
	defer endpoint.Close()
	deployments, err := endpoint.Client.Deployments()
	assert.NoError(t, err)
	assert.NotNil(t, deployments)
	assert.Equal(t, len(deployments), 1)
	deployment := deployments[0]
	assert.NotNil(t, deployment)
	assert.Equal(t, deployment.ID, "2620aa06-1001-4eea-8861-a51957d4fd80")
	assert.NotNil(t, deployment.Steps)
	assert.Equal(t, len(deployment.Steps), 2)
}

func TestDeleteDeployment(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()
	id, err := endpoint.Client.DeleteDeployment(fakeDeploymentID, false)
	assert.NoError(t, err)
	assert.NotNil(t, t)
	assert.Equal(t, id.DeploymentID, "0b1467fc-d5cd-4bbc-bac2-2805351cee1e")
	assert.Equal(t, id.Version, "2014-08-26T08:20:26.171Z")
}
