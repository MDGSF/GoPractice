package main

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func main() {
	output_1()

	fmt.Println()

	output_mpd()

	fmt.Println()

	parse_bookstore()
}

func output_mpd() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mpd := doc.CreateElement("MPD")
	mpd.CreateAttr("xmlns", "urn:mpeg:dash:schema:mpd:2011")
	mpd.CreateAttr("minBufferTime", "PT1.500S")
	mpd.CreateAttr("type", "static")
	mpd.CreateAttr("mediaPresentationDuration", "PT0H0M54.950S")
	mpd.CreateAttr("maxSegmentDuration", "PT0H0M20.000S")
	mpd.CreateAttr("profiles", "urn:mpeg:dash:profile:full:2011")

	period1 := mpd.CreateElement("Period")
	period1.CreateAttr("duration", "PT0H0M54.950S")

	programinformation := period1.CreateElement("ProgramInformation")
	programinformation.CreateAttr("moreInformationURL", "http://gpac.io")

	title := programinformation.CreateElement("Title")
	title.CreateText("1.mpd generated by GPAC")

	period2 := mpd.CreateElement("Period")
	period2.CreateAttr("duration", "PT0H0M54.950S")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}

func output_1() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	people := doc.CreateElement("People")
	people.CreateComment("These are all known people")

	jon := people.CreateElement("Person")
	jon.CreateAttr("name", "Jon")

	sally := people.CreateElement("Person")
	sally.CreateAttr("name", "Sally")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}

func parse_bookstore() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("bookstore.xml"); err != nil {
		panic(err)
	}

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}