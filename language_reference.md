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



## Conditionals

### EVALUATE

### IF..ELSE

## Date Functions

=======
  * [Date Functions](#date-functions)
>>>>>>> master
