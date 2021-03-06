package network

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	opv1alpha1 "github.com/kubevirt/cluster-network-addons-operator/pkg/apis/networkaddonsoperator/v1alpha1"
)

var _ = Describe("Testing nmstate", func() {
	Describe("changeSafeNMState", func() {
		Context("when it is kept disabled", func() {
			prev := &opv1alpha1.NetworkAddonsConfigSpec{}
			new := &opv1alpha1.NetworkAddonsConfigSpec{}
			It("should pass", func() {
				errorList := changeSafeNMState(prev, new)
				Expect(errorList).To(BeEmpty())
			})
		})

		Context("when there is no previous value", func() {
			prev := &opv1alpha1.NetworkAddonsConfigSpec{}
			new := &opv1alpha1.NetworkAddonsConfigSpec{NMState: &opv1alpha1.NMState{}}
			It("should accept any configuration", func() {
				errorList := changeSafeNMState(prev, new)
				Expect(errorList).To(BeEmpty())
			})
		})

		Context("when the previous and new configuration match", func() {
			prev := &opv1alpha1.NetworkAddonsConfigSpec{NMState: &opv1alpha1.NMState{}}
			new := &opv1alpha1.NetworkAddonsConfigSpec{NMState: &opv1alpha1.NMState{}}
			It("should accept the configuration", func() {
				errorList := changeSafeNMState(prev, new)
				Expect(errorList).To(BeEmpty())
			})
		})

		Context("when there is previous value, but the new one is empty (removing component)", func() {
			prev := &opv1alpha1.NetworkAddonsConfigSpec{NMState: &opv1alpha1.NMState{}}
			new := &opv1alpha1.NetworkAddonsConfigSpec{}
			It("should fail", func() {
				errorList := changeSafeNMState(prev, new)
				Expect(len(errorList)).To(Equal(1), "validation of safe change failed due to an unexpected error: %v", errorList)
				Expect(errorList[0].Error()).To(Equal("cannot modify NMState state handler configuration once it is deployed"))
			})
		})
	})
})
