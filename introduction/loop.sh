n=$1

if [ $n -gt 100 ]; then
	n=100
fi

for i in $(seq 1 $n); do
	echo "This is loop number $i"
done
