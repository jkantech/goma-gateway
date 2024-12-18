/*
 * Copyright 2024 Jonas Kaninda
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package pkg

// Gateway contains Goma Proxy Gateway's configs
type Gateway struct {
	// SSLCertFile  SSL Certificate file
	SSLCertFile string `yaml:"sslCertFile" env:"GOMA_SSL_CERT_FILE, overwrite"`
	// SSLKeyFile SSL Private key  file
	SSLKeyFile string `yaml:"sslKeyFile" env:"GOMA_SSL_KEY_FILE, overwrite"`
	// Redis contains redis database details
	Redis Redis `yaml:"redis"`
	// WriteTimeout defines proxy write timeout
	WriteTimeout int `yaml:"writeTimeout" env:"GOMA_WRITE_TIMEOUT, overwrite"`
	// ReadTimeout defines proxy read timeout
	ReadTimeout int `yaml:"readTimeout" env:"GOMA_READ_TIMEOUT, overwrite"`
	// IdleTimeout defines proxy idle timeout
	IdleTimeout int `yaml:"idleTimeout" env:"GOMA_IDLE_TIMEOUT, overwrite"`
	// RateLimit Defines the number of request peer minutes
	RateLimit int `yaml:"rateLimit" env:"GOMA_RATE_LIMIT, overwrite"`
	// BlockCommonExploits enable, disable block common exploits
	BlockCommonExploits bool   `yaml:"blockCommonExploits"`
	AccessLog           string `yaml:"accessLog" env:"GOMA_ACCESS_LOG, overwrite"`
	ErrorLog            string `yaml:"errorLog" env:"GOMA_ERROR_LOG=, overwrite"`
	LogLevel            string `yaml:"logLevel" env:"GOMA_LOG_LEVEL, overwrite"`

	// DisableHealthCheckStatus enable and disable routes health check
	DisableHealthCheckStatus bool `yaml:"disableHealthCheckStatus"`
	// DisableRouteHealthCheckError allows enabling and disabling backend healthcheck errors
	DisableRouteHealthCheckError bool `yaml:"disableRouteHealthCheckError"`
	// Disable allows enabling and disabling displaying routes on start
	DisableDisplayRouteOnStart bool `yaml:"disableDisplayRouteOnStart"`
	// DisableKeepAlive allows enabling and disabling KeepALive server
	DisableKeepAlive bool `yaml:"disableKeepAlive"`
	EnableMetrics    bool `yaml:"enableMetrics"`
	// InterceptErrors holds the status codes to intercept the error from backend
	InterceptErrors []int `yaml:"interceptErrors"`
	// Cors holds proxy global cors
	Cors Cors `yaml:"cors"`
	// ExtraRoutes additional routes from defined directory
	ExtraRoutes ExtraRouteConfig `yaml:"extraRoutes"`
	// Routes holds proxy routes
	Routes []Route `yaml:"routes"`
}
