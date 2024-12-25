export class Helper {
  static getRegExpEmail(): RegExp {
    return /^[\w\-._]+@[\w\-._]+\.[A-Za-z]{2,}$/;
  }

  static getRegExpPassword(): RegExp {
    return new RegExp(
      [
        "(?=.*[0-9])", // 少なくとも1つの数字
        "(?=.*[!@#$%^&*()_+\\-=\\[\\]{};':\"\\\\|,.<>/?])", // 少なくとも1つの特殊文字
        "(?=.*[A-Z])", // 少なくとも1つの大文字
        "(?=.*[a-z])", // 少なくとも1つの小文字
        ".{10,}", // 10文字以上
      ].join("")
    );
  }
}
