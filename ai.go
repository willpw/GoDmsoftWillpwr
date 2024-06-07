package dmsoft

import "github.com/go-ole/go-ole"

/*
整个Ai接口是通过外挂模块来实现的.   ai.module通过后台下载.

在调用Ai接口之前,需要保证LoadAi或者LoadAiMemory接口成功.

另外,由于是内存加载,卸载模块会导致程序出异常,所以只提供了加载接口(LoadAi和LoadAiMemory),没有提供卸载接口. 所以要特别注意,千万不能在LoadAi或者LoadAiMemory后卸载对象. 这样会导致模块没卸载,从而内存泄漏. 除非是程序结束.

关于Yolo部分,由于模型很占用内存空间,并且检测接口很占用内存和CPU,所以在多线程中AiYoloDetectXX系列接口不建议频繁调用,更不可以用此接口来代替找图等接口.

如果只是单线程调用,或者同一时间只有一个线程调用AiYoloDetectXX系列接口,那么没什么影响.

内部实现上,Yolo是使用了全局的静态模型. 所有的对象是共用模型. 所以在多线程的使用上要特别注意.

对于同一个序号的模型,在多线程上是排队执行的. 尤其是同一个脚本程序控制很多窗口时,那么多线程执行AiYoloDetect系列接口时,并且使用的序号是相同的,那么效率会大打折扣.

另外在脚本程序下,识别效率会不如Yolo综合工具里的效率. 因为32位程序的优化不如64位.

另外也不要问我为何没有GPU加速,因为cuda不支持32位程序.

具体的使用例子请查看我录制的视频.



注:
如果想提高检测效率，两个途径
1. 使用更小更快的预训练模型. 比如yolov5n
2. 运行的机器CPU核心数越多,效率越高. 因为检测函数内部是多线程执行的.

如果想提高检测精度,两个途径
1. 使用更大但是更慢的预训练模型. 比如yolov5x
2. 对于每个类尽可能多的提供训练图片. 尽可能多的提供各种复杂背景下的训练图片. 尽可能对每个类在各种复杂背景下都提供训练图片. 训练的轮次可以稍微多一些.

如果发现自己训练后的模型,会越训练越差,说明是你训练的过头了(过拟合),减少轮次,重新训练


*/

// AiYoloDetectObjects 需要先加载Ai模块. 在指定范围内检测对象.
func (com *Dmsoft) AiYoloDetectObjects(x1, y1, x2, y2 int, prob, iou float32) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjects"], x1, y1, x2, y2, prob, iou)
	defer ret.Clear()
	return ret.ToString()
}

