// 窗口

package dmsoft

import (
	ole "github.com/go-ole/go-ole"
)

func (com *Dmsoft) ClientToScreen(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod(DllM["ClientToScreen"], hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	_ = intx.Clear()
	_ = inty.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnumProcess(name string) string {
	ret, _ := com.dm.CallMethod(DllM["EnumProcess"], name)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod(DllM["EnumWindow"], parent, title, className, filter)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowByProcess(processName, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod(DllM["EnumWindowByProcess"], processName, title, className, filter)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod(DllM["EnumWindowByProcessId"], pid, title, className, filter)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.dm.CallMethod(DllM["EnumWindowSuper"], spec1, flag1, type1, spec2, flag2, type2, sort)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindWindow(class, title string) int {
	ret, _ := com.dm.CallMethod(DllM["FindWindow"], class, title)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.dm.CallMethod(DllM["FindWindowByProcess"], processName, class, title)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.dm.CallMethod(DllM["FindWindowByProcessId"], processId, class, title)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.dm.CallMethod(DllM["FindWindowEx"], parent, class, title)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.dm.CallMethod(DllM["FindWindowSuper"], spec1, flag1, type1, spec2, flag2, type2)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.dm.CallMethod(DllM["GetClientRect"], hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	// 清除对象变量内存
	_ = intx1.Clear()
	_ = inty1.Clear()
	_ = intx2.Clear()
	_ = inty2.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.dm.CallMethod(DllM["GetClientSize"], hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	_ = pWidth.Clear()
	_ = pHeight.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetForegroundFocus() int {
	ret, _ := com.dm.CallMethod(DllM["GetForegroundFocus"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetForegroundWindow() int {
	ret, _ := com.dm.CallMethod(DllM["GetForegroundWindow"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetMousePointWindow() int {
	ret, _ := com.dm.CallMethod(DllM["GetMousePointWindow"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetPointWindow(x, y int) int {
	ret, _ := com.dm.CallMethod(DllM["GetPointWindow"], x, y)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetProcessInfo(pid int) string {
	ret, _ := com.dm.CallMethod(DllM["GetProcessInfo"], pid)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetSpecialWindow(flag int) int {
	ret, _ := com.dm.CallMethod(DllM["GetSpecialWindow"], flag)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindow(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod(DllM["GetWindow"], hwnd, flag)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowClass(hwnd int) string {
	ret, _ := com.dm.CallMethod(DllM["GetWindowClass"], hwnd)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetWindowProcessId(hwnd int) int {
	ret, _ := com.dm.CallMethod(DllM["GetWindowProcessId"], hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.dm.CallMethod(DllM["GetWindowProcessPath"], hwnd)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.dm.CallMethod(DllM["GetWindowRect"], hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	_ = intx1.Clear()
	_ = inty1.Clear()
	_ = intx2.Clear()
	_ = inty2.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod(DllM["GetWindowState"], hwnd, flag)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowTitle(hwnd int) string {
	ret, _ := com.dm.CallMethod(DllM["GetWindowTitle"], hwnd)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.dm.CallMethod(DllM["MoveWindow"], hwnd, x, y)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.dm.CallMethod(DllM["ScreenToClient"], hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	_ = intx.Clear()
	_ = inty.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SendPaste(hwnd int) int {
	ret, _ := com.dm.CallMethod(DllM["SendPaste"], hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SendString(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod(DllM["SendString"], hwnd, str)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SendString2(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod(DllM["SendString2"], hwnd, str)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SendStringIme(str string) int {
	ret, _ := com.dm.CallMethod(DllM["SendStringIme"], str)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SendStringIme2(hwnd int, str string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["SendStringIme2"], hwnd, str, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod(DllM["SetClientSize"], hwnd, width, height)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWindowSize"], hwnd, width, height)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWindowState"], hwnd, flag)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowText(hwnd int, title string) int {
	ret, _ := com.dm.CallMethod(DllM["SetWindowText"], hwnd, title)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.dm.CallMethod(DllM["SetWindowTransparent"], hwnd, trans)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowThreadId(hwnd int) int {
	ret, _ := com.dm.CallMethod(DllM["GetWindowThreadId"], hwnd)
	defer ret.Clear()
	return int(ret.Val)
}
