package main

/*
File: 'egts.bin'
Packet data:
0100000B000B001538011104001538200101090100003CBC


EGTS Transport Layer:
---------------------
 Validating result   - 0 (OK)

 Protocol Version    - 1
 Security Key ID     - 0
 Flags               - 00000000b (0x00)
	  Prefix         - 00
	  Route          -   0
	  Encryption Alg -    00
	  Compression    -      0
	  Priority       -       00 (the highest)
 Header Length       - 11
 Header Encoding     - 0
 Frame Data Length   - 11
 Packet ID           - 14357
 No route info       -
 Header Check Sum    - 0x11

EGTS Service Layer:
---------------------
 Validating result   - 0 (OK)

 Packet Type         - EGTS_PT_APPDATA
 Service Layer CS    - 0xBC3C

   Service Layer Record:
   ---------------------
   Validating Result    - 0 (OK)

   Record Length               - 4
   Record Number               - 14357
   Record flags                -     00100000b (0x20)
	   Sourse Service On Device    - 0
	   Recipient Service On Device -  0
	   Group Flag                  -   1
	   Record Processing Priority  -    00 (the highest)
	   Time Field Exists           -      0
	   Event ID Field Exists       -       0
	   Object ID Field Exists      -        0
   Source Service Type         - 1 (EGTS_AUTH_SERVICE)
   Recipient Service Type      - 1 (EGTS_AUTH_SERVICE)

	  Subrecord Data:
	  ------------------
	  Validating Result   - 0 (OK)

	  Subrecord Type      - 9 (EGTS_SR_RESULT_CODE)
	  Subrecord Length    - 1
	  Result Code            - 0 (OK)

*/
