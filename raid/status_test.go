package raid_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/beanieboi/raidstatus/raid"
)

func TestStatus(t *testing.T) {
	status, err := raid.Status()

	if err != nil {
		fmt.Println("FAILED", err)
		t.Error(err)
	}

	fmt.Println(status)
}

func TestParser(t *testing.T) {
	file, err := os.Open("example.plist")

	if err != nil {
		t.Error(err)
	}

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)

	// read file content to buffer
	_, err = file.Read(buffer)

	if err != nil {
		t.Error(err)
	}

	fileBytes := bytes.NewReader(buffer)

	output, err := raid.Parser(fileBytes)

	if err != nil {
		t.Error(err)
	}

	if len(output.RaidSets) != 1 {
		t.Errorf("number of RaidSets) = %d; want 1", len(output.RaidSets))
	}

	set := output.RaidSets[0]

	expect(t, "AppleRAIDSetUUID", set.UUID, "29A25F24-BEA2-47BA-B0F9-323CDF5545EC")
	expect(t, "BSDName", set.BSDName, "disk4")
	expectInt(t, "ChunkCount", int(set.ChunkCount), 122083833)
	expectInt(t, "ChunkSize", int(set.ChunkSize), 32768)
	expect(t, "Content", set.Content, "7C3457EF-0000-11AA-AA11-00306543ECAC")
	expect(t, "Level", set.Level, "Mirror")
	expect(t, "Name", set.Name, "DataRaid")
	expect(t, "Rebuild", set.Rebuild, "Automatic")
	expect(t, "Status", set.Status, "Online")
	expectInt(t, "Size", int(set.Size), 4000443039744)
}

func expectInt(t *testing.T, name string, actual int, expected int) {
	if actual != expected {
		t.Errorf("%s: actual: %d, expected %d", name, actual, expected)
	}
}

func expect(t *testing.T, name string, actual string, expected string) {
	if actual != expected {
		t.Errorf("%s: actual: %s, expected %s", name, actual, expected)
	}
}
