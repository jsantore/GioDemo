package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"log"
	"os"
)
var sampleData = map[string]string{
	"comp510":"Topics in Programming Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course investigates programming language development from designer’s, user’s and implementer’s point of view. Topics include formal syntax and semantics, language system, extensible languages and control structures. There is also a survey of intralanguage features, covering ALGOL-60, ALGOL-68, Ada, Pascal, LISP, SNOBOL-4 APL, SIMULA-67, CLU, MODULA, and others. Offered periodically.",
	"comp520":"Operating Systems Principles\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course examines design principles such as optimal scheduling; file systems, system integrity and security, as well as the mathematical analysis of selected aspects of operating system design. Topics include queuing theory, disk scheduling, storage management and the working set model. Design and implementation of an operating system nucleus is also studied. Offered periodically.",
	"comp530":"Software Engineering\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include construction of reliable software, software tools, software testing methodologies, structured design, structured programming, software characteristics and quality and formal proofs of program correctness. Chief programmer teams and structure walk-throughs will be employed. Offered periodically.\n",
	"comp545":"Analysis of Algorithms\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course deals with techniques in the analysis of algorithms. Topics to be chosen from among the following: dynamic programming, search and traverse techniques, backtracking, numerical techniques, NP-hard and NP-complete problems, approximation algorithms and other topics in the analysis and design of algorithms. Offered fall semester.\n",
	"comp560":"Artificial Intelligence\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to LISP or another AI programming language. Topics are chosen from pattern recognition, theorem proving, learning, cognitive science and vision. It also presents introduction to the basic techniques of AI such as heuristic search, semantic nets, production systems, frames, planning and other AI topics. Offered periodically.\n",
	"comp570":"Robotics\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis is a project-oriented course in robotics. Topics are chosen from manipulator motion and control, motion planning, legged-motion, vision, touch sensing, grasping, programming languages for robots and automated factory design. Offered periodically."}

type ListElement struct{
	Element widget.Clickable
	Title string
	Desc string
}

type ClassList struct{
	list layout.List
	Items []ListElement
	selected int
}

var listControl ClassList
var appTheme *material.Theme

func setupList(){
	for key, value := range sampleData{
		listControl.Items = append(listControl.Items, ListElement{Title: key, Desc: value})
	}
}

func main(){
	setupList()
	go startApp()
	app.Main()
}

func startApp(){
	defer os.Exit(0) //if we leave this function then  exit with success
	mainWindow := app.NewWindow()
	err := mainEventLoop(mainWindow)
	if err != nil{
		log.Fatal(err)
	}
}

func mainEventLoop(mainWindow *app.Window)(err error){
	appTheme = material.NewTheme(gofont.Collection())

	var operationsQ op.Ops
	for{//for ever loop
		event := <- mainWindow.Events() //read from the events channel, will wait till there is an event if none
		switch eventType := event.(type){
		case system.DestroyEvent: //so the user closed the window
			return eventType.Err
		case system.FrameEvent: //time to draw the window
			graphicsContext := layout.NewContext(&operationsQ, eventType)
			drawGUI(graphicsContext, appTheme)
			eventType.Frame(graphicsContext.Ops)
		}
	}
}

func drawGUI(gContext layout.Context, theme *material.Theme)layout.Dimensions{
	retLayout := layout.Flex{Axis: layout.Vertical}.Layout(gContext, //now we begin building the layout tree toplevel is flex
		layout.Rigid(drawList(gContext, theme)),
		layout.Flexed(1,drawDisplay(gContext, theme)))
	return retLayout

}

func drawList(gContext layout.Context, theme *material.Theme)layout.Widget{
	return func(gtx layout.Context) layout.Dimensions { //goland helps - will autofill function declaration
		return listControl.list.Layout(gtx, len(listControl.Items),selectItem) //the listItem type is a function from context and int to dimensions
	}
}

func selectItem(graphicsContext layout.Context, selectedItem int) layout.Dimensions{
	userSelection := &listControl.Items[selectedItem]
	if userSelection.Element.Clicked() {
		listControl.selected = selectedItem
	}
	var itemHeight int
	//the layout.Stack.Layout function takes a context followed by possibly many StackChild Structs, each of which must be created
	//by using either layout.Explanded, or layout.Stacked. In either case the parameter is a function for context to dimensions
	return layout.Stack{Alignment: layout.W}.Layout(graphicsContext,
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //going with the anonymous function so we can use userSelection
				dimensions := material.Clickable(gtx, &userSelection.Element,
					func(gtx layout.Context) layout.Dimensions { //yes another!! anonymous function
						return layout.UniformInset(unit.Sp(12)).
							Layout(gtx, material.H6(appTheme, userSelection.Title).Layout)
					})
				itemHeight = dimensions.Size.Y
				return dimensions
			}),//thats the end of the first child
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //another one of those 'glorious anonymous functions
				if listControl.selected != selectedItem{
					return layout.Dimensions{} //if not selected - don't do anything special
				}
				paint.ColorOp{Color: appTheme.Color.Primary}.Add(gtx.Ops)//add a paint operation
				highlightWidth:= gtx.Px(unit.Dp(4)) //lets make it 4 device independent pixals
				paint.PaintOp{Rect: f32.Rectangle{ //paint a rectangle using 32 bit floats
					Max: f32.Point{
						X: float32(highlightWidth),
						Y: float32(itemHeight),
				}}}.Add(gtx.Ops)
				return layout.Dimensions{Size: image.Point{X: highlightWidth, Y: itemHeight}}
			},
		),
	)
}

func drawDisplay(gContext layout.Context, theme *material.Theme)layout.Widget{ //layout.Widget is a function from context to dimensions
	return func (ctx layout.Context) layout.Dimensions {
		displayText := material.Body1(theme, listControl.Items[listControl.selected].Desc)
		return layout.Center.Layout(ctx, displayText.Layout)
	}
}