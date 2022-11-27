package model

import (
	"fmt"
	"time"
)

type IndexProgress struct {
	FileCount    uint64     `json:"file_count"`
	IsDone       bool       `json:"is_done"`
	LastDoneTime *time.Time `json:"last_done_time"`
	Error        string     `json:"error"`
}

type SearchReq struct {
	Path     string `json:"path"`
	Keywords string `json:"keywords"`
	PageReq
}

type SearchNode struct {
	Parent string `json:"parent"`
	Name   string `json:"name"`
	IsDir  bool   `json:"is_dir"`
	Size   int64  `json:"size"`
}

func (p *SearchReq) Validate() error {
	if p.Page < 1 {
		return fmt.Errorf("page can't < 1")
	}
	if p.PerPage < 1 {
		return fmt.Errorf("per_page can't < 1")
	}
	return nil
}
