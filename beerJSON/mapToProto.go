package beerJSON

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"

	"github.com/beerproto/beerjson.go"
)

func MapToProto(i *beerjson.Beerjson) *beerproto.Recipe {
	if i == nil {
		return nil
	}

	output := &beerproto.Recipe{
		Mashes:                   []*beerproto.MashProcedureType{},
		Recipes:                  []*beerproto.RecipeType{},
		MiscellaneousIngredients: []*beerproto.MiscellaneousType{},
		Styles:                   []*beerproto.StyleType{},
		Fermentations:            []*beerproto.FermentationProcedureType{},
		Boil:                     []*beerproto.BoilProcedureType{},
		Version:                  float64(i.Version),
		Fermentables:             []*beerproto.FermentableType{},
		Cultures:                 []*beerproto.CultureInformation{},
		Equipments:               []*beerproto.EquipmentType{},
		Packaging:                []*beerproto.PackagingProcedureType{},
		HopVarieties:             []*beerproto.VarietyInformation{},
		Profiles:                 []*beerproto.WaterBase{},
	}

	for _, mash := range i.Mashes {
		output.Mashes = append(output.Mashes, ToProtoMashProcedureType(&mash))
	}

	for _, recipe := range i.Recipes {
		output.Recipes = append(output.Recipes, ToProtoRecipeType(&recipe))
	}

	for _, ingredients := range i.MiscellaneousIngredients {
		output.MiscellaneousIngredients = append(output.MiscellaneousIngredients, ToProtoMiscellaneousType(&ingredients))
	}

	for _, style := range i.Styles {
		output.Styles = append(output.Styles, ToProtoStyleType(&style))
	}

	for _, fermentation := range i.Fermentations {
		output.Fermentations = append(output.Fermentations, ToProtoFermentationProcedureType(&fermentation))
	}

	for _, boil := range i.Boil {
		output.Boil = append(output.Boil, ToProtoBoilProcedureType(&boil))
	}

	for _, fermentable := range i.Fermentables {
		output.Fermentables = append(output.Fermentables, ToProtoFermentableType(&fermentable))
	}

	for _, culture := range i.Cultures {
		output.Cultures = append(output.Cultures, ToProtoCultureInformation(&culture))
	}

	for _, equipment := range i.Equipments {
		output.Equipments = append(output.Equipments, ToProtoEquipmentType(&equipment))
	}

	for _, packaging := range i.Packaging {
		output.Packaging = append(output.Packaging, ToProtoPackagingProcedureType(&packaging))
	}

	for _, hopVariety := range i.HopVarieties {
		output.HopVarieties = append(output.HopVarieties, ToProtoVarietyInformation(&hopVariety))
	}

	for _, profile := range i.Profiles {
		output.Profiles = append(output.Profiles, ToProtoWaterBase(&profile))
	}

	return output
}

func ToProtoCarbonationUnitType(i *beerjson.CarbonationUnitType) beerproto.CarbonationUnit {
	if i == nil {
		return beerproto.CarbonationUnit_CARBONATION_UNIT_UNSPECIFIED
	}
	unit := beerproto.CarbonationUnit_value[strings.ToUpper(string(*i))]
	return beerproto.CarbonationUnit(unit)
}

func ToProtoColorRangeType(i *beerjson.ColorRangeType) *beerproto.ColorRangeType {
	if i == nil {
		return nil
	}
	return &beerproto.ColorRangeType{
		Minimum: ToProtoColorType(&i.Minimum),
		Maximum: ToProtoColorType(&i.Maximum),
	}
}

func ToProtoGravityRangeType(i *beerjson.GravityRangeType) *beerproto.GravityRangeType {
	if i == nil {
		return nil
	}

	return &beerproto.GravityRangeType{
		Minimum: ToProtoGravityType(&i.Minimum),
		Maximum: ToProtoGravityType(&i.Maximum),
	}
}

func ToProtoMiscellaneousType(i *beerjson.MiscellaneousType) *beerproto.MiscellaneousType {
	if i == nil {
		return nil
	}

	miscellaneousType := &beerproto.MiscellaneousType{}

	if i.UseFor != nil {
		miscellaneousType.UseFor = UseString(i.UseFor)
	}

	if i.Notes != nil {
		miscellaneousType.Notes = UseString(i.Notes)
	}

	if i.Name != nil {
		miscellaneousType.Name = UseString(i.Name)
	}

	if i.Producer != nil {
		miscellaneousType.Producer = UseString(i.Producer)
	}

	if i.ProductId != nil {
		miscellaneousType.ProductId = UseString(i.ProductId)
	}

	if i.MiscellaneousBaseType != nil {
		miscellaneousType.Type = ToProtoMiscellaneousBaseType(i.MiscellaneousBaseType)
	}

	if i.Inventory != nil {
		miscellaneousType.Inventory = ToProtoMiscellaneousInventoryType(i.Inventory)
	}

	return miscellaneousType
}

