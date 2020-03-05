package common

import (
	"io"
	"log"

	"github.com/hako/branca"
	"github.com/vmihailenco/msgpack"
)

// Message struct is used to exchange the control messages between the peers
// in the control stream.
type Message struct {
	Command string `msgpack:"command"`
	Context string `msgpack:"context"`
}

// NewMessage is used to create a new message object
func NewMessage(command, context string) *Message {
	return &Message{
		Command: command,
		Context: context,
	}
}

// EncodeTo is used to write the msgpack notation of the message
// struct
func (m *Message) EncodeTo(w io.Writer) error {
	return msgpack.NewEncoder(w).Encode(m)
}

// EncodeTo is used to write the msgpack notation of the message
// struct
func (m *Message) EnryptTo(w io.Writer) error {
	b := branca.NewBranca(Secretkey) // This key must be exactly 32 bytes long.
	// Encode String to Branca Token.
	token, err := b.EncodeToString(m.Context)
	log.Println("Encryptinh")
	if err != nil {
		log.Println("Encryptinh Failure", err)
		return err
	}

	log.Println("Encryptinh Success", token)

	return NewMessage(m.Command, token).EncodeTo(w)
}

func (m *Message) Authenticate() error {
	b := branca.NewBranca(Secretkey) // This key must be exactly 32 bytes long.
	// Encode String to Branca Token.
	message, err := b.DecodeToString(m.Context)
	if err != nil {
		return err
	}

	m.Context = message
	return nil
}

// DecodeFrom is used to decode the message into the message struct
func (m *Message) DecodeFrom(r io.Reader) (*Message, error) {
	return m, msgpack.NewDecoder(r).Decode(m)
}
