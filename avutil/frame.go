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

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
import "C"
import "unsafe"

type Frame struct {
	cPtrAVFrame *C.struct_AVFrame
}

func NewFrame() *Frame {
	c := C.av_frame_alloc()
	return &Frame{cPtrAVFrame: c}
}

func (f *Frame) CPointer() unsafe.Pointer {
	return unsafe.Pointer(f.cPtrAVFrame)
}

//Free the frame and any dynamically allocated objects in it,
//e.g. extended_data. If the frame is reference counted, it will be
//unreferenced first.
func (f *Frame) Free() {
	if f.cPtrAVFrame != nil {
		C.av_frame_free(&f.cPtrAVFrame)
	}
}

func (f *Frame) Ref(dst *Frame) error {
	code := C.av_frame_ref(dst.cPtrAVFrame, f.cPtrAVFrame)
	if code < 0 {
		return NewErrorFromCode(int(code))
	}
	return nil
}

func (f *Frame) Unref() {
	C.av_frame_unref(f.cPtrAVFrame)
}

func (f *Frame) KeyFrame() bool {
	return f.cPtrAVFrame.key_frame == 1
}

func (f *Frame) PixelFormat() AVPixelFormat {
	return AVPixelFormat(f.cPtrAVFrame.format)
}

func (f *Frame) PictureType() AVPictureType {
	return AVPictureType(f.cPtrAVFrame.pict_type)
}
