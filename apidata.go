package main

// Variables used for command line parameters
type config struct {
	Token     string
	replySelf string
	DeepAIkey string
}

type activiteit struct {
	Activity     string
	Participants int
	Price        float64
}

type catfact struct {
	Fact string
}

type doggyphoto struct {
	Message string
}

type fortunecookie struct {
	Fortune string
}

type text2img struct {
	Output_url string
}
