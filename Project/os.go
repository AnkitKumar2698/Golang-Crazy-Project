package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/Knetic/govaluate"

	// "fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	// "fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a:=app.New()
	w:=a.NewWindow("ALinux")
	fullScreen:=true;
	w.Resize(fyne.NewSize(900,700))
    dark:=false
	// startMenu:=false
	w.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

        if keyEvent.Name == fyne.KeyF11 {
			w.SetFullScreen(fullScreen)
			fullScreen=!fullScreen
        }
    })
	img:=canvas.NewImageFromFile("Alinux.jpg")
	img.FillMode = canvas.ImageFillStretch
	bar:=widget.NewToolbar(
		widget.NewToolbarAction(theme.VisibilityIcon(),func(){ 
			dark=!dark
			if dark {
				a.Settings().SetTheme(theme.DarkTheme())
			}else{
				a.Settings().SetTheme(theme.LightTheme())
			}
		 }),
		 widget.NewToolbarAction(theme.ViewFullScreenIcon(),func(){ 
			w.SetFullScreen(fullScreen)
			fullScreen=!fullScreen
		 }),
		 
	)
// iconFile1, _ := os.Open("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\Notepad.png")
// r1 := bufio.NewReader(iconFile1)
// b1, _ := ioutil.ReadAll(r1)
iconNotepad:=canvas.NewImageFromFile("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\Notepad.png")
iconNotepad.FillMode=canvas.ImageFillContain
notepadBtn:= widget.NewButton("", func() {
	notepad(a)
})
notepadbt:=container.NewPadded(iconNotepad,notepadBtn)
iconFile2, _ := os.Open("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\weather.jpg")
r2 := bufio.NewReader(iconFile2)
b2, _ := ioutil.ReadAll(r2)
WeatherBtn:= widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b2), func() {
	weather(a)
})	
iconFile3, _ := os.Open("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\calculator.png")
r3 := bufio.NewReader(iconFile3)
b3, _ := ioutil.ReadAll(r3)
calculatorBtn:= widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b3), func() {
	calculator(a)
})	
iconFile4, _ := os.Open("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\gallery.png")
r4 := bufio.NewReader(iconFile4)
b4, _ := ioutil.ReadAll(r4)
galleryBtn:= widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b4), func() {
	gallery(a)
})	
iconFile5, _ := os.Open("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\start.png")
r5 := bufio.NewReader(iconFile5)
b5, _ := ioutil.ReadAll(r5)
startBtn:= widget.NewButtonWithIcon("",fyne.NewStaticResource("icon", b5), func() {
  
	})
empty:=container.NewVBox()
sidebar:=container.NewGridWithRows(18,
	notepadbt,
	WeatherBtn,
	galleryBtn,
	calculatorBtn,
	startBtn,
)
split:=container.NewHSplit(sidebar,empty)
split.SetOffset(0.12)
	w.SetContent(container.NewVBox(
	container.New(layout.NewMaxLayout(),
     img,
	 container.NewVBox(
		  bar,
		 split,
		 
	   ),
     ),
	),)
	w.ShowAndRun()
}


func gallery(a fyne.App){
	w1:=a.NewWindow("Gallery")
	filepath:="C:\\test"
	files, err := ioutil.ReadDir(filepath)
    if err != nil {
        log.Fatal(err)
    }
	tabs:=container.NewAppTabs()

    for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())	
		if !file.IsDir() {
			expression:=  strings.Split(file.Name(), ".")[1]
		if expression=="png" || expression=="jpg" ||expression=="jpeg" {	
			image := canvas.NewImageFromFile(filepath+"\\"+file.Name())
			tabs.Append(container.NewTabItem(file.Name(),image))
		  } 
		}
    }
	tabs.SetTabLocation(container.TabLocationLeading)
	w1.Resize(fyne.NewSize(800,700))
	w1.SetContent(tabs) 
	w1.Show()
}

func notepad(a fyne.App){
    w3:=a.NewWindow("NotePad By Ankit")
	var count int =0
	 openedfiles:=binding.BindStringList(
		&[]string{},
	)
	input:=widget.NewMultiLineEntry()
   var inputtedtext= make(map[int]string)
	bar:=widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(),func(){ 
		
		 }),
		widget.NewToolbarAction(theme.FolderOpenIcon(),func(){
			openfiledialog:=dialog.NewFileOpen(
			  	func(r fyne.URIReadCloser,_ error){
					  ReadData,_:=ioutil.ReadAll(r)
					 
					  output:=fyne.NewStaticResource("Open"+strconv.Itoa(count),ReadData)
					  val := fmt.Sprintf(output.StaticName)
					  openedfiles.Append(val)
					  inputtedtext[count]=string(output.StaticContent)
					  count++
				  },w3)
				  openfiledialog.SetFilter(
					storage.NewExtensionFileFilter([]string{".txt"}),
				  )	
				  openfiledialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(),func(){
			saveFileDialog:=dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, _ error){
					textdata:=[]byte(input.Text)
					uc.Write(textdata)
				},w3)
				saveFileDialog.SetFileName("New File"+strconv.Itoa(count)+".txt")
				saveFileDialog.Show()
		}), 
		widget.NewToolbarAction(theme.ContentAddIcon(),func(){
			
				val := fmt.Sprintf("File %d",count)
				openedfiles.Append(val)
				inputtedtext[count]=""
				count++
		}),
	)
	listSide := widget.NewListWithData(openedfiles,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
     listSide.OnSelected=func(id widget.ListItemID){
		 if inputtedtext!=nil{
         _,present:=inputtedtext[id]
		  if present {
			input.SetText(inputtedtext[id])
		  }
	}
		
	}

	 listSide.OnUnselected=func(id widget.ListItemID){
		 inputtedtext[id]=input.Text
	 }
	list:=container.New(layout.NewHBoxLayout(),listSide)
	split:=container.NewHSplit(list,input)
	split.SetOffset(0.12)
	side:=container.New(layout.NewBorderLayout(bar,nil,nil,nil),bar,split)
	w3.Resize(fyne.NewSize(600,500))
	w3.SetContent(side) 
	w3.Show()
}

