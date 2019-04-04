# GOBOL Language Reference
---------------------------------------
  * [Looping](#looping)
    * [PERFORM](#perform)
  * [Conditionals](#conditionals)
    * [EVALUATE](#evaluate)
    * [IF..ELSE](#if-else)
  *[Date Functions](#date-functions)

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

####Syntax Rules
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


### IF..ELSE

## Date Functions


