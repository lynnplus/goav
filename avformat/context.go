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

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
//#include <stdlib.h>
//static const AVStream *go_av_streams_get(const AVStream **streams, unsigned int n)
//{
//  return streams[n];
//}
import "C"
import (
	"github.com/lynnplus/goav/avcodec"
	"github.com/lynnplus/goav/avutil"
	"unsafe"
)

type Context struct {
	cPtrAVFormatContext *C.struct_AVFormatContext
}

// NewContext Allocate an AVFormatContext.
// Context.Free() can be used to free the context and everything
// allocated by the framework within it.
func NewContext() *Context {
	ctx := C.avformat_alloc_context()
	p := unsafe.Pointer(ctx)
	return &Context{cPtrAVFormatContext: (*C.struct_AVFormatContext)(p)}
}

func (ctx *Context) Free() {
	if ctx.cPtrAVFormatContext != nil {
		C.avformat_free_context(ctx.cPtrAVFormatContext)
		ctx.cPtrAVFormatContext = nil
	}
}

func (ctx *Context) OpenInput(fileName string, inFmt *InputFormat, options *avutil.Dictionary) error {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	code := C.avformat_open_input(&ctx.cPtrAVFormatContext, cFileName, nil, nil)
	if code != 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}

func (ctx *Context) CloseInput() {
	C.avformat_close_input(&ctx.cPtrAVFormatContext)
}

// FindStreamInfo Read packets of a media file to get stream information. This
//is useful for file formats with no headers such as MPEG. This
//function also computes the real framerate in case of MPEG-2 repeat
//frame mode.
//The logical file position is not changed by this function;
//examined packets may be buffered for later processing.
func (ctx *Context) FindStreamInfo() error {
	code := C.avformat_find_stream_info(ctx.cPtrAVFormatContext, nil)
	if code != 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}

// FindBestStream Find the "best" stream in the file.
//The best stream is determined according to various heuristics as the most
//likely to be what the user expects.
//If the decoder parameter is non-NULL, av_find_best_stream will find the
//default decoder for the stream's codec; streams for which no decoder can
//be found are ignored.
// @wantedSnb user-requested stream number, or -1 for automatic selection
// @relatedStream try to find a stream related (eg. in the same program) to this one, or -1 if none
// @decoderRet if non-NULL, returns the decoder for the selected stream
// @flags none are currently defined
//
//the non-negative stream number in case of success,
// AVERROR_STREAM_NOT_FOUND if no stream with the requested type could be found,
// AVERROR_DECODER_NOT_FOUND if streams were found but no decoder
func (ctx *Context) FindBestStream(mediaType avutil.AVMediaType, wantedSnb int, relatedStream int) int {
	code := C.av_find_best_stream(ctx.cPtrAVFormatContext, C.enum_AVMediaType(mediaType), C.int(wantedSnb), C.int(relatedStream), nil, C.int(0))
	return int(code)
}

// NumberOfStreams Number of elements in AVFormatContext.streams.
//Set by avformat_new_stream(), must not be modified by any other code.
func (ctx *Context) NumberOfStreams() uint {
	return uint(ctx.cPtrAVFormatContext.nb_streams)
}

func (ctx *Context) Streams() []*Stream {
	num := ctx.NumberOfStreams()
	if num <= 0 {
		return nil
	}
	s := make([]*Stream, num, num)
	ptr := unsafe.Pointer(*ctx.cPtrAVFormatContext.streams)
	size := C.int(C.sizeof_AVStream)
	for i := 0; i < int(num); i++ {
		p := (*C.AVStream)(unsafe.Pointer(uintptr(ptr) + uintptr(size*C.int(i))))
		stream := newStreamFromC(p)
		s[i] = stream
	}
	return s
}

func (ctx *Context) ReadFrame(pkt *avcodec.Packet) error {
	code := C.av_read_frame(ctx.cPtrAVFormatContext, (*C.struct_AVPacket)(pkt.CPointer()))
	if code < 0 {
		return avutil.NewErrorFromCode(int(code))
	}
	return nil
}
