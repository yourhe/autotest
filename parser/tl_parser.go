// Code generated from tl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tl

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 6, 2, 16, 10, 2, 13, 2, 14, 2, 17, 3, 3, 3, 3, 3, 3, 6, 3, 23, 10, 3,
	13, 3, 14, 3, 24, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 48, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 59, 10, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8,
	2, 4, 6, 8, 10, 12, 2, 3, 4, 2, 19, 19, 23, 23, 2, 77, 2, 15, 3, 2, 2,
	2, 4, 19, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 29, 3, 2, 2, 2, 10, 62, 3,
	2, 2, 2, 12, 64, 3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16,
	17, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 3, 3, 2, 2,
	2, 19, 20, 7, 20, 2, 2, 20, 22, 7, 19, 2, 2, 21, 23, 5, 10, 6, 2, 22, 21,
	3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2,
	25, 5, 3, 2, 2, 2, 26, 27, 7, 20, 2, 2, 27, 28, 7, 19, 2, 2, 28, 7, 3,
	2, 2, 2, 29, 30, 5, 10, 6, 2, 30, 9, 3, 2, 2, 2, 31, 32, 7, 3, 2, 2, 32,
	63, 7, 23, 2, 2, 33, 34, 7, 4, 2, 2, 34, 63, 7, 23, 2, 2, 35, 36, 7, 4,
	2, 2, 36, 63, 7, 17, 2, 2, 37, 38, 7, 4, 2, 2, 38, 63, 7, 9, 2, 2, 39,
	63, 7, 4, 2, 2, 40, 63, 7, 9, 2, 2, 41, 42, 7, 6, 2, 2, 42, 63, 7, 23,
	2, 2, 43, 44, 7, 7, 2, 2, 44, 63, 7, 23, 2, 2, 45, 47, 7, 8, 2, 2, 46,
	48, 7, 19, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 63, 3, 2,
	2, 2, 49, 50, 7, 5, 2, 2, 50, 51, 5, 12, 7, 2, 51, 52, 7, 23, 2, 2, 52,
	63, 3, 2, 2, 2, 53, 54, 7, 10, 2, 2, 54, 63, 9, 2, 2, 2, 55, 63, 7, 11,
	2, 2, 56, 58, 7, 13, 2, 2, 57, 59, 7, 17, 2, 2, 58, 57, 3, 2, 2, 2, 58,
	59, 3, 2, 2, 2, 59, 63, 3, 2, 2, 2, 60, 61, 7, 12, 2, 2, 61, 63, 7, 19,
	2, 2, 62, 31, 3, 2, 2, 2, 62, 33, 3, 2, 2, 2, 62, 35, 3, 2, 2, 2, 62, 37,
	3, 2, 2, 2, 62, 39, 3, 2, 2, 2, 62, 40, 3, 2, 2, 2, 62, 41, 3, 2, 2, 2,
	62, 43, 3, 2, 2, 2, 62, 45, 3, 2, 2, 2, 62, 49, 3, 2, 2, 2, 62, 53, 3,
	2, 2, 2, 62, 55, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63,
	11, 3, 2, 2, 2, 64, 65, 7, 23, 2, 2, 65, 13, 3, 2, 2, 2, 7, 17, 24, 47,
	58, 62,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'open'", "'click'", "'type'", "'find'", "'findAll'", "'text'", "'random'",
	"'iframe'", "'activeElement'", "'sendKey'", "'wait'", "'/'", "'+'", "'-'",
	"", "", "", "'##'",
}
var symbolicNames = []string{
	"", "OPEN", "CLICK", "TYPE", "FIND", "FINDALL", "TEXT", "RANDOM", "IFRAME",
	"ACTIVEELEMENT", "SENDKEY", "WAIT", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
	"ID", "STAGE", "COMMENT", "COMMENTS", "QUOTED_STRING",
}

