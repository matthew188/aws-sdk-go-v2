//go:build integration && perftest
// +build integration,perftest

package downloader

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/matthew188/aws-sdk-go-v2/config"
	"github.com/matthew188/aws-sdk-go-v2/feature/s3/manager"
	"github.com/matthew188/aws-sdk-go-v2/feature/s3/manager/internal/integration"
	"github.com/matthew188/aws-sdk-go-v2/internal/awstesting"
	"github.com/matthew188/aws-sdk-go-v2/internal/sdkio"
	"github.com/matthew188/aws-sdk-go-v2/service/s3"
)

var benchConfig BenchmarkConfig

type BenchmarkConfig struct {
	bucket         string
	tempdir        string
	clientConfig   ClientConfig
	sizes          string
	parts          string
	concurrency    string
	bufferSize     string
	uploadPartSize int64
}

func (b *BenchmarkConfig) SetupFlags(prefix string, flagSet *flag.FlagSet) {
	flagSet.StringVar(&b.bucket, "bucket", "", "Bucket to use for benchmark")
	flagSet.StringVar(&b.tempdir, "temp", os.TempDir(), "location to create temporary files")

	flagSet.StringVar(&b.sizes, "size",
		fmt.Sprintf("%d,%d",
			5*sdkio.MebiByte,
			1*sdkio.GibiByte), "file sizes to benchmark separated by comma")

	flagSet.StringVar(&b.parts, "part",
		fmt.Sprintf("%d,%d,%d",
			manager.DefaultDownloadPartSize,
			25*sdkio.MebiByte,
			100*sdkio.MebiByte), "part sizes to benchmark separated by comma")

	flagSet.StringVar(&b.concurrency, "concurrency",
		fmt.Sprintf("%d,%d,%d",
			manager.DefaultDownloadConcurrency,
			2*manager.DefaultDownloadConcurrency,
			100),
		"part sizes to benchmark separated comma")

	flagSet.StringVar(&b.bufferSize, "buffer", fmt.Sprintf("%d,%d", 0, 1*sdkio.MebiByte), "part sizes to benchmark separated comma")
	flagSet.Int64Var(&b.uploadPartSize, "upload-part-size", 0, "upload part size, defaults to download part size if not specified")
	b.clientConfig.SetupFlags(prefix, flagSet)
}

func (b *BenchmarkConfig) BufferSizes() []int {
	ints, err := b.stringToInt(b.bufferSize)
	if err != nil {
		panic(fmt.Sprintf("failed to parse file sizes: %v", err))
	}

	return ints
}

func (b *BenchmarkConfig) FileSizes() []int64 {
	ints, err := b.stringToInt64(b.sizes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse file sizes: %v", err))
	}

	return ints
}

func (b *BenchmarkConfig) PartSizes() []int64 {
	ints, err := b.stringToInt64(b.parts)
	if err != nil {
		panic(fmt.Sprintf("failed to parse part sizes: %v", err))
	}

	return ints
}

func (b *BenchmarkConfig) Concurrences() []int {
	ints, err := b.stringToInt(b.concurrency)
	if err != nil {
		panic(fmt.Sprintf("failed to parse part sizes: %v", err))
	}

	return ints
}

func (b *BenchmarkConfig) stringToInt(s string) ([]int, error) {
	int64s, err := b.stringToInt64(s)
	if err != nil {
		return nil, err
	}

	var ints []int
	for i := range int64s {
		ints = append(ints, int(int64s[i]))
	}

	return ints, nil
}

func (b *BenchmarkConfig) stringToInt64(s string) ([]int64, error) {
	var sizes []int64

	split := strings.Split(s, ",")

	for _, size := range split {
		size = strings.Trim(size, " ")
		i, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %s: %v", size, err)
		}

		sizes = append(sizes, i)
	}

	return sizes, nil
}

