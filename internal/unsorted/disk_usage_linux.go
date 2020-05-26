/*
 */

// Copyright © 2020 Hedzr Yeh.

package tool

import (
	"syscall"
)

// DiskUsage is for disk usage of path/disk
func DiskUsage(path string) (disk *DiskStatus) {
	disk = &DiskStatus{}
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
