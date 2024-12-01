import { BasePostRequestAbs } from "../interfaces/base";
import { BaseApiResponse } from "./base";

export class CreateAccountRequest implements BasePostRequestAbs {
  name: string;
  email: string;
  password: string;

  constructor(name: string, email: string, password: string) {
    this.name = name;
    this.email = email;
    this.password = password;
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
