package samples_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	mapping "github.com/beerproto/parser/beerJSON"
	"github.com/go-test/deep"
)

func TestSchemas_Generate(t *testing.T) {

	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name: "Boil Whirlpool Chill",
			json: "boil_whirlpool_chill.json",
		},
		{
			name: "BrettDosedKegsSaison",
			json: "BrettDosedKegsSaison.json",
		},
		{
			name: "CheaterHops",
			json: "CheaterHops.json",
		},
		{
			name: "CorianderSpice",
			json: "CorianderSpice.json",
		},
		{
			name: "CrystalMaltSpecialtyGrain",
			json: "CrystalMaltSpecialtyGrain.json",
		},
		{
			name: "EquipmentSet",
			json: "EquipmentSet.json",
		},
		{
			name: "FermentableRecord",
			json: "FermentableRecord.json",
		},
		{
			name: "HoppedExtract",
			json: "HoppedExtract.json",
		},
		{
			name: "HopRecordWithAllFields",
			json: "HopRecordWithAllFields.json",
		},
		{
			name: "HopWithRequiredFieldsOnly",
			json: "HopWithRequiredFieldsOnly.json",
		},
		{
			name: "IrishMoss",
			json: "IrishMoss.json",
		},
		{
			name: "MashSingleStepInfusion",
			json: "MashSingleStepInfusion.json",
		},
		{
			name: "MashTwoStepTemperature",
			json: "MashTwoStepTemperature.json",
		},
		{
			name: "MedievalAle",
			json: "MedievalAle.json",
		},
		{
			name: "SampleWaterProfile",
			json: "SampleWaterProfile.json",
		},
		{
			name: "StyleBohemianPilsner",
			json: "StyleBohemianPilsner.json",
		},
		{
			name: "StyleDryIrishStoutWithAllFields",
			json: "StyleDryIrishStoutWithAllFields.json",
		},
		{
			name: "YeastWithMorePopularFields",
			json: "YeastWithMorePopularFields.json",
		},
		{
			name: "YeastWithRequiredFieldsOnly",
			json: "YeastWithRequiredFieldsOnly.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			doc, err := mapping.Load(tt.json)
			if err != nil {
				t.Error(err)
			}

			recipe := mapping.MapToProto(doc.Beer)

			j, err := mapping.MapToJSON(recipe)
			if err != nil {
				t.Error(err)
			}

			bytes, err := json.Marshal(j)
			if err != nil {
				t.Error(err)
			}

			trimedJSON, err := TrimJson(bytes)
			if err != nil {
				t.Error(err)
			}

			innerData, err := json.Marshal(doc.Beer)
			if err != nil {
				t.Error(err)
			}

			expectedJSON, err := TrimJson(innerData)
			if err != nil {
				t.Error(err)
			}

			diff, err := ShouldEqualJSONObject(expectedJSON, trimedJSON)
			if err != nil {
				if len(diff) > 0 {
					s := strings.Join(diff, "\n")
					t.Errorf("test \n test \n %s", s)
				}
				t.Error(err)
			}
		})
	}
}

func ShouldEqualJSONObject(data1, data2 []byte) ([]string, error) {
	x := make(map[string]interface{})
	err := json.Unmarshal(data1, &x)
	if err != nil {
		return []string{}, fmt.Errorf("unmarshal of data1 failed: %w", err)
	}
	y := make(map[string]interface{})
	err = json.Unmarshal(data2, &y)
	if err != nil {
		return []string{}, fmt.Errorf("unmarshal of data2 failed: %w", err)
	}

	if diff := deep.Equal(x, y); diff != nil {
		return diff, fmt.Errorf("object not equal")
	}

	return nil, nil
}

func TrimJson(data []byte) ([]byte, error) {
	x := make(map[string]interface{})
	err := json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}

	result := trimJson(x)

	results, err := json.Marshal(result)
	return results, err
}

func trimJson(x map[string]interface{}) map[string]interface{} {
	for key, value := range x {
		switch v := value.(type) {
		case string:
			if v == "" {
				delete(x, key)
			}
		case float64:
			if v == 0 {
				delete(x, key)
			}
		case bool:
			if v == false {
				delete(x, key)
			}
		case []interface{}:
			for p, iterator := range v {
				if m, ok := iterator.(map[string]interface{}); ok {
					m2 := trimJson(m)
					if len(m) > 0 {
						v[p] = m2
					} else {
						v = append(v[:p], v[p+1:]...)
					}
				}
			}
		case map[string]interface{}:
			m := trimJson(v)
			if len(m) > 0 {
				x[key] = m
			} else {
				delete(x, key)
			}
		}
	}
	return x
}
