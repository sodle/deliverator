<script lang="ts">
  import {
    TabItem,
    Table,
    TableHead,
    TableHeadCell,
    TableBody,
    TableBodyRow,
    TableBodyCell,
    Input,
    Button,
  } from "flowbite-svelte";
  import { Request, Header } from "../../Request.svelte";

  export let request: Request;
</script>

<TabItem open title="Headers">
  <Table>
    <TableHead>
      <TableHeadCell class="pe-0 ps-2">Name</TableHeadCell>
      <TableHeadCell class="pe-0 ps-2">Value</TableHeadCell>
      <TableHeadCell class="pe-0 ps-2">
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
            <Button
              class="block w-full rounded-none border border-gray-600"
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
            class="block h-full w-full rounded-none rounded-b-md border border-t-0 border-gray-600"
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
