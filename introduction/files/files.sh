#create files named a ! \ "

touch a \! \\ \"

#to reate a directory named `
mkdir -p \`

#copy ! into `
cp \! \`/

#conditional based on MOVE_A  variable
if [ "MOVE_A" = "yes" ]; then
	mv a \`/
elif [ "MOVE_A" = "no" ]; then
	rm a
else 
	echo "doing nothing"
fi
