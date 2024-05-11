<script lang="ts">
  import {
    Button,
    Input,
    Label,
    Select,
    TabItem,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Tabs,
    Textarea,
  } from "flowbite-svelte";
  import HeaderList from "./ui/request/HeaderList.svelte";

  import { Request, HttpVerb } from "./Request.svelte";
  import { type ResponseType, type Response } from "@tauri-apps/api/http";
  import { afterUpdate } from "svelte";
  import { appWindow } from "@tauri-apps/api/window";

  export let request: Request | undefined;
  let response: Response<ResponseType.Text> | undefined;

  afterUpdate(() => {
    if (request) {
      appWindow.setTitle(`${request.name} - Deliverator`);
    } else {
      appWindow.setTitle("Deliverator: REST Client");
    }
  });
</script>

<main class="m-5 w-2/3 overflow-auto">
  {#if request}
    <div class="m-5">
      <div class="flex w-full">
        <Label for="name" class="block w-full">
          <span class="sr-only">Name</span>
          <Input placeholder="Name" id="name" bind:value={request.name} />
        </Label>
      </div>
      <div class="mb-5 mt-5 flex w-full">
        <Label for="method" class="sr-only">Method</Label>
        <Select
          class="w-1/5 rounded-e-none"
          id="method"
          placeholder="Method"
          bind:value={request.method}
          items={Object.keys(HttpVerb).map((v) => {
            return { name: v, value: v };
          })}
        />
        <Label for="url" class="sr-only">URL</Label>
        <Input
          class="w-3/5 rounded-none"
          id="url"
          placeholder="URL"
          bind:value={request.url}
        />
        <Button
          class="w-1/5 rounded-s-none"
          on:click={() => {
            request.send().then((r) => (response = r));
          }}
        >
          Send
        </Button>
      </div>
      <Tabs contentClass="m-0">
        <HeaderList {request} />
        <TabItem title="Body">
          <Label class="sr-only" for="body">Body</Label>
          <Textarea
            name="body"
            class="block w-full rounded-t-none"
            rows="7"
            bind:value={request.body}
            placeholder="Body"
          />
        </TabItem>
        <TabItem title="Curl Command">
          <Textarea
            name="body"
            class="block w-full rounded-t-none font-mono"
            rows="3"
            disabled
            value={request.asCurl()}
          />
        </TabItem>
      </Tabs>
    </div>
    <hr class="mb-5 mt-5" />
    {#if response}
      <div
        class="m-5 rounded-md bg-gray-50 p-5 text-gray-900 dark:bg-gray-700 dark:text-white"
      >
        Status: {response.status}
      </div>
      <div
        class="m-5 mt-0 rounded-md bg-gray-50 pt-2 text-gray-900 dark:bg-gray-700 dark:text-white"
      >
        <h1 class="mx-5 text-gray-900 dark:text-white">Headers</h1>
        <Table>
          <TableHead>
            <TableHeadCell>Name</TableHeadCell>
            <TableHeadCell>Value</TableHeadCell>
          </TableHead>
          <TableBody>
            {#each Object.keys(response.headers) as name}
              <TableBodyRow>
                <TableBodyCell>{name}</TableBodyCell>
                <TableBodyCell>{response.headers[name]}</TableBodyCell>
              </TableBodyRow>
            {/each}
          </TableBody>
        </Table>
      </div>
      <div
        class="m-5 mt-0 rounded-md bg-gray-50 py-2 text-gray-900 dark:bg-gray-700 dark:text-white"
      >
        <h1 class="mx-5 text-gray-900 dark:text-white">Body</h1>
        <div class="m-5 font-mono text-gray-800 dark:text-gray-50">
          {response.data}
        </div>
      </div>
    {/if}
  {:else}
    <p>No request selected.</p>
  {/if}
</main>
