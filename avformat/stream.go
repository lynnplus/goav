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

package avformat

//#cgo pkg-config: libavformat libavcodec
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/lynnplus/goav/avcodec"
	"unsafe"
)

type Stream struct {
	cPtrAVStream *C.struct_AVStream
}

func newStreamFromC(in *C.struct_AVStream) *Stream {
	return &Stream{cPtrAVStream: (*C.struct_AVStream)(unsafe.Pointer(in))}
}

func (s *Stream) CodecContext() *avcodec.Context {
	if s.cPtrAVStream == nil || s.cPtrAVStream.codec == nil {
		return nil
	}
	p := unsafe.Pointer(s.cPtrAVStream.codec)
	return avcodec.NewContextFromC(p)
}
