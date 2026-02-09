class ErrorResponseDto {
  final bool success;
  final String message;
  final String errorCode;
  final int statusCode;

  ErrorResponseDto({
    required this.success,
    required this.message,
    required this.errorCode,
    required this.statusCode
  });

  // JSON -> Dart
  factory ErrorResponseDto.fromJson(Map<String, dynamic> json) {
    return ErrorResponseDto(
      success: json['success'] ?? false,
      message: json['message'] ?? 'Unknown error',
      errorCode: json['error_code'],
      statusCode: json['status_code'] ?? 500,
    );
  }

  // Dart -> JSON
  Map<String, dynamic> toJson() {
    return {
      'success': success,
      'message': message,
      'error_code': errorCode,
      'status_code': statusCode,
    };
  }
}