// Generated from c:\Users\angge\Documents\GitHub\OLC2_P1_201901055\Backend\SwiftLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "TSTRING", "BOOL", "CHAR", "VAR", "LET", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", 
			"RETURN", "WHILE", "FOR", "IN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", 
			"ISEMPTY", "COUNT", "REPEATING", "STRUCT", "SELF", "MUTATING", "FUNC", 
			"AT", "INOUT", "NIL", "NUMBER", "CHARACTER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", 
			"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORIZQ", "CORDER", "DOSP", "PUNTO", "COMA", "INTCE", "FLECHA", "AMP", 
			"PCOMA", "GBAJO", "TRESP", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2J\u01f1\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3"+
		"!\3!\3!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\6%\u016b\n"+
		"%\r%\16%\u016c\3%\3%\6%\u0171\n%\r%\16%\u0172\5%\u0175\n%\3&\3&\5&\u0179"+
		"\n&\3&\3&\3\'\3\'\7\'\u017f\n\'\f\'\16\'\u0182\13\'\3\'\3\'\3(\3(\7(\u0188"+
		"\n(\f(\16(\u018b\13(\3)\3)\3)\3*\3*\3*\3+\3+\3,\3,\3,\3-\3-\3-\3.\3.\3"+
		"/\3/\3/\3\60\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3"+
		"\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3"+
		"?\3@\3@\3A\3A\3B\3B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3F\3F\3G\6G\u01d3\nG\r"+
		"G\16G\u01d4\3G\3G\3H\3H\3H\3H\7H\u01dd\nH\fH\16H\u01e0\13H\3H\3H\3H\3"+
		"H\3H\3I\3I\3I\3I\7I\u01eb\nI\fI\16I\u01ee\13I\3I\3I\3\u01de\2J\3\3\5\4"+
		"\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w"+
		"=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091"+
		"J\3\2\t\3\2\62;\3\2))\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f\17\17\"\""+
		"^^\4\2\f\f\17\17\2\u01f9\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2"+
		"\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3"+
		"\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2"+
		"\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2"+
		"Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3"+
		"\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2"+
		"\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2"+
		"w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2"+
		"\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b"+
		"\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\3\u0093\3\2\2"+
		"\2\5\u0097\3\2\2\2\7\u009d\3\2\2\2\t\u00a4\3\2\2\2\13\u00a9\3\2\2\2\r"+
		"\u00b3\3\2\2\2\17\u00b7\3\2\2\2\21\u00bb\3\2\2\2\23\u00c0\3\2\2\2\25\u00c6"+
		"\3\2\2\2\27\u00cc\3\2\2\2\31\u00cf\3\2\2\2\33\u00d4\3\2\2\2\35\u00db\3"+
		"\2\2\2\37\u00e0\3\2\2\2!\u00e8\3\2\2\2#\u00ee\3\2\2\2%\u00f7\3\2\2\2\'"+
		"\u00fe\3\2\2\2)\u0104\3\2\2\2+\u0108\3\2\2\2-\u010b\3\2\2\2/\u0111\3\2"+
		"\2\2\61\u0118\3\2\2\2\63\u0123\3\2\2\2\65\u012a\3\2\2\2\67\u0132\3\2\2"+
		"\29\u0138\3\2\2\2;\u0142\3\2\2\2=\u0149\3\2\2\2?\u014e\3\2\2\2A\u0157"+
		"\3\2\2\2C\u015c\3\2\2\2E\u015f\3\2\2\2G\u0165\3\2\2\2I\u016a\3\2\2\2K"+
		"\u0176\3\2\2\2M\u017c\3\2\2\2O\u0185\3\2\2\2Q\u018c\3\2\2\2S\u018f\3\2"+
		"\2\2U\u0192\3\2\2\2W\u0194\3\2\2\2Y\u0197\3\2\2\2[\u019a\3\2\2\2]\u019c"+
		"\3\2\2\2_\u019f\3\2\2\2a\u01a2\3\2\2\2c\u01a4\3\2\2\2e\u01a6\3\2\2\2g"+
		"\u01a8\3\2\2\2i\u01aa\3\2\2\2k\u01ac\3\2\2\2m\u01ae\3\2\2\2o\u01b0\3\2"+
		"\2\2q\u01b2\3\2\2\2s\u01b4\3\2\2\2u\u01b6\3\2\2\2w\u01b8\3\2\2\2y\u01ba"+
		"\3\2\2\2{\u01bc\3\2\2\2}\u01be\3\2\2\2\177\u01c0\3\2\2\2\u0081\u01c2\3"+
		"\2\2\2\u0083\u01c4\3\2\2\2\u0085\u01c7\3\2\2\2\u0087\u01c9\3\2\2\2\u0089"+
		"\u01cb\3\2\2\2\u008b\u01cd\3\2\2\2\u008d\u01d2\3\2\2\2\u008f\u01d8\3\2"+
		"\2\2\u0091\u01e6\3\2\2\2\u0093\u0094\7K\2\2\u0094\u0095\7p\2\2\u0095\u0096"+
		"\7v\2\2\u0096\4\3\2\2\2\u0097\u0098\7H\2\2\u0098\u0099\7n\2\2\u0099\u009a"+
		"\7q\2\2\u009a\u009b\7c\2\2\u009b\u009c\7v\2\2\u009c\6\3\2\2\2\u009d\u009e"+
		"\7U\2\2\u009e\u009f\7v\2\2\u009f\u00a0\7t\2\2\u00a0\u00a1\7k\2\2\u00a1"+
		"\u00a2\7p\2\2\u00a2\u00a3\7i\2\2\u00a3\b\3\2\2\2\u00a4\u00a5\7D\2\2\u00a5"+
		"\u00a6\7q\2\2\u00a6\u00a7\7q\2\2\u00a7\u00a8\7n\2\2\u00a8\n\3\2\2\2\u00a9"+
		"\u00aa\7E\2\2\u00aa\u00ab\7j\2\2\u00ab\u00ac\7c\2\2\u00ac\u00ad\7t\2\2"+
		"\u00ad\u00ae\7c\2\2\u00ae\u00af\7e\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1"+
		"\7g\2\2\u00b1\u00b2\7t\2\2\u00b2\f\3\2\2\2\u00b3\u00b4\7x\2\2\u00b4\u00b5"+
		"\7c\2\2\u00b5\u00b6\7t\2\2\u00b6\16\3\2\2\2\u00b7\u00b8\7n\2\2\u00b8\u00b9"+
		"\7g\2\2\u00b9\u00ba\7v\2\2\u00ba\20\3\2\2\2\u00bb\u00bc\7v\2\2\u00bc\u00bd"+
		"\7t\2\2\u00bd\u00be\7w\2\2\u00be\u00bf\7g\2\2\u00bf\22\3\2\2\2\u00c0\u00c1"+
		"\7h\2\2\u00c1\u00c2\7c\2\2\u00c2\u00c3\7n\2\2\u00c3\u00c4\7u\2\2\u00c4"+
		"\u00c5\7g\2\2\u00c5\24\3\2\2\2\u00c6\u00c7\7r\2\2\u00c7\u00c8\7t\2\2\u00c8"+
		"\u00c9\7k\2\2\u00c9\u00ca\7p\2\2\u00ca\u00cb\7v\2\2\u00cb\26\3\2\2\2\u00cc"+
		"\u00cd\7k\2\2\u00cd\u00ce\7h\2\2\u00ce\30\3\2\2\2\u00cf\u00d0\7g\2\2\u00d0"+
		"\u00d1\7n\2\2\u00d1\u00d2\7u\2\2\u00d2\u00d3\7g\2\2\u00d3\32\3\2\2\2\u00d4"+
		"\u00d5\7u\2\2\u00d5\u00d6\7y\2\2\u00d6\u00d7\7k\2\2\u00d7\u00d8\7v\2\2"+
		"\u00d8\u00d9\7e\2\2\u00d9\u00da\7j\2\2\u00da\34\3\2\2\2\u00db\u00dc\7"+
		"e\2\2\u00dc\u00dd\7c\2\2\u00dd\u00de\7u\2\2\u00de\u00df\7g\2\2\u00df\36"+
		"\3\2\2\2\u00e0\u00e1\7f\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7h\2\2\u00e3"+
		"\u00e4\7c\2\2\u00e4\u00e5\7w\2\2\u00e5\u00e6\7n\2\2\u00e6\u00e7\7v\2\2"+
		"\u00e7 \3\2\2\2\u00e8\u00e9\7d\2\2\u00e9\u00ea\7t\2\2\u00ea\u00eb\7g\2"+
		"\2\u00eb\u00ec\7c\2\2\u00ec\u00ed\7m\2\2\u00ed\"\3\2\2\2\u00ee\u00ef\7"+
		"e\2\2\u00ef\u00f0\7q\2\2\u00f0\u00f1\7p\2\2\u00f1\u00f2\7v\2\2\u00f2\u00f3"+
		"\7k\2\2\u00f3\u00f4\7p\2\2\u00f4\u00f5\7w\2\2\u00f5\u00f6\7g\2\2\u00f6"+
		"$\3\2\2\2\u00f7\u00f8\7t\2\2\u00f8\u00f9\7g\2\2\u00f9\u00fa\7v\2\2\u00fa"+
		"\u00fb\7w\2\2\u00fb\u00fc\7t\2\2\u00fc\u00fd\7p\2\2\u00fd&\3\2\2\2\u00fe"+
		"\u00ff\7y\2\2\u00ff\u0100\7j\2\2\u0100\u0101\7k\2\2\u0101\u0102\7n\2\2"+
		"\u0102\u0103\7g\2\2\u0103(\3\2\2\2\u0104\u0105\7h\2\2\u0105\u0106\7q\2"+
		"\2\u0106\u0107\7t\2\2\u0107*\3\2\2\2\u0108\u0109\7k\2\2\u0109\u010a\7"+
		"p\2\2\u010a,\3\2\2\2\u010b\u010c\7i\2\2\u010c\u010d\7w\2\2\u010d\u010e"+
		"\7c\2\2\u010e\u010f\7t\2\2\u010f\u0110\7f\2\2\u0110.\3\2\2\2\u0111\u0112"+
		"\7c\2\2\u0112\u0113\7r\2\2\u0113\u0114\7r\2\2\u0114\u0115\7g\2\2\u0115"+
		"\u0116\7p\2\2\u0116\u0117\7f\2\2\u0117\60\3\2\2\2\u0118\u0119\7t\2\2\u0119"+
		"\u011a\7g\2\2\u011a\u011b\7o\2\2\u011b\u011c\7q\2\2\u011c\u011d\7x\2\2"+
		"\u011d\u011e\7g\2\2\u011e\u011f\7N\2\2\u011f\u0120\7c\2\2\u0120\u0121"+
		"\7u\2\2\u0121\u0122\7v\2\2\u0122\62\3\2\2\2\u0123\u0124\7t\2\2\u0124\u0125"+
		"\7g\2\2\u0125\u0126\7o\2\2\u0126\u0127\7q\2\2\u0127\u0128\7x\2\2\u0128"+
		"\u0129\7g\2\2\u0129\64\3\2\2\2\u012a\u012b\7k\2\2\u012b\u012c\7u\2\2\u012c"+
		"\u012d\7G\2\2\u012d\u012e\7o\2\2\u012e\u012f\7r\2\2\u012f\u0130\7v\2\2"+
		"\u0130\u0131\7{\2\2\u0131\66\3\2\2\2\u0132\u0133\7e\2\2\u0133\u0134\7"+
		"q\2\2\u0134\u0135\7w\2\2\u0135\u0136\7p\2\2\u0136\u0137\7v\2\2\u01378"+
		"\3\2\2\2\u0138\u0139\7t\2\2\u0139\u013a\7g\2\2\u013a\u013b\7r\2\2\u013b"+
		"\u013c\7g\2\2\u013c\u013d\7c\2\2\u013d\u013e\7v\2\2\u013e\u013f\7k\2\2"+
		"\u013f\u0140\7p\2\2\u0140\u0141\7i\2\2\u0141:\3\2\2\2\u0142\u0143\7u\2"+
		"\2\u0143\u0144\7v\2\2\u0144\u0145\7t\2\2\u0145\u0146\7w\2\2\u0146\u0147"+
		"\7e\2\2\u0147\u0148\7v\2\2\u0148<\3\2\2\2\u0149\u014a\7u\2\2\u014a\u014b"+
		"\7g\2\2\u014b\u014c\7n\2\2\u014c\u014d\7h\2\2\u014d>\3\2\2\2\u014e\u014f"+
		"\7o\2\2\u014f\u0150\7w\2\2\u0150\u0151\7v\2\2\u0151\u0152\7c\2\2\u0152"+
		"\u0153\7v\2\2\u0153\u0154\7k\2\2\u0154\u0155\7p\2\2\u0155\u0156\7i\2\2"+
		"\u0156@\3\2\2\2\u0157\u0158\7h\2\2\u0158\u0159\7w\2\2\u0159\u015a\7p\2"+
		"\2\u015a\u015b\7e\2\2\u015bB\3\2\2\2\u015c\u015d\7c\2\2\u015d\u015e\7"+
		"v\2\2\u015eD\3\2\2\2\u015f\u0160\7k\2\2\u0160\u0161\7p\2\2\u0161\u0162"+
		"\7q\2\2\u0162\u0163\7w\2\2\u0163\u0164\7v\2\2\u0164F\3\2\2\2\u0165\u0166"+
		"\7p\2\2\u0166\u0167\7k\2\2\u0167\u0168\7n\2\2\u0168H\3\2\2\2\u0169\u016b"+
		"\t\2\2\2\u016a\u0169\3\2\2\2\u016b\u016c\3\2\2\2\u016c\u016a\3\2\2\2\u016c"+
		"\u016d\3\2\2\2\u016d\u0174\3\2\2\2\u016e\u0170\7\60\2\2\u016f\u0171\t"+
		"\2\2\2\u0170\u016f\3\2\2\2\u0171\u0172\3\2\2\2\u0172\u0170\3\2\2\2\u0172"+
		"\u0173\3\2\2\2\u0173\u0175\3\2\2\2\u0174\u016e\3\2\2\2\u0174\u0175\3\2"+
		"\2\2\u0175J\3\2\2\2\u0176\u0178\7$\2\2\u0177\u0179\n\3\2\2\u0178\u0177"+
		"\3\2\2\2\u0178\u0179\3\2\2\2\u0179\u017a\3\2\2\2\u017a\u017b\7$\2\2\u017b"+
		"L\3\2\2\2\u017c\u0180\7$\2\2\u017d\u017f\n\4\2\2\u017e\u017d\3\2\2\2\u017f"+
		"\u0182\3\2\2\2\u0180\u017e\3\2\2\2\u0180\u0181\3\2\2\2\u0181\u0183\3\2"+
		"\2\2\u0182\u0180\3\2\2\2\u0183\u0184\7$\2\2\u0184N\3\2\2\2\u0185\u0189"+
		"\t\5\2\2\u0186\u0188\t\6\2\2\u0187\u0186\3\2\2\2\u0188\u018b\3\2\2\2\u0189"+
		"\u0187\3\2\2\2\u0189\u018a\3\2\2\2\u018aP\3\2\2\2\u018b\u0189\3\2\2\2"+
		"\u018c\u018d\7#\2\2\u018d\u018e\7?\2\2\u018eR\3\2\2\2\u018f\u0190\7?\2"+
		"\2\u0190\u0191\7?\2\2\u0191T\3\2\2\2\u0192\u0193\7#\2\2\u0193V\3\2\2\2"+
		"\u0194\u0195\7~\2\2\u0195\u0196\7~\2\2\u0196X\3\2\2\2\u0197\u0198\7(\2"+
		"\2\u0198\u0199\7(\2\2\u0199Z\3\2\2\2\u019a\u019b\7?\2\2\u019b\\\3\2\2"+
		"\2\u019c\u019d\7@\2\2\u019d\u019e\7?\2\2\u019e^\3\2\2\2\u019f\u01a0\7"+
		">\2\2\u01a0\u01a1\7?\2\2\u01a1`\3\2\2\2\u01a2\u01a3\7@\2\2\u01a3b\3\2"+
		"\2\2\u01a4\u01a5\7>\2\2\u01a5d\3\2\2\2\u01a6\u01a7\7\'\2\2\u01a7f\3\2"+
		"\2\2\u01a8\u01a9\7,\2\2\u01a9h\3\2\2\2\u01aa\u01ab\7\61\2\2\u01abj\3\2"+
		"\2\2\u01ac\u01ad\7-\2\2\u01adl\3\2\2\2\u01ae\u01af\7/\2\2\u01afn\3\2\2"+
		"\2\u01b0\u01b1\7*\2\2\u01b1p\3\2\2\2\u01b2\u01b3\7+\2\2\u01b3r\3\2\2\2"+
		"\u01b4\u01b5\7}\2\2\u01b5t\3\2\2\2\u01b6\u01b7\7\177\2\2\u01b7v\3\2\2"+
		"\2\u01b8\u01b9\7]\2\2\u01b9x\3\2\2\2\u01ba\u01bb\7_\2\2\u01bbz\3\2\2\2"+
		"\u01bc\u01bd\7<\2\2\u01bd|\3\2\2\2\u01be\u01bf\7\60\2\2\u01bf~\3\2\2\2"+
		"\u01c0\u01c1\7.\2\2\u01c1\u0080\3\2\2\2\u01c2\u01c3\7A\2\2\u01c3\u0082"+
		"\3\2\2\2\u01c4\u01c5\7/\2\2\u01c5\u01c6\7@\2\2\u01c6\u0084\3\2\2\2\u01c7"+
		"\u01c8\7(\2\2\u01c8\u0086\3\2\2\2\u01c9\u01ca\7=\2\2\u01ca\u0088\3\2\2"+
		"\2\u01cb\u01cc\7a\2\2\u01cc\u008a\3\2\2\2\u01cd\u01ce\7\60\2\2\u01ce\u01cf"+
		"\7\60\2\2\u01cf\u01d0\7\60\2\2\u01d0\u008c\3\2\2\2\u01d1\u01d3\t\7\2\2"+
		"\u01d2\u01d1\3\2\2\2\u01d3\u01d4\3\2\2\2\u01d4\u01d2\3\2\2\2\u01d4\u01d5"+
		"\3\2\2\2\u01d5\u01d6\3\2\2\2\u01d6\u01d7\bG\2\2\u01d7\u008e\3\2\2\2\u01d8"+
		"\u01d9\7\61\2\2\u01d9\u01da\7,\2\2\u01da\u01de\3\2\2\2\u01db\u01dd\13"+
		"\2\2\2\u01dc\u01db\3\2\2\2\u01dd\u01e0\3\2\2\2\u01de\u01df\3\2\2\2\u01de"+
		"\u01dc\3\2\2\2\u01df\u01e1\3\2\2\2\u01e0\u01de\3\2\2\2\u01e1\u01e2\7,"+
		"\2\2\u01e2\u01e3\7\61\2\2\u01e3\u01e4\3\2\2\2\u01e4\u01e5\bH\2\2\u01e5"+
		"\u0090\3\2\2\2\u01e6\u01e7\7\61\2\2\u01e7\u01e8\7\61\2\2\u01e8\u01ec\3"+
		"\2\2\2\u01e9\u01eb\n\b\2\2\u01ea\u01e9\3\2\2\2\u01eb\u01ee\3\2\2\2\u01ec"+
		"\u01ea\3\2\2\2\u01ec\u01ed\3\2\2\2\u01ed\u01ef\3\2\2\2\u01ee\u01ec\3\2"+
		"\2\2\u01ef\u01f0\bI\2\2\u01f0\u0092\3\2\2\2\f\2\u016c\u0172\u0174\u0178"+
		"\u0180\u0189\u01d4\u01de\u01ec\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}