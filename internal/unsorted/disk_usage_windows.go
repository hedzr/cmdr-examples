/*
 */

// Copyright © 2020 Hedzr Yeh.

package unsorted

// DiskUsage is for disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	disk.All = 0
	disk.Free = 0
	disk.Used = 0
	return
}
