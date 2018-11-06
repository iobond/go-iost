package tx

import (
	"github.com/golang/protobuf/proto"
	"github.com/iost-official/go-iost/common"
)

// StatusCode status code of transaction execution result
type StatusCode int32

// tx execution result
const (
	Success StatusCode = iota
	ErrorGasRunOut
	ErrorBalanceNotEnough
	ErrorParamter // parameter mismatch when calling function
	ErrorRuntime  // runtime error
	ErrorTimeout
	ErrorTxFormat         // tx format errors
	ErrorDuplicateSetCode // more than one set code action in a tx
	ErrorUnknown          // other errors
)

// Status status of transaction execution result, including code and message
type Status struct {
	Code    StatusCode
	Message string
}

// ReceiptType type of single receipt
type ReceiptType int32

const (
	// SystemDefined system receipt, recording info of calling a method
	SystemDefined ReceiptType = iota
	// UserDefined user defined receipt, usually a json string
	UserDefined
)

// Receipt generated when applying transaction
type Receipt struct {
	ReceiptRaw
}

// TxReceipt Transaction Receipt
type TxReceipt struct { //nolint:golint
	TxHash   []byte
	GasUsage int64
	RAMUsage map[string]int64
	/*
		CpuTimeUsage    uint64
		NetUsage    uint64
	*/
	Status   Status
	Returns  []*Return
	Receipts []Receipt
}

// NewTxReceipt generate tx receipt for a tx hash
func NewTxReceipt(txHash []byte) TxReceipt {
	var status = Status{
		Code:    Success,
		Message: "",
	}
	return TxReceipt{
		TxHash:   txHash,
		GasUsage: 0,
		Status:   status,
		Returns:  []*Return{},
		Receipts: []Receipt{},
	}
}

// ToTxReceiptRaw convert TxReceipt to proto buf data structure
func (r *TxReceipt) ToTxReceiptRaw() *TxReceiptRaw {
	tr := &TxReceiptRaw{
		TxHash:   r.TxHash,
		GasUsage: r.GasUsage,
		Status: &StatusRaw{
			Code:    int32(r.Status.Code),
			Message: r.Status.Message,
		},
		Returns: r.Returns,
	}
	for _, re := range r.Receipts {
		tr.Receipts = append(tr.Receipts, &re.ReceiptRaw)
	}
	return tr
}

// Encode TxReceipt as byte array
func (r *TxReceipt) Encode() []byte {
	b, err := proto.Marshal(r.ToTxReceiptRaw())
	if err != nil {
		panic(err)
	}
	return b
}

// FromTxReceiptRaw convert TxReceipt from proto buf data structure
func (r *TxReceipt) FromTxReceiptRaw(tr *TxReceiptRaw) {
	r.TxHash = tr.TxHash
	r.GasUsage = tr.GasUsage
	r.Status = Status{
		Code:    StatusCode(tr.Status.Code),
		Message: tr.Status.Message,
	}
	r.Returns = tr.Returns
	r.Receipts = []Receipt{}
	for _, re := range tr.Receipts {
		r.Receipts = append(r.Receipts, Receipt{*re})
	}
}

// Decode TxReceipt from byte array
func (r *TxReceipt) Decode(b []byte) error {
	tr := &TxReceiptRaw{}
	err := proto.Unmarshal(b, tr)
	if err != nil {
		return err
	}
	r.FromTxReceiptRaw(tr)
	return nil
}

// Hash return byte hash
func (r *TxReceipt) Hash() []byte {
	return common.Sha3(r.Encode())
}

func (r *TxReceipt) String() string {
	tr := &TxReceiptRaw{
		TxHash:   r.TxHash,
		GasUsage: r.GasUsage,
		Status: &StatusRaw{
			Code:    int32(r.Status.Code),
			Message: r.Status.Message,
		},
		Returns: r.Returns,
	}
	return tr.String()
}
