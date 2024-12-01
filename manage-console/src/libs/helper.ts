export class Helper {
  static getRegExpEmail(): RegExp {
    return />[\w\-._]+@[\w\-._]+\.[A-Za-z]+/;
  }

  static getRegExpPassword(): RegExp {
    /*
    Contains at least 1 number
    Contains at least 1 special character
    Contains at least 1 uppercase letter
    Contains at least 1 lowercase letter
    */
    return /^(?=.*[0-9])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?])(?=.*[A-Z])(?=.*[a-z]).{8,}$/;
  }
}
