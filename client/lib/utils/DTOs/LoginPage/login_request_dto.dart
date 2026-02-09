class LoginRequestDto {
  final String username;
  final String password;

  LoginRequestDto({
    required this.username,
    required this.password,
  });

  // Convert DTO to JSON
  Map<String, dynamic> toJson() => {
    'username': this.username,
    'password': this.password
  };
}