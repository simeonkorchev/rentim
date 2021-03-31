package rent_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/simeonkorchev/rentim/pkg/fakes"
	"github.com/simeonkorchev/rentim/pkg/money"
	"github.com/simeonkorchev/rentim/pkg/rent"
)

var _ = Describe("Service", func() {
	var (
		service  rent.PropertyService
		repo     rent.Repository
		property rent.Property
		properties []rent.Property
		err      error
		repoErr  = errors.New("repo error")
	)

	Describe("Saving", func() {
		JustBeforeEach(func() {
			err = service.SaveProperty(property)
		})

		Context("property with no location", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{}
				service = rent.NewPropertyService(repo)
				m, _ := money.Make(10, money.BGN)
				property = rent.NewProperty("id", rent.OneBedroom, rent.NewLocation("", ""), 10, m)
			})

			It("returns error", func() {
				Expect(err).To(MatchError("invalid property location, city and street can not be empty"))
			})
		})

		Context("property with invalid price", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{}
				service = rent.NewPropertyService(repo)
				m, _ := money.Make(0, money.BGN)
				property = rent.NewProperty("id", rent.OneBedroom, rent.NewLocation("a", "b"), 10, m)
			})

			It("returns error", func() {
				Expect(err).To(MatchError("invalid property price, can not be less than or equal  0"))
			})
		})

		Context("property with error from repository", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					AddPropertyFunc: func(property rent.Property) error {
						return repoErr
					},
				}
				service = rent.NewPropertyService(repo)
				m, _ := money.Make(10, money.BGN)
				property = rent.NewProperty("id", rent.OneBedroom, rent.NewLocation("a", "b"), 10, m)
			})

			It("returns error", func() {
				Expect(err).To(MatchError(repoErr))
			})
		})

		Context("valid property", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					AddPropertyFunc: func(property rent.Property) error {
						return nil
					},
				}
				service = rent.NewPropertyService(repo)
				m, _ := money.Make(10, money.BGN)
				property = rent.NewProperty("id", rent.OneBedroom, rent.NewLocation("a", "b"), 10, m)
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FindPropertyByID", func() {
		JustBeforeEach(func() {
			property, err = service.FindPropertyByID("any")
		})

		Context("when repo returns error", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					GetPropertyFunc: func(id string) (rent.Property, error) {
						return rent.Property{}, repoErr
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return error", func() {
				Expect(err).To(MatchError(repoErr))
			})

			It("should return zero property", func() {
				Expect(property).To(BeZero())
			})
		})

		Context("when property is found", func() {
			BeforeEach(func() {
				m, _ := money.Make(100, money.EUR)
				repo = fakes.RepositoryFake{
					GetPropertyFunc: func(id string) (rent.Property, error) {
						return rent.NewProperty("1", rent.OneBedroom, rent.NewLocation("a","b"), 10, m), nil
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return the property", func() {
				Expect(property.Id()).To(Equal("1"))
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("FindAllProperties", func() {
		JustBeforeEach(func() {
			properties, err = service.FindAllProperties()
		})

		Context("when repo returns error", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					GetAllPropertiesFunc: func() (properties []rent.Property, err error) {
						return nil, repoErr
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return error", func() {
				Expect(err).To(MatchError(repoErr))
			})
		})

		Context("when properties are found", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					GetAllPropertiesFunc: func() (properties []rent.Property, err error) {
						m, _ := money.Make(10, money.EUR)
						return []rent.Property{rent.NewProperty("1", rent.TwoBedroom, rent.NewLocation("a","b"), 10, m)}, nil
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return the properties", func() {
				m, _ := money.Make(10, money.EUR)
				Expect(properties).To(Equal([]rent.Property{rent.NewProperty("1", rent.TwoBedroom, rent.NewLocation("a","b"), 10, m)}))
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("DeletePropertyByID", func() {
		JustBeforeEach(func() {
			property, err = service.DeletePropertyByID("id")
		})

		Context("when repo returns error", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					DeletePropertyFunc: func(id string) (rent.Property, error) {
						return rent.Property{}, repoErr
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return error", func() {
				Expect(err).To(MatchError(repoErr))
			})
		})

		Context("when property is deleted", func() {
			BeforeEach(func() {
				repo = fakes.RepositoryFake{
					DeletePropertyFunc: func(id string) (rent.Property, error) {
						m, _ := money.Make(10, money.BGN)
						return rent.NewProperty("1", rent.TwoBedroom, rent.NewLocation("a","b"), 199, m), nil
					},
				}
				service = rent.NewPropertyService(repo)
			})

			It("should return the property", func() {
				m, _ := money.Make(10, money.BGN)
				Expect(property).To(Equal(rent.NewProperty("1", rent.TwoBedroom, rent.NewLocation("a","b"), 199, m)))
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
