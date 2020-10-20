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
	"comp502": "Research\n(3 credits)\nPrerequisite: Consent of the department; formal application required\nOriginal research is undertaken by the graduate student in their field. This course culminates in a capstone project. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Offered fall and spring semesters.",
	"comp503": "Directed Study\n(1-3 credits)\nPrerequisite: Consent of the department; formal application required\nDirected study is designed for the graduate student who desires to study selected topics in a specific field. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Repeatable: may earn a maximum of six credits. Offered fall and spring semesters.",
	"comp510": "Topics in Programming Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course investigates programming language development from designer’s, user’s and implementer’s point of view. Topics include formal syntax and semantics, language system, extensible languages and control structures. There is also a survey of intralanguage features, covering ALGOL-60, ALGOL-68, Ada, Pascal, LISP, SNOBOL-4 APL, SIMULA-67, CLU, MODULA, and others. Offered periodically.",
	"comp520": "Operating Systems Principles\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course examines design principles such as optimal scheduling; file systems, system integrity and security, as well as the mathematical analysis of selected aspects of operating system design. Topics include queuing theory, disk scheduling, storage management and the working set model. Design and implementation of an operating system nucleus is also studied. Offered periodically.",
	"comp525": "Design and Construction of Compilers\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include lexical and syntactic analysis; code generation; error detection and correction; optimization techniques; models of code generators; and incremental and interactive compiling. Students will design and implement a compiler. Offered periodically.",
	"comp530": "Software Engineering\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include construction of reliable software, software tools, software testing methodologies, structured design, structured programming, software characteristics and quality and formal proofs of program correctness. Chief programmer teams and structure walk-throughs will be employed. Offered periodically.\n",
	"comp540": "Automata, Computability and Formal Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include finite automata and regular languages, context- free languages, Turing machines and their variants, partial recursive functions and grammars, Church’s thesis, undecidable problems, complexity of algorithms and completeness. Offered periodically.",
	"comp545": "Analysis of Algorithms\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course deals with techniques in the analysis of algorithms. Topics to be chosen from among the following: dynamic programming, search and traverse techniques, backtracking, numerical techniques, NP-hard and NP-complete problems, approximation algorithms and other topics in the analysis and design of algorithms. Offered fall semester.\n",
	"comp560": "Artificial Intelligence\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to LISP or another AI programming language. Topics are chosen from pattern recognition, theorem proving, learning, cognitive science and vision. It also presents introduction to the basic techniques of AI such as heuristic search, semantic nets, production systems, frames, planning and other AI topics. Offered periodically.\n",
	"comp570": "Robotics\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis is a project-oriented course in robotics. Topics are chosen from manipulator motion and control, motion planning, legged-motion, vision, touch sensing, grasping, programming languages for robots and automated factory design. Offered periodically.",
	"comp580": "Database Systems\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include relational, hierarchical and network data models; design theory for relational databases and query optimization; classification of data models, data languages; concurrency, integrity, privacy; modeling and measurement of access strategies; and dedicated processors, information retrieval and real time applications. Offered periodically.",
	"comp590": "Computer Architecture\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to the internal structure of digital computers including design of gates, flip-fops, registers and memories to perform operations on numerical and other data represented in binary form; computer system analysis and design; organizational dependence on computations to be performed; and theoretical aspects of parallel and pipeline computation. Offered periodically.",
	"comp594": "Computer Networks\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course provides an introduction to fundamental concepts in computer networks, including their design and implementation. Topics include network architectures and protocols, placing emphasis on protocol used in the Internet; routing; data link layer issues; multimedia networking; network security; and network management. Offered periodically.\n",
	"comp596": "Topics in Computer Science\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics are chosen from program verification, formal semantics, formal language theory, concurrent programming, complexity or algorithms, programming language theory, graphics and other computer science topics. Repeatable for different topics. Offered as topics arise.",
	"comp598": " Computer Science Graduate Internship\n(3 credits)\nPrerequisite: Matriculation in the computer science master’s program; at least six credits of graduate-level course work in computer science (COMP); formal application required\nAn internship provides an opportunity to apply what has been learned in the classroom and allows the student to further professional skills. Faculty supervision allows for reflection on the internship experience and connects the applied portion of the academic study to other courses. Repeatable; may earn a maximum of six credits, however, only three credits can be used toward the degree. Graded on (P) Pass/(N) No Pass basis. Offered fall and spring semesters.\n",
}

