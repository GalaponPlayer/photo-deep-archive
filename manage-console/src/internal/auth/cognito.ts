import * as AmazonCognitoIdentity from "amazon-cognito-identity-js";

export class CognitoAuthenticator {
  cognitoUserPool: AmazonCognitoIdentity.CognitoUserPool;
  constructor() {
    this.cognitoUserPool = new AmazonCognitoIdentity.CognitoUserPool({
      UserPoolId: import.meta.env.VITE_COGNITO_USER_POOL_ID,
      ClientId: import.meta.env.VITE_COGNITO_APP_CLIENT_ID,
    });
  }

  async signIn(
    email: string,
    password: string
  ): Promise<CognitoAuthenticationResult> {
    const userData = {
      Username: email,
      Pool: this.cognitoUserPool,
    };
    const authenticationData = new AmazonCognitoIdentity.AuthenticationDetails({
      Username: email,
      Password: password,
    });

    const cognitoUser = new AmazonCognitoIdentity.CognitoUser(userData);
    let isError = false;
    let session = null;
    let error = new Error();

    await cognitoUser.authenticateUser(authenticationData, {
      onSuccess: (res) => {
        isError = false;
        session = res;
        error = new Error();
      },
      onFailure: (err) => {
        isError = true;
        error = err;
      },
    });

    return new CognitoAuthenticationResult(isError, session, error);
  }

  async isLogin(): Promise<boolean> {
    const cognitoUser = this.cognitoUserPool.getCurrentUser();
    let isLogin = false;
    if (cognitoUser) {
      cognitoUser.getSession(
        (
          err: Error | null,
          session: AmazonCognitoIdentity.CognitoUserSession | null
        ) => {
          if (err) {
            console.log(err);
            isLogin = false;
          }
          if (session) {
            if (session.isValid()) {
              isLogin = true;
            }
          } else {
            isLogin = false;
          }
        }
      );
    } else {
      isLogin = false;
    }

    return isLogin;
  }

  async getAccessJWT(): Promise<string> {
    const cognitoUser = this.cognitoUserPool.getCurrentUser();
    let accessToken = "";
    if (cognitoUser) {
      cognitoUser.getSession(
        (
          err: Error | null,
          session: AmazonCognitoIdentity.CognitoUserSession | null
        ) => {
          if (err) {
            console.log(err);
            accessToken = "";
          }
          if (session) {
            accessToken = session.getAccessToken().getJwtToken();
          } else {
            accessToken = "";
          }
        }
      );
    } else {
      accessToken = "";
    }

    return accessToken;
  }
}

export class CognitoAuthenticationResult {
  isError: boolean;
  session: AmazonCognitoIdentity.CognitoUserSession | null;
  error: Error;

  constructor(
    isError: boolean,
    session: AmazonCognitoIdentity.CognitoUserSession | null,
    error: Error
  ) {
    this.isError = isError;
    this.session = session;
    this.error = error;
  }

  isAuthError(): boolean {
    return this.isError;
  }

  getError(): Error {
    return this.error;
  }

  getAccessToken(): string {
    if (this.session) {
      return this.session.getAccessToken().getJwtToken();
    } else {
      return "";
    }
  }

  getIdToken(): string {
    if (this.session) {
      return this.session.getIdToken().getJwtToken();
    } else {
      return "";
    }
  }

  getRefreshToken(): string {
    if (this.session) {
      return this.session.getRefreshToken().getToken();
    } else {
      return "";
    }
  }
}
