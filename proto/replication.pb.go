// Code generated by protoc-gen-gogo.
// source: replication.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "gogoproto"

import fmt "fmt"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// KeyserverStep denotes the input to a single step of the keyserver state
// machine. Linearizable high-availability replication is achieved by
// replicating an in-order log of all steps and having each replica reproduce
// the state from them.
type KeyserverStep struct {
	UID uint64 `protobuf:"varint,1,opt,proto3" json:"UID,omitempty"`
	// Update is appended to the log when it is received from a client and
	// has passed pre-validation. However, since pre-validation is not
	// final, "success" should not be returned to the client until after the
	// update has been applied and ratified.
	// update is applied to the keyserver state as soon as it has been
	// committed to the log.
	Update *UpdateRequest `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// EpochDelimiter is appended to the log by a leader (not necessarily
	// unique) node when at least the duration EPOCH_INTERVAL_MIN and at
	// most EPOCH_INTERVAL_MAX after the previous epoch_delimiter has passed.
	// Between these times, an epoch delimiter is appended as soon as an
	// update is committed.
	// As the leader requirement for appending an epoch_delimiter is soft,
	// it may happen that an epoch delimiter with a epoch number not higher
	// than the previous gets committed to the log. It must be ignored.
	EpochDelimiter *EpochDelimiter `protobuf:"bytes,3,opt,name=epoch_delimiter" json:"epoch_delimiter,omitempty"`
	// ReplicaSigned for the last epoch is appended to the log
	// when the epoch_delimiter is committed.
	// After some majority of the replicas has signed the next
	// TimestampedEpochHead; their signatures make up the keyserver
	// signature. As combining signatures is deterministic, no new log
	// entry is appended to record that.
	ReplicaSigned *SignedEpochHead `protobuf:"bytes,4,opt,name=replica_signed" json:"replica_signed,omitempty"`
	// VerifierSigned is appended for each new SignedEpochHead received
	// from a verifier; these are used to provide proof of verification to
	// clients.
	VerifierSigned *SignedEpochHead `protobuf:"bytes,5,opt,name=verifier_signed" json:"verifier_signed,omitempty"`
}

func (m *KeyserverStep) Reset()      { *m = KeyserverStep{} }
func (*KeyserverStep) ProtoMessage() {}

func (m *KeyserverStep) GetUpdate() *UpdateRequest {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *KeyserverStep) GetEpochDelimiter() *EpochDelimiter {
	if m != nil {
		return m.EpochDelimiter
	}
	return nil
}

func (m *KeyserverStep) GetReplicaSigned() *SignedEpochHead {
	if m != nil {
		return m.ReplicaSigned
	}
	return nil
}

func (m *KeyserverStep) GetVerifierSigned() *SignedEpochHead {
	if m != nil {
		return m.VerifierSigned
	}
	return nil
}

type EpochDelimiter struct {
	EpochNumber uint64    `protobuf:"varint,1,opt,name=epoch_number,proto3" json:"epoch_number,omitempty"`
	Timestamp   Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp"`
}

func (m *EpochDelimiter) Reset()      { *m = EpochDelimiter{} }
func (*EpochDelimiter) ProtoMessage() {}

func (m *EpochDelimiter) GetTimestamp() Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return Timestamp{}
}