var ruleNames = []string{
	"start", "stages", "stage", "expressions", "gexpression", "target",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tlParser struct {
	*antlr.BaseParser
}

func NewtlParser(input antlr.TokenStream) *tlParser {
	this := new(tlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tl.g4"

	return this
}

// tlParser tokens.
const (
	tlParserEOF           = antlr.TokenEOF
	tlParserOPEN          = 1
	tlParserCLICK         = 2
	tlParserTYPE          = 3
	tlParserFIND          = 4
	tlParserFINDALL       = 5
	tlParserTEXT          = 6
	tlParserRANDOM        = 7
	tlParserIFRAME        = 8
	tlParserACTIVEELEMENT = 9
	tlParserSENDKEY       = 10
	tlParserWAIT          = 11
	tlParserDIV           = 12
	tlParserADD           = 13
	tlParserSUB           = 14
	tlParserNUMBER        = 15
	tlParserWHITESPACE    = 16
	tlParserID            = 17
	tlParserSTAGE         = 18
	tlParserCOMMENT       = 19
	tlParserCOMMENTS      = 20
	tlParserQUOTED_STRING = 21
)

// tlParser rules.
const (
	tlParserRULE_start       = 0
	tlParserRULE_stages      = 1
	tlParserRULE_stage       = 2
	tlParserRULE_expressions = 3
	tlParserRULE_gexpression = 4
	tlParserRULE_target      = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DS1Context struct {
	*StartContext
}

func NewDS1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DS1Context {
	var p = new(DS1Context)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *DS1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DS1Context) AllStages() []IStagesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStagesContext)(nil)).Elem())
	var tst = make([]IStagesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStagesContext)
		}
	}

	return tst
}

func (s *DS1Context) Stages(i int) IStagesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStagesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStagesContext)
}

func (s *DS1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterDS1(s)
	}
}

func (s *DS1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitDS1(s)
	}
}

func (p *tlParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tlParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewDS1Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tlParserSTAGE {
		{
			p.SetState(12)
			p.Stages()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStagesContext is an interface to support dynamic dispatch.
type IStagesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStagesContext differentiates from other interfaces.
	IsStagesContext()
}

type StagesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStagesContext() *StagesContext {
	var p = new(StagesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_stages
	return p
}

func (*StagesContext) IsStagesContext() {}

func NewStagesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StagesContext {
	var p = new(StagesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_stages

	return p
}

func (s *StagesContext) GetParser() antlr.Parser { return s.parser }

func (s *StagesContext) CopyFrom(ctx *StagesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StagesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RunStagesContext struct {
	*StagesContext
}

func NewRunStagesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunStagesContext {
	var p = new(RunStagesContext)

	p.StagesContext = NewEmptyStagesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StagesContext))

	return p
}

func (s *RunStagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStagesContext) STAGE() antlr.TerminalNode {
	return s.GetToken(tlParserSTAGE, 0)
}

func (s *RunStagesContext) ID() antlr.TerminalNode {
	return s.GetToken(tlParserID, 0)
}

func (s *RunStagesContext) AllGexpression() []IGexpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGexpressionContext)(nil)).Elem())
	var tst = make([]IGexpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGexpressionContext)
		}
	}

	return tst
}

func (s *RunStagesContext) Gexpression(i int) IGexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGexpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGexpressionContext)
}

func (s *RunStagesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterRunStages(s)
	}
}

func (s *RunStagesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitRunStages(s)
	}
}

func (p *tlParser) Stages() (localctx IStagesContext) {
	localctx = NewStagesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tlParserRULE_stages)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewRunStagesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(tlParserSTAGE)
	}
	{
		p.SetState(18)
		p.Match(tlParserID)
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlParserOPEN)|(1<<tlParserCLICK)|(1<<tlParserTYPE)|(1<<tlParserFIND)|(1<<tlParserFINDALL)|(1<<tlParserTEXT)|(1<<tlParserRANDOM)|(1<<tlParserIFRAME)|(1<<tlParserACTIVEELEMENT)|(1<<tlParserSENDKEY)|(1<<tlParserWAIT))) != 0) {
		{
			p.SetState(19)
			p.Gexpression()
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStageContext is an interface to support dynamic dispatch.
type IStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStageContext differentiates from other interfaces.
	IsStageContext()
}

type StageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageContext() *StageContext {
	var p = new(StageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_stage
	return p
}

func (*StageContext) IsStageContext() {}

func NewStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageContext {
	var p = new(StageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_stage

	return p
}

func (s *StageContext) GetParser() antlr.Parser { return s.parser }

func (s *StageContext) CopyFrom(ctx *StageContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DS3Context struct {
	*StageContext
}

func NewDS3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DS3Context {
	var p = new(DS3Context)

	p.StageContext = NewEmptyStageContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StageContext))

	return p
}

func (s *DS3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DS3Context) STAGE() antlr.TerminalNode {
	return s.GetToken(tlParserSTAGE, 0)
}

func (s *DS3Context) ID() antlr.TerminalNode {
	return s.GetToken(tlParserID, 0)
}

func (s *DS3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterDS3(s)
	}
}

func (s *DS3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitDS3(s)
	}
}

