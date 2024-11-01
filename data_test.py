import asyncio

from viam.app.viam_client import ViamClient
from viam.rpc.dial import DialOptions
import bson

async def main():
    viam = await ViamClient.create_from_dial_options(DialOptions.with_api_key(api_key="fc3aoh6xx5kmfzb135guvxjcf12l02k8", api_key_id="b5cf5cb8-5dcc-41d4-8b1f-ca79e28cef72"))
    dc = viam.data_client

    query = bson.encode({"$match": {"organization_id": "e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb"}})
    limit = bson.encode({"$limit": 1})

    data = await dc.tabular_data_by_mql("e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb", [query, limit])
    for item in data:
        for (key, value) in item.items():
            print("the value type is:", type(value))
            # assert(data[0], type(value))

    # print(data)

    viam.close()

if __name__ == "__main__":
    asyncio.run(main())
