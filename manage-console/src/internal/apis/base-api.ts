import axios, { AxiosRequestConfig } from "axios";
import {
  BaseApiResponseAbs,
  BaseGetRequestAbs,
  BasePostRequestAbs,
} from "../ports/interfaces/base";
import { BaseApiResponse } from "../ports/impls/base";

export class InternalApiBase {
  //TODO: env
  private baseUrl = "http://localhost:3000";
  private path = "";

  constructor(path: string = "") {
    this.baseUrl = import.meta.env.VITE_API_BASE_URL as string;
    this.path = path;
  }

  setPath(path: string) {
    this.path = path;
  }

  getEndPointUrl(): URL {
    return new URL(this.path, this.baseUrl);
  }

  async get(
    request: BaseGetRequestAbs,
    config: AxiosRequestConfig | undefined
  ): Promise<BaseApiResponseAbs> {
    //TODO: publicは別で作るので、AUthenticatorを使う
    const cfg = config
      ? config
      : { headers: { "content-type": "application/json" } };
    //TODO: cognitoの認証が必要
    const url = this.getEndPointUrl();
    Object.entries(request.params).forEach(([key, value]) => {
      url.searchParams.append(key, value);
    });
    const res = await axios.get(url.toString(), cfg);
    const response = new BaseApiResponse(
      res.data,
      res.status,
      res.statusText,
      res.headers,
      res.config,
      res.request
    );
    return response;
  }

  async post(
    request: BasePostRequestAbs,
    config: AxiosRequestConfig | undefined
  ): Promise<BaseApiResponse> {
    //TODO: publicは別で作るので、AUthenticatorを使う
    const cfg = config
      ? config
      : { headers: { "content-type": "application/json" } };
    //TODO: cognitoの認証が必要
    const url = this.getEndPointUrl();
    const res = await axios.post(
      url.toString(),
      request.getParamsObject(),
      cfg
    );
    const response = new BaseApiResponse(
      res.data,
      res.status,
      res.statusText,
      res.headers,
      res.config,
      res.request
    );
    return response;
  }
}
