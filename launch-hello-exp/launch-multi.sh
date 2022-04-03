FILENAME=$(readlink -f $0)
DIR=$(dirname $FILENAME)

$DIR/clean.sh

#HOSTS=blue01,blue02,blue03,blue04
#HOSTS=blue01,blue02,blue03,blue04,blue05,blue06,blue07,blue08,blue09,blue10,blue11,blue12,blue13,blue14,blue15,yellow04,yellow05,yellow07,yellow08,yellow09,yellow10,yellow11
HOSTS=yellow04,yellow05,yellow07,yellow08,yellow09,yellow10,yellow11

NUM=$1
echo "$NUM clients per machine"
pussh -h $HOSTS $DIR/launch.sh $NUM

