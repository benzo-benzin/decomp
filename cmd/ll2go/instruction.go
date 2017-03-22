package main

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// insts converts the given LLVM IR instructions to a corresponding list of Go
// statements.
func (d *decompiler) insts(insts []ir.Instruction) []ast.Stmt {
	var stmts []ast.Stmt
	for _, inst := range insts {
		if _, ok := inst.(*ir.InstPhi); ok {
			// PHI instructions are handled during the pre-processing of basic
			// blocks.
			continue
		}
		stmts = append(stmts, d.inst(inst))
	}
	return stmts
}

// inst converts the given LLVM IR instruction to a corresponding Go statement.
func (d *decompiler) inst(inst ir.Instruction) ast.Stmt {
	switch inst := inst.(type) {
	// Binary instructions
	case *ir.InstAdd:
		return d.instAdd(inst)
	case *ir.InstFAdd:
		return d.instFAdd(inst)
	case *ir.InstSub:
		return d.instSub(inst)
	case *ir.InstFSub:
		return d.instFSub(inst)
	case *ir.InstMul:
		return d.instMul(inst)
	case *ir.InstFMul:
		return d.instFMul(inst)
	case *ir.InstUDiv:
		return d.instUDiv(inst)
	case *ir.InstSDiv:
		return d.instSDiv(inst)
	case *ir.InstFDiv:
		return d.instFDiv(inst)
	case *ir.InstURem:
		return d.instURem(inst)
	case *ir.InstSRem:
		return d.instSRem(inst)
	case *ir.InstFRem:
		return d.instFRem(inst)
	// Bitwise instructions
	case *ir.InstShl:
		return d.instShl(inst)
	case *ir.InstLShr:
		return d.instLShr(inst)
	case *ir.InstAShr:
		return d.instAShr(inst)
	case *ir.InstAnd:
		return d.instAnd(inst)
	case *ir.InstOr:
		return d.instOr(inst)
	case *ir.InstXor:
		return d.instXor(inst)
	// Memory instructions
	case *ir.InstAlloca:
		return d.instAlloca(inst)
	case *ir.InstLoad:
		return d.instLoad(inst)
	case *ir.InstStore:
		return d.instStore(inst)
	case *ir.InstGetElementPtr:
		return d.instGetElementPtr(inst)
	// Conversion instructions
	case *ir.InstTrunc:
		return d.instTrunc(inst)
	case *ir.InstZExt:
		return d.instZExt(inst)
	case *ir.InstSExt:
		return d.instSExt(inst)
	case *ir.InstFPTrunc:
		return d.instFPTrunc(inst)
	case *ir.InstFPExt:
		return d.instFPExt(inst)
	case *ir.InstFPToUI:
		return d.instFPToUI(inst)
	case *ir.InstFPToSI:
		return d.instFPToSI(inst)
	case *ir.InstUIToFP:
		return d.instUIToFP(inst)
	case *ir.InstSIToFP:
		return d.instSIToFP(inst)
	case *ir.InstPtrToInt:
		return d.instPtrToInt(inst)
	case *ir.InstIntToPtr:
		return d.instIntToPtr(inst)
	case *ir.InstBitCast:
		return d.instBitCast(inst)
	case *ir.InstAddrSpaceCast:
		return d.instAddrSpaceCast(inst)
	// Other instructions
	case *ir.InstICmp:
		return d.instICmp(inst)
	case *ir.InstFCmp:
		return d.instFCmp(inst)
	case *ir.InstPhi:
		panic(fmt.Sprintf("unexpected PHI instruction `%v`", inst))
	case *ir.InstSelect:
		return d.instSelect(inst)
	case *ir.InstCall:
		return d.instCall(inst)
	default:
		panic(fmt.Sprintf("support for instruction %T not yet implemented", inst))
	}
}

// instAdd converts the given LLVM IR add instruction to a corresponding Go
// statement.
func (d *decompiler) instAdd(inst *ir.InstAdd) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.ADD, inst.Y)
}

// instFAdd converts the given LLVM IR fadd instruction to a corresponding Go
// statement.
func (d *decompiler) instFAdd(inst *ir.InstFAdd) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.ADD, inst.Y)
}

// instSub converts the given LLVM IR sub instruction to a corresponding Go
// statement.
func (d *decompiler) instSub(inst *ir.InstSub) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.SUB, inst.Y)
}

// instFSub converts the given LLVM IR fsub instruction to a corresponding Go
// statement.
func (d *decompiler) instFSub(inst *ir.InstFSub) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.SUB, inst.Y)
}

// instMul converts the given LLVM IR mul instruction to a corresponding Go
// statement.
func (d *decompiler) instMul(inst *ir.InstMul) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.MUL, inst.Y)

}

// instFMul converts the given LLVM IR fmul instruction to a corresponding Go
// statement.
func (d *decompiler) instFMul(inst *ir.InstFMul) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.MUL, inst.Y)
}

