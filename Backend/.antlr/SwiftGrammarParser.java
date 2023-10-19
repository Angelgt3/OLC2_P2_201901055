// Generated from c:\Users\angge\Documents\GitHub\OLC2_P1_201901055\Backend\SwiftGrammar.g4 by ANTLR 4.8

    import "Backend/interfaces"
    import "Backend/environment"
    import "Backend/expressions"
    import "Backend/instructions"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, TSTRING=3, BOOL=4, CHAR=5, VAR=6, LET=7, TRU=8, FAL=9, 
		PRINT=10, IF=11, ELSE=12, SWITCH=13, CASE=14, DEFAULT=15, BREAK=16, CONTINUE=17, 
		RETURN=18, WHILE=19, FOR=20, IN=21, GUARD=22, APPEND=23, REMOVELAST=24, 
		REMOVE=25, ISEMPTY=26, COUNT=27, REPEATING=28, STRUCT=29, SELF=30, MUTATING=31, 
		FUNC=32, AT=33, INOUT=34, NIL=35, NUMBER=36, CHARACTER=37, STRING=38, 
		ID=39, DIF=40, IG_IG=41, NOT=42, OR=43, AND=44, IG=45, MAY_IG=46, MEN_IG=47, 
		MAYOR=48, MENOR=49, MOD=50, MUL=51, DIV=52, ADD=53, SUB=54, PARIZQ=55, 
		PARDER=56, LLAVEIZQ=57, LLAVEDER=58, CORIZQ=59, CORDER=60, DOSP=61, PUNTO=62, 
		COMA=63, INTCE=64, FLECHA=65, AMP=66, PCOMA=67, GBAJO=68, TRESP=69, WHITESPACE=70, 
		COMMENT=71, LINE_COMMENT=72;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_structCreation = 3, 
		RULE_listStructDec = 4, RULE_funcstmt = 5, RULE_listParamsFunc = 6, RULE_funcallstmt = 7, 
		RULE_funcallestmt = 8, RULE_printstmt = 9, RULE_variablestmt = 10, RULE_typestmt = 11, 
		RULE_type2stmt = 12, RULE_funvecstmt = 13, RULE_guardstmt = 14, RULE_ifstmt = 15, 
		RULE_elifs = 16, RULE_switchstmt = 17, RULE_cases = 18, RULE_whilestmt = 19, 
		RULE_forstmt = 20, RULE_expr = 21, RULE_primitives = 22, RULE_listParams = 23, 
		RULE_listArray = 24, RULE_listAceso = 25, RULE_listStructExp = 26;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "structCreation", "listStructDec", "funcstmt", 
			"listParamsFunc", "funcallstmt", "funcallestmt", "printstmt", "variablestmt", 
			"typestmt", "type2stmt", "funvecstmt", "guardstmt", "ifstmt", "elifs", 
			"switchstmt", "cases", "whilestmt", "forstmt", "expr", "primitives", 
			"listParams", "listArray", "listAceso", "listStructExp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'var'", 
			"'let'", "'true'", "'false'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'break'", "'continue'", "'return'", "'while'", 
			"'for'", "'in'", "'guard'", "'append'", "'removeLast'", "'remove'", "'isEmpty'", 
			"'count'", "'repeating'", "'struct'", "'self'", "'mutating'", "'func'", 
			"'at'", "'inout'", "'nil'", null, null, null, null, "'!='", "'=='", "'!'", 
			"'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", 
			"'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "'.'", 
			"','", "'?'", "'->'", "'&'", "';'", "'_'", "'...'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "TSTRING", "BOOL", "CHAR", "VAR", "LET", "TRU", 
			"FAL", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", 
			"RETURN", "WHILE", "FOR", "IN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", 
			"ISEMPTY", "COUNT", "REPEATING", "STRUCT", "SELF", "MUTATING", "FUNC", 
			"AT", "INOUT", "NIL", "NUMBER", "CHARACTER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", 
			"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORIZQ", "CORDER", "DOSP", "PUNTO", "COMA", "INTCE", "FLECHA", "AMP", 
			"PCOMA", "GBAJO", "TRESP", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			((SContext)_localctx).block = block();
			setState(55);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(59); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(58);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(61); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public VariablestmtContext variablestmt;
		public GuardstmtContext guardstmt;
		public IfstmtContext ifstmt;
		public SwitchstmtContext switchstmt;
		public WhilestmtContext whilestmt;
		public ForstmtContext forstmt;
		public FunvecstmtContext funvecstmt;
		public Token BREAK;
		public Token CONTINUE;
		public Token RETURN;
		public ExprContext expr;
		public FuncallstmtContext funcallstmt;
		public FuncstmtContext funcstmt;
		public StructCreationContext structCreation;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public TerminalNode PCOMA() { return getToken(SwiftGrammarParser.PCOMA, 0); }
		public VariablestmtContext variablestmt() {
			return getRuleContext(VariablestmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public FunvecstmtContext funvecstmt() {
			return getRuleContext(FunvecstmtContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FuncallstmtContext funcallstmt() {
			return getRuleContext(FuncallstmtContext.class,0);
		}
		public FuncstmtContext funcstmt() {
			return getRuleContext(FuncstmtContext.class,0);
		}
		public StructCreationContext structCreation() {
			return getRuleContext(StructCreationContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(121);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				((InstructionContext)_localctx).printstmt = printstmt();
				setState(67);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(66);
					match(PCOMA);
					}
					break;
				}
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				((InstructionContext)_localctx).variablestmt = variablestmt();
				setState(73);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
				case 1:
					{
					setState(72);
					match(PCOMA);
					}
					break;
				}
				 _localctx.inst = ((InstructionContext)_localctx).variablestmt.vari
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guinst 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(80);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(83);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.swinst 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(86);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinst 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(89);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinst 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(92);
				((InstructionContext)_localctx).funvecstmt = funvecstmt();
				 _localctx.inst = ((InstructionContext)_localctx).funvecstmt.fvecisnt 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(95);
				((InstructionContext)_localctx).BREAK = match(BREAK);
				setState(97);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(96);
					match(PCOMA);
					}
					break;
				}
				 _localctx.inst = instructions.NewBreak((((InstructionContext)_localctx).BREAK!=null?((InstructionContext)_localctx).BREAK.getLine():0), (((InstructionContext)_localctx).BREAK!=null?((InstructionContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(100);
				((InstructionContext)_localctx).CONTINUE = match(CONTINUE);
				setState(102);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
				case 1:
					{
					setState(101);
					match(PCOMA);
					}
					break;
				}
				 _localctx.inst = instructions.NewContinue((((InstructionContext)_localctx).CONTINUE!=null?((InstructionContext)_localctx).CONTINUE.getLine():0), (((InstructionContext)_localctx).CONTINUE!=null?((InstructionContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(105);
				((InstructionContext)_localctx).RETURN = match(RETURN);
				setState(106);
				((InstructionContext)_localctx).expr = expr(0);
				setState(108);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
				case 1:
					{
					setState(107);
					match(PCOMA);
					}
					break;
				}
				 _localctx.inst = instructions.NewReturn((((InstructionContext)_localctx).RETURN!=null?((InstructionContext)_localctx).RETURN.getLine():0), (((InstructionContext)_localctx).RETURN!=null?((InstructionContext)_localctx).RETURN.getCharPositionInLine():0), ((InstructionContext)_localctx).expr.e) 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(112);
				((InstructionContext)_localctx).funcallstmt = funcallstmt();
				 _localctx.inst = ((InstructionContext)_localctx).funcallstmt.cfun 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(115);
				((InstructionContext)_localctx).funcstmt = funcstmt();
				 _localctx.inst = ((InstructionContext)_localctx).funcstmt.fun 
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(118);
				((InstructionContext)_localctx).structCreation = structCreation();
				 _localctx.inst = ((InstructionContext)_localctx).structCreation.dec 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StructCreationContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token STRUCT;
		public Token ID;
		public ListStructDecContext listStructDec;
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public StructCreationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structCreation; }
	}

	public final StructCreationContext structCreation() throws RecognitionException {
		StructCreationContext _localctx = new StructCreationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_structCreation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			((StructCreationContext)_localctx).STRUCT = match(STRUCT);
			setState(124);
			((StructCreationContext)_localctx).ID = match(ID);
			setState(125);
			match(LLAVEIZQ);
			setState(126);
			((StructCreationContext)_localctx).listStructDec = listStructDec(0);
			setState(127);
			match(LLAVEDER);
			 _localctx.dec = instructions.NewStruct((((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getLine():0), (((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getCharPositionInLine():0), (((StructCreationContext)_localctx).ID!=null?((StructCreationContext)_localctx).ID.getText():null), ((StructCreationContext)_localctx).listStructDec.l) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListStructDecContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructDecContext list;
		public Token ID;
		public TypestmtContext typestmt;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public TypestmtContext typestmt() {
			return getRuleContext(TypestmtContext.class,0);
		}
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public ListStructDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructDec; }
	}

	public final ListStructDecContext listStructDec() throws RecognitionException {
		return listStructDec(0);
	}

	private ListStructDecContext listStructDec(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructDecContext _localctx = new ListStructDecContext(_ctx, _parentState);
		ListStructDecContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_listStructDec, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(131);
				match(VAR);
				setState(132);
				((ListStructDecContext)_localctx).ID = match(ID);
				setState(133);
				match(DOSP);
				setState(134);
				((ListStructDecContext)_localctx).typestmt = typestmt();

				        var arr []interface{}
				        newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).typestmt.type)
				        arr = append(arr, newParams)
				        _localctx.l = arr
				    
				}
				break;
			case 2:
				{
				 _localctx.l = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(149);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructDecContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
					setState(140);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(141);
					match(VAR);
					setState(142);
					((ListStructDecContext)_localctx).ID = match(ID);
					setState(143);
					match(DOSP);
					setState(144);
					((ListStructDecContext)_localctx).typestmt = typestmt();

					                  var arr []interface{}
					                  newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).typestmt.type)
					                  arr = append(((ListStructDecContext)_localctx).list.l, newParams)
					                  _localctx.l = arr
					              
					}
					} 
				}
				setState(151);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class FuncstmtContext extends ParserRuleContext {
		public interfaces.Instruction fun;
		public Token FUNC;
		public Token ID;
		public BlockContext block;
		public ListParamsFuncContext listParamsFunc;
		public TypestmtContext typestmt;
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public TerminalNode FLECHA() { return getToken(SwiftGrammarParser.FLECHA, 0); }
		public TypestmtContext typestmt() {
			return getRuleContext(TypestmtContext.class,0);
		}
		public FuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcstmt; }
	}

	public final FuncstmtContext funcstmt() throws RecognitionException {
		FuncstmtContext _localctx = new FuncstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_funcstmt);
		try {
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(153);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(154);
				match(PARIZQ);
				setState(155);
				match(PARDER);
				setState(156);
				match(LLAVEIZQ);
				setState(157);
				((FuncstmtContext)_localctx).block = block();
				setState(158);
				match(LLAVEDER);
				 _localctx.fun = instructions.NewDeclarationFunc((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), environment.NULL,  nil ,((FuncstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(162);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(163);
				match(PARIZQ);
				setState(164);
				((FuncstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(165);
				match(PARDER);
				setState(166);
				match(LLAVEIZQ);
				setState(167);
				((FuncstmtContext)_localctx).block = block();
				setState(168);
				match(LLAVEDER);
				 _localctx.fun = instructions.NewDeclarationFunc((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), environment.NULL, ((FuncstmtContext)_localctx).listParamsFunc.lpf, ((FuncstmtContext)_localctx).block.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(171);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(172);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(173);
				match(PARIZQ);
				setState(174);
				match(PARDER);
				setState(175);
				match(FLECHA);
				setState(176);
				((FuncstmtContext)_localctx).typestmt = typestmt();
				setState(177);
				match(LLAVEIZQ);
				setState(178);
				((FuncstmtContext)_localctx).block = block();
				setState(179);
				match(LLAVEDER);
				 _localctx.fun = instructions.NewDeclarationFunc((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), ((FuncstmtContext)_localctx).typestmt.type,  nil ,((FuncstmtContext)_localctx).block.blk) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(182);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(183);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(184);
				match(PARIZQ);
				setState(185);
				((FuncstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(186);
				match(PARDER);
				setState(187);
				match(FLECHA);
				setState(188);
				((FuncstmtContext)_localctx).typestmt = typestmt();
				setState(189);
				match(LLAVEIZQ);
				setState(190);
				((FuncstmtContext)_localctx).block = block();
				setState(191);
				match(LLAVEDER);
				 _localctx.fun = instructions.NewDeclarationFunc((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), ((FuncstmtContext)_localctx).typestmt.type, ((FuncstmtContext)_localctx).listParamsFunc.lpf, ((FuncstmtContext)_localctx).block.blk) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListParamsFuncContext extends ParserRuleContext {
		public []interface{} lpf;
		public ListParamsFuncContext list;
		public Token id1;
		public Token id2;
		public Token INOUT;
		public TypestmtContext typestmt;
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public TypestmtContext typestmt() {
			return getRuleContext(TypestmtContext.class,0);
		}
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode INOUT() { return getToken(SwiftGrammarParser.INOUT, 0); }
		public TerminalNode GBAJO() { return getToken(SwiftGrammarParser.GBAJO, 0); }
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public ListParamsFuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsFunc; }
	}

	public final ListParamsFuncContext listParamsFunc() throws RecognitionException {
		return listParamsFunc(0);
	}

	private ListParamsFuncContext listParamsFunc(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsFuncContext _localctx = new ListParamsFuncContext(_ctx, _parentState);
		ListParamsFuncContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_listParamsFunc, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(198);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
				case 1:
					{
					setState(197);
					((ListParamsFuncContext)_localctx).id1 = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==ID || _la==GBAJO) ) {
						((ListParamsFuncContext)_localctx).id1 = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				setState(200);
				((ListParamsFuncContext)_localctx).id2 = match(ID);
				setState(201);
				match(DOSP);
				setState(203);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(202);
					((ListParamsFuncContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(205);
				((ListParamsFuncContext)_localctx).typestmt = typestmt();

				    _localctx.lpf = []interface{}{}
				    
				    newParam := instructions.NewParamsDeclaration((((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getLine():0), (((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getCharPositionInLine():0), (((ListParamsFuncContext)_localctx).id1!=null?((ListParamsFuncContext)_localctx).id1.getText():null),(((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getText():null),(((ListParamsFuncContext)_localctx).INOUT!=null?((ListParamsFuncContext)_localctx).INOUT.getText():null), ((ListParamsFuncContext)_localctx).typestmt.type)
				    _localctx.lpf = append(_localctx.lpf, newParam)
				    
				}
				break;
			case 2:
				{
				 _localctx.lpf = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(226);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsFuncContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsFunc);
					setState(211);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(212);
					match(COMA);
					setState(214);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
					case 1:
						{
						setState(213);
						((ListParamsFuncContext)_localctx).id1 = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ID || _la==GBAJO) ) {
							((ListParamsFuncContext)_localctx).id1 = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						}
						break;
					}
					setState(216);
					((ListParamsFuncContext)_localctx).id2 = match(ID);
					setState(217);
					match(DOSP);
					setState(219);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==INOUT) {
						{
						setState(218);
						((ListParamsFuncContext)_localctx).INOUT = match(INOUT);
						}
					}

					setState(221);
					((ListParamsFuncContext)_localctx).typestmt = typestmt();

					              var arr []interface{}
					              newParam := instructions.NewParamsDeclaration((((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getLine():0), (((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getCharPositionInLine():0), (((ListParamsFuncContext)_localctx).id1!=null?((ListParamsFuncContext)_localctx).id1.getText():null),(((ListParamsFuncContext)_localctx).id2!=null?((ListParamsFuncContext)_localctx).id2.getText():null),(((ListParamsFuncContext)_localctx).INOUT!=null?((ListParamsFuncContext)_localctx).INOUT.getText():null), ((ListParamsFuncContext)_localctx).typestmt.type)
					              arr = append(((ListParamsFuncContext)_localctx).list.lpf, newParam)
					              _localctx.lpf = arr
					              
					}
					} 
				}
				setState(228);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class FuncallstmtContext extends ParserRuleContext {
		public interfaces.Instruction cfun;
		public Token ID;
		public ListParamsContext listParams;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public FuncallstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcallstmt; }
	}

	public final FuncallstmtContext funcallstmt() throws RecognitionException {
		FuncallstmtContext _localctx = new FuncallstmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_funcallstmt);
		try {
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(229);
				((FuncallstmtContext)_localctx).ID = match(ID);
				setState(230);
				match(PARIZQ);
				setState(231);
				match(PARDER);
				 _localctx.cfun = instructions.NewFunCall((((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getLine():0), (((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getCharPositionInLine():0), (((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getText():null), nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(233);
				((FuncallstmtContext)_localctx).ID = match(ID);
				setState(234);
				match(PARIZQ);
				setState(235);
				((FuncallstmtContext)_localctx).listParams = listParams(0);
				setState(236);
				match(PARDER);
				 _localctx.cfun = instructions.NewFunCall((((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getLine():0), (((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getCharPositionInLine():0), (((FuncallstmtContext)_localctx).ID!=null?((FuncallstmtContext)_localctx).ID.getText():null), ((FuncallstmtContext)_localctx).listParams.l) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncallestmtContext extends ParserRuleContext {
		public interfaces.Expression cefun;
		public Token ID;
		public ListParamsContext listParams;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public FuncallestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcallestmt; }
	}

	public final FuncallestmtContext funcallestmt() throws RecognitionException {
		FuncallestmtContext _localctx = new FuncallestmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_funcallestmt);
		try {
			setState(251);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(241);
				((FuncallestmtContext)_localctx).ID = match(ID);
				setState(242);
				match(PARIZQ);
				setState(243);
				match(PARDER);
				 _localctx.cefun = expressions.NewFunCallE((((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getLine():0), (((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getCharPositionInLine():0), (((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getText():null), nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				((FuncallestmtContext)_localctx).ID = match(ID);
				setState(246);
				match(PARIZQ);
				setState(247);
				((FuncallestmtContext)_localctx).listParams = listParams(0);
				setState(248);
				match(PARDER);
				 _localctx.cefun = expressions.NewFunCallE((((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getLine():0), (((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getCharPositionInLine():0), (((FuncallestmtContext)_localctx).ID!=null?((FuncallestmtContext)_localctx).ID.getText():null), ((FuncallestmtContext)_localctx).listParams.l) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ExprContext expr;
		public Token PARIZQ;
		public ListParamsContext listParams;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_printstmt);
		try {
			setState(265);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(253);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(254);
				match(PARIZQ);
				setState(255);
				((PrintstmtContext)_localctx).expr = expr(0);
				setState(256);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(259);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(260);
				((PrintstmtContext)_localctx).PARIZQ = match(PARIZQ);
				setState(261);
				((PrintstmtContext)_localctx).listParams = listParams(0);
				setState(262);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),expressions.NewArray((((PrintstmtContext)_localctx).PARIZQ!=null?((PrintstmtContext)_localctx).PARIZQ.getLine():0), (((PrintstmtContext)_localctx).PARIZQ!=null?((PrintstmtContext)_localctx).PARIZQ.getCharPositionInLine():0), ((PrintstmtContext)_localctx).listParams.l))
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariablestmtContext extends ParserRuleContext {
		public interfaces.Instruction vari;
		public Token VAR;
		public Token ID;
		public TypestmtContext typestmt;
		public ExprContext expr;
		public Token LET;
		public PrimitivesContext primitives;
		public ExprContext e1;
		public ExprContext e2;
		public Token CORIZQ;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public TypestmtContext typestmt() {
			return getRuleContext(TypestmtContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode INTCE() { return getToken(SwiftGrammarParser.INTCE, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public PrimitivesContext primitives() {
			return getRuleContext(PrimitivesContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public VariablestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variablestmt; }
	}

	public final VariablestmtContext variablestmt() throws RecognitionException {
		VariablestmtContext _localctx = new VariablestmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_variablestmt);
		try {
			setState(376);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(267);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(268);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(269);
				match(DOSP);
				setState(270);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(271);
				match(IG);
				setState(272);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, ((VariablestmtContext)_localctx).expr.e, true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(275);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(276);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(277);
				match(IG);
				setState(278);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), environment.NULL, ((VariablestmtContext)_localctx).expr.e, true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(281);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(282);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(283);
				match(DOSP);
				setState(284);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(285);
				match(INTCE);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, expressions.NewPrimitive((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), nil, environment.NULL), true) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(288);
				((VariablestmtContext)_localctx).LET = match(LET);
				setState(289);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(290);
				match(IG);
				setState(291);
				((VariablestmtContext)_localctx).primitives = primitives();
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getLine():0), (((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), environment.NULL, ((VariablestmtContext)_localctx).primitives.p, false) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(294);
				((VariablestmtContext)_localctx).LET = match(LET);
				setState(295);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(296);
				match(DOSP);
				setState(297);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(298);
				match(IG);
				setState(299);
				((VariablestmtContext)_localctx).primitives = primitives();
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getLine():0), (((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, ((VariablestmtContext)_localctx).primitives.p, false) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(302);
				((VariablestmtContext)_localctx).LET = match(LET);
				setState(303);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(304);
				match(IG);
				setState(305);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getLine():0), (((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), environment.NULL, ((VariablestmtContext)_localctx).expr.e, false) 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(308);
				((VariablestmtContext)_localctx).LET = match(LET);
				setState(309);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(310);
				match(DOSP);
				setState(311);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(312);
				match(IG);
				setState(313);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getLine():0), (((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, ((VariablestmtContext)_localctx).expr.e, false) 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(316);
				((VariablestmtContext)_localctx).LET = match(LET);
				setState(317);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(318);
				match(IG);
				setState(319);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getLine():0), (((VariablestmtContext)_localctx).LET!=null?((VariablestmtContext)_localctx).LET.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), environment.NULL, ((VariablestmtContext)_localctx).expr.e, false) 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(322);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(323);
				match(IG);
				setState(324);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewAssigment((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).expr.e) 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(327);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(328);
				match(CORIZQ);
				setState(329);
				((VariablestmtContext)_localctx).e1 = expr(0);
				setState(330);
				match(CORDER);
				setState(331);
				match(IG);
				setState(332);
				((VariablestmtContext)_localctx).e2 = expr(0);
				 _localctx.vari = instructions.NewAssigmentVec((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).e1.e, ((VariablestmtContext)_localctx).e2.e) 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(335);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(336);
				((VariablestmtContext)_localctx).e1 = expr(0);
				setState(337);
				match(IG);
				setState(338);
				((VariablestmtContext)_localctx).e2 = expr(0);
				 _localctx.vari = instructions.NewAssigmentMa((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).e1.e, ((VariablestmtContext)_localctx).e2.e) 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(341);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(342);
				match(ADD);
				setState(343);
				match(IG);
				setState(344);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewSumAssigment((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).expr.e) 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(347);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(348);
				match(SUB);
				setState(349);
				match(IG);
				setState(350);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewResAssigment((((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).expr.e) 
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(353);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(354);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(355);
				match(DOSP);
				setState(356);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(357);
				match(IG);
				setState(358);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, ((VariablestmtContext)_localctx).expr.e, true) 
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(361);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(362);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(363);
				match(DOSP);
				setState(364);
				((VariablestmtContext)_localctx).typestmt = typestmt();
				setState(365);
				match(IG);
				setState(366);
				((VariablestmtContext)_localctx).CORIZQ = match(CORIZQ);
				setState(367);
				match(CORDER);

				    _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), ((VariablestmtContext)_localctx).typestmt.type, 
				    expressions.NewArray((((VariablestmtContext)_localctx).CORIZQ!=null?((VariablestmtContext)_localctx).CORIZQ.getLine():0), (((VariablestmtContext)_localctx).CORIZQ!=null?((VariablestmtContext)_localctx).CORIZQ.getCharPositionInLine():0), nil),
				    true)

				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(370);
				((VariablestmtContext)_localctx).VAR = match(VAR);
				setState(371);
				((VariablestmtContext)_localctx).ID = match(ID);
				setState(372);
				match(IG);
				setState(373);
				((VariablestmtContext)_localctx).expr = expr(0);
				 _localctx.vari = instructions.NewDeclaration((((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getLine():0), (((VariablestmtContext)_localctx).VAR!=null?((VariablestmtContext)_localctx).VAR.getCharPositionInLine():0), (((VariablestmtContext)_localctx).ID!=null?((VariablestmtContext)_localctx).ID.getText():null), environment.STRUCT , ((VariablestmtContext)_localctx).expr.e,true) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypestmtContext extends ParserRuleContext {
		public environment.TipoExpresion type;
		public Type2stmtContext type2stmt;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode TSTRING() { return getToken(SwiftGrammarParser.TSTRING, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(SwiftGrammarParser.CHAR, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public Type2stmtContext type2stmt() {
			return getRuleContext(Type2stmtContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TypestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typestmt; }
	}

	public final TypestmtContext typestmt() throws RecognitionException {
		TypestmtContext _localctx = new TypestmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_typestmt);
		try {
			setState(393);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(378);
				match(INT);
				 _localctx.type = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(380);
				match(FLOAT);
				 _localctx.type = environment.FLOAT 
				}
				break;
			case TSTRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(382);
				match(TSTRING);
				 _localctx.type = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(384);
				match(BOOL);
				 _localctx.type = environment.BOOLEAN 
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(386);
				match(CHAR);
				 _localctx.type = environment.CHAR 
				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(388);
				match(CORIZQ);
				setState(389);
				((TypestmtContext)_localctx).type2stmt = type2stmt();
				setState(390);
				match(CORDER);
				 _localctx.type = ((TypestmtContext)_localctx).type2stmt.type2 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type2stmtContext extends ParserRuleContext {
		public environment.TipoExpresion type2;
		public Type2stmtContext type2stmt;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode TSTRING() { return getToken(SwiftGrammarParser.TSTRING, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(SwiftGrammarParser.CHAR, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public Type2stmtContext type2stmt() {
			return getRuleContext(Type2stmtContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public Type2stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type2stmt; }
	}

	public final Type2stmtContext type2stmt() throws RecognitionException {
		Type2stmtContext _localctx = new Type2stmtContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_type2stmt);
		try {
			setState(410);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(395);
				match(INT);
				 _localctx.type2 = environment.A_INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(397);
				match(FLOAT);
				 _localctx.type2 = environment.A_FLOAT 
				}
				break;
			case TSTRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(399);
				match(TSTRING);
				 _localctx.type2 = environment.A_STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(401);
				match(BOOL);
				 _localctx.type2 = environment.A_BOOLEAN 
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(403);
				match(CHAR);
				 _localctx.type2 = environment.A_CHAR 
				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(405);
				match(CORIZQ);
				setState(406);
				((Type2stmtContext)_localctx).type2stmt = type2stmt();
				setState(407);
				match(CORDER);
				 _localctx.type2 = ((Type2stmtContext)_localctx).type2stmt.type2 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunvecstmtContext extends ParserRuleContext {
		public interfaces.Instruction fvecisnt;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(SwiftGrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode REMOVELAST() { return getToken(SwiftGrammarParser.REMOVELAST, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public TerminalNode AT() { return getToken(SwiftGrammarParser.AT, 0); }
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public FunvecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funvecstmt; }
	}

	public final FunvecstmtContext funvecstmt() throws RecognitionException {
		FunvecstmtContext _localctx = new FunvecstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_funvecstmt);
		try {
			setState(436);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(412);
				((FunvecstmtContext)_localctx).ID = match(ID);
				setState(413);
				match(PUNTO);
				setState(414);
				match(APPEND);
				setState(415);
				match(PARIZQ);
				setState(416);
				((FunvecstmtContext)_localctx).expr = expr(0);
				setState(417);
				match(PARDER);
				 _localctx.fvecisnt = instructions.NewAppend((((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getText():null) , ((FunvecstmtContext)_localctx).expr.e) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				((FunvecstmtContext)_localctx).ID = match(ID);
				setState(421);
				match(PUNTO);
				setState(422);
				match(REMOVELAST);
				setState(423);
				match(PARIZQ);
				setState(424);
				match(PARDER);
				 _localctx.fvecisnt = instructions.NewRemoveLast((((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(426);
				((FunvecstmtContext)_localctx).ID = match(ID);
				setState(427);
				match(PUNTO);
				setState(428);
				match(REMOVE);
				setState(429);
				match(PARIZQ);
				setState(430);
				match(AT);
				setState(431);
				match(DOSP);
				setState(432);
				((FunvecstmtContext)_localctx).expr = expr(0);
				setState(433);
				match(PARDER);
				 _localctx.fvecisnt = instructions.NewRemove((((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((FunvecstmtContext)_localctx).ID!=null?((FunvecstmtContext)_localctx).ID.getText():null) , ((FunvecstmtContext)_localctx).expr.e) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GuardstmtContext extends ParserRuleContext {
		public interfaces.Instruction guinst;
		public Token GUARD;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(439);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(440);
			match(ELSE);
			setState(441);
			match(LLAVEIZQ);
			setState(442);
			((GuardstmtContext)_localctx).block = block();
			setState(443);
			match(LLAVEDER);
			 _localctx.guinst = instructions.NewGuard((((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getLine():0), (((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardstmtContext)_localctx).expr.e, ((GuardstmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ifinst;
		public Token IF;
		public ExprContext expr;
		public BlockContext block;
		public BlockContext bif;
		public BlockContext belse;
		public ElifsContext elifs;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(SwiftGrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(SwiftGrammarParser.LLAVEIZQ, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(SwiftGrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(SwiftGrammarParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public ElifsContext elifs() {
			return getRuleContext(ElifsContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_ifstmt);
		try {
			setState(484);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(446);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(447);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(448);
				match(LLAVEIZQ);
				setState(449);
				((IfstmtContext)_localctx).block = block();
				setState(450);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil,nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(453);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(454);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(455);
				match(LLAVEIZQ);
				setState(456);
				((IfstmtContext)_localctx).bif = block();
				setState(457);
				match(LLAVEDER);
				setState(458);
				match(ELSE);
				setState(459);
				match(LLAVEIZQ);
				setState(460);
				((IfstmtContext)_localctx).belse = block();
				setState(461);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).bif.blk, ((IfstmtContext)_localctx).belse.blk,nil) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(464);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(465);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(466);
				match(LLAVEIZQ);
				setState(467);
				((IfstmtContext)_localctx).bif = block();
				setState(468);
				match(LLAVEDER);
				setState(469);
				((IfstmtContext)_localctx).elifs = elifs(0);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).bif.blk, nil, ((IfstmtContext)_localctx).elifs.elifinst) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(472);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(473);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(474);
				match(LLAVEIZQ);
				setState(475);
				((IfstmtContext)_localctx).bif = block();
				setState(476);
				match(LLAVEDER);
				setState(477);
				((IfstmtContext)_localctx).elifs = elifs(0);
				setState(478);
				match(ELSE);
				setState(479);
				match(LLAVEIZQ);
				setState(480);
				((IfstmtContext)_localctx).belse = block();
				setState(481);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).bif.blk, ((IfstmtContext)_localctx).belse.blk, ((IfstmtContext)_localctx).elifs.elifinst) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElifsContext extends ParserRuleContext {
		public []interface{} elifinst;
		public ElifsContext celif;
		public Token ELSE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElifsContext elifs() {
			return getRuleContext(ElifsContext.class,0);
		}
		public ElifsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elifs; }
	}

	public final ElifsContext elifs() throws RecognitionException {
		return elifs(0);
	}

	private ElifsContext elifs(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ElifsContext _localctx = new ElifsContext(_ctx, _parentState);
		ElifsContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_elifs, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(487);
			((ElifsContext)_localctx).ELSE = match(ELSE);
			setState(488);
			match(IF);
			setState(489);
			((ElifsContext)_localctx).expr = expr(0);
			setState(490);
			match(LLAVEIZQ);
			setState(491);
			((ElifsContext)_localctx).block = block();
			setState(492);
			match(LLAVEDER);

			    _localctx.elifinst = []interface{}{}
			    _localctx.elifinst = append(_localctx.elifinst, instructions.NewElif((((ElifsContext)_localctx).ELSE!=null?((ElifsContext)_localctx).ELSE.getLine():0), (((ElifsContext)_localctx).ELSE!=null?((ElifsContext)_localctx).ELSE.getCharPositionInLine():0), ((ElifsContext)_localctx).expr.e, ((ElifsContext)_localctx).block.blk))
			    

			}
			_ctx.stop = _input.LT(-1);
			setState(506);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ElifsContext(_parentctx, _parentState);
					_localctx.celif = _prevctx;
					_localctx.celif = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_elifs);
					setState(495);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(496);
					((ElifsContext)_localctx).ELSE = match(ELSE);
					setState(497);
					match(IF);
					setState(498);
					((ElifsContext)_localctx).expr = expr(0);
					setState(499);
					match(LLAVEIZQ);
					setState(500);
					((ElifsContext)_localctx).block = block();
					setState(501);
					match(LLAVEDER);

					              var arrif []interface{}
					              arrif = append(((ElifsContext)_localctx).celif.elifinst, instructions.NewElif((((ElifsContext)_localctx).ELSE!=null?((ElifsContext)_localctx).ELSE.getLine():0), (((ElifsContext)_localctx).ELSE!=null?((ElifsContext)_localctx).ELSE.getCharPositionInLine():0), ((ElifsContext)_localctx).expr.e, ((ElifsContext)_localctx).block.blk))
					              _localctx.elifinst = arrif
					          
					}
					} 
				}
				setState(508);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class SwitchstmtContext extends ParserRuleContext {
		public interfaces.Instruction swinst;
		public Token SWITCH;
		public ExprContext expr;
		public CasesContext cases;
		public BlockContext block;
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_switchstmt);
		try {
			setState(526);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(509);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(510);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(511);
				match(LLAVEIZQ);
				setState(512);
				((SwitchstmtContext)_localctx).cases = cases(0);
				setState(513);
				match(DEFAULT);
				setState(514);
				match(DOSP);
				setState(515);
				((SwitchstmtContext)_localctx).block = block();
				setState(516);
				match(LLAVEDER);
				 _localctx.swinst = instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, ((SwitchstmtContext)_localctx).cases.casesinst, ((SwitchstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(519);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(520);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(521);
				match(LLAVEIZQ);
				setState(522);
				((SwitchstmtContext)_localctx).cases = cases(0);
				setState(523);
				match(LLAVEDER);
				 _localctx.swinst = instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, ((SwitchstmtContext)_localctx).cases.casesinst, nil) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CasesContext extends ParserRuleContext {
		public []interface{} casesinst;
		public CasesContext ccases;
		public Token CASE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public CasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cases; }
	}

	public final CasesContext cases() throws RecognitionException {
		return cases(0);
	}

	private CasesContext cases(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CasesContext _localctx = new CasesContext(_ctx, _parentState);
		CasesContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_cases, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(529);
			((CasesContext)_localctx).CASE = match(CASE);
			setState(530);
			((CasesContext)_localctx).expr = expr(0);
			setState(531);
			match(DOSP);
			setState(532);
			((CasesContext)_localctx).block = block();

			    _localctx.casesinst = []interface{}{}
			    _localctx.casesinst = append(_localctx.casesinst, instructions.NewCase((((CasesContext)_localctx).CASE!=null?((CasesContext)_localctx).CASE.getLine():0), (((CasesContext)_localctx).CASE!=null?((CasesContext)_localctx).CASE.getCharPositionInLine():0), ((CasesContext)_localctx).expr.e, ((CasesContext)_localctx).block.blk))
			    

			}
			_ctx.stop = _input.LT(-1);
			setState(544);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new CasesContext(_parentctx, _parentState);
					_localctx.ccases = _prevctx;
					_localctx.ccases = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_cases);
					setState(535);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(536);
					((CasesContext)_localctx).CASE = match(CASE);
					setState(537);
					((CasesContext)_localctx).expr = expr(0);
					setState(538);
					match(DOSP);
					setState(539);
					((CasesContext)_localctx).block = block();

					              var arrcase []interface{}
					              arrcase = append(((CasesContext)_localctx).ccases.casesinst, instructions.NewCase((((CasesContext)_localctx).CASE!=null?((CasesContext)_localctx).CASE.getLine():0), (((CasesContext)_localctx).CASE!=null?((CasesContext)_localctx).CASE.getCharPositionInLine():0), ((CasesContext)_localctx).expr.e, ((CasesContext)_localctx).block.blk))
					              _localctx.casesinst = arrcase
					          
					}
					} 
				}
				setState(546);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class WhilestmtContext extends ParserRuleContext {
		public interfaces.Instruction whileinst;
		public Token WHILE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(547);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(548);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(549);
			match(LLAVEIZQ);
			setState(550);
			((WhilestmtContext)_localctx).block = block();
			setState(551);
			match(LLAVEDER);
			 _localctx.whileinst = instructions.NewWhile((((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getLine():0), (((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getCharPositionInLine():0), ((WhilestmtContext)_localctx).expr.e, ((WhilestmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction forinst;
		public Token FOR;
		public Token ID;
		public ExprContext e1;
		public ExprContext e2;
		public BlockContext block;
		public ExprContext expr;
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public TerminalNode TRESP() { return getToken(SwiftGrammarParser.TRESP, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_forstmt);
		try {
			setState(574);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(554);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(555);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(556);
				match(IN);
				setState(557);
				((ForstmtContext)_localctx).e1 = expr(0);
				setState(558);
				match(TRESP);
				setState(559);
				((ForstmtContext)_localctx).e2 = expr(0);
				setState(560);
				match(LLAVEIZQ);
				setState(561);
				((ForstmtContext)_localctx).block = block();
				setState(562);
				match(LLAVEDER);
				 _localctx.forinst = instructions.NewForRange((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0),(((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null) ,((ForstmtContext)_localctx).e1.e,((ForstmtContext)_localctx).e2.e ,((ForstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(565);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(566);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(567);
				match(IN);
				setState(568);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(569);
				match(LLAVEIZQ);
				setState(570);
				((ForstmtContext)_localctx).block = block();
				setState(571);
				match(LLAVEDER);
				 _localctx.forinst = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0),(((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null) ,((ForstmtContext)_localctx).expr.e ,((ForstmtContext)_localctx).block.blk) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext expr;
		public Token ID;
		public PrimitivesContext primitives;
		public Token INT;
		public Token FLOAT;
		public Token TSTRING;
		public ListArrayContext listArray;
		public Token CORIZQ;
		public ListParamsContext listParams;
		public ListAcesoContext listAceso;
		public ListStructExpContext listStructExp;
		public FuncallestmtContext funcallestmt;
		public ExprContext right;
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public PrimitivesContext primitives() {
			return getRuleContext(PrimitivesContext.class,0);
		}
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode TSTRING() { return getToken(SwiftGrammarParser.TSTRING, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
		public ListAcesoContext listAceso() {
			return getRuleContext(ListAcesoContext.class,0);
		}
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public FuncallestmtContext funcallestmt() {
			return getRuleContext(FuncallestmtContext.class,0);
		}
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SwiftGrammarParser.MOD, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(641);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				setState(577);
				((ExprContext)_localctx).op = match(NOT);
				setState(578);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(15);
				 _localctx.e = expressions.NewOperationUnary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null)) 
				}
				break;
			case 2:
				{
				setState(581);
				((ExprContext)_localctx).op = match(SUB);
				setState(582);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(14);
				 _localctx.e = expressions.NewOperationUnary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null)) 
				}
				break;
			case 3:
				{
				setState(585);
				match(PARIZQ);
				setState(586);
				((ExprContext)_localctx).expr = expr(0);
				setState(587);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(590);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewCallVariable((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null))
				}
				break;
			case 5:
				{
				setState(592);
				((ExprContext)_localctx).primitives = primitives();
				 _localctx.e = ((ExprContext)_localctx).primitives.p
				}
				break;
			case 6:
				{
				setState(595);
				((ExprContext)_localctx).INT = match(INT);
				setState(596);
				match(PARIZQ);
				setState(597);
				((ExprContext)_localctx).expr = expr(0);
				setState(598);
				match(PARDER);
				 _localctx.e = expressions.NewCastingInt((((ExprContext)_localctx).INT!=null?((ExprContext)_localctx).INT.getLine():0), (((ExprContext)_localctx).INT!=null?((ExprContext)_localctx).INT.getCharPositionInLine():0), ((ExprContext)_localctx).expr.e ) 
				}
				break;
			case 7:
				{
				setState(601);
				((ExprContext)_localctx).FLOAT = match(FLOAT);
				setState(602);
				match(PARIZQ);
				setState(603);
				((ExprContext)_localctx).expr = expr(0);
				setState(604);
				match(PARDER);
				 _localctx.e = expressions.NewCastingFloat((((ExprContext)_localctx).FLOAT!=null?((ExprContext)_localctx).FLOAT.getLine():0), (((ExprContext)_localctx).FLOAT!=null?((ExprContext)_localctx).FLOAT.getCharPositionInLine():0), ((ExprContext)_localctx).expr.e ) 
				}
				break;
			case 8:
				{
				setState(607);
				((ExprContext)_localctx).TSTRING = match(TSTRING);
				setState(608);
				match(PARIZQ);
				setState(609);
				((ExprContext)_localctx).expr = expr(0);
				setState(610);
				match(PARDER);
				 _localctx.e = expressions.NewCastingString((((ExprContext)_localctx).TSTRING!=null?((ExprContext)_localctx).TSTRING.getLine():0), (((ExprContext)_localctx).TSTRING!=null?((ExprContext)_localctx).TSTRING.getCharPositionInLine():0), ((ExprContext)_localctx).expr.e ) 
				}
				break;
			case 9:
				{
				setState(613);
				((ExprContext)_localctx).listArray = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).listArray.p
				}
				break;
			case 10:
				{
				setState(616);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(617);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(618);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 11:
				{
				setState(621);
				((ExprContext)_localctx).ID = match(ID);
				setState(622);
				match(PUNTO);
				setState(623);
				match(COUNT);
				 _localctx.e = expressions.NewCount((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 12:
				{
				setState(625);
				((ExprContext)_localctx).ID = match(ID);
				setState(626);
				match(PUNTO);
				setState(627);
				match(ISEMPTY);
				 _localctx.e = expressions.NewIsEmpty((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 13:
				{
				setState(629);
				((ExprContext)_localctx).listAceso = listAceso(0);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listAceso.l) 
				}
				break;
			case 14:
				{
				setState(632);
				((ExprContext)_localctx).ID = match(ID);
				setState(633);
				match(PARIZQ);
				setState(634);
				((ExprContext)_localctx).listStructExp = listStructExp(0);
				setState(635);
				match(PARDER);
				 _localctx.e = expressions.NewStructExp((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null), ((ExprContext)_localctx).listStructExp.l ) 
				}
				break;
			case 15:
				{
				setState(638);
				((ExprContext)_localctx).funcallestmt = funcallestmt();
				 _localctx.e = ((ExprContext)_localctx).funcallestmt.cefun 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(680);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(678);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(643);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(644);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MOD) | (1L << MUL) | (1L << DIV))) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(645);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(648);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(649);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(650);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(653);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(654);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(655);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(658);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(659);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(660);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(663);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(664);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(665);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(668);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(669);
						((ExprContext)_localctx).op = match(AND);
						setState(670);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(673);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(674);
						((ExprContext)_localctx).op = match(OR);
						setState(675);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperationBinary((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(682);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimitivesContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Token NUMBER;
		public Token CHARACTER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public Token NIL;
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public PrimitivesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitives; }
	}

	public final PrimitivesContext primitives() throws RecognitionException {
		PrimitivesContext _localctx = new PrimitivesContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_primitives);
		try {
			setState(695);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(683);
				((PrimitivesContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getLine():0),(((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getLine():0),(((PrimitivesContext)_localctx).NUMBER!=null?((PrimitivesContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 2);
				{
				setState(685);
				((PrimitivesContext)_localctx).CHARACTER = match(CHARACTER);

				        str := (((PrimitivesContext)_localctx).CHARACTER!=null?((PrimitivesContext)_localctx).CHARACTER.getText():null)
				        _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).CHARACTER!=null?((PrimitivesContext)_localctx).CHARACTER.getLine():0), (((PrimitivesContext)_localctx).CHARACTER!=null?((PrimitivesContext)_localctx).CHARACTER.getCharPositionInLine():0), str[1:len(str)-1],environment.CHAR)
				    
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(687);
				((PrimitivesContext)_localctx).STRING = match(STRING);

				        str := (((PrimitivesContext)_localctx).STRING!=null?((PrimitivesContext)_localctx).STRING.getText():null)
				        _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).STRING!=null?((PrimitivesContext)_localctx).STRING.getLine():0), (((PrimitivesContext)_localctx).STRING!=null?((PrimitivesContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				enterOuterAlt(_localctx, 4);
				{
				setState(689);
				((PrimitivesContext)_localctx).TRU = match(TRU);
				 _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).TRU!=null?((PrimitivesContext)_localctx).TRU.getLine():0), (((PrimitivesContext)_localctx).TRU!=null?((PrimitivesContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(691);
				((PrimitivesContext)_localctx).FAL = match(FAL);
				 _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).FAL!=null?((PrimitivesContext)_localctx).FAL.getLine():0), (((PrimitivesContext)_localctx).FAL!=null?((PrimitivesContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case NIL:
				enterOuterAlt(_localctx, 6);
				{
				setState(693);
				((PrimitivesContext)_localctx).NIL = match(NIL);
				 _localctx.p = expressions.NewPrimitive((((PrimitivesContext)_localctx).NIL!=null?((PrimitivesContext)_localctx).NIL.getLine():0), (((PrimitivesContext)_localctx).NIL!=null?((PrimitivesContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListParamsContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(698);
			((ListParamsContext)_localctx).expr = expr(0);

			    _localctx.l = []interface{}{}
			    _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)

			}
			_ctx.stop = _input.LT(-1);
			setState(708);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(701);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(702);
					match(COMA);
					setState(703);
					((ListParamsContext)_localctx).expr = expr(0);

					              var arr []interface{}
					              arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					              _localctx.l = arr
					          
					}
					} 
				}
				setState(710);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListArrayContext extends ParserRuleContext {
		public interfaces.Expression p;
		public ListArrayContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public ListArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listArray; }
	}

	public final ListArrayContext listArray() throws RecognitionException {
		return listArray(0);
	}

	private ListArrayContext listArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListArrayContext _localctx = new ListArrayContext(_ctx, _parentState);
		ListArrayContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(712);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVariable((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(727);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(725);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
					case 1:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(715);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(716);
						match(CORIZQ);
						setState(717);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(718);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					case 2:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(721);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(722);
						match(PUNTO);
						setState(723);
						((ListArrayContext)_localctx).ID = match(ID);
						 _localctx.p = expressions.NewStructAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))  
						}
						break;
					}
					} 
				}
				setState(729);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListAcesoContext extends ParserRuleContext {
		public []interface{} l;
		public ListAcesoContext list;
		public ExprContext expr;
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListAcesoContext listAceso() {
			return getRuleContext(ListAcesoContext.class,0);
		}
		public ListAcesoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAceso; }
	}

	public final ListAcesoContext listAceso() throws RecognitionException {
		return listAceso(0);
	}

	private ListAcesoContext listAceso(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListAcesoContext _localctx = new ListAcesoContext(_ctx, _parentState);
		ListAcesoContext _prevctx = _localctx;
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_listAceso, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(731);
			match(CORIZQ);
			setState(732);
			((ListAcesoContext)_localctx).expr = expr(0);
			setState(733);
			match(CORDER);

			    _localctx.l = []interface{}{}
			    _localctx.l = append(_localctx.l, ((ListAcesoContext)_localctx).expr.e)

			}
			_ctx.stop = _input.LT(-1);
			setState(744);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListAcesoContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listAceso);
					setState(736);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(737);
					match(CORIZQ);
					setState(738);
					((ListAcesoContext)_localctx).expr = expr(0);
					setState(739);
					match(CORDER);

					              var arr []interface{}
					              arr = append(((ListAcesoContext)_localctx).list.l, ((ListAcesoContext)_localctx).expr.e)
					              _localctx.l = arr
					          
					}
					} 
				}
				setState(746);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListStructExpContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructExpContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSP() { return getToken(SwiftGrammarParser.DOSP, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public ListStructExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructExp; }
	}

	public final ListStructExpContext listStructExp() throws RecognitionException {
		return listStructExp(0);
	}

	private ListStructExpContext listStructExp(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructExpContext _localctx = new ListStructExpContext(_ctx, _parentState);
		ListStructExpContext _prevctx = _localctx;
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_listStructExp, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(754);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(748);
				((ListStructExpContext)_localctx).ID = match(ID);
				setState(749);
				match(DOSP);
				setState(750);
				((ListStructExpContext)_localctx).expr = expr(0);

				    var arr []interface{}
				    StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
				    arr = append(arr, StrExp)
				    _localctx.l = arr

				}
				break;
			case 2:
				{

				    _localctx.l = []interface{}{}

				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(765);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructExpContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStructExp);
					setState(756);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(757);
					match(COMA);
					setState(758);
					((ListStructExpContext)_localctx).ID = match(ID);
					setState(759);
					match(DOSP);
					setState(760);
					((ListStructExpContext)_localctx).expr = expr(0);

					              var arr []interface{}
					              StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
					              arr = append(((ListStructExpContext)_localctx).list.l, StrExp)
					              _localctx.l = arr
					          
					}
					} 
				}
				setState(767);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return listStructDec_sempred((ListStructDecContext)_localctx, predIndex);
		case 6:
			return listParamsFunc_sempred((ListParamsFuncContext)_localctx, predIndex);
		case 16:
			return elifs_sempred((ElifsContext)_localctx, predIndex);
		case 18:
			return cases_sempred((CasesContext)_localctx, predIndex);
		case 21:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 23:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 24:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		case 25:
			return listAceso_sempred((ListAcesoContext)_localctx, predIndex);
		case 26:
			return listStructExp_sempred((ListStructExpContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listStructDec_sempred(ListStructDecContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean listParamsFunc_sempred(ListParamsFuncContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean elifs_sempred(ElifsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean cases_sempred(CasesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 22);
		case 5:
			return precpred(_ctx, 21);
		case 6:
			return precpred(_ctx, 20);
		case 7:
			return precpred(_ctx, 19);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 16);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return precpred(_ctx, 3);
		case 13:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean listAceso_sempred(ListAcesoContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listStructExp_sempred(ListStructExpContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3J\u0303\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\3\2\3\2\3\2\3\2\3\3\6\3>\n\3\r\3\16\3?"+
		"\3\3\3\3\3\4\3\4\5\4F\n\4\3\4\3\4\3\4\3\4\5\4L\n\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5"+
		"\4d\n\4\3\4\3\4\3\4\5\4i\n\4\3\4\3\4\3\4\3\4\5\4o\n\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4|\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008d\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\7\6\u0096\n\6\f\6\16\6\u0099\13\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00c5"+
		"\n\7\3\b\3\b\5\b\u00c9\n\b\3\b\3\b\3\b\5\b\u00ce\n\b\3\b\3\b\3\b\3\b\5"+
		"\b\u00d4\n\b\3\b\3\b\3\b\5\b\u00d9\n\b\3\b\3\b\3\b\5\b\u00de\n\b\3\b\3"+
		"\b\3\b\7\b\u00e3\n\b\f\b\16\b\u00e6\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\5\t\u00f2\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00fe"+
		"\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u010c\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u017b\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u018c\n\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u019d\n\16\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u01b7\n\17\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\5\21\u01e7\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\7\22\u01fb\n\22\f\22\16"+
		"\22\u01fe\13\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0211\n\23\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u0221\n\24\f\24\16"+
		"\24\u0224\13\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\5\26\u0241\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\5\27\u0284\n\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u02a9\n\27\f\27"+
		"\16\27\u02ac\13\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\5\30\u02ba\n\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\7\31\u02c5\n\31\f\31\16\31\u02c8\13\31\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\7\32\u02d8\n\32\f\32\16\32\u02db"+
		"\13\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\7\33"+
		"\u02e9\n\33\f\33\16\33\u02ec\13\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\5\34\u02f5\n\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\7\34\u02fe\n\34\f"+
		"\34\16\34\u0301\13\34\3\34\2\13\n\16\"&,\60\62\64\66\35\2\4\6\b\n\f\16"+
		"\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\66\2\b\4\2))FF\3\2\64\66\3"+
		"\2\678\4\2\60\60\62\62\4\2\61\61\63\63\3\2*+\2\u034a\28\3\2\2\2\4=\3\2"+
		"\2\2\6{\3\2\2\2\b}\3\2\2\2\n\u008c\3\2\2\2\f\u00c4\3\2\2\2\16\u00d3\3"+
		"\2\2\2\20\u00f1\3\2\2\2\22\u00fd\3\2\2\2\24\u010b\3\2\2\2\26\u017a\3\2"+
		"\2\2\30\u018b\3\2\2\2\32\u019c\3\2\2\2\34\u01b6\3\2\2\2\36\u01b8\3\2\2"+
		"\2 \u01e6\3\2\2\2\"\u01e8\3\2\2\2$\u0210\3\2\2\2&\u0212\3\2\2\2(\u0225"+
		"\3\2\2\2*\u0240\3\2\2\2,\u0283\3\2\2\2.\u02b9\3\2\2\2\60\u02bb\3\2\2\2"+
		"\62\u02c9\3\2\2\2\64\u02dc\3\2\2\2\66\u02f4\3\2\2\289\5\4\3\29:\7\2\2"+
		"\3:;\b\2\1\2;\3\3\2\2\2<>\5\6\4\2=<\3\2\2\2>?\3\2\2\2?=\3\2\2\2?@\3\2"+
		"\2\2@A\3\2\2\2AB\b\3\1\2B\5\3\2\2\2CE\5\24\13\2DF\7E\2\2ED\3\2\2\2EF\3"+
		"\2\2\2FG\3\2\2\2GH\b\4\1\2H|\3\2\2\2IK\5\26\f\2JL\7E\2\2KJ\3\2\2\2KL\3"+
		"\2\2\2LM\3\2\2\2MN\b\4\1\2N|\3\2\2\2OP\5\36\20\2PQ\b\4\1\2Q|\3\2\2\2R"+
		"S\5 \21\2ST\b\4\1\2T|\3\2\2\2UV\5$\23\2VW\b\4\1\2W|\3\2\2\2XY\5(\25\2"+
		"YZ\b\4\1\2Z|\3\2\2\2[\\\5*\26\2\\]\b\4\1\2]|\3\2\2\2^_\5\34\17\2_`\b\4"+
		"\1\2`|\3\2\2\2ac\7\22\2\2bd\7E\2\2cb\3\2\2\2cd\3\2\2\2de\3\2\2\2e|\b\4"+
		"\1\2fh\7\23\2\2gi\7E\2\2hg\3\2\2\2hi\3\2\2\2ij\3\2\2\2j|\b\4\1\2kl\7\24"+
		"\2\2ln\5,\27\2mo\7E\2\2nm\3\2\2\2no\3\2\2\2op\3\2\2\2pq\b\4\1\2q|\3\2"+
		"\2\2rs\5\20\t\2st\b\4\1\2t|\3\2\2\2uv\5\f\7\2vw\b\4\1\2w|\3\2\2\2xy\5"+
		"\b\5\2yz\b\4\1\2z|\3\2\2\2{C\3\2\2\2{I\3\2\2\2{O\3\2\2\2{R\3\2\2\2{U\3"+
		"\2\2\2{X\3\2\2\2{[\3\2\2\2{^\3\2\2\2{a\3\2\2\2{f\3\2\2\2{k\3\2\2\2{r\3"+
		"\2\2\2{u\3\2\2\2{x\3\2\2\2|\7\3\2\2\2}~\7\37\2\2~\177\7)\2\2\177\u0080"+
		"\7;\2\2\u0080\u0081\5\n\6\2\u0081\u0082\7<\2\2\u0082\u0083\b\5\1\2\u0083"+
		"\t\3\2\2\2\u0084\u0085\b\6\1\2\u0085\u0086\7\b\2\2\u0086\u0087\7)\2\2"+
		"\u0087\u0088\7?\2\2\u0088\u0089\5\30\r\2\u0089\u008a\b\6\1\2\u008a\u008d"+
		"\3\2\2\2\u008b\u008d\b\6\1\2\u008c\u0084\3\2\2\2\u008c\u008b\3\2\2\2\u008d"+
		"\u0097\3\2\2\2\u008e\u008f\f\5\2\2\u008f\u0090\7\b\2\2\u0090\u0091\7)"+
		"\2\2\u0091\u0092\7?\2\2\u0092\u0093\5\30\r\2\u0093\u0094\b\6\1\2\u0094"+
		"\u0096\3\2\2\2\u0095\u008e\3\2\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3\2"+
		"\2\2\u0097\u0098\3\2\2\2\u0098\13\3\2\2\2\u0099\u0097\3\2\2\2\u009a\u009b"+
		"\7\"\2\2\u009b\u009c\7)\2\2\u009c\u009d\79\2\2\u009d\u009e\7:\2\2\u009e"+
		"\u009f\7;\2\2\u009f\u00a0\5\4\3\2\u00a0\u00a1\7<\2\2\u00a1\u00a2\b\7\1"+
		"\2\u00a2\u00c5\3\2\2\2\u00a3\u00a4\7\"\2\2\u00a4\u00a5\7)\2\2\u00a5\u00a6"+
		"\79\2\2\u00a6\u00a7\5\16\b\2\u00a7\u00a8\7:\2\2\u00a8\u00a9\7;\2\2\u00a9"+
		"\u00aa\5\4\3\2\u00aa\u00ab\7<\2\2\u00ab\u00ac\b\7\1\2\u00ac\u00c5\3\2"+
		"\2\2\u00ad\u00ae\7\"\2\2\u00ae\u00af\7)\2\2\u00af\u00b0\79\2\2\u00b0\u00b1"+
		"\7:\2\2\u00b1\u00b2\7C\2\2\u00b2\u00b3\5\30\r\2\u00b3\u00b4\7;\2\2\u00b4"+
		"\u00b5\5\4\3\2\u00b5\u00b6\7<\2\2\u00b6\u00b7\b\7\1\2\u00b7\u00c5\3\2"+
		"\2\2\u00b8\u00b9\7\"\2\2\u00b9\u00ba\7)\2\2\u00ba\u00bb\79\2\2\u00bb\u00bc"+
		"\5\16\b\2\u00bc\u00bd\7:\2\2\u00bd\u00be\7C\2\2\u00be\u00bf\5\30\r\2\u00bf"+
		"\u00c0\7;\2\2\u00c0\u00c1\5\4\3\2\u00c1\u00c2\7<\2\2\u00c2\u00c3\b\7\1"+
		"\2\u00c3\u00c5\3\2\2\2\u00c4\u009a\3\2\2\2\u00c4\u00a3\3\2\2\2\u00c4\u00ad"+
		"\3\2\2\2\u00c4\u00b8\3\2\2\2\u00c5\r\3\2\2\2\u00c6\u00c8\b\b\1\2\u00c7"+
		"\u00c9\t\2\2\2\u00c8\u00c7\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00ca\3\2"+
		"\2\2\u00ca\u00cb\7)\2\2\u00cb\u00cd\7?\2\2\u00cc\u00ce\7$\2\2\u00cd\u00cc"+
		"\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00d0\5\30\r\2"+
		"\u00d0\u00d1\b\b\1\2\u00d1\u00d4\3\2\2\2\u00d2\u00d4\b\b\1\2\u00d3\u00c6"+
		"\3\2\2\2\u00d3\u00d2\3\2\2\2\u00d4\u00e4\3\2\2\2\u00d5\u00d6\f\5\2\2\u00d6"+
		"\u00d8\7A\2\2\u00d7\u00d9\t\2\2\2\u00d8\u00d7\3\2\2\2\u00d8\u00d9\3\2"+
		"\2\2\u00d9\u00da\3\2\2\2\u00da\u00db\7)\2\2\u00db\u00dd\7?\2\2\u00dc\u00de"+
		"\7$\2\2\u00dd\u00dc\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00df\3\2\2\2\u00df"+
		"\u00e0\5\30\r\2\u00e0\u00e1\b\b\1\2\u00e1\u00e3\3\2\2\2\u00e2\u00d5\3"+
		"\2\2\2\u00e3\u00e6\3\2\2\2\u00e4\u00e2\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5"+
		"\17\3\2\2\2\u00e6\u00e4\3\2\2\2\u00e7\u00e8\7)\2\2\u00e8\u00e9\79\2\2"+
		"\u00e9\u00ea\7:\2\2\u00ea\u00f2\b\t\1\2\u00eb\u00ec\7)\2\2\u00ec\u00ed"+
		"\79\2\2\u00ed\u00ee\5\60\31\2\u00ee\u00ef\7:\2\2\u00ef\u00f0\b\t\1\2\u00f0"+
		"\u00f2\3\2\2\2\u00f1\u00e7\3\2\2\2\u00f1\u00eb\3\2\2\2\u00f2\21\3\2\2"+
		"\2\u00f3\u00f4\7)\2\2\u00f4\u00f5\79\2\2\u00f5\u00f6\7:\2\2\u00f6\u00fe"+
		"\b\n\1\2\u00f7\u00f8\7)\2\2\u00f8\u00f9\79\2\2\u00f9\u00fa\5\60\31\2\u00fa"+
		"\u00fb\7:\2\2\u00fb\u00fc\b\n\1\2\u00fc\u00fe\3\2\2\2\u00fd\u00f3\3\2"+
		"\2\2\u00fd\u00f7\3\2\2\2\u00fe\23\3\2\2\2\u00ff\u0100\7\f\2\2\u0100\u0101"+
		"\79\2\2\u0101\u0102\5,\27\2\u0102\u0103\7:\2\2\u0103\u0104\b\13\1\2\u0104"+
		"\u010c\3\2\2\2\u0105\u0106\7\f\2\2\u0106\u0107\79\2\2\u0107\u0108\5\60"+
		"\31\2\u0108\u0109\7:\2\2\u0109\u010a\b\13\1\2\u010a\u010c\3\2\2\2\u010b"+
		"\u00ff\3\2\2\2\u010b\u0105\3\2\2\2\u010c\25\3\2\2\2\u010d\u010e\7\b\2"+
		"\2\u010e\u010f\7)\2\2\u010f\u0110\7?\2\2\u0110\u0111\5\30\r\2\u0111\u0112"+
		"\7/\2\2\u0112\u0113\5,\27\2\u0113\u0114\b\f\1\2\u0114\u017b\3\2\2\2\u0115"+
		"\u0116\7\b\2\2\u0116\u0117\7)\2\2\u0117\u0118\7/\2\2\u0118\u0119\5,\27"+
		"\2\u0119\u011a\b\f\1\2\u011a\u017b\3\2\2\2\u011b\u011c\7\b\2\2\u011c\u011d"+
		"\7)\2\2\u011d\u011e\7?\2\2\u011e\u011f\5\30\r\2\u011f\u0120\7B\2\2\u0120"+
		"\u0121\b\f\1\2\u0121\u017b\3\2\2\2\u0122\u0123\7\t\2\2\u0123\u0124\7)"+
		"\2\2\u0124\u0125\7/\2\2\u0125\u0126\5.\30\2\u0126\u0127\b\f\1\2\u0127"+
		"\u017b\3\2\2\2\u0128\u0129\7\t\2\2\u0129\u012a\7)\2\2\u012a\u012b\7?\2"+
		"\2\u012b\u012c\5\30\r\2\u012c\u012d\7/\2\2\u012d\u012e\5.\30\2\u012e\u012f"+
		"\b\f\1\2\u012f\u017b\3\2\2\2\u0130\u0131\7\t\2\2\u0131\u0132\7)\2\2\u0132"+
		"\u0133\7/\2\2\u0133\u0134\5,\27\2\u0134\u0135\b\f\1\2\u0135\u017b\3\2"+
		"\2\2\u0136\u0137\7\t\2\2\u0137\u0138\7)\2\2\u0138\u0139\7?\2\2\u0139\u013a"+
		"\5\30\r\2\u013a\u013b\7/\2\2\u013b\u013c\5,\27\2\u013c\u013d\b\f\1\2\u013d"+
		"\u017b\3\2\2\2\u013e\u013f\7\t\2\2\u013f\u0140\7)\2\2\u0140\u0141\7/\2"+
		"\2\u0141\u0142\5,\27\2\u0142\u0143\b\f\1\2\u0143\u017b\3\2\2\2\u0144\u0145"+
		"\7)\2\2\u0145\u0146\7/\2\2\u0146\u0147\5,\27\2\u0147\u0148\b\f\1\2\u0148"+
		"\u017b\3\2\2\2\u0149\u014a\7)\2\2\u014a\u014b\7=\2\2\u014b\u014c\5,\27"+
		"\2\u014c\u014d\7>\2\2\u014d\u014e\7/\2\2\u014e\u014f\5,\27\2\u014f\u0150"+
		"\b\f\1\2\u0150\u017b\3\2\2\2\u0151\u0152\7)\2\2\u0152\u0153\5,\27\2\u0153"+
		"\u0154\7/\2\2\u0154\u0155\5,\27\2\u0155\u0156\b\f\1\2\u0156\u017b\3\2"+
		"\2\2\u0157\u0158\7)\2\2\u0158\u0159\7\67\2\2\u0159\u015a\7/\2\2\u015a"+
		"\u015b\5,\27\2\u015b\u015c\b\f\1\2\u015c\u017b\3\2\2\2\u015d\u015e\7)"+
		"\2\2\u015e\u015f\78\2\2\u015f\u0160\7/\2\2\u0160\u0161\5,\27\2\u0161\u0162"+
		"\b\f\1\2\u0162\u017b\3\2\2\2\u0163\u0164\7\b\2\2\u0164\u0165\7)\2\2\u0165"+
		"\u0166\7?\2\2\u0166\u0167\5\30\r\2\u0167\u0168\7/\2\2\u0168\u0169\5,\27"+
		"\2\u0169\u016a\b\f\1\2\u016a\u017b\3\2\2\2\u016b\u016c\7\b\2\2\u016c\u016d"+
		"\7)\2\2\u016d\u016e\7?\2\2\u016e\u016f\5\30\r\2\u016f\u0170\7/\2\2\u0170"+
		"\u0171\7=\2\2\u0171\u0172\7>\2\2\u0172\u0173\b\f\1\2\u0173\u017b\3\2\2"+
		"\2\u0174\u0175\7\b\2\2\u0175\u0176\7)\2\2\u0176\u0177\7/\2\2\u0177\u0178"+
		"\5,\27\2\u0178\u0179\b\f\1\2\u0179\u017b\3\2\2\2\u017a\u010d\3\2\2\2\u017a"+
		"\u0115\3\2\2\2\u017a\u011b\3\2\2\2\u017a\u0122\3\2\2\2\u017a\u0128\3\2"+
		"\2\2\u017a\u0130\3\2\2\2\u017a\u0136\3\2\2\2\u017a\u013e\3\2\2\2\u017a"+
		"\u0144\3\2\2\2\u017a\u0149\3\2\2\2\u017a\u0151\3\2\2\2\u017a\u0157\3\2"+
		"\2\2\u017a\u015d\3\2\2\2\u017a\u0163\3\2\2\2\u017a\u016b\3\2\2\2\u017a"+
		"\u0174\3\2\2\2\u017b\27\3\2\2\2\u017c\u017d\7\3\2\2\u017d\u018c\b\r\1"+
		"\2\u017e\u017f\7\4\2\2\u017f\u018c\b\r\1\2\u0180\u0181\7\5\2\2\u0181\u018c"+
		"\b\r\1\2\u0182\u0183\7\6\2\2\u0183\u018c\b\r\1\2\u0184\u0185\7\7\2\2\u0185"+
		"\u018c\b\r\1\2\u0186\u0187\7=\2\2\u0187\u0188\5\32\16\2\u0188\u0189\7"+
		">\2\2\u0189\u018a\b\r\1\2\u018a\u018c\3\2\2\2\u018b\u017c\3\2\2\2\u018b"+
		"\u017e\3\2\2\2\u018b\u0180\3\2\2\2\u018b\u0182\3\2\2\2\u018b\u0184\3\2"+
		"\2\2\u018b\u0186\3\2\2\2\u018c\31\3\2\2\2\u018d\u018e\7\3\2\2\u018e\u019d"+
		"\b\16\1\2\u018f\u0190\7\4\2\2\u0190\u019d\b\16\1\2\u0191\u0192\7\5\2\2"+
		"\u0192\u019d\b\16\1\2\u0193\u0194\7\6\2\2\u0194\u019d\b\16\1\2\u0195\u0196"+
		"\7\7\2\2\u0196\u019d\b\16\1\2\u0197\u0198\7=\2\2\u0198\u0199\5\32\16\2"+
		"\u0199\u019a\7>\2\2\u019a\u019b\b\16\1\2\u019b\u019d\3\2\2\2\u019c\u018d"+
		"\3\2\2\2\u019c\u018f\3\2\2\2\u019c\u0191\3\2\2\2\u019c\u0193\3\2\2\2\u019c"+
		"\u0195\3\2\2\2\u019c\u0197\3\2\2\2\u019d\33\3\2\2\2\u019e\u019f\7)\2\2"+
		"\u019f\u01a0\7@\2\2\u01a0\u01a1\7\31\2\2\u01a1\u01a2\79\2\2\u01a2\u01a3"+
		"\5,\27\2\u01a3\u01a4\7:\2\2\u01a4\u01a5\b\17\1\2\u01a5\u01b7\3\2\2\2\u01a6"+
		"\u01a7\7)\2\2\u01a7\u01a8\7@\2\2\u01a8\u01a9\7\32\2\2\u01a9\u01aa\79\2"+
		"\2\u01aa\u01ab\7:\2\2\u01ab\u01b7\b\17\1\2\u01ac\u01ad\7)\2\2\u01ad\u01ae"+
		"\7@\2\2\u01ae\u01af\7\33\2\2\u01af\u01b0\79\2\2\u01b0\u01b1\7#\2\2\u01b1"+
		"\u01b2\7?\2\2\u01b2\u01b3\5,\27\2\u01b3\u01b4\7:\2\2\u01b4\u01b5\b\17"+
		"\1\2\u01b5\u01b7\3\2\2\2\u01b6\u019e\3\2\2\2\u01b6\u01a6\3\2\2\2\u01b6"+
		"\u01ac\3\2\2\2\u01b7\35\3\2\2\2\u01b8\u01b9\7\30\2\2\u01b9\u01ba\5,\27"+
		"\2\u01ba\u01bb\7\16\2\2\u01bb\u01bc\7;\2\2\u01bc\u01bd\5\4\3\2\u01bd\u01be"+
		"\7<\2\2\u01be\u01bf\b\20\1\2\u01bf\37\3\2\2\2\u01c0\u01c1\7\r\2\2\u01c1"+
		"\u01c2\5,\27\2\u01c2\u01c3\7;\2\2\u01c3\u01c4\5\4\3\2\u01c4\u01c5\7<\2"+
		"\2\u01c5\u01c6\b\21\1\2\u01c6\u01e7\3\2\2\2\u01c7\u01c8\7\r\2\2\u01c8"+
		"\u01c9\5,\27\2\u01c9\u01ca\7;\2\2\u01ca\u01cb\5\4\3\2\u01cb\u01cc\7<\2"+
		"\2\u01cc\u01cd\7\16\2\2\u01cd\u01ce\7;\2\2\u01ce\u01cf\5\4\3\2\u01cf\u01d0"+
		"\7<\2\2\u01d0\u01d1\b\21\1\2\u01d1\u01e7\3\2\2\2\u01d2\u01d3\7\r\2\2\u01d3"+
		"\u01d4\5,\27\2\u01d4\u01d5\7;\2\2\u01d5\u01d6\5\4\3\2\u01d6\u01d7\7<\2"+
		"\2\u01d7\u01d8\5\"\22\2\u01d8\u01d9\b\21\1\2\u01d9\u01e7\3\2\2\2\u01da"+
		"\u01db\7\r\2\2\u01db\u01dc\5,\27\2\u01dc\u01dd\7;\2\2\u01dd\u01de\5\4"+
		"\3\2\u01de\u01df\7<\2\2\u01df\u01e0\5\"\22\2\u01e0\u01e1\7\16\2\2\u01e1"+
		"\u01e2\7;\2\2\u01e2\u01e3\5\4\3\2\u01e3\u01e4\7<\2\2\u01e4\u01e5\b\21"+
		"\1\2\u01e5\u01e7\3\2\2\2\u01e6\u01c0\3\2\2\2\u01e6\u01c7\3\2\2\2\u01e6"+
		"\u01d2\3\2\2\2\u01e6\u01da\3\2\2\2\u01e7!\3\2\2\2\u01e8\u01e9\b\22\1\2"+
		"\u01e9\u01ea\7\16\2\2\u01ea\u01eb\7\r\2\2\u01eb\u01ec\5,\27\2\u01ec\u01ed"+
		"\7;\2\2\u01ed\u01ee\5\4\3\2\u01ee\u01ef\7<\2\2\u01ef\u01f0\b\22\1\2\u01f0"+
		"\u01fc\3\2\2\2\u01f1\u01f2\f\4\2\2\u01f2\u01f3\7\16\2\2\u01f3\u01f4\7"+
		"\r\2\2\u01f4\u01f5\5,\27\2\u01f5\u01f6\7;\2\2\u01f6\u01f7\5\4\3\2\u01f7"+
		"\u01f8\7<\2\2\u01f8\u01f9\b\22\1\2\u01f9\u01fb\3\2\2\2\u01fa\u01f1\3\2"+
		"\2\2\u01fb\u01fe\3\2\2\2\u01fc\u01fa\3\2\2\2\u01fc\u01fd\3\2\2\2\u01fd"+
		"#\3\2\2\2\u01fe\u01fc\3\2\2\2\u01ff\u0200\7\17\2\2\u0200\u0201\5,\27\2"+
		"\u0201\u0202\7;\2\2\u0202\u0203\5&\24\2\u0203\u0204\7\21\2\2\u0204\u0205"+
		"\7?\2\2\u0205\u0206\5\4\3\2\u0206\u0207\7<\2\2\u0207\u0208\b\23\1\2\u0208"+
		"\u0211\3\2\2\2\u0209\u020a\7\17\2\2\u020a\u020b\5,\27\2\u020b\u020c\7"+
		";\2\2\u020c\u020d\5&\24\2\u020d\u020e\7<\2\2\u020e\u020f\b\23\1\2\u020f"+
		"\u0211\3\2\2\2\u0210\u01ff\3\2\2\2\u0210\u0209\3\2\2\2\u0211%\3\2\2\2"+
		"\u0212\u0213\b\24\1\2\u0213\u0214\7\20\2\2\u0214\u0215\5,\27\2\u0215\u0216"+
		"\7?\2\2\u0216\u0217\5\4\3\2\u0217\u0218\b\24\1\2\u0218\u0222\3\2\2\2\u0219"+
		"\u021a\f\4\2\2\u021a\u021b\7\20\2\2\u021b\u021c\5,\27\2\u021c\u021d\7"+
		"?\2\2\u021d\u021e\5\4\3\2\u021e\u021f\b\24\1\2\u021f\u0221\3\2\2\2\u0220"+
		"\u0219\3\2\2\2\u0221\u0224\3\2\2\2\u0222\u0220\3\2\2\2\u0222\u0223\3\2"+
		"\2\2\u0223\'\3\2\2\2\u0224\u0222\3\2\2\2\u0225\u0226\7\25\2\2\u0226\u0227"+
		"\5,\27\2\u0227\u0228\7;\2\2\u0228\u0229\5\4\3\2\u0229\u022a\7<\2\2\u022a"+
		"\u022b\b\25\1\2\u022b)\3\2\2\2\u022c\u022d\7\26\2\2\u022d\u022e\7)\2\2"+
		"\u022e\u022f\7\27\2\2\u022f\u0230\5,\27\2\u0230\u0231\7G\2\2\u0231\u0232"+
		"\5,\27\2\u0232\u0233\7;\2\2\u0233\u0234\5\4\3\2\u0234\u0235\7<\2\2\u0235"+
		"\u0236\b\26\1\2\u0236\u0241\3\2\2\2\u0237\u0238\7\26\2\2\u0238\u0239\7"+
		")\2\2\u0239\u023a\7\27\2\2\u023a\u023b\5,\27\2\u023b\u023c\7;\2\2\u023c"+
		"\u023d\5\4\3\2\u023d\u023e\7<\2\2\u023e\u023f\b\26\1\2\u023f\u0241\3\2"+
		"\2\2\u0240\u022c\3\2\2\2\u0240\u0237\3\2\2\2\u0241+\3\2\2\2\u0242\u0243"+
		"\b\27\1\2\u0243\u0244\7,\2\2\u0244\u0245\5,\27\21\u0245\u0246\b\27\1\2"+
		"\u0246\u0284\3\2\2\2\u0247\u0248\78\2\2\u0248\u0249\5,\27\20\u0249\u024a"+
		"\b\27\1\2\u024a\u0284\3\2\2\2\u024b\u024c\79\2\2\u024c\u024d\5,\27\2\u024d"+
		"\u024e\7:\2\2\u024e\u024f\b\27\1\2\u024f\u0284\3\2\2\2\u0250\u0251\7)"+
		"\2\2\u0251\u0284\b\27\1\2\u0252\u0253\5.\30\2\u0253\u0254\b\27\1\2\u0254"+
		"\u0284\3\2\2\2\u0255\u0256\7\3\2\2\u0256\u0257\79\2\2\u0257\u0258\5,\27"+
		"\2\u0258\u0259\7:\2\2\u0259\u025a\b\27\1\2\u025a\u0284\3\2\2\2\u025b\u025c"+
		"\7\4\2\2\u025c\u025d\79\2\2\u025d\u025e\5,\27\2\u025e\u025f\7:\2\2\u025f"+
		"\u0260\b\27\1\2\u0260\u0284\3\2\2\2\u0261\u0262\7\5\2\2\u0262\u0263\7"+
		"9\2\2\u0263\u0264\5,\27\2\u0264\u0265\7:\2\2\u0265\u0266\b\27\1\2\u0266"+
		"\u0284\3\2\2\2\u0267\u0268\5\62\32\2\u0268\u0269\b\27\1\2\u0269\u0284"+
		"\3\2\2\2\u026a\u026b\7=\2\2\u026b\u026c\5\60\31\2\u026c\u026d\7>\2\2\u026d"+
		"\u026e\b\27\1\2\u026e\u0284\3\2\2\2\u026f\u0270\7)\2\2\u0270\u0271\7@"+
		"\2\2\u0271\u0272\7\35\2\2\u0272\u0284\b\27\1\2\u0273\u0274\7)\2\2\u0274"+
		"\u0275\7@\2\2\u0275\u0276\7\34\2\2\u0276\u0284\b\27\1\2\u0277\u0278\5"+
		"\64\33\2\u0278\u0279\b\27\1\2\u0279\u0284\3\2\2\2\u027a\u027b\7)\2\2\u027b"+
		"\u027c\79\2\2\u027c\u027d\5\66\34\2\u027d\u027e\7:\2\2\u027e\u027f\b\27"+
		"\1\2\u027f\u0284\3\2\2\2\u0280\u0281\5\22\n\2\u0281\u0282\b\27\1\2\u0282"+
		"\u0284\3\2\2\2\u0283\u0242\3\2\2\2\u0283\u0247\3\2\2\2\u0283\u024b\3\2"+
		"\2\2\u0283\u0250\3\2\2\2\u0283\u0252\3\2\2\2\u0283\u0255\3\2\2\2\u0283"+
		"\u025b\3\2\2\2\u0283\u0261\3\2\2\2\u0283\u0267\3\2\2\2\u0283\u026a\3\2"+
		"\2\2\u0283\u026f\3\2\2\2\u0283\u0273\3\2\2\2\u0283\u0277\3\2\2\2\u0283"+
		"\u027a\3\2\2\2\u0283\u0280\3\2\2\2\u0284\u02aa\3\2\2\2\u0285\u0286\f\30"+
		"\2\2\u0286\u0287\t\3\2\2\u0287\u0288\5,\27\31\u0288\u0289\b\27\1\2\u0289"+
		"\u02a9\3\2\2\2\u028a\u028b\f\27\2\2\u028b\u028c\t\4\2\2\u028c\u028d\5"+
		",\27\30\u028d\u028e\b\27\1\2\u028e\u02a9\3\2\2\2\u028f\u0290\f\26\2\2"+
		"\u0290\u0291\t\5\2\2\u0291\u0292\5,\27\27\u0292\u0293\b\27\1\2\u0293\u02a9"+
		"\3\2\2\2\u0294\u0295\f\25\2\2\u0295\u0296\t\6\2\2\u0296\u0297\5,\27\26"+
		"\u0297\u0298\b\27\1\2\u0298\u02a9\3\2\2\2\u0299\u029a\f\24\2\2\u029a\u029b"+
		"\t\7\2\2\u029b\u029c\5,\27\25\u029c\u029d\b\27\1\2\u029d\u02a9\3\2\2\2"+
		"\u029e\u029f\f\23\2\2\u029f\u02a0\7.\2\2\u02a0\u02a1\5,\27\24\u02a1\u02a2"+
		"\b\27\1\2\u02a2\u02a9\3\2\2\2\u02a3\u02a4\f\22\2\2\u02a4\u02a5\7-\2\2"+
		"\u02a5\u02a6\5,\27\23\u02a6\u02a7\b\27\1\2\u02a7\u02a9\3\2\2\2\u02a8\u0285"+
		"\3\2\2\2\u02a8\u028a\3\2\2\2\u02a8\u028f\3\2\2\2\u02a8\u0294\3\2\2\2\u02a8"+
		"\u0299\3\2\2\2\u02a8\u029e\3\2\2\2\u02a8\u02a3\3\2\2\2\u02a9\u02ac\3\2"+
		"\2\2\u02aa\u02a8\3\2\2\2\u02aa\u02ab\3\2\2\2\u02ab-\3\2\2\2\u02ac\u02aa"+
		"\3\2\2\2\u02ad\u02ae\7&\2\2\u02ae\u02ba\b\30\1\2\u02af\u02b0\7\'\2\2\u02b0"+
		"\u02ba\b\30\1\2\u02b1\u02b2\7(\2\2\u02b2\u02ba\b\30\1\2\u02b3\u02b4\7"+
		"\n\2\2\u02b4\u02ba\b\30\1\2\u02b5\u02b6\7\13\2\2\u02b6\u02ba\b\30\1\2"+
		"\u02b7\u02b8\7%\2\2\u02b8\u02ba\b\30\1\2\u02b9\u02ad\3\2\2\2\u02b9\u02af"+
		"\3\2\2\2\u02b9\u02b1\3\2\2\2\u02b9\u02b3\3\2\2\2\u02b9\u02b5\3\2\2\2\u02b9"+
		"\u02b7\3\2\2\2\u02ba/\3\2\2\2\u02bb\u02bc\b\31\1\2\u02bc\u02bd\5,\27\2"+
		"\u02bd\u02be\b\31\1\2\u02be\u02c6\3\2\2\2\u02bf\u02c0\f\4\2\2\u02c0\u02c1"+
		"\7A\2\2\u02c1\u02c2\5,\27\2\u02c2\u02c3\b\31\1\2\u02c3\u02c5\3\2\2\2\u02c4"+
		"\u02bf\3\2\2\2\u02c5\u02c8\3\2\2\2\u02c6\u02c4\3\2\2\2\u02c6\u02c7\3\2"+
		"\2\2\u02c7\61\3\2\2\2\u02c8\u02c6\3\2\2\2\u02c9\u02ca\b\32\1\2\u02ca\u02cb"+
		"\7)\2\2\u02cb\u02cc\b\32\1\2\u02cc\u02d9\3\2\2\2\u02cd\u02ce\f\5\2\2\u02ce"+
		"\u02cf\7=\2\2\u02cf\u02d0\5,\27\2\u02d0\u02d1\7>\2\2\u02d1\u02d2\b\32"+
		"\1\2\u02d2\u02d8\3\2\2\2\u02d3\u02d4\f\3\2\2\u02d4\u02d5\7@\2\2\u02d5"+
		"\u02d6\7)\2\2\u02d6\u02d8\b\32\1\2\u02d7\u02cd\3\2\2\2\u02d7\u02d3\3\2"+
		"\2\2\u02d8\u02db\3\2\2\2\u02d9\u02d7\3\2\2\2\u02d9\u02da\3\2\2\2\u02da"+
		"\63\3\2\2\2\u02db\u02d9\3\2\2\2\u02dc\u02dd\b\33\1\2\u02dd\u02de\7=\2"+
		"\2\u02de\u02df\5,\27\2\u02df\u02e0\7>\2\2\u02e0\u02e1\b\33\1\2\u02e1\u02ea"+
		"\3\2\2\2\u02e2\u02e3\f\4\2\2\u02e3\u02e4\7=\2\2\u02e4\u02e5\5,\27\2\u02e5"+
		"\u02e6\7>\2\2\u02e6\u02e7\b\33\1\2\u02e7\u02e9\3\2\2\2\u02e8\u02e2\3\2"+
		"\2\2\u02e9\u02ec\3\2\2\2\u02ea\u02e8\3\2\2\2\u02ea\u02eb\3\2\2\2\u02eb"+
		"\65\3\2\2\2\u02ec\u02ea\3\2\2\2\u02ed\u02ee\b\34\1\2\u02ee\u02ef\7)\2"+
		"\2\u02ef\u02f0\7?\2\2\u02f0\u02f1\5,\27\2\u02f1\u02f2\b\34\1\2\u02f2\u02f5"+
		"\3\2\2\2\u02f3\u02f5\b\34\1\2\u02f4\u02ed\3\2\2\2\u02f4\u02f3\3\2\2\2"+
		"\u02f5\u02ff\3\2\2\2\u02f6\u02f7\f\5\2\2\u02f7\u02f8\7A\2\2\u02f8\u02f9"+
		"\7)\2\2\u02f9\u02fa\7?\2\2\u02fa\u02fb\5,\27\2\u02fb\u02fc\b\34\1\2\u02fc"+
		"\u02fe\3\2\2\2\u02fd\u02f6\3\2\2\2\u02fe\u0301\3\2\2\2\u02ff\u02fd\3\2"+
		"\2\2\u02ff\u0300\3\2\2\2\u0300\67\3\2\2\2\u0301\u02ff\3\2\2\2(?EKchn{"+
		"\u008c\u0097\u00c4\u00c8\u00cd\u00d3\u00d8\u00dd\u00e4\u00f1\u00fd\u010b"+
		"\u017a\u018b\u019c\u01b6\u01e6\u01fc\u0210\u0222\u0240\u0283\u02a8\u02aa"+
		"\u02b9\u02c6\u02d7\u02d9\u02ea\u02f4\u02ff";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}