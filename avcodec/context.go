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

//#cgo pkg-config: libavcodec libavutil
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"errors"
	"github.com/lynnplus/goav/avutil"
	"unsafe"
)

type CPtrAVCodecContext *C.struct_AVCodecContext

type Context struct {
	cPtrAVCodecContext *C.struct_AVCodecContext
}

func NewContextFromC(c unsafe.Pointer) *Context {
	return &Context{cPtrAVCodecContext: (*C.struct_AVCodecContext)(c)}
}

func NewContext(codec *Codec) *Context {
	if codec.cPtrAVCodec == nil {
		panic("codec nil")
	}
	ctx := C.avcodec_alloc_context3(codec.cPtrAVCodec)
	return NewContextFromC(unsafe.Pointer(ctx))
}

func (ctx *Context) CodeID() CodecID {
	return CodecID(ctx.cPtrAVCodecContext.codec_id)
}

func (ctx *Context) Free() {
	if ctx.cPtrAVCodecContext != nil {
		C.avcodec_free_context(&ctx.cPtrAVCodecContext)
	}
}

func (ctx *Context) Open(codec *Codec) error {
	code := C.avcodec_open2(ctx.cPtrAVCodecContext, codec.cPtrAVCodec, nil)
	if code != 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}

func (ctx *Context) Close() {
	C.avcodec_close(ctx.cPtrAVCodecContext)
}

// CopyTo TODO lynn 需要使用新的接口
func (ctx *Context) CopyTo(dest *Context) error {
	if dest.cPtrAVCodecContext == nil || ctx.cPtrAVCodecContext == nil {
		return errors.New("dest param nil")
	}
	code := C.avcodec_copy_context(dest.cPtrAVCodecContext, ctx.cPtrAVCodecContext)
	if code < 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}

func (ctx *Context) DecodeVideo(frame *avutil.Frame, packet *Packet) (decompressed bool, err error) {
	var gotPic C.int
	picture := (*C.struct_AVFrame)(frame.CPointer())
	code := C.avcodec_decode_video2(ctx.cPtrAVCodecContext, picture, &gotPic, packet.cPtrAVPacket)
	if code < 0 {
		return false, avutil.NewErrorFromCode(int(code))
	}
	return gotPic != 0, nil
}
