package main

import (
	"encoding/json"
	"fmt"
	"km/src/km"
	"km/src/windows"
	"km/src/windows/keyboard"
	"km/src/windows/mouse"
	"net/http"
	"time"
)

func main() {

	//cmd := exec.Command("taskkill.exe", "/f", "/im", "km.exe")
	//cmd.Start()

	port := "7788"
	km.KillProcess(port)

	go keyboardHook()

	go mouseHook()

	files := http.FileServer(http.Dir("./webroot"))
	http.Handle("/", files)
	http.HandleFunc("/km", handlerKM)

	addr := fmt.Sprintf("localhost:%s", port)
	fmt.Printf("http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func keyboardHook() {

	keyboardChan := make(chan windows.KeyboardEvent, 100)
	keyboard.Hook(keyboardChan)

	defer keyboard.UnHook()

	// TODO get default value
	isCapsLock := false
	isShift := false

	for {
		k := <-keyboardChan
		// fmt.Printf("%v\n", k)
		if k.Message == 256 || k.Message == 260 { // key down
			if k.VKCode == 20 {
				isCapsLock = !isCapsLock
			} else if k.VKCode == 160 || k.VKCode == 161 {
				isShift = !isShift
			}
			code := k.VKCode
			if code == 13 && k.Flags == 1 {
				code = 108 // numpad enter
			}
			key := km.GetKey(k.VKCode, isCapsLock, isShift)
			km.LogToFile(km.KEYBOARD, key)
		}
		if k.Message == 257 || k.Message == 261 { // key up
			if k.VKCode == 20 {
				isCapsLock = !isCapsLock
			} else if k.VKCode == 160 || k.VKCode == 161 {
				isShift = !isShift
			}
		}
	}
}

func mouseHook() {
	mouseChan := make(chan windows.MouseEvent, 100)
	mouse.Hook(mouseChan)

	defer mouse.UnHook()
	for {
		k := <-mouseChan
		// fmt.Printf("%v\n", k)
		if k.Message == 513 { // mouse left down
			km.LogToFile(km.MOUSE, "1")
		}
		if k.Message == 516 { // mouse right down
			km.LogToFile(km.MOUSE, "2")
		}
		if k.Message == 519 { // mouse mind down
			km.LogToFile(km.MOUSE, "3")
		}
	}
}

func handlerKM(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	lastTime := query.Get("lastTime")
	day := query.Get("day")
	timeStartStr := query.Get("timeStart")
	timeEndStr := query.Get("timeEnd")

	var startTime time.Time
	var endTime time.Time
	if "" != lastTime {
		s, _ := time.ParseDuration(fmt.Sprintf("-%ss", lastTime))
		startTime = time.Now().Add(s)
		endTime = km.ParseTime(fmt.Sprintf("%s 23:59:59.999", day))
	} else {
		startTime = km.ParseTime(fmt.Sprintf("%s %s:00.000", day, timeStartStr))
		endTime = km.ParseTime(fmt.Sprintf("%s %s:00.000", day, timeEndStr))
	}

	result := map[string]interface{}{}
	result["keyboard"] = km.LogStat(km.KEYBOARD, startTime, endTime)
	result["mouse"] = km.LogStat(km.MOUSE, startTime, endTime)
	jsonBytes, _ := json.Marshal(result)
	_, _ = res.Write(jsonBytes)
}
