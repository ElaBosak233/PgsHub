package repositories

import (
	model "github.com/elabosak233/pgshub/internal/models/data"
	"github.com/elabosak233/pgshub/internal/models/request"
	"xorm.io/xorm"
)

type SubmissionRepository interface {
	Insert(submission model.Submission) (err error)
	Delete(id string) (err error)
	Find(req request.SubmissionFindRequest) (submissions []model.Submission, count int64, err error)
}

type SubmissionRepositoryImpl struct {
	Db *xorm.Engine
}

func NewSubmissionRepositoryImpl(Db *xorm.Engine) SubmissionRepository {
	return &SubmissionRepositoryImpl{Db: Db}
}

func (t *SubmissionRepositoryImpl) Insert(submission model.Submission) (err error) {
	_, err = t.Db.Table("submission").Insert(&submission)
	return err
}

func (t *SubmissionRepositoryImpl) Delete(id string) (err error) {
	_, err = t.Db.Table("submission").ID(id).Delete(&model.Submission{})
	return err
}

func (t *SubmissionRepositoryImpl) Find(req request.SubmissionFindRequest) (submissions []model.Submission, count int64, err error) {
	applyFilters := func(q *xorm.Session) *xorm.Session {
		if req.UserId != "" {
			q = q.Where("user_id = ?", req.UserId)
		}
		if req.ChallengeId != "" {
			q = q.Where("challenge_id = ?", req.ChallengeId)
		}
		if req.TeamId != "" {
			q = q.Where("team_id = ?", req.TeamId)
		}
		if req.GameId != -1 {
			q = q.Where("game_id = ?", req.GameId)
		}
		if req.Status != -1 {
			q = q.Where("status = ?", req.Status)
		}
		if req.IsDetailed == 0 {
			q = q.Omit("flag")
		}
		if req.IsAscend {
			q = q.Asc("created_at")
		} else {
			q = q.Desc("created_at")
		}
		return q
	}
	db := applyFilters(t.Db.Table("submission"))
	ct := applyFilters(t.Db.Table("submission"))
	count, err = ct.Count(&model.Submission{})
	if req.Page != -1 && req.Size != -1 {
		offset := (req.Page - 1) * req.Size
		db = db.Limit(req.Size, offset)
	}
	err = db.Find(&submissions)
	return submissions, count, err
}
