package beerJSON

import (
	"fmt"
	"strings"

	"github.com/beerproto/beerjson.go"
	beerproto "github.com/beerproto/beerproto_go"
)

func MapToJSON(i *beerproto.Recipe) (*beerjson.Beerjson, error) {
	if i == nil {
		return nil, nil
	}

	if i.Version == 0 {
		return nil, fmt.Errorf("version is required")
	}

	output := &beerjson.Beerjson{
		Mashes:                   []beerjson.MashProcedureType{},
		Recipes:                  []beerjson.RecipeType{},
		MiscellaneousIngredients: []beerjson.MiscellaneousType{},
		Styles:                   []beerjson.StyleType{},
		Fermentations:            []beerjson.FermentationProcedureType{},
		Boil:                     []beerjson.BoilProcedureType{},
		Version:                  beerjson.VersionType(i.Version),
		Fermentables:             []beerjson.FermentableType{},
		TimingObject:             ToJSONTimingType(i.TimingObject),
		Cultures:                 []beerjson.CultureInformation{},
		Equipments:               []beerjson.EquipmentType{},
		Packaging:                []beerjson.PackagingProcedureType{},
		HopVarieties:             []beerjson.VarietyInformation{},
		Profiles:                 []beerjson.WaterBase{},
	}

	for _, mash := range i.Mashes {
		if mash, err := ToJSONMashProcedureType(mash); err == nil {
			output.Mashes = append(output.Mashes, *mash)
		} else {
			return nil, err
		}
	}

	for _, recipe := range i.Recipes {
		if r, err := ToJSONRecipeType(recipe); err == nil {
			output.Recipes = append(output.Recipes, *r)
		} else {
			return nil, err
		}
	}

	for _, ingredients := range i.MiscellaneousIngredients {
		if mice, err := ToJSONMiscellaneousType(ingredients); err == nil {
			output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, *mice)
		}
	}

	for _, style := range i.Styles {
		if style, err := ToJSONStyleType(style); err == nil {
			output.Styles = append(output.Styles, *style)
		} else {
			return nil, err
		}
	}

	for _, fermentation := range i.Fermentations {
		if fermentationProcedureType, err := ToJSONFermentationProcedureType(fermentation); err == nil {
			output.Fermentations = append(output.Fermentations, *fermentationProcedureType)
		} else {
			return nil, err
		}
	}

	for _, boil := range i.Boil {
		if boil, err := ToJSONBoilProcedureType(boil); err == nil {
			output.Boil = append(output.Boil, *boil)
		} else {
			return nil, err
		}
	}

	for _, fermentable := range i.Fermentables {
		if fermentableType, err := ToJSONFermentableType(fermentable); err == nil {
			output.Fermentables = append(output.Fermentables, *fermentableType)
		} else {
			return nil, err
		}
	}

	for _, culture := range i.Cultures {
		if culture, err := ToJSONCultureInformation(culture); err == nil {
			output.Cultures = append(output.Cultures, *culture)
		} else {
			return nil, err
		}
	}

	for _, equipment := range i.Equipments {
		if equipmentType, err := ToJSONEquipmentType(equipment); err == nil{
			output.Equipments = append(output.Equipments, *equipmentType)
		} else {
			return nil, err
		}
	}

	for _, packaging := range i.Packaging {
		if packagingProcedure, err := ToJSONPackagingProcedureType(packaging); err == nil {
			output.Packaging = append(output.Packaging, *packagingProcedure)
		} else {
			return nil, err
		}
	}

	for _, hopVariety := range i.HopVarieties {
		if hop, err := ToJSONVarietyInformation(hopVariety); err == nil {
			output.HopVarieties = append(output.HopVarieties, *hop)
		} else {
			return nil, err
		}
	}

	for _, profile := range i.Profiles {
		if water, err := ToJSONWaterBase(profile); err == nil {
			output.Profiles = append(output.Profiles, *water)
		} else {
			return nil, err
		}
	}

	return output, nil
}

func ToJSONWaterBase(i *beerproto.WaterBase) (*beerjson.WaterBase, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Calcium == nil {
		return nil, fmt.Errorf("calcium is required")
	}

	if i.Bicarbonate == nil {
		return nil, fmt.Errorf("bicarbonate is required")
	}

	if i.Sulfate == nil {
		return nil, fmt.Errorf("sulfate is required")
	}

	if i.Chloride == nil {
		return nil, fmt.Errorf("chloride is required")
	}

	if i.Sodium == nil {
		return nil, fmt.Errorf("sodium is required")
	}

	if i.Magnesium == nil {
		return nil, fmt.Errorf("magnesium is required")
	}

	return &beerjson.WaterBase{
		Calcium:     *ToJSONConcentrationType(i.Calcium),
		Nitrite:     ToJSONConcentrationType(i.Nitrite),
		Chloride:    *ToJSONConcentrationType(i.Chloride),
		Name:        i.Name,
		Potassium:   ToJSONConcentrationType(i.Potassium),
		Carbonate:   ToJSONConcentrationType(i.Carbonate),
		Iron:        ToJSONConcentrationType(i.Iron),
		Flouride:    ToJSONConcentrationType(i.Flouride),
		Sulfate:     *ToJSONConcentrationType(i.Sulfate),
		Magnesium:   *ToJSONConcentrationType(i.Magnesium),
		Producer:    &i.Producer,
		Bicarbonate: *ToJSONConcentrationType(i.Bicarbonate),
		Nitrate:     ToJSONConcentrationType(i.Nitrate),
		Sodium:      *ToJSONConcentrationType(i.Sodium),
	}, nil
}

func ToJSONVarietyInformation(i *beerproto.VarietyInformation) (*beerjson.VarietyInformation, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.AlphaAcid == nil {
		return nil, fmt.Errorf("alphaAcid is required")
	}

	return &beerjson.VarietyInformation{
		Inventory:              ToJSONHopInventoryType(i.Inventory),
		VarietyInformationType: ToJSONVarietyInformationType(i.Type),
		OilContent:             ToJSONOilContentType(i.OilContent),
		PercentLost:            ToJSONPercentType(i.PercentLost),
		ProductId:              &i.ProductId,
		AlphaAcid:              ToJSONPercentType(i.AlphaAcid),
		BetaAcid:               ToJSONPercentType(i.BetaAcid),
		Name:                   &i.Name,
		Origin:                 &i.Origin,
		Substitutes:            &i.Substitutes,
		Year:                   &i.Year,
		HopVarietyBaseForm:     ToJSONHopVarietyBaseForm(i.Form),
		Producer:               &i.Producer,
		Notes:                  &i.Notes,
	}, nil
}

func ToJSONOilContentType(i *beerproto.OilContentType) *beerjson.OilContentType {
	if i == nil {
		return nil
	}

	return &beerjson.OilContentType{
		Polyphenols:       ToJSONPercentType(i.Polyphenols),
		TotalOilMlPer100g: &i.TotalOilMlPer_100G,
		Farnesene:         ToJSONPercentType(i.Farnesene),
		Limonene:          ToJSONPercentType(i.Limonene),
		Nerol:             ToJSONPercentType(i.Nerol),
		Geraniol:          ToJSONPercentType(i.Geraniol),
		BPinene:           ToJSONPercentType(i.BPinene),
		Linalool:          ToJSONPercentType(i.Linalool),
		Caryophyllene:     ToJSONPercentType(i.Caryophyllene),
		Cohumulone:        ToJSONPercentType(i.Cohumulone),
		Xanthohumol:       ToJSONPercentType(i.Xanthohumol),
		Humulene:          ToJSONPercentType(i.Humulene),
		Myrcene:           ToJSONPercentType(i.Myrcene),
		Pinene:            ToJSONPercentType(i.Pinene),
	}
}

func ToJSONVarietyInformationType(i beerproto.VarietyInformation_VarietyInformationType) *beerjson.VarietyInformationType {
	if i == beerproto.VarietyInformation_NULL_VARIETYINFORMATIONTYPE {
		return nil
	}

	unit := beerproto.VarietyInformation_VarietyInformationType_name[int32(i)]
	t := beerjson.VarietyInformationType(strings.ToLower(unit))
	return &t
}

