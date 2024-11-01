import * as VIAM from '@viamrobotics/sdk';
import { BSON } from 'bsonfy';

const ORG_ID = 'e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb';
const API_KEY_ID = 'b5cf5cb8-5dcc-41d4-8b1f-ca79e28cef72';
const API_KEY_SECRET = 'fc3aoh6xx5kmfzb135guvxjcf12l02k8';


async function connect(): Promise<VIAM.ViamClient> {
  const opts: VIAM.ViamClientOptions = {
    credentials: {
      type: 'api-key',
      authEntity: API_KEY_ID,
      payload: API_KEY_SECRET,
    },
  };

  const client = await VIAM.createViamClient(opts);
  return client;
}

const button = <HTMLButtonElement>document.getElementById('main-button');

async function run(client: VIAM.ViamClient) {
  try {
    button.disabled = true;
    const textElement = <HTMLParagraphElement>document.getElementById('text');
    textElement.innerHTML = 'waiting for data...';

    const query = BSON.serialize({ $match: { organization_id: "e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb" } });
    const limit = BSON.serialize({ $limit: 1 });

    const dataList = await client.dataClient.tabularDataByMQL(
      ORG_ID,
      [query, limit]
    );

    for (const item of dataList) {
      for (const [key, value] of Object.entries(item as Record<string, any>)) {
        if (value instanceof Date) {
          console.log(`The value for "${key}" is of type Date:`, value);
        } else {
          console.log(`The value type for "${key}" is:`, typeof value);
        }
      }
    }

    textElement.innerHTML = JSON.stringify(dataList, null, 2);
  } finally {
    button.disabled = false;
  }
}


async function main() {
  let client: VIAM.ViamClient;
  try {
    button.textContent = 'Connecting...';
    client = await connect();
    button.textContent = 'Click for data';
  } catch (error) {
    button.textContent = 'Unable to connect';
    console.error(error);
    return;
  }

  // Make the button in our app do something interesting
  button.addEventListener('click', async () => {
    await run(client);
  });
  button.disabled = false;
}

main();
