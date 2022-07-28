package storage

import (
	"fmt"
	"log"
)

type RemoteStorage interface {
	List(bucket string) []RemoteFile
}

type SystemWriter interface {
	PrintRemoteFiles(files []RemoteFile)
	PrintAWSConfiguration(config ConfigChangeEvent)
	PrintTestResults(results TestResults)
}

type RemoteFile struct {
	Key  string
	Size int64
}

type GeneratedFiles struct {
	Files []*GeneratedFile
}

type GeneratedFile struct {
	Path     string
	Contents []byte
}

func ListFilesFromStorage(bucket string, storage RemoteStorage, writer SystemWriter) {
	files := storage.List(bucket)
	for _, file := range files {
		file.Key = fmt.Sprintf("Prefix: %s", file.Key)
	}
	writer.PrintRemoteFiles(files)
}

type TestTriggerEvent struct {
	Key string
}

type ConfigChangeEvent struct {
	Key string
}

type TestResults struct {
	Success bool
}

func PerformTestTrigger(trigger TestTriggerEvent) {
	log.Println("XYZ")
}

func PerformConfigTrigger(trigger ConfigChangeEvent) {
	log.Println("ABC")
}
