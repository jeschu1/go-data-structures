package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLinkedList(t *testing.T) {
	list := &LinkedList{}

	assert.NotNil(t, list)
}

func Test_InsertCreatesNewHead(t *testing.T) {
	list := &LinkedList{}

	assert.Nil(t, list.head)
	list.Insert("testing")

	assert.NotNil(t, list.head)
	assert.Equal(t, "testing", list.head.data)
}

func Test_InsertMovesPreviousHeadToNext(t *testing.T) {
	list := &LinkedList{}

	list.Insert("first insert")
	list.Insert("second insert")

	assert.Equal(t, "second insert", list.head.data)
	assert.Equal(t, "first insert", list.head.next.data)
}