package wasi

import (
	"encoding/binary"
	"io"
	"unsafe"

	"github.com/stealthrocket/wazergo/types"
	"github.com/stealthrocket/wazergo/wasm"
	"github.com/tetratelabs/wazero/api"
)

func (f FDStat) ObjectSize() int                                  { return int(unsafe.Sizeof(FDStat{})) }
func (f FDStat) LoadObject(_ api.Memory, b []byte) FDStat         { return unsafeLoad[FDStat](b) }
func (f FDStat) StoreObject(_ api.Memory, b []byte)               { unsafeStore(b, f) }
func (f FDStat) FormatObject(w io.Writer, _ api.Memory, b []byte) { formatObject(w, b, f) }

func (f FileStat) ObjectSize() int                                  { return int(unsafe.Sizeof(FileStat{})) }
func (f FileStat) LoadObject(_ api.Memory, b []byte) FileStat       { return unsafeLoad[FileStat](b) }
func (f FileStat) StoreObject(_ api.Memory, b []byte)               { unsafeStore(b, f) }
func (f FileStat) FormatObject(w io.Writer, _ api.Memory, b []byte) { formatObject(w, b, f) }

func (p PreStat) ObjectSize() int                                  { return int(unsafe.Sizeof(PreStat{})) }
func (p PreStat) LoadObject(_ api.Memory, b []byte) PreStat        { return unsafeLoad[PreStat](b) }
func (p PreStat) StoreObject(_ api.Memory, b []byte)               { unsafeStore(b, p) }
func (p PreStat) FormatObject(w io.Writer, _ api.Memory, b []byte) { formatObject(w, b, p) }

func (s Subscription) ObjectSize() int { return int(unsafe.Sizeof(Subscription{})) }
func (s Subscription) LoadObject(_ api.Memory, b []byte) Subscription {
	return unsafeLoad[Subscription](b)
}
func (s Subscription) StoreObject(_ api.Memory, b []byte)               { unsafeStore(b, s) }
func (s Subscription) FormatObject(w io.Writer, _ api.Memory, b []byte) { formatObject(w, b, s) }

func (e Event) ObjectSize() int                                  { return int(unsafe.Sizeof(Event{})) }
func (e Event) LoadObject(_ api.Memory, b []byte) Event          { return unsafeLoad[Event](b) }
func (e Event) StoreObject(_ api.Memory, b []byte)               { unsafeStore(b, e) }
func (e Event) FormatObject(w io.Writer, _ api.Memory, b []byte) { formatObject(w, b, e) }

func (arg IOVec) ObjectSize() int {
	return 8
}

func (arg IOVec) LoadObject(memory api.Memory, object []byte) IOVec {
	offset := binary.LittleEndian.Uint32(object[:4])
	length := binary.LittleEndian.Uint32(object[4:])
	return wasm.Read(memory, offset, length)
}

func (arg IOVec) StoreObject(memory api.Memory, object []byte) {
	panic("BUG: i/o vectors cannot be stored back to wasm memory")
}

func (arg IOVec) FormatObject(w io.Writer, memory api.Memory, object []byte) {
	types.Bytes(arg.LoadObject(memory, object)).Format(w)
}

func formatObject[T types.Object[T]](w io.Writer, object []byte, typ T) {
	types.Format(w, typ.LoadObject(nil, object))
}

func unsafeStore[T types.Object[T]](b []byte, t T) {
	types.UnsafeStoreObject(b, t)
}

func unsafeLoad[T types.Object[T]](b []byte) T {
	return types.UnsafeLoadObject[T](b)
}
