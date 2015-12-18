#!/bin/bash

(

	echo "package gedcom"
	echo;
	echo "// File automatically generated with ./genMap.sh";
	echo;
	echo "import (";
	echo "	\"strconv\"";
	echo "	\"strings\"";
	echo ")";


	while read line; do
		OFS="$IFS";
		IFS=":"
		data=($line)
		IFS="$OFS";
		eType="${data[0]}";
		vType="string";
		first="${data[1]}";
		if [ -z "$first" ]; then
			first="${data[2]}";
			if [ "${first:0:1}" == "(" ]; then
				vType="$(echo "${first:1}" | cut -d')' -f1)";
				data[2]="$(echo "$first" | cut -d')' -f2)";
			fi;
		elif [ "${first:0:1}" == "(" ]; then
			vType="$(echo "${first:1}" | cut -d')' -f1)";
			data[1]="$(echo "$first" | cut -d')' -f2)";
		fi;
		echo;
		echo "type $eType $vType";
		echo;
		echo "func (e *$eType) parse(l Line) error {";
		if [ -z "${data[1]}" ]; then
			if [ -z "$(echo "${data[2]}" | tr -d "[:lower:]")" ]; then
				echo "	switch strings.ToLower(l.value) {"
			else
				echo "	switch strings.ToUpper(l.value) {"
			fi;
			for i in $(seq 2 $(( ${#data[@]} - 1 ))); do
				echo "	case \"${data[$i]}\":";
				if [ "$vType" = "string" ]; then
					echo "		*e = \"${data[$i]}\"";
				else
					echo "		*e = ${data[$i]}";
				fi;
			done;
			echo "	default:";
			echo "		return ErrInvalidValue{\"$eType\", l.value}";
			echo "	}";
			echo "	return nil";
		else
			echo "	if len(l.value) < ${data[1]} || len(l.value) > ${data[2]} {"
			echo "		return ErrInvalidLength{\"$eType\", l.value, ${data[1]}, ${data[2]}}";
			echo "	}";
			if [ "$vType" = "string" ]; then
				echo "	*e = $eType(l.value)";
			else
				num="$(echo "$vType" | tr -d "[:alpha:]")";
				if [ -z "$num" ]; then
					num="0"
				fi;
				echo "	n, err := strconv.ParseUint(l.value, 10, $num)";
				echo "	if err != nil {";
				echo "		return err";
				echo "	}";
				echo "	*e = $eType(n)";
			fi;
			echo "	return nil";
		fi;
		echo "}";
	done < types.gen

	echo;
	echo "type ErrInvalidValue struct {";
	echo "	Type, Value string";
	echo "}";
	echo;
	echo "func (e ErrInvalidValue) Error() string {";
	echo "	return \"Value for \" + e.Type + \" is invalid\"";
	echo "}";
	echo;
	echo "type ErrInvalidLength struct {"
	echo "	Type, Value string";
	echo "	Min, Max    uint";
	echo "}";
	echo;
	echo "func (e ErrInvalidLength) Error() string {";
	echo "	return \"Value for \" + e.Type + \" has an invalid length\"";
	echo "}";
) > types.go
