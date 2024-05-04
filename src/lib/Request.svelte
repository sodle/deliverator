<script lang="ts" context="module">
  import { Body, ResponseType, getClient, type HttpOptions, type Response } from "@tauri-apps/api/http";

  export enum HttpVerb {
    GET = "GET",
    POST = "POST",
    PUT = "PUT",
    PATCH = "PATCH",
    DELETE = "DELETE",
    HEAD = "HEAD",
    OPTIONS = "OPTIONS",
    TRACE = "TRACE",
    CONNECT = "CONNECT",
  }

  export class Header {
    name: string;
    value: string;

    constructor(name: string, value: string) {
      this.name = name;
      this.value = value;
    }
  }

  export class Request {
    name: string;
    url: string;
    method: HttpVerb;
    body: string | null;
    headers: Header[];

    constructor() {
      this.name = "New Request";
      this.url = "https://example.com";
      this.method = HttpVerb.GET;
      this.headers = [new Header("Content-Type", "application/json")];
      this.body = ""
    }

    async send(): Promise<Response<any>> {
      let options: HttpOptions = {
        url: this.url,
        method: this.method,
        responseType: ResponseType.Text
      }

      let headers: Record<string, any> = {}
      this.headers.forEach(h => {
        headers[h.name] = h.value
      })
      options.headers = headers

      if (this.body && this.body.length) {
        options.body = Body.text(this.body)
      }

      let client = await getClient()
      return client.request(options)
    }
  }
</script>
