grammar GoSubset;

program  : 'package' IDENT topLevel+ EOF ;
topLevel : funcDecl
         | typeDecl
         ;

funcDecl  : 'func' IDENT '(' paramList? ')' returnSpec? block ;
typeDecl  : 'type' IDENT typeDef ;
typeDef   : 'struct' '{' fieldDecl* '}'
          | 'interface' '{' methodDecl* '}'
          ;
fieldDecl : IDENT type ';' ;
methodDecl: IDENT '(' paramList? ')' returnSpec? ';' ;
paramList : param (',' param)* ;
param     : IDENT type ;
returnSpec : type
           | '(' type (',' type)* ')'
           ;

block : '{' statement* '}' ;

statement : varDecl
          | goStmt
          | multiShortDecl
          | shortVarDecl
          | fieldAssignStmt
          | indexAssignStmt
          | assignStmt
          | ifStmt
          | forStmt
          | breakStmt
          | continueStmt
          | returnStmt
          | callStmt
          | block
          | printStmt
          ;

varDecl         : 'var' IDENT type ( '=' expr )? ';' ;
shortVarDecl    : IDENT ':=' expr ';' ;
assignStmt      : IDENT '=' expr ';' ;
fieldAssignStmt : IDENT '.' IDENT '=' expr ';' ;
indexAssignStmt : IDENT '[' expr ']' '=' expr ';' ;

returnStmt    : 'return' exprList? ';' ;
callStmt      : IDENT '(' exprList? ')' ';' ;
multiShortDecl: IDENT (',' IDENT)+ ':=' IDENT '(' exprList? ')' ';' ;

ifStmt   : 'if' expr block elsePart? ;
elsePart : 'else' ifStmt
         | 'else' block
         ;

forStmt : 'for' block                                                         # ForInfinite
        | 'for' IDENT ',' IDENT ':=' 'range' expr block                       # ForRangeFull
        | 'for' IDENT ':=' 'range' expr block                                 # ForRangeIndex
        | 'for' init=simpleStmt? ';' cond=expr? ';' post=simpleStmt? block    # ForClause
        | 'for' expr block                                                    # ForWhile
        ;


simpleStmt : IDENT ':=' expr  # SimpleShort
           | IDENT '=' expr    # SimpleAssign
           ;

breakStmt    : 'break' ';' ;
continueStmt : 'continue' ';' ;

type : 'int' | 'float64' | 'string' | 'bool' | '[' ']' type | IDENT ;

fieldInitList : fieldInit (',' fieldInit)* ;
fieldInit     : IDENT ':' expr ;

printStmt : 'fmt' '.' ( 'Println' | 'Printf' ) '(' exprList? ')' ';' ;
goStmt    : 'go' ( IDENT '(' exprList? ')' | 'fmt' '.' ( 'Println' | 'Printf' ) '(' exprList? ')' ) ';' ;
exprList  : expr ( ',' expr )* ;

expr
    : '(' expr ')'                                  # ParenExpr
    | op=( '-' | '!' ) expr                         # UnaryExpr
    | expr '.' IDENT                                # SelectorExpr
    | expr '[' expr ']'                             # IndexExpr
    | expr op=( '*' | '/' | '%' ) expr              # MulExpr
    | expr op=( '+' | '-' ) expr                    # AddExpr
    | expr op=( '<' | '<=' | '>' | '>=' ) expr      # RelExpr
    | expr op=( '==' | '!=' ) expr                  # EqExpr
    | expr '&&' expr                                # AndExpr
    | expr '||' expr                                # OrExpr
    | INT                                           # IntExpr
    | FLOAT                                         # FloatExpr
    | STRING                                        # StringExpr
    | TRUE                                          # TrueExpr
    | FALSE                                         # FalseExpr
    | IDENT '(' exprList? ')'                       # CallExpr
    | '[' ']' type '{' exprList? '}'                # SliceLitExpr
    | IDENT '{' fieldInitList? '}'                  # StructLitExpr
    | IDENT                                         # IdentExpr
    ;

TRUE   : 'true' ;
FALSE  : 'false' ;

IDENT  : [a-zA-Z_] [a-zA-Z0-9_]* ;
FLOAT  : [0-9]+ '.' [0-9]+ ;
INT    : [0-9]+ ;
STRING : '"' ( ~["\\\r\n] | '\\' . )* '"' ;

WS            : [ \t\r\n]+    -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;
