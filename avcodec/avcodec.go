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
import "unsafe"

type (
	Descriptor C.struct_AVCodecDescriptor
)

type Codec struct {
	cPtrAVCodec *C.struct_AVCodec
}

func newCodec(in *C.struct_AVCodec) *Codec {
	return &Codec{cPtrAVCodec: (*C.struct_AVCodec)(unsafe.Pointer(in))}
}

func FindDecoderByID(codecID CodecID) *Codec {
	decoder := C.avcodec_find_decoder((C.enum_AVCodecID)(codecID))
	if decoder == nil {
		return nil
	}
	return newCodec(decoder)
}
