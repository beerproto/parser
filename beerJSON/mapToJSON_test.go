package beerJSON

import (
	"context"
	"testing"

	"github.com/beerproto/beerjson.go"
	beerprotov1 "github.com/beerproto/beerproto_go"
	"github.com/go-test/deep"
)

func TestToJSONRecipeType(t *testing.T) {
	type args struct {
		i *beerprotov1.RecipeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Required fields",
			args: args{
				i: &beerprotov1.RecipeType{
					IbuEstimate: &beerprotov1.IBUEstimateType{
						Method: beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_RAGER,
					},
					Efficiency: &beerprotov1.EfficiencyType{
						Brewhouse: &beerprotov1.PercentType{
							Value: 80,
							Unit:  beerprotov1.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
						},
					},
					Ingredients: &beerprotov1.IngredientsType{
						FermentableAdditions: []*beerprotov1.FermentableAdditionType{},
					},
					BatchSize: &beerprotov1.VolumeType{
						Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_L,
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

			got, err := ToJSONRecipeType(context.Background(), tt.args.i)
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
		i *beerprotov1.IngredientsType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.IngredientsType{
					FermentableAdditions:   []*beerprotov1.FermentableAdditionType{},
					HopAdditions:           []*beerprotov1.HopAdditionType{},
					MiscellaneousAdditions: []*beerprotov1.MiscellaneousAdditionType{},
					CultureAdditions:       []*beerprotov1.CultureAdditionType{},
					WaterAdditions:         []*beerprotov1.WaterAdditionType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.IngredientsType{
					FermentableAdditions: []*beerprotov1.FermentableAdditionType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.IngredientsType{},
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
			got, err := ToJSONIngredientsType(context.Background(), tt.args.i)
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
		i *beerprotov1.IBUEstimateType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Rager",
			args: args{
				i: &beerprotov1.IBUEstimateType{
					Method: beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_RAGER,
				},
			},
		},
		{
			name: "Garetz",
			args: args{
				i: &beerprotov1.IBUEstimateType{
					Method: beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_GARETZ,
				},
			},
		},
		{
			name: "Tinseth",
			args: args{
				i: &beerprotov1.IBUEstimateType{
					Method: beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_TINSETH,
				},
			},
		},
		{
			name: "Other",
			args: args{
				i: &beerprotov1.IBUEstimateType{
					Method: beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_OTHER,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.IBUEstimateType{},
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
			got := ToJSONIBUEstimateType(context.Background(), tt.args.i)
			have := ToProtoIBUEstimateType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONIBUEstimateType() %v", diff)
			}
		})
	}
}

func TestToJSONDiastaticPowerType(t *testing.T) {
	type args struct {
		i *beerprotov1.DiastaticPowerType
	}
	tests := []struct {
		name string
		args args
		want *beerjson.DiastaticPowerType
	}{
		{
			name: "Lintner",
			args: args{
				i: &beerprotov1.DiastaticPowerType{
					Value: float64(10),
					Unit:  beerprotov1.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_LINTNER,
				},
			},
		},
		{
			name: "WK",
			args: args{
				i: &beerprotov1.DiastaticPowerType{
					Value: float64(10),
					Unit:  beerprotov1.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.DiastaticPowerType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONDiastaticPowerType(context.Background(), tt.args.i)
			have := ToProtoDiastaticPowerType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONDiastaticPowerType() %v", diff)
			}
		})
	}
}

func TestToJSONWaterBase(t *testing.T) {
	type args struct {
		i *beerprotov1.WaterBase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.WaterBase{
					Calcium:     &beerprotov1.ConcentrationType{},
					Nitrite:     &beerprotov1.ConcentrationType{},
					Chloride:    &beerprotov1.ConcentrationType{},
					Name:        "Name",
					Potassium:   &beerprotov1.ConcentrationType{},
					Carbonate:   &beerprotov1.ConcentrationType{},
					Iron:        &beerprotov1.ConcentrationType{},
					Flouride:    &beerprotov1.ConcentrationType{},
					Sulfate:     &beerprotov1.ConcentrationType{},
					Magnesium:   &beerprotov1.ConcentrationType{},
					Producer:    "Producer",
					Bicarbonate: &beerprotov1.ConcentrationType{},
					Nitrate:     &beerprotov1.ConcentrationType{},
					Sodium:      &beerprotov1.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.WaterBase{
					Name: "Balanced Profile",
					Calcium: &beerprotov1.ConcentrationType{
						Value: 80,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
					Bicarbonate: &beerprotov1.ConcentrationType{
						Value: 100,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
					Sulfate: &beerprotov1.ConcentrationType{
						Value: 80,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
					Chloride: &beerprotov1.ConcentrationType{
						Value: 75,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
					Sodium: &beerprotov1.ConcentrationType{
						Value: 25,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
					Magnesium: &beerprotov1.ConcentrationType{
						Value: 5,
						Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.WaterBase{},
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
			got, err := ToJSONWaterBase(context.Background(), tt.args.i)

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
		i *beerprotov1.ConcentrationType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "MGL",
			args: args{
				i: &beerprotov1.ConcentrationType{
					Value: 80,
					Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL,
				},
			},
		},
		{
			name: "PPB",
			args: args{
				i: &beerprotov1.ConcentrationType{
					Value: 80,
					Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_PPB,
				},
			},
		},
		{
			name: "PPM",
			args: args{
				i: &beerprotov1.ConcentrationType{
					Value: 80,
					Unit:  beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_PPM,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.ConcentrationType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONConcentrationType(context.Background(), tt.args.i)
			have := ToProtoConcentrationType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONConcentrationType() %v", diff)
			}
		})
	}
}

func TestMapToJSON(t *testing.T) {
	type args struct {
		i *beerprotov1.Recipe
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.Recipe{
					Version:                  1,
					Mashes:                   []*beerprotov1.MashProcedureType{},
					Recipes:                  []*beerprotov1.RecipeType{},
					MiscellaneousIngredients: []*beerprotov1.MiscellaneousType{},
					Styles:                   []*beerprotov1.StyleType{},
					Fermentations:            []*beerprotov1.FermentationProcedureType{},
					Boil:                     []*beerprotov1.BoilProcedureType{},
					Fermentables:             []*beerprotov1.FermentableType{},
					Cultures:                 []*beerprotov1.CultureInformation{},
					Equipments:               []*beerprotov1.EquipmentType{},
					Packaging:                []*beerprotov1.PackagingProcedureType{},
					HopVarieties:             []*beerprotov1.VarietyInformation{},
					Profiles:                 []*beerprotov1.WaterBase{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.Recipe{},
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
			got, err := MapToJSON(context.Background(), tt.args.i)
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
		i *beerprotov1.VarietyInformation
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.VarietyInformation{
					Inventory:   &beerprotov1.HopInventoryType{},
					Type:        beerprotov1.VarietyInformationType_VARIETY_INFORMATION_TYPE_BITTERING,
					OilContent:  &beerprotov1.OilContentType{},
					PercentLost: &beerprotov1.PercentType{},
					ProductId:   "ProductId",
					AlphaAcid:   &beerprotov1.PercentType{},
					BetaAcid:    &beerprotov1.PercentType{},
					Name:        "Name",
					Origin:      "Origin",
					Substitutes: "Substitutes",
					Year:        "Year",
					Form:        beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PELLET,
					Producer:    "Producer",
					Notes:       "Notes",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.VarietyInformation{
					AlphaAcid: &beerprotov1.PercentType{},
					Name:      "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.VarietyInformation{},
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
			got, err := ToJSONVarietyInformation(context.Background(), tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONVarietyInformation() error = %v, wantErr %v", err, tt.wantErr)
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
		i *beerprotov1.OilContentType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.OilContentType{
					Polyphenols:        &beerprotov1.PercentType{},
					TotalOilMlPer_100G: 100,
					Farnesene:          &beerprotov1.PercentType{},
					Limonene:           &beerprotov1.PercentType{},
					Nerol:              &beerprotov1.PercentType{},
					Geraniol:           &beerprotov1.PercentType{},
					BPinene:            &beerprotov1.PercentType{},
					Linalool:           &beerprotov1.PercentType{},
					Caryophyllene:      &beerprotov1.PercentType{},
					Cohumulone:         &beerprotov1.PercentType{},
					Xanthohumol:        &beerprotov1.PercentType{},
					Humulene:           &beerprotov1.PercentType{},
					Myrcene:            &beerprotov1.PercentType{},
					Pinene:             &beerprotov1.PercentType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.OilContentType{},
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
			got := ToJSONOilContentType(context.Background(), tt.args.i)
			have := ToProtoOilContentType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONOilContentType() %v", diff)
			}
		})
	}
}

func TestToJSONHopInventoryType(t *testing.T) {
	type args struct {
		i *beerprotov1.HopInventoryType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.HopInventoryType{
					Amount: &beerprotov1.HopInventoryType_Volume{},
				},
			},
		},
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.HopInventoryType{
					Amount: &beerprotov1.HopInventoryType_Mass{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.HopInventoryType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONHopInventoryType(context.Background(), tt.args.i)
			have := ToProtoHopInventoryType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONHopInventoryType() %v", diff)
			}
		})
	}
}

func TestToJSONEquipmentType(t *testing.T) {
	type args struct {
		i *beerprotov1.EquipmentType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.EquipmentType{
					Name:           "Name",
					EquipmentItems: []*beerprotov1.EquipmentItemType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.EquipmentType{},
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
			got, err := ToJSONEquipmentType(context.Background(), tt.args.i)
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
		i *beerprotov1.EquipmentItemType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.EquipmentItemType{
					Name:                "Name",
					Notes:               "Notes",
					BoilRatePerHour:     &beerprotov1.VolumeType{},
					Type:                "Type",
					Form:                beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_MASH_TUN,
					DrainRatePerMinute:  &beerprotov1.VolumeType{},
					SpecificHeat:        &beerprotov1.SpecificHeatType{},
					GrainAbsorptionRate: &beerprotov1.SpecificVolumeType{},
					MaximumVolume:       &beerprotov1.VolumeType{},
					Weight:              &beerprotov1.MassType{},
					Loss:                &beerprotov1.VolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.EquipmentItemType{
					Name:          "Name",
					Type:          "Type",
					Form:          beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_MASH_TUN,
					MaximumVolume: &beerprotov1.VolumeType{},
					Loss:          &beerprotov1.VolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.EquipmentItemType{},
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
			got, err := ToJSONEquipmentItemType(context.Background(), tt.args.i)
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
		i *beerprotov1.SpecificHeatType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Calgc",
			args: args{
				i: &beerprotov1.SpecificHeatType{
					Value: 80,
					Unit:  beerprotov1.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_CALGC,
				},
			},
		},
		{
			name: "JKGK",
			args: args{
				i: &beerprotov1.SpecificHeatType{
					Value: 80,
					Unit:  beerprotov1.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_JKGK,
				},
			},
		},
		{
			name: "BTULBF",
			args: args{
				i: &beerprotov1.SpecificHeatType{
					Value: 80,
					Unit:  beerprotov1.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_BTULBF,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.SpecificHeatType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONSpecificHeatType(context.Background(), tt.args.i)
			have := ToProtoSpecificHeatType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONSpecificHeatType() %v", diff)
			}
		})
	}
}

func TestToJSONMassType(t *testing.T) {
	type args struct {
		i *beerprotov1.MassType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "G",
			args: args{
				i: &beerprotov1.MassType{
					Value: 90,
					Unit:  beerprotov1.MassUnit_MASS_UNIT_G,
				},
			},
		},
		{
			name: "KG",
			args: args{
				i: &beerprotov1.MassType{
					Value: 90,
					Unit:  beerprotov1.MassUnit_MASS_UNIT_KG,
				},
			},
		},
		{
			name: "LB",
			args: args{
				i: &beerprotov1.MassType{
					Value: 90,
					Unit:  beerprotov1.MassUnit_MASS_UNIT_LB,
				},
			},
		},
		{
			name: "MG",
			args: args{
				i: &beerprotov1.MassType{
					Value: 90,
					Unit:  beerprotov1.MassUnit_MASS_UNIT_MG,
				},
			},
		},
		{
			name: "OZ",
			args: args{
				i: &beerprotov1.MassType{
					Value: 90,
					Unit:  beerprotov1.MassUnit_MASS_UNIT_OZ,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.MassType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONMassType(context.Background(), tt.args.i)
			have := ToProtoMassType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONMassType() %v", diff)
			}
		})
	}
}

func TestToJSONVolumeType(t *testing.T) {
	type args struct {
		i *beerprotov1.VolumeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ML",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_ML,
				},
			},
		},
		{
			name: "L",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_L,
				},
			},
		},
		{
			name: "TSP",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_TSP,
				},
			},
		},
		{
			name: "TBSP",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_TBSP,
				},
			},
		},
		{
			name: "FLOZ",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_FLOZ,
				},
			},
		},
		{
			name: "CUP",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_CUP,
				},
			},
		},
		{
			name: "PT",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_PT,
				},
			},
		},
		{
			name: "QT",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_QT,
				},
			},
		},
		{
			name: "GAL",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_GAL,
				},
			},
		},
		{
			name: "BBL",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_BBL,
				},
			},
		},
		{
			name: "IFOZ",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_IFOZ,
				},
			},
		},
		{
			name: "IPT",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_IPT,
				},
			},
		},
		{
			name: "IQT",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_IQT,
				},
			},
		},
		{
			name: "IGAL",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_IGAL,
				},
			},
		},
		{
			name: "IBBL",
			args: args{
				i: &beerprotov1.VolumeType{
					Value: 50,
					Unit:  beerprotov1.VolumeUnit_VOLUME_UNIT_IBBL,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.VolumeType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONVolumeType(context.Background(), tt.args.i)
			have := ToProtoVolumeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONVolumeType() %v", diff)
			}
		})
	}
}

