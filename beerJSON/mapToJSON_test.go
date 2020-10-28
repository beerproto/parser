package beerJSON

import (
	"github.com/beerproto/beerjson.go"
	beerproto "github.com/beerproto/beerproto_go"
	"github.com/go-test/deep"
	"testing"
)

func TestToJSONRecipeType(t *testing.T) {
	type args struct {
		i *beerproto.RecipeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Required fields",
			args: args{
				i: &beerproto.RecipeType{
					IbuEstimate: &beerproto.IBUEstimateType{
						Method: beerproto.IBUEstimateType_TINSETH,
					},
					Efficiency: &beerproto.EfficiencyType{
						Brewhouse: &beerproto.PercentType{
							Value: 80,
							Unit:  beerproto.PercentType_PERCENT_SIGN,
						},
					},
					Ingredients: &beerproto.IngredientsType{
						FermentableAdditions: []*beerproto.FermentableAdditionType{},
					},
					BatchSize: &beerproto.VolumeType{
						Unit:  beerproto.VolumeType_L,
						Value: float64(30),
					},
					Name:   "Name",
					Author: "Author",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONRecipeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONRecipeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			have := ToProtoRecipeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONRecipeType() %v", diff)
			}
		})
	}
}

func TestToJSONIngredientsType(t *testing.T) {
	type args struct {
		i *beerproto.IngredientsType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.IngredientsType{
					FermentableAdditions: []*beerproto.FermentableAdditionType{},
					HopAdditions: []*beerproto.HopAdditionType{},
					MiscellaneousAdditions: []*beerproto.MiscellaneousAdditionType{},
					CultureAdditions: []*beerproto.CultureAdditionType{},
					WaterAdditions: []*beerproto.WaterAdditionType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.IngredientsType{
					FermentableAdditions: []*beerproto.FermentableAdditionType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.IngredientsType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONIngredientsType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONIngredientsType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			have := ToProtoIngredientsType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONIngredientsType() %v", diff)
			}
		})
	}
}

func TestToJSONIBUEstimateType(t *testing.T) {
	type args struct {
		i *beerproto.IBUEstimateType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Rager",
			args: args{
				i: &beerproto.IBUEstimateType{
					Method: beerproto.IBUEstimateType_RAGER,
				},
			},
		},
		{
			name: "Garetz",
			args: args{
				i: &beerproto.IBUEstimateType{
					Method: beerproto.IBUEstimateType_GARETZ,
				},
			},
		},
		{
			name: "Tinseth",
			args: args{
				i: &beerproto.IBUEstimateType{
					Method: beerproto.IBUEstimateType_TINSETH,
				},
			},
		},
		{
			name: "Other",
			args: args{
				i: &beerproto.IBUEstimateType{
					Method: beerproto.IBUEstimateType_OTHER,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.IBUEstimateType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONIBUEstimateType(tt.args.i)
			have := ToProtoIBUEstimateType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONIBUEstimateType() %v", diff)
			}
		})
	}
}

func TestToJSONDiastaticPowerType(t *testing.T) {
	type args struct {
		i *beerproto.DiastaticPowerType
	}
	tests := []struct {
		name string
		args args
		want *beerjson.DiastaticPowerType
	}{
		{
			name: "Lintner",
			args: args{
				i: &beerproto.DiastaticPowerType{
					Value: float64(10),
					Unit:  beerproto.DiastaticPowerUnitType_LINTNER,
				},
			},
		},
		{
			name: "WK",
			args: args{
				i: &beerproto.DiastaticPowerType{
					Value: float64(10),
					Unit:  beerproto.DiastaticPowerUnitType_WK,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.DiastaticPowerType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONDiastaticPowerType(tt.args.i)
			have := ToProtoDiastaticPowerType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONDiastaticPowerType() %v", diff)
			}
		})
	}
}

