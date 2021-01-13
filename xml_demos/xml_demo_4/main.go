package main

import (
	"os"

	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/MDGSF/utils"
	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()

	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	mpd := doc.CreateElement("MPD")

	mediaPresentationDuration := time.Duration(0)

	for i := 1; i <= 3; i++ {
		iStr := utils.IntToString(i)

		curDoc := etree.NewDocument()

		if err := curDoc.ReadFromFile(iStr + ".mpd"); err != nil {
			panic(err)
		}

		period := curDoc.FindElement("/MPD/Period")
		period.CreateAttr("id", "DID"+iStr)

		duration := period.SelectAttr("duration")
		durationTime, err := ParseDuration(duration.Value)
		if err != nil {
			panic(err)
		}

		mediaPresentationDuration += durationTime

		var prefix string
		if i == 1 {
			prefix = "runtime/video_testcase/"
		} else {
			prefix = "runtime/video_raw/"
		}

		baseurls := period.FindElements("./AdaptationSet/Representation/BaseURL")
		for _, baseurl := range baseurls {
			baseurl.SetText(prefix + baseurl.Text())
		}

		mpd.AddChild(period)
	}

	mpd.CreateAttr("xmlns", "urn:mpeg:dash:schema:mpd:2011")
	mpd.CreateAttr("minBufferTime", "PT1.500S")
	mpd.CreateAttr("type", "static")
	mpd.CreateAttr("mediaPresentationDuration", FormatDuration(mediaPresentationDuration))
	mpd.CreateAttr("maxSegmentDuration", "PT0H0M20.000S")
	mpd.CreateAttr("profiles", "urn:mpeg:dash:profile:full:2011")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}

var (
	// ErrBadFormat is returned when parsing fails
	ErrBadFormat = errors.New("bad format string")

	// ErrNoMonth is raised when a month is in the format string
	ErrNoMonth = errors.New("no months allowed")

	// full = regexp.MustCompile(`P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?`)
	full = regexp.MustCompile(`P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>([0-9]*[.])?[0-9]+)S)?)?`)
	week = regexp.MustCompile(`P((?P<week>\d+)W)`)
)

// adapted from https://github.com/BrianHicks/finch/duration
func ParseDuration(value string) (time.Duration, error) {
	var match []string
	var regex *regexp.Regexp

	if week.MatchString(value) {
		match = week.FindStringSubmatch(value)
		regex = week
	} else if full.MatchString(value) {
		match = full.FindStringSubmatch(value)
		regex = full
	} else {
		return time.Duration(0), ErrBadFormat
	}

	d := time.Duration(0)
	day := time.Hour * 24
	week := day * 7
	year := day * 365

	for i, name := range regex.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		if name == "second" {
			value, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return time.Duration(0), err
			}
			value *= 1000

			d += time.Millisecond * time.Duration(value)

			continue
		}

		value, err := strconv.Atoi(part)
		if err != nil {
			return time.Duration(0), err
		}
		switch name {
		case "year":
			d += year * time.Duration(value)
		case "month":
			if value != 0 {
				return time.Duration(0), ErrNoMonth
			}
		case "week":
			d += week * time.Duration(value)
		case "day":
			d += day * time.Duration(value)
		case "hour":
			d += time.Hour * time.Duration(value)
		case "minute":
			d += time.Minute * time.Duration(value)
		}
	}

	return d, nil
}

func FormatDuration(duration time.Duration) string {
	// we're not doing negative durations
	if duration.Seconds() <= 0 {
		return "PT0S"
	}

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	seconds := duration.Seconds() - float64(hours*3600+minutes*60)

	// we're not doing Y,M,W
	s := "PT"
	if hours > 0 {
		s = fmt.Sprintf("%s%dH", s, hours)
	}
	if minutes > 0 {
		s = fmt.Sprintf("%s%dM", s, minutes)
	}
	if seconds > 0 {
		s = fmt.Sprintf("%s%.3fS", s, seconds)
	}

	return s
}
