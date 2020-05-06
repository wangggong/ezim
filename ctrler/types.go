/*
 * Types for ctrler.
 * Wang Ruichao (793160615@qq.com)
 */

package ctrler

type basicResp struct {
	Ret  int64       `json:"ret"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
