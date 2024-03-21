package glua64

import (
	"testing"

	"github.com/yuin/gopher-lua"
)

func assert(t *testing.T, ls *lua.LState, source string) {
	t.Helper()
	err := ls.DoString(source)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func Test_64(t *testing.T) {
	ls := lua.NewState(lua.Options{})
	Loader(ls)

	n5 := int64(-5)
	ls.Env.RawSetString("i64p5", I64.New(ls, 5))
	ls.Env.RawSetString("i64p5_2", I64.New(ls, 5))
	ls.Env.RawSetString("u64p5", U64.New(ls, 5))
	ls.Env.RawSetString("u64p5_2", U64.New(ls, 5))
	ls.Env.RawSetString("i64n5", I64.New(ls, n5))
	ls.Env.RawSetString("u64n5", U64.New(ls, uint64(n5)))

	assert(t, ls, `
assert(i64p5 == u64p5)
assert(i64p5 ~= i64n5)
assert(i64n5 ~= u64n5)
assert(i64p5 == i64p5_2)
assert(u64p5 == u64p5_2)
assert(u64n5 > i64p5)
assert(i64p5 > i64n5)
`)
}