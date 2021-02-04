# By Ubiot
while getopts s:k: flag
do
    case "${flag}" in
        s) SEED=${OPTARG};;
        k) KEYINDEX=${OPTARG};;
    esac
done

echo "Seed: $SEED";
echo "Key Index: $KEYINDEX";

file='complex-contract.go'

cd smart-contract/contract-tutorial/

insertS='var walletSeed1 string = "'$SEED'"'
insertK="var walletKeyIndex uint64 = $KEYINDEX"

sed -i "24c\ $insertS" $file

sed -i "28c\ $insertK" $file

cd /home/rtorrealba/ubiot-network

./start-network.sh