func TestToJSONSpecificVolumeType(t *testing.T) {
	type args struct {
		i *beerprotov1.SpecificVolumeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "qt/lb",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_QTLB,
				},
			},
		},
		{
			name: "gal/lb",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALLB,
				},
			},
		},
		{
			name: "gal/oz",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALOZ,
				},
			},
		},
		{
			name: "l/g",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LG,
				},
			},
		},
		{
			name: "l/kg",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LKG,
				},
			},
		},
		{
			name: "floz/oz",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FLOZOZ,
				},
			},
		},
		{
			name: "m^3/kg",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_M3KG,
				},
			},
		},
		{
			name: "ft^3/lb",
			args: args{
				i: &beerprotov1.SpecificVolumeType{
					Value: 30,
					Unit:  beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FT3LB,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.SpecificVolumeType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONSpecificVolumeType(context.Background(), tt.args.i)
			have := ToProtoSpecificVolumeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONSpecificVolumeType() %v", diff)
			}
		})
	}
}

func TestToJSONCultureInformation(t *testing.T) {
	type args struct {
		i *beerprotov1.CultureInformation
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.CultureInformation{
					Form:             beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_DRY,
					Producer:         "Producer",
					TemperatureRange: nil,
					Notes:            "Notes",
					BestFor:          "BestFor",
					Inventory:        &beerprotov1.CultureInventoryType{},
					ProductId:        "ProductId",
					Name:             "Name",
					AlcoholTolerance: &beerprotov1.PercentType{},
					Glucoamylase:     true,
					Type:             beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_ALE,
					Flocculation:     beerprotov1.QualitativeRangeUnit_QUALITATIVE_RANGE_UNIT_HIGH,
					AttenuationRange: nil,
					MaxReuse:         2,
					Pof:              true,
					Zymocide:         &beerprotov1.Zymocide{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.CultureInformation{
					Form: beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_DRY,
					Name: "Name",
					Type: beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_ALE,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.CultureInformation{},
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
			got, err := ToJSONCultureInformation(context.Background(), tt.args.i)
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
		i *beerprotov1.TemperatureRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.TemperatureRangeType{
					Maximum: &beerprotov1.TemperatureType{},
					Minimum: &beerprotov1.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.TemperatureRangeType{
					Minimum: &beerprotov1.TemperatureType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.TemperatureRangeType{
					Maximum: &beerprotov1.TemperatureType{},
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
			got, err := ToJSONTemperatureRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.TemperatureType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "C",
			args: args{
				i: &beerprotov1.TemperatureType{
					Value: 5,
					Unit:  beerprotov1.TemperatureUnit_TEMPERATURE_UNIT_C,
				},
			},
		},
		{
			name: "F",
			args: args{
				i: &beerprotov1.TemperatureType{
					Value: 5,
					Unit:  beerprotov1.TemperatureUnit_TEMPERATURE_UNIT_F,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.TemperatureType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONTemperatureType(context.Background(), tt.args.i)
			have := ToProtoTemperatureType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTemperatureType() %v", diff)
			}
		})
	}
}

func TestToJSONPercentRangeType(t *testing.T) {
	type args struct {
		i *beerprotov1.PercentRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.PercentRangeType{
					Maximum: &beerprotov1.PercentType{},
					Minimum: &beerprotov1.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.PercentRangeType{
					Minimum: &beerprotov1.PercentType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.PercentRangeType{
					Maximum: &beerprotov1.PercentType{},
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
			got, err := ToJSONPercentRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.PercentType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Percent",
			args: args{
				i: &beerprotov1.PercentType{
					Value: 75,
					Unit:  beerprotov1.PercentUnit_PERCENT_UNIT_PERCENT_SIGN,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.PercentType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONPercentType(context.Background(), tt.args.i)
			have := ToProtoPercentType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONPercentType() %v", diff)
			}
		})
	}
}

func TestToJSONZymocide(t *testing.T) {
	type args struct {
		i *beerprotov1.Zymocide
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.Zymocide{
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
				i: &beerprotov1.Zymocide{},
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
			got := ToJSONZymocide(context.Background(), tt.args.i)
			have := ToProtoZymocide(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONZymocide() %v", diff)
			}
		})
	}
}

func TestToJSONCultureInventoryType(t *testing.T) {
	type args struct {
		i *beerprotov1.CultureInventoryType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.CultureInventoryType{
					Liquid:  &beerprotov1.VolumeType{},
					Dry:     &beerprotov1.MassType{},
					Slant:   &beerprotov1.VolumeType{},
					Culture: &beerprotov1.VolumeType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.CultureInventoryType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONCultureInventoryType(context.Background(), tt.args.i)
			have := ToProtoCultureInventoryType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCultureInventoryType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentableType(t *testing.T) {
	type args struct {
		i *beerprotov1.FermentableType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.FermentableType{
					MaxInBatch:     &beerprotov1.PercentType{},
					RecommendMash:  true,
					Protein:        &beerprotov1.PercentType{},
					ProductId:      "ProductId",
					GrainGroup:     beerprotov1.GrainGroup_GRAIN_GROUP_BASE,
					Yield:          &beerprotov1.YieldType{},
					Type:           beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_GRAIN,
					Producer:       "Producer",
					AlphaAmylase:   80,
					Color:          &beerprotov1.ColorType{},
					Name:           "Name",
					DiastaticPower: &beerprotov1.DiastaticPowerType{},
					Moisture:       &beerprotov1.PercentType{},
					Origin:         "Origin",
					Inventory:      &beerprotov1.FermentableInventoryType{},
					KolbachIndex:   &beerprotov1.PercentType{Value: 80},
					Glassy:         &beerprotov1.PercentType{},
					Plump:          &beerprotov1.PercentType{},
					Half:           &beerprotov1.PercentType{},
					Mealy:          &beerprotov1.PercentType{},
					Thru:           &beerprotov1.PercentType{},
					Friability:     &beerprotov1.PercentType{},
					DiPh:           &beerprotov1.AcidityType{},
					Viscosity:      &beerprotov1.ViscosityType{},
					DmsP:           &beerprotov1.ConcentrationType{},
					Fan:            &beerprotov1.ConcentrationType{},
					Fermentability: &beerprotov1.PercentType{},
					BetaGlucan:     &beerprotov1.ConcentrationType{},
					Notes:          "Notes",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.FermentableType{
					Yield: &beerprotov1.YieldType{},
					Type:  beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_GRAIN,
					Color: &beerprotov1.ColorType{},
					Name:  "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.FermentableType{},
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
			got, err := ToJSONFermentableType(context.Background(), tt.args.i)
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
		i *beerprotov1.ColorType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "EBC",
			args: args{
				i: &beerprotov1.ColorType{
					Value: 10,
					Unit:  beerprotov1.ColorUnit_COLOR_UNIT_EBC,
				},
			},
		},
		{
			name: "Lovi",
			args: args{
				i: &beerprotov1.ColorType{
					Value: 10,
					Unit:  beerprotov1.ColorUnit_COLOR_UNIT_LOVI,
				},
			},
		},
		{
			name: "SRM",
			args: args{
				i: &beerprotov1.ColorType{
					Value: 10,
					Unit:  beerprotov1.ColorUnit_COLOR_UNIT_SRM,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.ColorType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONColorType(context.Background(), tt.args.i)
			have := ToProtoColorType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONColorType() %v", diff)
			}
		})
	}
}

func TestToJSONAcidityType(t *testing.T) {
	type args struct {
		i *beerprotov1.AcidityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pH",
			args: args{
				i: &beerprotov1.AcidityType{
					Value: 10,
					Unit:  beerprotov1.AcidityUnit_ACIDITY_UNIT_PH,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.AcidityType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONAcidityType(context.Background(), tt.args.i)
			have := ToProtoAcidityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONAcidityType() %v", diff)
			}
		})
	}
}

func TestToJSONViscosityType(t *testing.T) {
	type args struct {
		i *beerprotov1.ViscosityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mPa-s",
			args: args{
				i: &beerprotov1.ViscosityType{
					Value: 10,
					Unit:  beerprotov1.ViscosityUnit_VISCOSITY_UNIT_MPAS,
				},
			},
		},
		{
			name: "cP",
			args: args{
				i: &beerprotov1.ViscosityType{
					Value: 10,
					Unit:  beerprotov1.ViscosityUnit_VISCOSITY_UNIT_CP,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.ViscosityType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONViscosityType(context.Background(), tt.args.i)
			have := ToProtoViscosityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONViscosityType() %v", diff)
			}
		})
	}
}

func TestToJSONStyleType(t *testing.T) {
	type args struct {
		i *beerprotov1.StyleType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.StyleType{
					Aroma:          "Aroma",
					Ingredients:    "Ingredients",
					CategoryNumber: 10,
					Notes:          "Notes",
					Flavor:         "Flavor",
					Mouthfeel:      "Mouthfeel",
					FinalGravity: &beerprotov1.GravityRangeType{
						Maximum: &beerprotov1.GravityType{},
						Minimum: &beerprotov1.GravityType{},
					},
					StyleGuide: "Style",
					Color: &beerprotov1.ColorRangeType{
						Minimum: &beerprotov1.ColorType{},
						Maximum: &beerprotov1.ColorType{},
					},
					OriginalGravity: &beerprotov1.GravityRangeType{
						Maximum: &beerprotov1.GravityType{},
						Minimum: &beerprotov1.GravityType{},
					},
					Examples: "Example",
					Name:     "Name",
					Carbonation: &beerprotov1.CarbonationRangeType{
						Maximum: &beerprotov1.CarbonationType{},
						Minimum: &beerprotov1.CarbonationType{},
					},
					AlcoholByVolume: &beerprotov1.PercentRangeType{
						Minimum: &beerprotov1.PercentType{},
						Maximum: &beerprotov1.PercentType{},
					},
					InternationalBitternessUnits: &beerprotov1.BitternessRangeType{
						Minimum: &beerprotov1.BitternessType{},
						Maximum: &beerprotov1.BitternessType{},
					},
					Appearance:        "Appearance",
					Category:          "Category",
					StyleLetter:       "StyleLetter",
					Type:              beerprotov1.StyleType_STYLE_CATEGORIES_BEER,
					OverallImpression: "Overall Impression",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.StyleType{
					Name:       "Name",
					StyleGuide: "StyleGuide",
					Category:   "Category",
					Type:       beerprotov1.StyleType_STYLE_CATEGORIES_CIDER,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.StyleType{},
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
			got, err := ToJSONStyleType(context.Background(), tt.args.i)
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
		i *beerprotov1.GravityRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.GravityRangeType{
					Minimum: &beerprotov1.GravityType{},
					Maximum: &beerprotov1.GravityType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.GravityRangeType{
					Minimum: &beerprotov1.GravityType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.GravityRangeType{
					Maximum: &beerprotov1.GravityType{},
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
			got, err := ToJSONGravityRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.GravityType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SG",
			args: args{
				i: &beerprotov1.GravityType{
					Value: 45,
					Unit:  beerprotov1.GravityUnit_GRAVITY_UNIT_SG,
				},
			},
		},
		{
			name: "BRIX",
			args: args{
				i: &beerprotov1.GravityType{
					Value: 45,
					Unit:  beerprotov1.GravityUnit_GRAVITY_UNIT_BRIX,
				},
			},
		},
		{
			name: "Plato",
			args: args{
				i: &beerprotov1.GravityType{
					Value: 45,
					Unit:  beerprotov1.GravityUnit_GRAVITY_UNIT_PLATO,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.GravityType{},
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
			got := ToJSONGravityType(context.Background(), tt.args.i)
			have := ToProtoGravityType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONGravityType() %v", diff)
			}
		})
	}
}

func TestToJSONColorRangeType(t *testing.T) {
	type args struct {
		i *beerprotov1.ColorRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.ColorRangeType{
					Minimum: &beerprotov1.ColorType{},
					Maximum: &beerprotov1.ColorType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.ColorRangeType{
					Minimum: &beerprotov1.ColorType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.ColorRangeType{
					Maximum: &beerprotov1.ColorType{},
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
			got, err := ToJSONColorRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.CarbonationRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.CarbonationRangeType{
					Minimum: &beerprotov1.CarbonationType{},
					Maximum: &beerprotov1.CarbonationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.CarbonationRangeType{
					Minimum: &beerprotov1.CarbonationType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.CarbonationRangeType{
					Maximum: &beerprotov1.CarbonationType{},
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
			got, err := ToJSONCarbonationRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.CarbonationType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Vols",
			args: args{
				i: &beerprotov1.CarbonationType{
					Value: 45,
					Unit:  beerprotov1.CarbonationUnit_CARBONATION_UNIT_VOLS,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.CarbonationType{},
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
			got := ToJSONCarbonationType(context.Background(), tt.args.i)
			have := ToProtoCarbonationType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONCarbonationType() %v", diff)
			}
		})
	}
}

func TestToJSONBitternessRangeType(t *testing.T) {
	type args struct {
		i *beerprotov1.BitternessRangeType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.BitternessRangeType{
					Minimum: &beerprotov1.BitternessType{},
					Maximum: &beerprotov1.BitternessType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Maximum",
			args: args{
				i: &beerprotov1.BitternessRangeType{
					Minimum: &beerprotov1.BitternessType{},
				},
			},
			wantErr: true,
		},
		{
			name: "Minimum",
			args: args{
				i: &beerprotov1.BitternessRangeType{
					Maximum: &beerprotov1.BitternessType{},
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
			got, err := ToJSONBitternessRangeType(context.Background(), tt.args.i)
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
		i *beerprotov1.BitternessType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "IBU",
			args: args{
				i: &beerprotov1.BitternessType{
					Value: 35,
					Unit:  beerprotov1.BitternessUnit_BITTERNESS_UNIT_IBUS,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.BitternessType{},
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
			got := ToJSONBitternessType(context.Background(), tt.args.i)
			have := ToProtoBitternessType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONBitternessType() %v", diff)
			}
		})
	}
}

func TestToJSONMiscellaneousType(t *testing.T) {
	type args struct {
		i *beerprotov1.MiscellaneousType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.MiscellaneousType{
					UseFor:    "UseFor",
					Notes:     "Notes",
					Name:      "Name",
					Producer:  "Producer",
					ProductId: "ProductId",
					Type:      beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_FLAVOR,
					Inventory: &beerprotov1.MiscellaneousInventoryType{
						Amount: &beerprotov1.MiscellaneousInventoryType_Mass{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.MiscellaneousType{
					Name: "Name",
					Type: beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_FLAVOR,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.MiscellaneousType{},
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
			got, err := ToJSONMiscellaneousType(context.Background(), tt.args.i)
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
		i *beerprotov1.MiscellaneousInventoryType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.MiscellaneousInventoryType{
					Amount: &beerprotov1.MiscellaneousInventoryType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerprotov1.MiscellaneousInventoryType{
					Amount: &beerprotov1.MiscellaneousInventoryType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.MiscellaneousInventoryType{
					Amount: &beerprotov1.MiscellaneousInventoryType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.MiscellaneousInventoryType{},
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
			got, err := ToJSONMiscellaneousInventoryType(context.Background(), tt.args.i)
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
		i *beerprotov1.TasteType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.TasteType{
					Notes:  "Notes",
					Rating: 10,
				},
			},
			wantErr: false,
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
			got, err := ToJSONTasteType(context.Background(), tt.args.i)
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
		i *beerprotov1.BoilProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.BoilProcedureType{
					PreBoilSize: &beerprotov1.VolumeType{},
					BoilTime: &beerprotov1.TimeType{
						Value: 60,
						Unit:  beerprotov1.TimeUnit_TIME_UNIT_MIN,
					},
					BoilSteps:   []*beerprotov1.BoilStepType{},
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
				i: &beerprotov1.BoilProcedureType{
					PreBoilSize: &beerprotov1.VolumeType{},
					BoilTime:    nil,
					BoilSteps:   []*beerprotov1.BoilStepType{},
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
			got, err := ToJSONBoilProcedureType(context.Background(), tt.args.i)
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
		i *beerprotov1.BoilStepType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Batch",
			args: args{
				i: &beerprotov1.BoilStepType{
					EndGravity:       &beerprotov1.GravityType{},
					ChillingType:     beerprotov1.BoilStepTypeChillingType_BOIL_STEP_TYPE_CHILLING_TYPE_BATCH,
					Description:      "Description",
					EndTemperature:   &beerprotov1.TemperatureType{},
					RampTime:         &beerprotov1.TimeType{},
					StepTime:         &beerprotov1.TimeType{},
					StartGravity:     &beerprotov1.GravityType{},
					StartPh:          &beerprotov1.AcidityType{},
					EndPh:            &beerprotov1.AcidityType{},
					Name:             "Name",
					StartTemperature: &beerprotov1.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Inline",
			args: args{
				i: &beerprotov1.BoilStepType{
					EndGravity:       &beerprotov1.GravityType{},
					ChillingType:     beerprotov1.BoilStepTypeChillingType_BOIL_STEP_TYPE_CHILLING_TYPE_INLINE,
					Description:      "Description",
					EndTemperature:   &beerprotov1.TemperatureType{},
					RampTime:         &beerprotov1.TimeType{},
					StepTime:         &beerprotov1.TimeType{},
					StartGravity:     &beerprotov1.GravityType{},
					StartPh:          &beerprotov1.AcidityType{},
					EndPh:            &beerprotov1.AcidityType{},
					Name:             "Name",
					StartTemperature: &beerprotov1.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.BoilStepType{
					Name: "Name",
				},
			},
			wantErr: false,
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
			got, err := ToJSONBoilStepType(context.Background(), tt.args.i)
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
		i *beerprotov1.TimeType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sec",
			args: args{
				i: &beerprotov1.TimeType{
					Value: 10,
					Unit:  beerprotov1.TimeUnit_TIME_UNIT_SEC,
				},
			},
		},
		{
			name: "Min",
			args: args{
				i: &beerprotov1.TimeType{
					Value: 10,
					Unit:  beerprotov1.TimeUnit_TIME_UNIT_MIN,
				},
			},
		},
		{
			name: "Hour",
			args: args{
				i: &beerprotov1.TimeType{
					Value: 10,
					Unit:  beerprotov1.TimeUnit_TIME_UNIT_HR,
				},
			},
		},
		{
			name: "Day",
			args: args{
				i: &beerprotov1.TimeType{
					Value: 10,
					Unit:  beerprotov1.TimeUnit_TIME_UNIT_DAY,
				},
			},
		},
		{
			name: "Week",
			args: args{
				i: &beerprotov1.TimeType{
					Value: 10,
					Unit:  beerprotov1.TimeUnit_TIME_UNIT_WEEK,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.TimeType{},
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
			got := ToJSONTimeType(context.Background(), tt.args.i)
			have := ToProtoTimeType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTimeType() %v", diff)
			}
		})
	}
}

func TestToJSONPackagingProcedureType(t *testing.T) {
	type args struct {
		i *beerprotov1.PackagingProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.PackagingProcedureType{
					Name:             "Name",
					PackagedVolume:   &beerprotov1.VolumeType{},
					Description:      "Description",
					Notes:            "Notes",
					PackagingVessels: []*beerprotov1.PackagingVesselType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.PackagingProcedureType{
					Name: "Name",
				},
			},
			wantErr: false,
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
			got, err := ToJSONPackagingProcedureType(context.Background(), tt.args.i)
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
		i *beerprotov1.PackagingVesselType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.PackagingVesselType{
					Type:             beerprotov1.PackagingVesselUnit_PACKAGING_VESSEL_UNIT_BOTTLE,
					StartGravity:     &beerprotov1.GravityType{},
					Name:             "Name",
					PackageDate:      "PackageDate",
					StepTime:         &beerprotov1.TimeType{},
					EndGravity:       &beerprotov1.GravityType{},
					VesselVolume:     &beerprotov1.VolumeType{},
					VesselQuantity:   10,
					Description:      "Description",
					StartPh:          &beerprotov1.AcidityType{},
					Carbonation:      80,
					StartTemperature: &beerprotov1.TemperatureType{},
					EndPh:            &beerprotov1.AcidityType{},
					EndTemperature:   &beerprotov1.TemperatureType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.PackagingVesselType{
					Name: "Name",
				},
			},
			wantErr: false,
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
			got, err := ToJSONPackagingVesselType(context.Background(), tt.args.i)
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
		i *beerprotov1.HopAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.HopAdditionType{
					BetaAcid:  &beerprotov1.PercentType{},
					Producer:  "Producer",
					Origin:    "Origin",
					Year:      "Year",
					Form:      beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PELLET,
					Timing:    &beerprotov1.TimingType{},
					Name:      "Name",
					ProductId: "ProductId",
					AlphaAcid: &beerprotov1.PercentType{},
					Amount:    &beerprotov1.HopAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.HopAdditionType{
					BetaAcid:  &beerprotov1.PercentType{},
					Producer:  "Producer",
					Origin:    "Origin",
					Year:      "Year",
					Form:      beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_POWDER,
					Timing:    &beerprotov1.TimingType{},
					Name:      "Name",
					ProductId: "ProductId",
					AlphaAcid: &beerprotov1.PercentType{},
					Amount:    &beerprotov1.HopAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.HopAdditionType{},
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
			got, err := ToJSONHopAdditionType(context.Background(), tt.args.i)
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
		i *beerprotov1.FermentableAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.FermentableAdditionType{
					Type:       beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_GRAIN,
					Origin:     "Origin",
					GrainGroup: beerprotov1.GrainGroup_GRAIN_GROUP_BASE,
					Yield:      &beerprotov1.YieldType{},
					Color:      &beerprotov1.ColorType{},
					Name:       "Name",
					Producer:   "Producer",
					ProductId:  "ProductId",
					Timing:     &beerprotov1.TimingType{},
					Amount:     &beerprotov1.FermentableAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.FermentableAdditionType{
					Type:       beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_DRY_EXTRACT,
					Origin:     "Origin",
					GrainGroup: beerprotov1.GrainGroup_GRAIN_GROUP_CARAMEL,
					Yield:      &beerprotov1.YieldType{},
					Color:      &beerprotov1.ColorType{},
					Name:       "Name",
					Producer:   "Producer",
					ProductId:  "ProductId",
					Timing:     &beerprotov1.TimingType{},
					Amount:     &beerprotov1.FermentableAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.FermentableAdditionType{
					Type:   beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_DRY_EXTRACT,
					Yield:  &beerprotov1.YieldType{},
					Color:  &beerprotov1.ColorType{},
					Name:   "Name",
					Timing: &beerprotov1.TimingType{},
					Amount: &beerprotov1.FermentableAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.FermentableAdditionType{},
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
			got, err := ToJSONFermentableAdditionType(context.Background(), tt.args.i)
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
		i *beerprotov1.YieldType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.YieldType{
					FineGrind:            &beerprotov1.PercentType{},
					CoarseGrind:          &beerprotov1.PercentType{},
					FineCoarseDifference: &beerprotov1.PercentType{},
					Potential:            &beerprotov1.GravityType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.YieldType{},
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
			got := ToJSONYieldType(context.Background(), tt.args.i)
			have := ToProtoYieldType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONYieldType() %v", diff)
			}
		})
	}
}

func TestToJSONWaterAdditionType(t *testing.T) {
	type args struct {
		i *beerprotov1.WaterAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.WaterAdditionType{
					Flouride:    &beerprotov1.ConcentrationType{},
					Sulfate:     &beerprotov1.ConcentrationType{},
					Producer:    "Producer",
					Nitrate:     &beerprotov1.ConcentrationType{},
					Nitrite:     &beerprotov1.ConcentrationType{},
					Chloride:    &beerprotov1.ConcentrationType{},
					Amount:      &beerprotov1.VolumeType{},
					Name:        "Name",
					Potassium:   &beerprotov1.ConcentrationType{},
					Magnesium:   &beerprotov1.ConcentrationType{},
					Iron:        &beerprotov1.ConcentrationType{},
					Bicarbonate: &beerprotov1.ConcentrationType{},
					Calcium:     &beerprotov1.ConcentrationType{},
					Carbonate:   &beerprotov1.ConcentrationType{},
					Sodium:      &beerprotov1.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.WaterAdditionType{
					Sulfate:     &beerprotov1.ConcentrationType{},
					Chloride:    &beerprotov1.ConcentrationType{},
					Name:        "Name",
					Magnesium:   &beerprotov1.ConcentrationType{},
					Bicarbonate: &beerprotov1.ConcentrationType{},
					Calcium:     &beerprotov1.ConcentrationType{},
					Sodium:      &beerprotov1.ConcentrationType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.WaterAdditionType{},
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
			got, err := ToJSONWaterAdditionType(context.Background(), tt.args.i)
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
		i *beerprotov1.CultureAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.CultureAdditionType{
					Form:              beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_DRY,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_ALE,
					Attenuation:       &beerprotov1.PercentType{},
					Timing:            &beerprotov1.TimingType{},
					Amount:            &beerprotov1.CultureAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerprotov1.CultureAdditionType{
					Form:              beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_CULTURE,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_LAGER,
					Attenuation:       &beerprotov1.PercentType{},
					Timing:            &beerprotov1.TimingType{},
					Amount:            &beerprotov1.CultureAdditionType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.CultureAdditionType{
					Form:              beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_LIQUID,
					ProductId:         "ProductId",
					Name:              "Name",
					CellCountBillions: 800,
					TimesCultured:     4,
					Producer:          "Producer",
					Type:              beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_BACTERIA,
					Attenuation:       &beerprotov1.PercentType{},
					Timing:            &beerprotov1.TimingType{},
					Amount:            &beerprotov1.CultureAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.CultureAdditionType{
					Form: beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_LIQUID,
					Name: "Name",
					Type: beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_BACTERIA,
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.CultureAdditionType{},
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
			got, err := ToJSONCultureAdditionType(context.Background(), tt.args.i)
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
		i *beerprotov1.MiscellaneousAdditionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Mass",
			args: args{
				i: &beerprotov1.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerprotov1.TimingType{},
					ProductId: "ProductId",
					Type:      beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_FLAVOR,
					Amount:    &beerprotov1.MiscellaneousAdditionType_Mass{},
				},
			},
			wantErr: false,
		},
		{
			name: "Unit",
			args: args{
				i: &beerprotov1.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerprotov1.TimingType{},
					ProductId: "ProductId",
					Type:      beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_FINING,
					Amount:    &beerprotov1.MiscellaneousAdditionType_Unit{},
				},
			},
			wantErr: false,
		},
		{
			name: "Volume",
			args: args{
				i: &beerprotov1.MiscellaneousAdditionType{
					Name:      "Name",
					Producer:  "Producer",
					Timing:    &beerprotov1.TimingType{},
					ProductId: "ProductId",
					Type:      beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_HERB,
					Amount:    &beerprotov1.MiscellaneousAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.MiscellaneousAdditionType{
					Name:   "Name",
					Type:   beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_SPICE,
					Timing: &beerprotov1.TimingType{},
					Amount: &beerprotov1.MiscellaneousAdditionType_Volume{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.MiscellaneousAdditionType{},
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
			got, err := ToJSONMiscellaneousAdditionType(context.Background(), tt.args.i)
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
		i *beerprotov1.UnitType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "one",
			args: args{
				i: &beerprotov1.UnitType{
					Value: 60,
					Unit:  beerprotov1.UnitUnit_UNIT_UNIT_ONE,
				},
			},
		},
		{
			name: "unit",
			args: args{
				i: &beerprotov1.UnitType{
					Value: 60,
					Unit:  beerprotov1.UnitUnit_UNIT_UNIT_UNIT,
				},
			},
		},
		{
			name: "each",
			args: args{
				i: &beerprotov1.UnitType{
					Value: 60,
					Unit:  beerprotov1.UnitUnit_UNIT_UNIT_EACH,
				},
			},
		},
		{
			name: "dimensionless",
			args: args{
				i: &beerprotov1.UnitType{
					Value: 60,
					Unit:  beerprotov1.UnitUnit_UNIT_UNIT_DIMENSIONLESS,
				},
			},
		},
		{
			name: "pkg",
			args: args{
				i: &beerprotov1.UnitType{
					Value: 60,
					Unit:  beerprotov1.UnitUnit_UNIT_UNIT_PKG,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.UnitType{},
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
			got := ToJSONUnitType(context.Background(), tt.args.i)
			have := ToProtoUnitType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONUnitType() %v", diff)
			}
		})
	}
}

func TestToJSONTimingType(t *testing.T) {
	type args struct {
		i *beerprotov1.TimingType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.TimingType{
					Time:            &beerprotov1.TimeType{},
					Duration:        &beerprotov1.TimeType{},
					Continuous:      true,
					SpecificGravity: &beerprotov1.GravityType{},
					Ph:              &beerprotov1.AcidityType{},
					Step:            4,
					Use:             1,
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.TimingType{},
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
			got := ToJSONTimingType(context.Background(), tt.args.i)
			have := ToProtoTimingType(got)
			if diff := deep.Equal(have, tt.args.i); diff != nil {
				t.Errorf("ToJSONTimingType() %v", diff)
			}
		})
	}
}

func TestToJSONFermentationProcedureType(t *testing.T) {
	type args struct {
		i *beerprotov1.FermentationProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.FermentationProcedureType{
					Description:       "Description",
					Notes:             "Notes",
					FermentationSteps: []*beerprotov1.FermentationStepType{},
					Name:              "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.FermentationProcedureType{
					Name:              "Name",
					FermentationSteps: []*beerprotov1.FermentationStepType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.FermentationProcedureType{},
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
			got, err := ToJSONFermentationProcedureType(context.Background(), tt.args.i)
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
		i *beerprotov1.FermentationStepType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.FermentationStepType{
					Name:             "Name",
					EndTemperature:   &beerprotov1.TemperatureType{},
					StepTime:         &beerprotov1.TimeType{},
					FreeRise:         true,
					StartGravity:     &beerprotov1.GravityType{},
					EndGravity:       &beerprotov1.GravityType{},
					StartPh:          &beerprotov1.AcidityType{},
					Description:      "Description",
					StartTemperature: &beerprotov1.TemperatureType{},
					EndPh:            &beerprotov1.AcidityType{},
					Vessel:           "Vessel",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.FermentationStepType{
					Name: "Name",
				},
			},
			wantErr: false,
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
			got, err := ToJSONFermentationStepType(context.Background(), tt.args.i)
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
		i *beerprotov1.RecipeStyleType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.RecipeStyleType{
					Type:           beerprotov1.RecipeStyleType_STYLE_CATEGORIES_BEER,
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
				i: &beerprotov1.RecipeStyleType{
					Type:       beerprotov1.RecipeStyleType_STYLE_CATEGORIES_BEER,
					Name:       "Name",
					Category:   "Category",
					StyleGuide: "StyleGuide",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.RecipeStyleType{},
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
			got, err := ToJSONRecipeStyleType(context.Background(), tt.args.i)
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
		i *beerprotov1.EfficiencyType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.EfficiencyType{
					Conversion: &beerprotov1.PercentType{},
					Lauter:     &beerprotov1.PercentType{},
					Mash:       &beerprotov1.PercentType{},
					Brewhouse:  &beerprotov1.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.EfficiencyType{
					Brewhouse: &beerprotov1.PercentType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.EfficiencyType{},
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
			got, err := ToJSONEfficiencyType(context.Background(), tt.args.i)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSONEfficiencyType() error = %v, wantErr %v", err, tt.wantErr)
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
		i *beerprotov1.MashProcedureType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.MashProcedureType{
					GrainTemperature: &beerprotov1.TemperatureType{},
					Notes:            "Notes",
					MashSteps:        []*beerprotov1.MashStepType{},
					Name:             "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.MashProcedureType{
					GrainTemperature: &beerprotov1.TemperatureType{},
					MashSteps:        []*beerprotov1.MashStepType{},
					Name:             "Name",
				},
			},
			wantErr: false,
		},
		{
			name: "Nil",
			args: args{
				i: &beerprotov1.MashProcedureType{},
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
			got, err := ToJSONMashProcedureType(context.Background(), tt.args.i)
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
		i *beerprotov1.MashStepType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				i: &beerprotov1.MashStepType{
					StepTime:          &beerprotov1.TimeType{},
					RampTime:          &beerprotov1.TimeType{},
					EndTemperature:    &beerprotov1.TemperatureType{},
					Description:       "Description",
					InfuseTemperature: &beerprotov1.TemperatureType{},
					StartPh:           &beerprotov1.AcidityType{},
					EndPh:             &beerprotov1.AcidityType{},
					Name:              "Name",
					Type:              beerprotov1.MashStepType_MASH_STEP_UNIT_DRAIN_MASH_TUN,
					Amount:            &beerprotov1.VolumeType{},
					StepTemperature:   &beerprotov1.TemperatureType{},
					WaterGrainRatio:   &beerprotov1.SpecificVolumeType{},
				},
			},
			wantErr: false,
		},
		{
			name: "Required",
			args: args{
				i: &beerprotov1.MashStepType{
					StepTime:        &beerprotov1.TimeType{},
					Name:            "Name",
					StepTemperature: &beerprotov1.TemperatureType{},
					Type:            beerprotov1.MashStepType_MASH_STEP_UNIT_DECOCTION,
				},
			},
			wantErr: false,
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
			got, err := ToJSONMashStepType(context.Background(), tt.args.i)
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
