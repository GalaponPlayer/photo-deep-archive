export interface BaseApiResponseAbs {
  data: object;
  status: number;
  statusText: string;
  headers: object;
  config: object;
  request: object;

  getData(): object;
  getStatus(): number;
  getStatusText(): string;
  getHeaders(): object;
  getConfig(): object;
  getRequest(): object;
}

export interface BaseGetRequestAbs {
  params: Record<string, string>;
}

export interface BasePostRequestAbs {
  getParamsObject(): object;
}
