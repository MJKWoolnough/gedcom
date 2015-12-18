#!/bin/bash

(
	cat <<HEREDOC
package gedcom

// File automatically generated with ./genTypes.sh

import (
	"strconv"
	"strings"
)
HEREDOC


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
		echo "// $eType is a GEDCOM base type";
		echo "type $eType $vType";
		echo;
		echo "func (e *$eType) parse(l Line) error {";
		if [ -z "${data[1]}" ]; then
			if [ -z "$(echo "${data[2]}" | tr -d "[:upper:]")" ]; then
				echo "	switch strings.ToUpper(l.value) {"
			elif [ -z "$(echo "${data[2]}" | tr -d "[:lower:]")" ]; then
				echo "	switch strings.ToLower(l.value) {"
			else
				echo "	switch l.value {"
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

	cat <<HEREDOC

// ErrInvalidValue is an error that is generated when a type is not one of the
// specified values
type ErrInvalidValue struct {
	Type, Value string
}

// Error is an implementation of the error interface
func (e ErrInvalidValue) Error() string {
	return "Value for " + e.Type + " is invalid"
}

// ErrInvalidLength is an error that is generated when a type is given more or
// less data than is required
type ErrInvalidLength struct {
	Type, Value string
	Min, Max    uint
}

// Error is an implementation of the error interface
func (e ErrInvalidLength) Error() string {
	return "Value for " + e.Type + " has an invalid length"
}
HEREDOC
) > types.go
