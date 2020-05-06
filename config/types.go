/*
 * Types for config.
 * Wang Ruichao (793160615@qq.com)
 */

package config

type configSt struct {
	HTTPPort string `yaml:"http_port"`
	MongoURL string `yaml:"mongo_url"`
}
