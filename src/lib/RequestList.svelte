<script lang="ts">
  import { Button, Sidebar } from "flowbite-svelte";
  import { Request } from "./Request.svelte";

  export let requests: Request[];
  export let selectedIndex: number;
  export let deleteHandler: (index: number) => void;
  export let addHandler: () => void;
</script>

<Sidebar
  class="box-border w-1/3 overflow-auto border-0 border-r border-black px-5 py-10 dark:border-white"
>
  {#each requests as request, i}
    <a
      class={selectedIndex == i
        ? "mb-2 block rounded-md border bg-gray-100 p-2 font-bold text-black dark:bg-gray-500 dark:text-white"
        : "mb-2 block rounded-md border bg-gray-50 p-2 text-black dark:bg-gray-700 dark:text-white"}
      on:click={() => {
        selectedIndex = i;
      }}
      href={"#" + i}
    >
      <div class="flex">
        <div class="flex-grow truncate">
          <h1 class="text-lg">{request.name}</h1>
          <h2 class="text-sm font-light">{request.method} {request.url}</h2>
        </div>
        <Button
          class="flex-shrink p-1"
          href="#delete"
          on:click={() => deleteHandler(i)}>X</Button
        >
      </div>
    </a>
  {/each}
  <Button class="w-full" on:click={() => addHandler()} href="#add"
    >+ New Request</Button
  >
</Sidebar>
