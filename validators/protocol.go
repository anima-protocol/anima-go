package validators

import (
	"fmt"

	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/utils"
)

func ValidateProtocol(anima *models.Protocol) error {
	if !utils.InArray(anima.Chain, models.AVAILABLE_CHAIN) {
		return fmt.Errorf("chain unavailable")
	}
	return nil
}