func ToProtoMiscellaneousInventoryType(i *beerjson.MiscellaneousInventoryType) *beerproto.MiscellaneousInventoryType {
	if i == nil {
		return nil
	}

	miscellaneousInventoryType := &beerproto.MiscellaneousInventoryType{}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok := i.Amount.(*beerjson.UnitType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		miscellaneousInventoryType.Amount = &beerproto.MiscellaneousInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return miscellaneousInventoryType
}

func ToProtoRecipeType(i *beerjson.RecipeType) *beerproto.RecipeType {
	if i == nil {
		return nil
	}

	created := ""
	if i.Created != nil {
		created = string(*i.Created)
	}
	return &beerproto.RecipeType{
		Efficiency:          ToProtoEfficiencyType(&i.Efficiency),
		Style:               ToProtoRecipeStyleType(i.Style),
		IbuEstimate:         ToProtoIBUEstimateType(i.IbuEstimate),
		ColorEstimate:       ToProtoColorType(i.ColorEstimate),
		BeerPh:              ToProtoAcidityType(i.BeerPH),
		Name:                i.Name,
		Type:                ToProtoRecipeTypeType(i.RecipeTypeType),
		Coauthor:            UseString(i.Coauthor),
		OriginalGravity:     ToProtoGravityType(i.OriginalGravity),
		FinalGravity:        ToProtoGravityType(i.FinalGravity),
		Carbonation:         UseFloat(i.Carbonation),
		Fermentation:        ToProtoFermentationProcedureType(i.Fermentation),
		Author:              i.Author,
		Ingredients:         ToProtoIngredientsType(&i.Ingredients),
		Mash:                ToProtoMashProcedureType(i.Mash),
		Packaging:           ToProtoPackagingProcedureType(i.Packaging),
		Boil:                ToProtoBoilProcedureType(i.Boil),
		Taste:               ToProtoTasteType(i.Taste),
		CaloriesPerPint:     UseFloat(i.CaloriesPerPint),
		Created:             created,
		BatchSize:           ToProtoVolumeType(&i.BatchSize),
		Notes:               UseString(i.Notes),
		AlcoholByVolume:     ToProtoPercentType(i.AlcoholByVolume),
		ApparentAttenuation: ToProtoPercentType(i.ApparentAttenuation),
	}
}

func ToProtoTasteType(i *beerjson.TasteType) *beerproto.TasteType {
	if i == nil {
		return nil
	}
	return &beerproto.TasteType{
		Notes:  i.Notes,
		Rating: i.Rating,
	}
}

func ToProtoBoilProcedureType(i *beerjson.BoilProcedureType) *beerproto.BoilProcedureType {
	if i == nil {
		return nil
	}

	boilProcedureType := &beerproto.BoilProcedureType{}

	if i.BoilSteps != nil {
		boilSteps := make([]*beerproto.BoilStepType, 0)

		for _, step := range i.BoilSteps {
			boilSteps = append(boilSteps, ToProtoBoilStepType(&step))
		}
		boilProcedureType.BoilSteps = boilSteps
	}

	if i.PreBoilSize != nil {
		boilProcedureType.PreBoilSize = ToProtoVolumeType(i.PreBoilSize)
	}

	boilProcedureType.BoilTime = ToProtoTimeType(&i.BoilTime)

	if i.Name != nil {
		boilProcedureType.Name = UseString(i.Name)
	}

	if i.Description != nil {
		boilProcedureType.Description = UseString(i.Description)
	}

	if i.Notes != nil {
		boilProcedureType.Notes = UseString(i.Notes)
	}

	return boilProcedureType
}

func ToProtoBoilStepType(i *beerjson.BoilStepType) *beerproto.BoilStepType {
	if i == nil {
		return nil
	}

	return &beerproto.BoilStepType{
		EndGravity:       ToProtoGravityType(i.EndGravity),
		ChillingType:     ToProtoChillingType(i.BoilStepTypeChillingType),
		Description:      UseString(i.Description),
		EndTemperature:   ToProtoTemperatureType(i.EndTemperature),
		RampTime:         ToProtoTimeType(i.RampTime),
		StepTime:         ToProtoTimeType(i.StepTime),
		StartGravity:     ToProtoGravityType(i.StartGravity),
		StartPh:          ToProtoAcidityType(i.StartPh),
		EndPh:            ToProtoAcidityType(i.EndPh),
		Name:             i.Name,
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
	}
}

func ToProtoChillingType(i *beerjson.BoilStepTypeChillingType) beerproto.BoilStepTypeChillingType {
	if i == nil {
		return beerproto.BoilStepTypeChillingType_BOIL_STEP_TYPE_CHILLING_TYPE_UNSPECIFIED
	}
	unit := beerproto.BoilStepTypeChillingType_value[strings.ToUpper(string(*i))]
	return beerproto.BoilStepTypeChillingType(unit)
}

func ToProtoPackagingProcedureType(i *beerjson.PackagingProcedureType) *beerproto.PackagingProcedureType {
	if i == nil {
		return nil
	}

	packagingProcedureType := &beerproto.PackagingProcedureType{}

	if i.PackagingVessels != nil {
		packagingVessels := make([]*beerproto.PackagingVesselType, 0)

		for _, vessels := range i.PackagingVessels {
			packagingVessels = append(packagingVessels, ToProtoPackagingVesselType(&vessels))
		}
		packagingProcedureType.PackagingVessels = packagingVessels
	}

	if i.Name != "" {
		packagingProcedureType.Name = i.Name
	}

	if i.Description != nil {
		packagingProcedureType.Description = UseString(i.Description)
	}

	if i.Notes != nil {
		packagingProcedureType.Notes = UseString(i.Notes)
	}

	if i.PackagedVolume != nil {
		packagingProcedureType.PackagedVolume = ToProtoVolumeType(i.PackagedVolume)
	}

	return packagingProcedureType
}

func ToProtoPackagingVesselType(i *beerjson.PackagingVesselType) *beerproto.PackagingVesselType {
	if i == nil {
		return nil
	}

	packageDate := ""
	if i.PackageDate != nil {
		packageDate = string(*i.PackageDate)
	}
	return &beerproto.PackagingVesselType{
		Type:             ToProtoPackagingVesselTypeType(i.PackagingVesselTypeType),
		StartGravity:     ToProtoGravityType(i.StartGravity),
		Name:             i.Name,
		PackageDate:      packageDate,
		StepTime:         ToProtoTimeType(i.StepTime),
		EndGravity:       ToProtoGravityType(i.EndGravity),
		VesselVolume:     ToProtoVolumeType(i.VesselVolume),
		VesselQuantity:   UseFloat(i.VesselQuantity),
		Description:      UseString(i.Description),
		StartPh:          ToProtoAcidityType(i.StartPh),
		Carbonation:      UseFloat(i.Carbonation),
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
		EndPh:            ToProtoAcidityType(i.EndPh),
		EndTemperature:   ToProtoTemperatureType(i.EndTemperature),
	}
}

func ToProtoPackagingVesselTypeType(i *beerjson.PackagingVesselTypeType) beerproto.PackagingVesselUnit {
	if i == nil {
		return beerproto.PackagingVesselUnit_PACKAGING_VESSEL_UNIT_UNSPECIFIED
	}
	unit := beerproto.PackagingVesselUnit_value[strings.ToUpper(string(*i))]
	return beerproto.PackagingVesselUnit(unit)
}

func ToProtoEquipmentItemType(i *beerjson.EquipmentItemType) *beerproto.EquipmentItemType {
	if i == nil {
		return nil
	}

	return &beerproto.EquipmentItemType{
		BoilRatePerHour:     ToProtoVolumeType(i.BoilRatePerHour),
		Type:                UseString(i.KeyType),
		Form:                ToProtoEquipmentBaseForm(i.EquipmentBaseForm),
		DrainRatePerMinute:  ToProtoVolumeType(i.DrainRatePerMinute),
		SpecificHeat:        ToProtoSpecificHeatType(i.SpecificHeat),
		GrainAbsorptionRate: ToProtoSpecificVolumeType(i.GrainAbsorptionRate),
		Name:                UseString(i.Name),
		MaximumVolume:       ToProtoVolumeType(i.MaximumVolume),
		Weight:              ToProtoMassType(i.Weight),
		Loss:                ToProtoVolumeType(&i.Loss),
		Notes:               UseString(i.Notes),
	}
}

func ToProtoWaterBase(i *beerjson.WaterBase) *beerproto.WaterBase {
	if i == nil {
		return nil
	}

	return &beerproto.WaterBase{
		Calcium:     ToProtoConcentrationType(&i.Calcium),
		Nitrite:     ToProtoConcentrationType(i.Nitrite),
		Chloride:    ToProtoConcentrationType(&i.Chloride),
		Name:        i.Name,
		Potassium:   ToProtoConcentrationType(i.Potassium),
		Carbonate:   ToProtoConcentrationType(i.Carbonate),
		Iron:        ToProtoConcentrationType(i.Iron),
		Flouride:    ToProtoConcentrationType(i.Flouride),
		Sulfate:     ToProtoConcentrationType(&i.Sulfate),
		Magnesium:   ToProtoConcentrationType(&i.Magnesium),
		Producer:    UseString(i.Producer),
		Bicarbonate: ToProtoConcentrationType(&i.Bicarbonate),
		Nitrate:     ToProtoConcentrationType(i.Nitrate),
		Sodium:      ToProtoConcentrationType(&i.Sodium),
	}
}

func ToProtoVarietyInformation(i *beerjson.VarietyInformation) *beerproto.VarietyInformation {
	if i == nil {
		return nil
	}
	return &beerproto.VarietyInformation{
		Inventory:   ToProtoHopInventoryType(i.Inventory),
		Type:        ToProtoVarietyInformationType(i.VarietyInformationType),
		OilContent:  ToProtoOilContentType(i.OilContent),
		PercentLost: ToProtoPercentType(i.PercentLost),
		ProductId:   UseString(i.ProductId),
		AlphaAcid:   ToProtoPercentType(i.AlphaAcid),
		BetaAcid:    ToProtoPercentType(i.BetaAcid),
		Name:        UseString(i.Name),
		Origin:      UseString(i.Origin),
		Substitutes: UseString(i.Substitutes),
		Year:        UseString(i.Year),
		Form:        ToProtoHopVarietyBaseForm(i.HopVarietyBaseForm),
		Producer:    UseString(i.Producer),
		Notes:       UseString(i.Notes),
	}
}

func ToProtoOilContentType(i *beerjson.OilContentType) *beerproto.OilContentType {
	if i == nil {
		return nil
	}

	return &beerproto.OilContentType{
		Polyphenols:        ToProtoPercentType(i.Polyphenols),
		TotalOilMlPer_100G: UseFloat(i.TotalOilMlPer100g),
		Farnesene:          ToProtoPercentType(i.Farnesene),
		Limonene:           ToProtoPercentType(i.Limonene),
		Nerol:              ToProtoPercentType(i.Nerol),
		Geraniol:           ToProtoPercentType(i.Geraniol),
		BPinene:            ToProtoPercentType(i.BPinene),
		Linalool:           ToProtoPercentType(i.Linalool),
		Caryophyllene:      ToProtoPercentType(i.Caryophyllene),
		Cohumulone:         ToProtoPercentType(i.Cohumulone),
		Xanthohumol:        ToProtoPercentType(i.Xanthohumol),
		Humulene:           ToProtoPercentType(i.Humulene),
		Myrcene:            ToProtoPercentType(i.Myrcene),
		Pinene:             ToProtoPercentType(i.Pinene),
	}
}

func ToProtoVarietyInformationType(i *beerjson.VarietyInformationType) beerproto.VarietyInformationType {
	if i == nil {
		return beerproto.VarietyInformationType_VARIETY_INFORMATION_TYPE_UNSPECIFIED
	}

	unit := beerproto.VarietyInformationType_value[strings.ToUpper(string(*i))]
	return beerproto.VarietyInformationType(unit)
}

func ToProtoHopInventoryType(i *beerjson.HopInventoryType) *beerproto.HopInventoryType {
	if i == nil {
		return nil
	}

	hopInventoryType := &beerproto.HopInventoryType{}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		hopInventoryType.Amount = &beerproto.HopInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		hopInventoryType.Amount = &beerproto.HopInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return hopInventoryType
}

func ToProtoEquipmentType(i *beerjson.EquipmentType) *beerproto.EquipmentType {
	if i == nil {
		return nil
	}

	equipmentType := &beerproto.EquipmentType{
		Name: i.Name,
	}

	if i.EquipmentItems != nil {
		equipmentItemType := []*beerproto.EquipmentItemType{}
		for _, item := range i.EquipmentItems {
			equipmentItemType = append(equipmentItemType, ToProtoEquipmentItemType(&item))
		}
		equipmentType.EquipmentItems = equipmentItemType
	}

	return equipmentType
}

func ToProtoSpecificHeatType(i *beerjson.SpecificHeatType) *beerproto.SpecificHeatType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.SpecificHeatType{}
	}

	return &beerproto.SpecificHeatType{
		Value: i.Value,
		Unit:  ToProtoSpecificHeatUnitType(&i.Unit),
	}
}

