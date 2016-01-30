package telegraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Something struct {
	Score int
	Info  string
}

var _ = Describe("Broadcaster", func() {
	Describe("LinkedBroadcaster", func() {
		It("Accept subscriptions", func() {
			broadcaster := NewLinkedBroadcaster()
			subscription := broadcaster.Listen()
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscription.Channel
		})

		It("Remove subscriptions", func() {
			broadcaster := NewLinkedBroadcaster()
			subscriptionOne := broadcaster.Listen()
			subscriptionTwo := broadcaster.Listen()
			subscriptionOne.RemoveFrom(broadcaster)
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscriptionTwo.Channel
		})

		It("Non blocking publications", func() {
			broadcaster := NewLinkedBroadcaster()
			subscriptionOne := broadcaster.Listen()
			if subscriptionOne.Channel != nil {
				//Do somthing
			}
			subscriptionTwo := broadcaster.Listen()
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscriptionTwo.Channel
		})

		It("Publish a struct", func() {
			const expectedInfo = "BOOM"
			const expectedScore = 5

			broadcaster := NewLinkedBroadcaster()
			subscriptionOne := broadcaster.Listen()
			go func() {
				broadcaster.Notify(Something{
					Info:  expectedInfo,
					Score: expectedScore,
				})
			}()
			item := <-subscriptionOne.Channel
			something := item.(Something)
			Expect(something.Info).To(Equal(expectedInfo))
			Expect(something.Score).To(Equal(expectedScore))
		})
	})
})