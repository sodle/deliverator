<script lang="ts">
  import { Request, HttpVerb, Header } from "./Request.svelte";
  import { type ResponseType, type Response } from "@tauri-apps/api/http";

  export let request: Request | undefined;
  let response: Response<ResponseType.Text> | undefined;
</script>

<div id="request-builder">
  {#if request}
    <div class="row">
      <fieldset id="fieldset-name">
        <label for="request-name">Name</label>
        <input
          bind:value={request.name}
          id="request-name"
          name="request-name"
        />
      </fieldset>
    </div>
    <div class="row">
      <fieldset id="fieldset-method">
        <label for="request-method">Method</label>
        <select
          bind:value={request.method}
          id="request-method"
          name="request-method"
        >
          {#each Object.keys(HttpVerb) as verb}
            <option>{verb}</option>
          {/each}
        </select>
      </fieldset>
      <fieldset id="fieldset-url">
        <label for="request-url">URL</label>
        <input bind:value={request.url} name="request-url" id="request-url" />
      </fieldset>
      <fieldset id="fieldset-send">
        <button
          on:click={() => {
            request.send().then((r) => (response = r));
          }}>Send</button
        >
      </fieldset>
    </div>
    <div class="row">
      <fieldset id="fieldset-headers">
        <label for="fieldset-headers">Headers</label>
        {#each request.headers as header, i}
          <div class="row">
            <fieldset id={"fieldset-header-" + i + "-key"}>
              <label for={"key-" + i}>Name</label>
              <input
                id={"key-" + i}
                name={"key-" + i}
                bind:value={header.name}
              />
            </fieldset>
            <fieldset id={"fieldset-header-" + i + "-value"}>
              <label for={"value-" + i}>Value</label>
              <input
                id={"value-" + i}
                name={"value-" + i}
                bind:value={header.value}
              />
            </fieldset>
            <fieldset>
              <button
                on:click={() => {
                  request.headers.splice(i, 1);
                  request.headers = request.headers;
                }}>Delete</button
              >
            </fieldset>
          </div>
        {/each}
        <div class="row">
          <button
            on:click={() => {
              request.headers.push(new Header("", ""));
              request.headers = request.headers;
            }}>+ Add header</button
          >
        </div>
      </fieldset>
    </div>
    <div class="row">
      <fieldset id="fieldset-body">
        <label for="request-body">Request body</label>
        <textarea
          id="request-body"
          name="request-body"
          bind:value={request.body}
          rows="10"
        />
      </fieldset>
    </div>
    <hr />
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
</div>

<style>
  #request-builder {
    flex: 3;
    flex-basis: 75vw;
    padding: 1em;
    height: 100vh;
    overflow: auto;
    box-sizing: border-box;
  }

  #response-window {
    text-align: left;
  }

  fieldset {
    flex: 1 1 auto;
    border: none;
  }

  label {
    display: block;
    text-align: left;
  }

  input,
  select,
  button,
  textarea {
    display: block;
    width: 100%;
    box-sizing: border-box;
  }

  select {
    background-color: black;
    color: white;
    font-size: 1em;
    /* appearance: none; */
    border: none;
    padding: 0.8em 1.2em;
  }

  button {
    height: 100%;
  }

  textarea {
    background: black;
    color: white;
    font-size: 1em;
    font-family: monospace;
  }

  #fieldset-url {
    flex-basis: 100%;
  }
</style>
