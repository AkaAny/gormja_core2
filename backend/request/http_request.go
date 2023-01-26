package request

import (
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"github.com/guonaihong/gout"
	"gormja_core2/jsutil"
	"io"
	"net/http"
	"reflect"
)

func MakeFetch(runtime *goja.Runtime) {
	if err := runtime.GlobalObject().Set("fetch", fetch); err != nil {
		panic(fmt.Errorf("set fetch to global obj with err:%w", err))
	}
}

type JSResponse struct {
	resp *http.Response
}

func (x *JSResponse) ToJSValue(runtime *goja.Runtime) (goja.Value, error) {
	var obj = runtime.NewObject()
	err := obj.Set("json", x.Json)
	if err != nil {
		return nil, fmt.Errorf("attach to vm with err:%w", err)
	}
	return obj, nil
}

func (x *JSResponse) Json(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	defer x.resp.Body.Close()
	promise, resolve, reject := runtime.NewPromise()
	rawData, err := io.ReadAll(x.resp.Body)
	if err != nil {
		reject(fmt.Errorf("read body with err:%w", err))
		return runtime.ToValue(promise)
	}
	var kvMap = make(map[string]interface{})
	if err := json.Unmarshal(rawData, &kvMap); err != nil {
		var goErr = fmt.Errorf("unmarshal as json with err:%w", err)
		reject(goErr)
		return runtime.ToValue(promise)
	}
	resolve(kvMap)
	return runtime.ToValue(promise)
}

func (x *JSResponse) wrapFetchErr(err error, runtime *goja.Runtime) goja.Value {
	var errObj = runtime.NewObject()
	respObj, err := x.ToJSValue(runtime)
	if err != nil {
		panic(err)
	}
	if err := errObj.Set("response", respObj); err != nil {
		panic(err)
	}
	return errObj
}

func fetch(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	promise, onFulfilled, onRejected := runtime.NewPromise()
	var url = call.Argument(0).String()
	var requestInitObj = call.Argument(1).ToObject(runtime)
	var method = requestInitObj.Get("method").String()
	var headerMapObj = requestInitObj.Get("headers").ToObject(runtime)
	var headerMap = http.Header{}
	//http.Header{}
	for _, key := range headerMapObj.Keys() {
		var valueObj = headerMapObj.Get(key).ToObject(runtime)
		if jsutil.IsJSArray(runtime, valueObj) {
			values, err := jsutil.ToGoArray[string](valueObj)
			if err != nil {
				onRejected(fmt.Errorf("key:%s refers to an array, convert to go got err:%w", key, err))
				return runtime.ToValue(promise)
			}
			for _, value := range values {
				headerMap.Add(key, value)
			}
		} else {
			headerMap.Set(key, valueObj.String())
		}
	}
	//http body
	var bodyData = make([]byte, 0)
	var bodyJSValue = requestInitObj.Get("body")
	var bodyObjType = bodyJSValue.ExportType()
	switch bodyObjType {
	case reflect.TypeOf(""): //string body
		var bodyStr = bodyJSValue.String()
		bodyData = []byte(bodyStr)
	default:
		onRejected(fmt.Errorf("unsupported body type:%s", bodyObjType))
		return runtime.ToValue(promise)
	}
	fmt.Println(bodyObjType)
	rawResp, err := gout.New().SetMethod(method).SetURL(url).SetHeader(headerMap).SetBody(bodyData).Response()
	if err != nil {
		onRejected(fmt.Errorf("request with err:%w", err))
	}
	var jsResponse = &JSResponse{resp: rawResp}
	jsResponseObj, err := jsResponse.ToJSValue(runtime)
	if err != nil {
		onRejected(fmt.Errorf(""))
	}
	onFulfilled(jsResponseObj)
	return runtime.ToValue(promise)
}
