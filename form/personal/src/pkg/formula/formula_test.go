package formula

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gookit/color"
)

func TestFormula_Run(t *testing.T) {
	type fields struct {
		Name, BirthDate, Country, State string
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
	}{
		{
			name: "Run",
			fields: fields{
				Name:      "Dennis",
				BirthDate: "9 de setembro de 1941",
				Country:   "EUA",
				State:     "Nova York",
			},
			wantWriter: func() string {
				return fmt.Sprintf("Hello World!\n") +
					color.FgGreen.Render(fmt.Sprintf("My name is Dennis.\n")) +
					color.FgYellow.Render(fmt.Sprintf("My birthday is 9 de setembro de 1941.\n")) +
					color.FgBlue.Render(fmt.Sprintf("My country is EUA.\n")) +
					color.FgGreen.Render(fmt.Sprintf("My state is Nova York.\n")) +
					""
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formula{
				Name:      tt.fields.Name,
				BirthDate: tt.fields.BirthDate,
				Country:   tt.fields.Country,
				State:     tt.fields.State,
			}
			writer := &bytes.Buffer{}
			f.Run(writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Run() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
