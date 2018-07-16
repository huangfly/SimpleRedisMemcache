package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFuncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Req
		rsp := make(map[string]interface{})
		defer func() {
			data, _ := json.MarshalIndent(rsp, "", " ")
			w.Write(data)
		}()
		//获取http的body信息
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Read request body error"
			log.Println(err)
			return //TODO
		}
		r.Body.Close()
		json.Unmarshal(body, &req)
		value, err := handler.GetValue(req.Key)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Get value from memcache or database error"
			return
		}
		rsp["retcode"] = 1
		rsp["contant"] = value
		log.Printf("GetFuncHandler: key is %s, value is %s \n", req.Key, value)
	}

}

func StoreFuncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Req
		rsp := make(map[string]interface{})
		defer func() {
			data, _ := json.MarshalIndent(rsp, "", " ")
			w.Write(data)
		}()
		//获取http的body信息
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Read request body error"
			log.Println(err)
			return //TODO
		}
		r.Body.Close()
		json.Unmarshal(body, &req)
		err = handler.StoreValue(req.Key, req.Value)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Stroe value to memcache or database error"
			return
		}
		rsp["retcode"] = 1
		rsp["contant"] = "Success"
		log.Printf("StoreFuncHandler : key is %s, value is %s \n", req.Key, req.Value)
	}
}

func DelFuncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Req
		rsp := make(map[string]interface{})
		defer func() {
			data, _ := json.MarshalIndent(rsp, "", " ")
			w.Write(data)
		}()
		//获取http的body信息
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Read request body error"
			log.Println(err)
			return //TODO
		}
		r.Body.Close()
		json.Unmarshal(body, &req)
		err = handler.Delete(req.Key)
		if err != nil {
			rsp["retcode"] = 2
			rsp["contant"] = "Del value from memcache or database error"
			return
		}
		rsp["retcode"] = 1
		rsp["contant"] = "Success"
		log.Printf("DelFuncHandler : key is %s, value is %s \n", req.Key)
	}
}

//打印logo
func printSoftInfo() {
	log.Println("..............CopyRight @ By XiaoFeiFei..................")
	log.Println("		    *     *      * * * *      * * * * ")
	log.Println("		   *     *      *     *      *       ")
	log.Println("		  * * * *      * * * *      * * * * ")
	log.Println("		 *     *      *            *       ")
	log.Println("		*     *      *            *       ")
	log.Println("..................Server Is Start........................")
}
