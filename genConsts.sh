#!/bin/bash

maxLength=0;

while read line; do
	if [ ${#line} -gt $maxLength ]; then
		maxLength=${#line};
	fi;
done < <(
	{
		echo -e "CONC\nCONT";
		grep "::" types.gen | cut -d':' -f3- | grep -v "^(" | tr ':' '\n';
		grep "	[A-Z]" structures.gen | cut -d':' -f1 | tr -d '	'
	} | sort | uniq;
);

(
	cat <<HEREDOC
package gedcom

// File automatically generated with ./genConsts.sh

const (
HEREDOC


	while read line; do
		name="$(echo "$line" | tr -d "-" | tr -d "/")";
		echo -n "	c$name";
		pad=$(( $maxLength - ${#name} ));
		while [ $pad -gt 0 ]; do
			let "pad--"
			echo -n " ";
		done;
		echo " = \"$line\"";
	done < <(
		{
			echo -e "0\n1\n2\n3\nCONC\nCONT";
			grep "::" types.gen | cut -d':' -f3- | grep -v "^(" | tr ':' '\n';
			grep "	[A-Z]" structures.gen | cut -d':' -f1 | tr -d '	'
		} | sort | uniq;
	);

	echo ")";
) > consts.go;
