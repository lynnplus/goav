/*
 * Copyright (c) 2021 The goav Author
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/lynnplus/goav/avutil"
	"unsafe"
)

type AVPacketFlag int

const (
	AvPktFlagKey     AVPacketFlag = C.AV_PKT_FLAG_KEY
	AvPktFlagCorrupt              = C.AV_PKT_FLAG_CORRUPT
	AvPktFlagDiscard              = C.AV_PKT_FLAG_DISCARD
)

type Packet struct {
	cPtrAVPacket *C.struct_AVPacket
}

func NewPacket() *Packet {
	c := C.av_packet_alloc()
	return &Packet{cPtrAVPacket: (*C.struct_AVPacket)(unsafe.Pointer(c))}
}

func (p *Packet) CPointer() unsafe.Pointer {
	return unsafe.Pointer(p.cPtrAVPacket)
}

// Free :release packet
func (p *Packet) Free() {
	C.av_packet_free(&p.cPtrAVPacket)
}

func (p *Packet) StreamIndex() int {
	return int(p.cPtrAVPacket.stream_index)
}

func (p *Packet) Flag() AVPacketFlag {
	return AVPacketFlag(p.cPtrAVPacket.flags)
}

func (p *Packet) FreePacket() {
	C.av_free_packet(p.cPtrAVPacket)
}

func (p *Packet) Ref(dst *Packet) error {
	code := C.av_packet_ref(dst.cPtrAVPacket, p.cPtrAVPacket)
	if code < 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}

func (p *Packet) Unref() {
	C.av_packet_unref(p.cPtrAVPacket)
}
