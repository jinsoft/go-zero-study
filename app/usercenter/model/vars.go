package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	LoginWithPhone = iota
	LoginWithEmail
)
