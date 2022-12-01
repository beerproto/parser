package beerJSON

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/beerproto/parser"

	"github.com/beerproto/beerjson.go"
	beerprotov1 "github.com/beerproto/beerproto_go"
)

func MapToJSON(context context.Context, i *beerprotov1.Recipe) (*beerjson.Beerjson, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	output := &beerjson.Beerjson{}

	stack.Push(i)

	if i.Version == 0 {
		stack.AppendError(fmt.Errorf("version is required"))
	}

	output.Mashes = []beerjson.MashProcedureType{}
	output.Recipes = []beerjson.RecipeType{}
	output.MiscellaneousIngredients = []beerjson.MiscellaneousType{}
	output.Styles = []beerjson.StyleType{}
	output.Fermentations = []beerjson.FermentationProcedureType{}
	output.Boil = []beerjson.BoilProcedureType{}
	output.Version = beerjson.VersionType(i.Version)
	output.Fermentables = []beerjson.FermentableType{}
	output.Cultures = []beerjson.CultureInformation{}
	output.Equipments = []beerjson.EquipmentType{}
	output.Packaging = []beerjson.PackagingProcedureType{}
	output.HopVarieties = []beerjson.VarietyInformation{}
	output.Profiles = []beerjson.WaterBase{}

	field, _ := reflect.TypeOf(i).Elem().FieldByName("Mashes")
	for index, mash := range i.Mashes {
		stack.PushMethod(index, field)
		if mash, err := ToJSONMashProcedureType(context, mash); err == nil {
			output.Mashes = append(output.Mashes, *mash)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Recipes")
	for index, recipe := range i.Recipes {
		stack.PushMethod(index, field)
		if r, err := ToJSONRecipeType(context, recipe); err == nil {
			output.Recipes = append(output.Recipes, *r)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("MiscellaneousIngredients")
	for index, ingredients := range i.MiscellaneousIngredients {
		stack.PushMethod(index, field)
		if mice, err := ToJSONMiscellaneousType(context, ingredients); err == nil {
			output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, *mice)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Styles")
	for index, style := range i.Styles {
		stack.PushMethod(index, field)
		if style, err := ToJSONStyleType(context, style); err == nil {
			output.Styles = append(output.Styles, *style)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Fermentations")
	for index, fermentation := range i.Fermentations {
		stack.PushMethod(index, field)
		if fermentationProcedureType, err := ToJSONFermentationProcedureType(context, fermentation); err == nil {
			output.Fermentations = append(output.Fermentations, *fermentationProcedureType)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Boil")
	for index, boil := range i.Boil {
		stack.PushMethod(index, field)
		if boil, err := ToJSONBoilProcedureType(context, boil); err == nil {
			output.Boil = append(output.Boil, *boil)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Fermentables")
	for index, fermentable := range i.Fermentables {
		stack.PushMethod(index, field)
		if fermentableType, err := ToJSONFermentableType(context, fermentable); err == nil {
			output.Fermentables = append(output.Fermentables, *fermentableType)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Cultures")
	for index, culture := range i.Cultures {
		stack.PushMethod(index, field)
		if culture, err := ToJSONCultureInformation(context, culture); err == nil {
			output.Cultures = append(output.Cultures, *culture)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Equipments")
	for index, equipment := range i.Equipments {
		stack.PushMethod(index, field)
		if equipmentType, err := ToJSONEquipmentType(context, equipment); err == nil {
			output.Equipments = append(output.Equipments, *equipmentType)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Packaging")
	for index, packaging := range i.Packaging {
		stack.PushMethod(index, field)
		if packagingProcedure, err := ToJSONPackagingProcedureType(context, packaging); err == nil {
			output.Packaging = append(output.Packaging, *packagingProcedure)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Packaging")
	for index, hopVariety := range i.HopVarieties {
		stack.PushMethod(index, field)
		if hop, err := ToJSONVarietyInformation(context, hopVariety); err == nil {
			output.HopVarieties = append(output.HopVarieties, *hop)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Packaging")
	for index, profile := range i.Profiles {
		stack.PushMethod(index, field)
		if water, err := ToJSONWaterBase(context, profile); err == nil {
			output.Profiles = append(output.Profiles, *water)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	return output, stack.Errors()
}

func ToJSONWaterBase(context context.Context, i *beerprotov1.WaterBase) (*beerjson.WaterBase, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	waterBase := &beerjson.WaterBase{}
	if i.Calcium == nil {
		stack.AppendError(fmt.Errorf("calcium is required"))
	} else {
		waterBase.Calcium = *ToJSONConcentrationType(context, i.Calcium)
	}

	if i.Bicarbonate == nil {
		stack.AppendError(fmt.Errorf("bicarbonate is required"))
	} else {
		waterBase.Bicarbonate = *ToJSONConcentrationType(context, i.Bicarbonate)
	}

	if i.Sulfate == nil {
		stack.AppendError(fmt.Errorf("sulfate is required"))
	} else {
		waterBase.Sulfate = *ToJSONConcentrationType(context, i.Sulfate)
	}

	if i.Chloride == nil {
		stack.AppendError(fmt.Errorf("chloride is required"))
	} else {
		waterBase.Chloride = *ToJSONConcentrationType(context, i.Chloride)
	}

	if i.Sodium == nil {
		stack.AppendError(fmt.Errorf("sodium is required"))
	} else {
		waterBase.Sodium = *ToJSONConcentrationType(context, i.Sodium)
	}

	if i.Magnesium == nil {
		stack.AppendError(fmt.Errorf("magnesium is required"))
	} else {
		waterBase.Magnesium = *ToJSONConcentrationType(context, i.Magnesium)
	}

	waterBase.Nitrite = ToJSONConcentrationType(context, i.Nitrite)
	waterBase.Name = i.Name
	waterBase.Potassium = ToJSONConcentrationType(context, i.Potassium)
	waterBase.Carbonate = ToJSONConcentrationType(context, i.Carbonate)
	waterBase.Iron = ToJSONConcentrationType(context, i.Iron)
	waterBase.Flouride = ToJSONConcentrationType(context, i.Flouride)
	waterBase.Producer = &i.Producer
	waterBase.Nitrate = ToJSONConcentrationType(context, i.Nitrate)

	return waterBase, stack.Errors()
}

func ToJSONVarietyInformation(context context.Context, i *beerprotov1.VarietyInformation) (*beerjson.VarietyInformation, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	if i.AlphaAcid == nil {
		stack.AppendError(fmt.Errorf("alphaAcid is required"))

	}

	return &beerjson.VarietyInformation{
		Inventory:              ToJSONHopInventoryType(context, i.Inventory),
		VarietyInformationType: ToJSONVarietyInformationType(context, i.Type),
		OilContent:             ToJSONOilContentType(context, i.OilContent),
		PercentLost:            ToJSONPercentType(context, i.PercentLost),
		ProductId:              &i.ProductId,
		AlphaAcid:              ToJSONPercentType(context, i.AlphaAcid),
		BetaAcid:               ToJSONPercentType(context, i.BetaAcid),
		Name:                   &i.Name,
		Origin:                 &i.Origin,
		Substitutes:            &i.Substitutes,
		Year:                   &i.Year,
		HopVarietyBaseForm:     ToJSONHopVarietyBaseForm(context, i.Form),
		Producer:               &i.Producer,
		Notes:                  &i.Notes,
	}, stack.Errors()
}

func ToJSONOilContentType(context context.Context, i *beerprotov1.OilContentType) *beerjson.OilContentType {
	if i == nil {
		return nil
	}

	return &beerjson.OilContentType{
		Polyphenols:       ToJSONPercentType(context, i.Polyphenols),
		TotalOilMlPer100g: &i.TotalOilMlPer_100G,
		Farnesene:         ToJSONPercentType(context, i.Farnesene),
		Limonene:          ToJSONPercentType(context, i.Limonene),
		Nerol:             ToJSONPercentType(context, i.Nerol),
		Geraniol:          ToJSONPercentType(context, i.Geraniol),
		BPinene:           ToJSONPercentType(context, i.BPinene),
		Linalool:          ToJSONPercentType(context, i.Linalool),
		Caryophyllene:     ToJSONPercentType(context, i.Caryophyllene),
		Cohumulone:        ToJSONPercentType(context, i.Cohumulone),
		Xanthohumol:       ToJSONPercentType(context, i.Xanthohumol),
		Humulene:          ToJSONPercentType(context, i.Humulene),
		Myrcene:           ToJSONPercentType(context, i.Myrcene),
		Pinene:            ToJSONPercentType(context, i.Pinene),
	}
}

func ToJSONVarietyInformationType(context context.Context, i beerprotov1.VarietyInformationType) *beerjson.VarietyInformationType {
	if i == beerprotov1.VarietyInformationType_VARIETY_INFORMATION_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.VarietyInformationType_name[int32(i)]
	t := beerjson.VarietyInformationType(strings.ToLower(unit))
	return &t
}

func ToJSONHopInventoryType(context context.Context, i *beerprotov1.HopInventoryType) *beerjson.HopInventoryType {
	if i == nil {
		return nil
	}

	hopInventoryType := &beerjson.HopInventoryType{}

	if mass, ok := i.Amount.(*beerprotov1.HopInventoryType_Mass); ok {
		hopInventoryType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if volume, ok := i.Amount.(*beerprotov1.HopInventoryType_Volume); ok {
		hopInventoryType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return hopInventoryType
}

func ToJSONEquipmentType(context context.Context, i *beerprotov1.EquipmentType) (*beerjson.EquipmentType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.EquipmentItems == nil {
		stack.AppendError(fmt.Errorf("equipmentItems is required"))
	}

	equipmentType := &beerjson.EquipmentType{
		Name: i.Name,
	}

	equipmentItemType := []beerjson.EquipmentItemType{}

	field, _ := reflect.TypeOf(i).Elem().FieldByName("EquipmentItems")
	for index, item := range i.EquipmentItems {
		stack.PushMethod(index, field)
		if equipmentType, err := ToJSONEquipmentItemType(context, item); err == nil {
			equipmentItemType = append(equipmentItemType, *equipmentType)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}

	equipmentType.EquipmentItems = equipmentItemType

	return equipmentType, stack.Errors()
}

func ToJSONEquipmentItemType(context context.Context, i *beerprotov1.EquipmentItemType) (*beerjson.EquipmentItemType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	equipmentItemType := &beerjson.EquipmentItemType{}

	if i.Form == beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("form is required"))
	}

	if i.MaximumVolume == nil {
		stack.AppendError(fmt.Errorf("maximumVolume is required"))
	}

	if i.Loss == nil {
		stack.AppendError(fmt.Errorf("loss is required"))
	} else {
		equipmentItemType.Loss = *ToJSONVolumeType(context, i.Loss)
	}

	equipmentItemType.BoilRatePerHour = ToJSONVolumeType(context, i.BoilRatePerHour)
	equipmentItemType.KeyType = &i.Type
	equipmentItemType.EquipmentBaseForm = ToJSONEquipmentBaseForm(context, i.Form)
	equipmentItemType.DrainRatePerMinute = ToJSONVolumeType(context, i.DrainRatePerMinute)
	equipmentItemType.SpecificHeat = ToJSONSpecificHeatType(context, i.SpecificHeat)
	equipmentItemType.GrainAbsorptionRate = ToJSONSpecificVolumeType(context, i.GrainAbsorptionRate)
	equipmentItemType.Name = &i.Name
	equipmentItemType.MaximumVolume = ToJSONVolumeType(context, i.MaximumVolume)
	equipmentItemType.Weight = ToJSONMassType(context, i.Weight)
	equipmentItemType.Notes = &i.Notes

	return equipmentItemType, stack.Errors()
}

func ToJSONSpecificHeatType(context context.Context, i *beerprotov1.SpecificHeatType) *beerjson.SpecificHeatType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_UNSPECIFIED {
		return &beerjson.SpecificHeatType{}
	}

	return &beerjson.SpecificHeatType{
		Value: i.Value,
		Unit:  *ToJSONSpecificHeatUnitType(context, i.Unit),
	}
}

func ToJSONSpecificHeatUnitType(context context.Context, i beerprotov1.SpecificHeatUnit) *beerjson.SpecificHeatUnitType {
	if i == beerprotov1.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.SpecificHeatUnit_name[int32(i)]
	t := beerjson.SpecificHeatUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONEquipmentBaseForm(context context.Context, i beerprotov1.EquipmentItemType_EquipmentBaseForm) *beerjson.EquipmentBaseForm {
	if i == beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED {
		return nil
	}

	var t beerjson.EquipmentBaseForm
	switch i {
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_HLT:
		t = beerjson.EquipmentBaseForm_HLT
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_MASH_TUN:
		t = beerjson.EquipmentBaseForm_MashTun
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_LAUTER_TUN:
		t = beerjson.EquipmentBaseForm_LauterTun
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_BREW_KETTLE:
		t = beerjson.EquipmentBaseForm_BrewKettle
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_FERMENTER:
		t = beerjson.EquipmentBaseForm_Fermenter
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_AGING_VESSEL:
		t = beerjson.EquipmentBaseForm_AgingVessel
	case beerprotov1.EquipmentItemType_EQUIPMENT_BASE_FORM_PACKAGING_VESSEL:
		t = beerjson.EquipmentBaseForm_PackagingVessel
	}

	return &t
}

func ToJSONCultureInformation(context context.Context, i *beerprotov1.CultureInformation) (*beerjson.CultureInformation, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.Form == beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("form is required"))
	}

	field, _ := reflect.TypeOf(i).Elem().FieldByName("TemperatureRange")
	stack.PushMethod(-1, field)
	temperatureRange, err := ToJSONTemperatureRangeType(context, i.TemperatureRange)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("AttenuationRange")
	stack.PushMethod(-1, field)
	attenuationRange, err := ToJSONPercentRangeType(context, i.AttenuationRange)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	return &beerjson.CultureInformation{
		CultureBaseForm:  ToJSONCultureBaseForm(context, i.Form),
		Producer:         &i.Producer,
		TemperatureRange: temperatureRange,
		Notes:            &i.Notes,
		BestFor:          &i.BestFor,
		Inventory:        ToJSONCultureInventoryType(context, i.Inventory),
		ProductId:        &i.ProductId,
		Name:             &i.Name,
		AlcoholTolerance: ToJSONPercentType(context, i.AlcoholTolerance),
		Glucoamylase:     &i.Glucoamylase,
		CultureBaseType:  ToJSONCultureBaseType(context, i.Type),
		Flocculation:     ToJSONQualitativeRangeType(context, i.Flocculation),
		AttenuationRange: attenuationRange,
		MaxReuse:         &i.MaxReuse,
		Pof:              &i.Pof,
		Zymocide:         ToJSONZymocide(context, i.Zymocide),
	}, stack.Errors()
}

func ToJSONZymocide(context context.Context, i *beerprotov1.Zymocide) *beerjson.Zymocide {
	if i == nil {
		return nil
	}
	return &beerjson.Zymocide{
		No1:     &i.No1,
		No2:     &i.No2,
		No28:    &i.No28,
		Klus:    &i.Klus,
		Neutral: &i.Neutral,
	}
}

func ToJSONQualitativeRangeType(context context.Context, i beerprotov1.QualitativeRangeUnit) *beerjson.QualitativeRangeType {
	if i == beerprotov1.QualitativeRangeUnit_QUALITATIVE_RANGE_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.QualitativeRangeUnit_name[int32(i)]
	t := beerjson.QualitativeRangeType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureBaseType(context context.Context, i beerprotov1.CultureBaseType) *beerjson.CultureBaseType {
	if i == beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.CultureBaseType_name[int32(i)]
	t := beerjson.CultureBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureInventoryType(context context.Context, i *beerprotov1.CultureInventoryType) *beerjson.CultureInventoryType {
	if i == nil {
		return nil
	}
	return &beerjson.CultureInventoryType{
		Liquid:  ToJSONVolumeType(context, i.Liquid),
		Dry:     ToJSONMassType(context, i.Dry),
		Slant:   ToJSONVolumeType(context, i.Slant),
		Culture: ToJSONVolumeType(context, i.Culture),
	}
}

func ToJSONTemperatureRangeType(context context.Context, i *beerprotov1.TemperatureRangeType) (*beerjson.TemperatureRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	temperatureRangeType := &beerjson.TemperatureRangeType{}
	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		temperatureRangeType.Minimum = *ToJSONTemperatureType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		temperatureRangeType.Maximum = *ToJSONTemperatureType(context, i.Maximum)
	}

	return temperatureRangeType, stack.Errors()
}

func ToJSONFermentableType(context context.Context, i *beerprotov1.FermentableType) (*beerjson.FermentableType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.Yield == nil {
		stack.AppendError(fmt.Errorf("yield is required"))
	}

	if i.Color == nil {
		stack.AppendError(fmt.Errorf("color is required"))
	}

	return &beerjson.FermentableType{
		MaxInBatch:                ToJSONPercentType(context, i.MaxInBatch),
		RecommendMash:             &i.RecommendMash,
		Protein:                   ToJSONPercentType(context, i.Protein),
		ProductId:                 &i.ProductId,
		FermentableBaseGrainGroup: ToJSONGrainGroup(context, i.GrainGroup),
		Yield:                     ToJSONYieldType(context, i.Yield),
		FermentableBaseType:       ToJSONFermentableBaseType(context, i.Type),
		Producer:                  &i.Producer,
		AlphaAmylase:              &i.AlphaAmylase,
		Color:                     ToJSONColorType(context, i.Color),
		Name:                      &i.Name,
		DiastaticPower:            ToJSONDiastaticPowerType(context, i.DiastaticPower),
		Moisture:                  ToJSONPercentType(context, i.Moisture),
		Origin:                    &i.Origin,
		Inventory:                 ToJSONFermentableInventoryType(context, i.Inventory),
		KolbachIndex:              ToJSONPercentType(context, i.KolbachIndex),
		Notes:                     &i.Notes,
		Glassy:                    ToJSONPercentType(context, i.Glassy),
		Plump:                     ToJSONPercentType(context, i.Plump),
		Half:                      ToJSONPercentType(context, i.Half),
		Mealy:                     ToJSONPercentType(context, i.Mealy),
		Thru:                      ToJSONPercentType(context, i.Thru),
		Friability:                ToJSONPercentType(context, i.Friability),
		DipH:                      ToJSONAcidityType(context, i.DiPh),
		Viscosity:                 ToJSONViscosityType(context, i.Viscosity),
		DMSP:                      ToJSONConcentrationType(context, i.DmsP),
		FAN:                       ToJSONConcentrationType(context, i.Fan),
		Fermentability:            ToJSONPercentType(context, i.Fermentability),
		BetaGlucan:                ToJSONConcentrationType(context, i.BetaGlucan),
	}, stack.Errors()
}

func ToJSONViscosityType(context context.Context, i *beerprotov1.ViscosityType) *beerjson.ViscosityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.ViscosityUnit_VISCOSITY_UNIT_UNSPECIFIED {
		return &beerjson.ViscosityType{}
	}
	return &beerjson.ViscosityType{
		Value: i.Value,
		Unit:  *ToJSONViscosityUnitType(context, i.Unit),
	}
}

func ToJSONViscosityUnitType(context context.Context, i beerprotov1.ViscosityUnit) *beerjson.ViscosityUnitType {
	if i == beerprotov1.ViscosityUnit_VISCOSITY_UNIT_UNSPECIFIED {
		return nil
	}

	if i == beerprotov1.ViscosityUnit_VISCOSITY_UNIT_CP {
		t := beerjson.ViscosityUnitType_CP
		return &t
	}

	if i == beerprotov1.ViscosityUnit_VISCOSITY_UNIT_MPAS {
		t := beerjson.ViscosityUnitType_MPAS
		return &t
	}

	if unit, ok := beerprotov1.ViscosityUnit_name[int32(i)]; ok {
		t := beerjson.ViscosityUnitType(strings.ToLower(unit))
		return &t
	}

	return nil
}

func ToJSONFermentableInventoryType(context context.Context, i *beerprotov1.FermentableInventoryType) *beerjson.FermentableInventoryType {
	if i == nil {
		return nil
	}

	fermentableInventoryType := &beerjson.FermentableInventoryType{}

	if mass, ok := i.Amount.(*beerprotov1.FermentableInventoryType_Mass); ok {
		fermentableInventoryType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if volume, ok := i.Amount.(*beerprotov1.FermentableInventoryType_Volume); ok {
		fermentableInventoryType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return fermentableInventoryType
}

func ToJSONDiastaticPowerType(context context.Context, i *beerprotov1.DiastaticPowerType) *beerjson.DiastaticPowerType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_UNSPECIFIED {
		return &beerjson.DiastaticPowerType{}
	}

	return &beerjson.DiastaticPowerType{
		Value: i.Value,
		Unit:  ToJSONDiastaticPowerUnitType(context, i.Unit),
	}
}

func ToJSONDiastaticPowerUnitType(context context.Context, i beerprotov1.DiastaticPowerUnit) beerjson.DiastaticPowerUnitType {
	if i == beerprotov1.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_UNSPECIFIED {
		return beerjson.DiastaticPowerUnitType_Lintner
	}

	if i == beerprotov1.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK {
		return beerjson.DiastaticPowerUnitType_WK
	}

	unit := beerprotov1.DiastaticPowerUnit_name[int32(i)]
	return beerjson.DiastaticPowerUnitType(strings.Title(strings.ToLower(unit)))
}

func ToJSONStyleType(context context.Context, i *beerprotov1.StyleType) (*beerjson.StyleType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.StyleType_STYLE_CATEGORIES_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	field, _ := reflect.TypeOf(i).Elem().FieldByName("AlcoholByVolume")
	stack.PushMethod(-1, field)
	alcoholByVolume, err := ToJSONPercentRangeType(context, i.AlcoholByVolume)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("FinalGravity")
	stack.PushMethod(-1, field)
	finalGravity, err := ToJSONGravityRangeType(context, i.FinalGravity)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("OriginalGravity")
	stack.PushMethod(-1, field)
	originalGravity, err := ToJSONGravityRangeType(context, i.OriginalGravity)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Color")
	stack.PushMethod(-1, field)
	color, err := ToJSONColorRangeType(context, i.Color)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Carbonation")
	stack.PushMethod(-1, field)
	carbonation, err := ToJSONCarbonationRangeType(context, i.Carbonation)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("InternationalBitternessUnits")
	stack.PushMethod(-1, field)
	internationalBitternessUnits, err := ToJSONBitternessRangeType(context, i.InternationalBitternessUnits)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	return &beerjson.StyleType{
		Aroma:                        &i.Aroma,
		Ingredients:                  &i.Ingredients,
		CategoryNumber:               &i.CategoryNumber,
		Notes:                        &i.Notes,
		Flavor:                       &i.Flavor,
		Mouthfeel:                    &i.Mouthfeel,
		FinalGravity:                 finalGravity,
		StyleGuide:                   &i.StyleGuide,
		Color:                        color,
		OriginalGravity:              originalGravity,
		Examples:                     &i.Examples,
		Name:                         &i.Name,
		Carbonation:                  carbonation,
		AlcoholByVolume:              alcoholByVolume,
		InternationalBitternessUnits: internationalBitternessUnits,
		Appearance:                   &i.Appearance,
		Category:                     &i.Category,
		StyleLetter:                  &i.StyleLetter,
		KeyType:                      ToJSONStyleType_StyleCategories(context, i.Type),
		OverallImpression:            &i.OverallImpression,
	}, stack.Errors()
}

func ToJSONStyleType_StyleCategories(context context.Context, i beerprotov1.StyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerprotov1.StyleType_STYLE_CATEGORIES_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.StyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONBitternessRangeType(context context.Context, i *beerprotov1.BitternessRangeType) (*beerjson.BitternessRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	bitternessRangeType := &beerjson.BitternessRangeType{}

	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		bitternessRangeType.Minimum = *ToJSONBitternessType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		bitternessRangeType.Maximum = *ToJSONBitternessType(context, i.Maximum)
	}

	return bitternessRangeType, stack.Errors()
}

func ToJSONBitternessType(context context.Context, i *beerprotov1.BitternessType) *beerjson.BitternessType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.BitternessUnit_BITTERNESS_UNIT_UNSPECIFIED {
		return &beerjson.BitternessType{}
	}

	return &beerjson.BitternessType{
		Value: i.Value,
		Unit:  ToJSONBitternessUnitType(context, i.Unit),
	}
}

func ToJSONBitternessUnitType(context context.Context, i beerprotov1.BitternessUnit) beerjson.BitternessUnitType {
	if i == beerprotov1.BitternessUnit_BITTERNESS_UNIT_UNSPECIFIED {
		return beerjson.BitternessUnitType_IBUs
	}

	switch i {
	case beerprotov1.BitternessUnit_BITTERNESS_UNIT_IBUS:
		return beerjson.BitternessUnitType_IBUs
	}

	return beerjson.BitternessUnitType_IBUs
}

func ToJSONPercentRangeType(context context.Context, i *beerprotov1.PercentRangeType) (*beerjson.PercentRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	percentRangeType := &beerjson.PercentRangeType{}
	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		percentRangeType.Minimum = *ToJSONPercentType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		percentRangeType.Maximum = *ToJSONPercentType(context, i.Maximum)
	}

	return percentRangeType, stack.Errors()
}

func ToJSONCarbonationRangeType(context context.Context, i *beerprotov1.CarbonationRangeType) (*beerjson.CarbonationRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	carbonationRangeType := &beerjson.CarbonationRangeType{}
	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		carbonationRangeType.Minimum = *ToJSONCarbonationType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		carbonationRangeType.Maximum = *ToJSONCarbonationType(context, i.Maximum)
	}

	return carbonationRangeType, stack.Errors()
}
func ToJSONCarbonationType(context context.Context, i *beerprotov1.CarbonationType) *beerjson.CarbonationType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.CarbonationUnit_CARBONATION_UNIT_UNSPECIFIED {
		return &beerjson.CarbonationType{}
	}

	return &beerjson.CarbonationType{
		Value: i.Value,
		Unit:  *ToJSONCarbonationUnitType(context, i.Unit),
	}
}

func ToJSONCarbonationUnitType(context context.Context, i beerprotov1.CarbonationUnit) *beerjson.CarbonationUnitType {
	if i == beerprotov1.CarbonationUnit_CARBONATION_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.CarbonationUnit_name[int32(i)]
	t := beerjson.CarbonationUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONColorRangeType(context context.Context, i *beerprotov1.ColorRangeType) (*beerjson.ColorRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	colorRangeType := &beerjson.ColorRangeType{}

	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		colorRangeType.Minimum = *ToJSONColorType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		colorRangeType.Maximum = *ToJSONColorType(context, i.Maximum)
	}

	return colorRangeType, stack.Errors()
}

func ToJSONGravityRangeType(context context.Context, i *beerprotov1.GravityRangeType) (*beerjson.GravityRangeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	gravityRangeType := &beerjson.GravityRangeType{}
	if i.Minimum == nil {
		stack.AppendError(fmt.Errorf("minimum is required"))
	} else {
		gravityRangeType.Minimum = *ToJSONGravityType(context, i.Minimum)
	}

	if i.Maximum == nil {
		stack.AppendError(fmt.Errorf("maximum is required"))
	} else {
		gravityRangeType.Maximum = *ToJSONGravityType(context, i.Maximum)
	}

	return gravityRangeType, stack.Errors()
}

func ToJSONMiscellaneousType(context context.Context, i *beerprotov1.MiscellaneousType) (*beerjson.MiscellaneousType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))

	}

	miscellaneousType := &beerjson.MiscellaneousType{}
	miscellaneousType.UseFor = &i.UseFor
	miscellaneousType.Notes = &i.Notes
	miscellaneousType.Name = &i.Name
	miscellaneousType.Producer = &i.Producer
	miscellaneousType.ProductId = &i.ProductId

	if i.Type != beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_UNSPECIFIED {
		miscellaneousType.MiscellaneousBaseType = ToJSONMiscellaneousBaseType(context, i.Type)
	}

	if i.Inventory != nil {
		field, _ := reflect.TypeOf(i).Elem().FieldByName("Inventory")
		stack.PushMethod(-1, field)
		inventory, err := ToJSONMiscellaneousInventoryType(context, i.Inventory)
		if err != nil {
			stack.AppendError(err)

		}
		miscellaneousType.Inventory = inventory
		stack.Pop()
	}

	return miscellaneousType, stack.Errors()
}

func ToJSONMiscellaneousInventoryType(context context.Context, i *beerprotov1.MiscellaneousInventoryType) (*beerjson.MiscellaneousInventoryType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Amount == nil {
		stack.AppendError(fmt.Errorf("amount is required"))
	}

	miscellaneousInventoryType := &beerjson.MiscellaneousInventoryType{}

	if mass, ok := i.Amount.(*beerprotov1.MiscellaneousInventoryType_Mass); ok {
		miscellaneousInventoryType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if unit, ok := i.Amount.(*beerprotov1.MiscellaneousInventoryType_Unit); ok {
		miscellaneousInventoryType.Amount = ToJSONUnitType(context, unit.Unit)
	}
	if volume, ok := i.Amount.(*beerprotov1.MiscellaneousInventoryType_Volume); ok {
		miscellaneousInventoryType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return miscellaneousInventoryType, stack.Errors()
}

func ToJSONRecipeType(context context.Context, i *beerprotov1.RecipeType) (*beerjson.RecipeType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	recipeType := &beerjson.RecipeType{}

	if i.Type != beerprotov1.RecipeUnit_RECIPE_UNIT_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.BatchSize == nil {
		stack.AppendError(fmt.Errorf("batchSize is required"))
	} else {
		recipeType.BatchSize = *ToJSONVolumeType(context, i.BatchSize)
	}

	if i.Efficiency == nil {
		stack.AppendError(fmt.Errorf("efficiency is required"))
	}

	if i.Ingredients == nil {
		stack.AppendError(fmt.Errorf("ingredients is required"))
	} else {
		field, _ := reflect.TypeOf(i).Elem().FieldByName("Ingredients")
		stack.PushMethod(-1, field)
		ingredients, err := ToJSONIngredientsType(context, i.Ingredients)
		if err != nil {
			stack.AppendError(err)
		} else {
			recipeType.Ingredients = *ingredients
		}
		stack.Pop()
	}

	var created beerjson.DateType
	if i.Created != "" {
		created = beerjson.DateType(i.Created)
	}

	field, _ := reflect.TypeOf(i).Elem().FieldByName("Efficiency")
	stack.PushMethod(-1, field)
	efficiency, err := ToJSONEfficiencyType(context, i.Efficiency)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Boil")
	stack.PushMethod(-1, field)
	boil, err := ToJSONBoilProcedureType(context, i.Boil)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Fermentation")
	stack.PushMethod(-1, field)
	fermentationProcedureType, err := ToJSONFermentationProcedureType(context, i.Fermentation)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Mash")
	stack.PushMethod(-1, field)
	mash, err := ToJSONMashProcedureType(context, i.Mash)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Taste")
	stack.PushMethod(-1, field)
	taste, err := ToJSONTasteType(context, i.Taste)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Packaging")
	stack.PushMethod(-1, field)
	packaging, err := ToJSONPackagingProcedureType(context, i.Packaging)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	field, _ = reflect.TypeOf(i).Elem().FieldByName("Style")
	stack.PushMethod(-1, field)
	style, err := ToJSONRecipeStyleType(context, i.Style)
	if err != nil {
		stack.AppendError(err)
	}
	stack.Pop()

	recipeType.Efficiency = *efficiency
	recipeType.Style = style
	recipeType.IbuEstimate = ToJSONIBUEstimateType(context, i.IbuEstimate)
	recipeType.ColorEstimate = ToJSONColorType(context, i.ColorEstimate)
	recipeType.BeerPH = ToJSONAcidityType(context, i.BeerPh)
	recipeType.Name = i.Name
	recipeType.RecipeTypeType = ToJSONRecipeTypeType(context, i.Type)
	recipeType.Coauthor = &i.Coauthor
	recipeType.OriginalGravity = ToJSONGravityType(context, i.OriginalGravity)
	recipeType.FinalGravity = ToJSONGravityType(context, i.FinalGravity)
	recipeType.Carbonation = &i.Carbonation
	recipeType.Fermentation = fermentationProcedureType
	recipeType.Author = i.Author
	recipeType.Mash = mash
	recipeType.Packaging = packaging
	recipeType.Boil = boil
	recipeType.Taste = taste
	recipeType.CaloriesPerPint = &i.CaloriesPerPint
	recipeType.Created = &created
	recipeType.Notes = &i.Notes
	recipeType.AlcoholByVolume = ToJSONPercentType(context, i.AlcoholByVolume)
	recipeType.ApparentAttenuation = ToJSONPercentType(context, i.ApparentAttenuation)

	return recipeType, stack.Errors()
}

func ToJSONTasteType(context context.Context, i *beerprotov1.TasteType) (*beerjson.TasteType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	return &beerjson.TasteType{
		Notes:  i.Notes,
		Rating: i.Rating,
	}, stack.Errors()
}

func ToJSONBoilProcedureType(context context.Context, i *beerprotov1.BoilProcedureType) (*beerjson.BoilProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	boilProcedureType := &beerjson.BoilProcedureType{}

	context, stack := parser.StackToContext(context)

	if i.BoilTime == nil {
		stack.AppendError(fmt.Errorf("boilTime is required"))
	} else {
		boilProcedureType.BoilTime = *ToJSONTimeType(context, i.BoilTime)
	}

	if i.BoilSteps != nil {
		boilSteps := make([]beerjson.BoilStepType, 0)
		field, _ := reflect.TypeOf(i).Elem().FieldByName("BoilSteps")

		for index, step := range i.BoilSteps {
			stack.PushMethod(index, field)

			if boilStep, err := ToJSONBoilStepType(context, step); err == nil {
				boilSteps = append(boilSteps, *boilStep)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		boilProcedureType.BoilSteps = boilSteps
	}

	if i.PreBoilSize != nil {
		boilProcedureType.PreBoilSize = ToJSONVolumeType(context, i.PreBoilSize)
	}

	boilProcedureType.Name = &i.Name
	boilProcedureType.Description = &i.Description
	boilProcedureType.Notes = &i.Notes

	return boilProcedureType, stack.Errors()
}

func ToJSONBoilStepType(context context.Context, i *beerprotov1.BoilStepType) (*beerjson.BoilStepType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	return &beerjson.BoilStepType{
		EndGravity:               ToJSONGravityType(context, i.EndGravity),
		BoilStepTypeChillingType: ToJSONChillingType(context, i.ChillingType),
		Description:              &i.Description,
		EndTemperature:           ToJSONTemperatureType(context, i.EndTemperature),
		RampTime:                 ToJSONTimeType(context, i.RampTime),
		StepTime:                 ToJSONTimeType(context, i.StepTime),
		StartGravity:             ToJSONGravityType(context, i.StartGravity),
		StartPh:                  ToJSONAcidityType(context, i.StartPh),
		EndPh:                    ToJSONAcidityType(context, i.EndPh),
		Name:                     i.Name,
		StartTemperature:         ToJSONTemperatureType(context, i.StartTemperature),
	}, stack.Errors()
}

func ToJSONChillingType(context context.Context, i beerprotov1.BoilStepTypeChillingType) *beerjson.BoilStepTypeChillingType {
	if i == beerprotov1.BoilStepTypeChillingType_BOIL_STEP_TYPE_CHILLING_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.BoilStepTypeChillingType_name[int32(i)]
	t := beerjson.BoilStepTypeChillingType(strings.ToLower(unit))
	return &t
}

func ToJSONPackagingProcedureType(context context.Context, i *beerprotov1.PackagingProcedureType) (*beerjson.PackagingProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	packagingProcedureType := &beerjson.PackagingProcedureType{}

	if i.PackagingVessels != nil {
		packagingVessels := make([]beerjson.PackagingVesselType, 0)
		field, _ := reflect.TypeOf(i).Elem().FieldByName("PackagingVessels")
		for index, vessels := range i.PackagingVessels {
			stack.PushMethod(index, field)
			if packagingVesselType, err := ToJSONPackagingVesselType(context, vessels); err == nil {
				packagingVessels = append(packagingVessels, *packagingVesselType)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}

		packagingProcedureType.PackagingVessels = packagingVessels
	}

	packagingProcedureType.Name = i.Name
	packagingProcedureType.Description = &i.Description
	packagingProcedureType.Notes = &i.Notes

	if i.PackagedVolume != nil {
		packagingProcedureType.PackagedVolume = ToJSONVolumeType(context, i.PackagedVolume)
	}

	return packagingProcedureType, stack.Errors()
}

func ToJSONPackagingVesselType(context context.Context, i *beerprotov1.PackagingVesselType) (*beerjson.PackagingVesselType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	var packageDate beerjson.DateType
	if i.PackageDate != "" {
		packageDate = beerjson.DateType(i.PackageDate)
	}
	return &beerjson.PackagingVesselType{
		PackagingVesselTypeType: ToJSONPackagingVesselTypeType(context, i.Type),
		StartGravity:            ToJSONGravityType(context, i.StartGravity),
		Name:                    i.Name,
		PackageDate:             &packageDate,
		StepTime:                ToJSONTimeType(context, i.StepTime),
		EndGravity:              ToJSONGravityType(context, i.EndGravity),
		VesselVolume:            ToJSONVolumeType(context, i.VesselVolume),
		VesselQuantity:          &i.VesselQuantity,
		Description:             &i.Description,
		StartPh:                 ToJSONAcidityType(context, i.StartPh),
		Carbonation:             &i.Carbonation,
		StartTemperature:        ToJSONTemperatureType(context, i.StartTemperature),
		EndPh:                   ToJSONAcidityType(context, i.EndPh),
		EndTemperature:          ToJSONTemperatureType(context, i.EndTemperature),
	}, stack.Errors()
}

func ToJSONPackagingVesselTypeType(context context.Context, i beerprotov1.PackagingVesselUnit) *beerjson.PackagingVesselTypeType {
	if i == beerprotov1.PackagingVesselUnit_PACKAGING_VESSEL_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.PackagingVesselUnit_name[int32(i)]
	t := beerjson.PackagingVesselTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONIngredientsType(context context.Context, i *beerprotov1.IngredientsType) (*beerjson.IngredientsType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	ingredientsType := &beerjson.IngredientsType{}

	if i.FermentableAdditions == nil {
		stack.AppendError(fmt.Errorf("fermentableAdditions is required"))
	}

	fermentableAdditions := make([]beerjson.FermentableAdditionType, 0)
	field, _ := reflect.TypeOf(i).Elem().FieldByName("FermentableAdditions")
	for index, fermentable := range i.FermentableAdditions {
		stack.PushMethod(index, field)
		if addition, err := ToJSONFermentableAdditionType(context, fermentable); err == nil {
			fermentableAdditions = append(fermentableAdditions, *addition)
		} else {
			stack.AppendError(err)
		}
		stack.Pop()
	}
	ingredientsType.FermentableAdditions = fermentableAdditions

	if i.MiscellaneousAdditions != nil {
		field, _ := reflect.TypeOf(i).Elem().FieldByName("MiscellaneousAdditions")
		miscellaneousAdditions := make([]beerjson.MiscellaneousAdditionType, 0)
		for index, misc := range i.MiscellaneousAdditions {
			stack.PushMethod(index, field)
			if misc, err := ToJSONMiscellaneousAdditionType(context, misc); err == nil {
				miscellaneousAdditions = append(miscellaneousAdditions, *misc)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		ingredientsType.MiscellaneousAdditions = miscellaneousAdditions
	}

	if i.CultureAdditions != nil {
		field, _ := reflect.TypeOf(i).Elem().FieldByName("CultureAdditions")
		cultureAdditions := make([]beerjson.CultureAdditionType, 0)
		for index, culture := range i.CultureAdditions {
			stack.PushMethod(index, field)
			if cultureAdditionType, err := ToJSONCultureAdditionType(context, culture); err == nil {
				cultureAdditions = append(cultureAdditions, *cultureAdditionType)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		ingredientsType.CultureAdditions = cultureAdditions
	}

	if i.WaterAdditions != nil {
		waterAdditions := make([]beerjson.WaterAdditionType, 0)
		field, _ := reflect.TypeOf(i).Elem().FieldByName("WaterAdditions")
		for index, water := range i.WaterAdditions {
			stack.PushMethod(index, field)
			if waterAdditionType, err := ToJSONWaterAdditionType(context, water); err == nil {
				waterAdditions = append(waterAdditions, *waterAdditionType)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		ingredientsType.WaterAdditions = waterAdditions
	}

	if i.HopAdditions != nil {
		hopAdditions := make([]beerjson.HopAdditionType, 0)
		field, _ := reflect.TypeOf(i).Elem().FieldByName("HopAdditions")
		for index, hop := range i.HopAdditions {
			stack.PushMethod(index, field)
			if hop, err := ToJSONHopAdditionType(context, hop); err == nil {
				hopAdditions = append(hopAdditions, *hop)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		ingredientsType.HopAdditions = hopAdditions
	}

	return ingredientsType, stack.Errors()
}

func ToJSONHopAdditionType(context context.Context, i *beerprotov1.HopAdditionType) (*beerjson.HopAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	hopAdditionType := &beerjson.HopAdditionType{}
	if i.Timing == nil {
		stack.AppendError(fmt.Errorf("timing is required"))
	} else {
		hopAdditionType.Timing = *ToJSONTimingType(context, i.Timing)
	}

	if i.Amount == nil {
		stack.AppendError(fmt.Errorf("amount is required"))
	}

	hopAdditionType.BetaAcid = ToJSONPercentType(context, i.BetaAcid)
	hopAdditionType.Producer = &i.Producer
	hopAdditionType.Origin = &i.Origin
	hopAdditionType.Year = &i.Year
	hopAdditionType.HopVarietyBaseForm = ToJSONHopVarietyBaseForm(context, i.Form)
	hopAdditionType.Name = &i.Name
	hopAdditionType.ProductId = &i.ProductId
	hopAdditionType.AlphaAcid = ToJSONPercentType(context, i.AlphaAcid)

	if mass, ok := i.Amount.(*beerprotov1.HopAdditionType_Mass); ok {
		hopAdditionType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if volume, ok := i.Amount.(*beerprotov1.HopAdditionType_Volume); ok {
		hopAdditionType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return hopAdditionType, stack.Errors()
}

func ToJSONHopVarietyBaseForm(context context.Context, i beerprotov1.HopVarietyBaseForm) *beerjson.HopVarietyBaseForm {
	if i == beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_UNSPECIFIED {
		return nil
	}

	var t beerjson.HopVarietyBaseForm

	switch i {
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_EXTRACT:
		t = beerjson.HopVarietyBaseForm_Extract
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_LEAF:
		t = beerjson.HopVarietyBaseForm_Leaf
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_LEAFWET:
		t = beerjson.HopVarietyBaseForm_LeafWet
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PELLET:
		t = beerjson.HopVarietyBaseForm_Pellet
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_POWDER:
		t = beerjson.HopVarietyBaseForm_Powder
	case beerprotov1.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PLUG:
		t = beerjson.HopVarietyBaseForm_Plug
	}

	return &t
}

func ToJSONFermentableAdditionType(context context.Context, i *beerprotov1.FermentableAdditionType) (*beerjson.FermentableAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.Yield == nil {
		stack.AppendError(fmt.Errorf("yield is required"))
	}

	if i.Color == nil {
		stack.AppendError(fmt.Errorf("color is required"))
	}

	if i.Amount == nil {
		stack.AppendError(fmt.Errorf("amount is required"))
	}

	fermentableAdditionType := &beerjson.FermentableAdditionType{
		FermentableBaseType:       ToJSONFermentableBaseType(context, i.Type),
		Origin:                    &i.Origin,
		FermentableBaseGrainGroup: ToJSONGrainGroup(context, i.GrainGroup),
		Yield:                     ToJSONYieldType(context, i.Yield),
		Color:                     ToJSONColorType(context, i.Color),
		Name:                      &i.Name,
		Producer:                  &i.Producer,
		ProductId:                 &i.ProductId,
		Timing:                    ToJSONTimingType(context, i.Timing),
	}

	if mass, ok := i.Amount.(*beerprotov1.FermentableAdditionType_Mass); ok {
		fermentableAdditionType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if volume, ok := i.Amount.(*beerprotov1.FermentableAdditionType_Volume); ok {
		fermentableAdditionType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return fermentableAdditionType, stack.Errors()
}

func ToJSONYieldType(context context.Context, i *beerprotov1.YieldType) *beerjson.YieldType {
	if i == nil {
		return nil
	}

	return &beerjson.YieldType{
		FineGrind:            ToJSONPercentType(context, i.FineGrind),
		CoarseGrind:          ToJSONPercentType(context, i.CoarseGrind),
		FineCoarseDifference: ToJSONPercentType(context, i.FineCoarseDifference),
		Potential:            ToJSONGravityType(context, i.Potential),
	}
}

func ToJSONGrainGroup(context context.Context, i beerprotov1.GrainGroup) *beerjson.FermentableBaseGrainGroup {
	if i == beerprotov1.GrainGroup_GRAIN_GROUP_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.GrainGroup_name[int32(i)]
	t := beerjson.FermentableBaseGrainGroup(strings.ToLower(unit))
	return &t
}

func ToJSONFermentableBaseType(context context.Context, i beerprotov1.FermentableBaseType) *beerjson.FermentableBaseType {
	if i == beerprotov1.FermentableBaseType_FERMENTABLE_BASE_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.FermentableBaseType_name[int32(i)]
	t := beerjson.FermentableBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONWaterAdditionType(context context.Context, i *beerprotov1.WaterAdditionType) (*beerjson.WaterAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Calcium == nil {
		stack.AppendError(fmt.Errorf("calcium is required"))
	}

	if i.Bicarbonate == nil {
		stack.AppendError(fmt.Errorf("bicarbonate is required"))
	}

	if i.Sulfate == nil {
		stack.AppendError(fmt.Errorf("sulfate is required"))
	}

	if i.Chloride == nil {
		stack.AppendError(fmt.Errorf("chloride is required"))
	}

	if i.Sodium == nil {
		stack.AppendError(fmt.Errorf("sodium is required"))
	}

	if i.Magnesium == nil {
		stack.AppendError(fmt.Errorf("magnesium is required"))
	}

	return &beerjson.WaterAdditionType{
		Flouride:    ToJSONConcentrationType(context, i.Flouride),
		Sulfate:     ToJSONConcentrationType(context, i.Sulfate),
		Producer:    &i.Producer,
		Nitrate:     ToJSONConcentrationType(context, i.Nitrate),
		Nitrite:     ToJSONConcentrationType(context, i.Nitrite),
		Chloride:    ToJSONConcentrationType(context, i.Chloride),
		Amount:      ToJSONVolumeType(context, i.Amount),
		Name:        &i.Name,
		Potassium:   ToJSONConcentrationType(context, i.Potassium),
		Magnesium:   ToJSONConcentrationType(context, i.Magnesium),
		Iron:        ToJSONConcentrationType(context, i.Iron),
		Bicarbonate: ToJSONConcentrationType(context, i.Bicarbonate),
		Calcium:     ToJSONConcentrationType(context, i.Calcium),
		Carbonate:   ToJSONConcentrationType(context, i.Carbonate),
		Sodium:      ToJSONConcentrationType(context, i.Sodium),
	}, stack.Errors()
}

func ToJSONConcentrationType(context context.Context, i *beerprotov1.ConcentrationType) *beerjson.ConcentrationType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_UNSPECIFIED {
		return &beerjson.ConcentrationType{}
	}

	return &beerjson.ConcentrationType{
		Value: i.Value,
		Unit:  *ToJSONConcentrationUnitType(context, i.Unit),
	}
}

func ToJSONConcentrationUnitType(context context.Context, i beerprotov1.ConcentrationUnit) *beerjson.ConcentrationUnitType {
	if i == beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_UNSPECIFIED {
		return nil
	}

	if i == beerprotov1.ConcentrationUnit_CONCENTRATION_UNIT_MGL {
		t := beerjson.ConcentrationUnitType_MgL
		return &t
	}

	if unit, ok := beerprotov1.ConcentrationUnit_name[int32(i)]; ok {
		t := beerjson.ConcentrationUnitType(strings.ToLower(unit))
		return &t
	}

	return nil
}

func ToJSONCultureAdditionType(context context.Context, i *beerprotov1.CultureAdditionType) (*beerjson.CultureAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.CultureBaseType_CULTURE_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.Form == beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("form is required"))
	}

	cultureAdditionType := &beerjson.CultureAdditionType{
		CultureBaseForm:   ToJSONCultureBaseForm(context, i.Form),
		ProductId:         &i.ProductId,
		Name:              &i.Name,
		CellCountBillions: &i.CellCountBillions,
		TimesCultured:     &i.TimesCultured,
		Producer:          &i.Producer,
		CultureBaseType:   ToJSONCultureBaseType(context, i.Type),
		Attenuation:       ToJSONPercentType(context, i.Attenuation),
		Timing:            ToJSONTimingType(context, i.Timing),
	}

	if mass, ok := i.Amount.(*beerprotov1.CultureAdditionType_Mass); ok {
		cultureAdditionType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if unit, ok := i.Amount.(*beerprotov1.CultureAdditionType_Unit); ok {
		cultureAdditionType.Amount = ToJSONUnitType(context, unit.Unit)
	}
	if volume, ok := i.Amount.(*beerprotov1.CultureAdditionType_Volume); ok {
		cultureAdditionType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return cultureAdditionType, stack.Errors()
}

func ToJSONCultureBaseForm(context context.Context, i beerprotov1.CultureBaseForm) *beerjson.CultureBaseForm {
	if i == beerprotov1.CultureBaseForm_CULTURE_BASE_FORM_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.CultureBaseForm_name[int32(i)]
	t := beerjson.CultureBaseForm(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousAdditionType(context context.Context, i *beerprotov1.MiscellaneousAdditionType) (*beerjson.MiscellaneousAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	if i.Timing == nil {
		stack.AppendError(fmt.Errorf("timing is required"))
	}

	if i.Amount == nil {
		stack.AppendError(fmt.Errorf("amount is required"))
	}

	miscellaneousAdditionType := &beerjson.MiscellaneousAdditionType{
		Name:                  &i.Name,
		Producer:              &i.Producer,
		Timing:                ToJSONTimingType(context, i.Timing),
		ProductId:             &i.ProductId,
		MiscellaneousBaseType: ToJSONMiscellaneousBaseType(context, i.Type),
	}

	if mass, ok := i.Amount.(*beerprotov1.MiscellaneousAdditionType_Mass); ok {
		miscellaneousAdditionType.Amount = ToJSONMassType(context, mass.Mass)
	}

	if unit, ok := i.Amount.(*beerprotov1.MiscellaneousAdditionType_Unit); ok {
		miscellaneousAdditionType.Amount = ToJSONUnitType(context, unit.Unit)
	}
	if volume, ok := i.Amount.(*beerprotov1.MiscellaneousAdditionType_Volume); ok {
		miscellaneousAdditionType.Amount = ToJSONVolumeType(context, volume.Volume)
	}

	return miscellaneousAdditionType, stack.Errors()
}

func ToJSONUnitType(context context.Context, i *beerprotov1.UnitType) *beerjson.UnitType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.UnitUnit_UNIT_UNIT_UNSPECIFIED {
		return &beerjson.UnitType{}
	}

	return &beerjson.UnitType{
		Value: i.Value,
		Unit:  *ToJSONUnitUnitType(context, i.Unit),
	}
}

func ToJSONUnitUnitType(context context.Context, i beerprotov1.UnitUnit) *beerjson.UnitUnitType {
	if i == beerprotov1.UnitUnit_UNIT_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.UnitUnit_name[int32(i)]
	t := beerjson.UnitUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMassType(context context.Context, i *beerprotov1.MassType) *beerjson.MassType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.MassUnit_MASS_UNIT_UNSPECIFIED {
		return &beerjson.MassType{}
	}

	return &beerjson.MassType{
		Value: i.Value,
		Unit:  *ToJSONMassUnitType(context, i.Unit),
	}
}

func ToJSONMassUnitType(context context.Context, i beerprotov1.MassUnit) *beerjson.MassUnitType {
	if i == beerprotov1.MassUnit_MASS_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.MassUnit_name[int32(i)]
	t := beerjson.MassUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousBaseType(context context.Context, i beerprotov1.MiscellaneousBaseType) *beerjson.MiscellaneousBaseType {
	if i == beerprotov1.MiscellaneousBaseType_MISCELLANEOUS_BASE_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.MiscellaneousBaseType_name[int32(i)]
	t := beerjson.MiscellaneousBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONTimingType(context context.Context, i *beerprotov1.TimingType) *beerjson.TimingType {
	if i == nil {
		return nil
	}
	return &beerjson.TimingType{
		Time:            ToJSONTimeType(context, i.Time),
		Duration:        ToJSONTimeType(context, i.Duration),
		Continuous:      &i.Continuous,
		SpecificGravity: ToJSONGravityType(context, i.SpecificGravity),
		PH:              ToJSONAcidityType(context, i.Ph),
		Step:            &i.Step,
		Use:             ToJSONUseType(context, i.Use),
	}
}

func ToJSONUseType(context context.Context, i beerprotov1.UseType) *beerjson.UseType {
	if i == beerprotov1.UseType_USE_TYPE_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.UseType_name[int32(i)]
	t := beerjson.UseType(strings.ToLower(unit))
	return &t
}

func ToJSONFermentationProcedureType(context context.Context, i *beerprotov1.FermentationProcedureType) (*beerjson.FermentationProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.FermentationSteps == nil {
		stack.AppendError(fmt.Errorf("fermentationSteps is required"))
	}

	fermentationProcedureType := &beerjson.FermentationProcedureType{}

	fermentationProcedureType.Name = i.Name
	if i.FermentationSteps != nil {
		field, _ := reflect.TypeOf(i).Elem().FieldByName("FermentationSteps")
		steps := make([]beerjson.FermentationStepType, 0)
		for index, step := range i.FermentationSteps {
			stack.PushMethod(index, field)
			if fermentationStepType, err := ToJSONFermentationStepType(context, step); err == nil {
				steps = append(steps, *fermentationStepType)
			} else {
				stack.AppendError(err)
			}
			stack.Pop()
		}
		fermentationProcedureType.FermentationSteps = steps
	}

	fermentationProcedureType.Description = &i.Description
	fermentationProcedureType.Notes = &i.Notes

	return fermentationProcedureType, stack.Errors()
}

func ToJSONFermentationStepType(context context.Context, i *beerprotov1.FermentationStepType) (*beerjson.FermentationStepType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	return &beerjson.FermentationStepType{
		Name:             i.Name,
		EndTemperature:   ToJSONTemperatureType(context, i.EndTemperature),
		StepTime:         ToJSONTimeType(context, i.StepTime),
		FreeRise:         &i.FreeRise,
		StartGravity:     ToJSONGravityType(context, i.StartGravity),
		EndGravity:       ToJSONGravityType(context, i.EndGravity),
		StartPh:          ToJSONAcidityType(context, i.StartPh),
		Description:      &i.Description,
		StartTemperature: ToJSONTemperatureType(context, i.StartTemperature),
		EndPh:            ToJSONAcidityType(context, i.EndPh),
		Vessel:           &i.Vessel,
	}, stack.Errors()
}

func ToJSONGravityType(context context.Context, i *beerprotov1.GravityType) *beerjson.GravityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.GravityUnit_GRAVITY_UNIT_UNSPECIFIED {
		return &beerjson.GravityType{}
	}

	return &beerjson.GravityType{
		Value: i.Value,
		Unit:  *ToJSONGravityUnitType(context, i.Unit),
	}
}

func ToJSONGravityUnitType(context context.Context, i beerprotov1.GravityUnit) *beerjson.GravityUnitType {
	if i == beerprotov1.GravityUnit_GRAVITY_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.GravityUnit_name[int32(i)]
	t := beerjson.GravityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONRecipeTypeType(context context.Context, i beerprotov1.RecipeUnit) beerjson.RecipeTypeType {
	if i == beerprotov1.RecipeUnit_RECIPE_UNIT_UNSPECIFIED {
		return beerjson.RecipeTypeType_AllGrain
	}

	return beerjson.RecipeTypeType(strings.ToLower(i.String()))
}

func ToJSONColorType(context context.Context, i *beerprotov1.ColorType) *beerjson.ColorType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.ColorUnit_COLOR_UNIT_UNSPECIFIED {
		return &beerjson.ColorType{}
	}

	return &beerjson.ColorType{
		Value: i.Value,
		Unit:  *ToJSONColorUnitType(context, i.Unit),
	}
}

func ToJSONColorUnitType(context context.Context, i beerprotov1.ColorUnit) *beerjson.ColorUnitType {
	if i == beerprotov1.ColorUnit_COLOR_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.ColorUnit_name[int32(i)]

	if i == beerprotov1.ColorUnit_COLOR_UNIT_LOVI {
		unit = strings.Title(strings.ToLower(unit))
	}
	t := beerjson.ColorUnitType(unit)
	return &t
}

func ToJSONIBUEstimateType(context context.Context, i *beerprotov1.IBUEstimateType) *beerjson.IBUEstimateType {
	if i == nil {
		return nil
	}

	return &beerjson.IBUEstimateType{
		Method: ToJSONIBUMethodType(context, i.Method),
	}
}

func ToJSONIBUMethodType(context context.Context, i beerprotov1.IBUMethodUnit) *beerjson.IBUMethodType {
	if i == beerprotov1.IBUMethodUnit_IBU_METHOD_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.IBUMethodUnit_name[int32(i)]
	t := beerjson.IBUMethodType(strings.Title(strings.ToLower(unit)))
	return &t
}

func ToJSONRecipeStyleType(context context.Context, i *beerprotov1.RecipeStyleType) (*beerjson.RecipeStyleType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.Type == beerprotov1.RecipeStyleType_STYLE_CATEGORIES_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))

	}

	return &beerjson.RecipeStyleType{
		KeyType:        ToJSONRecipeStyleType_StyleCategories(context, i.Type),
		Name:           &i.Name,
		Category:       &i.Category,
		CategoryNumber: &i.CategoryNumber,
		StyleGuide:     &i.StyleGuide,
		StyleLetter:    &i.StyleLetter,
	}, stack.Errors()
}

func ToJSONRecipeStyleType_StyleCategories(context context.Context, i beerprotov1.RecipeStyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerprotov1.RecipeStyleType_STYLE_CATEGORIES_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.RecipeStyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONEfficiencyType(context context.Context, i *beerprotov1.EfficiencyType) (*beerjson.EfficiencyType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	efficiencyType := &beerjson.EfficiencyType{}
	if i.Brewhouse == nil {
		stack.AppendError(fmt.Errorf("brewhouse is required"))
	} else {
		efficiencyType.Brewhouse = *ToJSONPercentType(context, i.Brewhouse)
	}

	efficiencyType.Conversion = ToJSONPercentType(context, i.Conversion)
	efficiencyType.Lauter = ToJSONPercentType(context, i.Lauter)
	efficiencyType.Mash = ToJSONPercentType(context, i.Mash)

	return efficiencyType, stack.Errors()
}

func ToJSONPercentType(context context.Context, i *beerprotov1.PercentType) *beerjson.PercentType {
	if i == nil {
		return nil
	}

	return &beerjson.PercentType{
		Value: i.Value,
		Unit:  ToJSONPercentUnitType(context, i.Unit),
	}
}

func ToJSONPercentUnitType(context context.Context, i beerprotov1.PercentUnit) beerjson.PercentUnitType {
	if i == beerprotov1.PercentUnit_PERCENT_UNIT_UNSPECIFIED {
		return beerjson.PercentUnitType_No
	}

	return beerjson.PercentUnitType(strings.ToLower(i.String()))
}

func ToJSONMashProcedureType(context context.Context, i *beerprotov1.MashProcedureType) (*beerjson.MashProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)
	mashProcedureType := &beerjson.MashProcedureType{}

	if i.MashSteps == nil {
		stack.AppendError(fmt.Errorf("mashSteps is required"))
	}

	if i.GrainTemperature == nil {
		stack.AppendError(fmt.Errorf("grainTemperature is required"))
	} else {
		mashProcedureType.GrainTemperature = *ToJSONTemperatureType(context, i.GrainTemperature)
	}

	if i.MashSteps != nil {
		mashSteps := []beerjson.MashStepType{}
		field, _ := reflect.TypeOf(i).Elem().FieldByName("MashSteps")
		for index, step := range i.MashSteps {
			stack.PushMethod(index, field)
			if step, err := ToJSONMashStepType(context, step); err == nil {
				mashSteps = append(mashSteps, *step)
			} else {
				stack.AppendError(err)
			}
		}
		mashProcedureType.MashSteps = mashSteps
	}

	mashProcedureType.Name = i.Name
	mashProcedureType.Notes = &i.Notes

	return mashProcedureType, stack.Errors()
}

func ToJSONMashStepType(context context.Context, i *beerprotov1.MashStepType) (*beerjson.MashStepType, error) {
	if i == nil {
		return nil, nil
	}

	context, stack := parser.StackToContext(context)

	if i.StepTemperature == nil {
		stack.AppendError(fmt.Errorf("stepTemperature is required"))
	}

	if i.StepTime == nil {
		stack.AppendError(fmt.Errorf("stepTime is required"))
	}

	if i.Type == beerprotov1.MashStepType_MASH_STEP_UNIT_UNSPECIFIED {
		stack.AppendError(fmt.Errorf("type is required"))
	}

	mashStepType := &beerjson.MashStepType{
		StepTime:          *ToJSONTimeType(context, i.StepTime),
		RampTime:          ToJSONTimeType(context, i.RampTime),
		EndTemperature:    ToJSONTemperatureType(context, i.EndTemperature),
		InfuseTemperature: ToJSONTemperatureType(context, i.InfuseTemperature),
		StartPH:           ToJSONAcidityType(context, i.StartPh),
		EndPH:             ToJSONAcidityType(context, i.EndPh),
		Name:              i.Name,
		MashStepTypeType:  *ToJSONMashStepTypeType(context, i.Type),
		Amount:            ToJSONVolumeType(context, i.Amount),
		StepTemperature:   *ToJSONTemperatureType(context, i.StepTemperature),
		WaterGrainRatio:   ToJSONSpecificVolumeType(context, i.WaterGrainRatio),
	}

	mashStepType.Description = &i.Description

	return mashStepType, stack.Errors()
}

func ToJSONVolumeType(context context.Context, i *beerprotov1.VolumeType) *beerjson.VolumeType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.VolumeUnit_VOLUME_UNIT_UNSPECIFIED {
		return &beerjson.VolumeType{}
	}

	return &beerjson.VolumeType{
		Value: i.Value,
		Unit:  ToJSONVolumeUnitType(context, i.Unit),
	}
}

func ToJSONVolumeUnitType(context context.Context, i beerprotov1.VolumeUnit) beerjson.VolumeUnitType {
	if i == beerprotov1.VolumeUnit_VOLUME_UNIT_UNSPECIFIED {
		return beerjson.VolumeUnitType_Ml
	}

	return beerjson.VolumeUnitType(strings.ToLower(i.String()))
}

func ToJSONSpecificVolumeType(context context.Context, i *beerprotov1.SpecificVolumeType) *beerjson.SpecificVolumeType {
	if i == nil {
		return nil
	}
	if i.Unit == beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_UNSPECIFIED {
		return &beerjson.SpecificVolumeType{}
	}
	return &beerjson.SpecificVolumeType{
		Value: i.Value,
		Unit:  ToJSONSpecificVolumeUnitType(context, i.Unit),
	}
}

func ToJSONSpecificVolumeUnitType(context context.Context, i beerprotov1.SpecificVolumeUnit) beerjson.SpecificVolumeUnitType {
	if i == beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_UNSPECIFIED {
		return beerjson.SpecificVolumeUnitType_LG
	}

	switch i {
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_QTLB:
		return beerjson.SpecificVolumeUnitType_QtLb
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALLB:
		return beerjson.SpecificVolumeUnitType_GalLb
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALOZ:
		return beerjson.SpecificVolumeUnitType_GalOz
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LG:
		return beerjson.SpecificVolumeUnitType_LG
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LKG:
		return beerjson.SpecificVolumeUnitType_LKg
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FLOZOZ:
		return beerjson.SpecificVolumeUnitType_FlozOz
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_M3KG:
		return beerjson.SpecificVolumeUnitType_M3Kg
	case beerprotov1.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FT3LB:
		return beerjson.SpecificVolumeUnitType_Ft3Lb
	}

	return beerjson.SpecificVolumeUnitType(i.String())
}

func ToJSONMashStepTypeType(context context.Context, i beerprotov1.MashStepType_MashStepUnit) *beerjson.MashStepTypeType {
	if i == beerprotov1.MashStepType_MASH_STEP_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.MashStepType_MashStepUnit_name[int32(i)]
	t := beerjson.MashStepTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONAcidityType(context context.Context, i *beerprotov1.AcidityType) *beerjson.AcidityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.AcidityUnit_ACIDITY_UNIT_UNSPECIFIED {
		return &beerjson.AcidityType{}
	}

	return &beerjson.AcidityType{
		Value: i.Value,
		Unit:  *ToJSONAcidityUnitType(context, i.Unit),
	}
}

func ToJSONAcidityUnitType(context context.Context, i beerprotov1.AcidityUnit) *beerjson.AcidityUnitType {
	if i == beerprotov1.AcidityUnit_ACIDITY_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.AcidityUnit_name[int32(i)]
	t := beerjson.AcidityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTimeType(context context.Context, i *beerprotov1.TimeType) *beerjson.TimeType {
	if i == nil {
		return nil
	}

	if i.Unit == beerprotov1.TimeUnit_TIME_UNIT_UNSPECIFIED {
		return &beerjson.TimeType{}
	}

	return &beerjson.TimeType{
		Value: i.Value,
		Unit:  *ToJSONTimeUnitType(context, i.Unit),
	}
}

func ToJSONTemperatureType(context context.Context, i *beerprotov1.TemperatureType) *beerjson.TemperatureType {
	if i == nil {
		return nil
	}
	if i.Unit == beerprotov1.TemperatureUnit_TEMPERATURE_UNIT_UNSPECIFIED {
		return &beerjson.TemperatureType{}
	}

	return &beerjson.TemperatureType{
		Value: i.Value,
		Unit:  *ToJSONTemperatureUnitType(context, i.Unit),
	}
}

func ToJSONTimeUnitType(context context.Context, i beerprotov1.TimeUnit) *beerjson.TimeUnitType {
	if i == beerprotov1.TimeUnit_TIME_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.TimeUnit_name[int32(i)]
	t := beerjson.TimeUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTemperatureUnitType(context context.Context, i beerprotov1.TemperatureUnit) *beerjson.TemperatureUnitType {
	if i == beerprotov1.TemperatureUnit_TEMPERATURE_UNIT_UNSPECIFIED {
		return nil
	}

	unit := beerprotov1.TemperatureUnit_name[int32(i)]
	t := beerjson.TemperatureUnitType(unit)
	return &t
}
