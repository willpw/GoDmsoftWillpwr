// 汇编

package dmsoft

func (com *Dmsoft) AsmAdd(asmIns string) int {
	ret, _ := com.dm.CallMethod(DllM["AsmAdd"], asmIns)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmCall(hwnd, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["AsmCall"], hwnd, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmCallEx(hwnd, mode int, baseAddr string) int64 {
	ret, _ := com.dm.CallMethod(DllM["AsmCallEx"], hwnd, mode, baseAddr)
	defer ret.Clear()
	return ret.Val
}

func (com *Dmsoft) AsmClear() int {
	ret, _ := com.dm.CallMethod(DllM["AsmClear"])
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmSetTimeout(timeOut, param int) int {
	ret, _ := com.dm.CallMethod(DllM["AsmSetTimeout"], timeOut, param)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Assemble(timeOut int64, is64bit int) string {
	ret, _ := com.dm.CallMethod(DllM["Assemble"], timeOut, is64bit)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) DisAssemble(asmCode string, baseAddr int64, is64bit int) string {
	ret, _ := com.dm.CallMethod(DllM["DisAssemble"], asmCode, baseAddr, is64bit)
	defer ret.Clear()
	return ret.ToString()
}
