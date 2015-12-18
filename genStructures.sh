#!/bin/bash

function processStructure {
	local structureName="$1";
	local longest=$2;
	shift;
	shift;
	local types=();
	local ID="";
	local lineValue=();
	local required=();
	local oneMost=();
	local maxes=();
	echo;
	echo "// $structureName is a GEDCOM structure type";
	echo "type $structureName struct {";
	for type;do
		IFS=":";
		parts=($type);
		IFS="$OFS";
		pTag="${parts[0]}";
		pType="${parts[1]}";
		pName="${parts[2]}";
		pMin="${parts[3]}";
		pMax="${parts[4]}";
		if [ -z "$pMin" ]; then
			pMin="1";
		fi;
		if [ -z "$pMax" ]; then
			pMax="1";
		fi;
		if [ -z "$pName" ]; then
			pName="$pType";
		fi;
		if [ "${parts[0]:0:1}" = "@" ]; then
			pType="Xref"
		fi;
		echo -n "	$pName ";
		for i in $(seq $(( longest - ${#pName} ))); do
			echo -n " ";
		done;
		if [ "$pMax" != "1" ]; then
			echo -n "[]";
		fi;
		echo "$pType";
		if [ "${pTag:0:1}" = "@" ]; then
			ID="$pName";
		elif [ "${pTag:0:1}" = "#" ]; then
			extends="false"
			if [ "${pTag:1:2}" = "*" ]; then
				extends="true";
			fi;
			lineValue=( "$pType" "$pName" extends );
		else
			if [ "$pMin" = "1" ]; then
				required+=( "$pName" );
			elif [ "$pMax" == "1" ]; then
				oneMost+=( "$pName" );
			elif [ "$pMax" != "M" ]; then
				maxes+=( "$pName" );
			fi;
			types+=("$pTag:$pType:$pName:$pMin:$pMax");
		fi;
	done;
	echo "}";
	echo;
	echo "func (s *$structureName) parse(l Line) error {";
	if [ ! -z "$ID" ]; then
		echo "	if err := s.${ID}.parse(Line{value: l.xrefID}); err != nil {";
		echo "		return ErrContext{\"$structureName\", \"xrefID\", err}";
		echo "	}";
	fi;
	if [ ${#lineValue[@]} -gt 0 ]; then
		echo "	if err := s.${lineValue[1]}.parse(l); err != nil {";
		echo "		return ErrContext{\"$structureName\", \"line_value\", err}";
		echo "	}";
		if [ ${lineValue[2]} = "true" ]; then
			echo "	for i := 0; i < len(l.Sub); i++ {";
			echo "		switch l.Sub[i].tag {";
			echo "		case \"CONT\":";
			echo "			s.${lineValue[1]} += \"\n\"";
			echo "			fallthrough";
			echo "		case \"CONC\":";
			echo "			var t ${lineValue[0]}";
			echo "			if err := t.parse(l.Sub[i]); err != nil {";
			echo "				return ErrContext{\"\$structureName\", \"line_value\", err}";
			echo "			}";
			echo "			s.${lineValue[1]} += t";
			echo "			l.Sub = append(l.Sub[:i], l.Sub[i+1]...)";
			echo "			i--";
			echo "		}";
			echo "	}";
		fi;
	fi;
	if [ ${#types[@]} -gt 0 ]; then
		if [ ${#required[@]} -gt 0 -o ${#oneMost[@]} -gt 0 ]; then
			echo -n "	var";
			local c=false;
			if [ ${#required[@]} -gt 0 ]; then
				for r in "${required[@]}"; do
					if $c; then
						echo -n ","
					fi;
					echo -n " ${r}Set";
					c=true;
				done;
			fi;
			if [ ${#oneMost[@]} -gt 0 ]; then
				for o in "${oneMost[@]}"; do
					if $c; then
						echo -n ","
					fi;
					echo -n " ${o}Set";
					c=true;
				done;
			fi;
			echo " bool";
		fi;
		if [ ${#maxes[@]} -gt 0 ]; then
			echo -n "	var";
			local c=false;
			for m in "${maxes[@]}"; do
				if $c; then
					echo -n ","
				fi;
				echo -n " ${m}Count";
				c=true;
			done;
			echo " int";
		fi;
		echo "	for _, sl := range l.Sub {"
		echo "		switch sl.tag {";
		for type in ${types[@]}; do
			IFS=":";
			parts=($type);
			IFS="$OFS";
			pTag="${parts[0]}";
			pType="${parts[1]}";
			pName="${parts[2]}";
			pMax="${parts[4]}";
			cont=false
			if [ "${pTag: -1}" = "*" ]; then
				pTag="${pTag:0:-1}";
				cont=true;
			fi;
			echo "		case \"$pTag\":";
			if [ "$pMax" = "1" ]; then
				echo "			if ${pName}Set {";
				echo "				return ErrContext{\"$structureName\", \"$pTag\", ErrSingleMultiple}";
				echo "			}";
				echo "			${pName}Set = true";
			elif [ "$pMax" != "M" ]; then
				echo "			if ${pName}Count == $pMax {";
				echo "				return ErrContext{\"$structureName\", \"$pTag\", ErrTooMany($pMax)}";
				echo "			}";
				echo "			${pName}Count++";
			fi;
			if [ "$pMax" = "1" ]; then
				echo "			if err := s.${pName}.parse(sl); err != nil {";
			else
				echo "			var t ${pType}";
				echo "			if err != t.parse(sl); err != nil {";
			fi;
			echo "				return ErrContext{\"$structureName\", \"$pTag\", err}";
			echo "			}";
			if [ "$pMax" != "1" ]; then
				echo "			s.${pName} = append(s.${pName}, t)";
			fi;
		done;
		echo "		default:";
		echo "			if len(sl.tag) < 1 || sl.tag[0] != '_' {";
		echo "				return ErrContext{\"$structureName\", sl.tag, ErrUnknownTag}";
		echo "			}";
		echo "			// possibly store in a Other field";
		echo "		}";
		echo "	}";
		if [ ${#required} -gt 0 ]; then
			for r in "${required[@]}"; do
				echo "	if !${r}Set {";
				echo "		return ErrContext{\"$structureName\", \"$r\", ErrRequiredMissing}";
				echo "	}";
			done;
		fi;
	else
		echo "	for _, sl := range l.Sub {"
		echo "		if len(sl.tag) < 1 || sl.tag[0] != '_' {";
		echo "			return ErrContext{\"$structureName\", sl.tag, ErrUnknownTag}";
		echo "		}";
		echo "		// possibly store in a Other field";
		echo "	}";
	fi;
	echo "	return nil";
	echo "}";
}

OFS="$IFS";

(
	echo "package gedcom";
	echo;
	echo "// File automatically generated with ./genStructures.sh";
	echo;
	echo "import \"errors\"";

	(
		read structureName;
		types=();
		longest=0;
		IFS="$(echo)";
		while read line;do 
			if [ "${line:0:1}" = "	" ]; then
				pType="$(echo "$line" | cut -d':' -f2)";
				pName="$(echo "$line" | cut -d':' -f3)";
				if [ -z "$pName" ]; then
					pName="$pType";
				fi;
				if [ ${#pName} -gt $longest ]; then
					longest=${#pName};
				fi;
				types+=("${line:1}");
			else
				IFS="$OFS";
				processStructure "$structureName" $longest "${types[@]}";
				IFS="$(echo)";
				longest=0;
				types=();
				structureName="$line";
			fi;
		done
		IFS="$OFS";
		processStructure "$structureName" $longest "${types[@]}";
	) < structures.gen
	
	echo;
	echo "// Errors";
	echo "var (";
	echo "	ErrRequiredMissing = errors.New(\"required tag missing\")";
	echo "	ErrSingleMultiple  = errors.New(\"tag was specified more than the one time allowed\")";
	echo "	ErrUnknownTag      = errors.New(\"unknown tag\")";
	echo ")";
	echo;
	echo "// ErrContext adds context to a returned error";
	echo "type ErrContext struct {"
	echo "	Structure, Tag string";
	echo "	Err            error";
	echo "}";
	echo;
	echo "// Error implements the error interface";
	echo "func (e ErrContext) Error() string {";
	echo "	return e.Tag + \":\" + t.Err.Error()";
	echo "}";
	echo;
	echo "// ErrTooMany is an error returned when too many of a particular tag exist";
	echo "type ErrTooMany int";
	echo;
	echo "// Error implements the error interface";
	echo "func (ErrTooMany) Error() string {";
	echo "	return \"too many tags\"";
	echo "}";

) > structures.go
