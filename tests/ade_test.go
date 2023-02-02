package tests

import (
	"context"
	"github.com/abc3354/CODEV-back/services/ade"
	"github.com/abc3354/CODEV-back/services/ent"
	"testing"
)

func TestADE(t *testing.T) {
	ent.Mock(t)
	defer ent.Close()

	ade.UpdateRooms()

	count, err := ent.Get().Room.Query().Count(context.Background())
	if err != nil {
		return
	}
	if count != 18 {
		t.Error(count)
	}

	ade.Update()

	count, err = ent.Get().AvailableRoom.Query().Count(context.Background())
	if err != nil {
		return
	}
	if count == 0 {
		t.Error("No available rooms found")
	}
}
