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

  import { Request, HttpVerb, Header } from "./Request.svelte";
  import { type ResponseType, type Response } from "@tauri-apps/api/http";

  export let request: Request | undefined;
  let response: Response<ResponseType.Text> | undefined;
</script>

<main class="m-5 w-2/3">
  {#if request}
    <div class="flex w-full">
      <Label for="name" class="block w-full">
        <span class="sr-only">Name</span>
        <Input placeholder="Name" id="name" bind:value={request.name} />
      </Label>
    </div>
    <div class="flex w-full mt-3 mb-3">
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
      <Input class="w-3/5 rounded-none" id="url" placeholder="URL" bind:value={request.url} />
      <Button
        class="w-1/5 rounded-s-none"
        on:click={() => {
          request.send().then((r) => (response = r));
        }}
      >
        Send
      </Button>
    </div>
    <Tabs
      tabStyle="full"
      defaultClass="flex rounded-lg divide-x rtl:divide-x-reverse"
    >
      <TabItem open title="Headers" class="w-full">
        <Table>
          <TableHead>
            <TableHeadCell class="ps-2 pe-0">Name</TableHeadCell>
            <TableHeadCell class="ps-2 pe-0">Value</TableHeadCell>
            <TableHeadCell class="ps-2 pe-0">
              <span class="sr-only">Delete</span>
            </TableHeadCell>
          </TableHead>
          <TableBody>
            {#each request.headers as header, i}
              <TableBodyRow>
                <TableBodyCell class="p-0">
                  <Input class="rounded-none" bind:value={header.name} />
                </TableBodyCell>
                <TableBodyCell class="p-0">
                  <Input class="rounded-none" bind:value={header.value} />
                </TableBodyCell>
                <TableBodyCell class="p-0">
                  <Button class="block w-full rounded-none border border-gray-600"
                    on:click={() => {
                      request.headers.splice(i, 1);
                      request.headers = request.headers;
                    }}
                  >
                    X
                  </Button>
                </TableBodyCell>
              </TableBodyRow>
            {/each}
            <TableBodyRow>
              <TableBodyCell colspan="3" class="p-0">
                <Button
                  class="block w-full h-full rounded-none border border-t-0 border-gray-600"
                  on:click={() => {
                    request.headers.push(new Header("", ""));
                    request.headers = request.headers;
                  }}
                >
                  + Add Header
                </Button>
              </TableBodyCell>
            </TableBodyRow>
          </TableBody>
        </Table>
      </TabItem>
      <TabItem title="Body" class="w-full">
        <Textarea class="block w-full" rows="7" bind:value={request.body} />
      </TabItem>
    </Tabs>
    <hr class="mb-5 mt-5" />
    {#if response}
      <div id="response-window">
        <p>Status: {response.status}</p>
        <div>
          <p>Headers:</p>
          <ul>
            {#each Object.keys(response.headers) as name}
              <li>{name}: {response.headers[name]}</li>
            {/each}
          </ul>
        </div>
        <div>
          <p>Body:</p>
          <p>{response.data}</p>
        </div>
      </div>
    {/if}
  {:else}
    <p>No request selected.</p>
  {/if}
</main>