func ToProtoSpecificHeatUnitType(i *beerjson.SpecificHeatUnitType) beerproto.SpecificHeatUnit {
	if i == nil {
		return beerproto.SpecificHeatUnit_SPECIFIC_HEAT_UNIT_UNSPECIFIED
	}

	unit := beerproto.SpecificHeatUnit_value[strings.ToUpper(string(*i))]
	return beerproto.SpecificHeatUnit(unit)
}

func ToProtoEquipmentBaseForm(i *beerjson.EquipmentBaseForm) beerproto.EquipmentItemType_EquipmentBaseForm {
	if i == nil {
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED
	}

	switch *i {
	case beerjson.EquipmentBaseForm_HLT:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_HLT
	case beerjson.EquipmentBaseForm_MashTun:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_MASH_TUN
	case beerjson.EquipmentBaseForm_LauterTun:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_LAUTER_TUN
	case beerjson.EquipmentBaseForm_BrewKettle:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_BREW_KETTLE
	case beerjson.EquipmentBaseForm_Fermenter:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_FERMENTER
	case beerjson.EquipmentBaseForm_AgingVessel:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_AGING_VESSEL
	case beerjson.EquipmentBaseForm_PackagingVessel:
		return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_PACKAGING_VESSEL
	}

	return beerproto.EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED
}

func ToProtoCultureInformation(i *beerjson.CultureInformation) *beerproto.CultureInformation {
	if i == nil {
		return nil
	}

	return &beerproto.CultureInformation{
		Form:             ToProtoCultureBaseForm(i.CultureBaseForm),
		Producer:         UseString(i.Producer),
		TemperatureRange: ToProtoTemperatureRangeType(i.TemperatureRange),
		Notes:            UseString(i.Notes),
		BestFor:          UseString(i.BestFor),
		Inventory:        ToProtoCultureInventoryType(i.Inventory),
		ProductId:        UseString(i.ProductId),
		Name:             UseString(i.Name),
		AlcoholTolerance: ToProtoPercentType(i.AlcoholTolerance),
		Glucoamylase:     UseBool(i.Glucoamylase),
		Type:             ToProtoCultureBaseType(i.CultureBaseType),
		Flocculation:     ToProtoQualitativeRangeType(i.Flocculation),
		AttenuationRange: ToProtoPercentRangeType(i.AttenuationRange),
		MaxReuse:         UseInt(i.MaxReuse),
		Pof:              UseBool(i.Pof),
		Zymocide:         ToProtoZymocide(i.Zymocide),
	}
}

