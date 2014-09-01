#!/bin/bash

FILENAME=errors.go

cat "$FILENAME.head" > "$FILENAME"

echo "" >> "$FILENAME"

echo "const (" >> "$FILENAME"
while read line; do
	if [[ $line == WU_* ]]; then
		VARNAME=$line
		CODE=""
	elif [[ $line == 0x* ]]; then
		CODE=$line
	fi
	if [[ ! -z $VARNAME ]] && [[ ! -z $CODE ]]; then
		echo -e "\t$VARNAME WUError = $CODE" >> "$FILENAME"
		VARNAME=""
		CODE=""
	fi
done < errors.txt
echo ")" >> "$FILENAME"

echo "" >> "$FILENAME"

echo "func ErrorDesc(errCode WUError) string {" >> "$FILENAME"
echo -e "\tswitch errCode {" >> "$FILENAME"
while read line; do
	if [[ $line == WU_* ]]; then
		VARNAME=$line
		DESC=""
	elif [[ $line == 0x* ]]; then
		DESC=""
	else
		DESC=$line
	fi
	if [[ ! -z $VARNAME ]] && [[ ! -z $DESC ]]; then
		echo -e "\tcase $VARNAME:" >> "$FILENAME"
		echo -e "\t\treturn \`$DESC\`" >> "$FILENAME"
		VARNAME=""
		DESC=""
	fi
done < errors.txt
echo -e "\tdefault: return \`Unknown error.\`" >> "$FILENAME"
echo -e "\t}" >> "$FILENAME"
echo "}" >> "$FILENAME"

echo "" >> "$FILENAME"

echo "func ErrorName(errCode WUError) string {" >> "$FILENAME"
echo -e "\tswitch errCode {" >> "$FILENAME"
while read line; do
	if [[ $line == WU_* ]]; then
		VARNAME=$line
	fi
	if [[ ! -z $VARNAME ]]; then
		echo -e "\tcase $VARNAME:" >> "$FILENAME"
		echo -e "\t\treturn \`$VARNAME\`" >> "$FILENAME"
		VARNAME=""
	fi
done < errors.txt
echo -e "\tdefault: return \`\`" >> "$FILENAME"
echo -e "\t}" >> "$FILENAME"
echo "}" >> "$FILENAME"

echo "" >> "$FILENAME"

cat "$FILENAME.tail" >> "$FILENAME"

echo "" >> "$FILENAME"
echo "// REFERENCES:" >> "$FILENAME"
echo "// http://msdn.microsoft.com/en-us/library/windows/desktop/hh968413(v=vs.85).aspx (success/error codes)" >> "$FILENAME"
echo "// http://msdn.microsoft.com/en-us/library/windows/desktop/aa387293(v=vs.85).aspx (network error codes)" >> "$FILENAME"

go fmt "$FILENAME"
