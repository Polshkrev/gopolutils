package fayl

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/Polshkrev/gopolutils"
)

// Gzip a given slice of bytes.
// If the data can not be written, an [gopolutils.IOError] is returned with a nil data pointer.
// If the [gzip.Writer] can not be closed, an [gopolutils.IOError] is returned with a nil data pointer.
func Gzip(content []byte) ([]byte, *gopolutils.Exception) {
	return runConcurrent(content, concurrentGzip)
}

// Gzip a given slice of bytes.
// If a [gzip.Reader] can not be acquired, an [gopolutils.IOError] is returned with a nil data pointer.
// If the data can not be read, an [gopolutils.IOError] is returned with a nil data pointer.
func Gunzip(content []byte) ([]byte, *gopolutils.Exception) {
	return runConcurrent(content, concurrentGunzip)
}

// Run a given function concurrently.
// Returns the result of the channel after being run by the given caller.
func runConcurrent(content []byte, caller func(content []byte, resultChannel chan<- []byte, errorChannel chan<- *gopolutils.Exception)) ([]byte, *gopolutils.Exception) {
	var byteChannel chan []byte = make(chan []byte, 1)
	var exceptionChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go caller(content, byteChannel, exceptionChannel)
	var result []byte = <-byteChannel
	var except *gopolutils.Exception = <-exceptionChannel
	return result, except
}

// Concurrently unzip a given gzipped slice of bytes.
// If a [gzip.Reader] can not be acquired, an [gopolutils.IOError] is sent to the exception channel.
// If the data can not be read, an [gopolutils.IOError] is sent to the exception channel.
func concurrentGunzip(content []byte, resultChannel chan<- []byte, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(exceptionChannel)
	var reader *gzip.Reader
	var readerError error
	reader, readerError = gzip.NewReader(bytes.NewReader(content))
	if readerError != nil {
		resultChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.IOError, readerError.Error())
		return
	}
	defer reader.Close()
	var uncompressed []byte
	var readError error
	uncompressed, readError = io.ReadAll(reader)
	if readError != nil {
		resultChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.IOError, readError.Error())
		return
	}
	resultChannel <- uncompressed
	exceptionChannel <- nil
}

// Concurrently gzip a given slice of bytes.
// If the data can not be written, an [gopolutils.IOError] is sent to the exception channel.
// If the writer can not be closed, an [gopolutils.IOError] is sent to the exception channel.
func concurrentGzip(content []byte, resultChannel chan<- []byte, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(exceptionChannel)
	var buffer *bytes.Buffer = new(bytes.Buffer)
	var writer *gzip.Writer = gzip.NewWriter(buffer)
	var writeError error
	_, writeError = writer.Write(content)
	if writeError != nil {
		resultChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.IOError, writeError.Error())
		return
	}
	var closeError error = writer.Close()
	if closeError != nil {
		resultChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.IOError, closeError.Error())
		return
	}
	resultChannel <- buffer.Bytes()
	exceptionChannel <- nil
}
