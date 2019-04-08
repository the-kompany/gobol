# GOBOL Language Reference
by Shawn M. Gordon
---------------------------------------
  * [Looping](#looping)
    * [PERFORM](#perform)
  * [Conditionals](#conditionals)
    * [IF..ELSE..END-IF](#if-else-end-if)
    * [EVALUATE](#evaluate)
  * [Date Operators](#date-operators)
    * [Date Addition](#date-addition)
    * [Date Subtraction](#date-subtraction)
    * [Date Comparison](#date-comparison)
  * [Functions](#functions)
    * [ACCEPT](#accept)
    * [DATE2STR](#date2str)
    * [DOWNSHIFT](#downshift)
    * [EXTRACT](#extract)
    * [GETENV](#getenv)
    * [ISDATE](#isdate)
    * [ISNUMERIC](#isnumeric)
    * [MOVE](#move)
    * [ROUND](#round)
    * [STR2DATE](#str2date)
    * [UPSHIFT](#upshift)
  * [File Types](#file-types)
  * [Some Psuedo Code](#some-psuedo-code)

---------------------------------------
## Looping
The PERFORM statement is used to define loops which are executed *until* a condition is true (not while true, which is more common in other languages).

### PERFORM

Perform-statement (the only looping construct).


		[ times-phrase   ]
	PERFORM	[ until-phrase   ] imperative-statement-1
		[ varying-phrase ]
		[ read-phrase    ]
	END-PERFORM
 
	times-phrase.
		{ identifier-1 } TIMES
		{ integer-1    }
 
	until-phrase.
		[ [ WITH ] TEST { BEFORE } ] UNTIL condition-1
                		{ AFTER  }
 
	varying-phrase.
		[ [ WITH ] TEST { BEFORE } ]
				{ AFTER  }
 
 		VARYING	{ identifier-2 }
			{ index-name-2 }
		FROM	{ identifier-3 }
                	{ index-name-3 }
               		{ literal-3    }
 
		[ BY	{ identifier-4 } ]
                	{ literal-4    }
 
		UNTIL condition-1
 			[ AFTER	{ identifier-2 }
				{ index-name-2 }
 
		FROM 	{ identifier-3 }
                   	{ index-name-3 }
                   	{ literal-3    }

		[ BY 	{ identifier-4 } ]
                   	{ literal-4    }
		UNTIL condition-1 ] ...

		read-phrase.
			PERFORM READ <file> into <record structure> [conditional tests as described above] UNTIL EOF
	// Maybe use this to perform a function in the code, have to think about how functions would be defined and variables passed or just make all variables global.



## Conditionals

## IF..ELSE..END-IF

**IF..ELSE..END-IF** Evaluates a Boolean expression and branches based on the outcome of the evaluation. Functions, variables, literals and expressions can be used in the statement.

**Examples:**

	IF (ISNUMERIC(“HELLO”)) THEN
	   MOVE “problem” 	TO var1
	ELSE
	   MOVE “sweet” 	TO var1
	END-IF

	IF (DOWNSHIFT(mystring) = “gobol) THEN
	   MOVE “20190401” TO MYDATE

	IF (numvar > 10) THEN…

	IF (ISDATE(MYDATE,"YYYYMMDD")) THEN ….

### EVALUATE
The EVALUATE statement causes multiple conditions to be evaluated. The subsequent action of the program depends on the results of these evaluations.
 
The EVALUATE statement is very similar to the CASE construct common in many other programming languages. The EVALUATE/CASE construct provides the ability to selectively execute one of a set of instruction alternatives based on the evaluation of a set of choice alternatives.
 
EVALUATE extends the power of the typical CASE construct by allowing multiple data items and conditions to be named in the EVALUATE phrase

	
	EVALUATE {subject}  [ ALSO {subject} ] ...  
             	{TRUE   }         {TRUE   }  
            	{FALSE  }         {FALSE  }  
    
     { { WHEN obj-phrase [ ALSO obj-phrase ] ... } ...  
                statement-1 } ...  
    
     [ WHEN OTHER statement-2 ]  
    
     [ END-EVALUATE ] 

	obj-phrase has the following format:
		{ ANY                                  }  
		{ TRUE                                 }  
		{ FALSE                                }  
		{ cond-obj                             }  
		{ [NOT] obj-item [ {THRU   } obj-item ]}  

#### Syntax Rules
1. *Subject* may be a literal, data item, arithmetic expression, or conditional expression.
1. *Cond-obj* is a conditional expression.
1. *Obj-item* may be a literal, data item, or arithmetic expression.
1. *Statement-1* and statement-2 are imperative statements.
1. Before the first WHEN phrase, *subject* and the words TRUE and FALSE are called "subjects," and all the subjects together are called the "subject set".
1. The operands and the words TRUE, FALSE, and ANY which appear in a WHEN phrase are called "objects," and the collection of objects in a single WHEN phrase is called the "object set".
1. Two *obj-items* connected by a THRU phrase must be of the same class. They are treated as a single object.
1. The number of objects within each object set must match the number of subjects in the subject set.
1. Each object within an object set must correspond to the subject having the same ordinal position as in the subject set. For each pair:
  1. *Obj-item* must be a valid operand for comparison to the corresponding *subject*.
  1. TRUE, FALSE, or *cond-obj* as an object must correspond to TRUE, FALSE, or a conditional expression as the subject.
  1. ANY may correspond to any type of subject.

#### General Rules
1. The EVALUATE statement operates as if each subject and object were evaluated and assigned a value or range of values. These values may be numeric, nonnumeric, truth values, or ranges of numeric or nonnumeric values. These values are determined as follows:
  1. Any subject or object that is a data item or literal, without either the THROUGH or the NOT phrase, is assigned the value and class of that data item or literal.
  1. Any subject or object that is an arithmetic expression, without either the THROUGH or the NOT phrase, is assigned a numeric value according to the rules for evaluating arithmetic expressions.
  1. Any subject or object that is a conditional expression is assigned a truth value according to the rules for evaluating conditional expressions.
  1. Any subject or object specified by the words TRUE or FALSE is assigned a truth value corresponding to that word.
  1. Any object specified by the word ANY is not evaluated.
  1. If the THROUGH phrase is specified for an object, without the NOT phrase, the range of values includes all permissible values of the corresponding subject that are greater than or equal to the first operand and less than or equal to the second operand, according to the rules for comparison.
  1. If the NOT phrase is specified for an object, the values assigned to that object are all permissible values of the corresponding subject not equal to the value, or range of values, that would have been assigned had the NOT phrase been omitted.
1. The EVALUATE statement then proceeds as if the values assigned to the subjects and objects were compared to determine if any WHEN phrase satisfies the subject set. Each object within the object set for the first WHEN phrase is compared to the subject having the same ordinal position within the subject set. The comparison is satisfied if one of the following is true:
  1. If the items being compared are assigned numeric or nonnumeric values, the comparison is satisfied if the value (or one of the range of values) assigned to the object is equal to the value assigned to the subject.
  1. If the items being compared are assigned truth values, the comparison is satisfied if the truth values are the same.
  1. If the object is the word ANY, the comparison is always satisfied.
1. If the comparison is satisfied for every object within the object set, the corresponding WHEN phrase is selected.
1. If the comparison is not satisfied for one or more objects within the object set, the procedure repeats for the next WHEN phrase. This is repeated until a WHEN phrase is selected or all the object sets have been tested.
1. If a WHEN phrase is selected, the corresponding statement-1 is executed.
1. If no WHEN phrase is selected and a WHEN OTHER phrase is specified, statement-2 is executed. If no WHEN OTHER phrase is present, control transfers to the end of the EVALUATE statement.
The scope of execution of the EVALUATE statement is terminated when the end of statement-1 or statement-2 is reached, or when no WHEN phrase is selected and no WHEN OTHER phrase is specified.

#### Code Examples
**Example 1:**

	EVALUATE AGE  
		WHEN 56 THRU 99  MOVE “S” TO PROSPECT-TYPE  
		WHEN 40 THRU 55  MOVE “M” TO PROSPECT-TYPE 
		WHEN 21 THRU 39  MOVE “Y” TO PROSPECT-TYPE  
		WHEN OTHER       MOVE “N” TO PROSPECT-TYPE   
	END-EVALUATE. 

**Example 2:**
  
	EVALUATE INCOME ALSO TRUE  
		WHEN 20000 THRU 39999  ALSO RISK_CLASS = "A"  
			SET LOW_INCOME_PROSPECT TO TRUE 
		WHEN 40000 THRU 59999  ALSO RISK_CLASS = "A"  
			SET MID_INCOME_PROSPECT TO TRUE 
		WHEN 60000 THRU 999999 ALSO RISK_CLASS = "A"  
			SET HIGH_INCOME_PROSPECT TO TRUE 
		WHEN 60000 THRU 999999 ALSO NOT RISK_CLASS = "A"  
			SET HIGH_INCOME_HIGH_RISK_PROSPECT TO TRUE  
		WHEN OTHER  
			SET UNCLASSIFIED_PROSPECT TO TRUE
	END-EVALUATE. 

**Highlights for first-time users**

1. Statement subjects (associated with the EVALUATE phrase) and statement objects (associated with the WHEN phrase) must be equal in number, correspond by position and be valid operands for comparison. Note the number and order of subjects in example 2 and the correspondent number and position of WHEN objects.
1. If all of the conditions in a WHEN phrase match, the associated imperative statement is executed. None of the remaining WHEN phrases is evaluated. Program execution then falls through to the end of the EVALUATE statement.
1. The WHEN OTHER phrase is an optional phrase for the handling of all remaining cases (the set of possible conditions not explicitly tested for by the preceding WHEN phrases). The WHEN OTHER phrase, if present, must be the last WHEN phrase in the statement.
1. The words TRUE and FALSE may be used in the subject or object phrase to specify a literal truth condition.
1. The word ANY may be used in the WHEN phrase to specify an unconditional match with the corresponding item in the subject phrase.
1. The word NOT may be used in the WHEN phrase to negate its associated condition.
1. The word THRU may be used in the WHEN phrase to describe a range of values. When combined with NOT, THRU describes an excluded set of values. For example, NOT 10 THRU 20 means that any object holding a value from 10 to 20, including the numbers 10 and 20, will result in a FALSE, or no match evaluation.


## Date Operators

Date operators are used to perform operations on date data. In general: two dates/times may be subtracted to produce an interval, a date/time and an interval may be added to produce another date/time, dates/times may be compared with each other.

### Date Addition
An INTERVAL may be added to a DATE or DATETIME resulting in a DATETIME.  An INTERVAL may also be added to a TIME or another INTERVAL resulting in an INTERVAL. The following table shows the result of date addition:

	DATETIME + DATETIME = Error 
	DATETIME + DATE     = Error 
	DATETIME + TIME     = DATETIME 
	DATETIME + INTERVAL = DATETIME 
	DATE     + DATETIME = Error 
	DATE     + DATE     = Error 
	DATE     + TIME     = DATETIME 
	DATE     + INTERVAL = DATETIME 
	TIME     + DATETIME = DATETIME 
	TIME     + DATE     = DATETIME 
	TIME     + TIME     = INTERVAL 
	TIME     + INTERVAL = INTERVAL 
	INTERVAL + DATETIME = DATETIME 
	INTERVAL + DATE     = DATETIME 
	INTERVAL + TIME     = INTERVAL 
	INTERVAL + INTERVAL = INTERVAL


### Date Subtraction
A DATE or DATETIME may be subtracted from another DATE or DATETIME resulting in an INTERVAL.  Two INTERVALs or two TIMEs may also be subtracted resulting in an INTERVAL. The following table shows the result of date subtraction:

	DATETIME - DATETIME = INTERVAL 
	DATETIME - DATE     = INTERVAL 
	DATETIME - TIME     = DATETIME 
	DATETIME - INTERVAL = DATETIME 
	DATE     - DATETIME = INTERVAL 
	DATE     - DATE     = INTERVAL 
	DATE     - TIME     = DATETIME 
	DATE     - INTERVAL = DATETIME 
	TIME     - DATETIME = Error
	TIME     - DATE     = Error 
	TIME     - TIME     = INTERVAL 
	TIME     - INTERVAL = INTERVAL 
	INTERVAL - DATETIME = Error 
	INTERVAL - DATE     = Error 
	INTERVAL - TIME     = INTERVAL 
	INTERVAL - INTERVAL = INTERVAL


### Date Comparison
Date comparison is done using the operators: <, <=, =, >=, >, and <>.  
A date comparison is legal between DATEs, and DATETIMEs, or between TIMEs and INTERVALs, but not between DATEs/DATETIMEs and TIMEs/INTERVALs. The following table shows the legal date compares based on the date subtype:

Date1 | Date2 | Is Legal?
------|-------|---------- 
DATETIME | DATETIME | Yes 
DATETIME | DATE | Yes 
DATETIME | TIME | No 
DATETIME | INTERVAL | No 
DATE | DATETIME | Yes 
DATE | DATE | Yes 
DATE | TIME | No 
DATE | INTERVAL | No 
TIME | DATETIME | No 
TIME | DATE | No 
TIME | TIME | Yes 
TIME | INTERVAL | Yes 
INTERVAL | DATETIME | No 
INTERVAL | DATE | No 
INTERVAL | TIME | Yes 
INTERVAL | INTERVAL | Yes


## Functions
GOBOL has a number of built-in functions for date, string and file management to make coding faster. They are as follows:

### ACCEPT 
Reads a string from the user’s standard input (stdin) after displaying a prompt string.

	Usage: ACCEPT(prompt-string) -> string

**Example:**

	MOVE ACCEPT("Enter your name: ") TO FIRST-NAME.

### DATE2STR 		
Converts a date, time, datetime, or interval to a string using a date format. DATE2STR is the inverse function of STR2DATE (which converts a string to a date).

	Usage: MOVE DATE2STR(date, fmt-string) TO FMT-DATE

*date* is the date, time, datetime, or interval type item to be converted to a string.

*fmt-string* is a string containing tokens that describe the format of date-string. The tokens are the same as those used in the print PIC of a date item and the **DATE2STR** function. The allowable tokens in fmt-string are:

	A.M.		AM/PM indicator with periods 
	AM 		AM/PM indicator 
	AY		Two character year and century where 00-99 is century
	CC		2 digit century 
	D		The day of week (1-7, Sun=1,Sat=7) 
	DAY		The 9 character name of day of week (SUNDAY-SATURDAY)
	DD		The 2 digit day number within month (1-31) 
	D*		The 1 or 2 digit day number within month (1-31) 
	DY		The 3 character name of day of week (SUN-SAT) 
	HH		The 2 digit hour in 12 hour time (0112) 
	HH12		The 2 digit hour in 12 hour time (0112) 
	HH24		The 2 digit hour in 24 hour time (0023) 
	H*		The 1 or 2 digit hour in 12 hour time (1-12) 
	H*12		The 1 or 2 digit hour in 12 hour time (1-12) 
	H*24		The 1 or 2 digit hour in 24 hour time (0-23) 
	MI		The 2 digit minute within the hour (00-59) 
	MM		The 2 digit month number within year (01-12) 
	M*		The 1 or 2 digit month number within year (1-12) 
	MON		The 3 character name of month (JAN-DEC) 
	MONTH		The 9 character name of month (JANUARY-DECEMBER)
	NNN... 		Number of days  (up to 9 Ns) 
	P.M. 		AM/PM indicator with periods 
	PM		AM/PM indicator 
	Q		Quarter within year (1-4) 
	SS		The 2 digit second within the minute (00-59) 
	SSSSS		The 5 digit second within the day (086399) 
	TTT...		Fractions of seconds (up to 9 Ts) 
	W		The week within the month (1-5) 
	WW		The 2 digit week within the year (0153) 
	Y		The last digit of the year 
	YY		The last 2 digits of the year. 
	YYY		The last 3 digits of the year 
	YYYY		The 4 digit year 
	Y,YYY		Year with comma 
	space		Space 
	:		Colon 
	/		Forward slash 
	-		Hyphen
	.		Period 
	,		Comma 
	;		Semicolon 
	"str" 		Quoted string


Date format items are case insensitive, except that the case determines the appearance of alphabetic date items.

**Examples:** 	

The following examples assume the following statements:

	DEFINE DTM : DATETIME 
	MOVE DATE2STR("960401 1504","YYMMDD HH24MI") TO DTM

	Expression				Result
	DATE2STR(DTM,"YYMMDD") 			960401 
	DATE2STR(DTM,"DD-MON-YY") 		1-APR-96
	DATE2STR(DTM,"DD-Mon-YY") 		1-Apr-96
	DATE2STR(DTM,"HH24:MI:SS") 		15:04:00 
	DATE2STR(DTM,"Month DD,YY")		April 1,96 
	DATE2STR(DTM,"Day, Mon DD")		Mon, Apr 1
	DATE2STR(DTM,"M*/D*/YY") 		4/1/96

### DOWNSHIFT	

Downshifts a string. DOWNSHIFT converts all uppercase characters in a given string to lowercase characters. Can be used on a variable in place, or in conjunction to the MOVE verb to modify the target value to be downshifted.

To upshift a string use the UPSHIFT function.

**Usage:**
 	
	MOVE DOWNSHIFT(string) TO VAR
		or
	DOWNSHIFT(VAR)


**Examples:** 		

	Expression			Result
	MOVE "Gobol" to VAR
	DOWNSHIFT(VAR)			"gobol"
	MOVE DOWNSHIFT("GOBOL") to VAR	"gobol"

### EXTRACT	
Used to extract data from a record or variable by position and length. The EXTRACT function requires three parameters.  The first parameter is a variable name, the second parameter is a byte index into the first parameter, and the third parameter is a length for the data to be extracted. It is possible to cast the results to another type with functions such as STR2DATE and STR2NUM using the MOVE verb during the operation.


**Usage:** 	

	EXTRACT (var, index, length) -> result


*var* can be a defined variable, a record name or a field from a record. Note:  You must be familiar with the data type of var to get the desired results.

*index* is the byte index of the first byte of data to be extracted.  The first byte is byte 1.

*length* is the length of the string you want to extract. 

**Examples:** 	

	The following examples assume the script:
		DEFINE R : RECORD 
				F : CHAR(4) 
				G : CHAR(6)
				H : CHAR(2) 
		END 

		MOVE "ABCD"   TO R.F
		MOVE "efghij" TO R.G
		MOVE “13”     TO R.H

		Expression                   			Result
		MOVE EXTRACT(R, 1, 3)   TO var			var = "ABC" 
		MOVE EXTRACT(R, 3, 3)   TO var			var = "CDe" 
		MOVE EXTRACT(R.F, 1, 3) TO var			var = "ABC" 
		MOVE EXTRACT(R.F, 3, 3) TO var			Error(1) 
		MOVE EXTRACT(R.G, 3, 3) TO var			var = "ghi" 
		MOVE STR2NUM(EXTRACT(R, 11, 2)) TO numvar	numvar = 13 
	
	(1) Field overflow.  R.F is only 4 characters long, so taking 4 characters starting at character 3 overflows the field.

### GETENV
Returns the value of an environment variable when passed the variable name. If there is no variable of the name given, a null string is returned.

**Usage:**

GETENV(string) -> string


**Example:** 	

	Expression                   		Result
	MOVE GETENV("PATH") TO var		var = "C:\WINDOWS\"

### ISDATE	

Determine if a variable is a date with an optional date format. This function would normally be used in a conditional IF statement.

**Usage:** 

	ISDATE(date-str) -> boolean
		or
	ISDATE(date-str, format-str) -> boolean

*date-str* is the string to be tested for a date value.
*format-str* is optional and if present is a string containing tokens that describe the format of date-str. See the *STR2DATE* function for information on allowable format strings.

If format-str is absent, ISDATE examines the string and returns TRUE if the string is in a recognizable format that can be converted to a date using the STR2DATE function.  ISDATE can automatically recognize dates, times,  datetimes, and intervals. See the STR2DATE function for information on how dates are interpreted.

**Examples:** 

	Expression                   		Result
	ISDATE("20180401","YYYYMMDD") 		True 
	ISDATE("1995/08/31","YYYY/MM/DD") 	True 
	ISDATE("121224","RRMMDD") 		True 
	ISDATE("121224","HHMISS") 		True 
	ISDATE("Sep  8,92","MON DD,YY") 	True 
	ISDATE("  960401","YYMMDD") 		False
	ISDATE("12-12-24","YY/MM/DD") 		False
	ISDATE("16:12:24","HH24:MI:SS") 	True 
	ISDATE("16:12:24","HH:MI:SS") 		False 
	ISDATE("960401") 			True 
	ISDATE("12/12/2)			True 
	ISDATE("12/12/1924") 			True 
	ISDATE("121224") 			True 
	ISDATE("12-Dec-24") 			True 
	ISDATE("Sep  8,92") 			True
	ISDATE("  960401")			True 
	ISDATE("12-12-24") 			True 
	ISDATE("16:12:24") 			True

### ISNUMERIC 	
Tests if a string can be converted to a numeric value. If the stringcan be successfully converted to a numeric value TRUE is returned, otherwise FALSE is returned. If TRUE is returned, then the NUMERIC function can convert the string to a numeric value without error. A numeric string may contain a decimal point, an E followed by an exponent, a single leading or trailing sign, a sign of CR or DB, and a comma as a 1000s separator. Leading and trailing spaces, asterisks, and dollar signs are ignored. A numeric string may not contain more than one sign, more than one decimal point, or a misplaced comma separator. A null string or a string containing all spaces is interpreted as zero. This function would normally be used in a conditional IF statement.


**Usage:** 

	ISNUMERIC(string) -> boolean
	string is the string to be tested for a numeric value.

**Examples:**

	Expression                   	Result
	ISNUMERIC("25") 		True 
	ISNUMERIC("    -16    ") 	True 
	ISNUMERIC("1.64") 		True 
	ISNUMERIC("  ") 		True 
	ISNUMERIC("1.64E+04") 		True 
	ISNUMERIC("-16.4E-4") 		True 
	ISNUMERIC("-    21") 		True 
	ISNUMERIC("-000021") 		True 
	ISNUMERIC("12,345,678") 	True 
	ISNUMERIC("12,345.678") 	True 
	ISNUMERIC("$1.64") 		True 
	ISNUMERIC("-$1.64") 		True 
	ISNUMERIC("44 CR") 		True 
	ISNUMERIC("55 DB") 		True 
	ISNUMERIC("1,2345.67") 		False 
	ISNUMERIC("1.64E+04.2") 	False 
	ISNUMERIC("44 AB") 		False
	ISNUMERIC("-44 CR") 		False 
	ISNUMERIC("-44-") 		False

### MOVE
Fundamentally the MOVE command will move data from one place to another. The source can be a literal or a variable, but the destination must be a variable that can accomodate what is being moved, ie., you cannot move a text string to a numeric variable type without throwing a runtime error. A useful aspect of the MOVE command is that it can operate on a single element of a data record, or the entire record itself. Many of the listed functions in this section will only work in concert with the MOVE command. 

**Usage:**

	MOVE data-element1 TO data-element2
		or
	MOVE "Gobol" TO data-element3
	
*data-element1* is the sending variable and *data-element2* is the receiving variable. The literal "Gobol" is being assigned to the alphabetic variable type *data-element3*.

#### CORRESPONDING
CORRESPONDING is a modifier for the MOVE command, MOVE CORRESPONDING will assign values between record structures to elements with identical names. It will not affect the elementary items in the sending or receiving structures that are not an exact name match.

If the sending field length is less than the receiving field length, it will be padded with spaces or zeroes depending on the data type.

If the sending field length is greater than the receiving field length, the data will be truncated based on the length.

Sending alphabetic data to a numeric field will throw a run time error.

Sending numeric data to an alphabetic field will typecast the data from numeric to alphabetic.

**Usage:**

	MOVE CORRESPONDING sending-record to receiving-record.

**Example:**
	
	The following examples assume the script:
	DEFINE sending-record : RECORD 
		F : CHAR(4) 
		G : CHAR(6)
		H : CHAR(2) 
	END
	DEFINE receiving-record : RECORD
		A : NUMERIC(10)
		F : CHAR(6)
		B : DATE
		G : CHAR(6)
		H : CHAR(2)
	END
	
	MOVE "ABCD"   TO sending-record.F
	MOVE "efghij" TO sending-record.G
	MOVE “13”     TO sending-record.H
	MOVE CORRESPONDING sending-record TO receiving-record
	
The result is that *receiving-record* will have the same values in elements F, G and H that *sending_record* has, and whatever data was in A and B will remain untouch. Since F is larger in the receiving field, it will simply be padded with blanks.
	
### ROUND 	
Rounds a number to a specified number of digits.

**Usage:** 

	ROUND(number, num-digits) -> number
	
*number* is the number to be rounded.
*num-digits* indicates the number of decimal digits to the right of the decimal to which number is to be rounded. If num-digits is equal to zero, number is rounded to the nearest integer. If num-digits is negative, number is rounded to the power of 10 indicated by negative num-digits.

**Examples:**
 
	Expression		Result
	ROUND(3.14159, 4) 	3.1416 
	ROUND(3.14159, 2)	3.14 
	ROUND(3.14159, 6)	3.141590 
	ROUND(-3.14159, 4) 	-3.1416 
	ROUND(7419.917, 0) 	7420 
	ROUND(7419.917, -1) 	7420 
	ROUND(7419.917, -2) 	7400 
	ROUND(7419.917, -4) 	10000

### STR2DATE 	

Converts a string to a date, time, datetime, or interval using a specified date format, or automatically recognizes a date if no format is specified. STR2DATE is the inverse function of DATE2STR (which converts a date to a string).

**Usage:** 	

	STR2DATE(date-str) -> date
		or
	STR2DATE(date-str, format-str) -> date
	
*date-str* is a string containing the date to be converted to a date type item.

*format-str* is optional and if present contains a string of tokens that describe the format of date-str.  The tokens are the same as those used in the print PIC of a date item and the DATE2STR function.  

The allowable tokens in format-str are:


	A.M. 	AM/PM indicator with periods 
	AM 	AM/PM indicator 
	CC 	2 digit century 
	D 	The day of week (1-7, Sun=1,Sat=7) 
	DAY 	The 9 character name of day of week (SUNDAY-SATURDAY) 
	DD 	The 2 digit day number within month (01-31) 
	D* 	The 1 or 2 digit day number within month (1-31) 
	DY 	The 3 character name of day of week (SUN-SAT) 
	HH12 	The 2 digit hour in 12 hour time (0112) 
	HH24 	The 2 digit hour in 24 hour time (0023) 
	H*12 	The 1 or 2 digit hour in 12 hour time (1-12)
	H*24 	The 1 or 2 digit hour in 24 hour time (0-23)
	MI 	The 2 digit minute within the hour (00-59) 
	MM 	The 2 digit month number within year (01-12) 
	M* 	The 1 or 2 digit month number within year (1-12) 
	MON 	The 3 character name of month (JAN-DEC) 
	MONTH 	The 9 character name of month (JANUARY-DECEMBER) 
	P.M. 	AM/PM indicator with periods 
	PM 	AM/PM indicator 
	SS 	The 2 digit second within the minute (00-59) 
	SSSSS 	The 5 digit second within the day (086399) 
	W 	The week within the month (1-5) 
	WW 	The 2 digit week within the year (0153) 
	YY 	The last 2 digits of the year.  If no century is specified (CC), the current century is assumed.
	YYYY 	The 4 digit year

	space 	Space 
	: 	Colon 
	/ 	Forward Slash 
	- 	Hyphen 
	. 	Period 
	, 	Comma 
	; 	Semicolon 
	"str" 	Quoted string


Date format items are case insensitive in the STR2DATE function.

Items are required to be the specified length. For example the MONTH token requires a 9 character month name, so APRIL must have 4 trailing  Spaces. 

If items are missing, default values are used as follows:

	Default century: 	the current century. 
	Default year: 		the current year. 
	Default month: 		1 (January) 
	Default day: 		1 
	Default hour: 		0 
	Default minute: 	0 
	Default second: 	0

If *format-str* is absent, STR2DATE examines *date-str* and converts it to a date if it is in a recognizable format. It is an error if the date-str is not in a recognizable format. The ISDATE function can be used to determine if the string is in a recognizable format. STR2DATE can automatically recognize dates, times, datetimes, and intervals.  If a string can be either a date or a time, then a date is recognized.  e.g. "121314" is interpreted as 13-Dec-2014, not 12:13:14 PM. A two digit year from 00 to 49 is interpreted to be from 2000 to 2049, and a two digit year from 50 to 99 is interpreted to be from 1950 to 1999. 

Recognizable formats are:

	Date formats: 
	MM/DD/YY 
	MM/DD/YYYY 
	MM/DD 
	MM-DD-YY 
	MM-DD-YYYY 
	MM-DD 
	DD-MON-YY
	DD-MON-YYYY 
	DD-MON 
	YYMMDD 
	YYYYMMDD 
	Month DD, YY 
	Month DD, YYYY 
	Month DD

	Time formats: 
	HH24:MI 
	HH24:MI:SS 	AM, PM, A.M. or P.M. may follow 

	Datetime formats: 
	date time 		Date in one of the above formats followed by a time in one of the above formats.


**Examples:**
 
	Expression                   		Result
	STR2DATE("180801","YYMMDD") 		01-AUG-2018
	STR2DATE("121224","HHMISS") 		12:12:24 PM 
	STR2DATE("Sep  8,1992","MON DD,YYYY") 	08-SEP-1992 
	STR2DATE("180704") 			04-JUL-2018 
	STR2DATE("11/12/24") 			12-DEC-2024 
	STR2DATE("12/12/1924") 			12-DEC-1924 
	STR2DATE("121224") 			12-DEC-2024 
	STR2DATE("12121964") 			12-DEC-1964
	STR2DATE("12-Dec-24") 			12-DEC-2024 
	STR2DATE("  130401") 			01-APR-2013 
	STR2DATE("12-12-24") 			12-DEC-2024 
	STR2DATE("12:12:24") 			12:12:24 PM 
	STR2DATE("12") 				12 Days 
	STR2DATE("12 4:14") 			12 Days 4 Hours 14 Minutes 
	STR2DATE("9/8/14 17:21") 		08-SEP-2014 5:21 PM 
	STR2DATE("May") 			* Error

### UPSHIFT	
UPSHIFT converts all lowercase characters in a given string to uppercase characters. Can be used on a variable in place, or in conjunction to the MOVE verb to modify the target value to be shifted. Default behavior is to upshift all characters in the String, but the parameters EACH or FIRST modify the behavior. EACH will upshift the first character of each string after a whitespace and the First character. Specifying FIRST will upshift just the first character.

**Usage:** 	

	MOVE UPSHIFT(string) TO VAR
		or
	MOVE UPSHIFT(string,FIRST) TO VAR

**Examples:**
 		
	Expression						Result
	MOVE "Gobol" to VAR									
	UPSHIFT(VAR)						"GOBOL"
	MOVE “george allan smith” 	TO VAR			“george allan smith”
	MOVE UPSHIFT(“george allan smith”) TO VAR		“GEORGE ALLAN SMITH”
	MOVE UPSHIFT(“george allan smith”,EACH) TO VAR		“George Allan Smith”
	MOVE UPSHIFT(“george allan smith”,FIRST) TO VAR		“George allan smith”

## FILE TYPES

	CSV 
	[DELIM=delimiter-character]
	[QUOTE=quotation-character]
	[ESCAPE=escape-character]
	[STRIP=strip-character]

*delimiter-character* is a single character that specifies how the fields are to be delimited in the output. If the delimiter is a space, it needs to be enclosed in quotation marks. The default delimiter-character is a comma (,). quotation-character is a single character that specifies how fields that contain the delimiter-character or quote-character are grouped into a single field. When a field that contains either the delimiter-character or the quote-character is written, it is enclosed in the quote-character. The default quote-character is a double quote (").

*escape-character* determines how the quote-character is handled in a quoted field. Fields that contain the quote-character are output surrounded by the quote-character and each time the quote-character actually occurs in the field, it is preceded by the escape-character.
The default escape-character is a double quote (").

*strip-character* indicates a character that is stripped at the beginning and end of each field. The default strip-character is a space. To indicate no characters are stripped, use STRIP="".


	FIXED <fixed width fields>
	SQL
	NOSQL
	JSON
	BSON

### RANDOM DESIGN NOTES, SOME OF WHICH ARE ABOVE

	RECORD MYTHING
		// GOBOL NAME, TYPE, SOURCE INDEX
		ID	INT		“id”
		NAME	STRING	“name:first”
	END
	DEFINE MYFILE “c:\stuff.csv”

	DEFINE MYRECORD : FORMAT MYTHING
	PERFORM READ MYFILE INTO MYRECORD UNTIL EOF
		DISPLAY MYRECORD.ID
	END-PERFORM

What about implementing threading so that you could read multiple input files into memory at the same time and then work on them there?


UNSTRING
Ex.	unstring input-buffer delimited by “,” into record-structure

INITIALIZE record-structure/variable //sets strings to spaces and numbers to zero

IMPORT <file-name> [CSV | JSON | database connection string/table ]
// this is part of the commercial module that would be in the Orsini GUI that will generate your record structure layout for you. In the command line version it would pull in the file, in the case of a CSV it assumes that line 1 has headers, if it doesn’t have headers, then you are screwed on that. The interpreter should have a basic editor in it or a plug in for someone other editor, but it needs to be able to save.


INSPECT <string> CONVERTING <value> TO <value> // this will allow you to change the length of things, so if you want to turn “ABC” to nothing you could do INSPECT MY-VAR CONVERTING “ABC” TO “”. This would not be a null. It is a way to strip things out as well. What about keywords like SYMBOLS, ALPHABETIC, ALPHANUMERIC and NUMERIC, for example if you want a phone number stripped of any symbols you could convert SYMBOLS to “”.

Make use of dataframes with this golang lib https://godoc.org/github.com/akualab/dataframe


Functions
A function can be part of a MOVE statement or stand-alone on a field that already contains the value.

MOVE STR2DATE(DATE-STRING) TO DATE-FIELD // this is a SQL DB date type field
DATE2STR // same thing about date type
RTRIM // remove trailing spaces
LTRIM // remove leading spaces
TRIM // remove leading and trailing spaces

SORT // apply to tables and files, look at cobol layout)
Move function(plot(<defined file.record.field>,<array>)) to kounter //will read the defined file and total all unique values for field in record and populate an array with field value and count of values.

MOVE UPSHIFT(STRING[All,First,Each]) TO NEW-FIELD // will upshift entire string by default or ‘all’ will upshift the first character of each word or ‘first’ will upshift just the first letter in the string. If NEW-FIELD is the same as STRING, then it will upshift the variable in place.

MOVE UPSHIFT(LTRIM(STRING)) TO NEW-FIELD

MOVE UNSTRING MY-STRING DELIMITED BY “,” TO VAR1

MOVE STR2DATE(VAR(DATE FORMAT) TO DATE-VAR // this lets you specify the date format in the string that will then be converted to a date type, so like YYYYMMDD or maybe it is “Month, day  year”.


MOVE EXPORT(ARRAY, <DELIMITER>, [NEW, APPEND OVERWRITE]) TO OUTPUT-FILE // will take the ARRAY and write it out one record at a time to OUTPUT-FILE, the optional flags will create, overwrite or append the file. Default behavior is NEW. A delimiter character will insert that between each field, this is useful for an export. Maybe an array should apply to a database table, which would then require selection criteria, need to think about how to lump functions together. Maybe the “from” source can even be the file/record structure.

MOVE CORRESPONDING RECORD1 TO RECORD2 //this will match the field names up between the two record structures and copy all the fields from RECORD1 that have the same names in RECORD2

## Some Pseudo Code
This example reads a tab delimited file, standardizes some input fields and writes it to a SQL type database. The database is not defined in this sample.

	BEGIN
	RECORD VOTER_ROLLS BEGIN
		lVoterUniqueID
		sAffNumber
		szStateVoterID
		sVoterTitle
		szNameLast
		szNameFirst
		szNameMiddle
		sNameSuffix
		szSitusAddress
		szSitusCity
		sSitusState
		sSitusZip
		sHouseNum
		sUnitAbbr
		sUnitNum
		szStreetName
		sStreetSuffix
		szMailAddress1
		szMailAddress2
		szMailAddress3
		szMailAddress4
		szMailZip
		szPhone
		szEmailAddress
		dtBirthDate
		sBirthPlace
		dtRegDate
		dtOrigRegDate
		dtLastUpdate_dt
		sStatusCode
		szStatusReasonDesc
		sUserCode1
		sUserCode2
		szPartyName
		szAVStatusAbbr
		szAVStatusDesc
		szPrecinctName
		sPrecinctID
		sPrecinctPortion
		sDistrictID_0
		iSubDistrict_0
		szDistrictName_0
	END-RECORD

	OPEN INPUT VR TAB VOTER_ROLLS	// this opens an input file that is TAB delimited and uses the VOTER_ROLLS record definition for a record structure called VR.
	OPEN OUTPUT table		// this is the output database table, which isn’t designed yet
	PERFORM READ VR UNTIL EOF
		INITIALIZE table.RECORD	// set all fields to default values in the database table
		move corresponding VR.RECORD  					to table.RECORD		// this will match up the field names between the source and destination and copy all the fields
		move STR2DATE(VR.dtBirthDate, "MM/DD/YYYY") 	to table.BirthDate	// this takes a string date and format information and casts it to a database date type
		move STR2DATE(VR.dtRegDate, "MM/DD/YYYY") 		to table.RegDate
		move STR2DATE(VR.dtOrigRegDate, "MM/DD/YYYY") 	to table.OrigRegDate
		move STR2DATE(VR.dtLastUpdateDt, "MM/DD/YYYY") 	to table.LastUpdateDt

		UPSHIFT(DOWNSHIFT(table.szNameLast),FIRST)		// downshift entire string and then upshift the first character of the string, the FIRST is optional, default is ALL
		UPSHIFT(DOWNSHIFT(table.szNameFirst),FIRST)
		UPSHIFT(DOWNSHIFT(table.szNameMiddle),FIRST)
		UPSHIFT(DOWNSHIFT(table.szMailAddress1),EACH)	// downshift entire string, then upshift the first letter of EACH word in the string 
		UPSHIFT(table.sSitusState)
		DOWNSHIFT(table.szEmailAddress)	// standardize emails as lowercase
		INSPECT table.szPhone CONVERTING SYMBOLS TO "" 	// we want to remove all non numeric symbols from the phone number
		WRITE table.RECORD 
	END-PERFORM

	stop run

