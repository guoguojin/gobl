package cache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"gitlab.com/gobl/gobl/pkg/cache"
)

func TestCache_Any(t *testing.T) {
	var c *cache.Cache[any]
	require.Nil(t, c)

	t.Run("New should create a new cache", func(t *testing.T) {
		c = cache.New[any]()
		require.NotNil(t, c)
	})

	t.Run("Set should set a value in the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, "bar", v)
	})

	t.Run("SetWithExpiry should set a value in the cache with an expiry", func(t *testing.T) {
		c.SetWithExpiry("foo", "bar", time.Second)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, "bar", v)
		time.Sleep(time.Second + time.Millisecond)
		v, ok = c.Get("foo")
		assert.False(t, ok)
		assert.Nil(t, v)
	})

	t.Run("Flush should remove all expired items from the cache", func(t *testing.T) {
		c.SetWithExpiry("foo", "bar", time.Second)
		c.SetWithExpiry("baz", "qux", time.Second)

		assert.Equal(t, 2, c.Len())
		time.Sleep(time.Second)
		c.Flush()
		assert.Equal(t, 0, c.Len())
	})

	t.Run("Delete should remove a value from the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		c.Delete("foo")
		v, ok := c.Get("foo")
		assert.False(t, ok)
		assert.Nil(t, v)
	})

	t.Run("Clear should remove all items from the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		c.Set("baz", "qux")
		assert.Equal(t, 2, c.Len())
		c.Clear()
		assert.Equal(t, 0, c.Len())
	})
}

func TestCache_String(t *testing.T) {
	var c *cache.Cache[string]
	require.Nil(t, c)

	t.Run("New should create a new cache", func(t *testing.T) {
		c = cache.New[string]()
		require.NotNil(t, c)
	})

	t.Run("Set should set a value in the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, "bar", v)
	})

	t.Run("SetWithExpiry should set a value in the cache with an expiry", func(t *testing.T) {
		c.SetWithExpiry("foo", "bar", time.Second)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, "bar", v)
		time.Sleep(time.Second + time.Millisecond)
		v, ok = c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, "", v)
	})

	t.Run("Flush should remove all expired items from the cache", func(t *testing.T) {
		c.SetWithExpiry("foo", "bar", time.Second)
		c.SetWithExpiry("baz", "qux", time.Second)

		assert.Equal(t, 2, c.Len())
		time.Sleep(time.Second)
		c.Flush()
		assert.Equal(t, 0, c.Len())
	})

	t.Run("Delete should remove a value from the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		c.Delete("foo")
		v, ok := c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, "", v)
	})

	t.Run("Clear should remove all items from the cache", func(t *testing.T) {
		c.Set("foo", "bar")
		c.Set("baz", "qux")
		assert.Equal(t, 2, c.Len())
		c.Clear()
		assert.Equal(t, 0, c.Len())
	})
}

func TestCache_Int(t *testing.T) {
	var c *cache.Cache[int]
	require.Nil(t, c)
	const (
		bar int = 100
		qux int = 200
		def int = 0
	)

	t.Run("New should create a new cache", func(t *testing.T) {
		c = cache.New[int]()
		require.NotNil(t, c)
	})

	t.Run("Set should set a value in the cache", func(t *testing.T) {
		c.Set("foo", bar)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, bar, v)
	})

	t.Run("SetWithExpiry should set a value in the cache with an expiry", func(t *testing.T) {
		c.SetWithExpiry("foo", bar, time.Second)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, bar, v)
		time.Sleep(time.Second + time.Millisecond)
		v, ok = c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, def, v)
	})

	t.Run("Flush should remove all expired items from the cache", func(t *testing.T) {
		c.SetWithExpiry("foo", bar, time.Second)
		c.SetWithExpiry("baz", qux, time.Second)

		assert.Equal(t, 2, c.Len())
		time.Sleep(time.Second)
		c.Flush()
		assert.Equal(t, 0, c.Len())
	})

	t.Run("Delete should remove a value from the cache", func(t *testing.T) {
		c.Set("foo", bar)
		c.Delete("foo")
		v, ok := c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, def, v)
	})

	t.Run("Clear should remove all items from the cache", func(t *testing.T) {
		c.Set("foo", bar)
		c.Set("baz", qux)
		assert.Equal(t, 2, c.Len())
		c.Clear()
		assert.Equal(t, 0, c.Len())
	})
}

func TestCache_Struct(t *testing.T) {
	type myStruct struct {
		Foo string
		Bar int
	}

	var c *cache.Cache[myStruct]
	require.Nil(t, c)

	var (
		bar = myStruct{Foo: "bar", Bar: 100}
		qux = myStruct{Foo: "qux", Bar: 200}
		def = myStruct{}
	)

	t.Run("New should create a new cache", func(t *testing.T) {
		c = cache.New[myStruct]()
		require.NotNil(t, c)
	})

	t.Run("Set should set a value in the cache", func(t *testing.T) {
		c.Set("foo", bar)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, bar, v)
	})

	t.Run("SetWithExpiry should set a value in the cache with an expiry", func(t *testing.T) {
		c.SetWithExpiry("foo", bar, time.Second)
		v, ok := c.Get("foo")
		assert.True(t, ok)
		assert.Equal(t, bar, v)
		time.Sleep(time.Second + time.Millisecond)
		v, ok = c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, def, v)
	})

	t.Run("Flush should remove all expired items from the cache", func(t *testing.T) {
		c.SetWithExpiry("foo", bar, time.Second)
		c.SetWithExpiry("baz", qux, time.Second)

		assert.Equal(t, 2, c.Len())
		time.Sleep(time.Second)
		c.Flush()
		assert.Equal(t, 0, c.Len())
	})

	t.Run("Delete should remove a value from the cache", func(t *testing.T) {
		c.Set("foo", bar)
		c.Delete("foo")
		v, ok := c.Get("foo")
		assert.False(t, ok)
		assert.Equal(t, def, v)
	})

	t.Run("Clear should remove all items from the cache", func(t *testing.T) {
		c.Set("foo", bar)
		c.Set("baz", qux)
		assert.Equal(t, 2, c.Len())
		c.Clear()
		assert.Equal(t, 0, c.Len())
	})
}