func ToJSONHopInventoryType(i *beerproto.HopInventoryType) *beerjson.HopInventoryType {
	if i == nil {
		return nil
	}

	hopInventoryType := &beerjson.HopInventoryType{}

	if mass, ok := i.Amount.(*beerproto.HopInventoryType_Mass); ok {
		hopInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.HopInventoryType_Volume); ok {
		hopInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return hopInventoryType
}

func ToJSONEquipmentType(i *beerproto.EquipmentType) (*beerjson.EquipmentType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.EquipmentItems == nil {
		return nil, fmt.Errorf("equipmentItems is required")
	}

	equipmentType := &beerjson.EquipmentType{
		Name: i.Name,
	}

	equipmentItemType := []beerjson.EquipmentItemType{}
	for _, item := range i.EquipmentItems {
		if equipmentType, err := ToJSONEquipmentItemType(item); err == nil {
			equipmentItemType = append(equipmentItemType, *equipmentType)
		} else {
			return nil, err
		}
	}

	equipmentType.EquipmentItems = equipmentItemType


	return equipmentType, nil
}

func ToJSONEquipmentItemType(i *beerproto.EquipmentItemType) (*beerjson.EquipmentItemType, error) {
	if i == nil {
		return nil, nil
	}
	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Form == beerproto.EquipmentItemType_NULL {
		return nil, fmt.Errorf("form is required")
	}

	if i.Type == "" {
		return nil, fmt.Errorf("type is required")
	}

	if i.MaximumVolume == nil  {
		return nil, fmt.Errorf("maximumVolume is required")
	}

	if i.Loss == nil  {
		return nil, fmt.Errorf("loss is required")
	}

	return &beerjson.EquipmentItemType{
		BoilRatePerHour:     ToJSONVolumeType(i.BoilRatePerHour),
		KeyType:             &i.Type,
		EquipmentBaseForm:   ToJSONEquipmentBaseForm(i.Form),
		DrainRatePerMinute:  ToJSONVolumeType(i.DrainRatePerMinute),
		SpecificHeat:        ToJSONSpecificHeatType(i.SpecificHeat),
		GrainAbsorptionRate: ToJSONSpecificVolumeType(i.GrainAbsorptionRate),
		Name:                &i.Name,
		MaximumVolume:       ToJSONVolumeType(i.MaximumVolume),
		Weight:              ToJSONMassType(i.Weight),
		Loss:                *ToJSONVolumeType(i.Loss),
		Notes: &i.Notes,
	}, nil
}

func ToJSONSpecificHeatType(i *beerproto.SpecificHeatType) *beerjson.SpecificHeatType {
	if i == nil {
		return nil
	}
	
	if i.Unit == beerproto.SpecificHeatUnitType_NULL_SPECIFICHEATUNITTYPE {
		return &beerjson.SpecificHeatType{
		}
	}

	return &beerjson.SpecificHeatType{
		Value: i.Value,
		Unit:  *ToJSONSpecificHeatUnitType(i.Unit),
	}
}

func ToJSONSpecificHeatUnitType(i beerproto.SpecificHeatUnitType) *beerjson.SpecificHeatUnitType {
	if i == beerproto.SpecificHeatUnitType_NULL_SPECIFICHEATUNITTYPE {
		return nil
	}

	unit := beerproto.SpecificHeatUnitType_name[int32(i)]
	t := beerjson.SpecificHeatUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONEquipmentBaseForm(i beerproto.EquipmentItemType_EquipmentBaseForm) *beerjson.EquipmentBaseForm {
	if i == beerproto.EquipmentItemType_NULL {
		return nil
	}

	var t beerjson.EquipmentBaseForm
	switch i {
	case beerproto.EquipmentItemType_HLT:
		t = beerjson.EquipmentBaseForm_HLT
	case beerproto.EquipmentItemType_MASH_TUN:
		t = beerjson.EquipmentBaseForm_MashTun
	case beerproto.EquipmentItemType_LAUTER_TUN:
		t = beerjson.EquipmentBaseForm_LauterTun
	case beerproto.EquipmentItemType_BREW_KETTLE:
		t = beerjson.EquipmentBaseForm_BrewKettle
	case beerproto.EquipmentItemType_FERMENTER:
		t = beerjson.EquipmentBaseForm_Fermenter
	case beerproto.EquipmentItemType_AGING_VESSEL:
		t = beerjson.EquipmentBaseForm_AgingVessel
	case beerproto.EquipmentItemType_PACKAGING_VESSEL:
		t = beerjson.EquipmentBaseForm_PackagingVessel
	}

	return &t
}

func ToJSONCultureInformation(i *beerproto.CultureInformation) (*beerjson.CultureInformation, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type == beerproto.CultureBaseType_NULL_CULTUREBASETYPE {
		return nil, fmt.Errorf("type is required")
	}


	if i.Form == beerproto.CultureBaseForm_NULL_CULTUREBASEFORM {
		return nil, fmt.Errorf("form is required")
	}

	temperatureRange, err := ToJSONTemperatureRangeType(i.TemperatureRange)
	if err != nil {
		return nil, err
	}

	attenuationRange, err := ToJSONPercentRangeType(i.AttenuationRange)
	if err != nil {
		return nil, err
	}

	return &beerjson.CultureInformation{
		CultureBaseForm:  ToJSONCultureBaseForm(i.Form),
		Producer:         &i.Producer,
		TemperatureRange: temperatureRange,
		Notes:            &i.Notes,
		BestFor:          &i.BestFor,
		Inventory:        ToJSONCultureInventoryType(i.Inventory),
		ProductId:        &i.ProductId,
		Name:             &i.Name,
		AlcoholTolerance: ToJSONPercentType(i.AlcoholTolerance),
		Glucoamylase:     &i.Glucoamylase,
		CultureBaseType:  ToJSONCultureBaseType(i.Type),
		Flocculation:     ToJSONQualitativeRangeType(i.Flocculation),
		AttenuationRange: attenuationRange,
		MaxReuse:         &i.MaxReuse,
		Pof:              &i.Pof,
		Zymocide:         ToJSONZymocide(i.Zymocide),
	}, nil
}

func ToJSONZymocide(i *beerproto.Zymocide) *beerjson.Zymocide {
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

func ToJSONQualitativeRangeType(i beerproto.QualitativeRangeType) *beerjson.QualitativeRangeType {
	if i == beerproto.QualitativeRangeType_NULL_QUALITATIVERANGETYPE {
		return nil
	}

	unit := beerproto.QualitativeRangeType_name[int32(i)]
	t := beerjson.QualitativeRangeType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureBaseType(i beerproto.CultureBaseType) *beerjson.CultureBaseType {
	if i == beerproto.CultureBaseType_NULL_CULTUREBASETYPE {
		return nil
	}

	unit := beerproto.CultureBaseType_name[int32(i)]
	t := beerjson.CultureBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONCultureInventoryType(i *beerproto.CultureInventoryType) *beerjson.CultureInventoryType {
	if i == nil {
		return nil
	}
	return &beerjson.CultureInventoryType{
		Liquid:  ToJSONVolumeType(i.Liquid),
		Dry:     ToJSONMassType(i.Dry),
		Slant:   ToJSONVolumeType(i.Slant),
		Culture: ToJSONVolumeType(i.Culture),
	}
}

func ToJSONTemperatureRangeType(i *beerproto.TemperatureRangeType) (*beerjson.TemperatureRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}

	return  &beerjson.TemperatureRangeType{
		Minimum: *ToJSONTemperatureType(i.Minimum),
		Maximum: *ToJSONTemperatureType(i.Maximum),
	}, nil
}

func ToJSONFermentableType(i *beerproto.FermentableType) (*beerjson.FermentableType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type == beerproto.FermentableBaseType_NULL_FERMENTABLEBASETYPE {
		return nil, fmt.Errorf("type is required")
	}

	if i.Yield == nil {
		return nil, fmt.Errorf("yield is required")
	}

	if i.Color == nil {
		return nil, fmt.Errorf("color is required")
	}

	return &beerjson.FermentableType{
		MaxInBatch:                ToJSONPercentType(i.MaxInBatch),
		RecommendMash:             &i.RecommendMash,
		Protein:                   ToJSONPercentType(i.Protein),
		ProductId:                 &i.ProductId,
		FermentableBaseGrainGroup: ToJSONGrainGroup(i.GrainGroup),
		Yield:                     ToJSONYieldType(i.Yield),
		FermentableBaseType:       ToJSONFermentableBaseType(i.Type),
		Producer:                  &i.Producer,
		AlphaAmylase:              &i.AlphaAmylase,
		Color:                     ToJSONColorType(i.Color),
		Name:                      &i.Name,
		DiastaticPower:            ToJSONDiastaticPowerType(i.DiastaticPower),
		Moisture:                  ToJSONPercentType(i.Moisture),
		Origin:                    &i.Origin,
		Inventory:                 ToJSONFermentableInventoryType(i.Inventory),
		KolbachIndex:              &i.KolbachIndex,
		Notes:                     &i.Notes,
		Glassy: ToJSONPercentType(i.Glassy),
		Plump: ToJSONPercentType(i.Plump),
		Half: ToJSONPercentType(i.Half),
		Mealy: ToJSONPercentType(i.Mealy),
		Thru: ToJSONPercentType(i.Thru),
		Friability: ToJSONPercentType(i.Friability),
		DipH: ToJSONAcidityType(i.DiPh),
		Viscosity:ToJSONViscosityType(i.Viscosity),
		DMSP: ToJSONConcentrationType(i.DmsP),
		FAN: ToJSONConcentrationType(i.Fan),
		Fermentability: ToJSONPercentType(i.Fermentability),
		BetaGlucan: ToJSONConcentrationType(i.BetaGlucan),
	}, nil
}

func ToJSONViscosityType(i *beerproto.ViscosityType) *beerjson.ViscosityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.ViscosityUnitType_NULL_VISCOSITYUNITTYPE {
		return &beerjson.ViscosityType{	}
	}
	return &beerjson.ViscosityType{
		Value: i.Value,
		Unit:  *ToJSONViscosityUnitType(i.Unit),
	}
}

func ToJSONViscosityUnitType(i beerproto.ViscosityUnitType) *beerjson.ViscosityUnitType {
	if i == beerproto.ViscosityUnitType_NULL_VISCOSITYUNITTYPE {
		return nil
	}

	if i == beerproto.ViscosityUnitType_CP {
		t := beerjson.ViscosityUnitType_CP
		return &t
	}

	if i == beerproto.ViscosityUnitType_MPAS {
		t := beerjson.ViscosityUnitType_MPAS
		return &t
	}

	if unit, ok := beerproto.ViscosityUnitType_name[int32(i)]; ok {
		t := beerjson.ViscosityUnitType(strings.ToLower(unit))
		return &t
	}

	return nil
}

func ToJSONFermentableInventoryType(i *beerproto.FermentableInventoryType) *beerjson.FermentableInventoryType {
	if i == nil {
		return nil
	}

	fermentableInventoryType := &beerjson.FermentableInventoryType{}

	if mass, ok := i.Amount.(*beerproto.FermentableInventoryType_Mass); ok {
		fermentableInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.FermentableInventoryType_Volume); ok {
		fermentableInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return fermentableInventoryType
}

func ToJSONDiastaticPowerType(i *beerproto.DiastaticPowerType) *beerjson.DiastaticPowerType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.DiastaticPowerUnitType_NULL_DIASTATICPOWERUNITTYPE {
		return &beerjson.DiastaticPowerType{}
	}

	return &beerjson.DiastaticPowerType{
		Value: i.Value,
		Unit:  ToJSONDiastaticPowerUnitType(i.Unit),
	}
}

func ToJSONDiastaticPowerUnitType(i beerproto.DiastaticPowerUnitType) beerjson.DiastaticPowerUnitType {
	if i == beerproto.DiastaticPowerUnitType_NULL_DIASTATICPOWERUNITTYPE {
		return beerjson.DiastaticPowerUnitType_Lintner
	}

	if i == beerproto.DiastaticPowerUnitType_WK {
		return beerjson.DiastaticPowerUnitType_WK
	}

	unit := beerproto.DiastaticPowerUnitType_name[int32(i)]
	return beerjson.DiastaticPowerUnitType(strings.Title(strings.ToLower(unit)))
}

func ToJSONStyleType(i *beerproto.StyleType) (*beerjson.StyleType, error) {
	if i == nil {
		return nil, nil
	}


	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Category == "" {
		return nil, fmt.Errorf("category is required")
	}

	if i.StyleGuide == "" {
		return nil, fmt.Errorf("styleGuide is required")
	}

	if i.Type == beerproto.StyleType_NULL_STYLECATEGORIES {
		return nil, fmt.Errorf("type is required")
	}

	alcoholByVolume, err := ToJSONPercentRangeType(i.AlcoholByVolume)
	if err != nil {
		return nil, err
	}
	finalGravity, err := ToJSONGravityRangeType(i.FinalGravity)
	if err != nil {
		return nil, err
	}

	originalGravity, err := ToJSONGravityRangeType(i.OriginalGravity)
	if err != nil {
		return nil, err
	}

	color, err := ToJSONColorRangeType(i.Color)
	if err != nil {
		return nil, err
	}

	carbonation, err := ToJSONCarbonationRangeType(i.Carbonation)
	if err != nil {
		return nil, err
	}

	internationalBitternessUnits, err := ToJSONBitternessRangeType(i.InternationalBitternessUnits)
	if err != nil {
		return nil, err
	}

	return &beerjson.StyleType{
		Aroma:                        &i.Aroma,
		Ingredients:                  &i.Ingredients,
		CategoryNumber:               &i.CategoryNumber,
		Notes:                        &i.Notes,
		Flavor:                       &i.Flavor,
		Mouthfeel:                    &i.Mouthfeel,
		FinalGravity:                 finalGravity,
		StyleGuide:                   &i.StyleGuide,
		Color:                       color,
		OriginalGravity:             originalGravity,
		Examples:                     &i.Examples,
		Name:                         &i.Name,
		Carbonation:                  carbonation,
		AlcoholByVolume:              alcoholByVolume,
		InternationalBitternessUnits: internationalBitternessUnits,
		Appearance:                   &i.Appearance,
		Category:                     &i.Category,
		StyleLetter:                  &i.StyleLetter,
		KeyType:                      ToJSONStyleType_StyleCategories(i.Type),
		OverallImpression:            &i.OverallImpression,
	}, nil
}

func ToJSONStyleType_StyleCategories(i beerproto.StyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerproto.StyleType_NULL_STYLECATEGORIES {
		return nil
	}

	unit := beerproto.StyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONBitternessRangeType(i *beerproto.BitternessRangeType) (*beerjson.BitternessRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}

	return &beerjson.BitternessRangeType{
		Minimum: *ToJSONBitternessType(i.Minimum),
		Maximum: *ToJSONBitternessType(i.Maximum),
	}, nil
}

func ToJSONBitternessType(i *beerproto.BitternessType) *beerjson.BitternessType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.BitternessType_NULL {
		return &beerjson.BitternessType{}
	}

	return &beerjson.BitternessType{
		Value: i.Value,
		Unit:  ToJSONBitternessUnitType(i.Unit),
	}
}

func ToJSONBitternessUnitType(i beerproto.BitternessType_BitternessUnitType) beerjson.BitternessUnitType {
	if i == beerproto.BitternessType_NULL {
		return beerjson.BitternessUnitType_IBUs
	}

	switch i {
	case beerproto.BitternessType_IBUs:
		return beerjson.BitternessUnitType_IBUs
	}

	return beerjson.BitternessUnitType_IBUs
}

func ToJSONPercentRangeType(i *beerproto.PercentRangeType) (*beerjson.PercentRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}

	return &beerjson.PercentRangeType{
		Minimum: *ToJSONPercentType(i.Minimum),
		Maximum: *ToJSONPercentType(i.Maximum),
	}, nil
}

func ToJSONCarbonationRangeType(i *beerproto.CarbonationRangeType) (*beerjson.CarbonationRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}


	return &beerjson.CarbonationRangeType{
		Minimum: *ToJSONCarbonationType(i.Minimum),
		Maximum: *ToJSONCarbonationType(i.Maximum),
	}, nil
}
func ToJSONCarbonationType(i *beerproto.CarbonationType) *beerjson.CarbonationType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.CarbonationUnitType_NULL_CARBONATIONUNITTYPE {
		return &beerjson.CarbonationType{
		}
	}

	return &beerjson.CarbonationType{
		Value: i.Value,
		Unit:  *ToJSONCarbonationUnitType(i.Unit),
	}
}

