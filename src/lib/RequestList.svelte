<script lang="ts">
  import { Request } from "./Request.svelte";

  export let requests: Request[];
  export let selectedIndex: number;
  export let deleteHandler: (index: number) => void;
  export let addHandler: () => void;
</script>

<div id="request-list">
  {#each requests as request, i}
    <a
      class={selectedIndex == i
        ? "request-list-item selected"
        : "request-list-item"}
      on:click={() => (selectedIndex = i)}
      href={"#" + i}
    >
      <div>
        <h1>{request.name}</h1>
        <h2>{request.method} {request.url}</h2>
        <button class="delete" on:click={() => deleteHandler(i)}>Delete</button>
      </div>
    </a>
  {/each}
  <a class="request-list-item add" on:click={() => addHandler()} href="#add">
    <div>+ New request</div>
  </a>
</div>

<style>
  #request-list {
    flex: 1;
    flex-basis: 25vw;
    border-right: 2px solid white;
    padding: 1em;
    height: 100vh;
    box-sizing: border-box;
    overflow: auto;
  }

  .request-list-item {
    display: block;
    border: 1px solid darkgray;
    border-radius: 1em;
    padding: 1em;
    margin-bottom: 0.5em;
  }
  .request-list-item>div {
    user-select: none;
    box-sizing: border-box;
    height: auto;
    overflow: visible;
  }

  .request-list-item:last-of-type {
    margin-bottom: 0;
  }

  .request-list-item:link,
  .request-list-item:visited,
  .request-list-item:hover,
  .request-list-item:focus,
  .request-list-item:active {
    color: inherit;
  }

  .request-list-item.selected {
    background-color: rgba(255, 255, 255, 0.3);
  }

  .request-list-item * {
    text-align: left;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow-x: hidden;
  }

  .request-list-item h1 {
    margin: 0;
    padding-top: 0.2em;
    padding-bottom: 0.2em;
    font-size: 2em;
    font-weight: normal;
  }

  .request-list-item h2 {
    margin: 0;
    padding-bottom: 0.5em;
    font-size: 1.2em;
    font-weight: normal;
  }

  .request-list-item .delete {
    width: 100%;
  }
  .request-list-item .delete:hover {
    background-color: red;
  }

  .request-list-item:not(.selected):hover {
    background-color: rgba(255, 255, 255, 0.1);
  }
</style>
