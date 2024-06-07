// 文字识别

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

// AddDict 给指定的字库中添加一条字库信息
func (com *Dmsoft) AddDict(index int, dictInfo string) int {
	ret, _ := com.dm.CallMethod(DllM["AddDict"], index, dictInfo)
	defer ret.Clear()
	return int(ret.Val)
}

// ClearDict 清空指定的字库
func (com *Dmsoft) ClearDict(index int) int {
	ret, _ := com.dm.CallMethod(DllM["ClearDict"], index)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableShareDict(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableShareDict"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FetchWord(x1, y1, x2, y2 int, color, word string) string {
	ret, _ := com.dm.CallMethod(DllM["FetchWord"], x1, y1, x2, y2, color, word)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindStr"], x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrE(x1, y1, x2, y2 int, str string, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrE"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrEx"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrExS"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindStrFast"], x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrFastE(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrFastE"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrFastEx"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrFastExS"], x1, y1, x2, y2, str, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrFastS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindStrFastS"], x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindStrS"], x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrWithFont(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, fontName string, fontSize, flag int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["FindStrWithFont"], x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindStrWithFontE(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrWithFontE"], x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStrWithFontEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod(DllM["FindStrWithFontEx"], x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDict(index, fontIndex int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDict"], index, fontIndex)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDictCount(index int) int {
	ret, _ := com.dm.CallMethod(DllM["GetDictCount"], index)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetDictInfo(str, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDictInfo"], str, fontName, fontSize, flag)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetNowDict() int {
	ret, _ := com.dm.CallMethod(DllM["GetNowDict"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetResultCount(str string) int {
	ret, _ := com.dm.CallMethod(DllM["GetResultCount"], str)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["GetResultPos"], str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultCount(str string) int {
	ret, _ := com.dm.CallMethod(DllM["GetWordResultCount"], str)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod(DllM["GetWordResultPos"], str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	_ = x.Clear()
	_ = y.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWordResultStr(str string, index int) string {
	ret, _ := com.dm.CallMethod(DllM["GetWordResultCount"], str, index)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetWords(x1, y1, x2, y2 int, color string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["GetWords"], x1, y1, x2, y2, color, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetWordsNoDict(x1, y1, x2, y2 int, color string) string {
	ret, _ := com.dm.CallMethod(DllM["GetWordsNoDict"], x1, y1, x2, y2, color)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["Ocr"], x1, y1, x2, y2, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["OcrEx"], x1, y1, x2, y2, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) OcrExOne(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["OcrExOne"], x1, y1, x2, y2, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) OcrInFile(x1, y1, x2, y2 int, picName, colorFormat string, sim float32) string {
	ret, _ := com.dm.CallMethod(DllM["OcrInFile"], x1, y1, x2, y2, picName, colorFormat, sim)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) SaveDict(index int, file string) int {
	ret, _ := com.dm.CallMethod(DllM["SaveDict"], index, file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetColGapNoDict(colGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetColGapNoDict"], colGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDict(index int, file string) int {
	ret, _ := com.dm.CallMethod(DllM["SetDict"], index, file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDictMem(index, addr, size int) int {
	ret, _ := com.dm.CallMethod(DllM["SetDictMem"], index, addr, size)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDictPwd(pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["SetDictPwd"], pwd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetExactOcr(exactOcr int) int {
	ret, _ := com.dm.CallMethod(DllM["SetExactOcr"], exactOcr)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetMinColGap(minColGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetMinColGap"], minColGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetMinRowGap(minRowGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetMinRowGap"], minRowGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetRowGapNoDict(RowGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetRowGapNoDict"], RowGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWordGap(wordGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWordGap"], wordGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWordGapNoDict(wordGap int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWordGapNoDict"], wordGap)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWordLineHeight(lineHeight int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWordLineHeight"], lineHeight)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWordLineHeightNoDict(lineHeight int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWordLineHeightNoDict"], lineHeight)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) UseDict(index int) int {
	ret, _ := com.dm.CallMethod(DllM["UseDict"], index)
	defer ret.Clear()
	return int(ret.Val)
}
