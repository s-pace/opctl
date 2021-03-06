package parallelloop

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	loopableFakes "github.com/opctl/opctl/sdks/go/opspec/interpreter/loopable/fakes"
)

var _ = Context("Interpreter", func() {
	Context("NewInterpreter", func() {
		It("shouldn't return nil", func() {
			/* arrange/act/assert */
			Expect(NewInterpreter()).To(Not(BeNil()))
		})
	})
	Context("Interpret", func() {
		It("should call loopableInterpreter.Interpret w/ expected args", func() {
			/* arrange */
			providedSCGParallelLoopCall := model.SCGParallelLoopCall{
				Range: "range",
			}

			providedScope := map[string]*model.Value{}

			fakeLoopableInterpreter := new(loopableFakes.FakeInterpreter)

			objectUnderTest := _interpreter{
				loopableInterpreter: fakeLoopableInterpreter,
			}

			/* act */
			objectUnderTest.Interpret(
				providedSCGParallelLoopCall,
				providedScope,
			)

			/* assert */
			actualSCGParallelLoopCallOn,
				actualScope := fakeLoopableInterpreter.InterpretArgsForCall(0)

			Expect(actualSCGParallelLoopCallOn).To(Equal(providedSCGParallelLoopCall.Range))
			Expect(actualScope).To(Equal(providedScope))
		})
		Context("loopableInterpreter.Interpret errs", func() {
			It("should return expected result", func() {
				/* arrange */
				fakeLoopableInterpreter := new(loopableFakes.FakeInterpreter)

				expectedError := errors.New("expectedError")
				fakeLoopableInterpreter.InterpretReturns(
					nil,
					expectedError,
				)

				objectUnderTest := _interpreter{
					loopableInterpreter: fakeLoopableInterpreter,
				}

				/* act */
				_, actualError := objectUnderTest.Interpret(
					model.SCGParallelLoopCall{
						Range: "range",
					},
					map[string]*model.Value{},
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
		It("should return expected result", func() {
			/* arrange */
			providedScgLoop := model.SCGParallelLoopCall{
				Range: "range",
			}

			fakeLoopableInterpreter := new(loopableFakes.FakeInterpreter)

			expectedDCGParallelLoopCallRange := &model.Value{Array: new([]interface{})}

			expectedResult := &model.DCGParallelLoopCall{
				Range: expectedDCGParallelLoopCallRange,
			}
			fakeLoopableInterpreter.InterpretReturns(expectedDCGParallelLoopCallRange, nil)

			objectUnderTest := _interpreter{
				loopableInterpreter: fakeLoopableInterpreter,
			}

			/* act */
			actualResult, _ := objectUnderTest.Interpret(
				providedScgLoop,
				map[string]*model.Value{},
			)

			/* assert */
			Expect(actualResult).To(Equal(expectedResult))
		})
	})
})
