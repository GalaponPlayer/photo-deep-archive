import { InternalApiBase } from "@/internal/apis/base-api";
import {
  CreateAccountRequest,
  CreateAccountResponse,
} from "../ports/impls/accounts";

export class AccountApi extends InternalApiBase {
  basePath: string = "/v1/user";
  constructor() {
    super();
    this.setPath(this.basePath);
  }

  async signup(request: CreateAccountRequest): Promise<CreateAccountResponse> {
    const res = await this.post(request, undefined);
    const response = new CreateAccountResponse(res);
    return response;
  }
}