func ToProtoZymocide(i *beerjson.Zymocide) *beerproto.Zymocide {
	if i == nil {
		return nil
	}
	return &beerproto.Zymocide{
		No1:     UseBool(i.No1),
		No2:     UseBool(i.No2),
		No28:    UseBool(i.No28),
		Klus:    UseBool(i.Klus),
		Neutral: UseBool(i.Neutral),
	}
}
func ToProtoQualitativeRangeType(i *beerjson.QualitativeRangeType) beerproto.QualitativeRangeUnit {
	if i == nil {
		return beerproto.QualitativeRangeUnit_QUALITATIVE_RANGE_UNIT_UNSPECIFIED
	}
	unit := beerproto.QualitativeRangeUnit_value[strings.ToUpper(string(*i))]
	return beerproto.QualitativeRangeUnit(unit)
}

func ToProtoCultureBaseType(i *beerjson.CultureBaseType) beerproto.CultureBaseType {
	if i == nil {
		return beerproto.CultureBaseType_CULTURE_BASE_TYPE_UNSPECIFIED
	}
	unit := beerproto.CultureBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.CultureBaseType(unit)
}

func ToProtoCultureInventoryType(i *beerjson.CultureInventoryType) *beerproto.CultureInventoryType {
	if i == nil {
		return nil
	}
	return &beerproto.CultureInventoryType{
		Liquid:  ToProtoVolumeType(i.Liquid),
		Dry:     ToProtoMassType(i.Dry),
		Slant:   ToProtoVolumeType(i.Slant),
		Culture: ToProtoVolumeType(i.Culture),
	}
}

func ToProtoTemperatureRangeType(i *beerjson.TemperatureRangeType) *beerproto.TemperatureRangeType {
	if i == nil {
		return nil
	}

	temperatureRangeType := &beerproto.TemperatureRangeType{
		Minimum: ToProtoTemperatureType(&i.Minimum),
		Maximum: ToProtoTemperatureType(&i.Maximum),
	}

	return temperatureRangeType
}

func ToProtoFermentableType(i *beerjson.FermentableType) *beerproto.FermentableType {
	if i == nil {
		return nil
	}
	return &beerproto.FermentableType{
		MaxInBatch:     ToProtoPercentType(i.MaxInBatch),
		RecommendMash:  UseBool(i.RecommendMash),
		Protein:        ToProtoPercentType(i.Protein),
		ProductId:      UseString(i.ProductId),
		GrainGroup:     ToProtoGrainGroup(i.FermentableBaseGrainGroup),
		Yield:          ToProtoYieldType(i.Yield),
		Type:           ToProtoFermentableBaseType(i.FermentableBaseType),
		Producer:       UseString(i.Producer),
		AlphaAmylase:   UseFloat(i.AlphaAmylase),
		Color:          ToProtoColorType(i.Color),
		Name:           UseString(i.Name),
		DiastaticPower: ToProtoDiastaticPowerType(i.DiastaticPower),
		Moisture:       ToProtoPercentType(i.Moisture),
		Origin:         UseString(i.Origin),
		Inventory:      ToProtoFermentableInventoryType(i.Inventory),
		KolbachIndex:   ToProtoPercentType(i.KolbachIndex),
		Notes:          UseString(i.Notes),
		BetaGlucan:     ToProtoConcentrationType(i.BetaGlucan),
		Glassy:         ToProtoPercentType(i.Glassy),
		Plump:          ToProtoPercentType(i.Plump),
		Half:           ToProtoPercentType(i.Half),
		Mealy:          ToProtoPercentType(i.Mealy),
		Thru:           ToProtoPercentType(i.Thru),
		Friability:     ToProtoPercentType(i.Friability),
		Fermentability: ToProtoPercentType(i.Fermentability),
		DiPh:           ToProtoAcidityType(i.DipH),
		Viscosity:      ToProtoViscosityType(i.Viscosity),
		DmsP:           ToProtoConcentrationType(i.DMSP),
		Fan:            ToProtoConcentrationType(i.FAN),
	}
}

func ToProtoFermentableInventoryType(i *beerjson.FermentableInventoryType) *beerproto.FermentableInventoryType {
	if i == nil {
		return nil
	}

	fermentableInventoryType := &beerproto.FermentableInventoryType{}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		fermentableInventoryType.Amount = &beerproto.FermentableInventoryType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		fermentableInventoryType.Amount = &beerproto.FermentableInventoryType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return fermentableInventoryType
}

func ToProtoViscosityType(i *beerjson.ViscosityType) *beerproto.ViscosityType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.ViscosityType{}
	}

	return &beerproto.ViscosityType{
		Value: i.Value,
		Unit:  ToProtoViscosityUnitType(&i.Unit),
	}
}

func ToProtoViscosityUnitType(i *beerjson.ViscosityUnitType) beerproto.ViscosityUnit {
	if i == nil {
		return beerproto.ViscosityUnit_VISCOSITY_UNIT_UNSPECIFIED
	}

	if *i == beerjson.ViscosityUnitType_MPAS {
		return beerproto.ViscosityUnit_VISCOSITY_UNIT_MPAS
	}

	if *i == beerjson.ViscosityUnitType_CP {
		return beerproto.ViscosityUnit_VISCOSITY_UNIT_CP
	}

	unit := beerproto.ViscosityUnit_value[strings.ToUpper(string(*i))]
	return beerproto.ViscosityUnit(unit)
}

func ToProtoDiastaticPowerType(i *beerjson.DiastaticPowerType) *beerproto.DiastaticPowerType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.DiastaticPowerType{}
	}

	return &beerproto.DiastaticPowerType{
		Value: i.Value,
		Unit:  ToProtoDiastaticPowerUnitType(&i.Unit),
	}
}

func ToProtoDiastaticPowerUnitType(i *beerjson.DiastaticPowerUnitType) beerproto.DiastaticPowerUnit {
	if i == nil {
		return beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_UNSPECIFIED
	}

	unit := beerproto.DiastaticPowerUnit_value[strings.ToUpper(string(*i))]
	return beerproto.DiastaticPowerUnit(unit)
}

