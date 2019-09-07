package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"text/template"

	"github.com/go-yaml/yaml"
	"github.com/gosimple/slug"
	"github.com/wcharczuk/go-chart"
)

type Languages []Language

func (ls Languages) Len() int {
	return len(ls)
}

func (ls Languages) Less(i, j int) bool {
	return len(ls[i].Keywords) < len(ls[j].Keywords)
}

func (ls Languages) Swap(i, j int) {
	tmp := ls[i]
	ls[i] = ls[j]
	ls[j] = tmp
}

type AlphaLanguages Languages

func (ls AlphaLanguages) Len() int {
	return len(ls)
}

func (ls AlphaLanguages) Less(i, j int) bool {
	switch {
	case ls[i].Name < ls[j].Name, ls[i].Name > ls[j].Name:
		return ls[i].Name < ls[j].Name
	default:
		return ls[i].Version < ls[j].Version
	}
}

func (ls AlphaLanguages) Swap(i, j int) {
	tmp := ls[i]
	ls[i] = ls[j]
	ls[j] = tmp
}

type Language struct {
	Name     string   `yaml:"name"`
	Version  string   `yaml:"version"`
	Keywords []string `yaml:"keywords"`
	Sources  []string `yaml:"sources"`
}

var funcMap = template.FuncMap{
	"chunk": func(sl []string, ch int) [][]string {
		l := len(sl) / ch
		if len(sl)%ch != 0 {
			l++
		}
		slsl := make([][]string, l)
		for i := 0; i < l; i++ {
			if ch*i+ch >= len(sl) {
				slsl[i] = sl[i*ch:]
			} else {
				slsl[i] = sl[i*ch : (i+1)*ch]
			}
		}
		return slsl
	},
	"alphabetize": func(ls Languages) Languages {
		tmp := make([]Language, len(ls))
		copy(tmp, ls)
		lsaf := AlphaLanguages(tmp)
		sort.Sort(lsaf)

		return Languages(lsaf)
	},
	"slug": slug.Make,
}

const tmplFile = "README.md.tmpl"
const dataFile = "chart.yaml"
const chartFile = "chart.png"
const readmeFile = "README.md"

func main() {
	//read in the YAML file and build the data set
	data := Languages{}
	dataBytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dataBytes, &data)
	sort.Sort(data)

	//render and save the chart
	chartData := make([]chart.Value, len(data))
	for i := range data {
		chartData[i] = chart.Value{Value: float64(len(data[i].Keywords)), Label: fmt.Sprintf("%s (%s)", data[i].Name, data[i].Version), Style: chart.StyleShow()}
	}

	//make ticks
	ticks := []chart.Tick{}
	max := len(data[len(data)-1].Keywords)
	i := 0
	for {
		ticks = append(ticks, chart.Tick{Value: float64(i * 10), Label: strconv.Itoa(i * 10)})
		if i*10 > max {
			break
		}
		i++
	}

	barchart := chart.BarChart{
		Title:      "Programming languages by keyword count",
		TitleStyle: chart.StyleShow(),
		XAxis: chart.Style{
			Show:                true,
			TextRotationDegrees: 90,
		},
		YAxis: chart.YAxis{
			Name:      "Keywords",
			Style:     chart.StyleShow(),
			Ascending: true,
			ValueFormatter: func(v interface{}) string {
				return fmt.Sprintf("%d", int(v.(float64)))
			},
			Ticks: ticks,
		},

		UseBaseValue: true,
		BaseValue:    0,
		Bars:         chartData,
	}

	buf := &bytes.Buffer{}
	err = barchart.Render(chart.PNG, buf)
	ioutil.WriteFile(chartFile, buf.Bytes(), 0644)

	//compile the final from the template
	slug.CustomSub = map[string]string{
		".": "",
	}
	t := template.Must(template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile))
	buf = &bytes.Buffer{}
	t.Execute(buf, data)
	ioutil.WriteFile(readmeFile, buf.Bytes(), 0644)
}