var prereqs = map[string][]string{
	"comp502": []string{"comp340", "comp442", "comp490", "comp545"},
	"comp503": []string{"comp442", "comp490"},
	"comp510": {"comp340", "comp350", "comp435", "comp430", "comp490", "comp250", "comp151", "comp152", "comp442", "comp399", "I'm kinda making stuff up",
		"math130", "comp999", "filling up the menu", "HELP there are too many prereqs", "I don't want to show too much of the song stuff",
		"but I want you all to have a baseline", "even if you are grad student"},
	"comp520": {"comp350", "comp250"},
	"comp525": {"comp340", "comp350", "comp435", "comp490", "math130", "math180"},
	"comp530": {"comp250", "comp442", "comp435", "comp490", "comp"},
	"comp545": {"Math161", "Comp435"},
	"comp560": {"math130", "Comp470", "math200", "comp490"},
	"comp570": {"math130", "math120", "comp460", "comp490"},
	"comp580": {"math130", "compXXX"},
	"comp594": {"comp430"},
	"comp596": {"comp490"},
	"comp598": {"comp530"},
}

type prereqElement struct {
	Element    widget.Clickable
	classTitle string
}

type ListElement struct {
	Element widget.Clickable
	Title   string
	Desc    string
}

type prereqList struct {
	list        layout.List
	Items       []prereqElement
	selectedNum int
}

type ClassList struct {
	list     layout.List
	Items    []ListElement
	selected int
}

var listControl ClassList
var secondList prereqList
var appTheme *material.Theme

//populate a static list at the beginning of the method
//also adjuests the orientation of a list control as a side effect
func setupList() {
	for key, value := range sampleData {
		listControl.Items = append(listControl.Items, ListElement{Title: key, Desc: value})
	}
	listControl.list.Axis = layout.Vertical
}

func populateSecondList(className string) {
	secondList.Items = []prereqElement{}
	for _, prereq := range prereqs[className] {
		secondList.Items = append(secondList.Items, prereqElement{classTitle: prereq})
	}
	secondList.list.Axis = layout.Vertical
}

func main() {
	setupList()
	go startApp()
	app.Main()
}

//
func startApp() {
	defer os.Exit(0) //if we leave this function then  exit with success
	mainWindow := app.NewWindow()
	err := mainEventLoop(mainWindow)
	if err != nil {
		log.Fatal(err)
	}
}

func mainEventLoop(mainWindow *app.Window) (err error) {
	appTheme = material.NewTheme(gofont.Collection())

	var operationsQ op.Ops
	for { //for ever loop
		event := <-mainWindow.Events() //read from the events channel, will wait till there is an event if none
		switch eventType := event.(type) {
		case system.DestroyEvent: //so the user closed the window
			return eventType.Err
		case system.FrameEvent: //time to draw the window
			graphicsContext := layout.NewContext(&operationsQ, eventType)
			drawGUI(graphicsContext, appTheme)
			eventType.Frame(graphicsContext.Ops)
		}
	}
}

func drawGUI(gContext layout.Context, theme *material.Theme) layout.Dimensions {
	retLayout := layout.Flex{Axis: layout.Horizontal}.Layout(gContext, //now we begin building the layout tree toplevel is flex
		layout.Rigid(drawList(gContext, theme)),
		layout.Flexed(1, drawSecondList(gContext, theme)))
	return retLayout

}

func drawList(gContext layout.Context, theme *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { //goland helps - will autofill function declaration
		return listControl.list.Layout(gtx, len(listControl.Items), selectItem) //the listItem type is a function from context and int to dimensions
	}
}

