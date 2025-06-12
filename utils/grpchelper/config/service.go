package config

import (
	"encoding/json"
	"fmt"
	"time"
)

// ServiceConfig defines the structure for Service Config
type ServiceConfig struct {
	LoadBalancingPolicy string         `json:"loadBalancingPolicy"`
	MethodConfigs       []MethodConfig `json:"methodConfig"`
}

// MethodConfig defines the structure for methodConfig
type MethodConfig struct {
	Names       []NameConfig `json:"name"`
	RetryPolicy RetryPolicy  `json:"retryPolicy"`
}

// NameConfig defines the structure for name field (service and method)
type NameConfig struct {
	Service string `json:"service"`
	Method  string `json:"method,omitempty"` // Optional, empty means applicable to all methods
}

// RetryPolicy defines the structure for retryPolicy
type RetryPolicy struct {
	MaxAttempts          int      `json:"maxAttempts"`
	InitialBackoff       string   `json:"initialBackoff"`
	MaxBackoff           string   `json:"maxBackoff"`
	BackoffMultiplier    float64  `json:"backoffMultiplier"`
	RetryableStatusCodes []string `json:"retryableStatusCodes"`
}

// GenerateServiceConfig generates a JSON string of gRPC Service Config
func GenerateServiceConfig(
	loadBalancingPolicy string,
	serviceNames []string, // List of service names, e.g. ["MorseBusiness"]
	methodNames []string, // Optional, list of method names, if empty, applies to all methods in the service
	maxAttempts int, // Maximum number of retry attempts
	initialBackoff time.Duration, // Initial backoff time
	maxBackoff time.Duration, // Maximum backoff time
	backoffMultiplier float64, // Backoff multiplier
	retryableStatusCodes []string, // Retryable error codes, e.g. ["UNAVAILABLE"]
) (string, error) {
	// Validate input parameters
	if loadBalancingPolicy == "" {
		return "", fmt.Errorf("loadBalancingPolicy cannot be empty")
	}
	if len(serviceNames) == 0 {
		return "", fmt.Errorf("serviceNames cannot be empty")
	}
	if maxAttempts < 2 {
		return "", fmt.Errorf("maxAttempts must be greater than or equal to 2")
	}
	if initialBackoff <= 0 {
		return "", fmt.Errorf("initialBackoff must be greater than 0")
	}
	if maxBackoff <= 0 {
		return "", fmt.Errorf("maxBackoff must be greater than 0")
	}
	if backoffMultiplier <= 0 {
		return "", fmt.Errorf("backoffMultiplier must be greater than 0")
	}
	if len(retryableStatusCodes) == 0 {
		return "", fmt.Errorf("retryableStatusCodes cannot be empty")
	}

	// Construct name configuration
	names := make([]NameConfig, 0, len(serviceNames)*len(methodNames))
	if len(methodNames) == 0 {
		// If no method is specified, it applies to all methods of the service
		for _, service := range serviceNames {
			names = append(names, NameConfig{Service: service})
		}
	} else {
		// Generate name for each combination of service and method
		for _, service := range serviceNames {
			for _, method := range methodNames {
				names = append(names, NameConfig{Service: service, Method: method})
			}
		}
	}

	// Construct Service Config
	config := ServiceConfig{
		LoadBalancingPolicy: loadBalancingPolicy,
		MethodConfigs: []MethodConfig{
			{
				Names: names,
				RetryPolicy: RetryPolicy{
					MaxAttempts:          maxAttempts,
					InitialBackoff:       fmt.Sprintf("%ds", int(initialBackoff.Seconds())),
					MaxBackoff:           fmt.Sprintf("%ds", int(maxBackoff.Seconds())),
					BackoffMultiplier:    backoffMultiplier,
					RetryableStatusCodes: retryableStatusCodes,
				},
			},
		},
	}

	// Convert to JSON
	configJSON, err := json.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to serialize Service Config: %v", err)
	}

	return string(configJSON), nil
}
