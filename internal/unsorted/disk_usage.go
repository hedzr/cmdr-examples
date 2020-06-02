/*
 */

// Copyright © 2020 Hedzr Yeh.

package unsorted

// DiskStatus provides the disk usages
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}