func drawSecondList(gContext layout.Context, theme *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return secondList.list.Layout(gtx, len(secondList.Items), subselectItem)
	}
}

func subselectItem(graphicsContext layout.Context, selectedItem int) layout.Dimensions {
	userSelection := &secondList.Items[selectedItem]
	if userSelection.Element.Clicked() {
		secondList.selectedNum = selectedItem
	}
	var itemHeight int
	return layout.Stack{Alignment: layout.NW}.Layout(graphicsContext,
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //going with the anonymous function so we can use userSelection
				dimensions := material.Clickable(gtx, &userSelection.Element,
					func(gtx layout.Context) layout.Dimensions { //yes another!! anonymous function
						return layout.UniformInset(unit.Sp(12)).
							Layout(gtx, material.H6(appTheme, userSelection.classTitle).Layout)
					})
				itemHeight = dimensions.Size.Y
				return dimensions
			}), //thats the end of the first child
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //another one of those 'glorious anonymous functions
				if secondList.selectedNum != selectedItem {
					return layout.Dimensions{} //if not selected - don't do anything special
				}
				paint.ColorOp{Color: appTheme.Color.Hint}.Add(gtx.Ops) //add a paint operation
				highlightWidth := gtx.Px(unit.Dp(4))                   //lets make it 4 device independent pixals
				paint.PaintOp{Rect: f32.Rectangle{                     //paint a rectangle using 32 bit floats
					Max: f32.Point{
						X: float32(highlightWidth),
						Y: float32(itemHeight),
					}}}.Add(gtx.Ops)
				return layout.Dimensions{Size: image.Point{X: highlightWidth, Y: itemHeight}}
			},
		),
	)
}

//select item is called whenever a clickable in the list is pushed
//it lets you respond to the event which in gio is both the actions to take
//and then we have to re layout the UI as well which is why we need both a graphics
//context and the selected itme number. returns a function (layout.Dimensions)
//
//that was just a dramatic pause up there to demo for the students.
func selectItem(graphicsContext layout.Context, selectedItem int) layout.Dimensions {
	userSelection := &listControl.Items[selectedItem]
	if userSelection.Element.Clicked() {
		listControl.selected = selectedItem
		populateSecondList(listControl.Items[selectedItem].Title)
	}
	var itemHeight int
	//the layout.Stack.Layout function takes a context followed by possibly many StackChild Structs, each of which must be created
	//by using either layout.Explanded, or layout.Stacked. In either case the parameter is a function for context to dimensions
	return layout.Stack{Alignment: layout.E}.Layout(graphicsContext,
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //going with the anonymous function so we can use userSelection
				dimensions := material.Clickable(gtx, &userSelection.Element,
					func(gtx layout.Context) layout.Dimensions { //yes another!! anonymous function
						return layout.UniformInset(unit.Sp(12)).
							Layout(gtx, material.H6(appTheme, userSelection.Title).Layout)
					})
				itemHeight = dimensions.Size.Y
				return dimensions
			}), //thats the end of the first child
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions { //another one of those 'glorious anonymous functions
				if listControl.selected != selectedItem {
					return layout.Dimensions{} //if not selected - don't do anything special
				}
				paint.ColorOp{Color: appTheme.Color.Primary}.Add(gtx.Ops) //add a paint operation
				highlightWidth := gtx.Px(unit.Dp(4))                      //lets make it 4 device independent pixals
				paint.PaintOp{Rect: f32.Rectangle{                        //paint a rectangle using 32 bit floats
					Max: f32.Point{
						X: float32(highlightWidth),
						Y: float32(itemHeight),
					}}}.Add(gtx.Ops)
				return layout.Dimensions{Size: image.Point{X: highlightWidth, Y: itemHeight}}
			},
		),
	)
}

func drawDisplay(gContext layout.Context, theme *material.Theme) layout.Widget { //layout.Widget is a function from context to dimensions
	return func(ctx layout.Context) layout.Dimensions {
		displayText := material.Body1(theme, listControl.Items[listControl.selected].Desc)
		return layout.Center.Layout(ctx, displayText.Layout)
	}
}
