package beerJSON

import (
	"github.com/beerproto/beerjson.go"
	beerproto "github.com/beerproto/beerproto_go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Recipe", func() {

	recipe := beerproto.Recipe{
		Mashes: []*beerproto.MashProcedureType{},
		Recipes: []*beerproto.RecipeType{},
		MiscellaneousIngredients: []*beerproto.MiscellaneousType{},
		Styles: []*beerproto.StyleType{},
		Fermentations: []*beerproto.FermentationProcedureType{},
		Boil: []*beerproto.BoilProcedureType{},
		Fermentables: []*beerproto.FermentableType{},
		TimingObject: &beerproto.TimingType{},
		Cultures: []*beerproto.CultureInformation{},
		Equipments: []*beerproto.EquipmentType{},
		Packaging: []*beerproto.PackagingProcedureType{},
		HopVarieties: []*beerproto.VarietyInformation{},
		Profiles: []*beerproto.WaterBase{},
	}

	Describe("Map Proto to JSON", func() {
		Context("Basic mapping", func() {

			It("should be json", func() {
				json, err := MapToJSON(&recipe)
				Expect(err).To(BeNil())
				Expect(json, Not(BeNil()))
				Expect(json.Mashes).To(BeEmpty())
				Expect(json.Recipes).To(BeEmpty())
				Expect(json.MiscellaneousIngredients).To(BeEmpty())
				Expect(json.Styles).To(BeEmpty())
				Expect(json.Fermentations).To(BeEmpty())
				Expect(json.Boil).To(BeEmpty())
				Expect(json.Fermentables).To(BeEmpty())
				Expect(json.Cultures).To(BeEmpty())
				Expect(json.Equipments).To(BeEmpty())
				Expect(json.Packaging).To(BeEmpty())
				Expect(json.HopVarieties).To(BeEmpty())
				Expect(json.Profiles).To(BeEmpty())
			})
		})
	})
})

var _ = Describe("RecipeType", func() {

	recipe := beerproto.RecipeType{
		IbuEstimate: &beerproto.IBUEstimateType{
			Method: beerproto.IBUEstimateType_TINSETH,
		},
		Efficiency: &beerproto.EfficiencyType{
			Brewhouse: &beerproto.PercentType{
				Value: 80,
				Unit: beerproto.PercentType_PERCENT_SIGN,
			},
		},
		Ingredients: &beerproto.IngredientsType{

		},
		BatchSize: &beerproto.VolumeType{
			Unit: beerproto.VolumeType_L,
			Value: float64(30),
		},
	}

	Describe("Map Proto to JSON", func() {
		var json *beerjson.RecipeType
		var err error
		Context("Basic mapping", func() {
			It("should be json", func() {
				json, err = ToJSONRecipeType(&recipe)
				Expect(err).To(BeNil())
				Expect(json, Not(BeNil()))

				By("IbuEstimate should be")
				Expect(json.IbuEstimate).To(Not(BeNil()))
				Expect(json.IbuEstimate.Method).To(Not(BeNil()))
				Expect(*json.IbuEstimate.Method).To(Equal(beerjson.IBUMethodType_Tinseth))
			})

			Context("Basic mapping", func() {
				It("should be proto", func() {
					result := ToProtoRecipeType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(&recipe))
				})
			})
		})
	})
})