func ToProtoStyleType(i *beerjson.StyleType) *beerproto.StyleType {
	if i == nil {
		return nil
	}

	return &beerproto.StyleType{
		Aroma:                        UseString(i.Aroma),
		Ingredients:                  UseString(i.Ingredients),
		CategoryNumber:               UseInt(i.CategoryNumber),
		Notes:                        UseString(i.Notes),
		Flavor:                       UseString(i.Flavor),
		Mouthfeel:                    UseString(i.Mouthfeel),
		FinalGravity:                 ToProtoGravityRangeType(i.FinalGravity),
		StyleGuide:                   UseString(i.StyleGuide),
		Color:                        ToProtoColorRangeType(i.Color),
		OriginalGravity:              ToProtoGravityRangeType(i.OriginalGravity),
		Examples:                     UseString(i.Examples),
		Name:                         UseString(i.Name),
		Carbonation:                  ToProtoCarbonationRangeType(i.Carbonation),
		AlcoholByVolume:              ToProtoPercentRangeType(i.AlcoholByVolume),
		InternationalBitternessUnits: ToProtoBitternessRangeType(i.InternationalBitternessUnits),
		Appearance:                   UseString(i.Appearance),
		Category:                     UseString(i.Category),
		StyleLetter:                  UseString(i.StyleLetter),
		Type:                         ToProtoStyleType_StyleCategories(i.KeyType),
		OverallImpression:            UseString(i.OverallImpression),
	}
}

func ToProtoStyleType_StyleCategories(i *beerjson.StyleCategories) beerproto.StyleType_StyleCategories {
	if i == nil {
		return beerproto.StyleType_STYLE_CATEGORIES_UNSPECIFIED
	}
	unit := beerproto.StyleType_StyleCategories_value[strings.ToUpper(string(*i))]
	return beerproto.StyleType_StyleCategories(unit)
}

func ToProtoBitternessRangeType(i *beerjson.BitternessRangeType) *beerproto.BitternessRangeType {
	if i == nil {
		return nil
	}
	return &beerproto.BitternessRangeType{
		Minimum: ToProtoBitternessType(&i.Minimum),
		Maximum: ToProtoBitternessType(&i.Maximum),
	}
}

func ToProtoBitternessType(i *beerjson.BitternessType) *beerproto.BitternessType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.BitternessType{}
	}

	return &beerproto.BitternessType{
		Value: i.Value,
		Unit:  ToProtoBitternessUnitType(&i.Unit),
	}
}

func ToProtoBitternessUnitType(i *beerjson.BitternessUnitType) beerproto.BitternessUnit {
	if i == nil {
		return beerproto.BitternessUnit_BITTERNESS_UNIT_UNSPECIFIED
	}

	if *i == beerjson.BitternessUnitType_IBUs {
		return beerproto.BitternessUnit_BITTERNESS_UNIT_IBUS
	}

	unit := beerproto.BitternessUnit_value[strings.ToUpper(string(*i))]
	return beerproto.BitternessUnit(unit)
}

func ToProtoPercentRangeType(i *beerjson.PercentRangeType) *beerproto.PercentRangeType {
	if i == nil {
		return nil
	}
	return &beerproto.PercentRangeType{
		Minimum: ToProtoPercentType(&i.Minimum),
		Maximum: ToProtoPercentType(&i.Maximum),
	}
}

func ToProtoCarbonationRangeType(i *beerjson.CarbonationRangeType) *beerproto.CarbonationRangeType {
	if i == nil {
		return nil
	}
	return &beerproto.CarbonationRangeType{
		Minimum: ToProtoCarbonationType(&i.Minimum),
		Maximum: ToProtoCarbonationType(&i.Maximum),
	}
}
func ToProtoCarbonationType(i *beerjson.CarbonationType) *beerproto.CarbonationType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.CarbonationType{}
	}

	return &beerproto.CarbonationType{
		Value: i.Value,
		Unit:  ToProtoCarbonationUnitType(&i.Unit),
	}
}

func ToProtoIngredientsType(i *beerjson.IngredientsType) *beerproto.IngredientsType {
	if i == nil {
		return nil
	}

	ingredientsType := &beerproto.IngredientsType{}

	if i.MiscellaneousAdditions != nil {
		miscellaneousAdditions := make([]*beerproto.MiscellaneousAdditionType, 0)
		for _, misc := range i.MiscellaneousAdditions {
			miscellaneousAdditions = append(miscellaneousAdditions, ToProtoMiscellaneousAdditionType(&misc))
		}
		ingredientsType.MiscellaneousAdditions = miscellaneousAdditions
	}

	if i.CultureAdditions != nil {
		cultureAdditions := make([]*beerproto.CultureAdditionType, 0)
		for _, culture := range i.CultureAdditions {
			cultureAdditions = append(cultureAdditions, ToProtoCultureAdditionType(&culture))
		}
		ingredientsType.CultureAdditions = cultureAdditions
	}

	if i.WaterAdditions != nil {
		waterAdditions := make([]*beerproto.WaterAdditionType, 0)
		for _, water := range i.WaterAdditions {
			waterAdditions = append(waterAdditions, ToProtoWaterAdditionType(&water))
		}
		ingredientsType.WaterAdditions = waterAdditions
	}

	if i.FermentableAdditions != nil {
		fermentableAdditions := make([]*beerproto.FermentableAdditionType, 0)
		for _, fermentable := range i.FermentableAdditions {
			fermentableAdditions = append(fermentableAdditions, ToProtoFermentableAdditionType(&fermentable))
		}
		ingredientsType.FermentableAdditions = fermentableAdditions
	}

	if i.HopAdditions != nil {
		hopAdditions := make([]*beerproto.HopAdditionType, 0)
		for _, hop := range i.HopAdditions {
			hopAdditions = append(hopAdditions, ToProtoHopAdditionType(&hop))
		}
		ingredientsType.HopAdditions = hopAdditions
	}

	return ingredientsType
}

