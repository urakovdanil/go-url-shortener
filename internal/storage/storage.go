package storage

import "github.com/urakovdanil/go-url-shortener/internal/storage/common"

var Used common.Storage

func SetUsed(s *common.Storage) {
	Used = *s
}
