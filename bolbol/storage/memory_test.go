package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"push_notification/entity"
	"testing"
)

func testNewTestMemory(m Storage, t *testing.T) {
	ctx := context.Background()

	m.Push(ctx, 10, entity.UnreadMessageNotification{Count: 1})
	m.Push(ctx, 10, entity.UnreadMessageNotification{Count: 2})
	m.Push(ctx, 10, entity.UnreadMessageNotification{Count: 3})
	c, _ := m.Count(ctx, 10)
	assert.Equal(t, 3, c)

	p, err := m.Pop(ctx, 10)
	assert.NoError(t, err)
	assert.Equal(t, 1, p.(entity.UnreadMessageNotification).Count)

	all, _ := m.PopAll(ctx, 10)
	assert.Equal(t, 2, len(all))

	for i := 0; i < 15; i++ {
		m.Push(ctx, 10, entity.UnreadMessageNotification{Count: i})
	}

	f, err := m.Pop(ctx, 10)
	assert.NoError(t, err)
	assert.Equal(t, 5, f.(entity.UnreadMessageNotification).Count)
}
