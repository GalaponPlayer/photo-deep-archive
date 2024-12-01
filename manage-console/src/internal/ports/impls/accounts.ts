import { BasePostRequestAbs } from "../interfaces/base";
import { BaseApiResponse } from "./base";

export class CreateAccountRequest implements BasePostRequestAbs {
  name: string;
  email: string;
  password: string;

  constructor(req: { name: string; email: string; password: string }) {
    this.name = req.name;
    this.email = req.email;
    this.password = req.password;
  }

  getParamsObject(): object {
    return {
      name: this.name,
      email: this.email,
      password: this.password,
    };
  }
}

export class CreateAccountResponse extends BaseApiResponse {
  constructor(res: BaseApiResponse) {
    super(
      res.getData(),
      res.getStatus(),
      res.getStatusText(),
      res.getHeaders(),
      res.getConfig(),
      res.getRequest()
    );
  }

  getData(): object {
    return super.getData();
  }
}