type Replica struct {
	PublicKey *PublicKey `protobuf:"bytes,1,opt,name=public_key" json:"public_key,omitempty"`
	Addr      string     `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *Replica) Reset()      { *m = Replica{} }
func (*Replica) ProtoMessage() {}

func (m *Replica) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (this *KeyserverStep) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep)
	if !ok {
		return fmt.Errorf("that is not of type *KeyserverStep")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStepbut is not nil && this == nil")
	}
	if this.UID != that1.UID {
		return fmt.Errorf("UID this(%v) Not Equal that(%v)", this.UID, that1.UID)
	}
	if !this.Update.Equal(that1.Update) {
		return fmt.Errorf("Update this(%v) Not Equal that(%v)", this.Update, that1.Update)
	}
	if !this.EpochDelimiter.Equal(that1.EpochDelimiter) {
		return fmt.Errorf("EpochDelimiter this(%v) Not Equal that(%v)", this.EpochDelimiter, that1.EpochDelimiter)
	}
	if !this.ReplicaSigned.Equal(that1.ReplicaSigned) {
		return fmt.Errorf("ReplicaSigned this(%v) Not Equal that(%v)", this.ReplicaSigned, that1.ReplicaSigned)
	}
	if !this.VerifierSigned.Equal(that1.VerifierSigned) {
		return fmt.Errorf("VerifierSigned this(%v) Not Equal that(%v)", this.VerifierSigned, that1.VerifierSigned)
	}
	return nil
}
func (this *KeyserverStep) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UID != that1.UID {
		return false
	}
	if !this.Update.Equal(that1.Update) {
		return false
	}
	if !this.EpochDelimiter.Equal(that1.EpochDelimiter) {
		return false
	}
	if !this.ReplicaSigned.Equal(that1.ReplicaSigned) {
		return false
	}
	if !this.VerifierSigned.Equal(that1.VerifierSigned) {
		return false
	}
	return true
}
func (this *EpochDelimiter) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EpochDelimiter)
	if !ok {
		return fmt.Errorf("that is not of type *EpochDelimiter")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EpochDelimiter but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EpochDelimiterbut is not nil && this == nil")
	}
	if this.EpochNumber != that1.EpochNumber {
		return fmt.Errorf("EpochNumber this(%v) Not Equal that(%v)", this.EpochNumber, that1.EpochNumber)
	}
	if !this.Timestamp.Equal(&that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	return nil
}
func (this *EpochDelimiter) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EpochDelimiter)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.EpochNumber != that1.EpochNumber {
		return false
	}
	if !this.Timestamp.Equal(&that1.Timestamp) {
		return false
	}
	return true
}
func (this *Replica) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Replica)
	if !ok {
		return fmt.Errorf("that is not of type *Replica")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Replica but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Replicabut is not nil && this == nil")
	}
	if !this.PublicKey.Equal(that1.PublicKey) {
		return fmt.Errorf("PublicKey this(%v) Not Equal that(%v)", this.PublicKey, that1.PublicKey)
	}
	if this.Addr != that1.Addr {
		return fmt.Errorf("Addr this(%v) Not Equal that(%v)", this.Addr, that1.Addr)
	}
	return nil
}
func (this *Replica) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Replica)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.PublicKey.Equal(that1.PublicKey) {
		return false
	}
	if this.Addr != that1.Addr {
		return false
	}
	return true
}
func (this *KeyserverStep) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.KeyserverStep{` +
		`UID:` + fmt.Sprintf("%#v", this.UID),
		`Update:` + fmt.Sprintf("%#v", this.Update),
		`EpochDelimiter:` + fmt.Sprintf("%#v", this.EpochDelimiter),
		`ReplicaSigned:` + fmt.Sprintf("%#v", this.ReplicaSigned),
		`VerifierSigned:` + fmt.Sprintf("%#v", this.VerifierSigned) + `}`}, ", ")
	return s
}
func (this *EpochDelimiter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.EpochDelimiter{` +
		`EpochNumber:` + fmt.Sprintf("%#v", this.EpochNumber),
		`Timestamp:` + strings.Replace(this.Timestamp.GoString(), `&`, ``, 1) + `}`}, ", ")
	return s
}
func (this *Replica) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.Replica{` +
		`PublicKey:` + fmt.Sprintf("%#v", this.PublicKey),
		`Addr:` + fmt.Sprintf("%#v", this.Addr) + `}`}, ", ")
	return s
}
func valueToGoStringReplication(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringReplication(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *KeyserverStep) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyserverStep) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintReplication(data, i, uint64(m.UID))
	}
	if m.Update != nil {
		data[i] = 0x12
		i++
		i = encodeVarintReplication(data, i, uint64(m.Update.Size()))
		n1, err := m.Update.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.EpochDelimiter != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochDelimiter.Size()))
		n2, err := m.EpochDelimiter.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ReplicaSigned != nil {
		data[i] = 0x22
		i++
		i = encodeVarintReplication(data, i, uint64(m.ReplicaSigned.Size()))
		n3, err := m.ReplicaSigned.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.VerifierSigned != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintReplication(data, i, uint64(m.VerifierSigned.Size()))
		n4, err := m.VerifierSigned.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *EpochDelimiter) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EpochDelimiter) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochNumber))
	}
	data[i] = 0x22
	i++
	i = encodeVarintReplication(data, i, uint64(m.Timestamp.Size()))
	n5, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *Replica) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Replica) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PublicKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintReplication(data, i, uint64(m.PublicKey.Size()))
		n6, err := m.PublicKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Addr) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintReplication(data, i, uint64(len(m.Addr)))
		i += copy(data[i:], m.Addr)
	}
	return i, nil
}

