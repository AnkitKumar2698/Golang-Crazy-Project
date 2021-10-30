package main

import (
	// "fmt"

	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	// "log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main(){
	a:=app.New()
	w:=a.NewWindow("Weather")
	res,err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q=Karnal&Appid=d0c9519b1dbd29fc02c4333c9faaaaa8")

  	if err!=nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	weather,err:=UnmarshalWelcome(body)
	img:=canvas.NewImageFromFile("weather.png")
	img.FillMode=canvas.ImageFillOriginal
	lable1:=canvas.NewText("Weather Details In Karnal",color.Black)
	lable1.TextStyle=fyne.TextStyle{Bold:true}
    label2:=canvas.NewText(fmt.Sprintf("Country: %s",weather.Sys.Country),color.Black)
	label3:=canvas.NewText(fmt.Sprintf("Wind Speed: %.2f ",weather.Wind.Speed),color.Black)
	label4:=canvas.NewText(fmt.Sprintf("Temperature: %2f k ",weather.Main.Temp),color.Black)
	// label5:=canvas.NewText("Country ",weather.Welcome.timezone)
	w.SetContent(container.NewVBox(
		lable1,
	    img,
		label2,
		label3,
		label4,
		// label5,
	))
	w.ShowAndRun()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}
