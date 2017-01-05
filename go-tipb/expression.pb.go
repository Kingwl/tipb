// Code generated by protoc-gen-gogo.
// source: expression.proto
// DO NOT EDIT!

/*
	Package tipb is a generated protocol buffer package.

	It is generated from these files:
		expression.proto
		schema.proto
		select.proto

	It has these top-level messages:
		Expr
		TableInfo
		ColumnInfo
		IndexInfo
		KeyRange
		ByItem
		SelectRequest
		Row
		Error
		SelectResponse
		Chunk
		RowMeta
*/
package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExprType int32

const (
	// Values are encoded bytes.
	ExprType_Null    ExprType = 0
	ExprType_Int64   ExprType = 1
	ExprType_Uint64  ExprType = 2
	ExprType_Float32 ExprType = 3
	ExprType_Float64 ExprType = 4
	ExprType_String  ExprType = 5
	ExprType_Bytes   ExprType = 6
	// Mysql specific types.
	ExprType_MysqlBit      ExprType = 101
	ExprType_MysqlDecimal  ExprType = 102
	ExprType_MysqlDuration ExprType = 103
	ExprType_MysqlEnum     ExprType = 104
	ExprType_MysqlHex      ExprType = 105
	ExprType_MysqlSet      ExprType = 106
	ExprType_MysqlTime     ExprType = 107
	// Encoded value list.
	ExprType_ValueList ExprType = 151
	// Column reference. value is int64 column ID.
	ExprType_ColumnRef ExprType = 201
	// Unary operations, children count 1.
	ExprType_Not    ExprType = 1001
	ExprType_Neg    ExprType = 1002
	ExprType_BitNeg ExprType = 1003
	// Comparison operations.
	ExprType_LT     ExprType = 2001
	ExprType_LE     ExprType = 2002
	ExprType_EQ     ExprType = 2003
	ExprType_NE     ExprType = 2004
	ExprType_GE     ExprType = 2005
	ExprType_GT     ExprType = 2006
	ExprType_NullEQ ExprType = 2007
	// Bit operations.
	ExprType_BitAnd    ExprType = 2101
	ExprType_BitOr     ExprType = 2102
	ExprType_BitXor    ExprType = 2103
	ExprType_LeftShift ExprType = 2104
	ExprType_RighShift ExprType = 2105
	// Arithmatic.
	ExprType_Plus   ExprType = 2201
	ExprType_Minus  ExprType = 2202
	ExprType_Mul    ExprType = 2203
	ExprType_Div    ExprType = 2204
	ExprType_IntDiv ExprType = 2205
	ExprType_Mod    ExprType = 2206
	// Logic operations.
	ExprType_And ExprType = 2301
	ExprType_Or  ExprType = 2302
	ExprType_Xor ExprType = 2303
	// Aggregate functions.
	ExprType_Count       ExprType = 3001
	ExprType_Sum         ExprType = 3002
	ExprType_Avg         ExprType = 3003
	ExprType_Min         ExprType = 3004
	ExprType_Max         ExprType = 3005
	ExprType_First       ExprType = 3006
	ExprType_GroupConcat ExprType = 3007
	// Math functions.
	ExprType_Abs   ExprType = 3101
	ExprType_Pow   ExprType = 3102
	ExprType_Round ExprType = 3103
	// String functions.
	ExprType_Concat         ExprType = 3201
	ExprType_ConcatWS       ExprType = 3202
	ExprType_Left           ExprType = 3203
	ExprType_Length         ExprType = 3204
	ExprType_Lower          ExprType = 3205
	ExprType_Repeat         ExprType = 3206
	ExprType_Replace        ExprType = 3207
	ExprType_Upper          ExprType = 3208
	ExprType_Strcmp         ExprType = 3209
	ExprType_Convert        ExprType = 3210
	ExprType_Cast           ExprType = 3211
	ExprType_Substring      ExprType = 3212
	ExprType_SubstringIndex ExprType = 3213
	ExprType_Locate         ExprType = 3214
	ExprType_Trim           ExprType = 3215
	// Control flow functions.
	ExprType_If     ExprType = 3301
	ExprType_NullIf ExprType = 3302
	ExprType_IfNull ExprType = 3303
	// Time functions.
	ExprType_Date        ExprType = 3401
	ExprType_DateAdd     ExprType = 3402
	ExprType_DateSub     ExprType = 3403
	ExprType_Year        ExprType = 3411
	ExprType_YearWeek    ExprType = 3412
	ExprType_Month       ExprType = 3421
	ExprType_Week        ExprType = 3431
	ExprType_Weekday     ExprType = 3432
	ExprType_WeekOfYear  ExprType = 3433
	ExprType_Day         ExprType = 3441
	ExprType_DayName     ExprType = 3442
	ExprType_DayOfYear   ExprType = 3443
	ExprType_DayOfMonth  ExprType = 3444
	ExprType_DayOfWeek   ExprType = 3445
	ExprType_Hour        ExprType = 3451
	ExprType_Minute      ExprType = 3452
	ExprType_Second      ExprType = 3453
	ExprType_Microsecond ExprType = 3454
	ExprType_Extract     ExprType = 3461
	// Other functions;
	ExprType_Coalesce ExprType = 3501
	ExprType_Greatest ExprType = 3502
	ExprType_Least    ExprType = 3503
	// Other expressions.
	ExprType_In      ExprType = 4001
	ExprType_IsTruth ExprType = 4002
	ExprType_IsNull  ExprType = 4003
	ExprType_ExprRow ExprType = 4004
	ExprType_Like    ExprType = 4005
	ExprType_RLike   ExprType = 4006
	ExprType_Case    ExprType = 4007
)

