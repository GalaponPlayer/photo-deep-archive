import { BaseApiResponseAbs } from "../interfaces/base";

export class BaseApiResponse implements BaseApiResponseAbs {
  data: object;
  status: number;
  statusText: string;
  headers: object;
  config: object;
  request: object;

  constructor(
    data: object = {},
    status: number = 200,
    statusText: string = "OK",
    headers: object = {},
    config: object = {},
    request: object = {}
  ) {
    this.data = data;
    this.status = status;
    this.statusText = statusText;
    this.headers = headers;
    this.config = config;
    this.request = request;
  }

  getData(): object {
    if (this.data === null) {
      return {};
    }
    return this.data;
  }

  getStatus(): number {
    return this.status;
  }

  getStatusText(): string {
    return this.statusText;
  }

  getHeaders(): object {
    return this.headers;
  }

  getConfig(): object {
    return this.config;
  }

  getRequest(): object {
    return this.request;
  }

  isSuccess(): boolean {
    return this.status === 200;
  }
}
