import axios, { AxiosRequestConfig } from "axios";
import {
  BaseApiResponseAbs,
  BaseGetRequestAbs,
  BasePostRequestAbs,
} from "../ports/interfaces/base";
import { BaseApiResponse } from "../ports/impls/base";

export class InternalApiBase {
  //TODO: env
  private baseUrl: string = "http://localhost:3000";
  private path: string = "";

  constructor(path: string = "") {
    // this.baseUrl = import.meta.env.VITE_API_BASE_URL as string;
    this.baseUrl = import.meta.env.VITE_API_BASE_URL.endsWith("/")
      ? import.meta.env.VITE_API_BASE_URL
      : import.meta.env.VITE_API_BASE_URL + "/";
    this.path = path;
  }

  setPath(path: string) {
    this.path = path;
  }

  getProdPath(): string {
    return "/prod" + this.path;
  }

  getEndPointUrl(): URL {
    return new URL(this.getProdPath(), this.baseUrl);
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

    let res = null;
    try {
      res = await axios.post(url.toString(), request.getParamsObject(), cfg);
      const response = new BaseApiResponse(
        res.data,
        res.status,
        res.statusText,
        res.headers,
        res.config,
        res.request
      );
      return response;
    } catch (e) {
      console.log(e);
      if (axios.isAxiosError(e) && e.response) {
        const res = e.response;
        const response = new BaseApiResponse(
          res.data,
          res.status,
          res.statusText,
          res.headers,
          res.config,
          res.request
        );
        return response;
      } else {
        const response = new BaseApiResponse(
          {},
          500,
          "Internal Server Error",
          {},
          {},
          {}
        );
        return response;
      }
    }
  }
}
