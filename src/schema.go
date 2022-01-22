package main

import (
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type BadPayload struct {
	Error string `json:"error"`
}

type FilePayload struct {
	Data string `json:"data"`
}

type DirPayload struct {
	Data []FormattedEntry `json:"data"`
}
type FormattedEntry struct {
	Filename    string `json:"filename"`
	OwnerID     string `json:"owner_id"`
	Size        string `json:"size"`
	Permissions string `json:"permissions"`
}

func getOwnership(info os.FileInfo) (uid uint32) {
	stat := info.Sys().(*syscall.Stat_t)
	return stat.Uid
}

func formatEntry(info os.FileInfo) FormattedEntry {
	uid := getOwnership(info)
	return FormattedEntry{
		Filename:    info.Name(),
		OwnerID:     fmt.Sprint(uid),
		Size:        fmt.Sprint(info.Size()),
		Permissions: fmt.Sprintf("%#o", info.Mode().Perm()),
	}
}

func FormatEntries(entries []fs.DirEntry) ([]FormattedEntry, error) {
	transformed := make([]FormattedEntry, len(entries))
	for i, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		transformed[i] = formatEntry(info)
	}
	return transformed, nil
}
