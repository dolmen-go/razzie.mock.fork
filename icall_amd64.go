package mock

import "unsafe"

const maxFrameSize = 256

var icall_array = []any{
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(0, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(1, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(2, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(3, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(4, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(5, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(6, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(7, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(8, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(9, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(10, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(11, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(12, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(13, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(14, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(15, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(16, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(17, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(18, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(19, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(20, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(21, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(22, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(23, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(24, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(25, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(26, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(27, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(28, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(29, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(30, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(31, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(32, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(33, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(34, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(35, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(36, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(37, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(38, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(39, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(40, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(41, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(42, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(43, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(44, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(45, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(46, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(47, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(48, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(49, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(50, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(51, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(52, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(53, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(54, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(55, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(56, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(57, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(58, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(59, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(60, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(61, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(62, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
	func(i1, i2, i3, i4, i5, i6, i7, i8, i9 unsafe.Pointer, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15 float64, frame [maxFrameSize]unsafe.Pointer) (r1, r2, r3, r4, r5, r6, r7, r8, r9 unsafe.Pointer, rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 float64) {
		r, rf := icall(63, i1, []unsafe.Pointer{i1, i2, i3, i4, i5, i6, i7, i8, i9}, []float64{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}, unsafe.Pointer(&frame[0]))
		r1, r2, r3, r4, r5, r6, r7, r8, r9 = r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8]
		rf1, rf2, rf3, rf4, rf5, rf6, rf7, rf8, rf9, rf10, rf11, rf12, rf13, rf14, rf15 = rf[0], rf[1], rf[2], rf[3], rf[4], rf[5], rf[6], rf[7], rf[8], rf[9], rf[10], rf[11], rf[12], rf[13], rf[14]
		return
	},
}