// instUDiv converts the given LLVM IR udiv instruction to a corresponding Go
// statement.
func (d *decompiler) instUDiv(inst *ir.InstUDiv) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.QUO, inst.Y)
}

// instSDiv converts the given LLVM IR sdiv instruction to a corresponding Go
// statement.
func (d *decompiler) instSDiv(inst *ir.InstSDiv) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.QUO, inst.Y)
}

// instFDiv converts the given LLVM IR fdiv instruction to a corresponding Go
// statement.
func (d *decompiler) instFDiv(inst *ir.InstFDiv) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.QUO, inst.Y)
}

// instURem converts the given LLVM IR urem instruction to a corresponding Go
// statement.
func (d *decompiler) instURem(inst *ir.InstURem) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.REM, inst.Y)
}

// instSRem converts the given LLVM IR srem instruction to a corresponding Go
// statement.
func (d *decompiler) instSRem(inst *ir.InstSRem) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.REM, inst.Y)
}

// instFRem converts the given LLVM IR frem instruction to a corresponding Go
// statement.
func (d *decompiler) instFRem(inst *ir.InstFRem) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.REM, inst.Y)
}

// instShl converts the given LLVM IR shl instruction to a corresponding Go
// statement.
func (d *decompiler) instShl(inst *ir.InstShl) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.SHL, inst.Y)
}

// instLShr converts the given LLVM IR lshr instruction to a corresponding Go
// statement.
func (d *decompiler) instLShr(inst *ir.InstLShr) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.SHR, inst.Y)
}

// instAShr converts the given LLVM IR ashr instruction to a corresponding Go
// statement.
func (d *decompiler) instAShr(inst *ir.InstAShr) ast.Stmt {
	// TODO: Differentiate between logical shift right and arithmetic shift
	// right.
	return d.instBinaryOp(inst.Name, inst.X, token.SHR, inst.Y)
}

// instAnd converts the given LLVM IR and instruction to a corresponding Go
// statement.
func (d *decompiler) instAnd(inst *ir.InstAnd) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.AND, inst.Y)
}

// instOr converts the given LLVM IR or instruction to a corresponding Go
// statement.
func (d *decompiler) instOr(inst *ir.InstOr) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.OR, inst.Y)
}

// instXor converts the given LLVM IR xor instruction to a corresponding Go
// statement.
func (d *decompiler) instXor(inst *ir.InstXor) ast.Stmt {
	return d.instBinaryOp(inst.Name, inst.X, token.XOR, inst.Y)
}

// instAlloca converts the given LLVM IR alloca instruction to a corresponding
// Go statement.
func (d *decompiler) instAlloca(inst *ir.InstAlloca) ast.Stmt {
	panic("not yet implemented")
}

// instLoad converts the given LLVM IR load instruction to a corresponding Go
// statement.
func (d *decompiler) instLoad(inst *ir.InstLoad) ast.Stmt {
	panic("not yet implemented")
}

// instStore converts the given LLVM IR store instruction to a corresponding Go
// statement.
func (d *decompiler) instStore(inst *ir.InstStore) ast.Stmt {
	panic("not yet implemented")
}

// instGetElementPtr converts the given LLVM IR getelementptr instruction to a
// corresponding Go statement.
func (d *decompiler) instGetElementPtr(inst *ir.InstGetElementPtr) ast.Stmt {
	panic("not yet implemented")
}

// instTrunc converts the given LLVM IR trunc instruction to a corresponding Go
// statement.
func (d *decompiler) instTrunc(inst *ir.InstTrunc) ast.Stmt {
	panic("not yet implemented")
}

// instZExt converts the given LLVM IR zext instruction to a corresponding Go
// statement.
func (d *decompiler) instZExt(inst *ir.InstZExt) ast.Stmt {
	panic("not yet implemented")
}

// instSExt converts the given LLVM IR sext instruction to a corresponding Go
// statement.
func (d *decompiler) instSExt(inst *ir.InstSExt) ast.Stmt {
	panic("not yet implemented")
}

// instFPTrunc converts the given LLVM IR fptrunc instruction to a corresponding
// Go statement.
func (d *decompiler) instFPTrunc(inst *ir.InstFPTrunc) ast.Stmt {
	panic("not yet implemented")
}

// instFPExt converts the given LLVM IR fpext instruction to a corresponding Go
// statement.
func (d *decompiler) instFPExt(inst *ir.InstFPExt) ast.Stmt {
	panic("not yet implemented")
}

// instFPToUI converts the given LLVM IR fptoui instruction to a corresponding
// Go statement.
func (d *decompiler) instFPToUI(inst *ir.InstFPToUI) ast.Stmt {
	panic("not yet implemented")
}

// instFPToSI converts the given LLVM IR fptosi instruction to a corresponding
// Go statement.
func (d *decompiler) instFPToSI(inst *ir.InstFPToSI) ast.Stmt {
	panic("not yet implemented")
}

// instUIToFP converts the given LLVM IR uitofp instruction to a corresponding
// Go statement.
func (d *decompiler) instUIToFP(inst *ir.InstUIToFP) ast.Stmt {
	panic("not yet implemented")
}

