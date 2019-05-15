from time import sleep

import socket

# TEST_FILE = '/Users/kuznetsovin/Projects/TMS/tools/receiver_testing/bnso_30277669.csv'
# TEST_FILE = '/Users/kuznetsovin/Projects/TMS/tools/receiver_testing/test_egts.csv'
TEST_FILE = 'test2.csv'

if __name__ == '__main__':
    TEST_ADDR = ('195.88.196.133', 5020)

    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(TEST_ADDR)

    BUFF = 2048

    with open(TEST_FILE) as f:
        for rec in f.readlines():
            print("send: {}".format(rec))
            package = bytes.fromhex(rec[:-1])            
            client.send(package)

            rec_package = client.recv(BUFF)
            print("received: {}".format(rec_package.hex()))
            sleep(1)

    client.close()
