package videoSplit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type VideoInfoData struct {
	Stream []Streams `json:"streams"` //视频数据流,包括视频音频字幕
	Format Format    `json:"format"`  //格式化
}

type Streams struct {
	CodecName string `json:"codec_name"`        //编码
	Width     int    `json:"width,omitempty"`   //宽
	Height    int    `json:"height,omitempty"`  //高
	PixFmt    string `json:"pix_fmt,omitempty"` // 分辨率
	Duration  string `json:"duration"`          //视频时长
}

type Format struct {
	BitRate string `json:"bit_rate"` //速率
}

// 视频资源解析
func PreTreatmentVideo(input string) (int, float64, error) {
	// 获取视频信息
	videoData, err := GetVideoInfo(input)
	if err != nil {
		return 0, 0, err
	}
	// 判断格式是否为h264
	if videoData.Stream[0].CodecName != "h264" {
		return 0, 0, errors.New("not h264")
	}

	//计算最大分辨率
	width := videoData.Stream[0].Width
	height := videoData.Stream[0].Height
	// 获取宽高最大的
	quality := Min(getWidthRes(width), getHeigthRes(height))

	//获取视频时长,单位s
	duration, _ := strconv.ParseFloat(videoData.Stream[0].Duration, 64)

	return quality, duration, err
}
func GetVideoInfo(input string) (videoData VideoInfoData, err error) {
	cmd := exec.Command("ffprobe", "-i", input, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams")
	out, err := runCmd(cmd)
	if err != nil {
		return videoData, err
	}

	// 反序列化
	err = json.Unmarshal(out.Bytes(), &videoData)
	if err != nil {
		return videoData, err
	}

	return videoData, nil
}
func runCmd(cmd *exec.Cmd) (bytes.Buffer, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return out, errors.New(stderr.String())
	}

	return out, nil
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 获取宽度支持的最大分辨率
func getWidthRes(width int) int {
	//1920*1080
	if width >= 1920 {
		return 1080
	}
	// 1280*720
	if width >= 1280 {
		return 720
	}
	//720*480
	if width >= 720 {
		return 480
	}
	return 360
}

// 获取高度支持的最大分辨率
func getHeigthRes(height int) int {
	//1920*1080
	if height >= 1080 {
		return 1080
	}
	// 1280*720
	if height >= 720 {
		return 720
	}
	//720*480
	if height >= 480 {
		return 480
	}
	return 360
}

// 转码 资源id,宽高最大值,保存的文件夹
func VideoTransCoding(resourceId uint, quality int, dirName, videoPath string) {
	localDir := dirName + "/"
	inputFile := localDir + videoPath

	// 提取音频
	audioFile, err := ExtractingAudio(inputFile, localDir)
	if err != nil {
		completeTransCoding(resourceId, 2300)
		return
	}

	// 生成不同分辨率的MP4
	videoFiles, err := PressingVideo(inputFile, localDir, quality)
	if err != nil {
		completeTransCoding(resourceId, 0)
		return
	}

	// 生成dash分片,参数为:视频文件,音频文件,保存文件夹地址,文件夹名称
	err = GenerateDash(videoFiles, audioFile, localDir, dirName)
	if err != nil {
		completeTransCoding(resourceId, 0)
		return
	}

	// 删除临时文件
	os.Remove(audioFile)
	for _, v := range videoFiles {
		os.Remove(v)
	}

	// 完成转码
	completeTransCoding(resourceId, 1)
}

// 提取音频
func ExtractingAudio(inputFile, outputDir string) (string, error) {
	output := outputDir + "audio.m4a"
	cmd := exec.Command("ffmpeg", "-hide_banner", "-i", inputFile, "-vn", "-c", "copy", output)
	_, err := runCmd(cmd)
	return output, err
}

// 压制视频
func PressingVideo(inputFile, outputDir string, quality int) ([]string, error) {
	outputFileList := make([]string, 0)
	command := []string{"-hide_banner", "-i", inputFile, "-crf", "20"}
	switch quality {
	case 1080:
		output1080 := outputDir + "tmp_1080p_3000k.mp4"
		outputFileList = append(outputFileList, output1080)
		command = append(command, "-c:v", "libx264", "-an", "-s", "1920x1080", "-r", "30000/1001", "-b:v", "3000k", output1080)
		fallthrough
	case 720:
		output720 := outputDir + "tmp_720p_2000k.mp4"
		outputFileList = append(outputFileList, output720)
		command = append(command, "-c:v", "libx264", "-an", "-s", "1080x720", "-r", "30000/1001", "-b:v", "2000k", output720)
		fallthrough
	case 480:
		output480 := outputDir + "tmp_480p_900k.mp4"
		outputFileList = append(outputFileList, output480)
		command = append(command, "-c:v", "libx264", "-an", "-s", "854x480", "-r", "30000/1001", "-b:v", "900k", output480)
		fallthrough
	case 360:
		output360 := outputDir + "tmp_360p_500k.mp4"
		outputFileList = append(outputFileList, output360)
		command = append(command, "-c:v", "libx264", "-an", "-s", "640x360", "-r", "30000/1001", "-b:v", "500k", output360)
	}

	// 翻转 outputFileList ，从低分辨率到高分辨率
	for i, j := 0, len(outputFileList)-1; i < j; i, j = i+1, j-1 {
		outputFileList[i], outputFileList[j] = outputFileList[j], outputFileList[i]
	}

	_, err := runCmd(exec.Command("ffmpeg", command...))
	return outputFileList, err
}

// TODO 视频解析失败处理
func completeTransCoding(resourceId uint, status int) {
	/*
		// 更新资源状态
		UpadteResourceStatus(resourceId, status)
		// 获取资源信息
		resource := SelectResourceByID(resourceId)
		// 获取转码中资源的数量
		count := SelectResourceCountByStatus(resource.Vid, common.VIDEO_PROCESSING)
		// 如果没有转码中的视频，则更新视频为待审核
		if count == 0 {
			// 获取视频审核状态
			video := GetVideoInfo(resource.Vid)
			if video.Status == common.SUBMIT_REVIEW {
				UpadteVideoStatus(video.ID, common.WAITING_REVIEW)
			}
		}
	*/

}
func GenerateDash(videoFiles []string, audioFile, outputDir, outputName string) error {
	mapCommand := make([]string, 0)
	command := make([]string, 0)

	mpdName := outputDir + "index.mpd"
	initStreamName := outputName + "-init-$RepresentationID$.m4s"
	chunkStreamName := outputName + "-$RepresentationID$-$Number%05d$.m4s"

	// 添加视频
	for i := 0; i < len(videoFiles); i++ {
		command = append(command, "-i", videoFiles[i])
		mapCommand = append(mapCommand, "-map", strconv.Itoa(i))
	}

	// 添加音频
	command = append(command, "-i", audioFile)
	mapCommand = append(mapCommand, "-map", strconv.Itoa(len(videoFiles)))

	// 合并命令
	command = append(command, "-c", "copy")
	command = append(command, mapCommand...)
	command = append(command, "-f", "dash", "-init_seg_name", initStreamName, "-media_seg_name", chunkStreamName, mpdName)
	str := strings.Join(command, " ")
	fmt.Println(str)
	_, err := runCmd(exec.Command("ffmpeg", command...))
	return err
}
