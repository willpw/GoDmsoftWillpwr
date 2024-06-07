// 后台

package dmsoft

func (com *Dmsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["BindWindow"], hwnd, display, mouse, keypad, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["BindWindowEx"], hwnd, display, mouse, keypad, public, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DownCPU(rate int) int {
	ret, _ := com.dm.CallMethod(DllM["DownCpu"], rate)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableBind(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableBind"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableFakeActive(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableFakeActive"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableIme(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableIme"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadMsg(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableKeypadMsg"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadPatch(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableKeypadPatch"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableKeypadSync"], enable, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseMsg(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableMouseMsg"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableMouseSync"], enable, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealKeypad(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableRealKeypad"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableRealMouse"], enable, mousedelay, mousestep)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableSpeedDx(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnableSpeedDx"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ForceUnBindWindow() int {
	ret, _ := com.dm.CallMethod(DllM["ForceUnBindWindow"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetBindWindow() int {
	ret, _ := com.dm.CallMethod(DllM["GetBindWindow"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetFps() int {
	ret, _ := com.dm.CallMethod(DllM["GetFps"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) HackSpeed(rate float32) int {
	ret, _ := com.dm.CallMethod(DllM["HackSpeed"], rate)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IsBind(hwnd int) int {
	ret, _ := com.dm.CallMethod(DllM["IsBind"], hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockDisplay(lock int) int {
	ret, _ := com.dm.CallMethod(DllM["LockDisplay"], lock)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockInput(lock int) int {
	ret, _ := com.dm.CallMethod(DllM["LockInput"], lock)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod(DllM["LockMouseRect"], x1, y1, x2, y2)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetAero(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["SetAero"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayDelay(time int) int {
	ret, _ := com.dm.CallMethod(DllM["SetDisplayDelay"], time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.dm.CallMethod(DllM["SetDisplayRefreshDelay"], time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SwitchBindWindow(hwnd int) int {
	ret, _ := com.dm.CallMethod(DllM["SwitchBindWindow"], hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) UnBindWindow() int {
	ret, _ := com.dm.CallMethod(DllM["UnBindWindow"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetInputDm(dm_id, rx, ry int) int {
	ret, _ := com.dm.CallMethod(DllM["SetInputDm"], dm_id, rx, ry)
	defer ret.Clear()
	return int(ret.Val)
}
