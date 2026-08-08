// SPDX-License-Identifier: Apache-2.0
// TypeScript client example against a running PEIRAVELA core api-server.

import { PeiravelaClient } from "../../sdk/ts/peiravela";

const api = new PeiravelaClient("http://127.0.0.1:8080", {});

const health = await api.health();
console.log("health:", health.status, health.version, health.store);

const worlds = await api.listWorlds();
console.log("worlds:", worlds.worlds.length);

export { health, worlds };