func TestToJSONWaterBase(t *testing.T) {
	type args struct {
		i *beerproto.WaterBase
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.WaterBase{
					Calcium:     &beerproto.ConcentrationType{},
					Nitrite:     &beerproto.ConcentrationType{},
					Chloride:    &beerproto.ConcentrationType{},
					Name:        "Name",
					Potassium:   &beerproto.ConcentrationType{},
					Carbonate:   &beerproto.ConcentrationType{},
					Iron:        &beerproto.ConcentrationType{},
					Flouride:    &beerproto.ConcentrationType{},
					Sulfate:     &beerproto.ConcentrationType{},
					Magnesium:   &beerproto.ConcentrationType{},
					Producer:    "Producer",
					Bicarbonate: &beerproto.ConcentrationType{},
					Nitrate:     &beerproto.ConcentrationType{},
					Sodium:      &beerproto.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.WaterBase{
					Name: "Balanced Profile",
					Calcium: &beerproto.ConcentrationType{
						Value: 80,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Bicarbonate: &beerproto.ConcentrationType{
						Value: 100,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Sulfate: &beerproto.ConcentrationType{
						Value: 80,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Chloride: &beerproto.ConcentrationType{
						Value: 75,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Sodium: &beerproto.ConcentrationType{
						Value: 25,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Magnesium: &beerproto.ConcentrationType{
						Value: 5,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.WaterBase{},
			},
			wantErr: true,
		},
		{
			name: "Required",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONWaterBase(tt.args.i)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONWaterBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoWaterBase(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONWaterBase() %v", diff)
			}
		})
	}
}

func TestToJSONConcentrationType(t *testing.T) {
	type args struct {
		i *beerproto.ConcentrationType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "MGL",
			args: args{
				i: &beerproto.ConcentrationType{
					Value: 80,
					Unit:  beerproto.ConcentrationUnitType_MGL,
				},
			},
		},
		{
			name: "PPB",
			args: args{
				i: &beerproto.ConcentrationType{
					Value: 80,
					Unit:  beerproto.ConcentrationUnitType_PPB,
				},
			},
		},
		{
			name: "PPM",
			args: args{
				i: &beerproto.ConcentrationType{
					Value: 80,
					Unit:  beerproto.ConcentrationUnitType_PPM,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.ConcentrationType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONConcentrationType(tt.args.i)
			have := ToProtoConcentrationType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONConcentrationType() %v", diff)
			}
		})
	}
}

func TestMapToJSON(t *testing.T) {
	type args struct {
		i *beerproto.Recipe
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.Recipe{
					Version: 1,
					Mashes:                   []*beerproto.MashProcedureType{},
					Recipes:                  []*beerproto.RecipeType{},
					MiscellaneousIngredients: []*beerproto.MiscellaneousType{},
					Styles:                   []*beerproto.StyleType{},
					Fermentations:            []*beerproto.FermentationProcedureType{},
					Boil:                     []*beerproto.BoilProcedureType{},
					Fermentables:             []*beerproto.FermentableType{},
					TimingObject:             &beerproto.TimingType{},
					Cultures:                 []*beerproto.CultureInformation{},
					Equipments:               []*beerproto.EquipmentType{},
					Packaging:                []*beerproto.PackagingProcedureType{},
					HopVarieties:             []*beerproto.VarietyInformation{},
					Profiles:                 []*beerproto.WaterBase{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.Recipe{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToJSON(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := MapToProto(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("MapToJSON() %v", diff)
			}
		})
	}
}

func TestToJSONVarietyInformation(t *testing.T) {
	type args struct {
		i *beerproto.VarietyInformation
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.VarietyInformation{
					Inventory:   &beerproto.HopInventoryType{},
					Type:        beerproto.VarietyInformation_BITTERING,
					OilContent:  &beerproto.OilContentType{},
					PercentLost: &beerproto.PercentType{},
					ProductId:   "ProductId",
					AlphaAcid:   &beerproto.PercentType{},
					BetaAcid:    &beerproto.PercentType{},
					Name:        "Name",
					Origin:      "Origin",
					Substitutes: "Substitutes",
					Year:        "Year",
					Form:        beerproto.HopVarietyBaseForm_PELLET,
					Producer:    "Producer",
					Notes:       "Notes",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.VarietyInformation{
					AlphaAcid:   &beerproto.PercentType{},
					Name:        "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.VarietyInformation{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONVarietyInformation(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONRecipeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoVarietyInformation(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONVarietyInformation() %v", diff)
			}
		})
	}
}

func TestToJSONOilContentType(t *testing.T) {
	type args struct {
		i *beerproto.OilContentType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.OilContentType{
					Polyphenols:        &beerproto.PercentType{},
					TotalOilMlPer_100G: 100,
					Farnesene:          &beerproto.PercentType{},
					Limonene:           &beerproto.PercentType{},
					Nerol:              &beerproto.PercentType{},
					Geraniol:           &beerproto.PercentType{},
					BPinene:            &beerproto.PercentType{},
					Linalool:           &beerproto.PercentType{},
					Caryophyllene:      &beerproto.PercentType{},
					Cohumulone:         &beerproto.PercentType{},
					Xanthohumol:        &beerproto.PercentType{},
					Humulene:           &beerproto.PercentType{},
					Myrcene:            &beerproto.PercentType{},
					Pinene:             &beerproto.PercentType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.OilContentType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i:nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONOilContentType(tt.args.i)
			have := ToProtoOilContentType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONOilContentType() %v", diff)
			}
		})
	}
}

func TestToJSONHopInventoryType(t *testing.T) {
	type args struct {
		i *beerproto.HopInventoryType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Volume",
			args: args{
				i: &beerproto.HopInventoryType{
					Amount: &beerproto.HopInventoryType_Volume{},
				},
			},
		},
		{
			name: "Mass",
			args: args{
				i: &beerproto.HopInventoryType{
					Amount: &beerproto.HopInventoryType_Mass{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.HopInventoryType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONHopInventoryType(tt.args.i)
			have := ToProtoHopInventoryType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONHopInventoryType() %v", diff)
			}
		})
	}
}

func TestToJSONEquipmentType(t *testing.T) {
	type args struct {
		i *beerproto.EquipmentType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.EquipmentType{
					Name:           "Name",
					EquipmentItems: []*beerproto.EquipmentItemType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.EquipmentType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONEquipmentType(tt.args.i)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONEquipmentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoEquipmentType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONEquipmentType() %v", diff)
			}
		})
	}
}

func TestToJSONEquipmentItemType(t *testing.T) {
	type args struct {
		i *beerproto.EquipmentItemType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.EquipmentItemType{
					Name:                "Name",
					Notes:               "Notes",
					BoilRatePerHour:     &beerproto.VolumeType{},
					Type:                "Type",
					Form:                beerproto.EquipmentItemType_MASH_TUN,
					DrainRatePerMinute:  &beerproto.VolumeType{},
					SpecificHeat:        &beerproto.SpecificHeatType{},
					GrainAbsorptionRate: &beerproto.SpecificVolumeType{},
					MaximumVolume:       &beerproto.VolumeType{},
					Weight:              &beerproto.MassType{},
					Loss:                &beerproto.VolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.EquipmentItemType{
					Name:                "Name",
					Type:                "Type",
					Form:                beerproto.EquipmentItemType_MASH_TUN,
					MaximumVolume:       &beerproto.VolumeType{},
					Loss:                &beerproto.VolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.EquipmentItemType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONEquipmentItemType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONEquipmentItemType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoEquipmentItemType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONEquipmentItemType() %v", diff)
			}
		})
	}
}

func TestToJSONSpecificHeatType(t *testing.T) {
	type args struct {
		i *beerproto.SpecificHeatType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Calgc",
			args: args{
				i: &beerproto.SpecificHeatType{
					Value: 80,
					Unit:  beerproto.SpecificHeatUnitType_CALGC,
				},
			},
		},
		{
			name: "JKGK",
			args: args{
				i: &beerproto.SpecificHeatType{
					Value: 80,
					Unit:  beerproto.SpecificHeatUnitType_JKGK,
				},
			},
		},
		{
			name: "BTULBF",
			args: args{
				i: &beerproto.SpecificHeatType{
					Value: 80,
					Unit:  beerproto.SpecificHeatUnitType_BTULBF,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.SpecificHeatType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONSpecificHeatType(tt.args.i)
			have := ToProtoSpecificHeatType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONSpecificHeatType() %v", diff)
			}
		})
	}
}

