package search_common

import (
	"context"
	"github.com/segmentfault/answer/internal/schema"
)

type SearchRepo interface {
	SearchContents(ctx context.Context, words []string, tagID, userID string, votes int, page, size int) (resp []schema.SearchResp, total int64, err error)
	SearchQuestions(ctx context.Context, words []string, limitNoAccepted bool, answers, page, size int) (resp []schema.SearchResp, total int64, err error)
	SearchAnswers(ctx context.Context, words []string, limitAccepted bool, questionID string, page, size int) (resp []schema.SearchResp, total int64, err error)
}