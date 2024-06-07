package dmsoft

func (com *Dmsoft) DoubleToData(value float64) string {
	ret, _ := com.dm.CallMethod("DoubleToData", value)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindData(hwnd int, addrRange string, data string) string {
	ret, _ := com.dm.CallMethod("FindData", hwnd, addrRange, data)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindDataEx(hwnd int, addrRange string, data string, step int, multiThread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindDataEx", hwnd, addrRange, data, step, multiThread, mode)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindDouble(hwnd int, addrRange string, doubleValueMin float64, doubleValueMax float64) string {
	ret, _ := com.dm.CallMethod("FindDouble", hwnd, addrRange, doubleValueMin, doubleValueMax)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindDoubleEx(hwnd int, addrRange string, doubleValueMin float64, doubleValueMax float64, step int, multiThread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindDoubleEx", hwnd, addrRange, doubleValueMin, doubleValueMax, step, multiThread, mode)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindFloat(hwnd int, addrRange string, floatValueMin float32, floatValueMax float32) string {
	ret, _ := com.dm.CallMethod("FindFloat", hwnd, addrRange, floatValueMin, floatValueMax)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindFloatEx(hwnd int, addrRange string, floatValueMin float32, floatValueMax float32, step int, multiThread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindFloatEx", hwnd, addrRange, floatValueMin, floatValueMax, step, multiThread, mode)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindInt(hwnd int, addrRange string, intValueMin int64, intValueMax int64, itype int) string {
	ret, _ := com.dm.CallMethod("FindInt", hwnd, addrRange, intValueMin, intValueMax, itype)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindIntEx(hwnd int, addr_range string, intValueMin int64, intValueMax int64, itype int, step int, multiThread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindIntEx", hwnd, addr_range, intValueMin, intValueMax, itype, step, multiThread, mode)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindString(hwnd int, addrRange string, stringValue string, itype int) string {
	ret, _ := com.dm.CallMethod("FindString", hwnd, addrRange, stringValue, itype)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FindStringEx(hwnd int, addrRange string, stringValue string, itype int, step int, multiThread int, mode int) string {
	ret, _ := com.dm.CallMethod("FindStringEx", hwnd, addrRange, stringValue, itype, step, multiThread, mode)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) FloatToData(value float32) string {
	ret, _ := com.dm.CallMethod("FloatToData", value)
	defer ret.Clear()
	return ret.ToString()
}

// string DoubleToData(value)
// string FindData(hwnd, addr_range, data)
// string FindDataEx(hwnd, addr_range, data,step,multi_thread,mode)
// string FindDouble(hwnd, addr_range, double_value_min, double_value_max)
// string FindDoubleEx(hwnd, addr_range, double_value_min, double_value_max,step,multi_thread,mode)
// string FindFloat(hwnd, addr_range, float_value_min, float_value_max)
// string FindFloatEx(hwnd, addr_range, float_value_min, float_value_max,step,multi_thread,mode)
// string FindInt(hwnd, addr_range, int_value_min, int_value_max,type)
// string FindIntEx(hwnd, addr_range, int_value_min, int_value_max,type,step,multi_thread,mode)
// string FindString(hwnd, addr_range, string_value,type)
// string FindStringEx(hwnd, addr_range, string_value,type,step,multi_thread,mode)
// string FloatToData(value)
// long FreeProcessMemory(hwnd)

func (com *Dmsoft) FreeProcessMemory(hwnd int) int {
	ret, _ := com.dm.CallMethod("FreeProcessMemory", hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCommandLine(hwnd int) string {
	ret, _ := com.dm.CallMethod(DllM["GetCommandLine"], hwnd)
	defer ret.Clear()
	return ret.ToString()
}

// LONGLONG GetModuleBaseAddr(hwnd,module)
// long GetModuleSize(hwnd,module)
// LONGLONG GetRemoteApiAddress(hwnd,base_addr,fun_name)

func (com *Dmsoft) GetModuleBaseAddr(hwnd int, module string) int64 {
	ret, _ := com.dm.CallMethod("GetModuleBaseAddr", hwnd, module)
	defer ret.Clear()
	return ret.Val
}

func (com *Dmsoft) GetModuleSize(hwnd int, module string) int {
	ret, _ := com.dm.CallMethod("GetModuleSize", hwnd, module)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetRemoteApiAddress(hwnd int, base_addr int64, fun_name string) int64 {
	ret, _ := com.dm.CallMethod("GetRemoteApiAddress", hwnd, base_addr, fun_name)
	defer ret.Clear()
	return ret.Val
}

// long Int64ToInt32(value)
// string IntToData(value,type)
func (com *Dmsoft) Int64ToInt32(value int64) int {
	ret, _ := com.dm.CallMethod("Int64ToInt32", value)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IntToData(value int64, itype int) string {
	ret, _ := com.dm.CallMethod("IntToData", value, itype)
	defer ret.Clear()
	return ret.ToString()
}
func (com *Dmsoft) OpenProcess(pid int) int {
	ret, _ := com.dm.CallMethod(DllM["OpenProcess"], pid)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadData(hwnd int, addr string, le int) string {
	ret, _ := com.dm.CallMethod(DllM["ReadData"], hwnd, addr, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddr(hwnd int, addr int64, le int) string {
	ret, _ := com.dm.CallMethod(DllM["ReadDataAddr"], hwnd, addr, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddrToBin(hwnd int, addr int64, le int) int {
	ret, _ := com.dm.CallMethod(DllM["ReadDataAddrToBin"], hwnd, addr, le)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadDataToBin(hwnd int, addr string, le int) int {
	ret, _ := com.dm.CallMethod(DllM["ReadDataToBin"], hwnd, addr, le)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadDouble(hwnd int, addr string) float64 {
	ret, _ := com.dm.CallMethod(DllM["ReadDouble"], hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float64)
}

func (com *Dmsoft) ReadDoubleAddr(hwnd int, addr int64) float64 {
	ret, _ := com.dm.CallMethod(DllM["ReadDoubleAddr"], hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float64)
}

func (com *Dmsoft) ReadFloat(hwnd int, addr string) float32 {
	ret, _ := com.dm.CallMethod(DllM["ReadFloat"], hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float32)
}

func (com *Dmsoft) ReadFloatAddr(hwnd int, addr int64) float32 {
	ret, _ := com.dm.CallMethod(DllM["ReadFloatAddr"], hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float32)
}

func (com *Dmsoft) ReadInt(hwnd int, addr string, typ int) int64 {
	ret, _ := com.dm.CallMethod(DllM["ReadInt"], hwnd, addr, typ)
	defer ret.Clear()
	return ret.Val

}

func (com *Dmsoft) ReadIntAddr(hwnd int, addr int64, typ int) int64 {
	ret, _ := com.dm.CallMethod(DllM["ReadIntAddr"], hwnd, addr, typ)
	defer ret.Clear()
	return ret.Val

}

func (com *Dmsoft) ReadString(hwnd int, addr string, typ, le int) string {
	ret, _ := com.dm.CallMethod(DllM["ReadString"], hwnd, addr, typ, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadStringAddr(hwnd int, addr int64, typ, le int) string {
	ret, _ := com.dm.CallMethod(DllM["ReadStringAddr"], hwnd, addr, typ, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) SetMemoryFindResultToFile(file string) int {
	ret, _ := com.dm.CallMethod(DllM["SetMemoryFindResultToFile"], file)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) SetMemoryHwndAsProcessId(en int) int {
	ret, _ := com.dm.CallMethod(DllM["SetMemoryHwndAsProcessId"], en)
	defer ret.Clear()
	return int(ret.Val)
}

// string StringToData(value,type)
// long TerminateProcess(pid)

func (com *Dmsoft) StringToData(value string, itype int) string {
	ret, _ := com.dm.CallMethod("StringToData", value, itype)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) TerminateProcess(pid int) int {
	ret, _ := com.dm.CallMethod("TerminateProcess", pid)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) VirtualAllocEx(hwnd int, addr int64, size, typ int) int64 {
	ret, _ := com.dm.CallMethod(DllM["VirtualAllocEx"], hwnd, addr, size, typ)
	defer ret.Clear()
	return ret.Val
}

func (com *Dmsoft) VirtualFreeEx(hwnd int, addr int64) int {
	ret, _ := com.dm.CallMethod(DllM["VirtualFreeEx"], hwnd, addr)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) VirtualProtectEx(hwnd int, addr int64, size, typ, oldProtect int) int {
	ret, _ := com.dm.CallMethod(DllM["VirtualProtectEx"], hwnd, addr, size, typ, oldProtect)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) VirtualQueryEx(hwnd int, addr int64, pmbi int) string {
	ret, _ := com.dm.CallMethod(DllM["VirtualQueryEx"], hwnd, addr, pmbi)
	defer ret.Clear()
	return ret.ToString()

}

func (com *Dmsoft) WriteData(hwnd int, addr, data string) int {
	ret, _ := com.dm.CallMethod(DllM["WriteData"], hwnd, addr, data)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) WriteDataAddr(hwnd int, addr, data string) int {
	ret, _ := com.dm.CallMethod(DllM["WriteDataAddr"], hwnd, addr, data)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDataAddrFromBin(hwnd int, addr int64, data, le int) int {
	ret, _ := com.dm.CallMethod(DllM["WriteDataAddrFromBin"], hwnd, addr, data, le)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDataFromBin(hwnd int, addr string, data, le int) int {
	ret, _ := com.dm.CallMethod(DllM["WriteDataFromBin"], hwnd, addr, data, le)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDouble(hwnd int, addr string, v float64) int {
	ret, _ := com.dm.CallMethod(DllM["WriteDouble"], hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) WriteDoubleAddr(hwnd int, addr int64, v float64) int {
	ret, _ := com.dm.CallMethod(DllM["WriteDoubleAddr"], hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteFloat(hwnd int, addr string, v float32) int {
	ret, _ := com.dm.CallMethod(DllM["WriteFloat"], hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteFloatAddr(hwnd int, addr int64, v float32) int {
	ret, _ := com.dm.CallMethod(DllM["WriteFloatAddr"], hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteInt(hwnd int, addr string, typ int, v int64) int {
	ret, _ := com.dm.CallMethod(DllM["WriteInt"], hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteIntAddr(hwnd int, addr int64, typ int, v int64) int {
	ret, _ := com.dm.CallMethod(DllM["WriteIntAddr"], hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteString(hwnd int, addr string, typ int, v string) int {
	ret, _ := com.dm.CallMethod(DllM["WriteString"], hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteStringAddr(hwnd int, addr int64, typ int, v string) int {
	ret, _ := com.dm.CallMethod(DllM["WriteStringAddr"], hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
