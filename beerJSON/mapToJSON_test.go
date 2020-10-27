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

					},
					BatchSize: &beerproto.VolumeType{
						Unit:  beerproto.VolumeType_L,
						Value: float64(30),
					},
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
		name string
		args args
	}{
		{
			name: "Nil fields",
			args: args{
				i: &beerproto.IngredientsType{

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONIngredientsType(tt.args.i)
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
	}{
		{
			name: "",
			args: args{
				i: &beerproto.WaterBase{
					Name: "Balanced Profile",
					Calcium: &beerproto.ConcentrationType{
						Value: 80,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Magnesium: &beerproto.ConcentrationType{
						Value: 5,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Sulfate: &beerproto.ConcentrationType{
						Value: 80,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Sodium: &beerproto.ConcentrationType{
						Value: 25,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Chloride: &beerproto.ConcentrationType{
						Value: 75,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
					Bicarbonate: &beerproto.ConcentrationType{
						Value: 100,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONWaterBase(tt.args.i)
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
			name: "",
			args: args{
				i: &beerproto.Recipe{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToJSON(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToJSON() error = %v, wantErr %v", err, tt.wantErr)
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
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.VarietyInformation{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONVarietyInformation(tt.args.i)
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
			name: "Basic",
			args: args{
				i: &beerproto.HopInventoryType{
					Amount: &beerproto.HopInventoryType_Volume{},
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
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.EquipmentType{
					Name:           "Name",
					EquipmentItems: []*beerproto.EquipmentItemType{},
				},
			},
		},
		{
			name: "Nil",
			args: args{
				i: &beerproto.EquipmentType{
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONEquipmentType(tt.args.i)
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONEquipmentItemType(tt.args.i)
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
					Unit: beerproto.SpecificVolumeType_QTLB,
				},
			},
		},
		{
			name: "gal/lb",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_GALLB,
				},
			},
		},
		{
			name: "gal/oz",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_GALOZ,
				},
			},
		},
		{
			name: "l/g",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_LG,
				},
			},
		},
		{
			name: "l/kg",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_LKG,
				},
			},
		},
		{
			name: "floz/oz",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_FLOZOZ,
				},
			},
		},
		{
			name: "m^3/kg",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_M3KG,
				},
			},
		},
		{
			name: "ft^3/lb",
			args: args{
				i: &beerproto.SpecificVolumeType{
					Value: 30,
					Unit: beerproto.SpecificVolumeType_FT3LB,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToJSONCultureInformation(tt.args.i)
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
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.TemperatureRangeType{
					Maximum: &beerproto.TemperatureType{},
					Minimum: &beerproto.TemperatureType{},
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
			got := ToJSONTemperatureRangeType(tt.args.i)
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
					Unit: beerproto.TemperatureUnitType_C,
				},
			},
		},
		{
			name: "F",
			args: args{
				i: &beerproto.TemperatureType{
					Value: 5,
					Unit: beerproto.TemperatureUnitType_F,
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
	}{
		{
			name: "Basic",
			args: args{
				i: &beerproto.PercentRangeType{
					Maximum: &beerproto.PercentType{},
					Minimum: &beerproto.PercentType{},
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

			got := ToJSONPercentRangeType(tt.args.i)
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
					Unit: beerproto.PercentType_PERCENT_SIGN,
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
					No1: true,
					No2: true,
					No28: true,
					Klus: true,
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
						Liquid: &beerproto.VolumeType{},
						Dry: &beerproto.MassType{},
						Slant: &beerproto.VolumeType{},
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToJSONFermentableType(tt.args.i)
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
					Unit: beerproto.ColorUnitType_EBC,
				},
			},
		},
		{
			name: "Lovi",
			args: args{
				i: &beerproto.ColorType{
					Value: 10,
					Unit: beerproto.ColorUnitType_LOVI,
				},
			},
		},
		{
			name: "SRM",
			args: args{
				i: &beerproto.ColorType{
					Value: 10,
					Unit: beerproto.ColorUnitType_SRM,
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
					Unit: beerproto.AcidityUnitType_PH,
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
					Unit: beerproto.ViscosityUnitType_MPAS,
				},
			},
		},
		{
			name: "cP",
			args: args{
				i: &beerproto.ViscosityType{
					Value: 10,
					Unit: beerproto.ViscosityUnitType_CP,
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

