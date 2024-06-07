// 答题

package dmsoft

// FaqCancel 可以把上次FaqPost的发送取消,接着下一次FaqPost
func (com *Dmsoft) FaqCancel() int {
	ret, _ := com.dm.CallMethod(DllM["FaqCancel"])
	defer ret.Clear()
	return int(ret.Val)
}

// FaqCapture 截取指定范围内的动画或者图像,并返回此句柄.
func (com *Dmsoft) FaqCapture(x1, y1, x2, y2, quality, delay, time int) int {
	ret, _ := com.dm.CallMethod(DllM["FaqCapture"], x1, y1, x2, y2, quality, delay, time)
	defer ret.Clear()
	return int(ret.Val)
}

// FaqCaptureFromFile 截取指定图片中的图像,并返回此句柄.
func (com *Dmsoft) FaqCaptureFromFile(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod(DllM["FaqCaptureFromFile"], x1, y1, x2, y2, file, quality)
	defer ret.Clear()
	return int(ret.Val)
}

// FaqCaptureString 从给定的字符串(也可以算是文字类型的问题),获取此句柄.
func (com *Dmsoft) FaqCaptureString(str string) int {
	ret, _ := com.dm.CallMethod(DllM["FaqCaptureString"], str)
	defer ret.Clear()
	return int(ret.Val)
}

// FaqFetch 获取由FaqPost发送后，由服务器返回的答案.
func (com *Dmsoft) FaqFetch() string {
	ret, _ := com.dm.CallMethod(DllM["FaqFetch"])
	defer ret.Clear()
	return ret.ToString()
}

// FaqGetSize 获取句柄所对应的数据包的大小,单位是字节
func (com *Dmsoft) FaqGetSize(handle int) int {
	ret, _ := com.dm.CallMethod(DllM["FaqGetSize"], handle)
	defer ret.Clear()
	return int(ret.Val)
}

// FaqIsPosted 用于判断当前对象是否有发送过答题(FaqPost)
func (com *Dmsoft) FaqIsPosted() int {
	ret, _ := com.dm.CallMethod(DllM["FaqIsPosted"])
	defer ret.Clear()
	return int(ret.Val)
}

// FaqPost 发送指定的图像句柄到指定的服务器,并立即返回(异步操作).
func (com *Dmsoft) FaqPost(server string, handle, requestType, timeOut int) int {
	ret, _ := com.dm.CallMethod(DllM["FaqPost"], server, handle, requestType, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

// FaqSend 发送指定的图像句柄到指定的服务器,并等待返回结果(同步等待).
func (com *Dmsoft) FaqSend(server string, handle, requestType, timeOut int) string {
	ret, _ := com.dm.CallMethod(DllM["FaqSend"], server, handle, requestType, timeOut)
	defer ret.Clear()
	return ret.ToString()
}
