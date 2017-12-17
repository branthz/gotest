package main

import (
	"fmt"
	"unsafe"
)

const (
	PtrSize = 4 << (^uintptr(0) >> 63) //==8
)

const (
	//bucket中存储的键值对数目
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	//bucket触发扩容的负载因子
	loadFactor = 6.5

	// Maximum key or value size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	// Fast versions cannot handle big values - the cutoff size for
	// fast versions in ../../cmd/internal/gc/walk.go must be at most this value.
	//最大存放的keysize
	maxKeySize   = 128
	maxValueSize = 128

	// bmap结构体内存偏移量
	dataOffset = unsafe.Offsetof(struct {
		b bmap
		v int64
	}{}.v)

	// Possible tophash values.  We reserve a few possibilities for special marks.
	// Each bucket (including its overflow buckets, if any) will have either all or none of its
	// entries in the evacuated* states (except during the evacuate() method, which only happens
	// during map writes and thus no one else can observe the map during that time).
	empty          = 0 // cell is empty
	evacuatedEmpty = 1 // cell is empty, bucket is evacuated.
	evacuatedX     = 2 // key/value is valid.  Entry has been evacuated to first half of larger table.
	evacuatedY     = 3 // same as above, but evacuated to second half of larger table.
	minTopHash     = 4 // minimum tophash for a normal filled cell.

	// flags
	iterator    = 1 // there may be an iterator using buckets
	oldIterator = 2 // there may be an iterator using oldbuckets
	hashWriting = 4 // a goroutine is writing to the map

	// 哨兵 bucket id（迭代检查）
	noCheck = 1<<(8*PtrSize) - 1 //uint64最大值
)

type bmap struct {
	tophash [bucketCnt]uint8
	//????
}

type hmap struct {
	count int // lenof(map)
	flags uint8
	B     uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	hash0 uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // 程序计数器for 清除(buckets 低于这个值 have been evacuated)

	//？？？？？？？
	//只有当key和value不为指针时才会用到overflow
	//overflow[0] 存放溢出buckerts
	//overflow[1] 存放溢出oldbuckets
	overflow *[2]*[]*bmap
}

//迭代器
type hiter struct {
	key   unsafe.Pointer // Must be in first position.  Write nil to indicate iteration end (see cmd/internal/gc/range.go).
	value unsafe.Pointer // Must be in second position (see cmd/internal/gc/range.go).
	//t           *maptype
	h           *hmap
	buckets     unsafe.Pointer // bucket ptr at hash_iter initialization time
	bptr        *bmap          // current bucket
	overflow    [2]*[]*bmap    // keeps overflow buckets alive
	startBucket uintptr        // bucket iteration started at
	offset      uint8          // intra-bucket offset to start from during iteration (should be big enough to hold bucketCnt-1)
	wrapped     bool           // already wrapped around from end of bucket array to beginning
	B           uint8
	i           uint8
	bucket      uintptr
	checkBucket uintptr
}

func main() {
	fmt.Printf("11111111\n")
	m := make(map[int]int, 1<<30)
	fmt.Println(len(m))
}