var ExprType_name = map[int32]string{
	0:    "Null",
	1:    "Int64",
	2:    "Uint64",
	3:    "Float32",
	4:    "Float64",
	5:    "String",
	6:    "Bytes",
	101:  "MysqlBit",
	102:  "MysqlDecimal",
	103:  "MysqlDuration",
	104:  "MysqlEnum",
	105:  "MysqlHex",
	106:  "MysqlSet",
	107:  "MysqlTime",
	151:  "ValueList",
	201:  "ColumnRef",
	1001: "Not",
	1002: "Neg",
	1003: "BitNeg",
	2001: "LT",
	2002: "LE",
	2003: "EQ",
	2004: "NE",
	2005: "GE",
	2006: "GT",
	2007: "NullEQ",
	2101: "BitAnd",
	2102: "BitOr",
	2103: "BitXor",
	2104: "LeftShift",
	2105: "RighShift",
	2201: "Plus",
	2202: "Minus",
	2203: "Mul",
	2204: "Div",
	2205: "IntDiv",
	2206: "Mod",
	2301: "And",
	2302: "Or",
	2303: "Xor",
	3001: "Count",
	3002: "Sum",
	3003: "Avg",
	3004: "Min",
	3005: "Max",
	3006: "First",
	3007: "GroupConcat",
	3101: "Abs",
	3102: "Pow",
	3103: "Round",
	3201: "Concat",
	3202: "ConcatWS",
	3203: "Left",
	3204: "Length",
	3205: "Lower",
	3206: "Repeat",
	3207: "Replace",
	3208: "Upper",
	3209: "Strcmp",
	3210: "Convert",
	3211: "Cast",
	3212: "Substring",
	3213: "SubstringIndex",
	3214: "Locate",
	3215: "Trim",
	3301: "If",
	3302: "NullIf",
	3303: "IfNull",
	3401: "Date",
	3402: "DateAdd",
	3403: "DateSub",
	3411: "Year",
	3412: "YearWeek",
	3421: "Month",
	3431: "Week",
	3432: "Weekday",
	3433: "WeekOfYear",
	3441: "Day",
	3442: "DayName",
	3443: "DayOfYear",
	3444: "DayOfMonth",
	3445: "DayOfWeek",
	3451: "Hour",
	3452: "Minute",
	3453: "Second",
	3454: "Microsecond",
	3461: "Extract",
	3501: "Coalesce",
	3502: "Greatest",
	3503: "Least",
	4001: "In",
	4002: "IsTruth",
	4003: "IsNull",
	4004: "ExprRow",
	4005: "Like",
	4006: "RLike",
	4007: "Case",
}
var ExprType_value = map[string]int32{
	"Null":           0,
	"Int64":          1,
	"Uint64":         2,
	"Float32":        3,
	"Float64":        4,
	"String":         5,
	"Bytes":          6,
	"MysqlBit":       101,
	"MysqlDecimal":   102,
	"MysqlDuration":  103,
	"MysqlEnum":      104,
	"MysqlHex":       105,
	"MysqlSet":       106,
	"MysqlTime":      107,
	"ValueList":      151,
	"ColumnRef":      201,
	"Not":            1001,
	"Neg":            1002,
	"BitNeg":         1003,
	"LT":             2001,
	"LE":             2002,
	"EQ":             2003,
	"NE":             2004,
	"GE":             2005,
	"GT":             2006,
	"NullEQ":         2007,
	"BitAnd":         2101,
	"BitOr":          2102,
	"BitXor":         2103,
	"LeftShift":      2104,
	"RighShift":      2105,
	"Plus":           2201,
	"Minus":          2202,
	"Mul":            2203,
	"Div":            2204,
	"IntDiv":         2205,
	"Mod":            2206,
	"And":            2301,
	"Or":             2302,
	"Xor":            2303,
	"Count":          3001,
	"Sum":            3002,
	"Avg":            3003,
	"Min":            3004,
	"Max":            3005,
	"First":          3006,
	"GroupConcat":    3007,
	"Abs":            3101,
	"Pow":            3102,
	"Round":          3103,
	"Concat":         3201,
	"ConcatWS":       3202,
	"Left":           3203,
	"Length":         3204,
	"Lower":          3205,
	"Repeat":         3206,
	"Replace":        3207,
	"Upper":          3208,
	"Strcmp":         3209,
	"Convert":        3210,
	"Cast":           3211,
	"Substring":      3212,
	"SubstringIndex": 3213,
	"Locate":         3214,
	"Trim":           3215,
	"If":             3301,
	"NullIf":         3302,
	"IfNull":         3303,
	"Date":           3401,
	"DateAdd":        3402,
	"DateSub":        3403,
	"Year":           3411,
	"YearWeek":       3412,
	"Month":          3421,
	"Week":           3431,
	"Weekday":        3432,
	"WeekOfYear":     3433,
	"Day":            3441,
	"DayName":        3442,
	"DayOfYear":      3443,
	"DayOfMonth":     3444,
	"DayOfWeek":      3445,
	"Hour":           3451,
	"Minute":         3452,
	"Second":         3453,
	"Microsecond":    3454,
	"Extract":        3461,
	"Coalesce":       3501,
	"Greatest":       3502,
	"Least":          3503,
	"In":             4001,
	"IsTruth":        4002,
	"IsNull":         4003,
	"ExprRow":        4004,
	"Like":           4005,
	"RLike":          4006,
	"Case":           4007,
}