func ToProtoHopAdditionType(i *beerjson.HopAdditionType) *beerproto.HopAdditionType {
	if i == nil {
		return nil
	}

	hopAdditionType := &beerproto.HopAdditionType{
		BetaAcid:  ToProtoPercentType(i.BetaAcid),
		Producer:  UseString(i.Producer),
		Origin:    UseString(i.Origin),
		Year:      UseString(i.Year),
		Form:      ToProtoHopVarietyBaseForm(i.HopVarietyBaseForm),
		Timing:    ToProtoTimingType(&i.Timing),
		Name:      UseString(i.Name),
		ProductId: UseString(i.ProductId),
		AlphaAcid: ToProtoPercentType(i.AlphaAcid),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		hopAdditionType.Amount = &beerproto.HopAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return hopAdditionType
}

func ToProtoHopVarietyBaseForm(i *beerjson.HopVarietyBaseForm) beerproto.HopVarietyBaseForm {
	if i == nil {
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_UNSPECIFIED
	}

	switch *i {
	case beerjson.HopVarietyBaseForm_Extract:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_EXTRACT
	case beerjson.HopVarietyBaseForm_Leaf:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_LEAF
	case beerjson.HopVarietyBaseForm_LeafWet:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_LEAFWET
	case beerjson.HopVarietyBaseForm_Pellet:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PELLET
	case beerjson.HopVarietyBaseForm_Powder:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_POWDER
	case beerjson.HopVarietyBaseForm_Plug:
		return beerproto.HopVarietyBaseForm_HOP_VARIETY_BASE_FORM_PLUG
	}

	unit := beerproto.HopVarietyBaseForm_value[strings.ToUpper(string(*i))]
	return beerproto.HopVarietyBaseForm(unit)
}

func ToProtoFermentableAdditionType(i *beerjson.FermentableAdditionType) *beerproto.FermentableAdditionType {
	if i == nil {
		return nil
	}

	fermentableAdditionType := &beerproto.FermentableAdditionType{
		Type:       ToProtoFermentableBaseType(i.FermentableBaseType),
		Origin:     UseString(i.Origin),
		GrainGroup: ToProtoGrainGroup(i.FermentableBaseGrainGroup),
		Yield:      ToProtoYieldType(i.Yield),
		Color:      ToProtoColorType(i.Color),
		Name:       UseString(i.Name),
		Producer:   UseString(i.Producer),
		ProductId:  UseString(i.ProductId),
		Timing:     ToProtoTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		fermentableAdditionType.Amount = &beerproto.FermentableAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return fermentableAdditionType
}

func ToProtoYieldType(i *beerjson.YieldType) *beerproto.YieldType {
	if i == nil {
		return nil
	}

	return &beerproto.YieldType{
		FineGrind:            ToProtoPercentType(i.FineGrind),
		CoarseGrind:          ToProtoPercentType(i.CoarseGrind),
		FineCoarseDifference: ToProtoPercentType(i.FineCoarseDifference),
		Potential:            ToProtoGravityType(i.Potential),
	}
}

func ToProtoGrainGroup(i *beerjson.FermentableBaseGrainGroup) beerproto.GrainGroup {
	if i == nil {
		return beerproto.GrainGroup_GRAIN_GROUP_UNSPECIFIED
	}

	unit := beerproto.GrainGroup_value[strings.ToUpper(string(*i))]
	return beerproto.GrainGroup(unit)
}

func ToProtoFermentableBaseType(i *beerjson.FermentableBaseType) beerproto.FermentableBaseType {
	if i == nil {
		return beerproto.FermentableBaseType_FERMENTABLE_BASE_TYPE_UNSPECIFIED
	}

	unit := beerproto.FermentableBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.FermentableBaseType(unit)
}

func ToProtoWaterAdditionType(i *beerjson.WaterAdditionType) *beerproto.WaterAdditionType {
	if i == nil {
		return nil
	}

	return &beerproto.WaterAdditionType{
		Flouride:    ToProtoConcentrationType(i.Flouride),
		Sulfate:     ToProtoConcentrationType(i.Sulfate),
		Producer:    UseString(i.Producer),
		Nitrate:     ToProtoConcentrationType(i.Nitrate),
		Nitrite:     ToProtoConcentrationType(i.Nitrite),
		Chloride:    ToProtoConcentrationType(i.Chloride),
		Amount:      ToProtoVolumeType(i.Amount),
		Name:        UseString(i.Name),
		Potassium:   ToProtoConcentrationType(i.Potassium),
		Magnesium:   ToProtoConcentrationType(i.Magnesium),
		Iron:        ToProtoConcentrationType(i.Iron),
		Bicarbonate: ToProtoConcentrationType(i.Bicarbonate),
		Calcium:     ToProtoConcentrationType(i.Calcium),
		Carbonate:   ToProtoConcentrationType(i.Carbonate),
		Sodium:      ToProtoConcentrationType(i.Sodium),
	}
}

func ToProtoConcentrationType(i *beerjson.ConcentrationType) *beerproto.ConcentrationType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.ConcentrationType{}
	}

	return &beerproto.ConcentrationType{
		Value: i.Value,
		Unit:  ToProtoConcentrationUnitType(i.Unit),
	}
}

func ToProtoConcentrationUnitType(i beerjson.ConcentrationUnitType) beerproto.ConcentrationUnit {
	if i == beerjson.ConcentrationUnitType_MgL {
		return beerproto.ConcentrationUnit_CONCENTRATION_UNIT_MGL
	}

	unit := beerproto.ConcentrationUnit_value[strings.ToUpper(string(i))]
	return beerproto.ConcentrationUnit(unit)
}

func ToProtoCultureAdditionType(i *beerjson.CultureAdditionType) *beerproto.CultureAdditionType {
	if i == nil {
		return nil
	}

	cultureAdditionType := &beerproto.CultureAdditionType{
		Form:              ToProtoCultureBaseForm(i.CultureBaseForm),
		ProductId:         UseString(i.ProductId),
		Name:              UseString(i.Name),
		CellCountBillions: UseInt(i.CellCountBillions),
		TimesCultured:     UseInt(i.TimesCultured),
		Producer:          UseString(i.Producer),
		Type:              ToProtoCultureBaseType(i.CultureBaseType),
		Attenuation:       ToProtoPercentType(i.Attenuation),
		Timing:            ToProtoTimingType(i.Timing),
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok := i.Amount.(*beerjson.UnitType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		cultureAdditionType.Amount = &beerproto.CultureAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return cultureAdditionType
}

func ToProtoCultureBaseForm(i *beerjson.CultureBaseForm) beerproto.CultureBaseForm {
	if i == nil {
		return beerproto.CultureBaseForm_CULTURE_BASE_FORM_UNSPECIFIED
	}
	unit := beerproto.CultureBaseForm_value[strings.ToUpper(string(*i))]
	return beerproto.CultureBaseForm(unit)
}

func ToProtoMiscellaneousAdditionType(i *beerjson.MiscellaneousAdditionType) *beerproto.MiscellaneousAdditionType {
	if i == nil {
		return nil
	}

	miscellaneousAdditionType := &beerproto.MiscellaneousAdditionType{
		Name:      UseString(i.Name),
		Producer:  UseString(i.Producer),
		Timing:    ToProtoTimingType(i.Timing),
		ProductId: UseString(i.ProductId),
	}

	if i.MiscellaneousBaseType != nil {
		miscellaneousAdditionType.Type = ToProtoMiscellaneousBaseType(i.MiscellaneousBaseType)
	}

	if mass, ok := i.Amount.(*beerjson.MassType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Mass{
			Mass: ToProtoMassType(mass),
		}
	}

	if unit, ok := i.Amount.(*beerjson.UnitType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Unit{
			Unit: ToProtoUnitType(unit),
		}
	}
	if volume, ok := i.Amount.(*beerjson.VolumeType); ok {
		miscellaneousAdditionType.Amount = &beerproto.MiscellaneousAdditionType_Volume{
			Volume: ToProtoVolumeType(volume),
		}
	}

	return miscellaneousAdditionType
}

func ToProtoUnitType(i *beerjson.UnitType) *beerproto.UnitType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.UnitType{}
	}

	return &beerproto.UnitType{
		Value: i.Value,
		Unit:  ToProtoUnitUnitType(i.Unit),
	}
}

