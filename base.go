package dmsoft

func (com *Dmsoft) EnablePicCache(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["EnablePicCache"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod(DllM["GetBasePath"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod(DllM["GetDmCount"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetID() int {
	ret, _ := com.dm.CallMethod(DllM["GetID"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetLastError() int {
	ret, _ := com.dm.CallMethod(DllM["GetLastError"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetPath() string {
	ret, _ := com.dm.CallMethod(DllM["GetPath"])
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) Reg(regCode string, verInfo string) int {
	ret, _ := com.dm.CallMethod(DllM["Reg"], regCode, verInfo)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod(DllM["RegEx"], regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod(DllM["RegExNoMac"], regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod(DllM["RegNoMac"], regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayInput(mode string) int {
	ret, _ := com.dm.CallMethod(DllM["SetDisplayInput"], mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetEnumWindowDelay(delay int) int {
	ret, _ := com.dm.CallMethod(DllM["SetEnumWindowDelay"], delay)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetPath(path string) int {
	ret, _ := com.dm.CallMethod(DllM["SetPath"], path)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.dm.CallMethod(DllM["SetShowErrorMsg"], show)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SpeedNormalGraphic(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["SpeedNormalGraphic"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

// Ver get version
func (com *Dmsoft) Ver() string {
	ver, _ := com.dm.CallMethod(DllM["Ver"])
	defer ver.Clear()
	return ver.ToString()
}
