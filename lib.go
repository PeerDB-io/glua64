package glua64

import (
	"strconv"

	"github.com/yuin/gopher-lua"
)

var (
	I64 = UserDataType[int64]{Name: "flatbuffers_i64"}
	U64 = UserDataType[uint64]{Name: "flatbuffers_u64"}
)

func Loader(ls *lua.LState) int {
	eq64 := ls.NewFunction(Lua64Eq)
	le64 := ls.NewFunction(Lua64Le)
	lt64 := ls.NewFunction(Lua64Lt)
	mt := I64.NewMetatable(ls)
	mt.RawSetString("__index", ls.NewFunction(LuaI64Index))
	mt.RawSetString("__tostring", ls.NewFunction(LuaI64String))
	mt.RawSetString("__eq", eq64)
	mt.RawSetString("__le", le64)
	mt.RawSetString("__lt", lt64)

	mt = U64.NewMetatable(ls)
	mt.RawSetString("__index", ls.NewFunction(LuaU64Index))
	mt.RawSetString("__tostring", ls.NewFunction(LuaU64String))
	mt.RawSetString("__eq", eq64)
	mt.RawSetString("__le", le64)
	mt.RawSetString("__lt", lt64)

	return 0
}

func Lua64Eq(ls *lua.LState) int {
	aud := ls.CheckUserData(1)
	bud := ls.CheckUserData(2)
	switch a := aud.Value.(type) {
	case int64:
		switch b := bud.Value.(type) {
		case int64:
			ls.Push(lua.LBool(a == b))
		case uint64:
			if a < 0 {
				ls.Push(lua.LFalse)
			} else {
				ls.Push(lua.LBool(uint64(a) == b))
			}
		default:
			return 0
		}
	case uint64:
		switch b := bud.Value.(type) {
		case int64:
			if b < 0 {
				ls.Push(lua.LFalse)
			} else {
				ls.Push(lua.LBool(a == uint64(b)))
			}
		case uint64:
			ls.Push(lua.LBool(a == b))
		default:
			return 0
		}
	default:
		return 0
	}
	return 1
}

func Lua64Le(ls *lua.LState) int {
	aud := ls.CheckUserData(1)
	bud := ls.CheckUserData(2)
	switch a := aud.Value.(type) {
	case int64:
		switch b := bud.Value.(type) {
		case int64:
			ls.Push(lua.LBool(a <= b))
		case uint64:
			if a < 0 {
				ls.Push(lua.LTrue)
			} else {
				ls.Push(lua.LBool(uint64(a) <= b))
			}
		default:
			return 0
		}
	case uint64:
		switch b := bud.Value.(type) {
		case int64:
			if b < 0 {
				ls.Push(lua.LFalse)
			} else {
				ls.Push(lua.LBool(a <= uint64(b)))
			}
		case uint64:
			ls.Push(lua.LBool(a <= b))
		default:
			return 0
		}
	default:
		return 0
	}
	return 1
}

func Lua64Lt(ls *lua.LState) int {
	aud := ls.CheckUserData(1)
	bud := ls.CheckUserData(2)
	switch a := aud.Value.(type) {
	case int64:
		switch b := bud.Value.(type) {
		case int64:
			ls.Push(lua.LBool(a < b))
		case uint64:
			if a < 0 {
				ls.Push(lua.LTrue)
			} else {
				ls.Push(lua.LBool(uint64(a) < b))
			}
		default:
			return 0
		}
	case uint64:
		switch b := bud.Value.(type) {
		case int64:
			if b < 0 {
				ls.Push(lua.LTrue)
			} else {
				ls.Push(lua.LBool(a < uint64(b)))
			}
		case uint64:
			ls.Push(lua.LBool(a < b))
		default:
			return 0
		}
	default:
		return 0
	}
	return 1
}

func LuaI64Index(ls *lua.LState) int {
	i64ud, i64 := I64.Check(ls, 1)
	key := ls.CheckString(2)
	switch key {
	case "i64":
		ls.Push(i64ud)
	case "u64":
		ls.Push(U64.New(ls, uint64(i64)))
	case "float64":
		ls.Push(lua.LNumber(i64))
	case "hi":
		ls.Push(lua.LNumber(uint32(i64 >> 32)))
	case "lo":
		ls.Push(lua.LNumber(uint32(i64)))
	default:
		return 0
	}
	return 1
}

func LuaU64Index(ls *lua.LState) int {
	u64ud, u64 := U64.Check(ls, 1)
	key := ls.CheckString(2)
	switch key {
	case "i64":
		ls.Push(I64.New(ls, int64(u64)))
	case "u64":
		ls.Push(u64ud)
	case "float64":
		ls.Push(lua.LNumber(u64))
	case "hi":
		ls.Push(lua.LNumber(uint32(u64 >> 32)))
	case "lo":
		ls.Push(lua.LNumber(uint32(u64)))
	default:
		return 0
	}
	return 1
}

func LuaI64String(ls *lua.LState) int {
	i64 := I64.StartMethod(ls)
	ls.Push(lua.LString(strconv.FormatInt(i64, 10)))
	return 1
}

func LuaU64String(ls *lua.LState) int {
	u64 := U64.StartMethod(ls)
	ls.Push(lua.LString(strconv.FormatUint(u64, 10)))
	return 1
}