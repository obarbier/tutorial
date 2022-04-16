package event

import (
	"bytes"
	"encoding/gob"
	"github.com/obarbier/cqrs-go/schema"

	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	meowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
}

var _ EventStore = &NatsEventStore{}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (mq *NatsEventStore) Close() {
	if mq.nc != nil {
		mq.nc.Close()
	}
	if mq.meowCreatedSubscription != nil {
		mq.meowCreatedSubscription.Unsubscribe()
	}
	close(mq.meowCreatedChan)
}

func (mq *NatsEventStore) PublishMeowCreated(meow schema.Meow) error {
	m := MeowCreatedMessage{meow.ID, meow.Body, meow.CreatedAt}
	data, err := mq.writeMessage(&m)
	if err != nil {
		return err
	}
	return mq.nc.Publish(m.Key(), data)
}

func (mq *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (mq *NatsEventStore) SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	m := MeowCreatedMessage{}
	mq.meowCreatedChan = make(chan MeowCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	mq.meowCreatedSubscription, err = mq.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				mq.readMessage(msg.Data, &m)
				mq.meowCreatedChan <- m
			}
		}
	}()
	return (<-chan MeowCreatedMessage)(mq.meowCreatedChan), nil
}
func (mq *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
func (mq *NatsEventStore) OnMeowCreated(f func(MeowCreatedMessage)) (err error) {
	m := MeowCreatedMessage{}
	mq.meowCreatedSubscription, err = mq.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		mq.readMessage(msg.Data, &m)
		f(m)
	})
	return
}
