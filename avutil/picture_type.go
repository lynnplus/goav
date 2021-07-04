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
//#include <libavutil/avutil.h>
import "C"

type AVPictureType C.enum_AVPictureType

const (
	AvPictureTypeNone AVPictureType = C.AV_PICTURE_TYPE_NONE ///< Undefined
	AvPictureTypeI                  = C.AV_PICTURE_TYPE_I    ///< Intra
	AvPictureTypeP                  = C.AV_PICTURE_TYPE_P    ///< Predicted
	AvPictureTypeB                  = C.AV_PICTURE_TYPE_B    ///< Bi-dir predicted
	AvPictureTypeS                  = C.AV_PICTURE_TYPE_S    ///< S(GMC)-VOP MPEG4
	AvPictureTypeSi                 = C.AV_PICTURE_TYPE_SI   ///< Switching Intra
	AvPictureTypeSp                 = C.AV_PICTURE_TYPE_SP   ///< Switching Predicted
	AvPictureTypeBi                 = C.AV_PICTURE_TYPE_BI   ///< BI type
)

func (p AVPictureType) String() string {
	switch p {
	case AvPictureTypeNone:
		return "AVPictureType(None)"
	case AvPictureTypeI:
		return "AVPictureType(I)"
	case AvPictureTypeP:
		return "AVPictureType(P)"
	case AvPictureTypeB:
		return "AVPictureType(B)"
	case AvPictureTypeS:
		return "AVPictureType(S)"
	case AvPictureTypeSi:
		return "AVPictureType(SI)"
	case AvPictureTypeSp:
		return "AVPictureType(SP)"
	case AvPictureTypeBi:
		return "AVPictureType(BI)"
	default:
		return "AVPictureType(unknown)"
	}
}
