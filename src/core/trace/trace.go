// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package trace

import (
	"github.com/primecitizens/pcz/std/core/arch"
)

const (
	// Timestamps in trace are cputicks/traceTickDiv.
	// This makes absolute values of timestamp diffs smaller,
	// and so they are encoded in less number of bytes.
	// 64 on x86 is somewhat arbitrary (one tick is ~20ns on a 3GHz machine).
	// The suggested increment frequency for PowerPC's time base register is
	// 512 MHz according to Power ISA v2.07 section 6.2, so we use 16 on ppc64
	// and ppc64le.
	TimeDiv = 16 + 48*(arch.Is386|arch.IsAmd64)
	// Maximum number of PCs in a single stack trace.
	// Since events contain only stack id rather than whole stack trace,
	// we can allow quite large values here.
	StackSize = 128
	// Identifier of a fake P that is used when we trace without a real P.
	GlobProc = -1
	// Maximum number of bytes to encode uint64 in base-128.
	BytesPerNumber = 10
	// Shift of the number of arguments in the first event byte.
	ArgCountShift = 6
)

// Event types in the trace, args are given in square brackets.
const (
	EvNone              = 0  // unused
	EvBatch             = 1  // start of per-P batch of events [pid, timestamp]
	EvFrequency         = 2  // contains tracer timer frequency [frequency (ticks per second)]
	EvStack             = 3  // stack [stack id, number of PCs, array of {PC, func string ID, file string ID, line}]
	EvGomaxprocs        = 4  // current value of GOMAXPROCS [timestamp, GOMAXPROCS, stack id]
	EvProcStart         = 5  // start of P [timestamp, thread id]
	EvProcStop          = 6  // stop of P [timestamp]
	EvGCStart           = 7  // GC start [timestamp, seq, stack id]
	EvGCDone            = 8  // GC done [timestamp]
	EvSTWStart          = 9  // STW start [timestamp, kind]
	EvSTWDone           = 10 // STW done [timestamp]
	EvGCSweepStart      = 11 // GC sweep start [timestamp, stack id]
	EvGCSweepDone       = 12 // GC sweep done [timestamp, swept, reclaimed]
	EvGoCreate          = 13 // goroutine creation [timestamp, new goroutine id, new stack id, stack id]
	EvGoStart           = 14 // goroutine starts running [timestamp, goroutine id, seq]
	EvGoEnd             = 15 // goroutine ends [timestamp]
	EvGoStop            = 16 // goroutine stops (like in select{}) [timestamp, stack]
	EvGoSched           = 17 // goroutine calls Gosched [timestamp, stack]
	EvGoPreempt         = 18 // goroutine is preempted [timestamp, stack]
	EvGoSleep           = 19 // goroutine calls Sleep [timestamp, stack]
	EvGoBlock           = 20 // goroutine blocks [timestamp, stack]
	EvGoUnblock         = 21 // goroutine is unblocked [timestamp, goroutine id, seq, stack]
	EvGoBlockSend       = 22 // goroutine blocks on chan send [timestamp, stack]
	EvGoBlockRecv       = 23 // goroutine blocks on chan recv [timestamp, stack]
	EvGoBlockSelect     = 24 // goroutine blocks on select [timestamp, stack]
	EvGoBlockSync       = 25 // goroutine blocks on Mutex/RWMutex [timestamp, stack]
	EvGoBlockCond       = 26 // goroutine blocks on Cond [timestamp, stack]
	EvGoBlockNet        = 27 // goroutine blocks on network [timestamp, stack]
	EvGoSysCall         = 28 // syscall enter [timestamp, stack]
	EvGoSysExit         = 29 // syscall exit [timestamp, goroutine id, seq, real timestamp]
	EvGoSysBlock        = 30 // syscall blocks [timestamp]
	EvGoWaiting         = 31 // denotes that goroutine is blocked when tracing starts [timestamp, goroutine id]
	EvGoInSyscall       = 32 // denotes that goroutine is in syscall when tracing starts [timestamp, goroutine id]
	EvHeapAlloc         = 33 // gcController.heapLive change [timestamp, heap_alloc]
	EvHeapGoal          = 34 // gcController.heapGoal() (formerly next_gc) change [timestamp, heap goal in bytes]
	EvTimerGoroutine    = 35 // not currently used; previously denoted timer goroutine [timer goroutine id]
	EvFutileWakeup      = 36 // not currently used; denotes that the previous wakeup of this goroutine was futile [timestamp]
	EvString            = 37 // string dictionary entry [ID, length, string]
	EvGoStartLocal      = 38 // goroutine starts running on the same P as the last event [timestamp, goroutine id]
	EvGoUnblockLocal    = 39 // goroutine is unblocked on the same P as the last event [timestamp, goroutine id, stack]
	EvGoSysExitLocal    = 40 // syscall exit on the same P as the last event [timestamp, goroutine id, real timestamp]
	EvGoStartLabel      = 41 // goroutine starts running with label [timestamp, goroutine id, seq, label string id]
	EvGoBlockGC         = 42 // goroutine blocks on GC assist [timestamp, stack]
	EvGCMarkAssistStart = 43 // GC mark assist start [timestamp, stack]
	EvGCMarkAssistDone  = 44 // GC mark assist done [timestamp]
	EvUserTaskCreate    = 45 // trace.NewTask [timestamp, internal task id, internal parent task id, name string, stack]
	EvUserTaskEnd       = 46 // end of a task [timestamp, internal task id, stack]
	EvUserRegion        = 47 // trace.WithRegion [timestamp, internal task id, mode(0:start, 1:end), name string, stack]
	EvUserLog           = 48 // trace.Log [timestamp, internal task id, key string id, stack, value string]
	EvCPUSample         = 49 // CPU profiling sample [timestamp, real timestamp, real P id (-1 when absent), goroutine id, stack]
	EvCount             = 50
	// Byte is used but only 6 bits are available for event type.
	// The remaining 2 bits are used to specify the number of arguments.
	// That means, the max event type value is 63.
)

// BlockReason is an enumeration of reasons a goroutine might block.
// This is the interface the rest of the runtime uses to tell the
// tracer why a goroutine blocked. The tracer then propagates this information
// into the trace however it sees fit.
//
// Note that BlockReasons should not be compared, since reasons that are
// distinct by name may *not* be distinct by value.
type BlockReason uint8

// For maximal efficiency, just map the trace block reason directly to a trace
// event.
const (
	BlockGeneric         BlockReason = EvGoBlock
	BlockForever                     = EvGoStop
	BlockNet                         = EvGoBlockNet
	BlockSelect                      = EvGoBlockSelect
	BlockCondWait                    = EvGoBlockCond
	BlockSync                        = EvGoBlockSync
	BlockChanSend                    = EvGoBlockSend
	BlockChanRecv                    = EvGoBlockRecv
	BlockGCMarkAssist                = EvGoBlockGC
	BlockGCSweep                     = EvGoBlock
	BlockSystemGoroutine             = EvGoBlock
	BlockPreempted                   = EvGoBlock
	BlockDebugCall                   = EvGoBlock
	BlockUntilGCEnds                 = EvGoBlock
	BlockSleep                       = EvGoSleep
)