func (x ExprType) Enum() *ExprType {
	p := new(ExprType)
	*p = x
	return p
}
func (x ExprType) String() string {
	return proto.EnumName(ExprType_name, int32(x))
}
func (x *ExprType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExprType_value, data, "ExprType")
	if err != nil {
		return err
	}
	*x = ExprType(value)
	return nil
}
func (ExprType) EnumDescriptor() ([]byte, []int) { return fileDescriptorExpression, []int{0} }

// Evaluators should implement evaluation functions for every expression type.
type Expr struct {
	Tp               ExprType `protobuf:"varint,1,opt,name=tp,enum=tipb.ExprType" json:"tp"`
	Val              []byte   `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	Children         []*Expr  `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Expr) Reset()                    { *m = Expr{} }
func (m *Expr) String() string            { return proto.CompactTextString(m) }
func (*Expr) ProtoMessage()               {}
func (*Expr) Descriptor() ([]byte, []int) { return fileDescriptorExpression, []int{0} }

func (m *Expr) GetTp() ExprType {
	if m != nil {
		return m.Tp
	}
	return ExprType_Null
}

func (m *Expr) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Expr) GetChildren() []*Expr {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Expr)(nil), "tipb.Expr")
	proto.RegisterEnum("tipb.ExprType", ExprType_name, ExprType_value)
}
func (m *Expr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Expr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintExpression(dAtA, i, uint64(m.Tp))
	if m.Val != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExpression(dAtA, i, uint64(len(m.Val)))
		i += copy(dAtA[i:], m.Val)
	}
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintExpression(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Expression(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Expression(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintExpression(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Expr) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovExpression(uint64(m.Tp))
	if m.Val != nil {
		l = len(m.Val)
		n += 1 + l + sovExpression(uint64(l))
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovExpression(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExpression(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExpression(x uint64) (n int) {
	return sovExpression(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Expr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Expr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Expr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (ExprType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthExpression
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = append(m.Val[:0], dAtA[iNdEx:postIndex]...)
			if m.Val == nil {
				m.Val = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpression
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Expr{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExpression(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpression
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExpression
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExpression
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExpression(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExpression = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpression   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("expression.proto", fileDescriptorExpression) }

var fileDescriptorExpression = []byte{
	// 962 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x54, 0x49, 0x70, 0x1b, 0x45,
	0x17, 0xb6, 0x6c, 0xd9, 0x92, 0xdb, 0xdb, 0xe7, 0xfe, 0x7f, 0xca, 0xd2, 0xb8, 0xca, 0x72, 0x51,
	0x14, 0x65, 0x38, 0x98, 0x2a, 0x93, 0xca, 0xdd, 0xb2, 0x15, 0x47, 0x55, 0xde, 0x22, 0x29, 0x04,
	0x8e, 0xe3, 0x51, 0x4b, 0x6a, 0x3c, 0x9a, 0x1e, 0x7a, 0x7a, 0x6c, 0xe9, 0xc8, 0x12, 0xf6, 0xe5,
	0xc8, 0x9a, 0x50, 0xec, 0x27, 0xe0, 0xc4, 0x12, 0xd6, 0x6b, 0x02, 0x17, 0x20, 0x06, 0x4e, 0x50,
	0x94, 0x29, 0x08, 0x81, 0x13, 0xdb, 0x09, 0x08, 0x54, 0xbf, 0x51, 0x94, 0xdb, 0xf7, 0xbd, 0xfe,
	0xbe, 0xf7, 0xde, 0xbc, 0xf7, 0x24, 0x06, 0xd1, 0x09, 0xb5, 0x88, 0x22, 0xa9, 0x82, 0xc5, 0x50,
	0x2b, 0xa3, 0x78, 0xda, 0xc8, 0x70, 0xc7, 0xf9, 0x7f, 0x53, 0x35, 0x15, 0x05, 0x6e, 0xb1, 0x28,
	0x79, 0xbb, 0xbe, 0xc1, 0xd2, 0xa5, 0x4e, 0xa8, 0xf9, 0x0d, 0x6c, 0xd0, 0x84, 0xb9, 0xd4, 0x7c,
	0x6a, 0x61, 0x72, 0x69, 0x72, 0xd1, 0x1a, 0x16, 0x6d, 0xbc, 0xd6, 0x0d, 0x45, 0x31, 0x7d, 0xfe,
	0xdb, 0xc2, 0x40, 0x65, 0xd0, 0x84, 0x1c, 0x6c, 0x68, 0xcf, 0xf5, 0x73, 0x83, 0xf3, 0xa9, 0x85,
	0xf1, 0x8a, 0x85, 0xfc, 0x46, 0x96, 0xf5, 0x5a, 0xd2, 0xaf, 0x6b, 0x11, 0xe4, 0x86, 0xe6, 0x87,
	0x16, 0xc6, 0x96, 0xd8, 0x35, 0x77, 0xa5, 0xff, 0x76, 0xf3, 0xc1, 0x28, 0xcb, 0x5e, 0x4d, 0xc8,
	0xb3, 0x2c, 0xbd, 0x19, 0xfb, 0x3e, 0x06, 0xf8, 0x28, 0x1b, 0x2e, 0x07, 0xe6, 0xe8, 0x11, 0xa4,
	0x38, 0x63, 0x23, 0x27, 0x25, 0xe1, 0x41, 0x3e, 0xc6, 0x32, 0xc7, 0x7c, 0xe5, 0x9a, 0x5b, 0x97,
	0x30, 0xd4, 0x27, 0x47, 0x8f, 0x20, 0x6d, 0x55, 0x55, 0xa3, 0x65, 0xd0, 0xc4, 0xb0, 0x35, 0x17,
	0xbb, 0x46, 0x44, 0x18, 0xe1, 0xe3, 0x2c, 0xbb, 0xd1, 0x8d, 0xee, 0xf2, 0x8b, 0xd2, 0x40, 0x70,
	0xb0, 0x71, 0x62, 0xab, 0xc2, 0x93, 0x6d, 0xd7, 0x47, 0x83, 0x4f, 0xb3, 0x89, 0x24, 0x12, 0x6b,
	0xd7, 0x48, 0x15, 0xa0, 0xc9, 0x27, 0xd8, 0x28, 0x85, 0x4a, 0x41, 0xdc, 0x46, 0xab, 0x9f, 0xe1,
	0xb8, 0xe8, 0x40, 0xf6, 0x59, 0x55, 0x18, 0xdc, 0xd9, 0x97, 0xd6, 0x64, 0x5b, 0x60, 0x97, 0x4f,
	0xb2, 0xd1, 0xdb, 0x5c, 0x3f, 0x16, 0xeb, 0x32, 0x32, 0x78, 0x32, 0x65, 0xf9, 0x8a, 0xf2, 0xe3,
	0x76, 0x50, 0x11, 0x0d, 0x5c, 0x48, 0xf1, 0x2c, 0x1b, 0xda, 0x54, 0x06, 0x97, 0x33, 0x84, 0x44,
	0x13, 0x3f, 0x67, 0xf8, 0x18, 0x1b, 0x29, 0x4a, 0x63, 0xc9, 0x2f, 0x19, 0x9e, 0x61, 0x83, 0xeb,
	0x35, 0x7c, 0x3e, 0x45, 0xa0, 0x84, 0x2f, 0x08, 0x94, 0x4e, 0xe0, 0x22, 0x81, 0xcd, 0x12, 0x0e,
	0x08, 0xac, 0x95, 0xf0, 0x65, 0x02, 0x6a, 0xf8, 0x6a, 0xca, 0xa6, 0xb0, 0x53, 0x2b, 0x9d, 0xc0,
	0xd7, 0x53, 0xbd, 0x7c, 0xcb, 0x41, 0x1d, 0x6f, 0x81, 0x33, 0x36, 0x5c, 0x94, 0x66, 0x4b, 0xe3,
	0x6d, 0xf4, 0x1e, 0x6e, 0x57, 0x1a, 0xef, 0xc0, 0x76, 0xb6, 0x2e, 0x1a, 0xa6, 0xda, 0x92, 0x0d,
	0x83, 0x77, 0x89, 0x57, 0x64, 0xb3, 0x95, 0xf0, 0x73, 0xe0, 0xa3, 0x2c, 0xbd, 0xed, 0xc7, 0x11,
	0x9e, 0x9e, 0xb6, 0x39, 0x36, 0x64, 0x10, 0x47, 0x78, 0x66, 0xda, 0xb6, 0xbd, 0x11, 0xfb, 0x78,
	0x96, 0xd0, 0xaa, 0xdc, 0xc3, 0x73, 0xd3, 0x36, 0x6f, 0x39, 0x30, 0x96, 0x9c, 0x49, 0x04, 0xaa,
	0x8e, 0xb3, 0x84, 0x6c, 0x13, 0xff, 0x4c, 0xdb, 0x3e, 0xb7, 0x34, 0xae, 0x50, 0xc8, 0x96, 0xff,
	0x97, 0x72, 0xae, 0xa8, 0x38, 0x30, 0x38, 0x37, 0x63, 0xa3, 0xd5, 0xb8, 0x8d, 0xf7, 0x08, 0x2d,
	0xef, 0x35, 0xf1, 0x3e, 0xa1, 0x0d, 0x19, 0xe0, 0x83, 0x04, 0xb9, 0x1d, 0x7c, 0x38, 0x63, 0x3d,
	0xc7, 0xa4, 0x8e, 0x0c, 0x3e, 0x9a, 0xe1, 0x60, 0x63, 0x6b, 0x5a, 0xc5, 0xe1, 0x8a, 0x0a, 0x3c,
	0xd7, 0xe0, 0xe3, 0xc4, 0xbb, 0x13, 0xe1, 0x4c, 0xce, 0xa2, 0x6d, 0xb5, 0x8f, 0xb3, 0x39, 0xeb,
	0xa8, 0xa8, 0x38, 0xa8, 0xe3, 0xf9, 0x9c, 0xed, 0xb2, 0x27, 0xbe, 0x3b, 0xcf, 0x27, 0x58, 0x36,
	0x21, 0xa7, 0xaa, 0xb8, 0x27, 0x6f, 0x3f, 0xd6, 0x0e, 0x03, 0xf7, 0xe6, 0xad, 0x6c, 0x5d, 0x04,
	0x4d, 0xd3, 0xc2, 0x7d, 0x79, 0xeb, 0x5f, 0x57, 0xfb, 0x42, 0xe3, 0x34, 0x3d, 0x54, 0x44, 0x28,
	0x5c, 0x83, 0xfb, 0xf3, 0x7c, 0x9c, 0x65, 0x2a, 0x22, 0xf4, 0x5d, 0x4f, 0xe0, 0x01, 0x92, 0x9d,
	0x0c, 0x43, 0xa1, 0xf1, 0x20, 0xc9, 0xaa, 0x46, 0x7b, 0xed, 0x10, 0x0f, 0x91, 0x6c, 0x45, 0x05,
	0x7b, 0x42, 0x1b, 0x3c, 0x4c, 0x55, 0x56, 0xdc, 0xc8, 0xe0, 0x91, 0xbc, 0x9d, 0x76, 0x35, 0xde,
	0x89, 0x92, 0x73, 0x7d, 0x34, 0xcf, 0xff, 0xc7, 0x26, 0xfb, 0xbc, 0x1c, 0xd4, 0x45, 0x07, 0x8f,
	0x25, 0xad, 0x28, 0xcf, 0x35, 0x02, 0x8f, 0x93, 0xb9, 0xa6, 0x65, 0x1b, 0x4f, 0xe4, 0xed, 0x38,
	0xcb, 0x0d, 0xfc, 0x90, 0xbf, 0xba, 0xf6, 0x72, 0x03, 0x3f, 0x12, 0x29, 0x37, 0xe8, 0xb7, 0x73,
	0x89, 0xd4, 0xab, 0xd6, 0x78, 0xc1, 0xb1, 0x3d, 0x58, 0xb8, 0x5c, 0xaf, 0xe3, 0x93, 0x3e, 0xab,
	0xc6, 0x3b, 0xf8, 0xd4, 0xb1, 0xb2, 0x3b, 0x84, 0xab, 0x71, 0xd1, 0xb1, 0x13, 0xb1, 0xf0, 0x94,
	0x10, 0xbb, 0x38, 0x70, 0x68, 0xe7, 0x2a, 0x30, 0x2d, 0x7c, 0x43, 0x2a, 0x0a, 0x5f, 0x22, 0xbb,
	0x85, 0x75, 0xb7, 0x8b, 0x9f, 0x1c, 0x3e, 0xc5, 0x98, 0x65, 0x5b, 0x0d, 0x4a, 0x72, 0xd9, 0xa1,
	0x9b, 0x70, 0xbb, 0xf8, 0xb5, 0x57, 0xa7, 0xbb, 0xe9, 0xb6, 0x05, 0x7e, 0x73, 0xec, 0xe7, 0xae,
	0xba, 0xdd, 0x9e, 0xee, 0x77, 0x32, 0x12, 0x4f, 0x4a, 0xfc, 0x71, 0x4d, 0x40, 0x75, 0xfe, 0xa4,
	0x92, 0xc7, 0x55, 0xac, 0xf1, 0x97, 0x63, 0xbf, 0xcb, 0x5e, 0x9f, 0x11, 0xf8, 0x9b, 0x48, 0x55,
	0x78, 0xca, 0x9e, 0x95, 0x63, 0x6f, 0x60, 0x43, 0x7a, 0x5a, 0x45, 0x49, 0xe4, 0x0a, 0x55, 0x2d,
	0x75, 0x8c, 0x76, 0x3d, 0x83, 0xd3, 0xb3, 0xc9, 0x92, 0x5d, 0x5f, 0x44, 0x9e, 0xc0, 0xeb, 0x44,
	0xd7, 0xb4, 0x70, 0x8d, 0x88, 0x0c, 0xde, 0x98, 0xa5, 0xdd, 0x0a, 0xbb, 0x8e, 0x37, 0x67, 0x69,
	0xa2, 0x01, 0x5e, 0x28, 0xd8, 0x04, 0xe5, 0xa8, 0xa6, 0x63, 0xd3, 0xc2, 0x8b, 0x05, 0x1a, 0x69,
	0x44, 0x23, 0x7d, 0xa9, 0x90, 0xe4, 0x0e, 0x75, 0x45, 0xed, 0xe3, 0xe5, 0x02, 0x5d, 0x8c, 0xdc,
	0x15, 0x78, 0xa5, 0x40, 0x47, 0x46, 0xf8, 0xd5, 0x42, 0x6f, 0xc5, 0x02, 0xaf, 0x15, 0x8a, 0x37,
	0x9d, 0x3f, 0x9c, 0x4b, 0x7d, 0x76, 0x38, 0x97, 0xfa, 0xee, 0x70, 0x2e, 0xf5, 0xd4, 0xf7, 0x73,
	0x03, 0xec, 0x3a, 0x4f, 0xb5, 0x17, 0x43, 0x19, 0x34, 0x3d, 0x37, 0x5c, 0x34, 0xb2, 0xbe, 0x43,
	0x7f, 0x87, 0xdb, 0xa9, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x77, 0x5c, 0xd9, 0x51, 0x98, 0x05,
	0x00, 0x00,
}
