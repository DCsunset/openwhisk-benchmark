FILENAME=$(readlink -f $0)
DIR=$(dirname $FILENAME)
HOSTNAME=$(hostname)

NUM=${1:-1}

cd $DIR
for i in $(seq 1 $NUM); do
	./launch-hello-exp "$HOSTNAME-$i" &
done
