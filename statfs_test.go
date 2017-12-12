package statfs

import (
	"testing"
)

func TestGetDiskInfo(t *testing.T) {
	disUsage, err := GetDiskInfo("/")
	if err != nil {
		t.Errorf("Failed to get disk info,because of %s", err.Error())
	}
	t.Logf("available %d, capacity %d, usage %d, inodes %d, inodesFree %d, inodesUsed %d",
		disUsage.Available, disUsage.Capacity, disUsage.Usage, disUsage.Inodes, disUsage.InodesFree, disUsage.InodesUsed)
}
