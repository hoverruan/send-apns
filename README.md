## Send apns command-line tool

Send apple push notification message using Go-Apns

### Usage

	Usage: ./send-apns [options...] token body [badge]
	Options:
	  -C="": Setting custom fields, separated with comma, eg: key1=value1,key2=value2
	  -p=false: Using production destination
	  -s="": Sound

### Requirement

Place your certificate files on current directory:

* dev-cert.pem, dev-key.pem: Certificate files for sandbox environment
* prod-cert.pem, prod-key.pem: Certificate files for production environment

Example for generating production certificate files:

	$ openssl x509 -in aps_production.cer -inform der -out prod-cert.pem
	
	$ openssl pkcs12 -nocerts -in Certificates.p12 -out prod-key-withpass.pem
	Enter Import Password:
	MAC verified OK
	Enter PEM pass phrase:
	Verifying - Enter PEM pass phrase:
	
	$ openssl rsa -in prod-key-withpass.pem -out prod-key.pem
	Enter pass phrase for prod-key-withpass.pem:
	writing RSA key
	
	$ rm prod-key-withpass.pem

Useful links:

* <http://sckl.blog.sohu.com/166505417.html> (Chinese)
* <http://www.raywenderlich.com/3443/apple-push-notification-services-tutorial-part-12>
