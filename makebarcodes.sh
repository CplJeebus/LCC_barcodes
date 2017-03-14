#!/bin/bash

file=$1

cat head.html.part 
url="https://requestb.in/15w9a031"
while read X 
	do 
	uName=$(python -c "import urllib, sys; print urllib.quote(sys.argv[1])" "${X}")
	#echo Hello $uName
	In=${url}?name=${uName}
	#echo GoodBye $In
	F=$(echo $X | tr ' ' '_')
	FF=${F}.png
	zint -b 58 --scale=2 --smalltext --border=20 -o $FF  -d "${In}"
        cat <<EOF
	<tr>
		<td><img src="./${FF}"></td>
		<td class="separator" />
		<td><img src="./${FF}"></td>
		<td class="separator" />
		<td><img src="./${FF}"></td>
	</tr>
        <tr>
                <td class="t"><center>${X}</center></td>
                <td class="separator" />
                <td class="t"><center>${X}</center></td>
                <td class="separator" />
                <td class="t"><center>${X}</center></td>
        </tr>
EOF
	
	done < $file

cat tail.html.part
