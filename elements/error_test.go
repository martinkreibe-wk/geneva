package elements_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/martinkreibe-wk/geneva/elements"
)

var _ = Describe("The Error constructs", func() {

	Context("creating simple errors", func() {
		It("should return the error as expected", func() {
			myMessage := "My special message"
			err := NewError(myMessage)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(Error("")))
			Expect(err.Error()).To(BeEquivalentTo(myMessage))
		})

		It("should ignore nil errors on append", func() {
			err := AppendError()
			Expect(err).To(BeNil())
			err = AppendError(nil)
			Expect(err).To(BeNil())
			err = AppendError(nil, nil, nil)
			Expect(err).To(BeNil())
		})

		It("should append errors then they are valid", func() {
			myMessage := "My special message"
			err := NewError(myMessage)

			err = AppendError(err)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(Error("")))
			Expect(err.Error()).To(BeEquivalentTo(myMessage))

			err = AppendError(nil, err) // should ignore the nil
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(Error("")))
			Expect(err.Error()).To(BeEquivalentTo(myMessage))
		})
	})

	Context("creating formatted errors", func() {
		It("should create the error with the format substitution", func() {
			myMessage := "My special message: %s"
			val := "foo"
			formatted := "My special message: foo"
			err := NewError(myMessage, val)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&FormatError{}))
			Expect(err.Error()).To(BeEquivalentTo(formatted))
		})

		It("should append like other errors", func() {
			myMessage := "My special message: %s"
			val := "foo"
			formatted := "My special message: foo"
			err := NewError(myMessage, val)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&FormatError{}))
			Expect(err.Error()).To(BeEquivalentTo(formatted))

			err = AppendError(err)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&FormatError{}))
			Expect(err.Error()).To(BeEquivalentTo(formatted))
		})
	})

	Context("creating cumulative errors", func() {
		It("should be created on append of multiple errors", func() {
			mySimpleMessage := "My special message"
			err1 := NewError(mySimpleMessage)
			Expect(err1).ToNot(BeNil())
			Expect(err1).To(BeAssignableToTypeOf(Error("")))
			Expect(err1.Error()).To(BeEquivalentTo(mySimpleMessage))

			myFormatMessage := "My special message: %s"
			val := "foo"
			formatted := "My special message: foo"
			err2 := NewError(myFormatMessage, val)
			Expect(err2).ToNot(BeNil())
			Expect(err2).To(BeAssignableToTypeOf(&FormatError{}))
			Expect(err2.Error()).To(BeEquivalentTo(formatted))

			v := &CumulativeError{}
			v.Append(err1, err2)

			err := AppendError(err1, err2)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&CumulativeError{}))
			Expect(err.Error()).To(BeEquivalentTo(v.Error()))

			err3 := Error("Another error")

			err = AppendError(err, err3)
			Expect(err).ToNot(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&CumulativeError{}))
			Expect(err.Error()).To(BeEquivalentTo("0: My special message\n1: My special message: foo\n2: Another error\n"))
		})

		It("should be created on append the cumulative items from another error onto the first", func() {
			mySimpleMessage := "My special message"
			err1 := NewError(mySimpleMessage)
			myFormatMessage := "My special message: %s"
			val := "foo"
			err2 := NewError(myFormatMessage, val)

			v1 := &CumulativeError{}
			v1.Append(err1)

			v2 := &CumulativeError{}
			v2.Append(err2)

			v1.Append(v2)
			Expect(v1.ErrorList()).To(HaveLen(2))
			Expect(v1.ErrorList()[0]).To(BeIdenticalTo(err1))
			Expect(v1.ErrorList()[1]).To(BeIdenticalTo(err2))
		})
	})
})
