package lang

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
)

func getParameterTypeArr(method *rtc.Method) *rtc.Obj {
	paramTypes := method.ParameterTypes()
	paramCount := len(paramTypes)

	classClass := rtc.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.Refs()
		for i, paramType := range paramTypes {
			classObjs[i] = paramType.JClass()
		}
	}

	return classArr
}

func getReturnType(method *rtc.Method) *rtc.Obj {
	goReturnType := method.ReturnType()
	return goReturnType.JClass()
}

func getExceptionTypeArr(method *rtc.Method) *rtc.Obj {
	exTypes := method.ExceptionTypes()
	exCount := len(exTypes)

	classClass := rtc.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.Refs()
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass()
		}
	}

	return classArr
}

func getParameterAnnotationDyteArr(method *rtc.Method) *rtc.Obj {
	if data := method.ParameterAnnotationData(); data != nil {
		bytes := util.CastUint8sToInt8s(data)
		return rtc.NewByteArray(bytes)
	}
	return nil
}
func getAnnotationDefaultData(method *rtc.Method) *rtc.Obj {
	if data := method.AnnotationDefaultData(); data != nil {
		bytes := util.CastUint8sToInt8s(data)
		return rtc.NewByteArray(bytes)
	}
	return nil
}

func getAnnotationByteArr(member *rtc.ClassMember) *rtc.Obj {
	if data := member.AnnotationData(); data != nil {
		bytes := util.CastUint8sToInt8s(data)
		return rtc.NewByteArray(bytes)
	}
	return nil
}

func getSignature(member *rtc.ClassMember) *rtc.Obj {
	if signature := member.Signature(); signature != "" {
		return rtda.NewJString(signature)
	}
	return nil
}
