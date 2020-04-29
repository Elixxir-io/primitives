////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package switchboard

import (
	"gitlab.com/elixxir/primitives/id"
)

type ListeningQueue chan Item

// ListenChannel sets up a listening queue and adds it to the switchboard.
func (lm *Switchboard) ListenChannel(
	messageType int32, sender *id.ID, channelBufferSize int) (id string,
	messageQueue ListeningQueue) {
	messageQueue = make(ListeningQueue, channelBufferSize)
	id = lm.Register(sender, messageType, messageQueue)
	return id, messageQueue
}

// Hear allows multiple threads to write to the buffer simultaneously through
// the switchboard.
func (l ListeningQueue) Hear(item Item, isHeardElsewhere bool, i ...interface{}) {
	l <- item
}
