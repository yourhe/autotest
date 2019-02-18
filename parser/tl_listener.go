// Code generated from tl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// tlListener is a complete listener for a parse tree produced by tlParser.
type tlListener interface {
	antlr.ParseTreeListener

	// EnterDS1 is called when entering the DS1 production.
	EnterDS1(c *DS1Context)

	// EnterRunStages is called when entering the RunStages production.
	EnterRunStages(c *RunStagesContext)

	// EnterDS3 is called when entering the DS3 production.
	EnterDS3(c *DS3Context)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterOpen is called when entering the Open production.
	EnterOpen(c *OpenContext)

	// EnterClick is called when entering the Click production.
	EnterClick(c *ClickContext)

	// EnterRandom is called when entering the Random production.
	EnterRandom(c *RandomContext)

	// EnterFind is called when entering the Find production.
	EnterFind(c *FindContext)

	// EnterFindAll is called when entering the FindAll production.
	EnterFindAll(c *FindAllContext)

	// EnterText is called when entering the Text production.
	EnterText(c *TextContext)

	// EnterType is called when entering the Type production.
	EnterType(c *TypeContext)

	// EnterIframe is called when entering the Iframe production.
	EnterIframe(c *IframeContext)

	// EnterActiveElement is called when entering the ActiveElement production.
	EnterActiveElement(c *ActiveElementContext)

	// EnterWait is called when entering the Wait production.
	EnterWait(c *WaitContext)

	// EnterSendKey is called when entering the SendKey production.
	EnterSendKey(c *SendKeyContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// ExitDS1 is called when exiting the DS1 production.
	ExitDS1(c *DS1Context)

	// ExitRunStages is called when exiting the RunStages production.
	ExitRunStages(c *RunStagesContext)

	// ExitDS3 is called when exiting the DS3 production.
	ExitDS3(c *DS3Context)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitOpen is called when exiting the Open production.
	ExitOpen(c *OpenContext)

	// ExitClick is called when exiting the Click production.
	ExitClick(c *ClickContext)

	// ExitRandom is called when exiting the Random production.
	ExitRandom(c *RandomContext)

	// ExitFind is called when exiting the Find production.
	ExitFind(c *FindContext)

	// ExitFindAll is called when exiting the FindAll production.
	ExitFindAll(c *FindAllContext)

	// ExitText is called when exiting the Text production.
	ExitText(c *TextContext)

	// ExitType is called when exiting the Type production.
	ExitType(c *TypeContext)

	// ExitIframe is called when exiting the Iframe production.
	ExitIframe(c *IframeContext)

	// ExitActiveElement is called when exiting the ActiveElement production.
	ExitActiveElement(c *ActiveElementContext)

	// ExitWait is called when exiting the Wait production.
	ExitWait(c *WaitContext)

	// ExitSendKey is called when exiting the SendKey production.
	ExitSendKey(c *SendKeyContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)
}
