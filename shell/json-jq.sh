jq --arg mod "acme" \
   --arg resolved "some_url" \
  '.dependencies[$mod].resolved=$resolved' \
  <in.json >out.json
