echo "Request 1"
vault write transit/encrypt/test plaintext=NDExMSAxMTExIDExMTEgMTExMQo=
echo "\nRequest 2"
vault write transit/encrypt/test plaintext=NDIyMiAyMjIyIDIyMjIgMjIyMgo=
echo "\nRequest 3"
vault write transit/encrypt/test plaintext=NDMzMyAzMzMzIDMzMzMgMzMzMwo=