func BenchmarkDownload(b *testing.B) {
	baseSdkConfig := SDKConfig{}

	for _, fileSize := range benchConfig.FileSizes() {
		b.Run(fmt.Sprintf("%s File", integration.SizeToName(int(fileSize))), func(b *testing.B) {
			for _, partSize := range benchConfig.PartSizes() {
				if partSize > fileSize {
					continue
				}
				uploadPartSize := getUploadPartSize(fileSize, benchConfig.uploadPartSize, partSize)
				b.Run(fmt.Sprintf("%s PartSize", integration.SizeToName(int(partSize))), func(b *testing.B) {
					b.Logf("setting up s3 file size")
					key, err := setupDownloadTest(benchConfig.bucket, fileSize, uploadPartSize)
					if err != nil {
						b.Fatalf("failed to setup download test: %v", err)
					}
					for _, concurrency := range benchConfig.Concurrences() {
						b.Run(fmt.Sprintf("%d Concurrency", concurrency), func(b *testing.B) {
							for _, bufferSize := range benchConfig.BufferSizes() {
								var name string
								if bufferSize == 0 {
									name = "unbuffered"
								} else {
									name = fmt.Sprintf("%s buffer", integration.SizeToName(bufferSize))
								}
								b.Run(name, func(b *testing.B) {
									sdkConfig := baseSdkConfig
									sdkConfig.Concurrency = concurrency
									sdkConfig.PartSize = partSize
									if bufferSize > 0 {
										sdkConfig.BufferProvider = manager.NewPooledBufferedWriterReadFromProvider(bufferSize)
									}

									b.ResetTimer()
									for i := 0; i < b.N; i++ {
										benchDownload(b, benchConfig.bucket, key, &awstesting.DiscardAt{}, sdkConfig, benchConfig.clientConfig)
									}
								})
							}
						})
					}
					b.Log("removing test file")
					err = teardownDownloadTest(benchConfig.bucket, key)
					if err != nil {
						b.Fatalf("failed to cleanup test file: %v", err)
					}
				})
			}
		})
	}
}

func benchDownload(b *testing.B, bucket, key string, body io.WriterAt, sdkConfig SDKConfig, clientConfig ClientConfig) {
	downloader := newDownloader(clientConfig, sdkConfig)
	_, err := downloader.Download(context.Background(), body, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		b.Fatalf("failed to download object, %v", err)
	}
}

func TestMain(m *testing.M) {
	if strings.EqualFold(os.Getenv("BUILD_ONLY"), "true") {
		os.Exit(0)
	}
	benchConfig.SetupFlags("", flag.CommandLine)
	flag.Parse()
	os.Exit(m.Run())
}

func setupDownloadTest(bucket string, fileSize, partSize int64) (key string, err error) {
	er := &awstesting.EndlessReader{}
	lr := io.LimitReader(er, fileSize)

	key = integration.MustUUID()

	defaultConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(defaultConfig)

	uploader := manager.NewUploader(client, func(u *manager.Uploader) {
		u.PartSize = partSize
		u.Concurrency = runtime.NumCPU() * 2
	})

	_, err = uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &bucket,
		Body:   lr,
		Key:    &key,
	})
	if err != nil {
		err = fmt.Errorf("failed to upload test object to s3: %v", err)
	}

	return
}

func teardownDownloadTest(bucket, key string) error {
	defaultConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := s3.NewFromConfig(defaultConfig)

	_, err = client.DeleteObject(context.Background(), &s3.DeleteObjectInput{Bucket: &bucket, Key: &key})
	return err
}

func newDownloader(clientConfig ClientConfig, sdkConfig SDKConfig) *manager.Downloader {
	defaultConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := s3.NewFromConfig(defaultConfig, func(options *s3.Options) {
		options.HTTPClient = NewHTTPClient(clientConfig)
	})

	downloader := manager.NewDownloader(client, func(d *manager.Downloader) {
		d.PartSize = sdkConfig.PartSize
		d.Concurrency = sdkConfig.Concurrency
		d.BufferProvider = sdkConfig.BufferProvider
	})

	return downloader
}

func getUploadPartSize(fileSize, uploadPartSize, downloadPartSize int64) int64 {
	partSize := uploadPartSize

	if partSize == 0 {
		partSize = downloadPartSize
	}
	if fileSize/partSize > int64(manager.MaxUploadParts) {
		partSize = (fileSize / int64(manager.MaxUploadParts)) + 1
	}

	return partSize
}