// AiYoloDetectObjectsToDataBmp 需要先加载Ai模块. 在指定范围内检测对象,把结果输出到BMP图像数据.用于二次开发.
func (com *Dmsoft) AiYoloDetectObjectsToDataBmp(x1, y1, x2, y2 int, prob, iou float32, data, size *int, mode int) int {
	d := ole.NewVariant(ole.VT_I4, int64(*data))
	s := ole.NewVariant(ole.VT_I4, int64(*size))
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjectsToDataBmp"], x1, y1, x2, y2, prob, iou, &d, &s, mode)
	*data = int(d.Val)
	*size = int(s.Val)
	_ = d.Clear()
	_ = s.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloDetectObjectsToFile 需要先加载Ai模块. 在指定范围内检测对象,把结果输出到指定的BMP文件.
func (com *Dmsoft) AiYoloDetectObjectsToFile(x1, y1, x2, y2 int, prob, iou float32, file string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjectsToFile"], x1, y1, x2, y2, prob, iou, file, mode)
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloFreeModel 需要先加载Ai模块. 卸载指定的模型
func (com *Dmsoft) AiYoloFreeModel(index int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloFreeModel"], index)
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloObjectsToString 需要先加载Ai模块. 把通过AiYoloDetectObjects或者是AiYoloSortsObjects的结果,按照顺序把class信息连接输出.
func (com *Dmsoft) AiYoloObjectsToString(objects string) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloObjectsToString"], objects)
	defer ret.Clear()
	return ret.ToString()
}

// AiYoloSetModel 需要先加载Ai模块. 从文件加载指定的模型.
func (com *Dmsoft) AiYoloSetModel(index int, file, pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetModel"], index, file, pwd)
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloSetModelMemory 需要先加载Ai模块. 从内存加载指定的模型. 仅支持dmx格式的内存
func (com *Dmsoft) AiYoloSetModelMemory(index, data, size int, pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetModelMemory"], index, data, size, pwd)
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloSetVersion 需要先加载Ai模块. 设置Yolo的版本
func (com *Dmsoft) AiYoloSetVersion(ver string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetVersion"], ver)
	defer ret.Clear()
	return int(ret.Val)
}

// AiYoloSortsObjects 需要先加载Ai模块. 把通过AiYoloDetectObjects的结果进行排序. 排序按照从上到下,从左到右.
func (com *Dmsoft) AiYoloSortsObjects(objects string, height int) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSortsObjects"], objects, height)
	defer ret.Clear()
	return ret.ToString()
}

// AiYoloUseModel 需要先加载Ai模块. 切换当前使用的模型序号.用于AiYoloDetectXX等系列接口.
func (com *Dmsoft) AiYoloUseModel(index int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloUseModel"], index)
	defer ret.Clear()
	return int(ret.Val)
}

// LoadAi 加载Ai模块. Ai模块从后台下载.
func (com *Dmsoft) LoadAi(file string) int {
	ret, _ := com.dm.CallMethod(DllM["LoadAi"], file)
	defer ret.Clear()
	return int(ret.Val)
}

// LoadAiMemory 从内存加载Ai模块. Ai模块从后台下载.
func (com *Dmsoft) LoadAiMemory(data, size int) int {
	ret, _ := com.dm.CallMethod(DllM["LoadAiMemory"], data, size)
	defer ret.Clear()
	return int(ret.Val)
}

/*
long AiEnableFindPicWindow(enable)
long AiFindPic(x1, y1, x2, y2, pic_name,sim, dir,intX, intY)
string AiFindPicEx(x1, y1, x2, y2, pic_name,sim, dir)
long AiFindPicMem(x1, y1, x2, y2, pic_info,sim, dir,intX, intY)
string AiFindPicMemEx(x1, y1, x2, y2, pic_info,sim, dir)
*/

// AiEnableFindPicWindow 设置是否在调用AiFindPicXX系列接口时,是否弹出找图结果的窗口.  方便调试. 默认是关闭的.
func (com *Dmsoft) AiEnableFindPicWindow(enable int) int {
	ret, _ := com.dm.CallMethod(DllM["AiEnableFindPicWindow"], enable)
	defer ret.Clear()
	return int(ret.Val)
}

// AiFindPic 查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.
//
// 这个函数可以查找多个图片,只返回第一个找到的X Y坐标.
//
// 此接口使用Ai模块来实现,比传统的FindPic的效果更好.
func (com *Dmsoft) AiFindPic(x1, y1, x2, y2 int, picName string, sim float32, dir int, intX, intY *int) int {
	ret, _ := com.dm.CallMethod(DllM["AiFindPic"], x1, y1, x2, y2, picName, sim, dir, intX, intY)
	defer ret.Clear()
	return int(ret.Val)
}

// AiFindPicEx 查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.
//
// 这个函数可以查找多个图片,并且返回所有找到的图像的坐标.
//
// 此接口使用Ai模块来实现,比传统的FindPicEx的效果更好.
func (com *Dmsoft) AiFindPicEx(x1, y1, x2, y2 int, picName string, sim float32, dir int) int {
	ret, _ := com.dm.CallMethod(DllM["AiFindPicEx"], x1, y1, x2, y2, picName, sim, dir)
	defer ret.Clear()
	return int(ret.Val)
}

// AiFindPicMem 查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.
//
// 这个函数可以查找多个图片,只返回第一个找到的X Y坐标. 这个函数要求图片是数据地址.
//
// 此接口使用Ai模块来实现,比传统的FindPicMem的效果更好.
func (com *Dmsoft) AiFindPicMem(x1, y1, x2, y2 int, picInfo string, sim float32, dir int, intX, intY *int) int {
	ret, _ := com.dm.CallMethod(DllM["AiFindPicMem"], x1, y1, x2, y2, picInfo, sim, dir, intX, intY)
	defer ret.Clear()
	return int(ret.Val)
}

// AiFindPicMemEx 查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.
//
// 这个函数可以查找多个图片,并且返回所有找到的图像的坐标. 这个函数要求图片是数据地址.
//
// 此接口使用Ai模块来实现,比传统的FindPicMemEx的效果更好.
func (com *Dmsoft) AiFindPicMemEx(x1, y1, x2, y2 int, picInfo string, sim float32, dir int) int {
	ret, _ := com.dm.CallMethod(DllM["AiFindPicMemEx"], x1, y1, x2, y2, picInfo, sim, dir)
	defer ret.Clear()
	return int(ret.Val)
}
