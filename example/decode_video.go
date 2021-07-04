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

package main

import (
	"fmt"
	"github.com/lynnplus/goav/avcodec"
	"github.com/lynnplus/goav/avformat"
	"github.com/lynnplus/goav/avutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	avformat.RegisterAll()

	pFormatCtx := avformat.NewContext()
	defer pFormatCtx.Free()

	e := pFormatCtx.OpenInput("test3.mp4", nil, nil)
	defer pFormatCtx.CloseInput()
	if e != nil {
		panic(e)
	}

	e = pFormatCtx.FindStreamInfo()
	if e != nil {
		panic(e)
	}

	videoIndex := pFormatCtx.FindBestStream(avutil.MediaTypeVideo, -1, -1)
	if videoIndex < 0 {
		panic("Cannot find a video stream in the input file")
	}
	stream := pFormatCtx.Streams()[videoIndex]

	pCodecCtx := stream.CodecContext()

	fmt.Println(pCodecCtx.CodeID())

	if pCodecCtx.CodeID() != avcodec.AvCodecIdH264 {
		panic("codec non h264")
	}

	pCodec := avcodec.FindDecoderByID(pCodecCtx.CodeID())
	if pCodec == nil {
		panic("Unsupported codec!")
	}

	pCtx := avcodec.NewContext(pCodec)
	if pCtx == nil {

		panic("nil")
	}

	e = pCodecCtx.CopyTo(pCtx)
	if e != nil {
		panic(e)
	}

	e = pCtx.Open(pCodec)
	if e != nil {
		panic(e)
	}
	defer pCtx.Close()

	pFrame := avutil.NewFrame()
	pkt := avcodec.NewPacket()

	for {
		err := pFormatCtx.ReadFrame(pkt)
		if err != nil {
			fmt.Println(err)
			break
		}

		if pkt.StreamIndex() != videoIndex {
			fmt.Println("skip----------")
			pkt.Unref()
			//pFrame.Free()
			continue
		}

		decompressed, er := pCtx.DecodeVideo(pFrame, pkt)
		if er != nil {
			fmt.Println(er)
			continue
		}
		if !decompressed {
			fmt.Println("decompressed fail")
			pkt.Unref()
			continue
		}

		fmt.Println(pFrame.KeyFrame(), pFrame.PictureType())

		pkt.Unref()
		pFrame.Unref()
	}

	pFrame.Free()
	pkt.Free()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGSTOP, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
