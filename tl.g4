grammar tl;

// Tokens
OPEN: 'open';
CLICK: 'click';
TYPE: 'type';
FIND: 'find';
FINDALL: 'findAll';
TEXT: 'text';
RANDOM: 'random';
IFRAME: 'iframe';
ACTIVEELEMENT: 'activeElement';
SENDKEY: 'sendKey';
WAIT: 'wait';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
ID: [a-zA-Z]+;
STAGE: '##';

COMMENT: '//' .*? '\n' -> skip;
COMMENTS: '/*' .*? '*/' -> skip;

fragment ESCAPED_QUOTE: '\\"';
fragment STRING: '"' ( ESCAPED_QUOTE | ~('\n' | '\r'))*? '"';
QUOTED_STRING: STRING;


// Rules
start: stages+ # DS1;

stages: STAGE ID gexpression+ # RunStages;

stage: STAGE ID # DS3;
expressions: gexpression;

gexpression:
	// gexpression op = ('*' | '/') gexpression	# MulDiv
	// | gexpression op = ('+' | '-') gexpression	# AddSub
	OPEN QUOTED_STRING						# Open
	| CLICK QUOTED_STRING						# Click
	| CLICK NUMBER								# Click // TODO
	| CLICK RANDOM								# Click // TODO
	| CLICK										# Click
	| RANDOM									# Random
	| FIND QUOTED_STRING						# Find
	| FINDALL QUOTED_STRING						# FindAll
	| TEXT ID?									# Text
	| TYPE target QUOTED_STRING					# Type
	| IFRAME (ID | QUOTED_STRING)				# Iframe
	| ACTIVEELEMENT								# ActiveElement
	| WAIT NUMBER?								# Wait
	| SENDKEY ID								# SendKey;
// | QUOTED_STRING #Click | command QUOTED_STRING #Commands;
// command: OPEN | CLICK | TYPE;

target: QUOTED_STRING;