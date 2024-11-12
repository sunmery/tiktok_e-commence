package pkg

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

// PgErrorCode
// 因为pgx驱动没有postgres返回的错误码对应的错误消息文本
// 所以使用自定义的错误消息文本提供给客户端
func PgErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ForeignKeyViolation:
			return "违反外键约束"
		case UniqueViolation:
			return "用户已存在"
		}
	}
	return fmt.Sprintf("%s %s", pgErr.Message, pgErr.Code)
}
