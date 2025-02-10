package main

import (
	"flag"
	"fmt"

	"os"
	"rtspranger/src/modules"
	"strings"
	"sync"

	"github.com/devchat-ai/gopool"
)

var creds []string = []string{"admin:12345", "default:tluafed", "user:user", "admin:123456", "admin:00000000", "admin:1111", "admin:admin", "admin:4321", "ubnt:ubnt", "admin:1111111", "admin:123123", "root:ikwd", "root:ikwb", "admin:qwerty", "admin:admin123", "admin:", "Dinion:", "admin:spycam", "admin:spy", "admin:security", "service:service", "Admin:1234", "admin:9999", "admin:meinsm", "Admin:123456", "administrator:", "admin:password", "root:root", "root:pass", "root:admin", "root:12345", "root:camera", "admin:admin1234", "admin:pass", "admin:1234", "admin:12345678", "admin:0000", "root:password", "user:password", "guest:guest", "admin:123", "admin:321", "admin:888888", "admin:666666", "admin:default", "admin:root", "user:admin", "guest:12345", "guest:123456", "admin:flir", "admin:wbox123", "root:system", "admin:system", "admin:netcam", "service:service", "supervisor:supervisor", "operator:operator", "root:123456", "admin:admin12345", "admin:54321", "admin:pass1234", "admin:letmein", "admin:password1", "admin:password123", "user:1234", "user:123456", "admin:changeme", "admin:12321", "admin:123456789", "admin:adminadmin", "admin:1234567", "admin:7654321", "admin:password1234", "admin:admin123456", "admin:112233", "admin:102030", "admin:admin1", "admin:passw0rd", "admin:password01", "admin:admin01", "admin:pass123", "admin:abc123", "admin:123abc", "admin:admin2019", "admin:admin2020", "admin:admin2021", "admin:admin2022", "admin:admin2023", "admin:q1w2e3r4", "admin:1q2w3e4r", "admin:qazwsx", "admin:asdfgh", "admin:zxcvbn", "admin:password!", "admin:p@ssw0rd", "admin:admin!", "admin:1234qwer", "admin:qwer1234", "admin:1qaz2wsx", "admin:jvc", "admin:qwertyuiop", "admin:asdfghjkl", "admin:zxcvbnm", "admin:1qaz!QAZ", "admin:2wsx@WSX", "admin:3edc#EDC", "admin:4rfv$RFV", "admin:5tgb%TGB", "admin:6yhn^YHN", "admin:7ujm&UJM", "admin:sunshine", "admin:654321", "admin:superman", "admin:iloveyou", "admin:monkey", "admin:11111111", "admin:000000", "admin:welcome", "admin:dragon", "admin:baseball", "admin:1qaz2wsx3edc", "admin:letmein!", "admin:abcdef", "admin:shadow", "admin:121212", "admin:lovely", "admin:freedom", "admin:jesus", "admin:ninja", "admin:charlie", "admin:banana", "admin:hunter", "admin:1234567890", "admin:summer", "admin:princess", "admin:sunshine1", "admin:football", "admin:superman1", "admin:jordan", "admin:harley", "admin:555555", "admin:yellow", "admin:purple", "admin:butterfly", "admin:7777777", "admin:ashley", "admin:6666666", "admin:tigger", "admin:jessica", "admin:michael", "admin:password2", "admin:password3", "admin:bubbles", "admin:987654321", "admin:11223344", "admin:pepper", "admin:daniel", "admin:aaaaaa", "admin:chocolate", "admin:chicken", "admin:maggie", "admin:thomas", "admin:fliradmin", "admin:ginger", "admin:liverpool", "admin:computer", "666666:666666", "888888:888888", "admin:88888888", "admin:1234abcd", "admin:mercedes", "admin:corvette", "admin:bigdog", "admin:cheese", "admin:matthew", "admin:cocacola", "admin:barney", "admin:maxwell", "admin:coffee", "admin:scooby", "admin:rabbit", "admin:mickey", "admin:987654", "admin:snoopy", "admin:cookie", "admin:stephen", "admin:arsenal", "admin:dolphin", "admin:rainbow", "admin:knight", "admin:midnight", "admin:123123123", "admin:badboy", "admin:123321", "admin:andrew", "admin:147258", "admin:99999999", "admin:sparky", "admin:silver", "admin:angela", "admin:coffee1", "admin:college", "admin:boston", "admin:jennifer", "admin:4444444", "admin:junior", "admin:redsox", "admin:gandalf", "admin:696969", "admin:password4", "admin:martin", "admin:jeffrey", "admin:george"}
var routes []string = []string{"/0/1:1/main", "/0/usrnm:pwd/main", "/0/video2", "/0/video1", "/5", "/4", "/3", "/2", "/1", "/1.AMP", "/1/h264major", "/2/h264major", "/1/stream1", "/1080p", "/11", "/12", "/125", "/1440p", "/480p", "/4K", "/666", "/720p", "/AVStream1_1", "/CAM_ID.password.mp2", "/CH001.sdp", "/GetData.cgi", "/HD", "/HighResolutionVideo", "/LowResolutionVideo", "/MediaInput/h264", "/MediaInput/mpeg4", "/ONVIF/MediaInput", "/ONVIF/MediaInput?profile=4_def_profile6", "/StdCh1", "/Streaming/Channels/3", "/Streaming/Channels/101/", "/Streaming/Channels/102", "/Streaming/Channels/2", "/Streaming/Channels/1", "/Streaming/Channels/301", "/Streaming/Channels/501", "/Streaming/channels/201", "/Streaming/channels/401", "/StreamingSetting?version=1.0&action=getRTSPStream&ChannelID=1&ChannelName=Channel1", "/VideoInput/1/h264/1", "/VideoInput/1/mpeg4/1", "/access_code", "/access_name_for_stream_1_to_5", "/api/mjpegvideo.cgi", "/av0_0", "/av2", "/avc", "/avn=2", "/axis-media/media.amp", "/axis-media/media.amp?camera=1", "/axis-media/media.amp?videocodec=h264", "/cam", "/cam/realmonitor", "/cam/realmonitor?channel=0&subtype=0", "/cam/realmonitor?channel=1&subtype=0", "/cam/realmonitor?channel=1&subtype=1", "/cam/realmonitor?channel=1&subtype=1&unicast=true&proto=Onvif", "/cam0", "/cam0_0", "/cam0_1", "/cam1", "/cam1/h264", "/cam1/h264/multicast", "/cam1/mjpeg", "/cam1/mpeg4", "/cam1/mpeg4?user='username'&pwd='password'", "/camera.stm", "/ch0", "/ch00/0", "/ch001.sdp", "/ch01.264", "/ch01.264?", "/ch01.264?ptype=tcp", "/ch0_0.h264", "/ch0_unicast_firststream", "/ch0_unicast_secondstream", "/ch1-s1", "/ch1/0", "/ch1_0", "/ch2/0", "/ch2_0", "/ch3/0", "/ch3_0", "/ch4/0", "/ch4_0", "/channel1", "/channel2", "/channel3", "/gnz_media/main", "/h264", "/h264.sdp", "/h264/ch1/sub/av_stream", "/h264/media.amp", "/h264Preview_01_main", "/h264Preview_01_sub", "/h264_vga.sdp", "/image.mpg", "/img/media.sav", "/img/media.sav?channel=1", "/img/video.asf", "/img/video.sav", "/ioImage/1", "/ipcam.sdp", "/ipcam_h264.sdp", "/ipcam_mjpeg.sdp", "/live", "/live.sdp", "/live/av0", "/live/ch0", "/live/ch00_0", "/live/ch01_0", "/live/h264", "/live/main", "/live/main0", "/live/mpeg4", "/live1.sdp", "/live3.sdp", "/live_mpeg4.sdp", "/live_st1", "/livestream", "/livestream/", "/main", "/media", "/media.amp", "/media.amp?streamprofile=Profile1", "/media/media.amp", "/media/video1", "/medias2", "/mjpeg/media.smp", "/mp4", "/mpeg/media.amp", "/mpeg4", "/mpeg4/1/media.amp", "/mpeg4/media.amp", "/mpeg4/media.smp", "/mpeg4unicast", "/mpg4/rtsp.amp", "/multicaststream", "/now.mp4", "/nph-h264.cgi", "/nphMpeg4/g726-640x", "/nphMpeg4/g726-640x48", "/nphMpeg4/g726-640x480", "/nphMpeg4/nil-320x240", "/onvif-media/media.amp", "/onvif1", "/pass@10.0.0.5:6667/blinkhd", "/play1.sdp", "/play2.sdp", "/profile2/media.smp", "/profile5/media.smp", "/rtpvideo1.sdp", "/rtsp_live0", "/rtsp_live1", "/rtsp_live2", "/rtsp_tunnel", "/rtsph264", "/snap.jpg", "/stream", "/stream.sdp", "/stream/0", "/stream/1", "/stream/live.sdp", "/stream1", "/streaming/channels/0", "/streaming/channels/1", "/streaming/channels/101", "/Streaming/Unicast/channels/101", "/tcp/av0_0", "/test", "/tmpfs/auto.jpg", "/trackID=1", "/ucast/11", "/udp/av0_0", "/udp/unicast/aiphone_H264", "/udpstream", "/user.pin.mp2", "/user=admin&password=&channel=1&stream=0.sdp?", "/user=admin&password=&channel=1&stream=0.sdp?real_stream", "/user=admin_password=?????_channel=1_stream=0.sdp?real_stream", "/user=admin_password=R5XFY888_channel=1_stream=0.sdp?real_stream", "/user_defined", "/v2", "/video", "/video.3gp", "/video.h264", "/video.mjpg", "/video.mp4", "/video.pro1", "/video.pro2", "/video.pro3", "/video0", "/video0.sdp", "/video1", "/video1+audio1", "/video1.sdp", "/videoMain", "/videoinput_1/h264_1/media.stm", "/videostream.asf", "/vis", "/wfov", "/PSIA/Streaming/channels/1?videoCodecType=MPEG4", "/cam1/onvif-h264", "/h264_stream", "/rtsph2641080p", "/0/av0", "/0/av1", "/0/av2", "/0/av3", "/live/0", "/live/1", "/live/2", "/live/3", "/main/av_stream", "/sub/av_stream", "/ONVIF/channels/1", "/ONVIF/channels/2", "/ONVIF/channels/3", "/1/h264minor", "/1/h264/main", "/1/h264/sub", "/2/h264/main", "/2/h264/sub", "/3/h264/main", "/3/h264/sub", "/channel2", "/channel3", "/channel4", "/video2", "/video3", "/video4", "/live/av1", "/live/av2", "/live/av3", "/cam/realmonitor?channel=2&subtype=0", "/cam/realmonitor?channel=3&subtype=0", "/cam/realmonitor?channel=4&subtype=0", "/Streaming/Channels/3", "/Streaming/Channels/4", "/Streaming/channels/202", "/Streaming/channels/203", "/Streaming/channels/204", "/1/h264/2", "/1/h264/3", "/1/h264/4", "/live/ch02_0", "/live/ch03_0", "/live/ch04_0", "/Streaming/Channels/1/Preview", "/Streaming/Channels/2/Preview", "/Streaming/Channels/3/Preview", "/Streaming/Channels/4/Preview", "/cam/realmonitor?channel=2&subtype=1", "/cam/realmonitor?channel=3&subtype=1", "/cam/realmonitor?channel=4&subtype=1", "/1/h264/preview", "/2/h264/preview", "/3/h264/preview", "/4/h264/preview", "/video2.sdp", "/video3.sdp", "/video4.sdp", "/ONVIF/channels/4", "/ONVIF/channels/5", "/ONVIF/channels/6"}

