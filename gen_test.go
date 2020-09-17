package ctxdata_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/peterbourgon/ctxdata/v4"
)

func TestDataNoDefaults(t *testing.T) {
	t.Parallel()

	_, d := ctxdata.New(context.Background())

	{
		v := "bar"
		d.Set("foo", v)
		s := d.GetString("foo")

		if want, have := v, s; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}

		var zeroval string
		s = d.GetString("unknown")

		if want, have := zeroval, s; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
	{
		v := 1
		d.Set("foo", v)
		i := d.GetInt("foo")

		if want, have := v, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}

		var zeroval int
		i = d.GetInt("unknown")

		if want, have := zeroval, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := int64(1)
		d.Set("foo", v)
		i := d.GetInt64("foo")

		if want, have := v, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}

		var zeroval int64
		i = d.GetInt64("unknown")

		if want, have := zeroval, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := uint64(1)
		d.Set("foo", v)
		u := d.GetUint64("foo")

		if want, have := v, u; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}

		var zeroval uint64
		u = d.GetUint64("unknown")

		if want, have := zeroval, u; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := 1.0
		d.Set("foo", v)
		f := d.GetFloat64("foo")

		if want, have := v, f; want != have {
			t.Fatalf("len: want %f, have %f", want, have)
		}

		var zeroval float64
		f = d.GetFloat64("unknown")

		if want, have := zeroval, f; want != have {
			t.Fatalf("len: want %f, have %f", want, have)
		}
	}
	{
		v := true
		d.Set("foo", v)
		b := d.GetBool("foo")

		if want, have := v, b; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}

		var zeroval bool
		b = d.GetBool("unknown")

		if want, have := zeroval, b; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}
	}
	{
		tn := time.Now()
		v := time.Since(tn)
		d.Set("foo", v)
		td := d.GetDuration("foo")

		if want, have := v, td; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}

		var zeroval time.Duration
		td = d.GetDuration("unknown")

		if want, have := zeroval, td; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
	{
		v := time.Now()
		d.Set("foo", v)
		tn := d.GetTime("foo")

		if want, have := v, tn; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}

		var zeroval time.Time
		tn = d.GetTime("unknown")

		if want, have := zeroval, tn; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
	{
		v := "whoops"
		d.Set("foo", errors.New(v))
		e := d.GetError("foo")

		if want, have := v, e.Error(); want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}

		var zeroval error
		e = d.GetError("unknown")

		if want, have := zeroval, e; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
}

func TestDataDefaults(t *testing.T) {
	t.Parallel()

	_, d := ctxdata.New(context.Background())

	{
		v := "bar"
		s := d.GetStringDefault("foo", v)

		if want, have := v, s; want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
	{
		v := 1
		i := d.GetIntDefault("foo", v)

		if want, have := v, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := int64(1)
		i := d.GetInt64Default("foo", v)

		if want, have := v, i; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := uint64(1)
		u := d.GetUint64Default("foo", v)

		if want, have := v, u; want != have {
			t.Fatalf("len: want %d, have %d", want, have)
		}
	}
	{
		v := 1.0
		f := d.GetFloat64Default("foo", v)

		if want, have := v, f; want != have {
			t.Fatalf("len: want %f, have %f", want, have)
		}
	}
	{
		v := true
		b := d.GetBoolDefault("foo", v)

		if want, have := v, b; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}
	}
	{
		tn := time.Now()
		v := time.Since(tn)
		td := d.GetDurationDefault("foo", v)

		if want, have := v, td; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}
	}
	{
		v := time.Now()
		tn := d.GetTimeDefault("foo", v)

		if want, have := v, tn; want != have {
			t.Fatalf("len: want %v, have %v", want, have)
		}
	}
	{
		v := "whoops"
		e := d.GetErrorDefault("foo", errors.New(v))

		if want, have := v, e.Error(); want != have {
			t.Fatalf("len: want %s, have %s", want, have)
		}
	}
}
