import asyncio
import nats
import json
import logging
from typing import Any, Dict

# Main async function for NATs
async def main(queue: asyncio.Queue, args: Dict[str, Any]):
    # Arguments
    logger = logging.getLogger()
    subject = args.get("subject")
    host = args.get("host")
    port = args.get("port")
    # Connect to the NATS topis with the concatinated port
    nc = await nats.connect(host+":"+str(port))
    # Start to subscribe to the topic
    sub = await nc.subscribe(subject)

    # Loop through the messages.
    try:
        async for msg in sub.messages:
            # Since it is json create a python dictionary out of it.
            data = json.loads(msg.data.decode())
            #https://github.com/ansible/event-driven-ansible/issues/60
            stringdict = str(data)
            a = stringdict.replace('/','_')
            b = a.replace('-','_')
            respdict = eval(b)
            await queue.put(respdict[0])
    except json.decoder.JSONDecodeError as e:
        logger.error(e)
    finally:
        logger.info("Stopping nats sub")
        await nc.close()

if __name__ == "__main__":
    class MockQueue:
        async def put(self, event):
            print(event)

    asyncio.run(main(MockQueue(), {}))