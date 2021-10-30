package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
)


func main(){
	a:=app.New()
	w:=a.NewWindow("ALinux")
	fullScreen:=true;
	monitor := w.getMonitorForWindow()
		mode := monitor.GetVideoMode()
		w.fullScreen = fullScreen
		if fullScreen {
			fmt.Println( 0, 0, mode.Width, mode.Height, mode.RefreshRate)
		} else {
		fmt.Println(nil, w.xpos, w.ypos, w.width, w.height, 0)
		}
	
}