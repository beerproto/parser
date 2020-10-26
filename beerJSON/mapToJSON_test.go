package beerJSON

import (
	"github.com/beerproto/beerjson.go"
	beerproto "github.com/beerproto/beerproto_go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Water", func() {
	var (
		water  beerproto.WaterBase
	)

	BeforeSuite(func() {
		water = beerproto.WaterBase{
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
	})

	Describe("Map water Proto to JSON", func() {
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

