# go-2c2p-sandbox
Sandbox purely written in Go lang because 2c2p sandbox is not good enough

# Usage
Just ran ./main

The binary will serve endpoint at "http://127.0.0.1:8080/payment".
The endpoint will do 2 actions.

The first action will make post request json call to "result_url_2".

The second action will show you the success page and if you click "Done", it will redirect back to "result_url_1".

# Mocking Result 
main success 

main pending

main fail

# Note
If you are using my sandbox, you cant compare hash value. Just assume as "True" as it is development purpose.

