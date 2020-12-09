/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package managedservices

// KafkaRequestPayload Schema for the request body sent to /kafkas POST
type KafkaRequestPayload struct {
	// The cloud provider where the Kafka cluster will be created in
	CloudProvider string `json:"cloud_provider,omitempty"`
	// Set this to true to configure the Kafka cluster to be multiAZ
	MultiAz bool `json:"multi_az,omitempty"`
	// The name of the Kafka cluster
	Name string `json:"name"`
	// The region where the Kafka cluster will be created in
	Region string `json:"region,omitempty"`
}