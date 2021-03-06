/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pipeline

type DeploymentListDeploymentResponse struct {

	Chart string `json:"chart,omitempty"`

	ChartName string `json:"chartName,omitempty"`

	ChartVersion string `json:"chartVersion,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	ReleaseName string `json:"releaseName,omitempty"`

	UpdatedAt string `json:"updatedAt,omitempty"`

	Version int32 `json:"version,omitempty"`
}
