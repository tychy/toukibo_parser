package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/tychy/toukibo_parser/toukibo"
	"gopkg.in/yaml.v3"
)

type TestData struct {
	HoujinKaku                string   `yaml:"HoujinKaku"`
	HoujinName                string   `yaml:"HoujinName"`
	HoujinAddress             string   `yaml:"HoujinAddress"`
	HoujinExecutiveNames      []string `yaml:"HoujinExecutiveNames"`
	HoujinRepresentativeNames []string `yaml:"HoujinRepresentativeNames"`
	HoujinDissolvedAt         string   `yaml:"HoujinDissolvedAt"`
	HoujinCapital             string   `yaml:"HoujinCapital"`
}

func TestToukiboParser(t *testing.T) {
	testCount := 190

	for i := 1; i <= testCount; i++ {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			pdfFileName := fmt.Sprintf("testdata/pdf/sample%d.pdf", i)
			yamlFileName := fmt.Sprintf("testdata/yaml/sample%d.yaml", i)
			content, err := readPdf(pdfFileName)
			if err != nil {
				t.Fatal(err)
			}
			h, err := toukibo.Parse(content)
			if err != nil {
				t.Fatal(err)
			}
			yamlContent, err := os.ReadFile(yamlFileName)
			if err != nil {
				t.Fatal(err)
			}
			td := TestData{}

			err = yaml.Unmarshal([]byte(yamlContent), &td)
			if err != nil {
				t.Fatal(err)
			}

			// check
			execNames, err := h.ListHoujinExecutives()
			if err != nil {
				t.Fatal(err)
			}

			if len(execNames) != len(td.HoujinExecutiveNames) {
				t.Fatalf("executive name count is not match,\nwant : %d,\ngot  : %d", len(td.HoujinExecutiveNames), len(execNames))
			}
			for i, v := range execNames {
				if v != td.HoujinExecutiveNames[i] {
					t.Fatalf("executive name is not match,\nwant : %s,\ngot  : %s", td.HoujinExecutiveNames[i], v)
				}
			}

			repNames, err := h.GetHoujinRepresentativeNames()
			if err != nil {
				t.Fatal(err)
			}
			if len(repNames) != len(td.HoujinRepresentativeNames) {
				t.Fatalf("representative name count is not match,\nwant : %d,\ngot  : %d", len(td.HoujinRepresentativeNames), len(repNames))
			}
			for i, v := range repNames {
				if v != td.HoujinRepresentativeNames[i] {
					t.Fatalf("representative name is not match,\nwant : %s,\ngot  : %s", td.HoujinRepresentativeNames[i], v)
				}
			}

			if h.GetHoujinKaku() != td.HoujinKaku {
				t.Fatalf("kaku is not match,\nwant : %s,\ngot  : %s,", td.HoujinKaku, h.GetHoujinKaku())
			}

			if h.GetHoujinName() != td.HoujinName {
				t.Fatalf("name is not match,\nwant : %s,\ngot  : %s,", td.HoujinName, h.GetHoujinName())
			}

			if h.GetHoujinAddress() != td.HoujinAddress {
				t.Fatalf("address is not match,\nwant : %s,\ngot  : %s,", td.HoujinAddress, h.GetHoujinAddress())
			}

			if h.GetHoujinDissolvedAt() != td.HoujinDissolvedAt {
				t.Fatalf("dissolved_at is not match,\nwant : %s,\ngot  : %s,", td.HoujinDissolvedAt, h.GetHoujinDissolvedAt())
			}

			if fmt.Sprint(h.GetHoujinCapital()) != td.HoujinCapital {
				t.Fatalf("capital is not match,\nwant : %s,\ngot  : %d,", td.HoujinCapital, h.GetHoujinCapital())
			}

		})
	}
}
