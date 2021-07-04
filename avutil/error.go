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

//#cgo pkg-config: libavformat libavutil
//#include <libavformat/avformat.h>
//#include <libavutil/error.h>
//#include <stdlib.h>
import "C"
import "unsafe"

var _ error = (*AVError)(nil)

type AVError struct {
	code int
	str  string
}

func (e *AVError) Error() string {
	return e.str
}

func NewErrorFromCode(code int) error {
	return &AVError{code: code, str: strerror(code)}
}

func strerror(code int) string {
	size := C.size_t(256)
	buf := (*C.char)(C.av_mallocz(size))
	defer C.av_free(unsafe.Pointer(buf))
	if C.av_strerror(C.int(code), buf, size-1) == 0 {
		return C.GoString(buf)
	}
	return "unknown error"
}