// instSIToFP converts the given LLVM IR sitofp instruction to a corresponding
// Go statement.
func (d *decompiler) instSIToFP(inst *ir.InstSIToFP) ast.Stmt {
	panic("not yet implemented")
}

// instPtrToInt converts the given LLVM IR ptrtoint instruction to a
// corresponding Go statement.
func (d *decompiler) instPtrToInt(inst *ir.InstPtrToInt) ast.Stmt {
	panic("not yet implemented")
}

// instIntToPtr converts the given LLVM IR inttoptr instruction to a
// corresponding Go statement.
func (d *decompiler) instIntToPtr(inst *ir.InstIntToPtr) ast.Stmt {
	panic("not yet implemented")
}

// instBitCast converts the given LLVM IR bitcast instruction to a corresponding
// Go statement.
func (d *decompiler) instBitCast(inst *ir.InstBitCast) ast.Stmt {
	panic("not yet implemented")
}

// instAddrSpaceCast converts the given LLVM IR addrspacecast instruction to a
// corresponding Go statement.
func (d *decompiler) instAddrSpaceCast(inst *ir.InstAddrSpaceCast) ast.Stmt {
	panic("not yet implemented")
}

// instICmp converts the given LLVM IR icmp instruction to a corresponding Go
// statement.
func (d *decompiler) instICmp(inst *ir.InstICmp) ast.Stmt {
	op := intPred(inst.Pred)
	return d.instBinaryOp(inst.Name, inst.X, op, inst.Y)
}

// instFCmp converts the given LLVM IR fcmp instruction to a corresponding Go
// statement.
func (d *decompiler) instFCmp(inst *ir.InstFCmp) ast.Stmt {
	op := floatPred(inst.Pred)
	return d.instBinaryOp(inst.Name, inst.X, op, inst.Y)
}

// instSelect converts the given LLVM IR select instruction to a corresponding
// Go statement.
func (d *decompiler) instSelect(inst *ir.InstSelect) ast.Stmt {
	panic("not yet implemented")
}

// instCall converts the given LLVM IR call instruction to a corresponding Go
// statement.
func (d *decompiler) instCall(inst *ir.InstCall) ast.Stmt {
	panic("not yet implemented")
}

// instBinaryOp converts the given LLVM IR binary operation to a corresponding
// Go statement.
func (d *decompiler) instBinaryOp(name string, x value.Value, op token.Token, y value.Value) ast.Stmt {
	// TODO: Handle type (inst.Typ).
	expr := &ast.BinaryExpr{
		X:  d.value(x),
		Op: op,
		Y:  d.value(y),
	}
	return &ast.AssignStmt{
		Lhs: []ast.Expr{d.local(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{expr},
	}
}

// intPred converts the given LLVM IR integer predicate to a corresponding Go
// token.
func intPred(pred ir.IntPred) token.Token {
	// TODO: Differentiate between unsigned and signed.
	switch pred {
	case ir.IntEQ:
		return token.EQL
	case ir.IntNE:
		return token.NEQ
	case ir.IntUGT:
		return token.GTR
	case ir.IntUGE:
		return token.GEQ
	case ir.IntULT:
		return token.LSS
	case ir.IntULE:
		return token.LEQ
	case ir.IntSGT:
		return token.GTR
	case ir.IntSGE:
		return token.GEQ
	case ir.IntSLT:
		return token.LSS
	case ir.IntSLE:
		return token.LEQ
	default:
		panic(fmt.Sprintf("support for integer predicate %v not yet implemented", pred))
	}
}

// floatPred converts the given LLVM IR floating-point predicate to a
// corresponding Go token.
func floatPred(pred ir.FloatPred) token.Token {
	// TODO: Differentiate between ordered and unordered.
	switch pred {
	case ir.FloatFalse:
		panic(`support for floating-point predicate "false" not yet implemented`)
	case ir.FloatOEQ:
		return token.EQL
	case ir.FloatOGT:
		return token.GTR
	case ir.FloatOGE:
		return token.GEQ
	case ir.FloatOLT:
		return token.LSS
	case ir.FloatOLE:
		return token.LEQ
	case ir.FloatONE:
		return token.NEQ
	case ir.FloatORD:
		panic(`support for floating-point predicate "ord" not yet implemented`)
	case ir.FloatUEQ:
		return token.EQL
	case ir.FloatUGT:
		return token.GTR
	case ir.FloatUGE:
		return token.GEQ
	case ir.FloatULT:
		return token.LSS
	case ir.FloatULE:
		return token.LEQ
	case ir.FloatUNE:
		return token.NEQ
	case ir.FloatUNO:
		panic(`support for floating-point predicate "uno" not yet implemented`)
	case ir.FloatTrue:
		panic(`support for floating-point predicate "true" not yet implemented`)
	default:
		panic(fmt.Sprintf("support for floating-point predicate %v not yet implemented", pred))
	}
}