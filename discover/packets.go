// Copyright (C) 2014 Jakob Borg and other contributors. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package discover

const (
	AnnouncementMagicV2 = 0x029E4C77
	QueryMagicV2        = 0x23D63A9A
)

type QueryV2 struct {
	Magic  uint32
	NodeID string // max:64
}

type AnnounceV2 struct {
	Magic uint32
	This  Node
	Extra []Node // max:16
}

type Node struct {
	ID        string    // max:64
	Addresses []Address // max:16
}

type Address struct {
	IP   []byte // max:16
	Port uint16
}