func ToProtoUnitUnitType(i beerjson.UnitUnitType) beerproto.UnitUnit {
	unit := beerproto.UnitUnit_value[strings.ToUpper(string(i))]
	return beerproto.UnitUnit(unit)
}

func ToProtoMassType(i *beerjson.MassType) *beerproto.MassType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.MassType{}
	}

	return &beerproto.MassType{
		Value: i.Value,
		Unit:  ToProtoMassUnitType(i.Unit),
	}
}

func ToProtoMassUnitType(i beerjson.MassUnitType) beerproto.MassUnit {
	unit := beerproto.MassUnit_value[strings.ToUpper(string(i))]
	return beerproto.MassUnit(unit)
}

func ToProtoMiscellaneousBaseType(i *beerjson.MiscellaneousBaseType) beerproto.MiscellaneousBaseType {
	unit := beerproto.MiscellaneousBaseType_value[strings.ToUpper(string(*i))]
	return beerproto.MiscellaneousBaseType(unit)
}

func ToProtoTimingType(i *beerjson.TimingType) *beerproto.TimingType {
	if i == nil {
		return nil
	}
	return &beerproto.TimingType{
		Time:            ToProtoTimeType(i.Time),
		Duration:        ToProtoTimeType(i.Duration),
		Continuous:      UseBool(i.Continuous),
		SpecificGravity: ToProtoGravityType(i.SpecificGravity),
		Ph:              ToProtoAcidityType(i.PH),
		Step:            UseInt(i.Step),
		Use:             ToProtoUseType(i.Use),
	}
}

func ToProtoUseType(i *beerjson.UseType) beerproto.UseType {
	if i == nil {
		return beerproto.UseType_USE_TYPE_UNSPECIFIED
	}
	unit := beerproto.UseType_value[strings.ToUpper(string(*i))]
	return beerproto.UseType(unit)
}

func ToProtoFermentationProcedureType(i *beerjson.FermentationProcedureType) *beerproto.FermentationProcedureType {
	if i == nil {
		return nil
	}
	steps := make([]*beerproto.FermentationStepType, 0)
	for _, step := range i.FermentationSteps {
		steps = append(steps, ToProtoFermentationStepType(&step))
	}
	return &beerproto.FermentationProcedureType{
		Description:       UseString(i.Description),
		Notes:             UseString(i.Notes),
		Name:              i.Name,
		FermentationSteps: steps,
	}
}

func ToProtoFermentationStepType(i *beerjson.FermentationStepType) *beerproto.FermentationStepType {
	if i == nil {
		return nil
	}

	return &beerproto.FermentationStepType{
		Name:             i.Name,
		EndTemperature:   ToProtoTemperatureType(i.EndTemperature),
		StepTime:         ToProtoTimeType(i.StepTime),
		FreeRise:         UseBool(i.FreeRise),
		StartGravity:     ToProtoGravityType(i.StartGravity),
		EndGravity:       ToProtoGravityType(i.EndGravity),
		StartPh:          ToProtoAcidityType(i.StartPh),
		Description:      UseString(i.Description),
		StartTemperature: ToProtoTemperatureType(i.StartTemperature),
		EndPh:            ToProtoAcidityType(i.EndPh),
		Vessel:           UseString(i.Vessel),
	}
}

func ToProtoGravityType(i *beerjson.GravityType) *beerproto.GravityType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.GravityType{}
	}

	return &beerproto.GravityType{
		Value: i.Value,
		Unit:  ToProtoGravityUnitType(i.Unit),
	}
}

func ToProtoGravityUnitType(i beerjson.GravityUnitType) beerproto.GravityUnit {
	unit := beerproto.GravityUnit_value[strings.ToUpper(string(i))]
	return beerproto.GravityUnit(unit)
}

func ToProtoRecipeTypeType(i beerjson.RecipeTypeType) beerproto.RecipeUnit {
	unit := beerproto.RecipeUnit_value[strings.ToUpper(string(i))]
	return beerproto.RecipeUnit(unit)
}

func ToProtoColorType(i *beerjson.ColorType) *beerproto.ColorType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.ColorType{}
	}

	return &beerproto.ColorType{
		Value: i.Value,
		Unit:  ToProtoColorUnitType(i.Unit),
	}
}

func ToProtoColorUnitType(i beerjson.ColorUnitType) beerproto.ColorUnit {
	unit := beerproto.ColorUnit_value[strings.ToUpper(string(i))]

	return beerproto.ColorUnit(unit)
}

func ToProtoIBUEstimateType(i *beerjson.IBUEstimateType) *beerproto.IBUEstimateType {
	if i == nil {
		return nil
	}

	return &beerproto.IBUEstimateType{
		Method: ToProtoIBUMethodType(i.Method),
	}
}

func ToProtoIBUMethodType(i *beerjson.IBUMethodType) beerproto.IBUMethodUnit {
	if i == nil {
		return beerproto.IBUMethodUnit_IBU_METHOD_UNIT_UNSPECIFIED
	}
	unit := beerproto.IBUMethodUnit_value[strings.ToUpper(string(*i))]
	return beerproto.IBUMethodUnit(unit)
}

func ToProtoRecipeStyleType(i *beerjson.RecipeStyleType) *beerproto.RecipeStyleType {
	if i == nil {
		return nil
	}
	return &beerproto.RecipeStyleType{
		Type:           ToProtoRecipeStyleType_StyleCategories(i.KeyType),
		Name:           UseString(i.Name),
		Category:       UseString(i.Category),
		CategoryNumber: UseInt(i.CategoryNumber),
		StyleGuide:     UseString(i.StyleGuide),
		StyleLetter:    UseString(i.StyleLetter),
	}
}

func ToProtoRecipeStyleType_StyleCategories(i *beerjson.StyleCategories) beerproto.RecipeStyleType_StyleCategories {
	if i == nil {
		return beerproto.RecipeStyleType_STYLE_CATEGORIES_UNSPECIFIED
	}

	unit := beerproto.RecipeStyleType_StyleCategories_value[strings.ToUpper(string(*i))]
	return beerproto.RecipeStyleType_StyleCategories(unit)
}

