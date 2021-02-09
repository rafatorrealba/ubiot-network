# By Ubiot
while getopts c: flag
do
    case "${flag}" in
        c) CHANNELNAME=${OPTARG};;
    esac
done

echo
echo Creating new channel
cd /home/rtorrealba/ubiot-network/test-network
bash network.sh createChannel -c $CHANNELNAME