func (p *tlParser) Stage() (localctx IStageContext) {
	localctx = NewStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tlParserRULE_stage)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewDS3Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(tlParserSTAGE)
	}
	{
		p.SetState(25)
		p.Match(tlParserID)
	}

	return localctx
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Gexpression() IGexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGexpressionContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *tlParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tlParserRULE_expressions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Gexpression()
	}

	return localctx
}

// IGexpressionContext is an interface to support dynamic dispatch.
type IGexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGexpressionContext differentiates from other interfaces.
	IsGexpressionContext()
}

type GexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGexpressionContext() *GexpressionContext {
	var p = new(GexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_gexpression
	return p
}

func (*GexpressionContext) IsGexpressionContext() {}

func NewGexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GexpressionContext {
	var p = new(GexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_gexpression

	return p
}

func (s *GexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GexpressionContext) CopyFrom(ctx *GexpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *GexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IframeContext struct {
	*GexpressionContext
}

func NewIframeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IframeContext {
	var p = new(IframeContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *IframeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IframeContext) IFRAME() antlr.TerminalNode {
	return s.GetToken(tlParserIFRAME, 0)
}

func (s *IframeContext) ID() antlr.TerminalNode {
	return s.GetToken(tlParserID, 0)
}

func (s *IframeContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *IframeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterIframe(s)
	}
}

func (s *IframeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitIframe(s)
	}
}

type TypeContext struct {
	*GexpressionContext
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(tlParserTYPE, 0)
}

func (s *TypeContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TypeContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitType(s)
	}
}

type ActiveElementContext struct {
	*GexpressionContext
}

func NewActiveElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ActiveElementContext {
	var p = new(ActiveElementContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *ActiveElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActiveElementContext) ACTIVEELEMENT() antlr.TerminalNode {
	return s.GetToken(tlParserACTIVEELEMENT, 0)
}

func (s *ActiveElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterActiveElement(s)
	}
}

func (s *ActiveElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitActiveElement(s)
	}
}

type RandomContext struct {
	*GexpressionContext
}

func NewRandomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RandomContext {
	var p = new(RandomContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *RandomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RandomContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(tlParserRANDOM, 0)
}

func (s *RandomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterRandom(s)
	}
}

func (s *RandomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitRandom(s)
	}
}

type FindAllContext struct {
	*GexpressionContext
}

func NewFindAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FindAllContext {
	var p = new(FindAllContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *FindAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindAllContext) FINDALL() antlr.TerminalNode {
	return s.GetToken(tlParserFINDALL, 0)
}

func (s *FindAllContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *FindAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterFindAll(s)
	}
}

func (s *FindAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitFindAll(s)
	}
}

type FindContext struct {
	*GexpressionContext
}

func NewFindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FindContext {
	var p = new(FindContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *FindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindContext) FIND() antlr.TerminalNode {
	return s.GetToken(tlParserFIND, 0)
}

func (s *FindContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *FindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterFind(s)
	}
}

func (s *FindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitFind(s)
	}
}

type ClickContext struct {
	*GexpressionContext
}

func NewClickContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClickContext {
	var p = new(ClickContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *ClickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClickContext) CLICK() antlr.TerminalNode {
	return s.GetToken(tlParserCLICK, 0)
}

func (s *ClickContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *ClickContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(tlParserNUMBER, 0)
}

func (s *ClickContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(tlParserRANDOM, 0)
}

func (s *ClickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterClick(s)
	}
}

func (s *ClickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitClick(s)
	}
}

type TextContext struct {
	*GexpressionContext
}

func NewTextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextContext {
	var p = new(TextContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(tlParserTEXT, 0)
}

func (s *TextContext) ID() antlr.TerminalNode {
	return s.GetToken(tlParserID, 0)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitText(s)
	}
}

type WaitContext struct {
	*GexpressionContext
}

func NewWaitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WaitContext {
	var p = new(WaitContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *WaitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaitContext) WAIT() antlr.TerminalNode {
	return s.GetToken(tlParserWAIT, 0)
}

func (s *WaitContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(tlParserNUMBER, 0)
}

func (s *WaitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterWait(s)
	}
}

func (s *WaitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitWait(s)
	}
}

type SendKeyContext struct {
	*GexpressionContext
}

func NewSendKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendKeyContext {
	var p = new(SendKeyContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *SendKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendKeyContext) SENDKEY() antlr.TerminalNode {
	return s.GetToken(tlParserSENDKEY, 0)
}

func (s *SendKeyContext) ID() antlr.TerminalNode {
	return s.GetToken(tlParserID, 0)
}

func (s *SendKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterSendKey(s)
	}
}

func (s *SendKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitSendKey(s)
	}
}

type OpenContext struct {
	*GexpressionContext
}

func NewOpenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpenContext {
	var p = new(OpenContext)

	p.GexpressionContext = NewEmptyGexpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GexpressionContext))

	return p
}

func (s *OpenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenContext) OPEN() antlr.TerminalNode {
	return s.GetToken(tlParserOPEN, 0)
}

func (s *OpenContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *OpenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterOpen(s)
	}
}

func (s *OpenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitOpen(s)
	}
}

func (p *tlParser) Gexpression() (localctx IGexpressionContext) {
	localctx = NewGexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tlParserRULE_gexpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpenContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(tlParserOPEN)
		}
		{
			p.SetState(30)
			p.Match(tlParserQUOTED_STRING)
		}

	case 2:
		localctx = NewClickContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(tlParserCLICK)
		}
		{
			p.SetState(32)
			p.Match(tlParserQUOTED_STRING)
		}

	case 3:
		localctx = NewClickContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(tlParserCLICK)
		}
		{
			p.SetState(34)
			p.Match(tlParserNUMBER)
		}

	case 4:
		localctx = NewClickContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
			p.Match(tlParserCLICK)
		}
		{
			p.SetState(36)
			p.Match(tlParserRANDOM)
		}

	case 5:
		localctx = NewClickContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(37)
			p.Match(tlParserCLICK)
		}

	case 6:
		localctx = NewRandomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(38)
			p.Match(tlParserRANDOM)
		}

	case 7:
		localctx = NewFindContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(39)
			p.Match(tlParserFIND)
		}
		{
			p.SetState(40)
			p.Match(tlParserQUOTED_STRING)
		}

	case 8:
		localctx = NewFindAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(41)
			p.Match(tlParserFINDALL)
		}
		{
			p.SetState(42)
			p.Match(tlParserQUOTED_STRING)
		}

	case 9:
		localctx = NewTextContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(43)
			p.Match(tlParserTEXT)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlParserID {
			{
				p.SetState(44)
				p.Match(tlParserID)
			}

		}

	case 10:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(47)
			p.Match(tlParserTYPE)
		}
		{
			p.SetState(48)
			p.Target()
		}
		{
			p.SetState(49)
			p.Match(tlParserQUOTED_STRING)
		}

	case 11:
		localctx = NewIframeContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(51)
			p.Match(tlParserIFRAME)
		}
		{
			p.SetState(52)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tlParserID || _la == tlParserQUOTED_STRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 12:
		localctx = NewActiveElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(53)
			p.Match(tlParserACTIVEELEMENT)
		}

	case 13:
		localctx = NewWaitContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(54)
			p.Match(tlParserWAIT)
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlParserNUMBER {
			{
				p.SetState(55)
				p.Match(tlParserNUMBER)
			}

		}

	case 14:
		localctx = NewSendKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(58)
			p.Match(tlParserSENDKEY)
		}
		{
			p.SetState(59)
			p.Match(tlParserID)
		}

	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(tlParserQUOTED_STRING, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *tlParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tlParserRULE_target)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(tlParserQUOTED_STRING)
	}

	return localctx
}
