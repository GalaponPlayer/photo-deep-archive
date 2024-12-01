import { CognitoAuthenticator } from "./cognito";

export class UserAuthenticator {
  cognitoAuthenticator: CognitoAuthenticator;
  constructor() {
    this.cognitoAuthenticator = new CognitoAuthenticator();
  }

  async signIn(email: string, password: string): Promise<SignInResult> {
    const res = await this.cognitoAuthenticator.signIn(email, password);
    if (res.isAuthError()) {
      return new SignInResult(true, new JWTtokens("", "", ""), res.getError());
    } else {
      const tokens = new JWTtokens(
        res.getAccessToken(),
        res.getIdToken(),
        res.getRefreshToken()
      );
      return new SignInResult(false, tokens, new Error());
    }
  }

  async isLogin(): Promise<boolean> {
    return this.cognitoAuthenticator.isLogin();
  }

  async getAccessJWT(): Promise<string> {
    return this.cognitoAuthenticator.getAccessJWT();
  }
}

export class SignInResult {
  isError: boolean;
  tokens: JWTtokens;
  error: Error;
  constructor(isError: boolean, tokens: JWTtokens, error: Error) {
    this.isError = isError;
    this.tokens = tokens;
    this.error = error;
  }
}

export class JWTtokens {
  accessToken: string;
  idToken: string;
  refreshToken: string;
  constructor(accessToken: string, idToken: string, refreshToken: string) {
    this.accessToken = accessToken;
    this.idToken = idToken;
    this.refreshToken = refreshToken;
  }

  getAccessToken(): string {
    return this.accessToken;
  }

  getIdToken(): string {
    return this.idToken;
  }

  getRefreshToken(): string {
    return this.refreshToken;
  }
}
