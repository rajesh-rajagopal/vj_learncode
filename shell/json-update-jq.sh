cid='12'
echo "{\"data\":{\"CLUSTER_ID\":\"$cid\"}}">test
# merge two json files 
jq -s '.[0] * .[1]' test1.json test
