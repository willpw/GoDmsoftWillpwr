// 鼠键

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *Dmsoft) EnableMouseAccuracy(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableMouseAccuracy"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod(DllM["GetCursorPos"], &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	_ = intx.Clear()
	_ = inty.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCursorShape() string {
	ret, _ := com.dm.CallMethod(DllM["GetCursorShape"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetCursorShapeEx(types int) string {
	ret, _ := com.dm.CallMethod(DllM["GetCursorShapeEx"], types)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetCursorSpot() string {
	ret, _ := com.dm.CallMethod(DllM["GetCursorSpot"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetKeyState(vkCode int) int {
	ret, _ := com.dm.CallMethod(DllM["GetKeyState"], vkCode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetMouseSpeed() int {
	ret, _ := com.dm.CallMethod(DllM["GetMouseSpeed"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyDown(vkCode int) int {
	ret, _ := com.dm.CallMethod(DllM["KeyDown"], vkCode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.dm.CallMethod(DllM["KeyDownChar"], keyStr)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyPress(vkCode int) int {
	ret, _ := com.dm.CallMethod(DllM["KeyPress"], vkCode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.dm.CallMethod(DllM["KeyPressChar"], keyStr)
	defer ret.Clear()
	return int(ret.Val)
}

// KeyPressStr KeyPressStr指定的字符串序列，依次按顺序按下其中的字符
func (com *Dmsoft) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.dm.CallMethod(DllM["KeyPressStr"], keyStr, delay)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyUp(vkCode int) int {
	ret, _ := com.dm.CallMethod(DllM["KeyUp"], vkCode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.dm.CallMethod(DllM["KeyUpChar"], keyStr)
	defer ret.Clear()
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *Dmsoft) LeftClick() int {
	ret, _ := com.dm.CallMethod(DllM["LeftClick"])
	defer ret.Clear()
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *Dmsoft) LeftDoubleClick() int {
	ret, _ := com.dm.CallMethod(DllM["LeftDoubleClick"])
	defer ret.Clear()
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *Dmsoft) LeftDown() int {
	ret, _ := com.dm.CallMethod(DllM["LeftDown"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LeftUp() int {
	ret, _ := com.dm.CallMethod(DllM["LeftUp"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MiddleClick() int {
	ret, _ := com.dm.CallMethod(DllM["MiddleClick"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MiddleDown() int {
	ret, _ := com.dm.CallMethod(DllM["MiddleDown"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MiddleUp() int {
	ret, _ := com.dm.CallMethod(DllM["MiddleUp"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MoveR(rx, ry int) int {
	ret, _ := com.dm.CallMethod(DllM["MoveR"], rx, ry)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MoveTo(x, y int) int {
	ret, _ := com.dm.CallMethod(DllM["MoveTo"], x, y)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.dm.CallMethod(DllM["MoveToEx"], x, y, w, h)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) RightClick() int {
	ret, _ := com.dm.CallMethod(DllM["RightClick"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RightDown() int {
	ret, _ := com.dm.CallMethod(DllM["RightDown"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RightUp() int {
	ret, _ := com.dm.CallMethod(DllM["RightUp"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod(DllM["SetKeypadDelay"], types, delay)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseDelay(types string, delay int) int {
	ret, _ := com.dm.CallMethod(DllM["SetMouseDelay"], types, delay)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseSpeed(speed int) int {
	ret, _ := com.dm.CallMethod(DllM["SetMouseSpeed"], speed)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetSimMode(mode int) int {
	ret, _ := com.dm.CallMethod(DllM["SetSimMode"], mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.dm.CallMethod(DllM["SetSimMode"], vkCode, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) WheelDown() int {
	ret, _ := com.dm.CallMethod(DllM["WheelDown"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) WheelUp() int {
	ret, _ := com.dm.CallMethod(DllM["WheelUp"])
	defer ret.Clear()
	return int(ret.Val)
}
