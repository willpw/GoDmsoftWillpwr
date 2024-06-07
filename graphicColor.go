// 图色

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *Dmsoft) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.dm.CallMethod(DllM["AppendPicAddr"], picInfo, addr, size)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod(DllM["Capture"], x1, y1, x2, y2, file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.dm.CallMethod(DllM["CaptureGif"], x1, y1, x2, y2, file, delay, time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod(DllM["CaptureJpg"], x1, y1, x2, y2, file, quality)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod(DllM["CapturePng"], x1, y1, x2, y2, file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CapturePre(file string) int {
	ret, _ := com.dm.CallMethod(DllM["CapturePre"], file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod(DllM["CmpColor"], x, y, color, sim)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableDisplayDebug"], enableDebug)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableFindPicMultithread(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableFindPicMultithread"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableGetColorByCapture"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindColor"], x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorBlock(x1, y1, x2, y2 int, color string, sim float32, count, width, height int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindColorBlock"], x1, y1, x2, y2, color, sim, count, width, height, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float32, count, width, height int) string {
	ret, _ := com.dm.CallMethod(DllM["FindColorBlockEx"], x1, y1, x2, y2, color, sim, count, width, height)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindColorE(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindColorE"], x1, y1, x2, y2, color, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindColorEx"], x1, y1, x2, y2, color, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindMulColor(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod(DllM["FindMulColor"], x1, y1, x2, y2, color, sim)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindMultiColor"], x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColorE(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindMultiColorE"], x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindMultiColorEx"], x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindPic"], x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicE(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindPicE"], x1, y1, x2, y2, picName, deltaColor, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindPicEx"], x1, y1, x2, y2, picName, deltaColor, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod(DllM["FindPicExS"], x1, y1, x2, y2, picName, deltaColor, sim, dir)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindPicMem(x1, y1, x2, y2 int, pic_info, delta_color string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindPicMem"], x1, y1, x2, y2, pic_info, delta_color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

// func (com *Dmsoft)FindPicMemE(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}
// func (com *Dmsoft)FindPicMemEx(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}

func (com *Dmsoft) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindPicS"], x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	defer ret.Clear()
	return ret.ToString()
}

// func (com *Dmsoft)FindShape(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) int{}
// func (com *Dmsoft)FindShapeE(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
// func (com *Dmsoft)FindShapeEx(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
func (com *Dmsoft) FreePic(pic_name string) int {
	ret, _ := com.dm.CallMethod(DllM["FreePic"], pic_name)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetAveHSV(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod(DllM["GetAveHSV"], x1, y1, x2, y2)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetAveRGB(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod(DllM["GetAveRGB"], x1, y1, x2, y2)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetColor(x, y int) string {
	ret, _ := com.dm.CallMethod(DllM["GetColor"], x, y)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetColorBGR(x, y int) string {
	ret, _ := com.dm.CallMethod(DllM["GetColorBGR"], x, y)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetColorHSV(x, y int) string {
	ret, _ := com.dm.CallMethod(DllM["GetColorHSV"], x, y)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetColorNum(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod(DllM["GetColorNum"], x1, y1, x2, y2, color, sim)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetPicSize(picName string) string {
	ret, _ := com.dm.CallMethod(DllM["GetPicSize"], picName)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetScreenData(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod(DllM["GetScreenData"], x1, y1, x2, y2)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenDataBmp(x1, y1, x2, y2 int, data, size *int) int {
	d := ole.NewVariant(ole.VT_I4, int64(*data))
	s := ole.NewVariant(ole.VT_I4, int64(*size))
	ret, _ := com.dm.CallMethod(DllM["GetScreenDataBmp"], x1, y1, x2, y2, &d, &s)
	*data = int(d.Val)
	*size = int(s.Val)
	d.Clear()
	s.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ImageToBmp(pic_name, bmp_name string) int {
	ret, _ := com.dm.CallMethod(DllM["ImageToBmp"], pic_name, bmp_name)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IsDisplayDead(x1, y1, x2, y2, time int) int {
	ret, _ := com.dm.CallMethod(DllM["IsDisplayDead"], x1, y1, x2, y2, time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LoadPic(pic_name string) int {
	ret, _ := com.dm.CallMethod(DllM["LoadPic"], pic_name)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LoadPicByte(addr, size int, pic_name string) int {
	ret, _ := com.dm.CallMethod(DllM["LoadPicByte"], addr, size, pic_name)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MatchPicName(picName string) string {
	ret, _ := com.dm.CallMethod(DllM["MatchPicName"], picName)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.dm.CallMethod(DllM["SetExcludeRegion"], mode, info)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetPicPwd(pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["SetExcludeRegion"], pwd)
	defer ret.Clear()
	return int(ret.Val)
}
func (com *Dmsoft) RGB2BGR(rgb_color string) string {
	ret, _ := com.dm.CallMethod(DllM["RGB2BGR"], rgb_color)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) BGR2RGB(bgr_color string) string {
	ret, _ := com.dm.CallMethod(DllM["BGR2RGB"], bgr_color)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) SetFindPicMultithreadCount(count int) int {
	ret, _ := com.dm.CallMethod(DllM["SetFindPicMultithreadCount"], count)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetFindPicMultithreadLimit(limit int) int {
	ret, _ := com.dm.CallMethod(DllM["SetFindPicMultithreadLimit"], limit)
	defer ret.Clear()
	return int(ret.Val)
}

// 增加了接口FindPicSim FindPicSimEx FindPicSimE和FindPicSimMem FindPicSimMemEx FindPicSimMemE
