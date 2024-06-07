// 系统

package dmsoft

func (com *Dmsoft) Beep(f, duration int) int {
	ret, _ := com.dm.CallMethod(DllM["Beep"], f, duration)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CheckFontSmooth() int {
	ret, _ := com.dm.CallMethod(DllM["CheckFontSmooth"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CheckUAC() int {
	ret, _ := com.dm.CallMethod(DllM["CheckUAC"])
	defer ret.Clear()
	return int(ret.Val)
}

// Delay 函数不推荐使用
// 建议time.Sleep()
func (com *Dmsoft) Delay(mis int) int {
	ret, _ := com.dm.CallMethod(DllM["Delay"])
	defer ret.Clear()
	return int(ret.Val)
}
func (com *Dmsoft) Delays(misMin, misMax int) int {
	ret, _ := com.dm.CallMethod(DllM["Delays"], misMin, misMax)
	defer ret.Clear()
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *Dmsoft) DisableCloseDisplayAndSleep() int {
	ret, _ := com.dm.CallMethod(DllM["DisableCloseDisplayAndSleep"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisableFontSmooth() int {
	ret, _ := com.dm.CallMethod(DllM["DisableFontSmooth"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisablePowerSave() int {
	ret, _ := com.dm.CallMethod(DllM["DisablePowerSave"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisableScreenSave() int {
	ret, _ := com.dm.CallMethod(DllM["DisableScreenSave"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableFontSmooth() int {
	ret, _ := com.dm.CallMethod(DllM["EnableFontSmooth"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ExitOs(types int) int {
	ret, _ := com.dm.CallMethod(DllM["ExitOs"], types)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetClipboard() string {
	ret, _ := com.dm.CallMethod(DllM["GetClipboard"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetCPUType() int {
	ret, _ := com.dm.CallMethod(DllM["GetCpuType"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCPUUsage() int {
	ret, _ := com.dm.CallMethod(DllM["GetCpuUsage"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetDir(types int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDir"], types)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *Dmsoft) GetDiskModel(index int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDiskModel"], index)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *Dmsoft) GetDiskReversion(index int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDiskReversion"], index)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *Dmsoft) GetDiskSerial(index int) string {
	ret, _ := com.dm.CallMethod(DllM["GetDiskSerial"], index)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDisplayInfo() string {
	ret, _ := com.dm.CallMethod(DllM["GetDisplayInfo"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDPI() int {
	ret, _ := com.dm.CallMethod(DllM["GetDPI"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetLocale() int {
	ret, _ := com.dm.CallMethod(DllM["GetLocale"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetMachineCode() string {
	ret, _ := com.dm.CallMethod(DllM["GetMachineCode"])
	defer ret.Clear()
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *Dmsoft) GetMachineCodeNoMac() string {
	ret, _ := com.dm.CallMethod(DllM["GetMachineCodeNoMac"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetMemoryUsage() int {
	ret, _ := com.dm.CallMethod(DllM["GetMemoryUsage"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetNetTime() string {
	ret, _ := com.dm.CallMethod(DllM["GetNetTime"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetNetTimeByIP(ip string) string {
	ret, _ := com.dm.CallMethod(DllM["GetNetTimeByIp"], ip)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetOsBuildNumber() int {
	ret, _ := com.dm.CallMethod(DllM["GetOsBuildNumber"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetOsType() int {
	ret, _ := com.dm.CallMethod(DllM["GetOsType"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenDepth() int {
	ret, _ := com.dm.CallMethod(DllM["GetScreenDepth"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenHeight() int {
	ret, _ := com.dm.CallMethod(DllM["GetScreenHeight"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenWidth() int {
	ret, _ := com.dm.CallMethod(DllM["GetScreenWidth"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetTime() int {
	ret, _ := com.dm.CallMethod(DllM["GetTime"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Is64Bit() int {
	ret, _ := com.dm.CallMethod(DllM["Is64Bit"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IsSurrpotVt() int {
	ret, _ := com.dm.CallMethod(DllM["IsSurrpotVt"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Play(mediaFile string) int {
	ret, _ := com.dm.CallMethod(DllM["Play"], mediaFile)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RunApp(appPath string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["RunApp"], appPath, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetClipboard(value string) int {
	ret, _ := com.dm.CallMethod(DllM["SetClipboard"], value)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayAcceler(level int) int {
	ret, _ := com.dm.CallMethod(DllM["SetDisplayAcceler"], level)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetLocale() int {
	ret, _ := com.dm.CallMethod(DllM["SetLocale"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetScreen(width, height, depth int) int {
	ret, _ := com.dm.CallMethod(DllM["SetScreen"], width, height, depth)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetUAC(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["SetUAC"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.dm.CallMethod(DllM["ShowTaskBarIcon"], hwnd, isShow)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Stop(id int) int {
	ret, _ := com.dm.CallMethod(DllM["Stop"], id)
	defer ret.Clear()
	return int(ret.Val)
}