func TestToJSONMassType(t *testing.T) {
	type args struct {
		i *beerproto.MassType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "G",
			args: args{
				i: &beerproto.MassType{
					Value: 90,
					Unit:  beerproto.MassUnitType_G,
				},
			},
		},
		{
			name: "KG",
			args: args{
				i: &beerproto.MassType{
					Value: 90,
					Unit:  beerproto.MassUnitType_KG,
				},
			},
		},
		{
			name: "LB",
			args: args{
				i: &beerproto.MassType{
					Value: 90,
					Unit:  beerproto.MassUnitType_LB,
				},
			},
		},
		{
			name: "MG",
			args: args{
				i: &beerproto.MassType{
					Value: 90,
					Unit:  beerproto.MassUnitType_MG,
				},
			},
		},
		{
			name: "OZ",
			args: args{
				i: &beerproto.MassType{
					Value: 90,
					Unit:  beerproto.MassUnitType_OZ,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MassType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONMassType(tt.args.i)
			have := ToProtoMassType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMassType() %v", diff)
			}
		})
	}
}

func TestToJSONVolumeType(t *testing.T) {
	type args struct {
		i *beerproto.VolumeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ML",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_ML,
				},
			},
		},
		{
			name: "L",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_L,
				},
			},
		},
		{
			name: "TSP",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_TSP,
				},
			},
		},
		{
			name: "TBSP",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_TBSP,
				},
			},
		},
		{
			name: "FLOZ",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_FLOZ,
				},
			},
		},
		{
			name: "CUP",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_CUP,
				},
			},
		},
		{
			name: "PT",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_PT,
				},
			},
		},
		{
			name: "QT",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_QT,
				},
			},
		},
		{
			name: "GAL",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_GAL,
				},
			},
		},
		{
			name: "BBL",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_BBL,
				},
			},
		},
		{
			name: "IFOZ",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_IFOZ,
				},
			},
		},
		{
			name: "IPT",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_IPT,
				},
			},
		},
		{
			name: "IQT",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_IQT,
				},
			},
		},
		{
			name: "IGAL",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_IGAL,
				},
			},
		},
		{
			name: "IBBL",
			args: args{
				i: &beerproto.VolumeType{
					Value: 50,
					Unit:  beerproto.VolumeType_IBBL,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.VolumeType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONVolumeType(tt.args.i)
			have := ToProtoVolumeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONVolumeType() %v", diff)
			}
		})
	}
}

func TestToJSONSpecificVolumeType(t *testing.T) {
	type args struct {
		i *beerproto.SpecificVolumeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "qt/lb",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_QTLB,
				},
			},
		},
		{
			name: "gal/lb",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_GALLB,
				},
			},
		},
		{
			name: "gal/oz",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_GALOZ,
				},
			},
		},
		{
			name: "l/g",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_LG,
				},
			},
		},
		{
			name: "l/kg",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_LKG,
				},
			},
		},
		{
			name: "floz/oz",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_FLOZOZ,
				},
			},
		},
		{
			name: "m^3/kg",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_M3KG,
				},
			},
		},
		{
			name: "ft^3/lb",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit:  beerproto.SpecificVolumeType_FT3LB,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.SpecificVolumeType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONSpecificVolumeType(tt.args.i)
			have := ToProtoSpecificVolumeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONSpecificVolumeType() %v", diff)
			}
		})
	}
}