func encodeFixed64Replication(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Replication(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintReplication(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKeyserverStep(r randyReplication, easy bool) *KeyserverStep {
	this := &KeyserverStep{}
	this.UID = uint64(uint64(r.Uint32()))
	if r.Intn(10) != 0 {
		this.Update = NewPopulatedUpdateRequest(r, easy)
	}
	if r.Intn(10) != 0 {
		this.EpochDelimiter = NewPopulatedEpochDelimiter(r, easy)
	}
	if r.Intn(10) != 0 {
		this.ReplicaSigned = NewPopulatedSignedEpochHead(r, easy)
	}
	if r.Intn(10) != 0 {
		this.VerifierSigned = NewPopulatedSignedEpochHead(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedEpochDelimiter(r randyReplication, easy bool) *EpochDelimiter {
	this := &EpochDelimiter{}
	this.EpochNumber = uint64(uint64(r.Uint32()))
	v1 := NewPopulatedTimestamp(r, easy)
	this.Timestamp = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedReplica(r randyReplication, easy bool) *Replica {
	this := &Replica{}
	if r.Intn(10) != 0 {
		this.PublicKey = NewPopulatedPublicKey(r, easy)
	}
	this.Addr = randStringReplication(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyReplication interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneReplication(r randyReplication) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringReplication(r randyReplication) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneReplication(r)
	}
	return string(tmps)
}
func randUnrecognizedReplication(r randyReplication, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldReplication(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldReplication(data []byte, r randyReplication, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateReplication(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateReplication(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateReplication(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateReplication(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateReplication(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateReplication(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateReplication(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *KeyserverStep) Size() (n int) {
	var l int
	_ = l
	if m.UID != 0 {
		n += 1 + sovReplication(uint64(m.UID))
	}
	if m.Update != nil {
		l = m.Update.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.EpochDelimiter != nil {
		l = m.EpochDelimiter.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.ReplicaSigned != nil {
		l = m.ReplicaSigned.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.VerifierSigned != nil {
		l = m.VerifierSigned.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}

func (m *EpochDelimiter) Size() (n int) {
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovReplication(uint64(m.EpochNumber))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovReplication(uint64(l))
	return n
}

func (m *Replica) Size() (n int) {
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}

func sovReplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReplication(x uint64) (n int) {
	return sovReplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *KeyserverStep) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Update:` + strings.Replace(fmt.Sprintf("%v", this.Update), "UpdateRequest", "UpdateRequest", 1) + `,`,
		`EpochDelimiter:` + strings.Replace(fmt.Sprintf("%v", this.EpochDelimiter), "EpochDelimiter", "EpochDelimiter", 1) + `,`,
		`ReplicaSigned:` + strings.Replace(fmt.Sprintf("%v", this.ReplicaSigned), "SignedEpochHead", "SignedEpochHead", 1) + `,`,
		`VerifierSigned:` + strings.Replace(fmt.Sprintf("%v", this.VerifierSigned), "SignedEpochHead", "SignedEpochHead", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EpochDelimiter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EpochDelimiter{`,
		`EpochNumber:` + fmt.Sprintf("%v", this.EpochNumber) + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Timestamp", "Timestamp", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Replica) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Replica{`,
		`PublicKey:` + strings.Replace(fmt.Sprintf("%v", this.PublicKey), "PublicKey", "PublicKey", 1) + `,`,
		`Addr:` + fmt.Sprintf("%v", this.Addr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReplication(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *KeyserverStep) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			m.UID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Update == nil {
				m.Update = &UpdateRequest{}
			}
			if err := m.Update.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDelimiter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EpochDelimiter == nil {
				m.EpochDelimiter = &EpochDelimiter{}
			}
			if err := m.EpochDelimiter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaSigned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReplicaSigned == nil {
				m.ReplicaSigned = &SignedEpochHead{}
			}
			if err := m.ReplicaSigned.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifierSigned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifierSigned == nil {
				m.VerifierSigned = &SignedEpochHead{}
			}
			if err := m.VerifierSigned.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipReplication(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *EpochDelimiter) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EpochNumber |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipReplication(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Replica) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipReplication(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipReplication(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipReplication(data[start:])
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
