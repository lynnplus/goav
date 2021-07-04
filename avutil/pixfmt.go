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
//#include <libavutil/pixfmt.h>
import "C"

type AVPixelFormat C.enum_AVPixelFormat

const (
	AvPixFmtNone      AVPixelFormat = C.AV_PIX_FMT_NONE
	AvPixFmtYuv420p                 = C.AV_PIX_FMT_YUV420P   ///< planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AvPixFmtYuyv422                 = C.AV_PIX_FMT_YUYV422   ///< packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AvPixFmtRgb24                   = C.AV_PIX_FMT_RGB24     ///< packed RGB 8:8:8, 24bpp, RGBRGB...
	AvPixFmtBgr24                   = C.AV_PIX_FMT_BGR24     ///< packed RGB 8:8:8, 24bpp, BGRBGR...
	AvPixFmtYuv422p                 = C.AV_PIX_FMT_YUV422P   ///< planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AvPixFmtYuv444p                 = C.AV_PIX_FMT_YUV444P   ///< planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AvPixFmtYuv410p                 = C.AV_PIX_FMT_YUV410P   ///< planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AvPixFmtYuv411p                 = C.AV_PIX_FMT_YUV411P   ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	AvPixFmtGray8                   = C.AV_PIX_FMT_GRAY8     ///<        Y        ,  8bpp
	AvPixFmtMonowhite               = C.AV_PIX_FMT_MONOWHITE ///<        Y        ,  1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb
	AvPixFmtMonoblack               = C.AV_PIX_FMT_MONOBLACK ///<        Y        ,  1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb
	AvPixFmtPal8                    = C.AV_PIX_FMT_PAL8      ///< 8 bit with PIX_FMT_RGB32 palette
	AvPixFmtYuvj420p                = C.AV_PIX_FMT_YUVJ420P  ///< planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of PIX_FMT_YUV420P and setting color_range
	AvPixFmtYuvj422p                = C.AV_PIX_FMT_YUVJ422P  ///< planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of PIX_FMT_YUV422P and setting color_range
	AvPixFmtYuvj444p                = C.AV_PIX_FMT_YUVJ444P  ///< planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of PIX_FMT_YUV444P and setting color_range
	AvPixFmtUyvy422                 = C.AV_PIX_FMT_UYVY422   ///< packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	AvPixFmtUyyvyy411               = C.AV_PIX_FMT_UYYVYY411 ///< packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	AvPixFmtBgr8                    = C.AV_PIX_FMT_BGR8      ///< packed RGB 3:3:2,  8bpp, (msb)2B 3G 3R(lsb)
	AvPixFmtBgr4                    = C.AV_PIX_FMT_BGR4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AvPixFmtRgb8                    = C.AV_PIX_FMT_RGB8      ///< packed RGB 3:3:2,  8bpp, (msb)2R 3G 3B(lsb)
	AvPixFmtRgb4                    = C.AV_PIX_FMT_RGB4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AvPixFmtNv12                    = C.AV_PIX_FMT_NV12      ///< planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AvPixFmtNv21                    = C.AV_PIX_FMT_NV21      ///< as above, but U and V bytes are swapped
	AvPixFmtArgb                    = C.AV_PIX_FMT_ARGB      ///< packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	AvPixFmtRgba                    = C.AV_PIX_FMT_RGBA      ///< packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	AvPixFmtAbgr                    = C.AV_PIX_FMT_ABGR      ///< packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	AvPixFmtBgra                    = C.AV_PIX_FMT_BGRA      ///< packed BGRA 8:8:8:8, 32bpp, BGRABGRA...
	AvPixFmtVdpau                   = C.AV_PIX_FMT_VDPAU     ///< HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface
	AvPixFmtYvyu422                 = C.AV_PIX_FMT_YVYU422   ///< packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb
	AvPixFmt0rgb                    = C.AV_PIX_FMT_0RGB      ///< packed RGB 8:8:8, 32bpp, 0RGB0RGB...
	AvPixFmtRgb0                    = C.AV_PIX_FMT_RGB0      ///< packed RGB 8:8:8, 32bpp, RGB0RGB0...
	AvPixFmt0bgr                    = C.AV_PIX_FMT_0BGR      ///< packed BGR 8:8:8, 32bpp, 0BGR0BGR...
	AvPixFmtBgr0                    = C.AV_PIX_FMT_BGR0      ///< packed BGR 8:8:8, 32bpp, BGR0BGR0...
	AvPixFmtYuva444p                = C.AV_PIX_FMT_YUVA444P  ///< planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	AvPixFmtYuva422p                = C.AV_PIX_FMT_YUVA422P  ///< planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
)

func (f AVPixelFormat) String() string {
	switch f {
	case AvPixFmtNone:
		return "AvPixFmtNone"
	case AvPixFmtYuv420p:
		return "AvPixFmtYuv420p"
	case AvPixFmtYuyv422:
		return "AvPixFmtYuyv422"
	case AvPixFmtRgb24:
		return "AvPixFmtRgb24"
	case AvPixFmtBgr24:
		return "AvPixFmtBgr24"
	case AvPixFmtYuv422p:
		return "AvPixFmtYuv422p"
	case AvPixFmtYuv444p:
		return "AvPixFmtYuv444p"
	case AvPixFmtYuv410p:
		return "AvPixFmtYuv410p"
	case AvPixFmtYuv411p:
		return "AvPixFmtYuv411p"
	case AvPixFmtGray8:
		return "AvPixFmtGray8"
	case AvPixFmtMonowhite:
		return "AvPixFmtMonowhite"
	case AvPixFmtMonoblack:
		return "AvPixFmtMonoblack"
	case AvPixFmtPal8:
		return "AvPixFmtPal8"
	case AvPixFmtYuvj420p:
		return "AvPixFmtYuvj420p"
	case AvPixFmtYuvj422p:
		return "AvPixFmtYuvj422p"
	case AvPixFmtYuvj444p:
		return "AvPixFmtYuvj444p"
	case AvPixFmtUyvy422:
		return "AvPixFmtUyvy422"
	case AvPixFmtUyyvyy411:
		return "AvPixFmtUyyvyy411"
	case AvPixFmtBgr8:
		return "AvPixFmtBgr8"
	case AvPixFmtBgr4:
		return "AvPixFmtBgr4"
	case AvPixFmtRgb8:
		return "AvPixFmtRgb8"
	case AvPixFmtRgb4:
		return "AvPixFmtRgb4"
	case AvPixFmtNv12:
		return "AvPixFmtNv12"
	case AvPixFmtNv21:
		return "AvPixFmtNv21"
	case AvPixFmtArgb:
		return "AvPixFmtArgb"
	case AvPixFmtRgba:
		return "AvPixFmtRgba"
	case AvPixFmtAbgr:
		return "AvPixFmtAbgr"
	case AvPixFmtBgra:
		return "AvPixFmtBgra"
	case AvPixFmtVdpau:
		return "AvPixFmtVdpau"
	case AvPixFmtYvyu422:
		return "AvPixFmtYvyu422"
	case AvPixFmt0rgb:
		return "AvPixFmt0rgb"
	case AvPixFmtRgb0:
		return "AvPixFmtRgb0"
	case AvPixFmt0bgr:
		return "AvPixFmt0bgr"
	case AvPixFmtBgr0:
		return "AvPixFmtBgr0"
	case AvPixFmtYuva444p:
		return "AvPixFmtYuva444p"
	case AvPixFmtYuva422p:
		return "AvPixFmtYuva422p"
	default:
		return "unknown"
	}
}