func ToJSONCarbonationUnitType(i beerproto.CarbonationUnitType) *beerjson.CarbonationUnitType {
	if i == beerproto.CarbonationUnitType_NULL_CARBONATIONUNITTYPE {
		return nil
	}

	unit := beerproto.CarbonationUnitType_name[int32(i)]
	t := beerjson.CarbonationUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONColorRangeType(i *beerproto.ColorRangeType) (*beerjson.ColorRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}


	return &beerjson.ColorRangeType{
		Minimum: *ToJSONColorType(i.Minimum),
		Maximum: *ToJSONColorType(i.Maximum),
	}, nil
}

func ToJSONGravityRangeType(i *beerproto.GravityRangeType) (*beerjson.GravityRangeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Minimum == nil {
		return nil, fmt.Errorf("minimum is required")
	}

	if i.Maximum == nil {
		return nil, fmt.Errorf("maximum is required")
	}

	return &beerjson.GravityRangeType{
		Minimum: *ToJSONGravityType(i.Minimum),
		Maximum: *ToJSONGravityType(i.Maximum),
	}, nil
}

func ToJSONMiscellaneousType(i *beerproto.MiscellaneousType) (*beerjson.MiscellaneousType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type == beerproto.MiscellaneousBaseType_NULL{
		return nil, fmt.Errorf("type is required")
	}

	miscellaneousType :=  &beerjson.MiscellaneousType{}

	if i.UseFor != "" {
		miscellaneousType.UseFor = &i.UseFor
	}

	if i.Notes != "" {
		miscellaneousType.Notes = &i.Notes
	}


	if i.Name != "" {
		miscellaneousType.Name = &i.Name
	}


	if i.Producer != "" {
		miscellaneousType.Producer = &i.Producer
	}

	if i.ProductId != "" {
		miscellaneousType.ProductId = &i.ProductId
	}

	if i.Type != beerproto.MiscellaneousBaseType_NULL {
		miscellaneousType.MiscellaneousBaseType = ToJSONMiscellaneousBaseType(i.Type)
	}

	if i.Inventory != nil{
		inventory, err := ToJSONMiscellaneousInventoryType(i.Inventory)
		if err != nil {
			return nil, err
		}
		miscellaneousType.Inventory = inventory
	}

	return miscellaneousType, nil
}