func TestToJSONCultureInformation(t *testing.T) {
	type args struct {
		i *beerproto.CultureInformation
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.CultureInformation{
					Form:             beerproto.CultureBaseForm_DRY,
					Producer:         "Producer",
					TemperatureRange: nil,
					Notes:            "Notes",
					BestFor:          "BestFor",
					Inventory:        &beerproto.CultureInventoryType{},
					ProductId:        "ProductId",
					Name:             "Name",
					AlcoholTolerance: &beerproto.PercentType{},
					Glucoamylase:     true,
					Type:             beerproto.CultureBaseType_ALE,
					Flocculation:     beerproto.QualitativeRangeType_HIGH,
					AttenuationRange: nil,
					MaxReuse:         2,
					Pof:              true,
					Zymocide:         &beerproto.Zymocide{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.CultureInformation{
					Form:             beerproto.CultureBaseForm_DRY,
					Name:             "Name",
					Type:             beerproto.CultureBaseType_ALE,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.CultureInformation{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONCultureInformation(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONCultureInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoCultureInformation(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCultureInformation() %v", diff)
			}
		})
	}
}

func TestToJSONTemperatureRangeType(t *testing.T) {
	type args struct {
		i *beerproto.TemperatureRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool

	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.TemperatureRangeType{
					Maximum: &beerproto.TemperatureType{},
					Minimum: &beerproto.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.TemperatureRangeType{
					Minimum: &beerproto.TemperatureType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.TemperatureRangeType{
					Maximum: &beerproto.TemperatureType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONTemperatureRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONTemperatureRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoTemperatureRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTemperatureRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONTemperatureType(t *testing.T) {
	type args struct {
		i *beerproto.TemperatureType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "C",
			args: args{
				i: &beerproto.TemperatureType{
					Value: 5,
					Unit:  beerproto.TemperatureUnitType_C,
				},
			},
		},
		{
			name: "F",
			args: args{
				i: &beerproto.TemperatureType{
					Value: 5,
					Unit:  beerproto.TemperatureUnitType_F,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.TemperatureType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONTemperatureType(tt.args.i)
			have := ToProtoTemperatureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTemperatureType() %v", diff)
			}
		})
	}
}

func TestToJSONPercentRangeType(t *testing.T) {
	type args struct {
		i *beerproto.PercentRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.PercentRangeType{
					Maximum: &beerproto.PercentType{},
					Minimum: &beerproto.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.PercentRangeType{
					Minimum: &beerproto.PercentType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.PercentRangeType{
					Maximum: &beerproto.PercentType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONPercentRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONPercentRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoPercentRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONPercentRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONPercentType(t *testing.T) {
	type args struct {
		i *beerproto.PercentType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Percent",
			args: args{
				i: &beerproto.PercentType{
					Value: 75,
					Unit:  beerproto.PercentType_PERCENT_SIGN,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.PercentType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONPercentType(tt.args.i)
			have := ToProtoPercentType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONPercentType() %v", diff)
			}
		})
	}
}

func TestToJSONZymocide(t *testing.T) {
	type args struct {
		i *beerproto.Zymocide
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.Zymocide{
					No1:     true,
					No2:     true,
					No28:    true,
					Klus:    true,
					Neutral: true,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.Zymocide{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONZymocide(tt.args.i)
			have := ToProtoZymocide(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONZymocide() %v", diff)
			}
		})
	}
}

func TestToJSONCultureInventoryType(t *testing.T) {
	type args struct {
		i *beerproto.CultureInventoryType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.CultureInventoryType{
					Liquid:  &beerproto.VolumeType{},
					Dry:     &beerproto.MassType{},
					Slant:   &beerproto.VolumeType{},
					Culture: &beerproto.VolumeType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.CultureInventoryType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONCultureInventoryType(tt.args.i)
			have := ToProtoCultureInventoryType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCultureInventoryType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentableType(t *testing.T) {
	type args struct {
		i *beerproto.FermentableType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.FermentableType{
					MaxInBatch:     &beerproto.PercentType{},
					RecommendMash:  true,
					Protein:        &beerproto.PercentType{},
					ProductId:      "ProductId",
					GrainGroup:     beerproto.GrainGroup_BASE,
					Yield:          &beerproto.YieldType{},
					Type:           beerproto.FermentableBaseType_GRAIN,
					Producer:       "Producer",
					AlphaAmylase:   80,
					Color:          &beerproto.ColorType{},
					Name:           "Name",
					DiastaticPower: &beerproto.DiastaticPowerType{},
					Moisture:       &beerproto.PercentType{},
					Origin:         "Origin",
					Inventory:      &beerproto.FermentableInventoryType{},
					KolbachIndex:   80,
					Glassy:         &beerproto.PercentType{},
					Plump:          &beerproto.PercentType{},
					Half:           &beerproto.PercentType{},
					Mealy:          &beerproto.PercentType{},
					Thru:           &beerproto.PercentType{},
					Friability:     &beerproto.PercentType{},
					DiPh:           &beerproto.AcidityType{},
					Viscosity:      &beerproto.ViscosityType{},
					DmsP:           &beerproto.ConcentrationType{},
					Fan:            &beerproto.ConcentrationType{},
					Fermentability: &beerproto.PercentType{},
					BetaGlucan:     &beerproto.ConcentrationType{},
					Notes:          "Notes",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.FermentableType{
					Yield:          &beerproto.YieldType{},
					Type:           beerproto.FermentableBaseType_GRAIN,
					Color:          &beerproto.ColorType{},
					Name:           "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.FermentableType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONFermentableType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONFermentableType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoFermentableType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONFermentableType() %v", diff)
			}
		})
	}
}

func TestToJSONColorType(t *testing.T) {
	type args struct {
		i *beerproto.ColorType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "EBC",
			args: args{
				i: &beerproto.ColorType{
					Value: 10,
					Unit:  beerproto.ColorUnitType_EBC,
				},
			},
		},
		{
			name: "Lovi",
			args: args{
				i: &beerproto.ColorType{
					Value: 10,
					Unit:  beerproto.ColorUnitType_LOVI,
				},
			},
		},
		{
			name: "SRM",
			args: args{
				i: &beerproto.ColorType{
					Value: 10,
					Unit:  beerproto.ColorUnitType_SRM,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.ColorType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONColorType(tt.args.i)
			have := ToProtoColorType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONColorType() %v", diff)
			}
		})
	}
}

func TestToJSONAcidityType(t *testing.T) {
	type args struct {
		i *beerproto.AcidityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pH",
			args: args{
				i: &beerproto.AcidityType{
					Value: 10,
					Unit:  beerproto.AcidityUnitType_PH,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.AcidityType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONAcidityType(tt.args.i)
			have := ToProtoAcidityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONAcidityType() %v", diff)
			}
		})
	}
}

func TestToJSONViscosityType(t *testing.T) {
	type args struct {
		i *beerproto.ViscosityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mPa-s",
			args: args{
				i: &beerproto.ViscosityType{
					Value: 10,
					Unit:  beerproto.ViscosityUnitType_MPAS,
				},
			},
		},
		{
			name: "cP",
			args: args{
				i: &beerproto.ViscosityType{
					Value: 10,
					Unit:  beerproto.ViscosityUnitType_CP,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.ViscosityType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONViscosityType(tt.args.i)
			have := ToProtoViscosityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONViscosityType() %v", diff)
			}
		})
	}
}

func TestToJSONStyleType(t *testing.T) {
	type args struct {
		i *beerproto.StyleType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.StyleType{
					Aroma:          "Aroma",
					Ingredients:    "Ingredients",
					CategoryNumber: 10,
					Notes:          "Notes",
					Flavor:         "Flavor",
					Mouthfeel:      "Mouthfeel",
					FinalGravity: &beerproto.GravityRangeType{
						Maximum: &beerproto.GravityType{},
						Minimum: &beerproto.GravityType{},
					},
					StyleGuide: "Style",
					Color: &beerproto.ColorRangeType{
						Minimum: &beerproto.ColorType{},
						Maximum: &beerproto.ColorType{},
					},
					OriginalGravity: &beerproto.GravityRangeType{
						Maximum: &beerproto.GravityType{},
						Minimum: &beerproto.GravityType{},
					},
					Examples: "Example",
					Name:     "Name",
					Carbonation: &beerproto.CarbonationRangeType{
						Maximum: &beerproto.CarbonationType{},
						Minimum: &beerproto.CarbonationType{},
					},
					AlcoholByVolume: &beerproto.PercentRangeType{
						Minimum: &beerproto.PercentType{},
						Maximum: &beerproto.PercentType{},
					},
					InternationalBitternessUnits: &beerproto.BitternessRangeType{
						Minimum: &beerproto.BitternessType{},
						Maximum: &beerproto.BitternessType{},
					},
					Appearance:        "Appearance",
					Category:          "Category",
					StyleLetter:       "StyleLetter",
					Type:              beerproto.StyleType_BEER,
					OverallImpression: "Overall Impression",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.StyleType{
					Name: "Name",
					StyleGuide: "StyleGuide",
					Category: "Category",
					Type: beerproto.StyleType_CIDER,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.StyleType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONStyleType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONStyleType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoStyleType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONStyleType() %v", diff)
			}
		})
	}
}

func TestToJSONGravityRangeType(t *testing.T) {
	type args struct {
		i *beerproto.GravityRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.GravityRangeType{
					Minimum: &beerproto.GravityType{},
					Maximum: &beerproto.GravityType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.GravityRangeType{
					Minimum: &beerproto.GravityType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.GravityRangeType{
					Maximum: &beerproto.GravityType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONGravityRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONGravityRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoGravityRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONGravityRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONGravityType(t *testing.T) {
	type args struct {
		i *beerproto.GravityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SG",
			args: args{
				i: &beerproto.GravityType{
					Value: 45,
					Unit:  beerproto.GravityUnitType_SG,
				},
			},
		},
		{
			name: "BRIX",
			args: args{
				i: &beerproto.GravityType{
					Value: 45,
					Unit:  beerproto.GravityUnitType_BRIX,
				},
			},
		},
		{
			name: "Plato",
			args: args{
				i: &beerproto.GravityType{
					Value: 45,
					Unit:  beerproto.GravityUnitType_PLATO,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.GravityType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONGravityType(tt.args.i)
			have := ToProtoGravityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONGravityType() %v", diff)
			}
		})
	}
}

func TestToJSONColorRangeType(t *testing.T) {
	type args struct {
		i *beerproto.ColorRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.ColorRangeType{
					Minimum: &beerproto.ColorType{},
					Maximum: &beerproto.ColorType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.ColorRangeType{
					Minimum: &beerproto.ColorType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.ColorRangeType{
					Maximum: &beerproto.ColorType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONColorRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONColorRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoColorRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONColorRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONCarbonationRangeType(t *testing.T) {
	type args struct {
		i *beerproto.CarbonationRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.CarbonationRangeType{
					Minimum: &beerproto.CarbonationType{},
					Maximum: &beerproto.CarbonationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.CarbonationRangeType{
					Minimum: &beerproto.CarbonationType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.CarbonationRangeType{
					Maximum: &beerproto.CarbonationType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ToJSONCarbonationRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONGravityRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoCarbonationRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCarbonationRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONCarbonationType(t *testing.T) {
	type args struct {
		i *beerproto.CarbonationType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Vols",
			args: args{
				i: &beerproto.CarbonationType{
					Value: 45,
					Unit:  beerproto.CarbonationUnitType_VOLS,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.CarbonationType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONCarbonationType(tt.args.i)
			have := ToProtoCarbonationType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCarbonationType() %v", diff)
			}
		})
	}
}

func TestToJSONBitternessRangeType(t *testing.T) {
	type args struct {
		i *beerproto.BitternessRangeType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.BitternessRangeType{
					Minimum: &beerproto.BitternessType{},
					Maximum: &beerproto.BitternessType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerproto.BitternessRangeType{
					Minimum: &beerproto.BitternessType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerproto.BitternessRangeType{
					Maximum: &beerproto.BitternessType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Nul",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONBitternessRangeType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONBitternessRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoBitternessRangeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONBitternessRangeType() %v", diff)
			}
		})
	}
}

func TestToJSONBitternessType(t *testing.T) {
	type args struct {
		i *beerproto.BitternessType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "IBU",
			args: args{
				i: &beerproto.BitternessType{
					Value: 35,
					Unit:  beerproto.BitternessType_IBUs,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.BitternessType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONBitternessType(tt.args.i)
			have := ToProtoBitternessType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONBitternessType() %v", diff)
			}
		})
	}
}

func TestToJSONMiscellaneousType(t *testing.T) {
	type args struct {
		i *beerproto.MiscellaneousType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.MiscellaneousType{
					UseFor:    "UseFor",
					Notes:     "Notes",
					Name:      "Name",
					Producer:  "Producer",
					ProductId: "ProductId",
					Type:      beerproto.MiscellaneousBaseType_FLAVOR,
					Inventory: &beerproto.MiscellaneousInventoryType{
						Amount: &beerproto.MiscellaneousInventoryType_Mass{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.MiscellaneousType{
					Name:      "Name",
					Type:      beerproto.MiscellaneousBaseType_FLAVOR,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MiscellaneousType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONMiscellaneousType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONMiscellaneousType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoMiscellaneousType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMiscellaneousType() %v", diff)
			}
		})
	}
}

func TestToJSONMiscellaneousInventoryType(t *testing.T) {
	type args struct {
		i *beerproto.MiscellaneousInventoryType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerproto.MiscellaneousInventoryType{
					Amount: &beerproto.MiscellaneousInventoryType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerproto.MiscellaneousInventoryType{
					Amount: &beerproto.MiscellaneousInventoryType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerproto.MiscellaneousInventoryType{
					Amount: &beerproto.MiscellaneousInventoryType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MiscellaneousInventoryType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONMiscellaneousInventoryType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONGravityRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoMiscellaneousInventoryType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMiscellaneousInventoryType() %v", diff)
			}
		})
	}
}

func TestToJSONTasteType(t *testing.T) {
	type args struct {
		i *beerproto.TasteType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.TasteType{
					Notes:  "Notes",
					Rating: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.TasteType{
					Notes:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.TasteType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONTasteType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONTasteType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoTasteType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTasteType() %v", diff)
			}
		})
	}
}

func TestToJSONBoilProcedureType(t *testing.T) {
	type args struct {
		i *beerproto.BoilProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.BoilProcedureType{
					PreBoilSize: &beerproto.VolumeType{},
					BoilTime: &beerproto.TimeType{
						Value: 60,
						Unit:  beerproto.TimeType_MIN,
					},
					BoilSteps:   []*beerproto.BoilStepType{},
					Name:        "Name",
					Description: "Description",
					Notes:       "Notes",
				},
			},
			wantErr: false,
		},
		{
			name: "BoilTime error",
			args: args{
				i: &beerproto.BoilProcedureType{
					PreBoilSize: &beerproto.VolumeType{},
					BoilTime:    nil,
					BoilSteps:   []*beerproto.BoilStepType{},
					Name:        "Name",
					Description: "Description",
					Notes:       "Notes",
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ToJSONBoilProcedureType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONBoilProcedureType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			have := ToProtoBoilProcedureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONBoilProcedureType() %v", diff)
			}
		})
	}
}

func TestToJSONBoilStepType(t *testing.T) {
	type args struct {
		i *beerproto.BoilStepType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Batch",
			args: args{
				i: &beerproto.BoilStepType{
					EndGravity:       &beerproto.GravityType{},
					ChillingType:     beerproto.BoilStepTypeChillingType_BATCH,
					Description:      "Description",
					EndTemperature:   &beerproto.TemperatureType{},
					RampTime:         &beerproto.TimeType{},
					StepTime:         &beerproto.TimeType{},
					StartGravity:     &beerproto.GravityType{},
					StartPh:          &beerproto.AcidityType{},
					EndPh:            &beerproto.AcidityType{},
					Name:             "Name",
					StartTemperature: &beerproto.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Inline",
			args: args{
				i: &beerproto.BoilStepType{
					EndGravity:       &beerproto.GravityType{},
					ChillingType:     beerproto.BoilStepTypeChillingType_INLINE,
					Description:      "Description",
					EndTemperature:   &beerproto.TemperatureType{},
					RampTime:         &beerproto.TimeType{},
					StepTime:         &beerproto.TimeType{},
					StartGravity:     &beerproto.GravityType{},
					StartPh:          &beerproto.AcidityType{},
					EndPh:            &beerproto.AcidityType{},
					Name:             "Name",
					StartTemperature: &beerproto.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.BoilStepType{
					Name: "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.BoilStepType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONBoilStepType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONGravityRangeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoBoilStepType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONBoilStepType() %v", diff)
			}
		})
	}
}

func TestToJSONTimeType(t *testing.T) {
	type args struct {
		i *beerproto.TimeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sec",
			args: args{
				i: &beerproto.TimeType{
					Value: 10,
					Unit:  beerproto.TimeType_SEC,
				},
			},
		},
		{
			name: "Min",
			args: args{
				i: &beerproto.TimeType{
					Value: 10,
					Unit:  beerproto.TimeType_MIN,
				},
			},
		},
		{
			name: "Hour",
			args: args{
				i: &beerproto.TimeType{
					Value: 10,
					Unit:  beerproto.TimeType_HR,
				},
			},
		},
		{
			name: "Day",
			args: args{
				i: &beerproto.TimeType{
					Value: 10,
					Unit:  beerproto.TimeType_DAY,
				},
			},
		},
		{
			name: "Week",
			args: args{
				i: &beerproto.TimeType{
					Value: 10,
					Unit:  beerproto.TimeType_WEEK,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.TimeType{},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONTimeType(tt.args.i)
			have := ToProtoTimeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTimeType() %v", diff)
			}
		})
	}
}

func TestToJSONPackagingProcedureType(t *testing.T) {
	type args struct {
		i *beerproto.PackagingProcedureType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.PackagingProcedureType{
					Name:             "Name",
					PackagedVolume:   &beerproto.VolumeType{},
					Description:      "Description",
					Notes:            "Notes",
					PackagingVessels: []*beerproto.PackagingVesselType{},
				},
			},
			wantErr: false,

		},
		{
			name: "Required",
			args: args{
				i: &beerproto.PackagingProcedureType{
					Name: "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.PackagingProcedureType{},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONPackagingProcedureType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONPackagingProcedureType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoPackagingProcedureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONPackagingProcedureType() %v", diff)
			}
		})
	}
}

func TestToJSONPackagingVesselType(t *testing.T) {
	type args struct {
		i *beerproto.PackagingVesselType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.PackagingVesselType{
					Type:             beerproto.PackagingVesselType_BOTTLE,
					StartGravity:     &beerproto.GravityType{},
					Name:             "Name",
					PackageDate:      "PackageDate",
					StepTime:         &beerproto.TimeType{},
					EndGravity:       &beerproto.GravityType{},
					VesselVolume:     &beerproto.VolumeType{},
					VesselQuantity:   10,
					Description:      "Description",
					StartPh:          &beerproto.AcidityType{},
					Carbonation:      80,
					StartTemperature: &beerproto.TemperatureType{},
					EndPh:            &beerproto.AcidityType{},
					EndTemperature:   &beerproto.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.PackagingVesselType{
					Name: "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.PackagingVesselType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONPackagingVesselType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONPackagingVesselType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoPackagingVesselType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONPackagingVesselType() %v", diff)
			}
		})
	}
}

func TestToJSONHopAdditionType(t *testing.T) {
	type args struct {
		i *beerproto.HopAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerproto.HopAdditionType{
					BetaAcid:  &beerproto.PercentType{},
					Producer:  "Producer",
					Origin:    "Origin",
					Year:      "Year",
					Form:      beerproto.HopVarietyBaseForm_PELLET,
					Timing:    &beerproto.TimingType{},
					Name:      "Name",
					ProductId: "ProductId",
					AlphaAcid: &beerproto.PercentType{},
					Amount:    &beerproto.HopAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerproto.HopAdditionType{
					BetaAcid:  &beerproto.PercentType{},
					Producer:  "Producer",
					Origin:    "Origin",
					Year:      "Year",
					Form:      beerproto.HopVarietyBaseForm_POWDER,
					Timing:    &beerproto.TimingType{},
					Name:      "Name",
					ProductId: "ProductId",
					AlphaAcid: &beerproto.PercentType{},
					Amount:    &beerproto.HopAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.HopAdditionType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONHopAdditionType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONHopAdditionType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			have := ToProtoHopAdditionType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONHopAdditionType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentableAdditionType(t *testing.T) {
	type args struct {
		i *beerproto.FermentableAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerproto.FermentableAdditionType{
					Type:       beerproto.FermentableBaseType_GRAIN,
					Origin:     "Origin",
					GrainGroup: beerproto.GrainGroup_BASE,
					Yield:      &beerproto.YieldType{},
					Color:      &beerproto.ColorType{},
					Name:       "Name",
					Producer:   "Producer",
					ProductId:  "ProductId",
					Timing:     &beerproto.TimingType{},
					Amount:     &beerproto.FermentableAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerproto.FermentableAdditionType{
					Type:       beerproto.FermentableBaseType_DRY_EXTRACT,
					Origin:     "Origin",
					GrainGroup: beerproto.GrainGroup_CARAMEL,
					Yield:      &beerproto.YieldType{},
					Color:      &beerproto.ColorType{},
					Name:       "Name",
					Producer:   "Producer",
					ProductId:  "ProductId",
					Timing:     &beerproto.TimingType{},
					Amount:     &beerproto.FermentableAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.FermentableAdditionType{
					Type:       beerproto.FermentableBaseType_DRY_EXTRACT,
					Yield:      &beerproto.YieldType{},
					Color:      &beerproto.ColorType{},
					Name:       "Name",
					Timing:     &beerproto.TimingType{},
					Amount:     &beerproto.FermentableAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.FermentableAdditionType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONFermentableAdditionType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONFermentableAdditionType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			have := ToProtoFermentableAdditionType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONFermentableAdditionType() %v", diff)
			}
		})
	}
}

func TestToJSONYieldType(t *testing.T) {
	type args struct {
		i *beerproto.YieldType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.YieldType{
					FineGrind:            &beerproto.PercentType{},
					CoarseGrind:          &beerproto.PercentType{},
					FineCoarseDifference: &beerproto.PercentType{},
					Potential:            &beerproto.GravityType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.YieldType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONYieldType(tt.args.i)
			have := ToProtoYieldType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONYieldType() %v", diff)
			}
		})
	}
}

func TestToJSONWaterAdditionType(t *testing.T) {
	type args struct {
		i *beerproto.WaterAdditionType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.WaterAdditionType{
					Flouride:    &beerproto.ConcentrationType{},
					Sulfate:     &beerproto.ConcentrationType{},
					Producer:    "Producer",
					Nitrate:     &beerproto.ConcentrationType{},
					Nitrite:     &beerproto.ConcentrationType{},
					Chloride:    &beerproto.ConcentrationType{},
					Amount:      &beerproto.VolumeType{},
					Name:        "Name",
					Potassium:   &beerproto.ConcentrationType{},
					Magnesium:   &beerproto.ConcentrationType{},
					Iron:        &beerproto.ConcentrationType{},
					Bicarbonate: &beerproto.ConcentrationType{},
					Calcium:     &beerproto.ConcentrationType{},
					Carbonate:   &beerproto.ConcentrationType{},
					Sodium:      &beerproto.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.WaterAdditionType{
					Sulfate:     &beerproto.ConcentrationType{},
					Chloride:    &beerproto.ConcentrationType{},
					Name:        "Name",
					Magnesium:   &beerproto.ConcentrationType{},
					Bicarbonate: &beerproto.ConcentrationType{},
					Calcium:     &beerproto.ConcentrationType{},
					Sodium:      &beerproto.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.WaterAdditionType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONWaterAdditionType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONWaterAdditionType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoWaterAdditionType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONWaterAdditionType() %v", diff)
			}
		})
	}
}

func TestToJSONCultureAdditionType(t *testing.T) {
	type args struct {
		i *beerproto.CultureAdditionType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerproto.CultureAdditionType{
					Form:              beerproto.CultureBaseForm_DRY,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerproto.CultureBaseType_ALE,
					Attenuation:       &beerproto.PercentType{},
					Timing:            &beerproto.TimingType{},
					Amount:            &beerproto.CultureAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerproto.CultureAdditionType{
					Form:              beerproto.CultureBaseForm_CULTURE,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerproto.CultureBaseType_LAGER,
					Attenuation:       &beerproto.PercentType{},
					Timing:            &beerproto.TimingType{},
					Amount:            &beerproto.CultureAdditionType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerproto.CultureAdditionType{
					Form:              beerproto.CultureBaseForm_LIQUID,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerproto.CultureBaseType_BACTERIA,
					Attenuation:       &beerproto.PercentType{},
					Timing:            &beerproto.TimingType{},
					Amount:            &beerproto.CultureAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.CultureAdditionType{
					Form:              beerproto.CultureBaseForm_LIQUID,
					Name:              "Name",
					Type:              beerproto.CultureBaseType_BACTERIA,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.CultureAdditionType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONCultureAdditionType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONCultureAdditionType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoCultureAdditionType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCultureAdditionType() %v", diff)
			}
		})
	}
}

func TestToJSONMiscellaneousAdditionType(t *testing.T) {
	type args struct {
		i *beerproto.MiscellaneousAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerproto.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerproto.TimingType{},
					ProductId: "ProductId",
					Type:      beerproto.MiscellaneousBaseType_FLAVOR,
					Amount:    &beerproto.MiscellaneousAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerproto.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerproto.TimingType{},
					ProductId: "ProductId",
					Type:      beerproto.MiscellaneousBaseType_FINING,
					Amount:    &beerproto.MiscellaneousAdditionType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerproto.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerproto.TimingType{},
					ProductId: "ProductId",
					Type:      beerproto.MiscellaneousBaseType_HERB,
					Amount:    &beerproto.MiscellaneousAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.MiscellaneousAdditionType{
					Name: "Name",
					Type: beerproto.MiscellaneousBaseType_SPICE,
					Timing: &beerproto.TimingType{},
					Amount: &beerproto.MiscellaneousAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MiscellaneousAdditionType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONMiscellaneousAdditionType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONMiscellaneousAdditionType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			have := ToProtoMiscellaneousAdditionType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMiscellaneousAdditionType() %v", diff)
			}
		})
	}
}

func TestToJSONUnitType(t *testing.T) {
	type args struct {
		i *beerproto.UnitType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "one",
			args: args{
				i: &beerproto.UnitType{
					Value: 60,
					Unit:  beerproto.UnitUnitType_ONE,
				},
			},
		},
		{
			name: "unit",
			args: args{
				i: &beerproto.UnitType{
					Value: 60,
					Unit:  beerproto.UnitUnitType_UNIT,
				},
			},
		},
		{
			name: "each",
			args: args{
				i: &beerproto.UnitType{
					Value: 60,
					Unit:  beerproto.UnitUnitType_EACH,
				},
			},
		},
		{
			name: "dimensionless",
			args: args{
				i: &beerproto.UnitType{
					Value: 60,
					Unit:  beerproto.UnitUnitType_DIMENSIONLESS,
				},
			},
		},
		{
			name: "pkg",
			args: args{
				i: &beerproto.UnitType{
					Value: 60,
					Unit:  beerproto.UnitUnitType_PKG,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.UnitType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONUnitType(tt.args.i)
			have := ToProtoUnitType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONUnitType() %v", diff)
			}
		})
	}
}

func TestToJSONTimingType(t *testing.T) {
	type args struct {
		i *beerproto.TimingType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.TimingType{
					Time:            &beerproto.TimeType{},
					Duration:        &beerproto.TimeType{},
					Continuous:      true,
					SpecificGravity: &beerproto.GravityType{},
					Ph:              &beerproto.AcidityType{},
					Step:            4,
					Use:             1,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.TimingType{
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONTimingType(tt.args.i)
			have := ToProtoTimingType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTimingType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentationProcedureType(t *testing.T) {
	type args struct {
		i *beerproto.FermentationProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.FermentationProcedureType{
					Description:       "Description",
					Notes:             "Notes",
					FermentationSteps: []*beerproto.FermentationStepType{},
					Name:              "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.FermentationProcedureType{
					Name:              "Name",
					FermentationSteps: []*beerproto.FermentationStepType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Basic",
			args: args{
				i: &beerproto.FermentationProcedureType{
				},
			},
			wantErr: true,
		},
		{
			name: "Basic",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONFermentationProcedureType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONFermentationProcedureType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoFermentationProcedureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONFermentationProcedureType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentationStepType(t *testing.T) {
	type args struct {
		i *beerproto.FermentationStepType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.FermentationStepType{
					Name:             "Name",
					EndTemperature:   &beerproto.TemperatureType{},
					StepTime:         &beerproto.TimeType{},
					FreeRise:         true,
					StartGravity:     &beerproto.GravityType{},
					EndGravity:       &beerproto.GravityType{},
					StartPh:          &beerproto.AcidityType{},
					Description:      "Description",
					StartTemperature: &beerproto.TemperatureType{},
					EndPh:            &beerproto.AcidityType{},
					Vessel:           "Vessel",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.FermentationStepType{
					Name: "Name",

				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.FermentationStepType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONFermentationStepType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONFermentationStepType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoFermentationStepType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONFermentationStepType() %v", diff)
			}
		})
	}
}

func TestToJSONRecipeStyleType(t *testing.T) {
	type args struct {
		i *beerproto.RecipeStyleType
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.RecipeStyleType{
					Type:           beerproto.RecipeStyleType_BEER,
					Name:           "Name",
					Category:       "Category",
					CategoryNumber: 10,
					StyleLetter:    "StyleLetter",
					StyleGuide:     "StyleGuide",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.RecipeStyleType{
					Type:           beerproto.RecipeStyleType_BEER,
					Name:           "Name",
					Category:       "Category",
					StyleGuide:     "StyleGuide",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.RecipeStyleType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONRecipeStyleType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONRecipeStyleType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoRecipeStyleType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONRecipeStyleType() %v", diff)
			}
		})
	}
}

func TestToJSONEfficiencyType(t *testing.T) {
	type args struct {
		i *beerproto.EfficiencyType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.EfficiencyType{
					Conversion: &beerproto.PercentType{},
					Lauter:     &beerproto.PercentType{},
					Mash:       &beerproto.PercentType{},
					Brewhouse:  &beerproto.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.EfficiencyType{
					Brewhouse: &beerproto.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.EfficiencyType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONEfficiencyType(tt.args.i)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONRecipeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoEfficiencyType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONEfficiencyType() %v", diff)
			}
		})
	}
}

func TestToJSONMashProcedureType(t *testing.T) {
	type args struct {
		i *beerproto.MashProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.MashProcedureType{
					GrainTemperature: &beerproto.TemperatureType{},
					Notes:            "Notes",
					MashSteps:        []*beerproto.MashStepType{},
					Name:             "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.MashProcedureType{
					GrainTemperature: &beerproto.TemperatureType{},
					MashSteps:        []*beerproto.MashStepType{},
					Name:             "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MashProcedureType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ToJSONMashProcedureType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONMashProcedureType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoMashProcedureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMashProcedureType() %v", diff)
			}
		})
	}
}

func TestToJSONMashStepType(t *testing.T) {
	type args struct {
		i *beerproto.MashStepType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.MashStepType{
					StepTime:          &beerproto.TimeType{},
					RampTime:          &beerproto.TimeType{},
					EndTemperature:    &beerproto.TemperatureType{},
					Description:       "Description",
					InfuseTemperature: &beerproto.TemperatureType{},
					StartPh:           &beerproto.AcidityType{},
					EndPh:             &beerproto.AcidityType{},
					Name:              "Name",
					Type:              beerproto.MashStepType_DRAIN_MASH_TUN,
					Amount:            &beerproto.VolumeType{},
					StepTemperature:   &beerproto.TemperatureType{},
					WaterGrainRatio:   &beerproto.SpecificVolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerproto.MashStepType{
					StepTime:        &beerproto.TimeType{},
					Name:            "Name",
					StepTemperature: &beerproto.TemperatureType{},
					Type:            beerproto.MashStepType_DECOCTION,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.MashStepType{
				},
			},
			wantErr: true,
		},
		{
			name: "Nil",
			args: args{
				i: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSONMashStepType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONMashProcedureType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			have := ToProtoMashStepType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMashStepType() %v", diff)
			}
		})
	}
}
