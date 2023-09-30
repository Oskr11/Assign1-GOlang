package main

import (
	"fmt"
)

func main() {
	youtubeChannel := &Channel{}

	subscriber1 := NewSubscriber("Subscriber 1")
	subscriber2 := NewSubscriber("Subscriber 2")
	subscriber3 := NewSubscriber("Subscriber 3")

	youtubeChannel.Subscribe(subscriber1)
	youtubeChannel.Subscribe(subscriber2)
	youtubeChannel.Subscribe(subscriber3)

	youtubeChannel.UploadVideo("Коллеги вы меня разочаровываете")

	youtubeChannel.Unsubscribe(subscriber2)

	youtubeChannel.UploadVideo("Топ 5 способов как не разочаровать коллегу")
}

type Observer interface {
	Update(string)
}
type Channel struct {
	subscribers []Observer
	videoTitle  string
}
func (c *Channel) Subscribe(observer Observer) {
	c.subscribers = append(c.subscribers, observer)
}

func (c *Channel) Unsubscribe(observer Observer) {
	indexToRemove := -1
	for i, subscriber := range c.subscribers {
		if subscriber == observer {
			indexToRemove = i
			break
		}
	}

	if indexToRemove != -1 {
		c.subscribers = append(c.subscribers[:indexToRemove], c.subscribers[indexToRemove+1:]...)
	}
}

func (c *Channel) NotifySubscribers() {
	for _, subscriber := range c.subscribers {
		subscriber.Update(c.videoTitle)
	}
}
func (c *Channel) UploadVideo(videoTitle string) {
	c.videoTitle = videoTitle
	fmt.Printf("New video uploaded: %s\n", c.videoTitle)
	c.NotifySubscribers()
}

type Subscriber struct {
	name string
}
func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name}
}

func (s *Subscriber) Update(videoTitle string) {
	fmt.Printf("%s received a notification: New video is uploaded - %s\n", s.name, videoTitle)
}