func ToJSONMiscellaneousInventoryType(i *beerproto.MiscellaneousInventoryType) (*beerjson.MiscellaneousInventoryType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Amount == nil {
		return nil, fmt.Errorf("amount is required")
	}

	miscellaneousInventoryType := &beerjson.MiscellaneousInventoryType{}

	if mass, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Mass); ok {
		miscellaneousInventoryType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Unit); ok {
		miscellaneousInventoryType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.MiscellaneousInventoryType_Volume); ok {
		miscellaneousInventoryType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return miscellaneousInventoryType, nil
}

func ToJSONRecipeType(i *beerproto.RecipeType) (*beerjson.RecipeType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type != beerproto.RecipeType_NULL_RECIPETYPETYPE {
		return nil, fmt.Errorf("type is required")
	}

	if i.Author == "" {
		return nil, fmt.Errorf("author is required")
	}

	if i.BatchSize == nil {
		return nil, fmt.Errorf("batchSize is required")
	}

	if i.Efficiency == nil {
		return nil, fmt.Errorf("efficiency is required")
	}

	if i.Ingredients == nil {
		return nil, fmt.Errorf("ingredients is required")
	}

	var created beerjson.DateType
	if i.Created != "" {
		created = beerjson.DateType(i.Created)
	}
	efficiency, err := ToJSONEfficiencyType(i.Efficiency)
	if err != nil {
		return nil, err
	}

	ingredients, err := ToJSONIngredientsType(i.Ingredients)
	if err != nil {
		return nil, err
	}

	batchSize := ToJSONVolumeType(i.BatchSize)

	boil, err := ToJSONBoilProcedureType(i.Boil)
	if err != nil {
		return nil, err
	}

	fermentationProcedureType, err := ToJSONFermentationProcedureType(i.Fermentation)
	if err != nil {
		return nil, err
	}

	mash, err := ToJSONMashProcedureType(i.Mash)
	if err != nil {
		return nil, err
	}

	taste, err := ToJSONTasteType(i.Taste)
	if err != nil {
		return nil, err
	}

	packaging, err := ToJSONPackagingProcedureType(i.Packaging)
	if err != nil {
		return nil, err
	}

	style, err := ToJSONRecipeStyleType(i.Style)
	if err != nil {
		return nil, err
	}
	return &beerjson.RecipeType{
		Efficiency:         *efficiency,
		Style:               style,
		IbuEstimate:         ToJSONIBUEstimateType(i.IbuEstimate),
		ColorEstimate:       ToJSONColorType(i.ColorEstimate),
		BeerPH:              ToJSONAcidityType(i.BeerPh),
		Name:                i.Name,
		RecipeTypeType:      ToJSONRecipeTypeType(i.Type),
		Coauthor:            &i.Coauthor,
		OriginalGravity:     ToJSONGravityType(i.OriginalGravity),
		FinalGravity:        ToJSONGravityType(i.FinalGravity),
		Carbonation:         &i.Carbonation,
		Fermentation:        fermentationProcedureType,
		Author:              i.Author,
		Ingredients:         *ingredients,
		Mash:               mash,
		Packaging:           packaging,
		Boil:                boil,
		Taste:               taste,
		CaloriesPerPint:     &i.CaloriesPerPint,
		Created:             &created,
		BatchSize:           *batchSize,
		Notes:               &i.Notes,
		AlcoholByVolume:     ToJSONPercentType(i.AlcoholByVolume),
		ApparentAttenuation: ToJSONPercentType(i.ApparentAttenuation),
	}, nil
}

func ToJSONTasteType(i *beerproto.TasteType) (*beerjson.TasteType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Notes == "" {
		return nil, fmt.Errorf("notes is required")
	}

	return &beerjson.TasteType{
		Notes:  i.Notes,
		Rating: i.Rating,
	}, nil
}

func ToJSONBoilProcedureType(i *beerproto.BoilProcedureType) (*beerjson.BoilProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	boilProcedureType :=  &beerjson.BoilProcedureType{}

	if i.BoilTime == nil {
		return nil, fmt.Errorf("boilTime is required")
	}

	boilProcedureType.BoilTime = *ToJSONTimeType(i.BoilTime)

	if i.BoilSteps != nil {
		boilSteps := make([]beerjson.BoilStepType, 0)

		for _, step := range i.BoilSteps {
			if boilStep, err := ToJSONBoilStepType(step); err == nil {
				boilSteps = append(boilSteps, *boilStep)
			} else {
				return nil, err
			}
		}
		boilProcedureType.BoilSteps = boilSteps
	}

	if i.PreBoilSize != nil {
		boilProcedureType.PreBoilSize = ToJSONVolumeType(i.PreBoilSize)
	}

	if i.Name != "" {
		boilProcedureType.Name = &i.Name
	}

	if i.Description != "" {
		boilProcedureType.Description = &i.Description
	}

	if i.Notes != "" {
		boilProcedureType.Notes = &i.Notes
	}

	return boilProcedureType, nil
}

