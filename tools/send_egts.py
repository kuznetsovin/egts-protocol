#!/usr/bin/env python3
from time import sleep
import argparse

import socket

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("-s", "--host", help="Адрес сервера для приема ЕГТС", default='localhost', type=str)
    parser.add_argument("-p", "--port", help="Порт сервера для приема ЕГТС", default=6000, type=int)
    parser.add_argument("file", help="Тестовый файл с пакетами", type=str)
    
    args = parser.parse_args()

    TEST_ADDR = (args.host, args.port)

    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(TEST_ADDR)

    BUFF = 2048

    with open(args.file) as f:
        for rec in f.readlines():
            print("send: {}".format(rec))
            package = bytes.fromhex(rec[:-1])            
            client.send(package)

            rec_package = client.recv(BUFF)
            print("received: {}".format(rec_package.hex()))
            sleep(1)

    client.close()
