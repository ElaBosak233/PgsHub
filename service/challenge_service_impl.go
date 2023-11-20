package service

import (
	"errors"
	model "github.com/elabosak233/pgshub/model/data"
	"github.com/elabosak233/pgshub/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ChallengeServiceImpl struct {
	ChallengeRepository repository.ChallengeRepository
	Validate            *validator.Validate
}

func NewChallengeServiceImpl(appRepository repository.AppRepository) ChallengeService {
	return &ChallengeServiceImpl{
		ChallengeRepository: appRepository.ChallengeRepository,
		Validate:            validator.New(),
	}
}

// Create implements UserService
func (t *ChallengeServiceImpl) Create(req model.Challenge) error {
	challengeModel := model.Challenge{
		Id:            uuid.NewString(),
		UploaderId:    req.UploaderId,
		Title:         req.Title,
		Description:   req.Description,
		HasAttachment: req.HasAttachment,
		AttachmentId:  req.AttachmentId,
		Flag:          req.Flag,
		Difficulty:    req.Difficulty,
		ImageName:     req.ImageName,
		IsPractice:    req.IsPractice,
		IsDynamic:     req.IsDynamic,
	}
	err := t.ChallengeRepository.Insert(challengeModel)
	return err
}

// Update implements UserService
func (t *ChallengeServiceImpl) Update(req model.Challenge) error {
	challengeData, err := t.ChallengeRepository.FindById(req.Id)
	if err != nil || challengeData.Id == "" {
		return errors.New("题目不存在")
	}
	challengeData2 := model.Challenge{
		Id:            challengeData.Id,
		Title:         req.Title,
		Description:   req.Description,
		HasAttachment: req.HasAttachment,
		AttachmentId:  req.AttachmentId,
		Flag:          req.Flag,
		Difficulty:    req.Difficulty,
		ImageName:     req.ImageName,
		IsPractice:    req.IsPractice,
		IsDynamic:     req.IsDynamic,
	}
	err = t.ChallengeRepository.Update(challengeData2)
	return err
}

// Delete implements UserService
func (t *ChallengeServiceImpl) Delete(id string) error {
	err := t.ChallengeRepository.Delete(id)
	return err
}

// FindAll implements UserService
func (t *ChallengeServiceImpl) FindAll() []model.Challenge {
	result := t.ChallengeRepository.FindAll()
	var challenges []model.Challenge
	for _, value := range result {
		challenges = append(challenges, value)
	}
	return challenges
}

// FindById implements UserService
func (t *ChallengeServiceImpl) FindById(id string) model.Challenge {
	challengeData, _ := t.ChallengeRepository.FindById(id)
	return challengeData
}