func ToJSONBoilStepType(i *beerproto.BoilStepType) (*beerjson.BoilStepType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return &beerjson.BoilStepType{
		EndGravity:               ToJSONGravityType(i.EndGravity),
		BoilStepTypeChillingType: ToJSONChillingType(i.ChillingType),
		Description:              &i.Description,
		EndTemperature:           ToJSONTemperatureType(i.EndTemperature),
		RampTime:                 ToJSONTimeType(i.RampTime),
		StepTime:                 ToJSONTimeType(i.StepTime),
		StartGravity:             ToJSONGravityType(i.StartGravity),
		StartPh:                  ToJSONAcidityType(i.StartPh),
		EndPh:                    ToJSONAcidityType(i.EndPh),
		Name:                     i.Name,
		StartTemperature:         ToJSONTemperatureType(i.StartTemperature),
	}, nil
}

func ToJSONChillingType(i beerproto.BoilStepTypeChillingType) *beerjson.BoilStepTypeChillingType {
	if i == beerproto.BoilStepTypeChillingType_NULL_BOILSTEPTYPECHILLINGTYPE {
		return nil
	}

	unit := beerproto.BoilStepTypeChillingType_name[int32(i)]
	t := beerjson.BoilStepTypeChillingType(strings.ToLower(unit))
	return &t
}

func ToJSONPackagingProcedureType(i *beerproto.PackagingProcedureType) (*beerjson.PackagingProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	packagingProcedureType := &beerjson.PackagingProcedureType{}

	if i.PackagingVessels != nil {
		packagingVessels := make([]beerjson.PackagingVesselType, 0)

		for _, vessels := range i.PackagingVessels {
			if packagingVesselType, err := ToJSONPackagingVesselType(vessels); err == nil {
				packagingVessels = append(packagingVessels, *packagingVesselType)
			} else {
				return nil, err
			}
		}

		packagingProcedureType.PackagingVessels = packagingVessels
	}

	if i.Name != "" {
		packagingProcedureType.Name = i.Name
	}

	if i.Description != "" {
		packagingProcedureType.Description = &i.Description
	}

	if i.Notes != "" {
		packagingProcedureType.Notes = &i.Notes
	}

	if i.PackagedVolume != nil {
		packagingProcedureType.PackagedVolume = ToJSONVolumeType(i.PackagedVolume)
	}

	return packagingProcedureType, nil
}

func ToJSONPackagingVesselType(i *beerproto.PackagingVesselType) (*beerjson.PackagingVesselType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	var packageDate beerjson.DateType
	if i.PackageDate != "" {
		packageDate = beerjson.DateType(i.PackageDate)
	}
	return &beerjson.PackagingVesselType{
		PackagingVesselTypeType: ToJSONPackagingVesselTypeType(i.Type),
		StartGravity:            ToJSONGravityType(i.StartGravity),
		Name:                    i.Name,
		PackageDate:             &packageDate,
		StepTime:                ToJSONTimeType(i.StepTime),
		EndGravity:              ToJSONGravityType(i.EndGravity),
		VesselVolume:            ToJSONVolumeType(i.VesselVolume),
		VesselQuantity:          &i.VesselQuantity,
		Description:             &i.Description,
		StartPh:                 ToJSONAcidityType(i.StartPh),
		Carbonation:             &i.Carbonation,
		StartTemperature:        ToJSONTemperatureType(i.StartTemperature),
		EndPh:                   ToJSONAcidityType(i.EndPh),
		EndTemperature:          ToJSONTemperatureType(i.EndTemperature),
	}, nil
}

