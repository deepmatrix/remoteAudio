// Copyright © 2016 Tobias Wellnitz, DH1TW <Tobias.Wellnitz@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Stream Audio through a specfic transportation protocol",
	Long:  `Stream Audio through a specfic transportation protocol`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please select a transportation protocol (--help for available options)")
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringP("input_device", "i", "default", "Input device")
	serveCmd.PersistentFlags().StringP("output_device", "o", "default", "Output device")
	serveCmd.PersistentFlags().Float64P("samplingrate", "s", 48000, "Sampling rate for the input device")
	serveCmd.PersistentFlags().IntP("buffersize", "b", 1024, "Frames per Buffer")
	serveCmd.PersistentFlags().DurationP("input_latency", "q", time.Millisecond*5, "Input latency")
	serveCmd.PersistentFlags().DurationP("output_latency", "w", time.Millisecond*5, "Output latency")
	serveCmd.PersistentFlags().StringP("input_channels", "m", "mono", "Input Channels")
	serveCmd.PersistentFlags().StringP("output_channels", "n", "stereo", "Output Channels")
	viper.BindPFlag("audio.input_device", serveCmd.PersistentFlags().Lookup("input_device"))
	viper.BindPFlag("audio.output_device", serveCmd.PersistentFlags().Lookup("output_device"))
	viper.BindPFlag("audio.samplingrate", serveCmd.PersistentFlags().Lookup("samplingrate"))
	viper.BindPFlag("audio.buffersize", serveCmd.PersistentFlags().Lookup("buffersize"))
	viper.BindPFlag("audio.input_latency", serveCmd.PersistentFlags().Lookup("input_latency"))
	viper.BindPFlag("audio.output_latency", serveCmd.PersistentFlags().Lookup("output_latency"))
	viper.BindPFlag("audio.input_channels", serveCmd.PersistentFlags().Lookup("input_channels"))
	viper.BindPFlag("audio.output_channels", serveCmd.PersistentFlags().Lookup("output_channels"))
}
