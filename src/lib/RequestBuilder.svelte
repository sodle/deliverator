<script lang="ts">
  import {
    Button,
    Input,
    Label,
    Select,
    TabItem,
    Tabs,
    Textarea,
  } from "flowbite-svelte";

  import { Request, HttpVerb, Header } from "./Request.svelte";
  import { type ResponseType, type Response } from "@tauri-apps/api/http";
    import HeaderList from "./ui/request/HeaderList.svelte";

  export let request: Request | undefined;
  let response: Response<ResponseType.Text> | undefined;
</script>

<main class="m-5 w-2/3 overflow-auto">
  {#if request}
    <div class="flex w-full">
      <Label for="name" class="block w-full">
        <span class="sr-only">Name</span>
        <Input placeholder="Name" id="name" bind:value={request.name} />
      </Label>
    </div>
    <div class="mb-3 mt-3 flex w-full">
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
        <Textarea name="body" class="block w-full rounded-t-none" rows="7" bind:value={request.body} placeholder="Body" />
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
