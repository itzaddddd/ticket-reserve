package service

import (
	"errors"

	"github.com/itzaddddd/ticket-reserve/repository"
	util "github.com/itzaddddd/ticket-reserve/util/constant"
)

func (s *Service) ReserveTicket(userId uint, eventId uint, numToReserve uint) error {
	err := s.repo.Transaction(func(repo repository.Repository) error {
		// check current quota
		event, err := repo.GetEvent(eventId)
		if err != nil {
			return err
		}

		if event.RemainQuota <= 0 {
			return errors.New(util.ErrQuotaLimitExceed)
		}

		// if num to serve > current quota, reserve remaining quota
		numToReserveCheckedDiffQuota := int(numToReserve)
		diffQuota := int(event.RemainQuota) - int(numToReserve)
		if diffQuota < 0 {
			numToReserveCheckedDiffQuota = -diffQuota
		}

		if err := repo.DecreseQuota(eventId, uint(numToReserveCheckedDiffQuota)); err != nil {
			return err
		}

		// insert reserve log
		if err := repo.InsertTicketReserveLog(userId, eventId, uint(numToReserveCheckedDiffQuota)); err != nil {
			return err
		}

		return nil

	})

	return err
}

func (s *Service) CancelTicket(eventId uint, numToCancel int) error {
	return nil
}