func Raw_Connect(target string) bool {
	auth := modules.AUTH{Username: "", Password: "", Method: "0"}

	client := modules.CreateRTSPClient(target, "554", 10, auth)

	if client.Connect("554") {

		cre := client.Auth.Username + ":" + client.Auth.Password
		if client.Authorize(cre, routes[1]) {
			if client.OkAuth() {

				return true
			} else {

				return false
			}
		}

	}

	return false
}

func Operation(target string, output *os.File) {
	var mutex sync.Mutex
	fmt.Println("Trying ", target)

	if Raw_Connect(target) {
		auth := modules.AUTH{Username: "", Password: ""}
		client := modules.CreateRTSPClient(target, "554", 5, auth)

		cre := client.Auth.Username + ":" + client.Auth.Password
		if client.Connect("554") {
			for _, route := range routes {
				if client.Status == 1 && client.Socket != nil {
					if client.Authorize(cre, route) {
						if client.OkRoute(route) {
							mutex.Lock()
							fmt.Println("R okay ", cre, route, client.GetRTSPUrl())
							output.WriteString(client.GetRTSPUrl() + "\n")

							mutex.Unlock()
							break
						}
					}

				}
			}
		}
		if client.Socket != nil {
			client.Socket.Close()
		}

	}
}
func main() {

	var file string
	out_file, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		out_file.Close()
	}
	flag.StringVar(&file, "file", "", "")
	flag.Parse()

	content, err := os.ReadFile(file)
	lines := strings.Split(string(content), "\n")

	if err != nil {
		//Do something
		fmt.Println("Unable to read input file... ")
		panic(err)
	}

	pool := gopool.NewGoPool(300)
	defer pool.Release()

	for _, target := range lines {

		pool.AddTask(func() (interface{}, error) {
			Operation(target, out_file)

			return nil, nil
		})
	}

	pool.Wait()

}
