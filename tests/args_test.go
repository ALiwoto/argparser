package args_test

import (
	"log"
	"testing"

	"github.com/ALiwoto/argparser"
)

const (
	cmd1 = `/command1 hello`
	cmd2 = `/command2 --action ban --word:hello`
	cmd3 = `/command3 --action warn --trigger:how are you?`
	cmd4 = `/command4 --action gban --trigger:I hate you --reason = 
"trolling, spamming, using unsuitable words for the group
and raiding."`
)

func TestArgs1(t *testing.T) {
	arg, err := argparser.ParseArg(cmd1)
	if err != nil {
		t.Error(err)
		return
	}
	if arg == nil {
		t.Error("arg is nil")
		return
	}
	if !arg.CompareCommand("/command1") {
		t.Error("command is not /command1")
		return
	}
	if arg.GetLength() != 0 {
		t.Error("arg length is not zero")
		return
	}
	if !arg.HasRawData() {
		t.Error("arg has no raw data")
		return
	}
	log.Println(arg.GetAsStringOrRaw())
}

func TestArgs2(t *testing.T) {
	arg, err := argparser.ParseArg(cmd2)
	if err != nil {
		t.Error(err)
		return
	}
	if arg == nil {
		t.Error("arg is nil")
		return
	}
	if !arg.CompareCommand("/command2") {
		t.Error("command is not /command2")
		return
	}
	if arg.GetLength() != 2 {
		t.Error("arg length is not 2")
		return
	}
	if arg.HasRawData() {
		t.Error("arg has raw data")
		return
	}
	act := arg.GetAsStringOrRaw("action")
	if act != "ban" {
		t.Error("action is not ban")
		return
	}
	if arg.GetAsString("keyword", "word", "trigger") != "hello" {
		t.Error("keyword is not hello")
		return
	}
	log.Println("we will", arg.GetAsString("action"),
		"users if they use \""+
			arg.GetAsString("keyword", "word", "trigger")+
			"\" in the group")
}

func TestArgs3(t *testing.T) {
	arg, err := argparser.ParseArg(cmd3)
	if err != nil {
		t.Error(err)
		return
	}
	if arg == nil {
		t.Error("arg is nil")
		return
	}
	if !arg.CompareCommand("/command3") {
		t.Error("command is not /command3")
		return
	}
	if arg.GetLength() != 2 {
		t.Error("arg length is not 2")
		return
	}
	if arg.HasRawData() {
		t.Error("arg has raw data")
		return
	}
	act := arg.GetAsStringOrRaw("action")
	if act != "warn" {
		t.Error("action is not warn")
		return
	}
	if arg.GetAsString("keyword", "word", "trigger") != "how are you?" {
		t.Error("keyword is not \"how are you?\"")
		return
	}
	log.Println("we will", arg.GetAsString("action"),
		"users if they use \""+
			arg.GetAsString("keyword", "word", "trigger")+
			"\" in the group")
}

func TestArgs4(t *testing.T) {
	arg, err := argparser.ParseArg(cmd4)
	if err != nil {
		t.Error(err)
		return
	}
	if arg == nil {
		t.Error("arg is nil")
		return
	}
	if !arg.CompareCommand("/command4") {
		t.Error("command is not /command4")
		return
	}
	if arg.GetLength() != 3 {
		t.Error("arg length is not 3")
		return
	}
	if arg.HasRawData() {
		t.Error("arg has raw data")
		return
	}
	act := arg.GetAsStringOrRaw("action")
	if act != "gban" {
		t.Error("action is not gban")
		return
	}
	if arg.GetAsString("keyword", "word", "trigger") != "I hate you" {
		t.Error("keyword is not \"I hate you\"")
		return
	}
	reason := arg.GetAsString("reason")
	log.Println("we will", act,
		"users if they use \""+
			arg.GetAsString("keyword", "word", "trigger")+
			"\" in the group.", "the reason will be: \""+reason+"\"")
}
