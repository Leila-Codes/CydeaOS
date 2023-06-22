package game_management

import (
	"cydeaos/events"
	"cydeaos/log"
	"fmt"
	"github.com/sirupsen/logrus"
)

var (
	logger = log.GetLogger()
)

func GMHandler(event events.Event) (events.Event, error) {
	req := event.(*events.GameManagementEvent)

	logger.WithFields(logrus.Fields{
		"event":  req,
		"action": req.Action,
	}).Info("GMHandler rx event")

	switch req.Action {
	case events.GetGame:
		{
			game, err := service.GetGame(*req.GameCode)
			if err != nil {
				return req.Error(err), err
			}

			return req.Success(game), nil
		}
	case events.CreateGame:
		{
			game := service.CreateGame(req.Config, req.Player)
			req.GameCode = &game.Code
			return req.Success(game), nil
		}
	case events.StartGame:
		{

			err := service.StartGame(*req.GameCode)
			if err != nil {
				return req.Error(err), err
			} else {
				return req.Success(nil), nil
			}
		}
	default:
		{
			return req.Error(fmt.Errorf("not yet implemented")), nil
		}
	}
}
