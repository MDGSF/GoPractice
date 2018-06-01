

openssl genrsa -aes256 -out private/cakey.pem -passout pass:123456
openssl req -new -key private/cakey.pem -out private/ca.csr -subj "/C=CN/ST=myprovince/L=mycity/O=myorganization/OU=mygroup/CN=localhost" -passin pass:123456
openssl x509 -req -days 365 -sha1 -extensions v3_ca -signkey private/cakey.pem -in private/ca.csr -out certs/ca.cer -passin pass:123456

openssl genrsa -aes256 -out private/sever-key.pem -passout pass:123456
openssl req -new -key private/server-key.pem -out private/server.csr -subj "/C=CN/ST=myprovince/L=mycity/O=myorganization/OU=mygroup/CN=localhost" -passin pass:123456
openssl x509 -req -days 365 -sha1 -extensions v3_req -CA certs/ca.cer -CAkey private/cakey.pem -CAserial ca.srl -CAcreateserial -in private/server.csr -out certs/server.cer -passin pass:123456

openssl genrsa -aes256 -out private/client-key.pem -passout pass:123456
openssl req -new -key private/client-key.pem -out private/client.csr -subj "/C=CN/ST=myprovince/L=mycity/O=myorganization/OU=mygroup/CN=localhost" -passin pass:123456
openssl x509 -req -days 365 -sha1 -extensions v3_req -CA certs/ca.cer -CAkey private/cakey.pem -CAserial ca.srl -in private/client.csr -out certs/client.cer -passin pass:123456

openssl rsa -in private/server-key.pem -out private/server-key-out.pem -passin pass:123456
openssl rsa -in private/client-key.pem -out private/client-key-out.pem -passin pass:123456

