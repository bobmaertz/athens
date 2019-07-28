package actions

import (
	"testing"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGetStorage(t *testing.T) {
	tests := []struct {
		testName      string
		storageType   string
		storageConfig *config.StorageConfig
	}{
		//{"create disk storage configuration", "disk", new(config.StorageConfig)},
		{"create memory storage configuration", "memory", new(config.StorageConfig)}, // test memory storage creation
		//{"create mongo storage client", "mongo", &config.StorageConfig{Disk: nil, GCP: nil, Minio: nil, Mongo: new(config.MongoConfig), S3: nil, AzureBlob: nil}}, // test mongo storage creation
		//{"azureblob", &config.StorageConfig{Disk: nil, GCP: nil, Minio: nil, Mongo: nil, S3: nil, AzureBlob: new(config.AzureBlobConfig)}}, // test mongo storage creation
		//{"s3", &config.StorageConfig{Disk: nil, GCP: nil, Minio: nil, Mongo: nil, S3: new(config.S3Config), AzureBlob: nil}},               // test mongo storage creation
		//{"create minio storage client", "minio", &config.StorageConfig{Disk: nil, GCP: nil, Minio: new(config.MinioConfig), Mongo: nil, S3: nil, AzureBlob: nil}}, // test mongo storage creation
		//{"gcp", &config.StorageConfig{Disk: nil, GCP: new(config.GCPConfig), Minio: nil, Mongo: nil, S3: nil, AzureBlob: nil}}, // test mongo storage creation
	}

	timeout := time.Duration(300) //Timeout is passed through to constructors

	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			backend, err := GetStorage(testCase.storageType, testCase.storageConfig, timeout)
			require.NoError(t, err)
			require.NotNil(t, backend)
		})
	}
}

//Test error cases generated from GetStorage. No dependency on storage backend and can be run locally.
func TestGetStorageErrors(t *testing.T) {
	storageConfig := &config.StorageConfig{Disk: nil, GCP: nil, Minio: nil, Mongo: nil, S3: nil, AzureBlob: nil}

	tests := []struct {
		testName      string
		storageType   string
		storageConfig *config.StorageConfig
	}{
		{"nil disk storage configuration", "disk", storageConfig},
		{"nil mongo storage configuration", "mongo", storageConfig},
		{"nil azureblob storage configuration", "azureblob", storageConfig},
		{"nil s3 storage configuration", "s3", storageConfig},
		{"nil minio storage configuration", "minio", storageConfig},
		{"nil gcp storage configuration", "gcp", storageConfig},
		{"unrecognized storage type", "fake_type", storageConfig},
	}

	timeout := time.Duration(300) //Timeout is passed through to constructors

	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			backend, err := GetStorage(testCase.storageType, testCase.storageConfig, timeout)
			require.Error(t, err)
			require.Nil(t, backend)
		})
	}
}