func ToProtoEfficiencyType(i *beerjson.EfficiencyType) *beerproto.EfficiencyType {
	if i == nil {
		return nil
	}

	return &beerproto.EfficiencyType{
		Conversion: ToProtoPercentType(i.Conversion),
		Lauter:     ToProtoPercentType(i.Lauter),
		Mash:       ToProtoPercentType(i.Mash),
		Brewhouse:  ToProtoPercentType(&i.Brewhouse),
	}
}

func ToProtoPercentType(i *beerjson.PercentType) *beerproto.PercentType {
	if i == nil {
		return nil
	}
	return &beerproto.PercentType{
		Value: i.Value,
		Unit:  ToProtoPercentUnitType(i.Unit),
	}
}

func ToProtoPercentUnitType(i beerjson.PercentUnitType) beerproto.PercentUnit {
	unit := beerproto.PercentUnit_value[strings.ToUpper(string(i))]
	return beerproto.PercentUnit(unit)
}

func ToProtoMashProcedureType(i *beerjson.MashProcedureType) *beerproto.MashProcedureType {
	if i == nil {
		return nil
	}

	mashSteps := []*beerproto.MashStepType{}
	for _, step := range i.MashSteps {
		mashSteps = append(mashSteps, ToProtoMashStepType(&step))
	}

	return &beerproto.MashProcedureType{
		Name:             i.Name,
		Notes:            UseString(i.Notes),
		GrainTemperature: ToProtoTemperatureType(&i.GrainTemperature),
		MashSteps:        mashSteps,
	}
}

func ToProtoMashStepType(i *beerjson.MashStepType) *beerproto.MashStepType {
	if i == nil {
		return nil
	}

	return &beerproto.MashStepType{
		StepTime:          ToProtoTimeType(&i.StepTime),
		RampTime:          ToProtoTimeType(i.RampTime),
		EndTemperature:    ToProtoTemperatureType(i.EndTemperature),
		Description:       UseString(i.Description),
		InfuseTemperature: ToProtoTemperatureType(i.InfuseTemperature),
		StartPh:           ToProtoAcidityType(i.StartPH),
		EndPh:             ToProtoAcidityType(i.EndPH),
		Name:              i.Name,
		Type:              ToProtoMashStepTypeType(i.MashStepTypeType),
		Amount:            ToProtoVolumeType(i.Amount),
		StepTemperature:   ToProtoTemperatureType(&i.StepTemperature),
		WaterGrainRatio:   ToProtoSpecificVolumeType(i.WaterGrainRatio),
	}
}

func ToProtoVolumeType(i *beerjson.VolumeType) *beerproto.VolumeType {
	if i == nil {
		return nil
	}
	if i.Unit == "" {
		return &beerproto.VolumeType{}
	}

	return &beerproto.VolumeType{
		Value: i.Value,
		Unit:  ToProtoVolumeUnitType(i.Unit),
	}
}

func ToProtoVolumeUnitType(i beerjson.VolumeUnitType) beerproto.VolumeUnit {
	unit := beerproto.VolumeUnit_value[strings.ToUpper(string(i))]
	return beerproto.VolumeUnit(unit)
}

func ToProtoSpecificVolumeType(i *beerjson.SpecificVolumeType) *beerproto.SpecificVolumeType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.SpecificVolumeType{}
	}

	return &beerproto.SpecificVolumeType{
		Value: i.Value,
		Unit:  ToProtoSpecificVolumeUnitType(i.Unit),
	}
}

func ToProtoSpecificVolumeUnitType(i beerjson.SpecificVolumeUnitType) beerproto.SpecificVolumeUnit {
	unit := beerproto.SpecificVolumeUnit_value[strings.ToUpper(string(i))]

	switch i {
	case beerjson.SpecificVolumeUnitType_QtLb:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_QTLB
	case beerjson.SpecificVolumeUnitType_GalLb:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALLB
	case beerjson.SpecificVolumeUnitType_GalOz:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_GALOZ
	case beerjson.SpecificVolumeUnitType_LG:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LG
	case beerjson.SpecificVolumeUnitType_LKg:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_LKG
	case beerjson.SpecificVolumeUnitType_FlozOz:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FLOZOZ
	case beerjson.SpecificVolumeUnitType_M3Kg:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_M3KG
	case beerjson.SpecificVolumeUnitType_Ft3Lb:
		return beerproto.SpecificVolumeUnit_SPECIFIC_VOLUME_UNIT_FT3LB
	}

	return beerproto.SpecificVolumeUnit(unit)
}

func ToProtoMashStepTypeType(i beerjson.MashStepTypeType) beerproto.MashStepType_MashStepUnit {
	unit := beerproto.MashStepType_MashStepUnit_value[strings.ToUpper(string(i))]
	return beerproto.MashStepType_MashStepUnit(unit)
}

func ToProtoAcidityType(i *beerjson.AcidityType) *beerproto.AcidityType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.AcidityType{}
	}

	return &beerproto.AcidityType{
		Value: i.Value,
		Unit:  ToProtoAcidityUnitType(i.Unit),
	}
}

func ToProtoAcidityUnitType(i beerjson.AcidityUnitType) beerproto.AcidityUnit {
	unit := beerproto.AcidityUnit_value[strings.ToUpper(string(i))]
	return beerproto.AcidityUnit(unit)
}

func ToProtoTimeType(i *beerjson.TimeType) *beerproto.TimeType {
	if i == nil {
		return nil
	}

	if i.Unit == "" {
		return &beerproto.TimeType{}
	}

	return &beerproto.TimeType{
		Value: i.Value,
		Unit:  ToProtoTimeUnitType(i.Unit),
	}
}

func ToProtoTemperatureType(i *beerjson.TemperatureType) *beerproto.TemperatureType {
	if i == nil {
		return nil
	}
	if i.Unit == "" {
		return &beerproto.TemperatureType{}
	}
	return &beerproto.TemperatureType{
		Value: i.Value,
		Unit:  ToProtoTemperatureUnitType(i.Unit),
	}
}

func ToProtoTimeUnitType(i beerjson.TimeUnitType) beerproto.TimeUnit {
	unit := beerproto.TimeUnit_value[strings.ToUpper(string(i))]
	return beerproto.TimeUnit(unit)
}

func ToProtoTemperatureUnitType(i beerjson.TemperatureUnitType) beerproto.TemperatureUnit {
	unit := beerproto.TemperatureUnit_value[strings.ToUpper(string(i))]
	return beerproto.TemperatureUnit(unit)
}

func UseString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func UseFloat(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}

func UseBool(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

func UseInt(s *int32) int32 {
	if s == nil {
		return 0
	}
	return *s
}
