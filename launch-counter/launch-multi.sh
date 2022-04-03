FILENAME=$(readlink -f $0)
DIR=$(dirname $FILENAME)

#HOSTS=blue01,blue02
#HOSTS=blue04,yellow13
#HOSTS=blue04,yellow13,yellow14
HOSTS=blue01,blue02,blue03,blue04

pussh -h $HOSTS $DIR/launch.sh

