class LoginResponseDto {
  final String uuid;
  final String username;
  final DateTime expiredAt;

  LoginResponseDto({
    required this.uuid,
    required this.username,
    required this.expiredAt
  });

  // Convert JSON to Dart object
  factory LoginResponseDto.fromJson(Map<String, dynamic> json) {
    final user = json['data']['user'];

    return LoginResponseDto(
      uuid: user['uuid'],
      username: user['username'],
      expiredAt: DateTime.parse(user['expired_at']),
    );
  }
}