#!/bin/bash

function processStructure() {
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
	local namedTags=0;
	local embedded="";
	echo;
	echo "// $structureName is a GEDCOM structure type.";
	echo "type $structureName struct {";
	for type; do
		IFS=":";
		parts=( $type );
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
			pType="Xref";
		fi;
		echo -n "	";
		if [ ! -z "$pTag" ]; then
			echo -n "$pName ";
			for i in $(seq $(( longest - ${#pName} ))); do
				echo -n " ";
			done;
			if [ "$pMax" != "1" ]; then
				echo -n "[]";
			fi;
		fi;
		echo "$pType";
		if [ "${pTag:0:1}" = "@" ]; then
			ID="$pName";
		elif [ "${pTag:0:1}" = "#" ]; then
			lineValue=( "$pType" "$pName" );
		elif [ -z "$pTag" ]; then
			embedded="$pType";
			types+=( "$pTag:$pType:$pName:$pMin:$pMax" );
		else
			if [ "$pMin" = "1" ]; then
				required+=( "$pName" );
			elif [ "$pMax" == "1" ]; then
				oneMost+=( "$pName" );
			elif [ "$pMax" != "M" ]; then
				maxes+=( "$pType:$pName:$pMax" );
			fi;
			let "namedTags++";
			types+=( "$pTag:$pType:$pName:$pMin:$pMax" );
		fi;
	done;
	echo "}";
	echo;
	echo "func (s *$structureName) parse(l *Line, o options) error {";
	if [ ! -z "$ID" ]; then
		echo "	if err := s.${ID}.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {";
		echo "		return ErrContext{\"$structureName\", \"xrefID\", err}";
		echo "	}";
		echo "";
	fi;
	if [ ${#lineValue[@]} -gt 0 ]; then
		echo "	if err := s.${lineValue[1]}.parse(l, o); err != nil {";
		echo "		return ErrContext{\"$structureName\", \"line_value\", err}";
		echo "	}";
		echo "";
	fi;
	if [ ${#types[@]} -gt 0 ]; then
		if [ ${#required[@]} -gt 0 -o ${#oneMost[@]} -gt 0 ]; then
			echo -n "	var";
			local c=false;
			if [ ${#required[@]} -gt 0 ]; then
				for r in "${required[@]}"; do
					if $c; then
						echo -n ",";
					fi;
					echo -n " ${r}Set";
					c=true;
				done;
			fi;
			if [ ${#oneMost[@]} -gt 0 ]; then
				for o in "${oneMost[@]}"; do
					if $c; then
						echo -n ",";
					fi;
					echo -n " ${o}Set";
					c=true;
				done;
			fi;
			echo " bool";
			echo "";
		fi;
		if [ $namedTags -gt 0 ]; then
			if [ ${#maxes[@]} -gt 0 ]; then
				for m in "${maxes[@]}"; do
					pType="$(echo "$m" | cut -d':' -f1)";
					pName="$(echo "$m" | cut -d':' -f2)";
					pMax="$(echo "$m" | cut -d':' -f3)";
					if [ "$pMax" != "M" ]; then
						echo "	s.$pName = make([]$pType, 0, $pMax)";
						echo "";
					fi;
				done;
			fi;
			if [ -z "$embedded" ]; then
				echo "	for _, sl := range l.Sub {";
			else
				echo "	for i := 0; i < len(l.Sub); i++ {";
				echo "		sl := l.Sub[i]";
			fi;
			echo "		switch sl.tag {";
			for type in ${types[@]}; do
				IFS=":";
				local parts=( $type );
				IFS="$OFS";
				pTag="${parts[0]}";
				pType="${parts[1]}";
				pName="${parts[2]}";
				pMax="${parts[4]}";
				if [ ! -z "$pTag" ]; then
					cont=false;
					if [ "${pTag: -1}" = "*" ]; then
						pTag="${pTag:0:-1}";
						cont=true;
					fi;
					echo "		case c$pTag:" | tr -d '/' | tr -d '-';
					if [ "$pMax" = "1" ]; then
						echo "			if ${pName}Set {";
						echo "				if !o.allowMoreThanAllowed {";
						echo "					continue";
						echo "				}";
						echo "";
						echo "				return ErrContext{\"$structureName\", c$pTag, ErrSingleMultiple}";
						echo "			}";
						echo "";
						echo "			${pName}Set = true";
						echo "";
					elif [ "$pMax" != "M" ]; then
						echo "			if len(s.$pName) == $pMax {";
						echo "				if !o.allowMoreThanAllowed {";
						echo "					continue";
						echo "				}";
						echo "";
						echo "				return ErrContext{\"$structureName\", c$pTag, ErrTooMany($pMax)}";
						echo "			}";

						#echo "			${pName}Count++";
					fi;
					if [ "$pMax" = "1" ]; then
						echo "			if err := s.${pName}.parse(&sl, o); err != nil {";
					else
						echo "";
						echo "			var t ${pType}";
						echo "";
						echo "			if err := t.parse(&sl, o); err != nil {";
					fi;
					echo "				return ErrContext{\"$structureName\", c$pTag, err}";
					echo "			}";
					if [ "$pMax" != "1" ]; then
						echo "			s.${pName} = append(s.${pName}, t)";
						echo "";
					fi;
					if [ ! -z "$embedded" ]; then
						echo "			l.Sub = append(l.Sub[:i], l.Sub[i+1:]...)";
						echo "";
						echo "			i--";
					fi;
				fi;
			done;
			if [ -z "$embedded" ]; then
				echo "		default:";
				echo "			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {";
				echo "				return ErrContext{\"$structureName\", sl.tag, ErrUnknownTag}";
				echo "			}";
				echo "			// possibly store in a Other field";
			fi;
			echo "		}";
			echo "	}";
			echo "";
			if [ ${#required} -gt 0 ]; then
				echo "	if !o.allowMissingRequired {";
				for r in "${required[@]}"; do
					echo "		if !${r}Set {";
					echo "			return ErrContext{\"$structureName\", \"$r\", ErrRequiredMissing}";
					echo "		}";
				done;
				echo "	}";
				echo "";
			fi;
		fi;
		if [ ! -z "$embedded" ]; then
			echo "	return s.${embedded}.parse(l, o)";
		fi;
	else
		echo "	for _, sl := range l.Sub {";
		echo "		if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {";
		echo "			return ErrContext{\"$structureName\", sl.tag, ErrUnknownTag}";
		echo "		}";
		echo "		// possibly store in a Other field";
		echo "	}";
	fi;
	if [ -z "$embedded" ]; then
		echo "	return nil";
	fi;
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
		while read line; do
			if [ "${line:0:1}" = "	" ]; then
				IFS=":";
				parts=( ${line:1} );
				IFS="$(echo)";
				pTag="${parts[0]}";
				pType="${parts[1]}";
				pName="${parts[2]}";
				if [ -z "$pName" ]; then
					pName="$pType";
				fi;
				if [ ! -z "$pTag" -a ${#pName} -gt $longest ]; then
					longest=${#pName};
				fi;
				types+=( "${line:1}" );
			else
				IFS="$OFS";
				processStructure "$structureName" $longest "${types[@]}";
				IFS="$(echo)";
				longest=0;
				types=();
				structureName="$line";
			fi;
		done;
		IFS="$OFS";
		processStructure "$structureName" $longest "${types[@]}";
	) < structures.gen;

	cat <<HEREDOC

// MultimediaLink splite between MultimediaLinkID and MultimediaLinkFile
type MultimediaLink struct {
	Data Record
}

func (s *MultimediaLink) parse(l *Line, o options) error {
	var err error
	if l.xrefID != "" {
		t := &MultimediaLinkID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &MultimediaLinkFile{}
		err = t.parse(l, o)
		s.Data = t
	}

	return err
}

// NoteStructure splits between NoteID and NoteText.
type NoteStructure struct {
	Data Record
}

func (s *NoteStructure) parse(l *Line, o options) error {
	var err error

	if l.xrefID != "" {
		t := &NoteID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &NoteText{}
		err = t.parse(l, o)
		s.Data = t
	}

	return err
}

// SourceCitation splits between SourceID and SourceText
type SourceCitation struct {
	Data Record
}

func (s *SourceCitation) parse(l *Line, o options) error {
	var err error

	if l.xrefID != "" {
		t := &SourceID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &SourceText{}
		err = t.parse(l, o)
		s.Data = t
	}

	return err
}

// Trailer type
type Trailer struct{}

func (s *Trailer) parse(*Line, options) error {
	return nil
}

// Errors.
var (
	ErrRequiredMissing = errors.New("required tag missing")
	ErrSingleMultiple  = errors.New("tag was specified more than the one time allowed")
	ErrUnknownTag      = errors.New("unknown tag")
)

// ErrContext adds context to a returned error.
type ErrContext struct {
	Structure, Tag string
	Err            error
}

// Error implements the error interface.
func (e ErrContext) Error() string {
	return e.Tag + ":" + e.Err.Error()
}

// Unwrap goes through the error list to retrieve the underlying
// (non-ErrContext) error.
func (e ErrContext) Unwrap() error {
	err := e.Err
	for {
		if er, ok := err.(ErrContext); ok {
			err = er.Err
			continue
		}
		return err
	}
}

// ErrTooMany is an error returned when too many of a particular tag exist.
type ErrTooMany int

// Error implements the error interface.
func (ErrTooMany) Error() string {
	return "too many tags"
}
HEREDOC
) > structures.go;