func ToJSONPackagingVesselTypeType(i beerproto.PackagingVesselType_PackagingVesselTypeType) *beerjson.PackagingVesselTypeType {
	if i == beerproto.PackagingVesselType_NULL_PACKAGINGVESSELTYPETYPE {
		return nil
	}

	unit := beerproto.PackagingVesselType_PackagingVesselTypeType_name[int32(i)]
	t := beerjson.PackagingVesselTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONIngredientsType(i *beerproto.IngredientsType) (*beerjson.IngredientsType, error) {
	if i == nil {
		return nil, nil
	}

	ingredientsType := &beerjson.IngredientsType{}

	if i.FermentableAdditions == nil {
		return nil, fmt.Errorf("fermentableAdditions is required")
	}

	fermentableAdditions := make([]beerjson.FermentableAdditionType, 0)
	for _, fermentable := range i.FermentableAdditions {
		if addition, err := ToJSONFermentableAdditionType(fermentable); err == nil {
			fermentableAdditions = append(fermentableAdditions, *addition)
		} else {
			return nil, err
		}
	}
	ingredientsType.FermentableAdditions = fermentableAdditions

	if i.MiscellaneousAdditions != nil {
		miscellaneousAdditions := make([]beerjson.MiscellaneousAdditionType, 0)
		for _, misc := range i.MiscellaneousAdditions {
			if misc, err := ToJSONMiscellaneousAdditionType(misc); err == nil {
				miscellaneousAdditions = append(miscellaneousAdditions, *misc)
			} else {
				return nil, err
			}
		}
		ingredientsType.MiscellaneousAdditions = miscellaneousAdditions
	}

	if i.CultureAdditions != nil {
		cultureAdditions := make([]beerjson.CultureAdditionType, 0)
		for _, culture := range i.CultureAdditions {
			if cultureAdditionType, err := ToJSONCultureAdditionType(culture); err == nil {
				cultureAdditions = append(cultureAdditions, *cultureAdditionType)
			} else {
				return nil, err
			}
		}
		ingredientsType.CultureAdditions = cultureAdditions
	}

	if i.WaterAdditions != nil {
		waterAdditions := make([]beerjson.WaterAdditionType, 0)
		for _, water := range i.WaterAdditions {
			if waterAdditionType, err := ToJSONWaterAdditionType(water); err == nil {
				waterAdditions = append(waterAdditions, *waterAdditionType)
			} else {
				return nil, err
			}
		}
		ingredientsType.WaterAdditions = waterAdditions
	}

	if i.HopAdditions != nil {
		hopAdditions := make([]beerjson.HopAdditionType, 0)
		for _, hop := range i.HopAdditions {
			if hop, err := ToJSONHopAdditionType(hop); err == nil {
				hopAdditions = append(hopAdditions, *hop)
			} else {
				return nil, err
			}
		}
		ingredientsType.HopAdditions = hopAdditions
	}


	return ingredientsType, nil
}

func ToJSONHopAdditionType(i *beerproto.HopAdditionType) (*beerjson.HopAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Timing == nil {
		return nil, fmt.Errorf("timing is required")
	}

	if i.Amount == nil {
		return nil, fmt.Errorf("amount is required")
	}

	hopAdditionType := &beerjson.HopAdditionType{
		BetaAcid:           ToJSONPercentType(i.BetaAcid),
		Producer:           &i.Producer,
		Origin:             &i.Origin,
		Year:               &i.Year,
		HopVarietyBaseForm: ToJSONHopVarietyBaseForm(i.Form),
		Timing:             *ToJSONTimingType(i.Timing),
		Name:               &i.Name,
		ProductId:          &i.ProductId,
		AlphaAcid:          ToJSONPercentType(i.AlphaAcid),
	}

	if mass, ok := i.Amount.(*beerproto.HopAdditionType_Mass); ok {
		hopAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.HopAdditionType_Volume); ok {
		hopAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return hopAdditionType, nil
}

func ToJSONHopVarietyBaseForm(i beerproto.HopVarietyBaseForm) *beerjson.HopVarietyBaseForm {
	if i == beerproto.HopVarietyBaseForm_NULL_HOPVARIETYBASEFORM {
		return nil
	}

	var t beerjson.HopVarietyBaseForm

	switch i {
	case beerproto.HopVarietyBaseForm_EXTRACT_HOPVARIETYBASEFORM:
		t = beerjson.HopVarietyBaseForm_Extract
	case beerproto.HopVarietyBaseForm_LEAF:
		t = beerjson.HopVarietyBaseForm_Leaf
	case beerproto.HopVarietyBaseForm_LEAFWET:
		t = beerjson.HopVarietyBaseForm_LeafWet
	case beerproto.HopVarietyBaseForm_PELLET:
		t = beerjson.HopVarietyBaseForm_Pellet
	case beerproto.HopVarietyBaseForm_POWDER:
		t = beerjson.HopVarietyBaseForm_Powder
	case beerproto.HopVarietyBaseForm_PLUG:
		t = beerjson.HopVarietyBaseForm_Plug
	}

	return &t
}

func ToJSONFermentableAdditionType(i *beerproto.FermentableAdditionType) (*beerjson.FermentableAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type == beerproto.FermentableBaseType_NULL_FERMENTABLEBASETYPE {
		return nil, fmt.Errorf("type is required")
	}

	if i.Yield == nil {
		return nil, fmt.Errorf("yield is required")
	}

	if i.Color == nil {
		return nil, fmt.Errorf("color is required")
	}

	if i.Timing == nil {
		return nil, fmt.Errorf("timing is required")
	}

	if i.Amount == nil {
		return nil, fmt.Errorf("amount is required")
	}

	fermentableAdditionType := &beerjson.FermentableAdditionType{
		FermentableBaseType:       ToJSONFermentableBaseType(i.Type),
		Origin:                    &i.Origin,
		FermentableBaseGrainGroup: ToJSONGrainGroup(i.GrainGroup),
		Yield:                     ToJSONYieldType(i.Yield),
		Color:                     ToJSONColorType(i.Color),
		Name:                      &i.Name,
		Producer:                  &i.Producer,
		ProductId:                 &i.ProductId,
		Timing:                    ToJSONTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerproto.FermentableAdditionType_Mass); ok {
		fermentableAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if volume, ok := i.Amount.(*beerproto.FermentableAdditionType_Volume); ok {
		fermentableAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return fermentableAdditionType, nil
}

func ToJSONYieldType(i *beerproto.YieldType) *beerjson.YieldType {
	if i == nil {
		return nil
	}

	return &beerjson.YieldType{
		FineGrind:            ToJSONPercentType(i.FineGrind),
		CoarseGrind:          ToJSONPercentType(i.CoarseGrind),
		FineCoarseDifference: ToJSONPercentType(i.FineCoarseDifference),
		Potential:            ToJSONGravityType(i.Potential),
	}
}

func ToJSONGrainGroup(i beerproto.GrainGroup) *beerjson.FermentableBaseGrainGroup {
	if i == beerproto.GrainGroup_NULL_GRAINGROUP {
		return nil
	}

	unit := beerproto.GrainGroup_name[int32(i)]
	t := beerjson.FermentableBaseGrainGroup(strings.ToLower(unit))
	return &t
}

func ToJSONFermentableBaseType(i beerproto.FermentableBaseType) *beerjson.FermentableBaseType {
	if i == beerproto.FermentableBaseType_NULL_FERMENTABLEBASETYPE {
		return nil
	}

	unit := beerproto.FermentableBaseType_name[int32(i)]
	t := beerjson.FermentableBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONWaterAdditionType(i *beerproto.WaterAdditionType) (*beerjson.WaterAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Calcium == nil {
		return nil, fmt.Errorf("calcium is required")
	}

	if i.Bicarbonate == nil {
		return nil, fmt.Errorf("bicarbonate is required")
	}

	if i.Sulfate == nil {
		return nil, fmt.Errorf("sulfate is required")
	}

	if i.Chloride == nil {
		return nil, fmt.Errorf("chloride is required")
	}

	if i.Sodium == nil {
		return nil, fmt.Errorf("sodium is required")
	}

	if i.Magnesium == nil {
		return nil, fmt.Errorf("magnesium is required")
	}

	return &beerjson.WaterAdditionType{
		Flouride:    ToJSONConcentrationType(i.Flouride),
		Sulfate:     ToJSONConcentrationType(i.Sulfate),
		Producer:    &i.Producer,
		Nitrate:     ToJSONConcentrationType(i.Nitrate),
		Nitrite:     ToJSONConcentrationType(i.Nitrite),
		Chloride:    ToJSONConcentrationType(i.Chloride),
		Amount:      ToJSONVolumeType(i.Amount),
		Name:        &i.Name,
		Potassium:   ToJSONConcentrationType(i.Potassium),
		Magnesium:   ToJSONConcentrationType(i.Magnesium),
		Iron:        ToJSONConcentrationType(i.Iron),
		Bicarbonate: ToJSONConcentrationType(i.Bicarbonate),
		Calcium:     ToJSONConcentrationType(i.Calcium),
		Carbonate:   ToJSONConcentrationType(i.Carbonate),
		Sodium:      ToJSONConcentrationType(i.Sodium),
	}, nil
}


func ToJSONConcentrationType(i *beerproto.ConcentrationType) *beerjson.ConcentrationType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.ConcentrationUnitType_NULL_CONCENTRATIONUNITTYPE {
		return &beerjson.ConcentrationType{	}
	}

	return &beerjson.ConcentrationType{
		Value: i.Value,
		Unit:  *ToJSONConcentrationUnitType(i.Unit),
	}
}

func ToJSONConcentrationUnitType(i beerproto.ConcentrationUnitType) *beerjson.ConcentrationUnitType {
	if i == beerproto.ConcentrationUnitType_NULL_CONCENTRATIONUNITTYPE {
		return nil
	}

	if i == beerproto.ConcentrationUnitType_MGL {
		t := beerjson.ConcentrationUnitType_MgL
		return &t
	}

	if unit, ok := beerproto.ConcentrationUnitType_name[int32(i)]; ok {
		t := beerjson.ConcentrationUnitType(strings.ToLower(unit))
		return &t
	}


	return nil
}

func ToJSONCultureAdditionType(i *beerproto.CultureAdditionType) (*beerjson.CultureAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if i.Type == beerproto.CultureBaseType_NULL_CULTUREBASETYPE {
		return nil, fmt.Errorf("type is required")
	}

	if i.Form == beerproto.CultureBaseForm_NULL_CULTUREBASEFORM {
		return nil, fmt.Errorf("form is required")
	}

	cultureAdditionType := &beerjson.CultureAdditionType{
		CultureBaseForm:   ToJSONCultureBaseForm(i.Form),
		ProductId:         &i.ProductId,
		Name:              &i.Name,
		CellCountBillions: &i.CellCountBillions,
		TimesCultured:     &i.TimesCultured,
		Producer:          &i.Producer,
		CultureBaseType:   ToJSONCultureBaseType(i.Type),
		Attenuation:       ToJSONPercentType(i.Attenuation),
		Timing:            ToJSONTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerproto.CultureAdditionType_Mass); ok {
		cultureAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.CultureAdditionType_Unit); ok {
		cultureAdditionType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.CultureAdditionType_Volume); ok {
		cultureAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return cultureAdditionType, nil
}

func ToJSONCultureBaseForm(i beerproto.CultureBaseForm) *beerjson.CultureBaseForm {
	if i == beerproto.CultureBaseForm_NULL_CULTUREBASEFORM {
		return nil
	}

	unit := beerproto.CultureBaseForm_name[int32(i)]
	t := beerjson.CultureBaseForm(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousAdditionType(i *beerproto.MiscellaneousAdditionType) (*beerjson.MiscellaneousAdditionType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Type == beerproto.MiscellaneousBaseType_NULL {
		return nil, fmt.Errorf("type is required")
	}

	if i.Timing == nil {
		return nil, fmt.Errorf("timing is required")
	}

	if i.Amount == nil {
		return nil, fmt.Errorf("amount is required")
	}

	miscellaneousAdditionType := &beerjson.MiscellaneousAdditionType{
		Name:                  &i.Name,
		Producer:              &i.Producer,
		Timing:                ToJSONTimingType(i.Timing),
		ProductId:             &i.ProductId,
		MiscellaneousBaseType: ToJSONMiscellaneousBaseType(i.Type),
	}

	if mass, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Mass); ok {
		miscellaneousAdditionType.Amount = ToJSONMassType(mass.Mass)
	}

	if unit, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Unit); ok {
		miscellaneousAdditionType.Amount = ToJSONUnitType(unit.Unit)
	}
	if volume, ok := i.Amount.(*beerproto.MiscellaneousAdditionType_Volume); ok {
		miscellaneousAdditionType.Amount = ToJSONVolumeType(volume.Volume)
	}

	return miscellaneousAdditionType, nil
}

func ToJSONUnitType(i *beerproto.UnitType) *beerjson.UnitType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.UnitUnitType_NULL_UNITUNITTYPE {
		return &beerjson.UnitType{	}
	}

	return &beerjson.UnitType{
		Value: i.Value,
		Unit:  *ToJSONUnitUnitType(i.Unit),
	}
}

func ToJSONUnitUnitType(i beerproto.UnitUnitType) *beerjson.UnitUnitType {
	if i == beerproto.UnitUnitType_NULL_UNITUNITTYPE {
		return nil
	}

	unit := beerproto.UnitUnitType_name[int32(i)]
	t := beerjson.UnitUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMassType(i *beerproto.MassType) *beerjson.MassType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.MassUnitType_NULL_MASSUNITTYPE {
		return &beerjson.MassType{}
	}

	return &beerjson.MassType{
		Value: i.Value,
		Unit:  *ToJSONMassUnitType(i.Unit),
	}
}

func ToJSONMassUnitType(i beerproto.MassUnitType) *beerjson.MassUnitType {
	if i == beerproto.MassUnitType_NULL_MASSUNITTYPE {
		return nil
	}

	unit := beerproto.MassUnitType_name[int32(i)]
	t := beerjson.MassUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONMiscellaneousBaseType(i beerproto.MiscellaneousBaseType) *beerjson.MiscellaneousBaseType {
	if i == beerproto.MiscellaneousBaseType_NULL {
		return nil
	}

	unit := beerproto.MiscellaneousBaseType_name[int32(i)]
	t := beerjson.MiscellaneousBaseType(strings.ToLower(unit))
	return &t
}

func ToJSONTimingType(i *beerproto.TimingType) *beerjson.TimingType {
	if i == nil {
		return nil
	}
	return &beerjson.TimingType{
		Time:            ToJSONTimeType(i.Time),
		Duration:        ToJSONTimeType(i.Duration),
		Continuous:      &i.Continuous,
		SpecificGravity: ToJSONGravityType(i.SpecificGravity),
		PH:              ToJSONAcidityType(i.Ph),
		Step:            &i.Step,
		Use:             ToJSONUseType(i.Use),
	}
}

func ToJSONUseType(i beerproto.UseType) *beerjson.UseType {
	if i == beerproto.UseType_NULL_USETYPE {
		return nil
	}

	unit := beerproto.UseType_name[int32(i)]
	t := beerjson.UseType(strings.ToLower(unit))
	return &t
}

func ToJSONFermentationProcedureType(i *beerproto.FermentationProcedureType) (*beerjson.FermentationProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.FermentationSteps == nil {
		return nil, fmt.Errorf("fermentationSteps is required")
	}

	fermentationProcedureType := &beerjson.FermentationProcedureType{}

	fermentationProcedureType.Name = i.Name
	if i.FermentationSteps != nil {
		steps := make([]beerjson.FermentationStepType, 0)
		for _, step := range i.FermentationSteps {
			if fermentationStepType, err := ToJSONFermentationStepType(step); err == nil {
				steps = append(steps, *fermentationStepType)
			} else {
				return nil, err
			}
		}
		fermentationProcedureType.FermentationSteps = steps
	}

	if i.Description != "" {
		fermentationProcedureType.Description = &i.Description
	}

	if i.Notes != "" {
		fermentationProcedureType.Notes = &i.Notes
	}

	return fermentationProcedureType, nil
}

func ToJSONFermentationStepType(i *beerproto.FermentationStepType) (*beerjson.FermentationStepType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return &beerjson.FermentationStepType{
		Name:             i.Name,
		EndTemperature:   ToJSONTemperatureType(i.EndTemperature),
		StepTime:         ToJSONTimeType(i.StepTime),
		FreeRise:         &i.FreeRise,
		StartGravity:     ToJSONGravityType(i.StartGravity),
		EndGravity:       ToJSONGravityType(i.EndGravity),
		StartPh:          ToJSONAcidityType(i.StartPh),
		Description:      &i.Description,
		StartTemperature: ToJSONTemperatureType(i.StartTemperature),
		EndPh:            ToJSONAcidityType(i.EndPh),
		Vessel:           &i.Vessel,
	}, nil
}

func ToJSONGravityType(i *beerproto.GravityType) *beerjson.GravityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.GravityUnitType_NULL_GRAVITYUNITTYPE {
		return &beerjson.GravityType{}
	}

	return &beerjson.GravityType{
		Value: i.Value,
		Unit:  *ToJSONGravityUnitType(i.Unit),
	}
}

func ToJSONGravityUnitType(i beerproto.GravityUnitType) *beerjson.GravityUnitType {
	if i == beerproto.GravityUnitType_NULL_GRAVITYUNITTYPE {
		return nil
	}

	unit := beerproto.GravityUnitType_name[int32(i)]
	t := beerjson.GravityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONRecipeTypeType(i beerproto.RecipeType_RecipeTypeType) beerjson.RecipeTypeType {
	if i == beerproto.RecipeType_NULL_RECIPETYPETYPE {
		return beerjson.RecipeTypeType_AllGrain
	}

	return beerjson.RecipeTypeType(strings.ToLower(i.String()))
}

func ToJSONColorType(i *beerproto.ColorType) *beerjson.ColorType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.ColorUnitType_NULL_COLORUNITTYPE {
		return &beerjson.ColorType{	}
	}

	return &beerjson.ColorType{
		Value: i.Value,
		Unit:  *ToJSONColorUnitType(i.Unit),
	}
}

func ToJSONColorUnitType(i beerproto.ColorUnitType) *beerjson.ColorUnitType {
	if i == beerproto.ColorUnitType_NULL_COLORUNITTYPE {
		return nil
	}

	unit := beerproto.ColorUnitType_name[int32(i)]

	if i == beerproto.ColorUnitType_LOVI {
		unit = strings.Title(strings.ToLower(unit))
	}
	t := beerjson.ColorUnitType(unit)
	return &t
}

func ToJSONIBUEstimateType(i *beerproto.IBUEstimateType) *beerjson.IBUEstimateType {
	if i == nil {
		return nil
	}

	return &beerjson.IBUEstimateType{
		Method: ToJSONIBUMethodType(i.Method),
	}
}

func ToJSONIBUMethodType(i beerproto.IBUEstimateType_IBUMethodType) *beerjson.IBUMethodType {
	if i == beerproto.IBUEstimateType_NULL_IBUMETHODTYPE {
		return nil
	}

	unit := beerproto.IBUEstimateType_IBUMethodType_name[int32(i)]
	t := beerjson.IBUMethodType(strings.Title(strings.ToLower(unit)))
	return &t
}

func ToJSONRecipeStyleType(i *beerproto.RecipeStyleType) (*beerjson.RecipeStyleType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.Category == "" {
		return nil, fmt.Errorf("category is required")
	}

	if i.StyleGuide == "" {
		return nil, fmt.Errorf("styleGuide is required")
	}

	if i.Type == beerproto.RecipeStyleType_NULL_STYLECATEGORIES {
		return nil, fmt.Errorf("type is required")
	}

	return &beerjson.RecipeStyleType{
		KeyType:        ToJSONRecipeStyleType_StyleCategories(i.Type),
		Name:           &i.Name,
		Category:       &i.Category,
		CategoryNumber: &i.CategoryNumber,
		StyleGuide:     &i.StyleGuide,
		StyleLetter:    &i.StyleLetter,
	}, nil
}

func ToJSONRecipeStyleType_StyleCategories(i beerproto.RecipeStyleType_StyleCategories) *beerjson.StyleCategories {
	if i == beerproto.RecipeStyleType_NULL_STYLECATEGORIES {
		return nil
	}

	unit := beerproto.RecipeStyleType_StyleCategories_name[int32(i)]
	t := beerjson.StyleCategories(strings.ToLower(unit))
	return &t
}

func ToJSONEfficiencyType(i *beerproto.EfficiencyType) (*beerjson.EfficiencyType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Brewhouse == nil {
		return nil, fmt.Errorf("brewhouse is required")
	}

	return &beerjson.EfficiencyType{
		Conversion: ToJSONPercentType(i.Conversion),
		Lauter:     ToJSONPercentType(i.Lauter),
		Mash:       ToJSONPercentType(i.Mash),
		Brewhouse:  *ToJSONPercentType(i.Brewhouse),
	}, nil
}

func ToJSONPercentType(i *beerproto.PercentType) *beerjson.PercentType {
	if i == nil {
		return nil
	}

	return &beerjson.PercentType{
		Value: i.Value,
		Unit:  ToJSONPercentUnitType(i.Unit),
	}
}

func ToJSONPercentUnitType(i beerproto.PercentType_PercentUnitType) beerjson.PercentUnitType {
	if i == beerproto.PercentType_NULL {
		return beerjson.PercentUnitType_No
	}

	return beerjson.PercentUnitType(strings.ToLower(i.String()))
}

func ToJSONMashProcedureType(i *beerproto.MashProcedureType) (*beerjson.MashProcedureType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.MashSteps == nil {
		return nil, fmt.Errorf("mashSteps is required")
	}

	if i.GrainTemperature == nil {
		return nil, fmt.Errorf("grainTemperature is required")
	}
	mashProcedureType := &beerjson.MashProcedureType{}
	if i.MashSteps != nil {
		mashSteps := []beerjson.MashStepType{}
		for _, step := range i.MashSteps {
			if step, err := ToJSONMashStepType(step); err == nil {
				mashSteps = append(mashSteps, *step)
			} else {
				return nil, err
			}
		}
		mashProcedureType.MashSteps = mashSteps
	}

	mashProcedureType.GrainTemperature = *ToJSONTemperatureType(i.GrainTemperature)

	mashProcedureType.Name = i.Name


	if i.Notes != "" {
		mashProcedureType.Notes = &i.Notes
	}
	return mashProcedureType, nil
}

func ToJSONMashStepType(i *beerproto.MashStepType) (*beerjson.MashStepType, error) {
	if i == nil {
		return nil, nil
	}

	if i.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if i.StepTemperature == nil {
		return nil, fmt.Errorf("stepTemperature is required")
	}

	if i.StepTime == nil {
		return nil, fmt.Errorf("stepTime is required")
	}

	if i.Type == beerproto.MashStepType_NULL {
		return nil, fmt.Errorf("type is required")
	}

	mashStepType :=  &beerjson.MashStepType{
		StepTime:          *ToJSONTimeType(i.StepTime),
		RampTime:          ToJSONTimeType(i.RampTime),
		EndTemperature:    ToJSONTemperatureType(i.EndTemperature),
		InfuseTemperature: ToJSONTemperatureType(i.InfuseTemperature),
		StartPH:           ToJSONAcidityType(i.StartPh),
		EndPH:             ToJSONAcidityType(i.EndPh),
		Name:              i.Name,
		MashStepTypeType:  *ToJSONMashStepTypeType(i.Type),
		Amount:            ToJSONVolumeType(i.Amount),
		StepTemperature:   *ToJSONTemperatureType(i.StepTemperature),
		WaterGrainRatio:   ToJSONSpecificVolumeType(i.WaterGrainRatio),
	}

	if i.Description != "" {
		mashStepType.Description = &i.Description
	}

	return mashStepType, nil
}

func ToJSONVolumeType(i *beerproto.VolumeType) *beerjson.VolumeType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.VolumeType_NULL {
		return &beerjson.VolumeType{
		}
	}

	return &beerjson.VolumeType{
		Value: i.Value,
		Unit:  ToJSONVolumeUnitType(i.Unit),
	}
}

func ToJSONVolumeUnitType(i beerproto.VolumeType_VolumeUnitType) beerjson.VolumeUnitType {
	if i == beerproto.VolumeType_NULL {
		return beerjson.VolumeUnitType_Ml
	}

	return beerjson.VolumeUnitType(strings.ToLower(i.String()))
}

func ToJSONSpecificVolumeType(i *beerproto.SpecificVolumeType) *beerjson.SpecificVolumeType {
	if i == nil {
		return nil
	}
	if i.Unit == beerproto.SpecificVolumeType_NULL {
		return &beerjson.SpecificVolumeType{
		}
	}
	return &beerjson.SpecificVolumeType{
		Value: i.Value,
		Unit:  ToJSONSpecificVolumeUnitType(i.Unit),
	}
}

func ToJSONSpecificVolumeUnitType(i beerproto.SpecificVolumeType_SpecificVolumeUnitType) beerjson.SpecificVolumeUnitType {
	if i == beerproto.SpecificVolumeType_NULL {
		return beerjson.SpecificVolumeUnitType_LG
	}

	switch i {
	case beerproto.SpecificVolumeType_QTLB:
		return beerjson.SpecificVolumeUnitType_QtLb
	case beerproto.SpecificVolumeType_GALLB:
		return beerjson.SpecificVolumeUnitType_GalLb
	case beerproto.SpecificVolumeType_GALOZ:
		return beerjson.SpecificVolumeUnitType_GalOz
	case beerproto.SpecificVolumeType_LG:
		return beerjson.SpecificVolumeUnitType_LG
	case beerproto.SpecificVolumeType_LKG:
		return beerjson.SpecificVolumeUnitType_LKg
	case beerproto.SpecificVolumeType_FLOZOZ:
		return beerjson.SpecificVolumeUnitType_FlozOz
	case beerproto.SpecificVolumeType_M3KG:
		return beerjson.SpecificVolumeUnitType_M3Kg
	case beerproto.SpecificVolumeType_FT3LB:
		return beerjson.SpecificVolumeUnitType_Ft3Lb
	}

	return beerjson.SpecificVolumeUnitType(i.String())
}

func ToJSONMashStepTypeType(i beerproto.MashStepType_MashStepTypeType) *beerjson.MashStepTypeType {
	if i == beerproto.MashStepType_NULL {
		return nil
	}

	unit := beerproto.MashStepType_MashStepTypeType_name[int32(i)]
	t := beerjson.MashStepTypeType(strings.ToLower(unit))
	return &t
}

func ToJSONAcidityType(i *beerproto.AcidityType) *beerjson.AcidityType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.AcidityUnitType_NULL_ACIDITYUNITTYPE {
		return &beerjson.AcidityType{}
	}

	return &beerjson.AcidityType{
		Value: i.Value,
		Unit:  *ToJSONAcidityUnitType(i.Unit),
	}
}

func ToJSONAcidityUnitType(i beerproto.AcidityUnitType) *beerjson.AcidityUnitType {
	if i == beerproto.AcidityUnitType_NULL_ACIDITYUNITTYPE {
		return nil
	}

	unit := beerproto.AcidityUnitType_name[int32(i)]
	t := beerjson.AcidityUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTimeType(i *beerproto.TimeType) *beerjson.TimeType {
	if i == nil {
		return nil
	}

	if i.Unit == beerproto.TimeType_NULL {
		return &beerjson.TimeType{	}
	}

	return &beerjson.TimeType{
		Value: i.Value,
		Unit:  *ToJSONTimeUnitType(i.Unit),
	}
}

func ToJSONTemperatureType(i *beerproto.TemperatureType) *beerjson.TemperatureType {
	if i == nil {
		return nil
	}
	if i.Unit == beerproto.TemperatureUnitType_NULL_TEMPERATUREUNITTYPE {
		return &beerjson.TemperatureType{}
	}

	return &beerjson.TemperatureType{
		Value: i.Value,
		Unit:  *ToJSONTemperatureUnitType(i.Unit),
	}
}

func ToJSONTimeUnitType(i beerproto.TimeType_TimeUnitType) *beerjson.TimeUnitType {
	if i == beerproto.TimeType_NULL {
		return nil
	}

	unit := beerproto.TimeType_TimeUnitType_name[int32(i)]
	t := beerjson.TimeUnitType(strings.ToLower(unit))
	return &t
}

func ToJSONTemperatureUnitType(i beerproto.TemperatureUnitType) *beerjson.TemperatureUnitType {
	if i == beerproto.TemperatureUnitType_NULL_TEMPERATUREUNITTYPE {
		return nil
	}

	unit := beerproto.TemperatureUnitType_name[int32(i)]
	t := beerjson.TemperatureUnitType(unit)
	return &t
}