func calculator( a fyne.App){
	w4:=a.NewWindow("Calculator")
	output :=""
	input:=widget.NewLabel(output);
	showHistory:=false;
	history:=""
	historylabel:=widget.NewLabel(history)
	var historyarr [] string 
    historyBtn:=widget.NewButton("History",func(){
		if !showHistory {
         for i:=(len(historyarr)-1);i>=0;i-- {
			history+=historyarr[i]
			history+="\n"
		 }	 
		}else{
			history=""
		}
		showHistory=!showHistory
		historylabel.SetText(history); 
	})
	backBtn:=widget.NewButton("Back",func(){
		if output=="error"{
			output=""
			input.SetText(output)	
		}
		if len(output)>0{
		 output=output[:len(output)-1]
		 input.SetText(output)
		}		
	})
	clearBtn:=widget.NewButton("Clear",func(){
		output=""
		input.SetText(output)
	})
	openBtn:=widget.NewButton("(",func(){
		output=output+"("
		input.SetText(output)
	})
	clostBtn:=widget.NewButton(")",func(){
		output=output+")"
		input.SetText(output)
	})
	devideBtn:=widget.NewButton("/",func(){
		output=output+"/"
		input.SetText(output)	
	})
	nineBtn:=widget.NewButton("9",func(){
		output=output+"9"
		input.SetText(output)
	})
	eightBtn:=widget.NewButton("8",func(){
		output=output+"8"
		input.SetText(output)
	})
	sevenBtn:=widget.NewButton("7",func(){
		output=output+"7"
		input.SetText(output)
	})
	multiplyBtn:=widget.NewButton("*",func(){
		output=output+"*"
		input.SetText(output)
	})
	sixBtn:=widget.NewButton("6",func(){
		output=output+"6"
		input.SetText(output)
	})
	fiveBtn:=widget.NewButton("5",func(){
		output=output+"5"
		input.SetText(output)
	})
	fourBtn:=widget.NewButton("4",func(){
		output=output+"4"
		input.SetText(output)
	})
	minusBtn:=widget.NewButton("-",func(){
		output=output+"-"
		input.SetText(output)
	})
	threeBtn:=widget.NewButton("3",func(){
		output=output+"3"
		input.SetText(output)
	})
	twoBtn:=widget.NewButton("2",func(){
		output=output+"2"
		input.SetText(output)
	})
	oneBtn:=widget.NewButton("1",func(){
		output=output+"1"
		input.SetText(output)
	})
	plusBtn:=widget.NewButton("+",func(){
		output=output+"+"
		input.SetText(output)
	})
	dotBtn:=widget.NewButton(".",func(){
		output=output+"."
		input.SetText(output)
	})
	zeroBtn:=widget.NewButton("0",func(){
		output=output+"0"
		input.SetText(output)
	})
	equalBtn:=widget.NewButton("=",func(){
		
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				
				res:=strconv.FormatFloat(result.(float64),'f',-1,64)
				reshistory:=output+"="+res
				historyarr=append(historyarr,reshistory)
				output=res
			}
		}else{
			output="error"
		}

		input.SetText(output)
	})
	w4.SetContent(container.NewVBox(
	    input,
		historylabel,
		container.NewGridWithColumns(1,
		  container.NewGridWithColumns(2,
			historyBtn,
			backBtn,
		   ),
		),
		container.NewGridWithColumns(4,
		   clearBtn,
		   openBtn,
		   clostBtn,
		   devideBtn,
		   ),
		container.NewGridWithColumns(4,
		 sevenBtn,
		 eightBtn,
		 nineBtn,
		 multiplyBtn,
	   ),
	   container.NewGridWithColumns(4,
	    fourBtn,
		fiveBtn,
		sixBtn,
		minusBtn,
	   ),
	   container.NewGridWithColumns(4,
	    oneBtn,
		twoBtn,
		threeBtn,
		plusBtn,
	   ),
	   container.NewGridWithColumns(2,
	    container.NewGridWithColumns(2,
		  dotBtn,
		  zeroBtn,
		),
		equalBtn,
	   ),

	))
	w4.Show()
}

func weather(a fyne.App){
	w5:=a.NewWindow("Weather")
	res,err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q=Karnal&Appid=d0c9519b1dbd29fc02c4333c9faaaaa8")

  	if err!=nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	weather,err:=UnmarshalWelcome(body)
	img:=canvas.NewImageFromFile("weather.jpg")
	img.FillMode=canvas.ImageFillOriginal
	lable1:=canvas.NewText("Weather Details In Karnal",color.Black)
	lable1.TextStyle=fyne.TextStyle{Bold:true}
    label2:=canvas.NewText(fmt.Sprintf("Country: %s",weather.Sys.Country),color.Black)
	label3:=canvas.NewText(fmt.Sprintf("Wind Speed: %.2f ",weather.Wind.Speed),color.Black)
	label4:=canvas.NewText(fmt.Sprintf("Temperature: %2f k ",weather.Main.Temp),color.Black)
	// label5:=canvas.NewText("Country ",weather.Welcome.timezone)
	w5.SetContent(container.NewVBox(
		lable1,
	    img,
		label2,
		label3,
		label4,
	))
	w5.Show()
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
