package raftmdb

import (
	"github.com/hashicorp/raft/bench"
	"os"
	"testing"
)

func BenchmarkMDBStore_FirstIndex(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.FirstIndex(b, store)
}

func BenchmarkMDBStore_LastIndex(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.LastIndex(b, store)
}

func BenchmarkMDBStore_GetLog(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.GetLog(b, store)
}

func BenchmarkMDBStore_StoreLog(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.StoreLog(b, store)
}

func BenchmarkMDBStore_StoreLogs(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.StoreLogs(b, store)
}

func BenchmarkMDBStore_DeleteRange(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.DeleteRange(b, store)
}

func BenchmarkMDBStore_Set(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.Set(b, store)
}

func BenchmarkMDBStore_Get(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.Get(b, store)
}

func BenchmarkMDBStore_SetUint64(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.SetUint64(b, store)
}

func BenchmarkMDBStore_GetUint64(b *testing.B) {
	dir, store := MDBTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.GetUint64(b, store)
}
