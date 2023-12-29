import yaml
import argparse
import asyncio
import aio_pika
import aio_pika.abc
import json
from pprint import pprint

async def main(conStr, exchange, routing_key):
    connection = await aio_pika.connect_robust(
        conStr
    )
    async with connection:
        channel: aio_pika.abc.AbstractChannel = await connection.channel()
        queue: aio_pika.abc.AbstractQueue = await channel.declare_queue(auto_delete=True)
        exchange = await channel.declare_exchange(exchange)
        await queue.bind(exchange, routing_key)
        async with queue.iterator() as queue_iter:
            async for message in queue_iter:
                async with message.process():
                    data = json.loads(message.body)
                    pprint(data)
                    print('\n\n')




if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-c", "--config", help="Конфигурационный файл", type=str)


    args = parser.parse_args()

    with open(args.config) as f:
        c = yaml.safe_load(f)
        config = c['storage']['rabbitmq']

    conStr = "amqp://%s:%s@%s:%s/%s" % (config["user"], config["password"], config["host"], config["port"], config["virtual_host"])


    asyncio.run(main(conStr, config["exchange"], config["key"]))