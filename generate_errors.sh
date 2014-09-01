#!/bin/bash

FILENAME=errors.go

echo "package wu" > "$FILENAME"
echo "" >> "$FILENAME"
echo "const (" >> "$FILENAME"
while read line; do
	if [[ $line == WU_* ]]; then
		VARNAME=$line
		CODE=""
		DESC=""
	elif [[ $line == 0x* ]]; then
		CODE=$line
		DESC=""
	else
		DESC=$line
	fi
	if [[ ! -z $VARNAME ]] && [[ ! -z $CODE ]] && [[ ! -z $DESC ]]; then
		echo -e "\t${VARNAME}      = $CODE" >> "$FILENAME"
		echo -e "\t${VARNAME}_DESC = \"$DESC\"" >> "$FILENAME"
	fi
done < errors.txt
echo ")" >> "$FILENAME"



echo "" >> "$FILENAME"
echo "// REFERENCES:" >> "$FILENAME"
echo "// http://msdn.microsoft.com/en-us/library/windows/desktop/hh968413(v=vs.85).aspx (success/error codes)" >> "$FILENAME"
echo "// http://msdn.microsoft.com/en-us/library/windows/desktop/aa387293(v=vs.85).aspx (network error codes)" >> "$FILENAME"

go fmt "$FILENAME"
