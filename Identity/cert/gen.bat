DEL  *%CD%pem

@REM 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj \C=ID\ST=Yogyakarta\L=Sleman\O=Artworkfairs\OU=Marketplace\CN=*%CD%yudafatah.com\emailAddress=yudafatah@gmail.com

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

@REM 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj \C=ID\ST=Yogyakarta\L=Sleman\O=Artworkfairs\OU=Marketplace\CN=*%CD%yudafatah.com\emailAddress=yudafatah@gmail.com

@REM 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

@REM 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj \C=ID\ST=Yogyakarta\L=Sleman\O=Artworkfairs\OU=Marketplace\CN=*%CD%yudafatah.com\emailAddress=yudafatah@gmail.com

@REM 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text