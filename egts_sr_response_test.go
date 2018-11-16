package main

/*
File: 'egts.bin'
Packet data:
0100000B0010008600001886000006005F002001010003005F00001373


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
 Frame Data Length   - 16
 Packet ID           - 134
 No route info       -
 Header Check Sum    - 0x18

EGTS Service Layer:
---------------------
 Validating result   - 0 (OK)

 Packet Type         - EGTS_PT_RESPONSE
 Service Layer CS    - 0x7313
 Responded Packet ID - 134
 Processing Result   - 0 (OK)

   Service Layer Record:
   ---------------------
   Validating Result    - 0 (OK)

   Record Length               - 6
   Record Number               - 95
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

	  Subrecord Type      - 0 (EGTS_SR_RESPONSE)
	  Subrecord Length    - 3
	  Confirmed Record Number- 95
	  Record Status          - 0 (OK)
*/

