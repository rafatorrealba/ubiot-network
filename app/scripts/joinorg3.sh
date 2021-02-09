# By Ubiot
while getopts c: flag
do
    case "${flag}" in
        c) CHANNELNAME=${OPTARG};;
    esac
done

echo
echo Joining Org3 to the channel
cd /home/rtorrealba/ubiot-network/test-network/addOrg3
bash addOrg3.sh up -c $CHANNELNAME