var _ = Describe("IBUEstimateType", func() {

	Describe("Rager", func() {
		var json *beerjson.IBUEstimateType

		IBUEstimate := &beerproto.IBUEstimateType{
			Method: beerproto.IBUEstimateType_RAGER,
		}

		Context("Mapping", func() {
			It("should be json", func() {
				json  = ToJSONIBUEstimateType(IBUEstimate)
				Expect(json, Not(BeNil()))

				By("IbuEstimate should be")
				Expect(json).To(Not(BeNil()))
				Expect(json.Method).To(Not(BeNil()))
				Expect(*json.Method).To(Equal(beerjson.IBUMethodType_Rager))
			})

			Context("revert mapping", func() {
				It("should be proto", func() {
					result := ToProtoIBUEstimateType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(IBUEstimate))
				})
			})
		})
	})

	Describe("Garetz", func() {
		var json *beerjson.IBUEstimateType

		IBUEstimate := &beerproto.IBUEstimateType{
			Method: beerproto.IBUEstimateType_GARETZ,
		}

		Context("Mapping", func() {
			It("should be json", func() {
				json  = ToJSONIBUEstimateType(IBUEstimate)
				Expect(json, Not(BeNil()))

				By("IbuEstimate should be")
				Expect(json).To(Not(BeNil()))
				Expect(json.Method).To(Not(BeNil()))
				Expect(*json.Method).To(Equal(beerjson.IBUMethodType_Garetz))
			})

			Context("revert mapping", func() {
				It("should be proto", func() {
					result := ToProtoIBUEstimateType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(IBUEstimate))
				})
			})
		})
	})

	Describe("Tinseth", func() {
		var json *beerjson.IBUEstimateType

		IBUEstimate := &beerproto.IBUEstimateType{
			Method: beerproto.IBUEstimateType_TINSETH,
		}

		Context("Mapping", func() {
			It("should be json", func() {
				json  = ToJSONIBUEstimateType(IBUEstimate)
				Expect(json, Not(BeNil()))

				By("IbuEstimate should be")
				Expect(json).To(Not(BeNil()))
				Expect(json.Method).To(Not(BeNil()))
				Expect(*json.Method).To(Equal(beerjson.IBUMethodType_Tinseth))
			})

			Context("revert mapping", func() {
				It("should be proto", func() {
					result := ToProtoIBUEstimateType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(IBUEstimate))
				})
			})
		})
	})

	Describe("Other", func() {
		var json *beerjson.IBUEstimateType

		IBUEstimate := &beerproto.IBUEstimateType{
			Method: beerproto.IBUEstimateType_OTHER,
		}

		Context("Mapping", func() {
			It("should be json", func() {
				json  = ToJSONIBUEstimateType(IBUEstimate)
				Expect(json, Not(BeNil()))

				By("IbuEstimate should be")
				Expect(json).To(Not(BeNil()))
				Expect(json.Method).To(Not(BeNil()))
				Expect(*json.Method).To(Equal(beerjson.IBUMethodType_Other))
			})

			Context("revert mapping", func() {
				It("should be proto", func() {
					result := ToProtoIBUEstimateType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(IBUEstimate))
				})
			})
		})
	})
})

var _ = FDescribe("DiastaticPowerType", func() {

	Describe("Rager", func() {
		var json *beerjson.DiastaticPowerType

		diastaticPower := &beerproto.DiastaticPowerType{
			Value: float64(10),
			Unit: beerproto.DiastaticPowerUnitType_LINTNER,
		}

		Context("Mapping", func() {
			It("should be json", func() {
				json  = ToJSONDiastaticPowerType(diastaticPower)
				Expect(json, Not(BeNil()))


				Expect(json.Value).To(Equal(float64(10)))
				Expect(json.Unit).To(Equal(beerjson.DiastaticPowerUnitType_Lintner))
			})

			Context("revert mapping", func() {
				It("should be proto", func() {
					result := ToProtoDiastaticPowerType(json)
					Expect(result, Not(BeNil()))
					Expect(result).To(Equal(diastaticPower))
				})
			})
		})
	})
})

var _ = Describe("Water", func() {
	water := beerproto.WaterBase{
		Id:  "c1063779-1393-4047-9337-c3f99c76f6fb",
		Name: "Balanced Profile",
		Calcium: &beerproto.ConcentrationType{
			Value: 80,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
		Magnesium: &beerproto.ConcentrationType{
			Value: 5,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
		Sulfate: &beerproto.ConcentrationType{
			Value: 80,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
		Sodium: &beerproto.ConcentrationType{
			Value: 25,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
		Chloride: &beerproto.ConcentrationType{
			Value: 75,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
		Bicarbonate: &beerproto.ConcentrationType{
			Value: 100,
			Unit: beerproto.ConcentrationUnitType_MGL,
		},
	}

	Describe("Map Proto to JSON", func() {
		Context("Basic mapping", func() {

			var json *beerjson.WaterBase

			It("should be json", func() {
				json = ToJSONWaterBase(&water)
				Expect(json, Not(BeNil()))
			})

			It("Name should be", func() {
				Expect(json.Name).To(Equal("Balanced Profile"))
			})

			It("Calcium should be", func() {
				Expect(json.Calcium.Value).To(Equal(float64(80)))
				Expect(json.Calcium.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})

			It("Magnesium should be", func() {
				Expect(json.Magnesium.Value).To(Equal(float64(5)))
				Expect(json.Magnesium.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})

			It("Sulfate should be", func() {
				Expect(json.Sulfate.Value).To(Equal(float64(80)))
				Expect(json.Sulfate.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})

			It("Sodium should be", func() {
				Expect(json.Sodium.Value).To(Equal(float64(25)))
				Expect(json.Sodium.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})

			It("Chloride should be", func() {
				Expect(json.Chloride.Value).To(Equal(float64(75)))
				Expect(json.Chloride.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})

			It("Bicarbonate should be", func() {
				Expect(json.Bicarbonate.Value).To(Equal(float64(100)))
				Expect(json.Bicarbonate.Unit).To(Equal(beerjson.ConcentrationUnitType_MgL))
			})
		})
	})
})

