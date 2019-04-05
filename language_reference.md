# GOBOL Language Reference
---------------------------------------
  * [Looping](#looping)
    * [PERFORM](#perform)
  * [Conditionals](#conditionals)
    * [EVALUATE](#evaluate)
    * [IF..ELSE..END-IF](#if-else-end-if)
  * [Functions](#functions)
    * [ACCEPT](#accept)
  * [Date Operators](#date-operators)

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
 
		FROM 		{ identifier-3 }
                   	{ index-name-3 }
                   	{ literal-3    }

			[ BY 	{ identifier-4 } ]
                   	{ literal-4    }
		UNTIL condition-1 ] ...

		read-phrase.
			PERFORM READ <file> into <record structure> [conditional tests as described above] UNTIL EOF
	// Maybe use this to perform a function in the code, have to think about how functions would be defined and variables passed or just make all variables global.



## Conditionals

### EVALUATE
The EVALUATE statement causes multiple conditions to be evaluated. The subsequent action of the program depends on the results of these evaluations.
 
The EVALUATE statement is very similar to the CASE construct common in many other programming languages. The EVALUATE/CASE construct provides the ability to selectively execute one of a set of instruction alternatives based on the evaluation of a set of choice alternatives.
 
EVALUATE extends the power of the typical CASE construct by allowing multiple data items and conditions to be named in the EVALUATE phrase

	
	EVALUATE  {subject}  [ ALSO {subject} ] ...  
             	TRUE   }         {TRUE   }  
            	 FALSE  }         {FALSE  }  
    
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

## Functions
GOBOL has a number of built-in functions for date, string and file management to make coding faster. They are as follows:

### ACCEPT 
Reads a string from the user’s standard input (stdin) after displaying a prompt string.

Usage ACCEPT(prompt-string) -> string

**Example:**

		MOVE ACCEPT("Enter your name: ") TO FIRST-NAME.



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
------|----_--|---------- 
DATETIME | DATETIME | Yes 
DATETIME 		DATE 			Yes 
DATETIME 		TIME 			No 
DATETIME 		INTERVAL 		No 
DATE 			DATETIME 		Yes 
DATE 			DATE 			Yes 
DATE 			TIME 			No 
DATE 			INTERVAL 		No 
TIME 			DATETIME 		No 
TIME 			DATE 			No 
TIME 			TIME 			Yes 
TIME 			INTERVAL 		Yes 
INTERVAL 		DATETIME 		No 
INTERVAL 		DATE 			No 
INTERVAL 		TIME 			Yes 
INTERVAL 		INTERVAL 		Yes

