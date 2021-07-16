package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	disk := os.Getenv("DEST_DISK")
	fileOut, err := os.OpenFile(disk, os.O_CREATE|os.O_WRONLY, 0644)
	defer fileOut.Close()
	if err != nil {
		fmt.Printf("unable to open the target disk %s: %v\n", disk, err)
		return
	}

	// Do the equivalent of partprobe on the device
	if err := fileOut.Sync(); err != nil {
		fmt.Printf("failed to sync the block device: %v\n", err)
		return
	}

	if err := unix.IoctlSetInt(int(fileOut.Fd()), unix.BLKRRPART, 0); err != nil {
		fmt.Printf("error re-probing the partitions for the specified device: %v\n", err)
		return
	}
}
