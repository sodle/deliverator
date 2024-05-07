<script lang="ts">
  import "./app.css";

  import RequestList from "./lib/RequestList.svelte";
  import RequestBuilder from "./lib/RequestBuilder.svelte";
  import { confirm } from "@tauri-apps/api/dialog";

  import { Request } from "./lib/Request.svelte";
  let requests = [new Request()];
  let selectedIndex = 0;

  function deleteRequest(index: number) {
    confirm("Really delete?", { type: "warning" }).then((confirmed) => {
      if (confirmed) {
        requests.splice(index, 1);
        requests = requests;
      }
    });
  }

  function createRequest() {
    requests.push(new Request());
    requests = requests;
  }
</script>

<main class="container">
  <div class="row">
    <RequestList
      {requests}
      bind:selectedIndex
      deleteHandler={deleteRequest}
      addHandler={createRequest}
    ></RequestList>
    <RequestBuilder bind:request={requests[selectedIndex]}></RequestBuilder>
  </div>
</main>
